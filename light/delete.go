package light

import (
	"github.com/genelet/sqlproto/xast"
	"github.com/genelet/sqlproto/xlight"
)

func XDeleteTo(stmt *xlight.DeleteStmt) *xast.DeleteStmt {
	if stmt == nil { return nil }

	return &xast.DeleteStmt{
		Delete: xposTo(),
		TableName: xobjectnameTo(stmt.TableName),
		Selection: xwhereNodeTo(stmt.Selection)}
}

func DeleteTo(stmt *xast.DeleteStmt) *xlight.DeleteStmt {
	if stmt == nil { return nil }

	return &xlight.DeleteStmt{
		TableName: objectnameTo(stmt.TableName),
		Selection: whereNodeTo(stmt.Selection)}
}
