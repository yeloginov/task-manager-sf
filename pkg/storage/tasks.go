// Пакет storage
// Реализует работу с БД Postgress
//
// task.go: Структура и CRUD-операции для Задач
//
// Проект Task Manager
// Автор: Егор Логинов (GO-11) по заданию SkillFactory начиная с модуля 30.8

package storage

import "context"

// Структура Task
type Task struct {
	ID         int    `json:"id"` // (!) pgx не поддерживает SructScan() используя db:"id"
	Created    int64  `json:"created"`
	Closed     int64  `json:"closed"`
	AuthorID   int    `json:"author_id"`
	AssignedID int    `json:"assigned_id"`
	Title      string `json:"title"`
	Content    string `json:"content"`
}

// GetTasks возвращает список задач из БД с фильтром по ID задачи и/или ID Автора
// Если в качестве параметра передан 0, то выборка по всем задачам / авторам
func (s *Storage) GetTasks(id, aId int) ([]Task, error) {
	sql := `
		SELECT id, created, closed, author_id, assigned_id, title, content
		FROM tasks
		WHERE ($1 = 0 OR id = $1) AND ($2 = 0 OR author_id = $2)
		ORDER BY id;
	`
	rows, err := s.pool.Query(context.Background(), sql, id, aId)
	if err != nil {
		return nil, err
	}

	var tset []Task
	// Сканирование строк результата запроса в структуру
	for rows.Next() {
		var t Task
		err = rows.Scan(&t.ID, &t.Created, &t.Closed, &t.AuthorID, &t.AssignedID, &t.Title, &t.Content)
		if err != nil {
			return nil, err
		}
		tset = append(tset, t)
	}

	// TODO: из комментария Автора модуля - разобраться с rows.Err() (как может проявитсья)
	return tset, rows.Err()
}

// NewTask создаёт новую задачу и возвращает её id.
func (s *Storage) NewTask(t Task) (int, error) {
	var id int
	sql := `INSERT INTO tasks (title, content) 
			VALUES ($1, $2) 
			RETURNING id;`
	err := s.pool.QueryRow(context.Background(), sql, t.Title, t.Content).Scan(&id)

	return id, err
}
