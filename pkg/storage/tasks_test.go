// Пакет storage
// Реализует работу с БД Postgress
//
// task.go: Структура и CRUD-операции для Task
//
// Проект Task Manager
// Автор: Егор Логинов (GO-11) по заданию SkillFactory начиная с модуля 30.8

package storage

import (
	"testing"

	"github.com/jackc/pgx/v4/pgxpool"
)

func TestStorage_NewTask(t *testing.T) {
	type fields struct {
		pool *pgxpool.Pool
	}
	type args struct {
		t Task
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		{name: "Test 1", args: args{
			Task{Title: "Выполнить тестирование", Content: "Разобраться с тестовыми данными для Unit test", AuthorID: 1}},
			want: 0, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Storage{
				pool: tt.fields.pool,
			}
			got, err := s.NewTask(tt.args.t)
			if (err != nil) != tt.wantErr {
				t.Errorf("Storage.NewTask() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Storage.NewTask() = %v, want %v", got, tt.want)
			}
		})
	}
}
