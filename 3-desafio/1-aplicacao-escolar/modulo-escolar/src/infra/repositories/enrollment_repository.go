package repositories

import (
	"fmt"
	"log"
	"modulo-escolar/src/core/db"
	"modulo-escolar/src/domain/entities"
	"strings"
)

const (
	SELECT_ALL_ENROLLMENTS = `SELECT e.id, e.installments, e.status, e.created_at, s.id, s.name, c.id, c.name
		FROM enrollments e
		JOIN students s ON e.student_id = s.id
		JOIN courses c ON e.course_id = c.id
		WHERE 1=1 AND ($1 = '' OR (LOWER(s.name) LIKE $1)) AND ($2 = '' OR (LOWER(c.name) LIKE $2))`
	INSERT_ENROLLMENT = `INSERT INTO enrollments(student_id, course_id, installments, status) VALUES($1, $2, $3, $4) RETURNING id, created_at`
	DELETE_ENROLLMENT = `DELETE FROM enrollments WHERE id=$1`
)

func FindAllEnrollments(studentName, courseName *string) ([]entities.Enrollment, error) {
	if studentName != nil {
		*studentName = fmt.Sprintf("%%%s%%", strings.ToLower(*studentName))
	}

	if courseName != nil {
		*courseName = fmt.Sprintf("%%%s%%", strings.ToLower(*courseName))
	}

	db, err := db.Connect()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, errQuery := db.Query(SELECT_ALL_ENROLLMENTS, studentName, courseName)
	if errQuery != nil {
		return nil, err
	}
	defer rows.Close()

	list := []entities.Enrollment{}
	for rows.Next() {
		var record entities.Enrollment
		if errScan := rows.Scan(&record.ID, &record.Installments, &record.Status, &record.CreatedAt,
			&record.Student.ID, &record.Student.Name, &record.Course.ID, &record.Course.Name); errScan != nil {
			log.Println(errScan.Error())
			return nil, err
		}
		list = append(list, record)
	}

	return list, nil
}

func CreateEnrollment(model *entities.Enrollment) error {
	db, err := db.Connect()
	if err != nil {
		return err
	}
	defer db.Close()

	err = db.QueryRow(INSERT_ENROLLMENT, model.Student.ID, model.Course.ID, model.Installments, model.Status).Scan(&model.ID, &model.CreatedAt)
	if err != nil {
		log.Printf("%v", err)
		return err
	}

	return nil
}

func DeleteEnrollment(id *uint64) error {
	db, err := db.Connect()
	if err != nil {
		return err
	}
	defer db.Close()

	stmt, errStt := db.Prepare(DELETE_ENROLLMENT)
	if errStt != nil {
		return err
	}
	defer stmt.Close()

	if _, err = stmt.Exec(id); err != nil {
		return err
	}

	return nil
}
