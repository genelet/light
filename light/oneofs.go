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


func xtableElementTo(item *xlight.TableElement) *xast.TableElement {
	element := &xast.TableElement{}
	if x := item.GetColumnDefElement(); x != nil {
		element.TableElementClause = &xast.TableElement_ColumnDefElement{ColumnDefElement: xcolumnDefTo(x)}
	} else if x := item.GetTableConstraintElement(); x != nil {
		element.TableElementClause = &xast.TableElement_TableConstraintElement{TableConstraintElement: xtableConstraintTo(x)}
	} else {
		return nil
	}

	return element
}

func tableElementTo(item *xast.TableElement) *xlight.TableElement {
	element := &xlight.TableElement{}
	if x := item.GetColumnDefElement(); x != nil {
		element.TableElementClause = &xlight.TableElement_ColumnDefElement{ColumnDefElement: columnDefTo(x)}
	} else if x := item.GetTableConstraintElement(); x != nil {
		element.TableElementClause = &xlight.TableElement_TableConstraintElement{TableConstraintElement: tableConstraintTo(x)}
	} else {
		return nil
	}

	return element
}

func xtableOptionTo(item *xlight.TableOption) *xast.TableOption {
	output := &xast.TableOption{}
    if x := item.GetMyEngineOption(); x != nil {
		output.TableOptionClause = &xast.TableOption_MyEngineOption{MyEngineOption: &xast.MyEngine{
                Engine: xposTo(x),
                Equal: x.Equal,
                Name: xidentTo(x.Name)}}
    } else if x := item.GetMyCharsetOption(); x != nil {
		output.TableOptionClause = &xast.TableOption_MyCharsetOption{MyCharsetOption: &xast.MyCharset{
                IsDefault: x.IsDefault,
                Default: xposTo(x.IsDefault),
                Charset: xposTo(x),
                Equal: x.Equal,
                Name: xidentTo(x.Name)}}
	} else {
		return nil
	}

	return output
}

func tableOptionTo(item *xast.TableOption) *xlight.TableOption {
	output := &xlight.TableOption{}
    if x := item.GetMyEngineOption(); x != nil {
		output.TableOptionClause = &xlight.TableOption_MyEngineOption{MyEngineOption: &xlight.MyEngine{
            Equal: x.Equal,
            Name: identTo(x.Name)}}
    } else if x := item.GetMyCharsetOption(); x != nil {
		output.TableOptionClause = &xlight.TableOption_MyCharsetOption{MyCharsetOption: &xlight.MyCharset{
            IsDefault: x.IsDefault,
			Equal: x.Equal,
            Name: identTo(x.Name)}}
	} else {
		return nil
	}

	return output
}

func xtableConstraintSpecTo(item *xlight.TableConstraintSpec) *xast.TableConstraintSpec {
	if item == nil { return nil }

	output := &xast.TableConstraintSpec{}
	if x := item.GetReferenceItem(); x != nil {
		output.TableContraintSpecClause = &xast.TableConstraintSpec_ReferenceItem{ReferenceItem: xreferentialTableConstraintTo(x)}
	} else if x := item.GetUniqueItem(); x != nil {
		output.TableContraintSpecClause = &xast.TableConstraintSpec_UniqueItem{UniqueItem: xuniqueTableConstraintTo(x)}
	} else if x := item.GetCheckItem(); x != nil {
		output.TableContraintSpecClause = &xast.TableConstraintSpec_CheckItem{CheckItem: xcheckTableConstraintTo(x)}
	} else {
		return nil
	}

	return output
}

func tableConstraintSpecTo(item *xast.TableConstraintSpec) *xlight.TableConstraintSpec {
	if item == nil { return nil }

	output := &xlight.TableConstraintSpec{}
	if x := item.GetReferenceItem(); x != nil {
		output.TableContraintSpecClause = &xlight.TableConstraintSpec_ReferenceItem{ReferenceItem: referentialTableConstraintTo(x)}
	} else if x := item.GetUniqueItem(); x != nil {
		output.TableContraintSpecClause = &xlight.TableConstraintSpec_UniqueItem{UniqueItem: uniqueTableConstraintTo(x)}
	} else if x := item.GetCheckItem(); x != nil {
		output.TableContraintSpecClause = &xlight.TableConstraintSpec_CheckItem{CheckItem: checkTableConstraintTo(x)}
	} else {
		return nil
	}

	return output
}

func xcolumnConstraintSpecTo(item *xlight.ColumnConstraintSpec) *xast.ColumnConstraintSpec {
	if item == nil { return nil }

    output := &xast.ColumnConstraintSpec{}
	if x := item.GetCheckItem(); x != nil {
		output.ColumnConstraintSpecClause = &xast.ColumnConstraintSpec_CheckItem{CheckItem: xcheckColumnSpecTo(x)}
	} else if x := item.GetUniqueItem(); x != nil {
		output.ColumnConstraintSpecClause = &xast.ColumnConstraintSpec_UniqueItem{UniqueItem: xuniqueColumnSpecTo(x)}
	} else if x := item.GetNotNullItem(); x != xlight.NotNullColumnSpec_NotNullColumnSpecUnknown {
        output.ColumnConstraintSpecClause = &xast.ColumnConstraintSpec_NotNullItem{NotNullItem: xnotNullColumnSpecTo(x)}
    } else if x := item.GetReferenceItem(); x != nil {
		output.ColumnConstraintSpecClause = &xast.ColumnConstraintSpec_ReferenceItem{ReferenceItem: xreferencesColumnSpecTo(x)}
	} else {
		return nil
	}

    return output
}

func columnConstraintSpecTo(item *xast.ColumnConstraintSpec) *xlight.ColumnConstraintSpec {
	if item == nil { return nil }

    output := &xlight.ColumnConstraintSpec{}
	if x := item.GetCheckItem(); x != nil {
		output.ColumnConstraintSpecClause = &xlight.ColumnConstraintSpec_CheckItem{CheckItem: checkColumnSpecTo(x)}
	} else if x := item.GetUniqueItem(); x != nil {
		output.ColumnConstraintSpecClause = &xlight.ColumnConstraintSpec_UniqueItem{UniqueItem: uniqueColumnSpecTo(x)}
	} else if x := item.GetNotNullItem(); x != nil {
        output.ColumnConstraintSpecClause = &xlight.ColumnConstraintSpec_NotNullItem{NotNullItem: notNullColumnSpecTo(x)}
    } else if x := item.GetReferenceItem(); x != nil {
		output.ColumnConstraintSpecClause = &xlight.ColumnConstraintSpec_ReferenceItem{ReferenceItem: referencesColumnSpecTo(x)}
	} else {
		return nil
	}

    return output
}

func xtypeTo(item *xlight.Type) *xast.Type {
	if item == nil { return nil }

	output := &xast.Type{}
	if x := item.GetIntData(); x != nil {
		output.TypeClause = &xast.Type_IntData{IntData: xintTo(x)}
	} else if x := item.GetSmallIntData(); x != nil {
		output.TypeClause = &xast.Type_SmallIntData{SmallIntData: xsmallIntTo(x)}
	} else if x := item.GetBigIntData(); x != nil {
		output.TypeClause = &xast.Type_BigIntData{BigIntData: xbigIntTo(x)}
	} else if x := item.GetDecimalData(); x != nil {
		output.TypeClause = &xast.Type_DecimalData{DecimalData: xdecimalTo(x)}
    } else if x := item.GetTimestampData(); x != nil {
        output.TypeClause = &xast.Type_TimestampData{TimestampData: xtimestampTo(x)}
    } else if x := item.GetUUIDData(); x != xlight.DataTypeSingle_DataTypeSingleUnknown {
        output.TypeClause = &xast.Type_UUIDData{UUIDData: xuuidTo(x)}
	} else if x := item.GetCharData(); x != nil {
		output.TypeClause = &xast.Type_CharData{CharData: xcharTypeTo(x)}
	} else if x := item.GetVarcharData(); x != nil {
		output.TypeClause = &xast.Type_VarcharData{VarcharData: xvarcharTypeTo(x)}
	} else {
		return nil
	}

	return output
}

func typeTo(item *xast.Type) *xlight.Type {
	if item == nil { return nil }

	output := &xlight.Type{}
	if x := item.GetIntData(); x != nil {
		output.TypeClause = &xlight.Type_IntData{IntData: intTo(x)}
	} else if x := item.GetSmallIntData(); x != nil {
		output.TypeClause = &xlight.Type_SmallIntData{SmallIntData: smallIntTo(x)}
	} else if x := item.GetBigIntData(); x != nil {
		output.TypeClause = &xlight.Type_BigIntData{BigIntData: bigIntTo(x)}
	} else if x := item.GetDecimalData(); x != nil {
		output.TypeClause = &xlight.Type_DecimalData{DecimalData: decimalTo(x)}
    } else if x := item.GetTimestampData(); x != nil {
        output.TypeClause = &xlight.Type_TimestampData{TimestampData: timestampTo(x)}
    } else if x := item.GetUUIDData(); x != nil {
        output.TypeClause = &xlight.Type_UUIDData{UUIDData: uuidTo(x)}
	} else if x := item.GetCharData(); x != nil {
		output.TypeClause = &xlight.Type_CharData{CharData: charTypeTo(x)}
	} else if x := item.GetVarcharData(); x != nil {
		output.TypeClause = &xlight.Type_VarcharData{VarcharData: varcharTypeTo(x)}
	} else {
		return nil
	}

	return output
}

// end create table

// insert

func xinsertSourceTo(item *xlight.InsertSource) *xast.InsertSource {
    if item == nil { return nil }

    output := &xast.InsertSource{}
	if x := item.GetSubItem(); x != nil {
		output.InsertSourceClause = &xast.InsertSource_SubItem{SubItem: xsubQuerySource(x)}
	} else if x := item.GetStructorItem(); x != nil {
		output.InsertSourceClause = &xast.InsertSource_StructorItem{StructorItem: xconstructorSourceTo(x)}
	} else {
		return nil
	}

    return output
}

func insertSourceTo(item *xast.InsertSource) *xlight.InsertSource {
    if item == nil { return nil }

    output := &xlight.InsertSource{}
	if x := item.GetSubItem(); x != nil {
		output.InsertSourceClause = &xlight.InsertSource_SubItem{SubItem: subQuerySource(x)}
	} else if x := item.GetStructorItem(); x != nil {
		output.InsertSourceClause = &xlight.InsertSource_StructorItem{StructorItem: constructorSourceTo(x)}
	} else {
		return nil
	}

    return output
}

//

// alter table

func xalterTableActionTo(item *xlight.AlterTableAction) *xast.AlterTableAction {
    if item == nil { return nil }

	output := &xast.AlterTableAction{}
    if x := item.GetAddColumnItem(); x != nil {
        output.AlterTableActionClause = &xast.AlterTableAction_AddColumnItem{AddColumnItem: xaddColumnTableActionTo(x)}
    } else if x := item.GetAlterColumnItem(); x != nil {
        output.AlterTableActionClause = &xast.AlterTableAction_AlterColumnItem{AlterColumnItem: xalterColumnTableActionTo(x)}
    } else if x := item.GetAddConstraintItem(); x != nil {
        output.AlterTableActionClause = &xast.AlterTableAction_AddConstraintItem{AddConstraintItem: xaddConstraintTableActionTo(x)}
    } else if x := item.GetDropConstraintItem(); x != nil {
        output.AlterTableActionClause = &xast.AlterTableAction_DropConstraintItem{DropConstraintItem: xdropConstraintTableActionTo(x)}
    } else if x := item.GetRemoveColumnItem(); x != nil {
        output.AlterTableActionClause = &xast.AlterTableAction_RemoveColumnItem{RemoveColumnItem: xremoveColumnTableActionTo(x)}
    } else {
    	return nil
	}

	return output
}

func alterTableActionTo(item *xast.AlterTableAction) *xlight.AlterTableAction {
    if item == nil { return nil }

	output := &xlight.AlterTableAction{}
    if x := item.GetAddColumnItem(); x != nil {
        output.AlterTableActionClause = &xlight.AlterTableAction_AddColumnItem{AddColumnItem: addColumnTableActionTo(x)}
    } else if x := item.GetAlterColumnItem(); x != nil {
        output.AlterTableActionClause = &xlight.AlterTableAction_AlterColumnItem{AlterColumnItem: alterColumnTableActionTo(x)}
    } else if x := item.GetAddConstraintItem(); x != nil {
        output.AlterTableActionClause = &xlight.AlterTableAction_AddConstraintItem{AddConstraintItem: addConstraintTableActionTo(x)}
    } else if x := item.GetDropConstraintItem(); x != nil {
        output.AlterTableActionClause = &xlight.AlterTableAction_DropConstraintItem{DropConstraintItem: dropConstraintTableActionTo(x)}
    } else if x := item.GetRemoveColumnItem(); x != nil {
        output.AlterTableActionClause = &xlight.AlterTableAction_RemoveColumnItem{RemoveColumnItem: removeColumnTableActionTo(x)}
    } else {
    	return nil
	}

	return output
}

func xalterColumnActionTo(item *xlight.AlterColumnAction) *xast.AlterColumnAction{
    if item == nil { return nil }

	output := &xast.AlterColumnAction{}
    if x := item.GetSetItem(); x != nil {
        output.AlterColumnActionClause = &xast.AlterColumnAction_SetItem{SetItem: xsetDefaultColumnActionTo(x)}
    } else if x := item.GetDropItem(); x != xlight.DropDefaultColumnAction_DropDefaultColumnActionUnknown  {
        output.AlterColumnActionClause = &xast.AlterColumnAction_DropItem{DropItem: xdropDefaultColumnActionTo(x)}
    } else if x := item.GetPGSetItem(); x != xlight.PGSetNotNullColumnAction_PGSetNotNullColumnActionUnknown {
        output.AlterColumnActionClause = &xast.AlterColumnAction_PGSetItem{PGSetItem: xpgSetNotNullColumnActionTo(x)}
    } else if x := item.GetPGDropItem(); x != xlight.PGDropNotNullColumnAction_PGDropNotNullColumnActionUnknown {
        output.AlterColumnActionClause = &xast.AlterColumnAction_PGDropItem{PGDropItem: xpgDropNotNullColumnActionTo(x)}
    } else if x := item.GetPGAlterItem(); x != nil {
        output.AlterColumnActionClause = &xast.AlterColumnAction_PGAlterItem{PGAlterItem: xpgAlterDataTypeColumnActionTo(x)}
    } else {
    	return nil
	}

	return output
}

func alterColumnActionTo(item *xast.AlterColumnAction) *xlight.AlterColumnAction {
    if item == nil { return nil }

	output := &xlight.AlterColumnAction{}
    if x := item.GetSetItem(); x != nil {
        output.AlterColumnActionClause = &xlight.AlterColumnAction_SetItem{SetItem: setDefaultColumnActionTo(x)}
    } else if x := item.GetDropItem(); x != nil {
        output.AlterColumnActionClause = &xlight.AlterColumnAction_DropItem{DropItem: dropDefaultColumnActionTo(x)}
    } else if x := item.GetPGSetItem(); x != nil {
        output.AlterColumnActionClause = &xlight.AlterColumnAction_PGSetItem{PGSetItem: pgSetNotNullColumnActionTo(x)}
    } else if x := item.GetPGDropItem(); x != nil {
        output.AlterColumnActionClause = &xlight.AlterColumnAction_PGDropItem{PGDropItem: pgDropNotNullColumnActionTo(x)}
    } else if x := item.GetPGAlterItem(); x != nil {
        output.AlterColumnActionClause = &xlight.AlterColumnAction_PGAlterItem{PGAlterItem: pgAlterDataTypeColumnActionTo(x)}
    } else {
    	return nil
	}

	return output
}

