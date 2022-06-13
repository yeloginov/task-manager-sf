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

	// -----------------------------------------
	// Заполнение БД тестовыми данными
	// Пользователи
	u1 := storage.User{Name: "Егор"}
	u1.ID, _ = db.NewUser(u1)
	//fmt.Println(db.GetUser(u1.ID))
	u2 := storage.User{Name: "Андрей"}
	u2.ID, _ = db.NewUser(u2)
	// Метки
	l1 := storage.Label{Name: "Важная"}
	l1.ID, _ = db.NewLabel(l1)
	//l1.ID, _ = db.NewLabel(l1) // контроль уникальности
	l2 := storage.Label{Name: "Срочная"}
	l2.ID, _ = db.NewLabel(l2)
	//fmt.Println(db.GetLabel(l1.ID))
	//fmt.Println(db.GetLabel(l2.ID))
	// -----------------------------------------

	// Задание модуля 30.8
	// 1. Создаем новые задачи
	tt := make([]storage.Task, 0)
	tt = append(tt, storage.Task{Title: "CRUD для Меток", Content: "Создать файл labels.go и реализовать методы", AuthorID: u1.ID})
	tt = append(tt, storage.Task{Title: "Demo-данные", Content: "Продемонстрировать работу интерфейсов для Users и Labels", AuthorID: u2.ID})
	tt = append(tt, storage.Task{Title: "Релиз к заданию 30.8", Content: "Смержить develop в master и присвоить версию по завершении", AuthorID: u1.ID})
	for _, t := range tt {
		t.ID, err = db.NewTask(t)
		if err != nil {
			log.Fatal(err)
		}
	}
	// 2. Получаем список всех задач
	tt, err = db.GetTasks(0, 0)
	if err != nil {
		log.Fatal(err)
	}
	for _, t := range tt {
		fmt.Println(t)
	}
	fmt.Println()
	// Задаем метки задачам
	db.SetLabelToTask(l1.ID, tt[0].ID)
	db.SetLabelToTask(l2.ID, tt[0].ID)
	db.SetLabelToTask(l1.ID, tt[1].ID)
	db.SetLabelToTask(l2.ID, tt[2].ID)

	// 3. Получаем список задач для автора Егор
	tt, err = db.GetTasks(0, 1)
	if err != nil {
		log.Fatal(err)
	}
	for _, t := range tt {
		fmt.Println(t)
	}
	fmt.Println()

	// 4. Получаем список задач по метке с ID = 1
	tt, err = db.GetTaskByLabels(l1.ID)
	if err != nil {
		log.Fatal(err)
	}
	for _, t := range tt {
		fmt.Println(t)
	}
	fmt.Println()

	// 4a. Получаем список меток для задачи с ID = 1
	ll, err := db.GetLabelsByTask(1)
	if err != nil {
		log.Fatal(err)
	}
	for _, l := range ll {
		fmt.Println(l)
	}
	fmt.Println()

	// 5. Обновляем задачу по ID (ID = 1)
	tt, err = db.GetTasks(1, 0)
	if err != nil {
		log.Fatal(err)
	}
	t := tt[0]
	t.Title = t.Title + " (изменил)"
	t.Content = t.Content + " (изменил)"
	t.AuthorID = u2.ID
	err = db.UpdateTask(t)
	if err != nil {
		log.Fatal(err)
	}
	tt, err = db.GetTasks(1, 0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(tt)
	fmt.Println()

	// 6. Удаляем задачу по ID (ID = 2)
	err = db.DeleteTask(2)
	if err != nil {
		log.Fatal(err)
	}
	// Выводим список задач
	tt, err = db.GetTasks(0, 0)
	if err != nil {
		log.Fatal(err)
	}
	for _, t := range tt {
		fmt.Println(t)
	}
}
