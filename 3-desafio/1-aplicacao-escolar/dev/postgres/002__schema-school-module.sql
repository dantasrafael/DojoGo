\connect dojo_modulo_escolar;

CREATE TABLE students
(
    id         SERIAL4   NOT NULL,
    name       VARCHAR   NOT NULL,
    email      VARCHAR   NOT NULL,
    birthday   TIMESTAMP NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    CONSTRAINT students_email_un UNIQUE (email),
    CONSTRAINT students_pk PRIMARY KEY (id)
);

CREATE TABLE courses
(
    id         SERIAL4        NOT NULL,
    name       VARCHAR        NOT NULL,
    value      NUMERIC(19, 2) NOT NULL,
    created_at TIMESTAMP      NOT NULL DEFAULT NOW(),
    CONSTRAINT courses_pk PRIMARY KEY (id)
);

CREATE TABLE enrollments
(
    id           SERIAL4   NOT NULL,
    student_id   INT4      NOT NULL,
    course_id    INT4      NOT NULL,
    installments INT2      NOT NULL,
    status       VARCHAR   NOT NULL,
    created_at   TIMESTAMP NOT NULL DEFAULT NOW(),
    CONSTRAINT enrollments_pk PRIMARY KEY (id),
    CONSTRAINT enrollments_students_fk FOREIGN KEY (student_id) REFERENCES students (id) ON DELETE CASCADE ON UPDATE CASCADE,
    CONSTRAINT enrollments_courses_fk FOREIGN KEY (course_id) REFERENCES courses (id) ON DELETE CASCADE ON UPDATE CASCADE
);
