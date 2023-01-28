package light

import (
	"github.com/genelet/sqlproto/xast"
	"github.com/genelet/sqlproto/xlight"
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
	x, err := xcolumnDefTo(item.Column)
	return &xast.AddColumnTableAction{
		Add: xposTo(item.Add),
		Column: x}, err
}

func addColumnTableActionTo(item *xast.AddColumnTableAction) *sqlast.AddColumnTableAction {
	return &sqlast.AddColumnTableAction{
		Add: posTo(item.Add),
		Column: columnDefTo(item.Column)}
}

func xalterColumnTableActionTo(item *sqlast.AlterColumnTableAction) (*xast.AlterColumnTableAction, error) {
	x, err := xalterColumnActionTo(item.Action)
	return &xast.AlterColumnTableAction{
		ColumnName: xidentTo(item.ColumnName),
		Alter: xposTo(item.Alter),
		Action: x}, err
}

func alterColumnTableActionTo(item *xast.AlterColumnTableAction) *sqlast.AlterColumnTableAction {
	return &sqlast.AlterColumnTableAction{
		Alter: posTo(item.Alter),
		ColumnName: identTo(item.ColumnName).(*sqlast.Ident),
		Action: alterColumnActionTo(item.Action)}
}

func xaddConstraintTableActionTo(item *sqlast.AddConstraintTableAction) (*xast.AddConstraintTableAction, error) {
	x, err := xtableConstraintTo(item.Constraint)
	return &xast.AddConstraintTableAction{
		Add: xposTo(item.Add),
		Constraint: x}, err
}

func addConstraintTableActionTo(item *xast.AddConstraintTableAction) *sqlast.AddConstraintTableAction {
	return &sqlast.AddConstraintTableAction{
		Add: posTo(item.Add),
		Constraint: tableConstraintTo(item.Constraint)}
}

func xdropConstraintTableActionTo(item *sqlast.DropConstraintTableAction) (*xast.DropConstraintTableAction, error) {
	return &xast.DropConstraintTableAction{
		Name: xidentTo(item.Name),
		Drop: xposTo(item.Drop),
		Cascade: item.Cascade,
		CascadePos: xposTo(item.CascadePos)}, nil
}

func dropConstraintTableActionTo(item *xast.DropConstraintTableAction) *sqlast.DropConstraintTableAction {
	return &sqlast.DropConstraintTableAction{
		Name: identTo(item.Name).(*sqlast.Ident),
		Drop: posTo(item.Drop),
		Cascade: item.Cascade,
		CascadePos: posTo(item.CascadePos)}
}

func xremoveColumnTableActionTo(item *sqlast.RemoveColumnTableAction) (*xast.RemoveColumnTableAction, error) {
	return &xast.RemoveColumnTableAction{
		Name: xidentTo(item.Name),
		Drop: xposTo(item.Drop),
		Cascade: item.Cascade,
		CascadePos: xposTo(item.CascadePos)}, nil
}

func removeColumnTableActionTo(item *xast.RemoveColumnTableAction) *sqlast.RemoveColumnTableAction {
	return &sqlast.RemoveColumnTableAction{
		Name: identTo(item.Name).(*sqlast.Ident),
		Drop: posTo(item.Drop),
		Cascade: item.Cascade,
		CascadePos: posTo(item.CascadePos)}
}
