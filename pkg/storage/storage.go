// Пакет storage
// Реализует работу с БД Postgress
//
// Проект Task Manager
// Автор: Егор Логинов (GO-11) по заданию SkillFactory начиная с модуля 30.8

package storage

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
)

// Хранилище данных
type Storage struct {
	db *pgxpool.Pool
}

// Конструктор объекта БД
func New(p string) (*Storage, error) {
	db, err := pgxpool.Connect(context.Background(), p)
	if err != nil {
		return nil, err
	}

	s := Storage{
		db: db,
	}
	return &s, nil
}
