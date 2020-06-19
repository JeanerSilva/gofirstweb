package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

//GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO usuario;

/*

create user usuario;
create database alura_loja;
\c alura_loja;
create table produtos (
	id serial primary key,
	nome varchar,
	descricao varchar,
	preco decimal,
	quantidade integer
);

alter table produtos add column xxx varchar;

insert into produtos (nome, descricao, preco, quantidade) values ('nome1', 'descricao1', 1, 1)

*/

func ConectaComBancoDeDados() *sql.DB {
	//conexao := "user=usuario dbname=alura_loja password=postgres host=localhost sslmode=disable"
	conexao := "user=usuario dbname=alura_loja password=postgres host=104.41.56.220 sslmode=disable"
	db, err := sql.Open("postgres", conexao)
	if err != nil {
		panic(err.Error())
	}
	return db
}
