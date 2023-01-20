package ast

import (
	"fmt"
	"github.com/genelet/sqlproto/xast"
	"github.com/akito0107/xsqlparser/sqlast"
)

func XCreateTableTo(stmt *sqlast.CreateTableStmt) (*xast.CreateTableStmt, error) {
	output := &xast.CreateTableStmt{
		Create: xposTo(stmt.Create),
		Name: xobjectnameTo(stmt.Name),
		NotExists: stmt.NotExists}
	if stmt.Location != nil {
		output.Location = *stmt.Location
	}

	for _, item := range stmt.Elements {
		v, err := xtableElementTo(item)
		if err != nil { return nil, err }
		output.Elements = append(output.Elements, v)
	}
	for _, item := range stmt.Options {
		v, err := xtableOptionTo(item)
		if err != nil { return nil, err }
		output.Options = append(output.Options, v)
	}
	return output, nil
}

func CreateTableTo(stmt *xast.CreateTableStmt) *sqlast.CreateTableStmt {
	output := &sqlast.CreateTableStmt{
		Create: posTo(stmt.Create),
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

func xtableConstraintTo(item *sqlast.TableConstraint) (*xast.TableConstraint, error) {
	x, err := xtableConstraintSpecTo(item.Spec)
	return &xast.TableConstraint{
		Constraint: xposTo(item.Constraint),
		Name: xidentTo(item.Name),
		Spec: x}, err
}

func tableConstraintTo(item *xast.TableConstraint) *sqlast.TableConstraint {
	output := &sqlast.TableConstraint{
		Constraint: posTo(item.Constraint),
		Spec: tableConstraintSpecTo(item.Spec)}
	if item.Name != nil {
		output.Name = identTo(item.Name).(*sqlast.Ident)
	}
	return output
}

func xuniqueTableConstraintTo(item *sqlast.UniqueTableConstraint) (*xast.UniqueTableConstraint, error) {
	output := &xast.UniqueTableConstraint{
		IsPrimary: item.IsPrimary,
		Primary: xposTo(item.Primary),
		Unique: xposTo(item.Unique),
		RParen: xposTo(item.RParen)}
	for _, column := range item.Columns {
		output.Columns = append(output.Columns, xidentTo(column))
	}
	return output, nil
}

func uniqueTableConstraintTo(item *xast.UniqueTableConstraint) *sqlast.UniqueTableConstraint {
	output := &sqlast.UniqueTableConstraint{
		IsPrimary: item.IsPrimary,
		Primary: posTo(item.Primary),
		Unique: posTo(item.Unique),
		RParen: posTo(item.RParen)}
	for _, column := range item.Columns {
		output.Columns = append(output.Columns, identTo(column).(*sqlast.Ident))
	}
	return output
}

func xreferentialTableConstraintTo(item *sqlast.ReferentialTableConstraint) (*xast.ReferentialTableConstraint, error) {
	referenceKeyExpr, err := xreferenceKeyExprTo(item.KeyExpr)
	if err != nil { return nil, err }
	output := &xast.ReferentialTableConstraint{
		Foreign: xposTo(item.Foreign),
		KeyExpr: referenceKeyExpr}
	for _, column := range item.Columns {
		output.Columns = append(output.Columns, xidentTo(column))
	}
	return output, nil
}

func referentialTableConstraintTo(item *xast.ReferentialTableConstraint) *sqlast.ReferentialTableConstraint {
	output := &sqlast.ReferentialTableConstraint{
		Foreign: posTo(item.Foreign),
		KeyExpr: referenceKeyExprTo(item.KeyExpr)}
	for _, column := range item.Columns {
		output.Columns = append(output.Columns, identTo(column).(*sqlast.Ident))
	}
	return output
}

func xreferenceKeyExprTo(item *sqlast.ReferenceKeyExpr) (*xast.ReferenceKeyExpr, error) {
	output := &xast.ReferenceKeyExpr{
		TableName: xidentTo(item.TableName),
		RParen: xposTo(item.RParen)}
	for _, column := range item.Columns {
		output.Columns = append(output.Columns, xidentTo(column))
	}
	return output, nil
}

func referenceKeyExprTo(item *xast.ReferenceKeyExpr) *sqlast.ReferenceKeyExpr {
	output := &sqlast.ReferenceKeyExpr{
		TableName: identTo(item.TableName).(*sqlast.Ident),
		RParen: posTo(item.RParen)}
	for _, column := range item.Columns {
		output.Columns = append(output.Columns, identTo(column).(*sqlast.Ident))
	}
	return output
}

func xcolumnDefTo(item *sqlast.ColumnDef) (*xast.ColumnDef, error) {
	x, err := xvalueNodeTo(item.Default)
	if err != nil { return nil, err }
	y, err := xtypeTo(item.DataType)
	if err != nil { return nil, err }
    columnDef := &xast.ColumnDef{
		Name: xidentTo(item.Name),
		Default: x,
		DataType: y}
	for _, mydeco := range item.MyDataTypeDecoration {
		x, err := xmyDataTypeDecorationTo(mydeco)
		if err != nil { return nil, err }
		columnDef.MyDecos = append(columnDef.MyDecos, x)
	}

	for _, constraint := range item.Constraints {
		x, err := xcolumnConstraintTo(constraint)
		if err != nil { return nil, err }
		columnDef.Constraints = append(columnDef.Constraints, x)
	}

	return columnDef, nil
}

func columnDefTo(item *xast.ColumnDef) *sqlast.ColumnDef {
	output := &sqlast.ColumnDef{
		Name: identTo(item.Name).(*sqlast.Ident),
		Default: valueNodeTo(item.Default),
		DataType: typeTo(item.DataType)}
	for _, mydeco := range item.MyDecos {
		output.MyDataTypeDecoration = append(output.MyDataTypeDecoration, myDataTypeDecorationTo(mydeco))
    }
	
	for _, constraint := range item.Constraints {
		x := columnConstraintTo(constraint)
		output.Constraints = append(output.Constraints, x)
	}

	return output
}

func xmyDataTypeDecorationTo(item sqlast.MyDataTypeDecoration) (*xast.MyDataTypeDecoration, error) {
	x, ok := item.(*sqlast.AutoIncrement)
	if !ok {
		return nil, fmt.Errorf("missing my data decoration for %T", item)
	}
	return &xast.MyDataTypeDecoration{
		Automent: &xast.AutoIncrement{
			Auto: xposTo(x.Auto),
			Increment: xposTo(x.Increment)}}, nil
}

func myDataTypeDecorationTo(item *xast.MyDataTypeDecoration) sqlast.MyDataTypeDecoration {
	return &sqlast.AutoIncrement{
			Auto: posTo(item.Automent.Auto),
			Increment: posTo(item.Automent.Increment)}
}

func xcolumnConstraintTo(item *sqlast.ColumnConstraint) (*xast.ColumnConstraint, error) {
	x, err := xcolumnConstraintSpecTo(item.Spec)
    return &xast.ColumnConstraint{
        Name: xidentTo(item.Name),
		Constraint: xposTo(item.Constraint),
		Spec: x}, err
}

func columnConstraintTo(item *xast.ColumnConstraint) *sqlast.ColumnConstraint {
    output := &sqlast.ColumnConstraint{
		Constraint: posTo(item.Constraint),
		Spec: columnConstraintSpecTo(item.Spec)}
	if item.Name != nil {
		output.Name = identTo(item.Name).(*sqlast.Ident)
	}
	return output
}

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
