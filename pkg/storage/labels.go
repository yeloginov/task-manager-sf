// Пакет storage
// Реализует работу с БД Postgress
//
// labels.go: Структура и CRUD-операции для Меток
//
// Проект Task Manager
// Автор: Егор Логинов (GO-11) по заданию SkillFactory начиная с модуля 30.8

package storage

import "context"

// Структура Label
type Label struct {
	ID   int    `json:"id"` // (!) pgx не поддерживает SructScan() используя db:"id"
	Name string `json:"name"`
}

// GetLabel возвращает метку по id.
func (s *Storage) GetLabel(id int) (Label, error) {
	sql := `
		SELECT id, name
		FROM labels
		WHERE id = $1;
	`
	row := s.pool.QueryRow(context.Background(), sql, id)

	var l Label
	// Сканирование результата запроса в структуру
	err := row.Scan(&l.ID, &l.Name)
	if err != nil {
		return l, err
	}

	return l, nil
}

// GetLabelsByTask возвращает set меток по id задачи.
func (s *Storage) GetLabelsByTask(id int) ([]Label, error) {
	sql := `
		SELECT labels.id AS id, labels.name AS name
		FROM tasks_labels INNER JOIN labels
		ON tasks_labels.label_id = labels.id
		WHERE tasks_labels.task_id = $1
		ORDER BY id;
	`
	rows, err := s.pool.Query(context.Background(), sql, id)
	if err != nil {
		return nil, err
	}

	var lset []Label
	// Сканирование строк результата запроса в структуру
	for rows.Next() {
		var l Label
		err = rows.Scan(&l.ID, &l.Name)
		if err != nil {
			return nil, err
		}
		lset = append(lset, l)
	}

	// TODO: из комментария Автора модуля - разобраться с rows.Err() (как может проявитсья)
	return lset, rows.Err()
}

// NewLabel создаёт новую метку и возвращает её id.
func (s *Storage) NewLabel(l Label) (int, error) {
	var id int
	sql := `INSERT INTO labels (name) VALUES ($1) RETURNING id;`
	err := s.pool.QueryRow(context.Background(), sql, l.Name).Scan(&id)
	if err != nil {
		// Возвращаем id имеющейся метки
		sql := `SELECT id FROM labels WHERE name = $1;`
		row := s.pool.QueryRow(context.Background(), sql, l.Name)

		var ll Label
		// Сканирование результата запроса в структуру
		err := row.Scan(&ll.ID)
		if err != nil {
			return 0, err
		}

		id = ll.ID
	}

	return id, nil
}
