package storage

import typesutils "github.com/mahin19/students-api/internal/typesUtils"

type Storage interface {
	CreateStudent(name string, email string, age int) (int64, error)
	GetStudentById(id int64) (typesutils.Student, error)
	GetAllStudent() ([]typesutils.Student, error)
	UpdateStudent(id int64,name string, email string, age int) (int64, error)
}
