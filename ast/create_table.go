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

func xtableElementTo(item sqlast.TableElement) (*xast.TableElement, error) {
	element := new(xast.TableElement)
	switch t := item.(type) {
	case *sqlast.ColumnDef:
		x, err := xcolumnDefTo(t)
		if err != nil { return nil, err }
		element.TableElementClause = &xast.TableElement_ColumnDefElement{
			ColumnDefElement:x}
	case *sqlast.TableConstraint:
		x, err := xtableConstraintTo(t)
		if err != nil { return nil, err }
		element.TableElementClause = &xast.TableElement_TableConstraintElement{
			TableConstraintElement: x}
	default:
		return nil, fmt.Errorf("missing table element type %T", t)
	}
	return element, nil
}

func tableElementTo(item *xast.TableElement) sqlast.TableElement {
	if item.GetColumnDefElement() != nil {
		return columnDefTo(item.GetColumnDefElement())
	}
	return tableConstraintTo(item.GetTableConstraintElement())
}

func xtableOptionTo(item sqlast.TableOption) (*xast.TableOption, error) {
	output := &xast.TableOption{}
	switch t := item.(type) {
	case *sqlast.MyEngine:
		output.TableOptionClause = &xast.TableOption_MyEngineOption{MyEngineOption:
            &xast.MyEngine{
                Engine: xposTo(t.Engine),
                Equal: t.Equal,
                Name: xidentTo(t.Name)}}
	case *sqlast.MyCharset:
		output.TableOptionClause = &xast.TableOption_MyCharsetOption{MyCharsetOption:
            &xast.MyCharset{
                IsDefault: t.IsDefault,
                Default: xposTo(t.Default),
                Charset: xposTo(t.Charset),
                Equal: t.Equal,
                Name: xidentTo(t.Name)}}
	default:
		return nil, fmt.Errorf("missing table element type %T", item)
	}
	return output, nil
}

func tableOptionTo(item *xast.TableOption) sqlast.TableOption {
    if x := item.GetMyEngineOption(); x != nil {
        return &sqlast.MyEngine{
            Engine: posTo(x.Engine),
            Equal: x.Equal,
            Name: identTo(x.Name).(*sqlast.Ident)}
    } else if x := item.GetMyCharsetOption(); x != nil {
		return &sqlast.MyCharset{
            IsDefault: x.IsDefault,
            Default: posTo(x.Default),
            Charset: posTo(x.Charset),
			Equal: x.Equal,
            Name: identTo(x.Name).(*sqlast.Ident)}
	}
	return nil
}

func xtableConstraintTo(item *sqlast.TableConstraint) (*xast.TableConstraint, error) {
	output := &xast.TableConstraint{
		Constraint: xposTo(item.Constraint),
		Name: xidentTo(item.Name)}
	switch t := item.Spec.(type) {
	case *sqlast.ReferentialTableConstraint:
		x, err := xreferentialTableConstraintTo(t)
		if err != nil { return nil, err }
		output.Spec = &xast.TableConstraint_SpecReference{SpecReference: x}
	case *sqlast.UniqueTableConstraint:
		x, err := xuniqueTableConstraintTo(t)
		if err != nil { return nil, err }
		output.Spec = &xast.TableConstraint_SpecUnique{SpecUnique: x}
	case *sqlast.CheckTableConstraint:
		switch s := t.Expr.(type) {
		case *sqlast.BinaryExpr:
			x, err := xbinaryExprTo(s)
			if err != nil { return nil, err }
			output.Spec = &xast.TableConstraint_SpecCheck{
				SpecCheck: &xast.CheckTableConstraint{
					Check: xposTo(t.Check),
					RParen: xposTo(t.RParen),
					Expr: x}}
		default:
			return nil, fmt.Errorf("missing type in table constaint Spec: %T", s)
		}
	default:
		return nil, fmt.Errorf("missing type in table constaint: %T", t)
	}
	return output, nil
}

func tableConstraintTo(item *xast.TableConstraint) *sqlast.TableConstraint {
	output := &sqlast.TableConstraint{
		Constraint: posTo(item.Constraint)}
	if item.Name != nil {
		output.Name = identTo(item.Name).(*sqlast.Ident)
	}
	if x := item.GetSpecReference(); x != nil {
		output.Spec = referentialTableConstraintTo(x)
	} else if x := item.GetSpecUnique(); x != nil {
		output.Spec = uniqueTableConstraintTo(x)
	} else {
		x := item.GetSpecCheck()
		output.Spec = &sqlast.CheckTableConstraint{
			Check: posTo(x.Check),
			RParen: posTo(x.RParen),
			Expr: binaryExprTo(x.Expr)}
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
    columnDef := &xast.ColumnDef{
		Name: xidentTo(item.Name)}
	for _, mydeco := range item.MyDataTypeDecoration {
		x, err := xmyDataTypeDecorationTo(mydeco)
		if err != nil { return nil, err }
		columnDef.MyDecos = append(columnDef.MyDecos, x)
	}
	switch t := item.DataType.(type) {
	case *sqlast.Int:
		columnDef.DataType = &xast.ColumnDef_IntData{IntData: xintTo(t)}
		if item.Default != nil {
			columnDef.Default = &xast.ColumnDef_LongDefault{LongDefault: xlongTo(item.Default.(*sqlast.LongValue))}
		}
	case *sqlast.SmallInt:
		columnDef.DataType = &xast.ColumnDef_SmallIntData{SmallIntData: xsmallIntTo(t)}
		if item.Default != nil {
			columnDef.Default = &xast.ColumnDef_LongDefault{LongDefault: xlongTo(item.Default.(*sqlast.LongValue))}
		}
	case *sqlast.CharType:
		columnDef.DataType = &xast.ColumnDef_CharData{CharData: xcharTypeTo(t)}
		if item.Default != nil {
			columnDef.Default = &xast.ColumnDef_StringDefault{StringDefault: xstringTo(item.Default.(*sqlast.SingleQuotedString))}
		}
	case *sqlast.VarcharType:
		columnDef.DataType = &xast.ColumnDef_VarcharData{VarcharData: xvarcharTypeTo(t)}
		if item.Default != nil {
			columnDef.Default = &xast.ColumnDef_StringDefault{StringDefault: xstringTo(item.Default.(*sqlast.SingleQuotedString))}
		}
	default:
		return nil, fmt.Errorf("missing column def type: %T", t)
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
		Name: identTo(item.Name).(*sqlast.Ident)}
	for _, mydeco := range item.MyDecos {
		output.MyDataTypeDecoration = append(output.MyDataTypeDecoration, myDataTypeDecorationTo(mydeco))
    }

	if item.GetIntData() != nil {
		output.DataType = intTo(item.GetIntData())
		if item.Default != nil {
			output.Default = longTo(item.GetLongDefault())
		}
	} else if item.GetSmallIntData() != nil {
		output.DataType = smallIntTo(item.GetSmallIntData())
		if item.Default != nil {
			output.Default = longTo(item.GetLongDefault())
		}
	} else if item.GetCharData() != nil {
		output.DataType = charTypeTo(item.GetCharData())
		if item.Default != nil {
			output.Default = stringTo(item.GetStringDefault())
		}
	} else { // GetVarcharData()
		output.DataType = varcharTypeTo(item.GetVarcharData())
		if item.Default != nil {
			output.Default = stringTo(item.GetStringDefault())
		}
	}
	for _, constraint := range item.Constraints {
		output.Constraints = append(output.Constraints, columnConstraintTo(constraint))
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
    output := &xast.ColumnConstraint{
        Name: xidentTo(item.Name),
		Constraint: xposTo(item.Constraint)}
    switch t := item.Spec.(type) {
    case *sqlast.CheckColumnSpec:
		switch s := t.Expr.(type) {
		case *sqlast.BinaryExpr:
        	x, err := xbinaryExprTo(s)
			if err != nil { return nil, err }
        	output.Spec = &xast.ColumnConstraint_CheckSpec{CheckSpec:
				&xast.CheckColumnSpec{
					Expr: x,
					Check: xposTo(t.Check),
					RParen: xposTo(t.RParen)}}
		default:
			return nil, fmt.Errorf("missing column constraint Expr type: %T", s)
		}
    case *sqlast.UniqueColumnSpec:
        output.Spec = &xast.ColumnConstraint_UniqueSpec{UniqueSpec:
			&xast.UniqueColumnSpec{
				IsPrimaryKey: t.IsPrimaryKey,
				Primary: xposTo(t.Primary),
				Key: xposTo(t.Key),
				Unique: xposTo(t.Unique)}}
    case *sqlast.NotNullColumnSpec:
        output.Spec = &xast.ColumnConstraint_NotNullSpec{NotNullSpec:
			&xast.NotNullColumnSpec{
				Not: xposTo(t.Not),
				Null: xposTo(t.Null)}}
    default:
        return nil, fmt.Errorf("missing column constraint type: %T", t)
    }

    return output, nil
}

func columnConstraintTo(item *xast.ColumnConstraint) *sqlast.ColumnConstraint {
    output := &sqlast.ColumnConstraint{
		Constraint: posTo(item.Constraint)}
	if item.Name != nil {
		output.Name = identTo(item.Name).(*sqlast.Ident)
	}
	if x := item.GetUniqueSpec(); x != nil {
		output.Spec = &sqlast.UniqueColumnSpec{
			IsPrimaryKey: x.IsPrimaryKey,
			Primary: posTo(x.Primary),
			Key: posTo(x.Key),
			Unique: posTo(x.Unique)}
	} else if x := item.GetNotNullSpec(); x != nil {
		output.Spec = &sqlast.NotNullColumnSpec{
			Not: posTo(x.Not),
			Null: posTo(x.Null)}
	} else {
		x := item.GetCheckSpec()
		output.Spec = &sqlast.CheckColumnSpec{
			Expr: binaryExprTo(x.Expr),
			Check: posTo(x.Check),
			RParen: posTo(x.RParen)}
	}
	return output
}
