package repositories

import (
	"database/sql"
	"fmt"
	"log"
	"modulo-escolar/src/core/db"
	"modulo-escolar/src/domain/entities"
	"strings"
)

const (
	SELECT_ALL_STUDENTS  = `SELECT s.id, s.name, s.email, s.birthday, s.created_at FROM students s WHERE ($1 = '' OR (LOWER(s.name) LIKE $1))`
	SELECT_STUDENT_BY_ID = `SELECT s.id, s.name, s.email, s.birthday, s.created_at FROM students s WHERE s.id = $1`
	INSERT_STUDENT       = `INSERT INTO students(name, email, birthday) VALUES($1, $2, $3) RETURNING id, created_at`
	UPDATE_STUDENT       = `UPDATE students SET name=$1 WHERE id=$2`
	DELETE_STUDENT       = `DELETE FROM students WHERE id=$1`
)

func FindAllStudents(name *string) ([]entities.Student, error) {
	if name != nil {
		*name = fmt.Sprintf("%%%s%%", strings.ToLower(*name))
	}

	db, err := db.Connect()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, errQuery := db.Query(SELECT_ALL_STUDENTS, name)
	if errQuery != nil {
		return nil, err
	}
	defer rows.Close()

	list := []entities.Student{}
	for rows.Next() {
		var record entities.Student
		if errScan := rows.Scan(&record.ID, &record.Name, &record.Email, &record.Birthday, &record.CreatedAt); errScan != nil {
			log.Println(errScan.Error())
			return nil, err
		}
		list = append(list, record)
	}

	return list, nil
}

func FindStudentById(id *uint64) (*entities.Student, error) {
	db, err := db.Connect()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var record entities.Student
	if err = db.QueryRow(SELECT_STUDENT_BY_ID, id).Scan(&record.ID, &record.Name, &record.Email, &record.Birthday, &record.CreatedAt); err != nil {
		if err == sql.ErrNoRows {
			return &record, nil
		}
		return nil, err
	}

	return &record, nil
}

func CreateStudent(model *entities.Student) error {
	db, err := db.Connect()
	if err != nil {
		return err
	}
	defer db.Close()

	err = db.QueryRow(INSERT_STUDENT, model.Name, model.Email, model.Birthday).Scan(&model.ID, &model.CreatedAt)
	if err != nil {
		log.Printf("%v", err)
		return err
	}

	return nil
}

func UpdateStudent(id *uint64, model *entities.Student) error {
	db, err := db.Connect()
	if err != nil {
		return err
	}
	defer db.Close()

	stmt, errStt := db.Prepare(UPDATE_STUDENT)
	if errStt != nil {
		return err
	}
	defer stmt.Close()

	if _, err = stmt.Exec(model.Name, id); err != nil {
		return err
	}

	return nil
}

func DeleteStudent(id *uint64) error {
	db, err := db.Connect()
	if err != nil {
		return err
	}
	defer db.Close()

	stmt, errStt := db.Prepare(DELETE_STUDENT)
	if errStt != nil {
		return err
	}
	defer stmt.Close()

	if _, err = stmt.Exec(id); err != nil {
		return err
	}

	return nil
}
