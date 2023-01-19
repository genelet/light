package ast

import (
	"fmt"
	"github.com/genelet/sqlproto/xast"
	"github.com/akito0107/xsqlparser/sqlast"
)

func XCreateViewTo(stmt *sqlast.CreateViewStmt) (*xast.CreateViewStmt, error) {
	query, err := XQueryTo(stmt.Query)
	return &xast.CreateViewStmt{
		Create: xposTo(stmt.Create),
		Name: xobjectnameTo(stmt.Name),
		Query: query,
		Materialized: stmt.Materialized}, err
}

func CreateViewTo(stmt *xast.CreateViewStmt) *sqlast.CreateViewStmt {
	return &xast.CreateViewStmt{
		Create: posTo(stmt.Create),
		Name: objectnameTo(stmt.Name),
		Query: QueryTo(stmt.Query),
		Materialized: stmt.Materialized}
}
