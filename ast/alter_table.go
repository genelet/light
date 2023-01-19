package ast

import (
	"github.com/genelet/sqlproto/xast"
	"github.com/akito0107/xsqlparser/sqlast"
)

func XAlterTableTo(stmt *sqlast.AlterTableStmt) (*xast.AlterTableStmt, error) {
	action, err := xalterTableActionTo(stmt.Action)
	return &xast.AlterTableStmt{
		Alter: xposTo(stmt.Alter),
		TableName: xobjectnameTo(stmt.TableName),
		Action: action}, err
}

func AlterTableTo(stmt *xast.AlterTableStmt) *sqlast.AlterTableStmt {
	return &sqlast.AlterTableStmt{
		Alter: posTo(stmt.Alter),
		TableName: objectnameTo(stmt.TableName),
		Action: alterTableActionTo(stmt.Action)}
}

func xaddColumnTableActionTo(item *sqlast.AddColumnTableAction) (*xast.AddColumnTableAction, error) {
	return nil, nil
}

func addColumnTableActionTo(item *xast.AddColumnTableAction) *sqlast.AddColumnTableAction {
	return nil
}

func xaddConstraintTableActionTo(item *sqlast.AddConstraintTableAction) (*xast.AddConstraintTableAction, error) {
	return nil, nil
}

func addConstraintTableActionTo(item *xast.AddConstraintTableAction) *sqlast.AddConstraintTableAction {
	return nil
}

func xdropConstraintTableActionTo(item *sqlast.DropConstraintTableAction) (*xast.DropConstraintTableAction, error) {
	return nil, nil
}

func dropConstraintTableActionTo(item *xast.DropConstraintTableAction) *sqlast.DropConstraintTableAction {
	return nil
}

func xremoveColumnTableActionTo(item *sqlast.RemoveColumnTableAction) (*xast.RemoveColumnTableAction, error) {
	return nil, nil
}

func removeColumnTableActionTo(item *xast.RemoveColumnTableAction) *sqlast.RemoveColumnTableAction {
	return nil
}
