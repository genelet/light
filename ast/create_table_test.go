package ast

import (
	"bytes"
	"strings"
	"testing"

	"github.com/k0kubun/pp/v3"

	"github.com/akito0107/xsqlparser"
	"github.com/akito0107/xsqlparser/sqlast"
	"github.com/akito0107/xsqlparser/dialect"
)

func TestCreateTable(t *testing.T) {
	strs := []string{
	`CREATE TABLE test (
    col0 int primary key,
    col1 integer constraint test_constraint check (10 < col1 and col1 < 100),
    col3 varchar(255) default 'abc',
    col4 integer default 100,
    foreign key (col0, col1) references test2(col1, col2),
    CONSTRAINT test_constraint check(col1 > 10)
);`,
	`CREATE TABLE def_state (
  state_id smallint NOT NULL AUTO_INCREMENT,
  country_id smallint NOT NULL,
  state_code char(2),
  state_name varchar(255),
  english_name varchar(255),
  PRIMARY KEY (state_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;`,
	`CREATE TABLE persons (
 person_id UUID PRIMARY KEY NOT NULL,
 first_name varchar(255) UNIQUE,
 last_name character varying(255) NOT NULL,
 created_at timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL);`,
	`CREATE TABLE persons (
person_id int PRIMARY KEY NOT NULL,
last_name character varying(255) NOT NULL,
test_id int NOT NULL REFERENCES test(id1),
email character varying(255) UNIQUE NOT NULL,
age int NOT NULL CHECK(age > 0 AND age < 100),
created_at timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL
);`,
	`CREATE TABLE persons (
person_id int,
CONSTRAINT production UNIQUE(test_column),
PRIMARY KEY(person_id),
CHECK(id > 100),
FOREIGN KEY(test_id) REFERENCES other_table(col1, col2)
);`}

	for i, str := range strs {
		//if i != 1 { continue }
		parser, err := xsqlparser.NewParser(bytes.NewBufferString(str), &dialect.GenericSQLDialect{}, xsqlparser.ParseComment())
		if err != nil { t.Fatal(err) }

		file, err := parser.ParseFile()
		if err != nil { t.Fatal(err) }

		createTableStmt := file.Stmts[0].(*sqlast.CreateTableStmt)
pp.Println(createTableStmt)

		createTable, err := XCreateTableTo(createTableStmt)
		if err != nil { t.Fatal(err) }

		reverse := CreateTableTo(createTable)
//pp.Println(reverse)
		if strings.ToLower(createTableStmt.ToSQLString()) != strings.ToLower(reverse.ToSQLString()) {
			t.Errorf("%d=>%s", i, createTableStmt.ToSQLString())
			t.Errorf("%d=>%s", i, reverse.ToSQLString())
		}
	}
}
