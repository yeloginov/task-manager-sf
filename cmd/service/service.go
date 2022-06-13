// Проект Task Manager
// Автор: Егор Логинов (GO-11) по заданию SkillFactory начиная с модуля 30.8
//
// Основной исполняемый файл сервиса

package main

import (
	"fmt"
	"log"
	"taskmanager/pkg/storage"
)

// Параметры подключения к БД
const (
	DBHost     = "89.223.121.125"
	DBPort     = "5432"
	DBName     = "taskmanager"
	DBUser     = "tm_external"
	DBPassword = "sdf23lLp39n"
)

func main() {

	// Подключение к БД (создание абстрактного объекта БД)
	db, err := storage.New(fmt.Sprintf("postgres://%s:%s@%s:%s/%s", DBUser, DBPassword, DBHost, DBPort, DBName))
	if err != nil {
		log.Fatal(err)
	}

	// Инициация – создание таблиц
	err = db.CreateTables("././pkg/storage/schema.sql")
	if err != nil {
		log.Fatal(err)
	}

	// Заполнение БД тестовыми данными
	u1 := storage.User{Name: "Егор"}
	u1.ID, _ = db.NewUser(u1)
	fmt.Println(db.GetUser(u1.ID))
}
