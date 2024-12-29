package sqlite

import (
	"database/sql"
	"fmt"

	"github.com/mahin19/students-api/internal/config"
	typesutils "github.com/mahin19/students-api/internal/typesUtils"
	_ "github.com/mattn/go-sqlite3"
)

type Sqlite struct {
	Db *sql.DB
}

func New(cfg *config.Config) (*Sqlite, error) {
	db, err := sql.Open("sqlite3", cfg.StoragePath)
	if err != nil {
		return nil, err
	}
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS students(
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT,
	email TEXT,
	age INTEGER
	)`)

	if err != nil {
		return nil, err
	}
	return &Sqlite{
		Db: db,
	}, nil

}
func (s *Sqlite) CreateStudent(name string, email string, age int) (int64, error) {
	stmt, err := s.Db.Prepare("INSERT INTO students (name,email,age) values(?,?,?)")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	result, err := stmt.Exec(name, email, age)
	if err != nil {
		return 0, err
	}
	lastId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return lastId, nil
}
func (s *Sqlite) UpdateStudent(id int64, name string, email string, age int) (int64, error) {
	stmt, err := s.Db.Prepare("UPDATE students set name =? ,email=?,age=?  where id=? ")
	if err != nil {

		return 0, err
	}
	defer stmt.Close()
	result, err := stmt.Exec(name, email, age, id)
	if err != nil {
		return 0, err
	}
	effectedId, err := result.RowsAffected()
	fmt.Print(effectedId, id)
	if err != nil {
		return 0, err
	}
	return effectedId, nil
}
func (s *Sqlite) GetStudentById(id int64) (typesutils.Student, error) {
	stmt, err := s.Db.Prepare("select id as id , name as name,email as email,age as age from students s where s.id=? limit 1")
	if err != nil {
		return typesutils.Student{}, err
	}
	defer stmt.Close()
	var student typesutils.Student
	err = stmt.QueryRow(id).Scan(&student.Id, &student.Name, &student.Email, &student.Age)
	if err != nil {
		if err == sql.ErrNoRows {
			return typesutils.Student{}, fmt.Errorf("No student found with id %s", fmt.Sprint(id))
		}
		return typesutils.Student{}, err
	}
	return student, nil
}

func (s *Sqlite) GetAllStudent() ([]typesutils.Student, error) {
	stmt, err := s.Db.Prepare("select id as id , name as name,email as email,age as age from students s")
	if err != nil {
		return []typesutils.Student{}, err
	}
	defer stmt.Close()
	var students []typesutils.Student
	rows, err := stmt.Query()
	if err != nil {
		if err == sql.ErrNoRows {
			return []typesutils.Student{}, fmt.Errorf("No student found with")
		}
		return []typesutils.Student{}, err
	}
	for rows.Next() {
		var student typesutils.Student
		err := rows.Scan(&student.Id, &student.Name, &student.Email, &student.Age)
		if err != nil {
			return nil, err
		}
		students = append(students, student)
	}
	return students, nil
}
