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
	SELECT_ALL_COURSES  = `SELECT c.id, c.name, c.value, c.created_at FROM courses c WHERE ($1 = '' OR (LOWER(c.name) LIKE $1))`
	SELECT_COURSE_BY_ID = `SELECT c.id, c.name, c.value, c.created_at FROM courses c WHERE c.id = $1`
	INSERT_COURSE       = `INSERT INTO courses(name, value) VALUES($1, $2) RETURNING id, created_at`
	UPDATE_COURSE       = `UPDATE courses SET name=$1, value=$2 WHERE id=$3`
	DELETE_COURSE       = `DELETE FROM courses WHERE id=$1`
)

func FindAllCourses(name *string) ([]entities.Course, error) {
	if name != nil {
		*name = fmt.Sprintf("%%%s%%", strings.ToLower(*name))
	}

	db, err := db.Connect()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, errQuery := db.Query(SELECT_ALL_COURSES, name)
	if errQuery != nil {
		return nil, err
	}
	defer rows.Close()

	list := []entities.Course{}
	for rows.Next() {
		var record entities.Course
		if errScan := rows.Scan(&record.ID, &record.Name, &record.Value, &record.CreatedAt); errScan != nil {
			log.Println(errScan.Error())
			return nil, err
		}
		list = append(list, record)
	}

	return list, nil
}

func FindCourseById(id *uint64) (*entities.Course, error) {
	db, err := db.Connect()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var record entities.Course
	if err = db.QueryRow(SELECT_COURSE_BY_ID, id).Scan(&record.ID, &record.Name, &record.Value, &record.CreatedAt); err != nil {
		if err == sql.ErrNoRows {
			return &record, nil
		}
		return nil, err
	}

	return &record, nil
}

func CreateCourse(model *entities.Course) error {
	db, err := db.Connect()
	if err != nil {
		return err
	}
	defer db.Close()

	err = db.QueryRow(INSERT_COURSE, model.Name, model.Value).Scan(&model.ID, &model.CreatedAt)
	if err != nil {
		log.Printf("%v", err)
		return err
	}

	return nil
}

func UpdateCourse(id *uint64, model *entities.Course) error {
	db, err := db.Connect()
	if err != nil {
		return err
	}
	defer db.Close()

	stmt, errStt := db.Prepare(UPDATE_COURSE)
	if errStt != nil {
		return err
	}
	defer stmt.Close()

	if _, err = stmt.Exec(model.Name, model.Value, id); err != nil {
		return err
	}

	return nil
}

func DeleteCourse(id *uint64) error {
	db, err := db.Connect()
	if err != nil {
		return err
	}
	defer db.Close()

	stmt, errStt := db.Prepare(DELETE_COURSE)
	if errStt != nil {
		return err
	}
	defer stmt.Close()

	if _, err = stmt.Exec(id); err != nil {
		return err
	}

	return nil
}
