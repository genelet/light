package light

import (
	"github.com/genelet/sqlproto/xast"
	"github.com/genelet/sqlproto/xlight"
)

func XInsertTo(stmt *xlight.InsertStmt) *xast.InsertStmt {
	output := &xast.InsertStmt{
		Insert: xposTo(),
		TableName: xobjectnameTo(stmt.TableName),
		Source: xinsertSourceTo(stmt.Source)}
	for _, column := range stmt.Columns {
		output.Columns = append(output.Columns, xidentTo(column))
	}
	for _, assignment := range stmt.UpdateAssignments {
		output.UpdateAssignments = append(output.UpdateAssignments, xassignmentTo(assignment))
	}
	return output
}

func InsertTo(stmt *xast.InsertStmt) *xlight.InsertStmt {
	output := &xlight.InsertStmt{
		TableName: objectnameTo(stmt.TableName),
		Source: insertSourceTo(stmt.Source)}
	for _, column := range stmt.Columns {
		output.Columns = append(output.Columns, identTo(column))
	}
	for _, assignment := range stmt.UpdateAssignments {
		output.UpdateAssignments = append(output.UpdateAssignments, assignmentTo(assignment))
	}
	return output
}

func xconstructorSourceTo(item *xlight.ConstructorSource) *xast.ConstructorSource {
	output := &xast.ConstructorSource{
		Values: xposTo()}
	for _, row := range item.Rows {
		output.Rows = append(output.Rows, xrowValueExprTo(row))
	}
	return output
}

func constructorSourceTo(item *xast.ConstructorSource) *xlight.ConstructorSource {
	output := &xlight.ConstructorSource{}
	for _, row := range item.Rows {
		output.Rows = append(output.Rows, rowValueExprTo(row))
	}
	return output
}

func xrowValueExprTo(item *xlight.RowValueExpr) *xast.RowValueExpr {
	output := &xast.RowValueExpr{
		LParen: xposTo(),
		RParen: xposTo(item)}
	for _, value := range item.Values {
		output.Values = append(output.Values, xvalueNodeTo(value))
	}
	return output
}

func rowValueExprTo(item *xast.RowValueExpr) *xlight.RowValueExpr {
	output := &xlight.RowValueExpr{}
	for _, value := range item.Values {
		output.Values = append(output.Values, valueNodeTo(value))
	}
	return output
}

func xassignmentTo(item *xlight.Assignment) *xast.Assignment {
	return &xast.Assignment{
		ID: xidentTo(item.ID),
		Value: xvalueNodeTo(item.Value)}
}

func assignmentTo(item *xast.Assignment) *xlight.Assignment {
	return &xlight.Assignment{
		ID: identTo(item.ID),
		Value: valueNodeTo(item.Value)}
}

func xsubQuerySource(item *xlight.QueryStmt) *xast.SubQuerySource {
	if item == nil { return nil }

	
	return &xast.SubQuerySource{
		SubQuery: XQueryTo(item)}
}

func subQuerySource(item *xast.SubQuerySource) *xlight.QueryStmt {
	if item == nil { return nil }

	return QueryTo(item.SubQuery)
}


