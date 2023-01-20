package ast

import (
	"github.com/genelet/sqlproto/xast"
	"github.com/akito0107/xsqlparser/sqlast"
)

func XCreateIndexTo(stmt *sqlast.CreateIndexStmt) (*xast.CreateIndexStmt, error) {
	selection, err := xwhereNodeTo(stmt.Selection)
	if err != nil { return nil, err }

	output := &xast.CreateIndexStmt{
		Create: xposTo(stmt.Create),
		TableName: xobjectnameTo(stmt.TableName),
		IsUnique: stmt.IsUnique,
		IndexName: xidentTo(stmt.IndexName),
		MethodName: xidentTo(stmt.MethodName),
		RParen: xposTo(stmt.RParen),
		Selection: selection}
	for _, cname := range stmt.ColumnNames {
		output.ColumnNames = append(output.ColumnNames, xidentTo(cname))
	}

	return output, nil
}

func CreateIndexTo(stmt *xast.CreateIndexStmt) *sqlast.CreateIndexStmt {
	output := &sqlast.CreateIndexStmt{
		Create: posTo(stmt.Create),
		RParen: posTo(stmt.RParen),
		TableName: objectnameTo(stmt.TableName),
		IsUnique: stmt.IsUnique,
		Selection: whereNodeTo(stmt.Selection)}
	if iname := identTo(stmt.IndexName); iname != nil {
		output.IndexName = iname.(*sqlast.Ident)
	}
	if mname := identTo(stmt.MethodName); mname != nil {
		output.MethodName = mname.(*sqlast.Ident)
	}
	for _, cname := range stmt.ColumnNames {
		output.ColumnNames = append(output.ColumnNames, identTo(cname).(*sqlast.Ident))
	}

	return output
}

func XDropIndexTo(stmt *sqlast.DropIndexStmt) *xast.DropIndexStmt {
    output := &xast.DropIndexStmt{
        Drop: xposTo(stmt.Drop)}
    for _, name := range stmt.IndexNames {
        output.IndexNames = append(output.IndexNames, xidentTo(name))
    }
    return output
}

func DropIndexTo(stmt *xast.DropIndexStmt) *sqlast.DropIndexStmt {
    output := &sqlast.DropIndexStmt{
        Drop: posTo(stmt.Drop)}
    for _, name := range stmt.IndexNames {
        output.IndexNames = append(output.IndexNames, identTo(name).(*sqlast.Ident))
    }
    return output
}
