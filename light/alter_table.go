package light

import (
	"github.com/genelet/sqlproto/xast"
	"github.com/genelet/sqlproto/xlight"
)

func XAlterTableTo(stmt *xlight.AlterTableStmt) *xast.AlterTableStmt {
	return &xast.AlterTableStmt{
		Alter: xposTo(),
		TableName: xobjectnameTo(stmt.TableName),
		Action: xalterTableActionTo(stmt.Action)}
}

func AlterTableTo(stmt *xast.AlterTableStmt) *xlight.AlterTableStmt {
	return &xlight.AlterTableStmt{
		TableName: objectnameTo(stmt.TableName),
		Action: alterTableActionTo(stmt.Action)}
}

func xaddColumnTableActionTo(item *xlight.ColumnDef) *xast.AddColumnTableAction {
	if item == nil { return nil }

	return &xast.AddColumnTableAction{
		Add: xposTo(),
		Column: xcolumnDefTo(item)}
}

func addColumnTableActionTo(item *xast.AddColumnTableAction) *xlight.ColumnDef {
	if item == nil { return nil }

	return columnDefTo(item.Column)
}

func xalterColumnTableActionTo(item *xlight.AlterColumnTableAction) *xast.AlterColumnTableAction {
	if item == nil { return nil }

	return &xast.AlterColumnTableAction{
		ColumnName: xidentTo(item.ColumnName),
		Alter: xposTo(),
		Action: xalterColumnActionTo(item.Action)}
}

func alterColumnTableActionTo(item *xast.AlterColumnTableAction) *xlight.AlterColumnTableAction {
	if item == nil { return nil }

	return &xlight.AlterColumnTableAction{
		ColumnName: identTo(item.ColumnName),
		Action: alterColumnActionTo(item.Action)}
}

func xaddConstraintTableActionTo(item *xlight.TableConstraint) *xast.AddConstraintTableAction {
	if item == nil { return nil }

	return &xast.AddConstraintTableAction{
		Add: xposTo(),
		Constraint: xtableConstraintTo(item)}
}

func addConstraintTableActionTo(item *xast.AddConstraintTableAction) *xlight.TableConstraint {
	if item == nil { return nil }

	return tableConstraintTo(item.Constraint)
}

func xdropConstraintTableActionTo(item *xlight.DropConstraintTableAction) *xast.DropConstraintTableAction {
	if item == nil { return nil }

	return &xast.DropConstraintTableAction{
		Name: xidentTo(item.Name),
		Drop: xposTo(),
		Cascade: item.Cascade,
		CascadePos: xposTo(item.Cascade)}
}

func dropConstraintTableActionTo(item *xast.DropConstraintTableAction) *xlight.DropConstraintTableAction {
	if item == nil { return nil }

	return &xlight.DropConstraintTableAction{
		Name: identTo(item.Name),
		Cascade: item.Cascade}
}

func xremoveColumnTableActionTo(item *xlight.RemoveColumnTableAction) *xast.RemoveColumnTableAction {
	if item == nil { return nil }

	return &xast.RemoveColumnTableAction{
		Name: xidentTo(item.Name),
		Drop: xposTo(),
		Cascade: item.Cascade,
		CascadePos: xposTo(item.Cascade)}
}

func removeColumnTableActionTo(item *xast.RemoveColumnTableAction) *xlight.RemoveColumnTableAction {
	if item == nil { return nil }

	return &xlight.RemoveColumnTableAction{
		Name: identTo(item.Name),
		Cascade: item.Cascade}
}

func xsetDefaultColumnActionTo(item *xlight.ValueNode) *xast.SetDefaultColumnAction {
	return &xast.SetDefaultColumnAction{
		Set: xposTo(),
		Default: xvalueNodeTo(item)}
}

func setDefaultColumnActionTo(item *xast.SetDefaultColumnAction) *xlight.ValueNode {
	return valueNodeTo(item.Default)
}

func xpgAlterDataTypeColumnActionTo(item *xlight.Type) *xast.PGAlterDataTypeColumnAction {
	return &xast.PGAlterDataTypeColumnAction {
		Type: xposTo(),
		DataType: xtypeTo(item)}
}

func pgAlterDataTypeColumnActionTo(item *xast.PGAlterDataTypeColumnAction) *xlight.Type {
	return typeTo(item.DataType)
}

func xdropDefaultColumnActionTo(item xlight.DropDefaultColumnAction) *xast.DropDefaultColumnAction {
	if item == xlight.DropDefaultColumnAction_DropDefaultColumnActionUnknown { return nil }
	return &xast.DropDefaultColumnAction{
		Drop: xposTo(),
		Default: xposTo(item)}
}

func dropDefaultColumnActionTo(item *xast.DropDefaultColumnAction) xlight.DropDefaultColumnAction {
	if item == nil { return xlight.DropDefaultColumnAction_DropDefaultColumnActionUnknown }
	return xlight.DropDefaultColumnAction_DropDefaultColumnActionConfirm
}

func xpgSetNotNullColumnActionTo(item xlight.PGSetNotNullColumnAction) *xast.PGSetNotNullColumnAction {
	if item == xlight.PGSetNotNullColumnAction_PGSetNotNullColumnActionUnknown { return nil }
	return &xast.PGSetNotNullColumnAction{
		Set: xposTo(),
		Null: xposTo(item)}
}

func pgSetNotNullColumnActionTo(item *xast.PGSetNotNullColumnAction) xlight.PGSetNotNullColumnAction {
	if item == nil { return xlight.PGSetNotNullColumnAction_PGSetNotNullColumnActionUnknown }
	return xlight.PGSetNotNullColumnAction_PGSetNotNullColumnActionConfirm
}

func xpgDropNotNullColumnActionTo(item xlight.PGDropNotNullColumnAction) *xast.PGDropNotNullColumnAction {
	if item == xlight.PGDropNotNullColumnAction_PGDropNotNullColumnActionUnknown { return nil }
	return &xast.PGDropNotNullColumnAction{
		Drop: xposTo(),
		Null: xposTo(item)}
}

func pgDropNotNullColumnActionTo(item *xast.PGDropNotNullColumnAction) xlight.PGDropNotNullColumnAction {
	if item == nil { return xlight.PGDropNotNullColumnAction_PGDropNotNullColumnActionUnknown }
	return xlight.PGDropNotNullColumnAction_PGDropNotNullColumnActionConfirm
}
