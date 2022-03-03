\connect dojo_modulo_financeiro;

CREATE TABLE account
(
    id           uuid PRIMARY KEY   DEFAULT gen_random_uuid(),
    client_id    varchar   NOT NULL,
    course_id    varchar   NOT NULL,
    external_id  varchar   NOT NULL,
    installments INT2      NOT NULL,
    total        decimal   NOT NULL,
    status       varchar   not null,
    created_at   TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE invoice
(
    id          uuid PRIMARY KEY   DEFAULT gen_random_uuid(),
    account_id  uuid      NOT NULL,
    installment int2      NOT NULL,
    due_date    date      not null,
    value       decimal   not null,
    created_at  timestamp not null default now(),
    paid_at     timestamp,
    CONSTRAINT invoice_account_fk FOREIGN KEY (account_id) REFERENCES account (id) ON DELETE CASCADE ON UPDATE CASCADE
)
