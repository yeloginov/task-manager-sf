/* 
    Схема БД

    Проект Task Manager
    Автор: Егор Логинов (GO-11) по заданию SkillFactory начиная с модуля 30.8
*/

DROP TABLE IF EXISTS tasks_labels, tasks, labels, users;
GRANT usage ON SCHEMA public TO public;

-- Пользователи
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL
);

-- Метки задач
CREATE TABLE labels (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL
);

-- Задачи
CREATE TABLE tasks (
    id SERIAL PRIMARY KEY,
    created BIGINT NOT NULL DEFAULT extract(epoch from now()),
    closed BIGINT DEFAULT 0,
    author_id INTEGER REFERENCES users(id) DEFAULT 0,
    assigned_id INTEGER REFERENCES users(id) DEFAULT 0,
    title TEXT,
    content TEXT
);

-- Связь между задачами и метками
CREATE TABLE tasks_labels (
    task_id INTEGER REFERENCES tasks(id),
    label_id INTEGER REFERENCES labels(id)
);