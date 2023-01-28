package light

import (
	"github.com/genelet/sqlproto/xast"
	"github.com/genelet/sqlproto/xlight"
)

func XCreateIndexTo(stmt *xlight.CreateIndexStmt) *xast.CreateIndexStmt {
	output := &xast.CreateIndexStmt{
		Create: xposTo(),
		TableName: xobjectnameTo(stmt.TableName),
		IsUnique: stmt.IsUnique,
		IndexName: xidentTo(stmt.IndexName),
		MethodName: xidentTo(stmt.MethodName),
		RParen: xposTo(stmt),
		Selection: xwhereNodeTo(stmt.Selection)}
	for _, cname := range stmt.ColumnNames {
		output.ColumnNames = append(output.ColumnNames, xidentTo(cname))
	}
	return output
}

func CreateIndexTo(stmt *xast.CreateIndexStmt) *xlight.CreateIndexStmt {
	output := &xlight.CreateIndexStmt{
		TableName: objectnameTo(stmt.TableName),
		IsUnique: stmt.IsUnique,
		IndexName: identTo(stmt.IndexName),
		MethodName: identTo(stmt.MethodName),
		Selection: whereNodeTo(stmt.Selection)}
	for _, cname := range stmt.ColumnNames {
		output.ColumnNames = append(output.ColumnNames, identTo(cname))
	}
	return output
}

func XDropIndexTo(stmt *xlight.DropIndexStmt) *xast.DropIndexStmt {
    output := &xast.DropIndexStmt{
        Drop: xposTo()}
    for _, name := range stmt.IndexNames {
        output.IndexNames = append(output.IndexNames, xidentTo(name))
    }
    return output
}

func DropIndexTo(stmt *xast.DropIndexStmt) *xlight.DropIndexStmt {
    output := &xlight.DropIndexStmt{}
    for _, name := range stmt.IndexNames {
        output.IndexNames = append(output.IndexNames, identTo(name))
    }
    return output
}
