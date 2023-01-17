package ast

import (
	"bytes"
	"strings"
	"testing"

//	"github.com/k0kubun/pp/v3"

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
);
`}

	for i, str := range strs {
		//if i != 17 { continue }
		parser, err := xsqlparser.NewParser(bytes.NewBufferString(str), &dialect.GenericSQLDialect{}, xsqlparser.ParseComment())
		if err != nil { t.Fatal(err) }

		file, err := parser.ParseFile()
		if err != nil { t.Fatal(err) }

		createTableStmt := file.Stmts[0].(*sqlast.CreateTableStmt)
//pp.Println(createTableStmt)

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
