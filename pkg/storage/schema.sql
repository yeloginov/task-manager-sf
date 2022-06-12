/* 
    Схема БД

    Проект Task Manager
    Автор: Егор Логинов (GO-11) по заданию SkillFactory начиная с модуля 30.8
*/

DROP TABLE IF EXISTS tasks_labels, tasks, labels, users;

-- Пользователи
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL
);
