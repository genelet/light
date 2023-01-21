package ast

import (
	"fmt"
	"github.com/genelet/sqlproto/xast"
	"github.com/akito0107/xsqlparser/sqlast"
)

func xwhereNodeTo(item sqlast.Node) (*xast.WhereNode, error ) {
	if item == nil { return nil, nil }

	output := &xast.WhereNode{}
    switch t := item.(type) {
    case *sqlast.InSubQuery:
        where, err := xinsubqueryTo(t)
        if err != nil { return nil, err }
        output.WhereNodeClause = &xast.WhereNode_InQuery{InQuery: where}
    case *sqlast.BinaryExpr:
        where, err := xbinaryExprTo(t)
        if err != nil { return nil, err }
        output.WhereNodeClause = &xast.WhereNode_BinExpr{BinExpr: where}
    default:
        return nil, fmt.Errorf("missing where type %T", t)
    }

	return output, nil
}

func whereNodeTo(item *xast.WhereNode) sqlast.Node {
	if item == nil { return nil }

	if x := item.GetInQuery(); x != nil {
        return insubqueryTo(x)
    } else if x := item.GetBinExpr(); x != nil {
        return binaryExprTo(x)
    }
	return nil
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

func xvalueNodeTo(item sqlast.Node) (*xast.ValueNode, error ) {
	if item == nil { return nil, nil }

	output := &xast.ValueNode{}
    switch t := item.(type) {
    case *sqlast.SingleQuotedString:
        output.ValueNodeClause = &xast.ValueNode_StringItem{StringItem: xstringTo(t)}
    case *sqlast.LongValue:
        output.ValueNodeClause = &xast.ValueNode_LongItem{LongItem: xlongTo(t)}
    case *sqlast.DoubleValue:
        output.ValueNodeClause = &xast.ValueNode_DoubleItem{DoubleItem: xdoubleTo(t)}
    case *sqlast.Ident:
        output.ValueNodeClause = &xast.ValueNode_CompoundItem{CompoundItem: xidentsTo(t)}
	case *sqlast.CompoundIdent:
		output.ValueNodeClause = &xast.ValueNode_CompoundItem{CompoundItem: xcompoundTo(t)}
	case *sqlast.Wildcard:
		output.ValueNodeClause = &xast.ValueNode_CompoundItem{CompoundItem: xwildcardsTo(t)}
    default:
        return nil, fmt.Errorf("missing value item type %T", t)
    }

	return output, nil
}

func valueNodeTo(item *xast.ValueNode) sqlast.Node {
	if item == nil { return nil }

	if x := item.GetLongItem(); x != nil {
		return longTo(x)
	} else if x := item.GetDoubleItem(); x != nil {
		return doubleTo(x)
	} else if x := item.GetCompoundItem(); x != nil {
		return compoundTo(x)
	} else if x := item.GetStringItem(); x != nil {
		return stringTo(x)
	}

	return nil
}

func xinsertSourceTo(item sqlast.InsertSource) (*xast.InsertSource, error ) {
	if item == nil { return nil, nil }

	output := &xast.InsertSource{}
    switch t := item.(type) {
    case *sqlast.SubQuerySource:
		// definition sqlast.SubQuerySource{SubQuery: q}
		source, err := XQueryTo(t.SubQuery)
		if err != nil { return nil, err }
		output.InsertSourceClause = &xast.InsertSource_SubItem{SubItem: &xast.SubQuerySource{SubQuery: source}}
    case *sqlast.ConstructorSource:
        source, err := xconstructorSourceTo(t)
        if err != nil { return nil, err }
        output.InsertSourceClause = &xast.InsertSource_StructorItem{StructorItem: source}
    default:
        return nil, fmt.Errorf("missing source type %T", t)
    }

	return output, nil
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

func xalterTableActionTo(item sqlast.AlterTableAction) (*xast.AlterTableAction, error) {
	if item == nil { return nil, nil }

	output := &xast.AlterTableAction{}
    switch t := item.(type) {
    case *sqlast.AddColumnTableAction:
        x, err := xaddColumnTableActionTo(t)
        if err != nil { return nil, err }
        output.AlterTableActionClause = &xast.AlterTableAction_AddColumnItem{AddColumnItem: x}
    case *sqlast.AddConstraintTableAction:
        x, err := xaddConstraintTableActionTo(t)
        if err != nil { return nil, err }
        output.AlterTableActionClause = &xast.AlterTableAction_AddConstraintItem{AddConstraintItem: x}
    case *sqlast.DropConstraintTableAction:
        x, err := xdropConstraintTableActionTo(t)
        if err != nil { return nil, err }
        output.AlterTableActionClause = &xast.AlterTableAction_DropConstraintItem{DropConstraintItem: x}
    case *sqlast.RemoveColumnTableAction:
        x, err := xremoveColumnTableActionTo(t)
        if err != nil { return nil, err }
        output.AlterTableActionClause = &xast.AlterTableAction_RemoveColumnItem{RemoveColumnItem: x}
    default:
        return nil, fmt.Errorf("missing actio node type %T", t)
    }

	return output, nil
}

func alterTableActionTo(item *xast.AlterTableAction) sqlast.AlterTableAction {
	if item == nil { return nil }

	if x := item.GetAddColumnItem(); x != nil {
        return addColumnTableActionTo(x)
    } else if x := item.GetAddConstraintItem(); x != nil {
        return addConstraintTableActionTo(x)
    } else if x := item.GetDropConstraintItem(); x != nil {
        return dropConstraintTableActionTo(x)
    } else if x := item.GetRemoveColumnItem(); x != nil {
        return removeColumnTableActionTo(x)
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

func xtypeTo(item sqlast.Type) (*xast.Type, error) {
	if item == nil { return nil, nil }

    output := &xast.Type{}
	switch t := item.(type) {
	case *sqlast.Int:
		output.TypeClause = &xast.Type_IntData{IntData: xintTo(t)}
	case *sqlast.SmallInt:
		output.TypeClause = &xast.Type_SmallIntData{SmallIntData: xsmallIntTo(t)}
	case *sqlast.Timestamp:
        output.TypeClause = &xast.Type_TimestampData{TimestampData: xtimestampTo(t)}
	case *sqlast.UUID:
        output.TypeClause = &xast.Type_UUIDData{UUIDData: xuuidTo(t)}
	case *sqlast.CharType:
		output.TypeClause = &xast.Type_CharData{CharData: xcharTypeTo(t)}
	case *sqlast.VarcharType:
		output.TypeClause = &xast.Type_VarcharData{VarcharData: xvarcharTypeTo(t)}
	default:
		return nil, fmt.Errorf("missing column def type: %T", t)
	}

	return output, nil
}

func typeTo(item *xast.Type) sqlast.Type {
	if item == nil { return nil }

	if item.GetIntData() != nil {
		return intTo(item.GetIntData())
	} else if item.GetSmallIntData() != nil {
		return smallIntTo(item.GetSmallIntData())
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

func xconditionNodeTo(item sqlast.Node) (*xast.ConditionNode, error) {
	if item == nil { return nil, nil }

	output := &xast.ConditionNode{}
	switch t := item.(type) {
	case *sqlast.BinaryExpr:
		x, err := xbinaryExprTo(t)
		if err != nil { return nil, err }
		output.ConditionNodeClause = &xast.ConditionNode_BinaryItem{BinaryItem: x}
	default:	
		return nil, fmt.Errorf("missing condition type in CaseExpr %T", t)
	}

	return output, nil
}

func conditionNodeTo(item *xast.ConditionNode) sqlast.Node {
	if item == nil { return nil }

	if x := item.GetBinaryItem(); x != nil {
		return binaryExprTo(x)
	}
	return nil
}

func xargsNodeTo(item sqlast.Node) (*xast.ArgsNode, error) {
	if item == nil { return nil, nil }

	output := &xast.ArgsNode{}
	switch t := item.(type) {
	case *sqlast.Ident, *sqlast.CompoundIdent, *sqlast.Wildcard, *sqlast.SingleQuotedString, *sqlast.LongValue, *sqlast.DoubleValue:
		x, err := xvalueNodeTo(item)
		if err != nil { return nil, err }
		output.ArgsNodeClause = &xast.ArgsNode_ValueItem{ValueItem: x}
	case *sqlast.BinaryExpr, *sqlast.InSubQuery:
		x, err := xwhereNodeTo(t)
		if err != nil { return nil, err }
		output.ArgsNodeClause = &xast.ArgsNode_WhereItem{WhereItem: x}
	case *sqlast.Function:
		x, err := xfunctionTo(t)
		if err != nil { return nil, err }
		output.ArgsNodeClause = &xast.ArgsNode_FunctionItem{FunctionItem: x}
	case *sqlast.CaseExpr:
		x, err := xcaseExprTo(t)
		if err != nil { return nil, err }
		output.ArgsNodeClause = &xast.ArgsNode_CaseItem{CaseItem: x}
	case *sqlast.Nested:
		x, err := xnestedTo(t)
		if err != nil { return nil, err }
		output.ArgsNodeClause = &xast.ArgsNode_NestedItem{NestedItem: x}
	case *sqlast.UnaryExpr:
		x, err := xunaryExprTo(t)
		if err != nil { return nil, err }
		output.ArgsNodeClause = &xast.ArgsNode_UnaryItem{UnaryItem: x}
	default:	
		return nil, fmt.Errorf("missing args type in args node %T", t)
	}

	return output, nil
}

func argsNodeTo(item *xast.ArgsNode) sqlast.Node {
	if item == nil { return nil }

	if x := item.GetValueItem(); x != nil {
		return valueNodeTo(x)
	} else if x := item.GetFunctionItem(); x != nil {
		return functionTo(x)	
	} else if x := item.GetCaseItem(); x != nil {
		return caseExprTo(x)	
	} else if x := item.GetNestedItem(); x != nil {
		return nestedTo(x)	
	} else if x := item.GetUnaryItem(); x != nil {
		return unaryExprTo(x)	
	} else if x := item.GetWhereItem(); x != nil {
		return whereNodeTo(x)	
	}

	return nil
}

func xsqlSelectItemTo(item sqlast.SQLSelectItem) (*xast.SQLSelectItem, error) {
	if item == nil { return nil, nil }

	output := &xast.SQLSelectItem{}
	switch t := item.(type) {
	case *sqlast.UnnamedSelectItem:
		x, err := xargsNodeTo(t.Node)
		if err != nil { return nil, err }
		output.SQLSelectItemClause = &xast.SQLSelectItem_UnnamedItem{UnnamedItem:
			&xast.UnnamedSelectItem{Node:x}}
	case *sqlast.AliasSelectItem:
		x, err := xargsNodeTo(t.Expr)
		if err != nil { return nil, err }
		output.SQLSelectItemClause = &xast.SQLSelectItem_AliasItem{AliasItem:
			&xast.AliasSelectItem{Expr:x, Alias: xidentTo(t.Alias)}}
	case *sqlast.QualifiedWildcardSelectItem:
		output.SQLSelectItemClause = &xast.SQLSelectItem_WildcardItem{WildcardItem:
			&xast.QualifiedWildcardSelectItem{Prefix: xobjectnameTo(t.Prefix)}}
	default:
		return nil, fmt.Errorf("missing select item type %T", t)
	}

	return output, nil
}

func sqlSelectItemTo(item *xast.SQLSelectItem) sqlast.SQLSelectItem {
	if item == nil { return nil }

	if x := item.GetUnnamedItem(); x != nil {
		return &sqlast.UnnamedSelectItem{Node: argsNodeTo(x.Node)}
	} else if x := item.GetAliasItem(); x != nil {
		return &sqlast.AliasSelectItem{Expr: argsNodeTo(x.Expr), Alias: identTo(x.Alias).(*sqlast.Ident)}
	} else if x := item.GetWildcardItem(); x != nil {
		return &sqlast.QualifiedWildcardSelectItem{Prefix: objectnameTo(x.Prefix)}
	}

	return nil
}

func xsqlSetExprTo(item sqlast.SQLSetExpr) (*xast.SQLSetExpr, error) {
	if item == nil { return nil, nil }

	output := &xast.SQLSetExpr{}
	switch t := item.(type) {
    case *sqlast.SQLSelect:
        x, err := xselectTo(t)
        if err != nil { return nil, err }
		output.SQLSetExprClause = &xast.SQLSetExpr_SelectItem{SelectItem: x}
    case *sqlast.SetOperationExpr:
		x, err := xsetOperationExprTo(t)
        if err != nil { return nil, err }
		output.SQLSetExprClause = &xast.SQLSetExpr_ExprItem{ExprItem: x}
    default:
    	return nil, fmt.Errorf("missing set expr type  %T", t)
    }
    return output, nil
}

func sqlSetExprTo(item *xast.SQLSetExpr) sqlast.SQLSetExpr {
    if item == nil { return nil }

    if x := item.GetSelectItem(); x != nil {
		return selectTo(x)
    } else if x := item.GetExprItem(); x != nil {
		return setOperationExprTo(x)
    }
    return nil
}
