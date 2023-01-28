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

func TestCreateView(t *testing.T) {
	strs := []string{
	"CREATE VIEW customers_view AS " +
				"SELECT customer_name, contract_name " +
				"FROM customers " +
				"WHERE country = 'Brazil'"}

	for i, str := range strs {
		//if i != 17 { continue }
		parser, err := xsqlparser.NewParser(bytes.NewBufferString(str), &dialect.GenericSQLDialect{})
		if err != nil { t.Fatal(err) }

		istmt, err := parser.ParseStatement()
		if err != nil { t.Fatal(err) }
		stmt := istmt.(*sqlast.CreateViewStmt)
//pp.Println(stmt)

		xcreateView, err := ast.XCreateViewTo(stmt)
		if err != nil { t.Fatal(err) }

		createView := CreateViewTo(xcreateView)
		reverse2 := XCreateViewTo(createView)
		reverse3 := ast.CreateViewTo(reverse2)
//pp.Println(reverse)
		if strings.ToLower(stmt.ToSQLString()) != strings.ToLower(reverse3.ToSQLString()) {
			t.Errorf("%d=>%s", i, stmt.ToSQLString())
			t.Errorf("%d=>%s", i, reverse.ToSQLString())
		}
	}
}
