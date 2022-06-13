// Пакет storage
// Реализует работу с БД Postgress
//
// Проект Task Manager
// Автор: Егор Логинов (GO-11) по заданию SkillFactory начиная с модуля 30.8

package storage

import (
	"context"
	"io/ioutil"

	"github.com/jackc/pgx/v4/pgxpool"
)

// Хранилище данных
type Storage struct {
	pool *pgxpool.Pool
}

// Конструктор объекта БД
func New(p string) (*Storage, error) {
	db, err := pgxpool.Connect(context.Background(), p)
	if err != nil {
		return nil, err
	}

	s := Storage{
		pool: db,
	}
	return &s, nil
}

// Создание таблиц БД на основе файла схемы schm
func (s *Storage) CreateTables(schm string) error {

	// Читаем файл SQL-запроса со схемой БД
	buf, err := ioutil.ReadFile(schm)
	if err != nil {
		return err
	}

	// Выполняем SQL-запрос создания структуры БД
	sql := string(buf)
	_, err = s.pool.Exec(context.Background(), sql)
	if err != nil {
		return err
	}

	return nil
}
