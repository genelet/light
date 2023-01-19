package ast

import (
	"github.com/genelet/sqlproto/xast"
	"github.com/akito0107/xsqlparser/sqlast"
)

func XInsertTo(stmt *sqlast.InsertStmt) (*xast.InsertStmt, error) {
	source, err := xsourceStmtTo(stmt.Source)
	if err != nil { return nil, err }

	output := &xast.InsertStmt{
		Insert: xposTo(stmt.Insert),
		TableName: xobjectnameTo(stmt.TableName),
		Source: source}
	for _, column := range stmt.Columns {
		output.Columns = append(output.Columns, xidentTo(column))
	}
	for _, assignment := range stmt.UpdateAssignments {
		x, err := xassignmentTo(assignment)
		if err != nil { return nil, err }	
		output.UpdateAssignments = append(output.UpdateAssignments, x)
	}
	return output, nil
}

func InsertTo(stmt *xast.InsertStmt) *sqlast.InsertStmt {
	output := &sqlast.InsertStmt{
		Insert: posTo(stmt.Insert),
		TableName: objectnameTo(stmt.TableName),
		Source: sourceStmtTo(stmt.Source)}
	for _, column := range stmt.Columns {
		output.Columns = append(output.Columns, identTo(column).(*sqlast.Ident))
	}
	for _, assignment := range stmt.UpdateAssignments {
		output.UpdateAssignments = append(output.UpdateAssignments, assignmentTo(assignment))
	}
	return output
}

func xconstructorSourceTo(item *sqlast.ConstructorSource) (*xast.ConstructorSource, error) {
	output := &xast.ConstructorSource{
		Values: xposTo(item.Values)}
	for _, row := range item.Rows {
		x, err := xrowValueExprTo(row)
		if err != nil { return nil, err }
		output.Rows = append(output.Rows, x)
	}
	return output, nil
}

func constructorSourceTo(item *xast.ConstructorSource) *sqlast.ConstructorSource {
	output := &sqlast.ConstructorSource{
		Values: posTo(item.Values)}
	for _, row := range item.Rows {
		output.Rows = append(output.Rows, rowValueExprTo(row))
	}
	return output
}

func xrowValueExprTo(item *sqlast.RowValueExpr) (*xast.RowValueExpr, error) {
	output := &xast.RowValueExpr{
		LParen: xposTo(item.LParen),
		RParen: xposTo(item.RParen)}
	for _, value := range item.Values {
		x, err := xvalueStmtTo(value)
		if err != nil { return nil, err }
		output.Values = append(output.Values, x)
	}
	return output, nil
}

func rowValueExprTo(item *xast.RowValueExpr) *sqlast.RowValueExpr {
	output := &sqlast.RowValueExpr{
		LParen: posTo(item.LParen),
		RParen: posTo(item.RParen)}
	for _, value := range item.Values {
		output.Values = append(output.Values, valueStmtTo(value))
	}
	return output
}

func xassignmentTo(item *sqlast.Assignment) (*xast.Assignment, error) {
	x, err := xvalueStmtTo(item.Value)
	if err != nil { return nil, err }
	return &xast.Assignment{
		ID: xidentTo(item.ID),
		Value: x}, nil
}

func assignmentTo(item *xast.Assignment) *sqlast.Assignment {
	return &sqlast.Assignment{
		ID: identTo(item.ID).(*sqlast.Ident),
		Value: valueStmtTo(item.Value)}
}
