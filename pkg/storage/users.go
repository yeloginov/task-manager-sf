// Пакет storage
// Реализует работу с БД Postgress
//
// users.go: Структура и CRUD-операции для Пользоваталей
//
// Проект Task Manager
// Автор: Егор Логинов (GO-11) по заданию SkillFactory начиная с модуля 30.8

package storage

import "context"

// Структура User
type User struct {
	ID   int    `json:"id"` // (!) pgx не поддерживает SructScan() используя db:"id"
	Name string `json:"name"`
}

// GetUser возвращает пользователя по ID
func (s *Storage) GetUser(id int) (User, error) {
	sql := `
		SELECT id, name
		FROM users
		WHERE id = $1;
	`
	row := s.pool.QueryRow(context.Background(), sql, id)

	var u User
	// Сканирование результата запроса в структуру
	err := row.Scan(&u.ID, &u.Name)
	if err != nil {
		return u, err
	}

	return u, nil
}

// NewTask создаёт новую задачу и возвращает её id.
func (s *Storage) NewUser(u User) (int, error) {
	var id int
	sql := `INSERT INTO users (name) 
			VALUES ($1) 
			RETURNING id;`
	err := s.pool.QueryRow(context.Background(), sql, u.Name).Scan(&id)

	return id, err
}
