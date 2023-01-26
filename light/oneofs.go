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
