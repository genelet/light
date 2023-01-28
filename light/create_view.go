package light

import (
	"github.com/genelet/sqlproto/xast"
	"github.com/genelet/sqlproto/xlight"
)

func XCreateViewTo(stmt *xlight.CreateViewStmt) *xast.CreateViewStmt {
	return &xast.CreateViewStmt{
		Create: xposTo(),
		Name: xobjectnameTo(stmt.Name),
		Query: XQueryTo(stmt.Query),
		Materialized: stmt.Materialized}
}

func CreateViewTo(stmt *xast.CreateViewStmt) *xlight.CreateViewStmt {
	return &xlight.CreateViewStmt{
		Name: objectnameTo(stmt.Name),
		Query: QueryTo(stmt.Query),
		Materialized: stmt.Materialized}
}
