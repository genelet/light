package light

import (
	"github.com/genelet/sqlproto/xast"
	"github.com/genelet/sqlproto/xlight"
)

func XCreateTableTo(stmt *xlight.CreateTableStmt) *xast.CreateTableStmt {
	output := &xast.CreateTableStmt{
		Create: xposTo(),
		Location: stmt.Location,
		Name: xobjectnameTo(stmt.Name),
		NotExists: stmt.NotExists}
	for _, item := range stmt.Elements {
		output.Elements = append(output.Elements, xtableElementTo(item))
	}
	for _, item := range stmt.Options {
		output.Options = append(output.Options, xtableOptionTo(item))
	}
	return output
}

func CreateTableTo(stmt *xast.CreateTableStmt) *xlight.CreateTableStmt {
	output := &xlight.CreateTableStmt{
		Location: stmt.Location,
		Name: objectnameTo(stmt.Name),
		NotExists: stmt.NotExists}
	for _, item := range stmt.Elements {
		output.Elements = append(output.Elements, tableElementTo(item))
	}
	for _, item := range stmt.Options {
		output.Options = append(output.Options, tableOptionTo(item))
	}
	return output
}

func xtableConstraintTo(item *xlight.TableConstraint) *xast.TableConstraint {
	if item == nil { return nil }

	return &xast.TableConstraint{
		Constraint: xposTo(),
		Name: xidentTo(item.Name),
		Spec: xtableConstraintSpecTo(item.Spec)}
}

func tableConstraintTo(item *xast.TableConstraint) *xlight.TableConstraint {
	if item == nil { return nil }

	return &xlight.TableConstraint{
		Name: identTo(item.Name),
		Spec: tableConstraintSpecTo(item.Spec)}
}

func xcheckTableConstraintTo(item *xlight.BinaryExpr) *xast.CheckTableConstraint {
	if item == nil { return nil }
	return &xast.CheckTableConstraint{
		Expr: xbinaryExprTo(item),
		Check: xposTo(),
		RParen: xposTo(item)}
}

func checkTableConstraintTo(item *xast.CheckTableConstraint) *xlight.BinaryExpr  {
	if item == nil { return nil }
	return binaryExprTo(item.Expr)
}

func xcheckColumnSpecTo(item *xlight.BinaryExpr) *xast.CheckColumnSpec {
	if item == nil { return nil }
	return &xast.CheckColumnSpec{
		Expr: xbinaryExprTo(item),
		Check: xposTo(),
		RParen: xposTo(item)}
}

func checkColumnSpecTo(item *xast.CheckColumnSpec) *xlight.BinaryExpr {
	if item == nil { return nil }
	return binaryExprTo(item.Expr)
}

func xuniqueColumnSpecTo(item *xlight.UniqueColumnSpec) *xast.UniqueColumnSpec {
	if item == nil { return nil }
	return &xast.UniqueColumnSpec{
		IsPrimaryKey: item.IsPrimaryKey,
		Primary: xposTo(item.IsPrimaryKey),
		Key: xposTo(),
		Unique: xposTo(item)}
}

func uniqueColumnSpecTo(item *xast.UniqueColumnSpec) *xlight.UniqueColumnSpec {
	if item == nil { return nil }
	return &xlight.UniqueColumnSpec{
		IsPrimaryKey: item.IsPrimaryKey}
}

func xnotNullColumnSpecTo(item xlight.NotNullColumnSpecType) *xast.NotNullColumnSpec {
	if item == xlight.NotNullColumnSpecType_NotNullColumnSpecTypeUnknown {
		return nil
	}

	return &xast.NotNullColumnSpec{
		Not: xposTo(),
		Null: xposTo(item)}
}

func notNullColumnSpecTo(item *xast.NotNullColumnSpec) xlight.NotNullColumnSpecType {
	if item == nil { return xlight.NotNullColumnSpecType_NotNullColumnSpecTypeUnknown }

	return xlight.NotNullColumnSpecType_NotNullColumnSpec
}

func xreferencesColumnSpecTo(item *xlight.ReferencesColumnSpec) *xast.ReferencesColumnSpec {
	if item == nil { return nil }
	output := &xast.ReferencesColumnSpec{
		References: xposTo(),
		RParen: xposTo(item),
		TableName: xobjectnameTo(item.TableName)}
	for _, column := range item.Columns {
		output.Columns = append(output.Columns, xidentTo(column))
	}
	return output
}

func referencesColumnSpecTo(item *xast.ReferencesColumnSpec) *xlight.ReferencesColumnSpec {
	if item == nil { return nil }
	output := &xlight.ReferencesColumnSpec{
		TableName: objectnameTo(item.TableName)}
	for _, column := range item.Columns {
		output.Columns = append(output.Columns, identTo(column))
	}
	return output
}

func xuniqueTableConstraintTo(item *xlight.UniqueTableConstraint) *xast.UniqueTableConstraint {
	if item == nil { return nil }

	output := &xast.UniqueTableConstraint{
		IsPrimary: item.IsPrimary,
		Primary: xposTo(item.IsPrimary),
		Unique: xposTo(),
		RParen: xposTo(item)}
	for _, column := range item.Columns {
		output.Columns = append(output.Columns, xidentTo(column))
	}
	return output
}

func uniqueTableConstraintTo(item *xast.UniqueTableConstraint) *xlight.UniqueTableConstraint {
	if item == nil { return nil }

	output := &xlight.UniqueTableConstraint{
		IsPrimary: item.IsPrimary}
	for _, column := range item.Columns {
		output.Columns = append(output.Columns, identTo(column))
	}
	return output
}

func xreferentialTableConstraintTo(item *xlight.ReferentialTableConstraint) *xast.ReferentialTableConstraint {
	if item == nil { return nil }

	output := &xast.ReferentialTableConstraint{
		Foreign: xposTo(),
		KeyExpr: xreferenceKeyExprTo(item.KeyExpr)}
	for _, column := range item.Columns {
		output.Columns = append(output.Columns, xidentTo(column))
	}
	return output
}

func referentialTableConstraintTo(item *xast.ReferentialTableConstraint) *xlight.ReferentialTableConstraint {
	if item == nil { return nil }

	output := &xlight.ReferentialTableConstraint{
		KeyExpr: referenceKeyExprTo(item.KeyExpr)}
	for _, column := range item.Columns {
		output.Columns = append(output.Columns, identTo(column))
	}
	return output
}

func xreferenceKeyExprTo(item *xlight.ReferenceKeyExpr) *xast.ReferenceKeyExpr {
	if item == nil { return nil }

	output := &xast.ReferenceKeyExpr{
		TableName: xidentTo(item.TableName),
		RParen: xposTo(item)}
	for _, column := range item.Columns {
		output.Columns = append(output.Columns, xidentTo(column))
	}
	return output
}

func referenceKeyExprTo(item *xast.ReferenceKeyExpr) *xlight.ReferenceKeyExpr {
	if item == nil { return nil }

	output := &xlight.ReferenceKeyExpr{
		TableName: identTo(item.TableName)}
	for _, column := range item.Columns {
		output.Columns = append(output.Columns, identTo(column))
	}
	return output
}

func xcolumnDefTo(item *xlight.ColumnDef) *xast.ColumnDef {
	if item == nil { return nil }

    columnDef := &xast.ColumnDef{
		Name: xidentTo(item.Name),
		Default: xvalueNodeTo(item.Default),
		DataType: xtypeTo(item.DataType)}
	for _, mydeco := range item.MyDecos {
		columnDef.MyDecos = append(columnDef.MyDecos, xmyDataTypeDecorationTo(mydeco))
	}
	for _, constraint := range item.Constraints {
		columnDef.Constraints = append(columnDef.Constraints, xcolumnConstraintTo(constraint))
	}

	return columnDef
}

func columnDefTo(item *xast.ColumnDef) *xlight.ColumnDef {
	if item == nil { return nil }

	output := &xlight.ColumnDef{
		Name: identTo(item.Name),
		Default: valueNodeTo(item.Default),
		DataType: typeTo(item.DataType)}
	for _, mydeco := range item.MyDecos {
		output.MyDecos = append(output.MyDecos, myDataTypeDecorationTo(mydeco))
    }
	for _, constraint := range item.Constraints {
		output.Constraints = append(output.Constraints, columnConstraintTo(constraint))
	}

	return output
}

func xmyDataTypeDecorationTo(item xlight.AutoIncrementType) *xast.MyDataTypeDecoration {
	return &xast.MyDataTypeDecoration{
		Automent: &xast.AutoIncrement{
			Auto: xposTo(),
			Increment: xposTo(item)}}
}

func myDataTypeDecorationTo(item *xast.MyDataTypeDecoration) xlight.AutoIncrementType {
	return xlight.AutoIncrementType_AutoIncrement
}

func xcolumnConstraintTo(item *xlight.ColumnConstraint) *xast.ColumnConstraint {
    return &xast.ColumnConstraint{
        Name: xidentTo(item.Name),
		Constraint: xposTo(),
		Spec: xcolumnConstraintSpecTo(item.Spec)}
}

func columnConstraintTo(item *xast.ColumnConstraint) *xlight.ColumnConstraint {
    return &xlight.ColumnConstraint{
		Name: identTo(item.Name),
		Spec: columnConstraintSpecTo(item.Spec)}
}

/*
func XDropTableTo(stmt *sqlast.DropTableStmt) *xast.DropTableStmt {
    output := &xast.DropTableStmt{
    	Cascade: stmt.Cascade,
    	CascadePos: xposTo(stmt.CascadePos),
    	IfExists: stmt.IfExists,
    	Drop:  xposTo(stmt.Drop)}
	for _, name := range stmt.TableNames {
		output.TableNames = append(output.TableNames, xobjectnameTo(name))
	}
	return output
}

func DropTableTo(stmt *xast.DropTableStmt) *sqlast.DropTableStmt {
    output := &sqlast.DropTableStmt{
    	Cascade: stmt.Cascade,
    	CascadePos: posTo(stmt.CascadePos),
    	IfExists: stmt.IfExists,
    	Drop:  posTo(stmt.Drop)}
	for _, name := range stmt.TableNames {
		output.TableNames = append(output.TableNames, objectnameTo(name))
	}
	return output
}
*/
