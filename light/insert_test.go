package light

import (
	"bytes"
	"strings"
	"testing"

	"github.com/genelet/sqlproto/ast"

//	"github.com/k0kubun/pp/v3"

	"github.com/akito0107/xsqlparser"
	"github.com/akito0107/xsqlparser/sqlast"
	"github.com/akito0107/xsqlparser/dialect"
)

func TestInsert(t *testing.T) {
	strs := []string{
	"INSERT INTO customers (customer_name, contract_name) VALUES ('Cardinal', 'Tom B. Erichsen')",
	"INSERT INTO customers (customer_name, contract_name) VALUES ('Cardinal', 'Tom B. Erichsen'), ('Cardinal2', 'Tom B. Erichsen2'), ('Cardinal3', 'Tom B. Erichsen3')",
	"INSERT INTO customers (customer_name, contract_name) SELECT * FROM customers2"}

	for i, str := range strs {
		//if i != 17 { continue }
		parser, err := xsqlparser.NewParser(bytes.NewBufferString(str), &dialect.GenericSQLDialect{})
		if err != nil { t.Fatal(err) }

		istmt, err := parser.ParseStatement()
		if err != nil { t.Fatal(err) }
		stmt := istmt.(*sqlast.InsertStmt)
//pp.Println(stmt)

		xinsert, err := ast.XInsertTo(stmt)
		if err != nil { t.Fatal(err) }

		insert := InsertTo(xinsert)
		reverse2 := XInsertTo(insert)
		reverse3 := ast.InsertTo(reverse2)
//pp.Println(reverse)
		if strings.ToLower(stmt.ToSQLString()) != strings.ToLower(reverse3.ToSQLString()) {
			t.Errorf("%d=>%s", i, stmt.ToSQLString())
			t.Errorf("%d=>%s", i, reverse3.ToSQLString())
		}
	}
}
