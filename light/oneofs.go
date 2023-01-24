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

/*
func xtableElementTo(item sqlast.TableElement) *xast.TableElement {
	element := new(xast.TableElement)
	switch t := item.(type) {
	case *sqlast.ColumnDef:
		x := xcolumnDefTo(t)
		element.TableElementClause = &xast.TableElement_ColumnDefElement{
			ColumnDefElement:x}
	case *sqlast.TableConstraint:
		x := xtableConstraintTo(t)
		element.TableElementClause = &xast.TableElement_TableConstraintElement{
			TableConstraintElement: x}
	default:
		return nil
	}
	return element
}

func tableElementTo(item *xast.TableElement) sqlast.TableElement {
	if item.GetColumnDefElement() != nil {
		return columnDefTo(item.GetColumnDefElement())
	}
	return tableConstraintTo(item.GetTableConstraintElement())
}

func xtableOptionTo(item sqlast.TableOption) *xast.TableOption {
	output := &xast.TableOption{}
	switch t := item.(type) {
	case *sqlast.MyEngine:
		output.TableOptionClause = &xast.TableOption_MyEngineOption{MyEngineOption:
            &xast.MyEngine{
                Engine: xposTo(t),
                Equal: t.Equal,
                Name: xidentTo(t.Name)}}
	case *sqlast.MyCharset:
		output.TableOptionClause = &xast.TableOption_MyCharsetOption{MyCharsetOption:
            &xast.MyCharset{
                IsDefault: t.IsDefault,
                Default: xposTo(),
                Charset: xposTo(),
                Equal: t.Equal,
                Name: xidentTo(t.Name)}}
	default:
		return nil
	}
	return output
}

func tableOptionTo(item *xast.TableOption) sqlast.TableOption {
    if x := item.GetMyEngineOption(); x != nil {
        return &sqlast.MyEngine{
            Equal: x.Equal,
            Name: identTo(x.Name)}
    } else if x := item.GetMyCharsetOption(); x != nil {
		return &sqlast.MyCharset{
            IsDefault: x.IsDefault,
			Equal: x.Equal,
            Name: identTo(x.Name)}
	}
	return nil
}
*/

func xvalueNodeTo(item *xlight.ValueNode) *xast.ValueNode {
	if item == nil { return nil }

	output := &xast.ValueNode{}
	if x := item.GetStringItem(); x != "" {
        output.ValueNodeClause = &xast.ValueNode_StringItem{StringItem: xstringTo(x)}
	} else if x := item.GetLongItem(); x != 0 {
        output.ValueNodeClause = &xast.ValueNode_LongItem{LongItem: xlongTo(x)}
	} else if x := item.GetDoubleItem(); x != 0 {
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

/*
func xinsertSourceTo(item sqlast.InsertSource) *xast.InsertSource {
	if item == nil { return nil, nil }

	output := &xast.InsertSource{}
    switch t := item.(type) {
    case *sqlast.SubQuerySource:
		source := XQueryTo(t.SubQuery)
		output.InsertSourceClause = &xast.InsertSource_SubItem{SubItem: &xast.SubQuerySource{SubQuery: source}}
    case *sqlast.ConstructorSource:
        source := xconstructorSourceTo(t)
        output.InsertSourceClause = &xast.InsertSource_StructorItem{StructorItem: source}
    default:
        return nil
    }

	return output
}

func insertSourceTo(item *xast.InsertSource) sqlast.InsertSource {
	if item == nil { return nil }

	if x := item.GetSubItem(); x != nil {
        return &sqlast.SubQuerySource{SubQuery: QueryTo(x.SubQuery)}
    } else if x := item.GetStructorItem(); x != nil {
        return constructorSourceTo(x)
    }
	return nil
}

func xalterTableActionTo(item sqlast.AlterTableAction) *xast.AlterTableAction {
	if item == nil { return nil }

	output := &xast.AlterTableAction{}
    switch t := item.(type) {
    case *sqlast.AddColumnTableAction:
        x := xaddColumnTableActionTo(t)
        output.AlterTableActionClause = &xast.AlterTableAction_AddColumnItem{AddColumnItem: x}
    case *sqlast.AlterColumnTableAction:
        x := xalterColumnTableActionTo(t)
        output.AlterTableActionClause = &xast.AlterTableAction_AlterColumnItem{AlterColumnItem: x}
    case *sqlast.AddConstraintTableAction:
        x := xaddConstraintTableActionTo(t)
        output.AlterTableActionClause = &xast.AlterTableAction_AddConstraintItem{AddConstraintItem: x}
    case *sqlast.DropConstraintTableAction:
        x := xdropConstraintTableActionTo(t)
        output.AlterTableActionClause = &xast.AlterTableAction_DropConstraintItem{DropConstraintItem: x}
    case *sqlast.RemoveColumnTableAction:
        x := xremoveColumnTableActionTo(t)
        output.AlterTableActionClause = &xast.AlterTableAction_RemoveColumnItem{RemoveColumnItem: x}
    default:
        return nil
    }

	return output
}

func alterTableActionTo(item *xast.AlterTableAction) sqlast.AlterTableAction {
	if item == nil { return nil }

	if x := item.GetAddColumnItem(); x != nil {
        return addColumnTableActionTo(x)
	} else if x := item.GetAlterColumnItem(); x != nil {
        return alterColumnTableActionTo(x)
    } else if x := item.GetAddConstraintItem(); x != nil {
        return addConstraintTableActionTo(x)
    } else if x := item.GetDropConstraintItem(); x != nil {
        return dropConstraintTableActionTo(x)
    } else if x := item.GetRemoveColumnItem(); x != nil {
        return removeColumnTableActionTo(x)
    }
	return nil
}

func xtableConstraintSpecTo(item sqlast.TableConstraintSpec) *xast.TableConstraintSpec {
	if item == nil { return nil }

	output := &xast.TableConstraintSpec{}
	switch t := item.(type) {
	case *sqlast.ReferentialTableConstraint:
		x := xreferentialTableConstraintTo(t)
		output.TableContraintSpecClause = &xast.TableConstraintSpec_ReferenceItem{ReferenceItem: x}
	case *sqlast.UniqueTableConstraint:
		x := xuniqueTableConstraintTo(t)
		output.TableContraintSpecClause = &xast.TableConstraintSpec_UniqueItem{UniqueItem: x}
	case *sqlast.CheckTableConstraint:
		switch s := t.Expr.(type) {
		case *sqlast.BinaryExpr:
			x := xbinaryExprTo(s)
			output.TableContraintSpecClause = &xast.TableConstraintSpec_CheckItem{
				CheckItem: &xast.CheckTableConstraint{
					Check: xposTo(),
					RParen: xposTo(t),
					Expr: x}}
		default:
			return nil
		}
	default:
		return nil
	}
	return output
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
			Expr: binaryExprTo(x.Expr)}
	}
	return nil
}

func xcolumnConstraintSpecTo(item sqlast.ColumnConstraintSpec) *xast.ColumnConstraintSpec {
	if item == nil { return nil }

    output := &xast.ColumnConstraintSpec{}
    switch t := item.(type) {
    case *sqlast.CheckColumnSpec:
		switch s := t.Expr.(type) {
		case *sqlast.BinaryExpr:
        	x := xbinaryExprTo(s)
        	output.ColumnConstraintSpecClause = &xast.ColumnConstraintSpec_CheckItem{CheckItem:
				&xast.CheckColumnSpec{
					Expr: x,
					Check: xposTo(),
					RParen: xposTo(x)}}
		default:
			return nil
		}
    case *sqlast.UniqueColumnSpec:
        output.ColumnConstraintSpecClause = &xast.ColumnConstraintSpec_UniqueItem{UniqueItem:
			&xast.UniqueColumnSpec{
				IsPrimaryKey: t.IsPrimaryKey,
				Primary: xposTo(),
				Key: xposTo(t),
				Unique: xposTo(t)}}
    case *sqlast.NotNullColumnSpec:
        output.ColumnConstraintSpecClause = &xast.ColumnConstraintSpec_NotNullItem{NotNullItem:
			&xast.NotNullColumnSpec{
				Not: xposTo(),
				Null: xposTo(t)}}
    case *sqlast.ReferencesColumnSpec:
		ref := &xast.ReferencesColumnSpec{
			References: xposTo(),
			RParen: xposTo(t),
			TableName: xobjectnameTo(t.TableName)}
		for _, column := range t.Columns {
			ref.Columns = append(ref.Columns, xidentTo(column))
		}
        output.ColumnConstraintSpecClause = &xast.ColumnConstraintSpec_ReferenceItem{ReferenceItem: ref}
    default:
        return nil
    }

    return output
}

func columnConstraintSpecTo(item *xast.ColumnConstraintSpec) sqlast.ColumnConstraintSpec {
	if item == nil { return nil }

	if x := item.GetUniqueItem(); x != nil {
		return &sqlast.UniqueColumnSpec{
			IsPrimaryKey: x.IsPrimaryKey}
	} else if x := item.GetNotNullItem(); x != nil {
		return &sqlast.NotNullColumnSpec{}
	} else if x := item.GetReferenceItem(); x != nil {
		ref := &sqlast.ReferencesColumnSpec{
			TableName: objectnameTo(x.TableName)}
		for _, column := range x.Columns {
			ref.Columns = append(ref.Columns, identTo(column))
		}
		return ref
	} else {
		x := item.GetCheckItem()
		return &sqlast.CheckColumnSpec{
			Expr: binaryExprTo(x.Expr)}
	}
	return nil
}

func xtypeTo(item sqlast.Type) *xast.Type {
	if item == nil { return nil }

    output := &xast.Type{}
	switch t := item.(type) {
	case *sqlast.Int:
		output.TypeClause = &xast.Type_IntData{IntData: xintTo(t)}
	case *sqlast.SmallInt:
		output.TypeClause = &xast.Type_SmallIntData{SmallIntData: xsmallIntTo(t)}
	case *sqlast.BigInt:
		output.TypeClause = &xast.Type_BigIntData{BigIntData: xbigIntTo(t)}
	case *sqlast.Decimal:
		output.TypeClause = &xast.Type_DecimalData{DecimalData: xdecimalTo(t)}
	case *sqlast.Timestamp:
        output.TypeClause = &xast.Type_TimestampData{TimestampData: xtimestampTo(t)}
	case *sqlast.UUID:
        output.TypeClause = &xast.Type_UUIDData{UUIDData: xuuidTo(t)}
	case *sqlast.CharType:
		output.TypeClause = &xast.Type_CharData{CharData: xcharTypeTo(t)}
	case *sqlast.VarcharType:
		output.TypeClause = &xast.Type_VarcharData{VarcharData: xvarcharTypeTo(t)}
	default:
		return nil
	}

	return output
}

func typeTo(item *xast.Type) sqlast.Type {
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
*/

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
			&xast.UnnamedSelectItem{Node: xargsNodeTo(x.Node)}}
	} else if x := item.GetAliasItem(); x != nil {
		output.SQLSelectItemClause = &xast.SQLSelectItem_AliasItem{AliasItem:
			&xast.AliasSelectItem{Expr:xargsNodeTo(x.Expr), Alias: xidentTo(x.Alias)}}
	} else if x := item.GetWildcardItem(); x != nil {
		output.SQLSelectItemClause = &xast.SQLSelectItem_WildcardItem{WildcardItem:
			&xast.QualifiedWildcardSelectItem{Prefix: xobjectnameTo(x.Prefix)}}
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
			&xlight.UnnamedSelectItem{Node: argsNodeTo(x.Node)}}
	} else if x := item.GetAliasItem(); x != nil {
		output.SQLSelectItemClause = &xlight.SQLSelectItem_AliasItem{AliasItem:
			&xlight.AliasSelectItem{Expr:argsNodeTo(x.Expr), Alias: identTo(x.Alias)}}
	} else if x := item.GetWildcardItem(); x != nil {
		output.SQLSelectItemClause = &xlight.SQLSelectItem_WildcardItem{WildcardItem:
			&xlight.QualifiedWildcardSelectItem{Prefix: objectnameTo(x.Prefix)}}
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

/*
func xalterColumnActionTo(item sqlast.AlterColumnAction) *xast.AlterColumnAction {
	if item == nil { return nil }

    output := &xast.AlterColumnAction{}
    switch t := item.(type) {
    case *sqlast.SetDefaultColumnAction:
		x := xvalueNodeTo(t.Default)
        output.AlterColumnActionClause = &xast.AlterColumnAction_SetItem{SetItem:
			&xast.SetDefaultColumnAction{
				Set: xposTo(),
				Default: x}}
    case *sqlast.DropDefaultColumnAction:
        output.AlterColumnActionClause = &xast.AlterColumnAction_DropItem{DropItem:
			&xast.DropDefaultColumnAction{
				Drop: xposTo(),
				Default: xposTo()}}
    case *sqlast.PGSetNotNullColumnAction:
        output.AlterColumnActionClause = &xast.AlterColumnAction_PGSetItem{PGSetItem:
			&xast.PGSetNotNullColumnAction{
				Set: xposTo(),
				Null: xposTo()}}
    case *sqlast.PGDropNotNullColumnAction:
        output.AlterColumnActionClause = &xast.AlterColumnAction_PGDropItem{PGDropItem:
			&xast.PGDropNotNullColumnAction{
				Drop: xposTo(),
				Null: xposTo()}}
    case *sqlast.PGAlterDataTypeColumnAction:
		x := xtypeTo(t.DataType)
        output.AlterColumnActionClause = &xast.AlterColumnAction_PGAlterItem{PGAlterItem:
			&xast.PGAlterDataTypeColumnAction{
				Type: xposTo(),
				DataType: x}}
	default:
        return nil
    }

    return output
}

func alterColumnActionTo(item *xast.AlterColumnAction) sqlast.AlterColumnAction {
	if item == nil { return nil }

	if x := item.GetSetItem(); x != nil {
		return &sqlast.SetDefaultColumnAction{
			Set: posTo(x.Set),
			Default: valueNodeTo(x.Default)}
	} else if x := item.GetDropItem(); x != nil {
		return &sqlast.DropDefaultColumnAction{
			Drop: posTo(x.Drop),
			Default: posTo(x.Default)}
	} else if x := item.GetPGSetItem(); x != nil {
		return &sqlast.PGSetNotNullColumnAction{
			Set: posTo(x.Set),
			Null: posTo(x.Null)}
	} else if x := item.GetPGDropItem(); x != nil {
		return &sqlast.PGDropNotNullColumnAction{
			Drop: posTo(x.Drop),
			Null: posTo(x.Null)}
	} else if x := item.GetPGAlterItem(); x != nil {
		return &sqlast.PGAlterDataTypeColumnAction{
			Type: posTo(x.Type),
			DataType: typeTo(x.DataType)}
	}
	return nil
}
*/
