\connect dojo_modulo_folha;

CREATE TABLE colaborador
(
    id          integer NOT NULL,
    nome        VARCHAR NOT NULL,
    salario      DECIMAL NOT NULL,
    CONSTRAINT colaborador_pk PRIMARY KEY (id)
);

CREATE TABLE folha
(
    id          serial4   NOT NULL,
    colaborador_id integer   NOT NULL,
    colaborador_nome varchar   NOT NULL,
    colaborador_salario numeric(19,2) NOT NULL,
    inss    numeric(19,2)   not null,
    irpf    numeric(19,2)   not null,
    total       numeric(19,2)   not null,
    CONSTRAINT folha_pk PRIMARY KEY (id)
);
