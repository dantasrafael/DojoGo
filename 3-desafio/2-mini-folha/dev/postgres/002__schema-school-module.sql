\connect dojo_modulo_folha;

CREATE TABLE employee
(
    id          SERIAL4 NOT NULL,
    name        VARCHAR NOT NULL,
    external_id varchar not null,
    salary      DECIMAL NOT NULL,
    CONSTRAINT employee_pk PRIMARY KEY (id)
);

CREATE TABLE payroll
(
    id          SERIAL4   NOT NULL,
    employee_id VARCHAR   NOT NULL,
    reference   TIMESTAMP NOT NULL,
    salary_base decimal   not null,
    discount    decimal   not null,
    total       decimal   not null,
    created_at  TIMESTAMP NOT NULL DEFAULT NOW(),
    CONSTRAINT payroll_pk PRIMARY KEY (id),
    CONSTRAINT payroll_employee_fk FOREIGN KEY (employee_id) REFERENCES employee (id) ON DELETE CASCADE ON UPDATE CASCADE
);
