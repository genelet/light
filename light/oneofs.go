package light

import (
	"github.com/genelet/sqlproto/xlight"
	"github.com/genelet/sqlproto/xast"
)

func xwhereNodeTo(item *xlight.WhereNode) *xast.WhereNode {
	if item == nil { return nil }

	output := &xast.WhereNode{}
    if x := item.GetInQuery(); x != nil {
        where := xinsubqueryTo(x)
        output.WhereNodeClause = &xast.WhereNode_InQuery{InQuery: where}
   	} else if x := item.GetBinExpr(); x != nil { 
        where := xbinaryExprTo(x)
        output.WhereNodeClause = &xast.WhereNode_BinExpr{BinExpr: where}
	} else {
		return nil
    }

	return output
}

func whereNodeTo(item *xast.WhereNode) *xlight.WhereNode {
	if item == nil { return nil }

	output := &xlight.WhereNode{}
	if x := item.GetInQuery(); x != nil {
        where := insubqueryTo(x)
        output.WhereNodeClause = &xlight.WhereNode_InQuery{InQuery: where}
    } else if x := item.GetBinExpr(); x != nil {
        where := binaryExprTo(x)
        output.WhereNodeClause = &xlight.WhereNode_BinExpr{BinExpr: where}
    } else {
		return nil
	}

	return output
}

func xvalueNodeTo(item *xlight.ValueNode) *xast.ValueNode {
	if item == nil { return nil }

	output := &xast.ValueNode{}
	if x := item.GetStringItem(); x != nil {
        output.ValueNodeClause = &xast.ValueNode_StringItem{StringItem: xstringTo(x)}
	} else if x := item.GetLongItem(); x != nil {
        output.ValueNodeClause = &xast.ValueNode_LongItem{LongItem: xlongTo(x)}
	} else if x := item.GetDoubleItem(); x != nil {
        output.ValueNodeClause = &xast.ValueNode_DoubleItem{DoubleItem: xdoubleTo(x)}
	} else if x := item.GetCompoundItem(); x != nil {
		output.ValueNodeClause = &xast.ValueNode_CompoundItem{CompoundItem: xcompoundTo(x)}
    } else {
        return nil
    }

	return output
}

func valueNodeTo(item *xast.ValueNode) *xlight.ValueNode {
	if item == nil { return nil }

	output := &xlight.ValueNode{}
	if x := item.GetStringItem(); x != nil {
        output.ValueNodeClause = &xlight.ValueNode_StringItem{StringItem: stringTo(x)}
	} else if x := item.GetLongItem(); x != nil {
        output.ValueNodeClause = &xlight.ValueNode_LongItem{LongItem: longTo(x)}
	} else if x := item.GetDoubleItem(); x != nil {
        output.ValueNodeClause = &xlight.ValueNode_DoubleItem{DoubleItem: doubleTo(x)}
	} else if x := item.GetCompoundItem(); x != nil {
		output.ValueNodeClause = &xlight.ValueNode_CompoundItem{CompoundItem: compoundTo(x)}
    } else {
        return nil
    }
	return output
}

func xconditionNodeTo(item *xlight.ConditionNode) *xast.ConditionNode {
	if item == nil { return nil }

	output := &xast.ConditionNode{}
	if x := item.GetBinaryItem(); x != nil {
		output.ConditionNodeClause = &xast.ConditionNode_BinaryItem{BinaryItem: xbinaryExprTo(x)}
	} else {
		return nil
	}

	return output
}

func conditionNodeTo(item *xast.ConditionNode) *xlight.ConditionNode {
	if item == nil { return nil }

	output := &xlight.ConditionNode{}
	if x := item.GetBinaryItem(); x != nil {
		output.ConditionNodeClause = &xlight.ConditionNode_BinaryItem{BinaryItem: binaryExprTo(x)}
	} else {
		return nil
	}
	return output
}

func xargsNodeTo(item *xlight.ArgsNode) *xast.ArgsNode {
	if item == nil { return nil }

	output := &xast.ArgsNode{}
	if x := item.GetValueItem(); x != nil {
		output.ArgsNodeClause = &xast.ArgsNode_ValueItem{ValueItem: xvalueNodeTo(x)}
	} else if x := item.GetWhereItem(); x != nil {
		output.ArgsNodeClause = &xast.ArgsNode_WhereItem{WhereItem: xwhereNodeTo(x)}
	} else if x := item.GetFunctionItem(); x != nil {
		output.ArgsNodeClause = &xast.ArgsNode_FunctionItem{FunctionItem: xfunctionTo(x)}
	} else if x := item.GetCaseItem(); x != nil {
		output.ArgsNodeClause = &xast.ArgsNode_CaseItem{CaseItem: xcaseExprTo(x)}
	} else if x := item.GetNestedItem(); x != nil {
		output.ArgsNodeClause = &xast.ArgsNode_NestedItem{NestedItem: xnestedTo(x)}
	} else if x := item.GetUnaryItem(); x != nil {
		output.ArgsNodeClause = &xast.ArgsNode_UnaryItem{UnaryItem: xunaryExprTo(x)}
	} else {
		return nil
	}

	return output
}

func argsNodeTo(item *xast.ArgsNode) *xlight.ArgsNode {
	if item == nil { return nil }

	output := &xlight.ArgsNode{}
	if x := item.GetValueItem(); x != nil {
		output.ArgsNodeClause = &xlight.ArgsNode_ValueItem{ValueItem: valueNodeTo(x)}
	} else if x := item.GetWhereItem(); x != nil {
		output.ArgsNodeClause = &xlight.ArgsNode_WhereItem{WhereItem: whereNodeTo(x)}
	} else if x := item.GetFunctionItem(); x != nil {
		output.ArgsNodeClause = &xlight.ArgsNode_FunctionItem{FunctionItem: functionTo(x)}
	} else if x := item.GetCaseItem(); x != nil {
		output.ArgsNodeClause = &xlight.ArgsNode_CaseItem{CaseItem: caseExprTo(x)}
	} else if x := item.GetNestedItem(); x != nil {
		output.ArgsNodeClause = &xlight.ArgsNode_NestedItem{NestedItem: nestedTo(x)}
	} else if x := item.GetUnaryItem(); x != nil {
		output.ArgsNodeClause = &xlight.ArgsNode_UnaryItem{UnaryItem: unaryExprTo(x)}
	} else {
		return nil
	}

	return output
}

func xsqlSelectItemTo(item *xlight.SQLSelectItem) *xast.SQLSelectItem {
	if item == nil { return nil }

	output := &xast.SQLSelectItem{}
	if x := item.GetUnnamedItem(); x != nil {
		output.SQLSelectItemClause = &xast.SQLSelectItem_UnnamedItem{UnnamedItem:
			xunnamedSelectItemTo(x)}
	} else if x := item.GetAliasItem(); x != nil {
		output.SQLSelectItemClause = &xast.SQLSelectItem_AliasItem{AliasItem:
			&xast.AliasSelectItem{Expr:xargsNodeTo(x.Expr), Alias: xidentTo(x.Alias)}}
	} else if x := item.GetWildcardItem(); x != nil {
		output.SQLSelectItemClause = &xast.SQLSelectItem_WildcardItem{WildcardItem:
			xqualifiedWildcardSelectItemTo(x)}
	} else {
		return nil
	}

	return output
}

func sqlSelectItemTo(item *xast.SQLSelectItem) *xlight.SQLSelectItem {
	if item == nil { return nil }

	output := &xlight.SQLSelectItem{}
	if x := item.GetUnnamedItem(); x != nil {
		output.SQLSelectItemClause = &xlight.SQLSelectItem_UnnamedItem{UnnamedItem:
			unnamedSelectItemTo(x)}
	} else if x := item.GetAliasItem(); x != nil {
		output.SQLSelectItemClause = &xlight.SQLSelectItem_AliasItem{AliasItem:
			&xlight.AliasSelectItem{Expr:argsNodeTo(x.Expr), Alias: identTo(x.Alias)}}
	} else if x := item.GetWildcardItem(); x != nil {
		output.SQLSelectItemClause = &xlight.SQLSelectItem_WildcardItem{WildcardItem:
			qualifiedWildcardSelectItemTo(x)}
	} else {
		return nil
	}

	return output
}

func xsqlSetExprTo(item *xlight.SQLSetExpr) *xast.SQLSetExpr {
	if item == nil { return nil }

	output := &xast.SQLSetExpr{}
    if x := item.GetSelectItem(); x != nil {
		output.SQLSetExprClause = &xast.SQLSetExpr_SelectItem{SelectItem: xselectTo(x)}
    } else if x := item.GetExprItem(); x != nil {
		output.SQLSetExprClause = &xast.SQLSetExpr_ExprItem{ExprItem: xsetOperationExprTo(x)}
    } else {
    	return nil
    }

    return output
}

func sqlSetExprTo(item *xast.SQLSetExpr) *xlight.SQLSetExpr {
    if item == nil { return nil }

	output := &xlight.SQLSetExpr{}
    if x := item.GetSelectItem(); x != nil {
		output.SQLSetExprClause = &xlight.SQLSetExpr_SelectItem{SelectItem: selectTo(x)}
    } else if x := item.GetExprItem(); x != nil {
		output.SQLSetExprClause = &xlight.SQLSetExpr_ExprItem{ExprItem: setOperationExprTo(x)}
    } else {
    	return nil
    }

    return output
}

// start create table


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

func xtableConstraintSpecTo(item sqlast.TableConstraintSpec) (*xast.TableConstraintSpec, error) {
	if item == nil { return nil, nil }

	output := &xast.TableConstraintSpec{}
	switch t := item.(type) {
	case *sqlast.ReferentialTableConstraint:
		x, err := xreferentialTableConstraintTo(t)
		if err != nil { return nil, err }
		output.TableContraintSpecClause = &xast.TableConstraintSpec_ReferenceItem{ReferenceItem: x}
	case *sqlast.UniqueTableConstraint:
		x, err := xuniqueTableConstraintTo(t)
		if err != nil { return nil, err }
		output.TableContraintSpecClause = &xast.TableConstraintSpec_UniqueItem{UniqueItem: x}
	case *sqlast.CheckTableConstraint:
		switch s := t.Expr.(type) {
		case *sqlast.BinaryExpr:
			x, err := xbinaryExprTo(s)
			if err != nil { return nil, err }
			output.TableContraintSpecClause = &xast.TableConstraintSpec_CheckItem{
				CheckItem: &xast.CheckTableConstraint{
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

func tableConstraintSpecTo(item *xast.TableConstraintSpec) sqlast.TableConstraintSpec {
	if item == nil { return nil }

	if x := item.GetReferenceItem(); x != nil {
		return referentialTableConstraintTo(x)
	} else if x := item.GetUniqueItem(); x != nil {
		return uniqueTableConstraintTo(x)
	} else {
		x := item.GetCheckItem()
		return &sqlast.CheckTableConstraint{
			Check: posTo(x.Check),
			RParen: posTo(x.RParen),
			Expr: binaryExprTo(x.Expr)}
	}
	return nil
}

func xcolumnConstraintSpecTo(item sqlast.ColumnConstraintSpec) (*xast.ColumnConstraintSpec, error) {
	if item == nil { return nil, nil }

    output := &xast.ColumnConstraintSpec{}
    switch t := item.(type) {
    case *sqlast.CheckColumnSpec:
		switch s := t.Expr.(type) {
		case *sqlast.BinaryExpr:
        	x, err := xbinaryExprTo(s)
			if err != nil { return nil, err }
        	output.ColumnConstraintSpecClause = &xast.ColumnConstraintSpec_CheckItem{CheckItem:
				&xast.CheckColumnSpec{
					Expr: x,
					Check: xposTo(t.Check),
					RParen: xposTo(t.RParen)}}
		default:
			return nil, fmt.Errorf("missing column constraint Expr type: %T", s)
		}
    case *sqlast.UniqueColumnSpec:
        output.ColumnConstraintSpecClause = &xast.ColumnConstraintSpec_UniqueItem{UniqueItem:
			&xast.UniqueColumnSpec{
				IsPrimaryKey: t.IsPrimaryKey,
				Primary: xposTo(t.Primary),
				Key: xposTo(t.Key),
				Unique: xposTo(t.Unique)}}
    case *sqlast.NotNullColumnSpec:
        output.ColumnConstraintSpecClause = &xast.ColumnConstraintSpec_NotNullItem{NotNullItem:
			&xast.NotNullColumnSpec{
				Not: xposTo(t.Not),
				Null: xposTo(t.Null)}}
    case *sqlast.ReferencesColumnSpec:
		ref := &xast.ReferencesColumnSpec{
			References: xposTo(t.References),
			RParen: xposTo(t.RParen),
			TableName: xobjectnameTo(t.TableName)}
		for _, column := range t.Columns {
			ref.Columns = append(ref.Columns, xidentTo(column))
		}
        output.ColumnConstraintSpecClause = &xast.ColumnConstraintSpec_ReferenceItem{ReferenceItem: ref}
    default:
        return nil, fmt.Errorf("missing column constraint type: %T", t)
    }

    return output, nil
}

func columnConstraintSpecTo(item *xast.ColumnConstraintSpec) sqlast.ColumnConstraintSpec {
	if item == nil { return nil }

	if x := item.GetUniqueItem(); x != nil {
		return &sqlast.UniqueColumnSpec{
			IsPrimaryKey: x.IsPrimaryKey,
			Primary: posTo(x.Primary),
			Key: posTo(x.Key),
			Unique: posTo(x.Unique)}
	} else if x := item.GetNotNullItem(); x != nil {
		return &sqlast.NotNullColumnSpec{
			Not: posTo(x.Not),
			Null: posTo(x.Null)}
	} else if x := item.GetReferenceItem(); x != nil {
		ref := &sqlast.ReferencesColumnSpec{
			References: posTo(x.References),
			RParen: posTo(x.RParen),
			TableName: objectnameTo(x.TableName)}
		for _, column := range x.Columns {
			ref.Columns = append(ref.Columns, identTo(column).(*sqlast.Ident))
		}
		return ref
	} else {
		x := item.GetCheckItem()
		return &sqlast.CheckColumnSpec{
			Expr: binaryExprTo(x.Expr),
			Check: posTo(x.Check),
			RParen: posTo(x.RParen)}
	}
	return nil
}

func xtypeTo(item *xlight.Type) (*xast.Type, error) {
	if item == nil { return nil }

	if item.GetIntData() != nil {
		return xintTo(item.GetIntData())
	} else if item.GetSmallIntData() != nil {
		return xsmallIntTo(item.GetSmallIntData())
	} else if item.GetBigIntData() != nil {
		return xbigIntTo(item.GetBigIntData())
	} else if item.GetDecimalData() != nil {
		return xdecimalTo(item.GetDecimalData())
    } else if item.GetTimestampData() != nil {
        return xtimestampTo(item.GetTimestampData())
    } else if item.GetUUIDData() != nil {
        return xuuidTo(item.GetUUIDData())
	} else if item.GetCharData() != nil {
		return xcharTypeTo(item.GetCharData())
	} else { // GetVarcharData()
		return xvarcharTypeTo(item.GetVarcharData())
	}

	return output, nil
}

func typeTo(item *xast.Type) *xlight.Type {
	if item == nil { return nil }

	if item.GetIntData() != nil {
		return intTo(item.GetIntData())
	} else if item.GetSmallIntData() != nil {
		return smallIntTo(item.GetSmallIntData())
	} else if item.GetBigIntData() != nil {
		return bigIntTo(item.GetBigIntData())
	} else if item.GetDecimalData() != nil {
		return decimalTo(item.GetDecimalData())
    } else if item.GetTimestampData() != nil {
        return timestampTo(item.GetTimestampData())
    } else if item.GetUUIDData() != nil {
        return uuidTo(item.GetUUIDData())
	} else if item.GetCharData() != nil {
		return charTypeTo(item.GetCharData())
	} else { // GetVarcharData()
		return varcharTypeTo(item.GetVarcharData())
	}

	return nil
}

// end create table
