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

func TestQuery(t *testing.T) {
	strs := []string{
	"SELECT a from test_table",
	"SELECT * from test_table",
	"SELECT test_table.* from test_table",
	"SELECT y as z from test_table where x > 6 order by x asc, y.z desc",
	"SELECT aa from test_table where bb = 1.414 and cc = 'john'",
	"SELECT COUNT(*) from test_table group by x.a having x.a > 6",
	"SELECT *, aa, bb as s, COUNT(*), count(*) t, count(cc), count(cc) as u from test_table group by x having x > 6",
	"SELECT aa from test_table where c=1 and bb in (SELECT region FROM top_regions)",
	"SELECT orders.product FROM orders LEFT JOIN accounts as acs ON orders.account_id = accounts.id INNER JOIN accounts_type ON accounts_type.type_id = accounts.type_id ",
	"SELECT orders.product as prod, SUM(orders.quantity) AS product_units, accounts.* FROM orders LEFT JOIN accounts ON orders.account_id = accounts.id INNER JOIN accounts_type ON accounts_type.type_id = accounts.type_id WHERE orders.region IN (SELECT region FROM top_regions) ORDER BY product_units ASC LIMIT 100",
	"WITH regional_sales AS (" +
	"SELECT region, SUM(amount) AS total_sales " +
	"FROM orders GROUP BY region) " +
	"SELECT product, SUM(quantity) AS product_units " +
	"FROM orders " +
	"WHERE region IN (SELECT region FROM top_regions) " +
	"GROUP BY region, product",
	"SELECT x FROM y UNION SELECT x FROM z",
	"SELECT name FROM stadium EXCEPT SELECT T2.name FROM concert AS T1 JOIN stadium AS T2 ON T1.stadium_id  =  T2.stadium_id WHERE T1.year  =  2014",
	`SELECT stock_name, -1*price AS capital_gain_loss, 5.0*price AS cool FROM Stocks`,
	`SELECT id, CASE WHEN nb_childrens >0 AND nb_parents >0 THEN "Inner" WHEN nb_childrens > 0 THEN "Root" ELSE "Leaf" END AS type FROM l2`,
	`SELECT stock_name, (CASE WHEN operation = "Buy" THEN -1*price ELSE price END) AS capital_gain_loss FROM Stocks`,
	`WITH l1 AS (SELECT t.id, c.id AS c_id, t.p_id FROM Tree t LEFT JOIN Tree c ON c.p_id = t.id), l2 AS ( SELECT id, COUNT(c_id) AS nb_childrens, COUNT(p_id) AS nb_parents FROM l1 GROUP BY id) SELECT id, CASE WHEN nb_childrens >0 AND nb_parents >0 THEN "Inner" WHEN nb_childrens > 0 THEN "Root" ELSE "Leaf" END AS type FROM l2`,
	`SELECT stock_name, SUM( CASE WHEN operation = "Buy" THEN -1*price ELSE price END) AS  capital_gain_loss FROM Stocks GROUP BY stock_name`,
	"SELECT x FROM a UNION SELECT x FROM b EXCEPT select x FROM c"}

	for i, str := range strs {
		//if i != 17 { continue }
		parser, err := xsqlparser.NewParser(bytes.NewBufferString(str), &dialect.GenericSQLDialect{})
		if err != nil { t.Fatal(err) }

		istmt, err := parser.ParseStatement()
		if err != nil { t.Fatal(err) }
		stmt := istmt.(*sqlast.QueryStmt)
//pp.Println(stmt)

		xquery, err := XQueryTo(stmt)
		if err != nil { t.Fatal(err) }

		reverse := QueryTo(xquery)
//pp.Println(reverse)
		if strings.ToLower(stmt.ToSQLString()) != strings.ToLower(reverse.ToSQLString()) {
			t.Errorf("%d=>%s", i, stmt.ToSQLString())
			t.Errorf("%d=>%s", i, reverse.ToSQLString())
		}
	}
}
