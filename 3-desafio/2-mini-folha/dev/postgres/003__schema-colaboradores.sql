\connect dojo_modulo_colaborador;

create table colaboradores
(
    id      serial4        not null primary key,
    nome    varchar(100)   not null,
    funcao  varchar(100)   not null,
    salario numeric(19, 2) not null
);