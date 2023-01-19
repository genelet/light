package ast

import (
	"github.com/genelet/sqlproto/xast"
	"github.com/akito0107/xsqlparser/sqlast"
)

func XDeleteTo(stmt *sqlast.DeleteStmt) (*xast.DeleteStmt, error) {
	selection, err := xwhereStmtTo(stmt.Selection)
	if err != nil { return nil, err }

	return &xast.DeleteStmt{
		Delete: xposTo(stmt.Delete),
		TableName: xobjectnameTo(stmt.TableName),
		Selection: selection}, nil
}

func DeleteTo(stmt *xast.DeleteStmt) *sqlast.DeleteStmt {
	return &sqlast.DeleteStmt{
		Delete: posTo(stmt.Delete),
		TableName: objectnameTo(stmt.TableName),
		Selection: whereStmtTo(stmt.Selection)}
}
