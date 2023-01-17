package ast

import (
"github.com/k0kubun/pp/v3"
	"fmt"
	"github.com/genelet/sqlproto/xast"
	"github.com/akito0107/xsqlparser/sqlast"
)

func xselectTo(body *sqlast.SQLSelect) (*xast.QueryStmt_SQLSelect, error) {
	query := &xast.QueryStmt_SQLSelect{
		DistinctBool: body.Distinct,
		Select: xposTo(body.Select)}

	for _, item := range body.Projection {
		selectItem, err := xselectitemTo(item)
		if err != nil { return nil, err }
		query.Projection = append(query.Projection, selectItem)
	}

	for _, item := range body.FromClause {
		from, err := xtablereferenceTo(item)
		if err != nil { return nil, err }
		query.FromClause = append(query.FromClause, from)
	}

	if body.WhereClause != nil {
		switch t := body.WhereClause.(type) {
		case *sqlast.InSubQuery:
			where, err := xinsubqueryTo(t)
			if err != nil { return nil, err }
			inQuery := &xast.QueryStmt_SQLSelect_InQuery{InQuery: where}
			query.WhereClause = inQuery
		case *sqlast.BinaryExpr:
			where, err := xbinaryexprTo(t)
			if err != nil { return nil, err }
			binExpr := &xast.QueryStmt_SQLSelect_BinExpr{BinExpr: where}
			query.WhereClause = binExpr
		default:
			return nil, fmt.Errorf("'where' type %#v", t)
		}
	}

	for _, item := range body.GroupByClause {
		switch t := item.(type) {
		case *sqlast.Ident:
			query.GroupByClause = append(query.GroupByClause, xidentsTo(t))
		case *sqlast.CompoundIdent:
			query.GroupByClause = append(query.GroupByClause, xcompoundTo(t))
		default:
			return nil, fmt.Errorf("'groupby' type %#v", t)
		}
	}

	if body.HavingClause != nil {
		havingBin, ok := body.HavingClause.(*sqlast.BinaryExpr)
		if !ok {
			return nil, fmt.Errorf("'having' type %#v", body.HavingClause)
		}
		having, err := xbinaryexprTo(havingBin)
		if err != nil { return nil, err }
		query.HavingClause = having
	}

	return query, nil
}

func selectTo(body *xast.QueryStmt_SQLSelect) *sqlast.SQLSelect {
	if body == nil { return nil }
	query := &sqlast.SQLSelect{
		Distinct: body.DistinctBool,
		Select: posTo(body.Select)}

	for _, item := range body.Projection {
		query.Projection = append(query.Projection, selectitemTo(item))
	}
	for _, item := range body.FromClause {
		query.FromClause = append(query.FromClause, tablereferenceTo(item))
	}

	if v := body.GetInQuery(); v != nil {
		query.WhereClause = insubqueryTo(v)
	} else if v := body.GetBinExpr(); v != nil {
		query.WhereClause = binaryexprTo(v)
	}

	for _, item := range body.GroupByClause {
		query.GroupByClause = append(query.GroupByClause, compoundTo(item))
	}
	if body.HavingClause != nil {
		query.HavingClause = binaryexprTo(body.HavingClause)
	}

	return query
}

func xitemToXsql(selectItem *xast.QueryStmt_SQLSelect_SQLSelectItem, t sqlast.Node) error {
	switch s := t.(type) {
	case *sqlast.Ident:
		selectItem.SelectItemClause = &xast.QueryStmt_SQLSelect_SQLSelectItem_FieldIdents{FieldIdents: xidentsTo(s)}
	case *sqlast.CompoundIdent:
		selectItem.SelectItemClause = &xast.QueryStmt_SQLSelect_SQLSelectItem_FieldIdents{FieldIdents: xcompoundTo(s)}
	case *sqlast.Function: // single function name
pp.Println(s)
		selectItem.SelectItemClause = &xast.QueryStmt_SQLSelect_SQLSelectItem_FieldFunction{FieldFunction: xfunctionTo(s)}
	case *sqlast.Wildcard:
		selectItem.SelectItemClause = &xast.QueryStmt_SQLSelect_SQLSelectItem_FieldIdents{FieldIdents: xwildcardsTo(s)}
	case *sqlast.CaseExpr:
		fieldCase, err := xcaseExprTo(s)
		if err != nil { return err }
		selectItem.SelectItemClause = &xast.QueryStmt_SQLSelect_SQLSelectItem_FieldCase{FieldCase: fieldCase}
	case *sqlast.Nested:
		fieldNested, err := xnestedTo(s)
		if err != nil { return err }
		selectItem.SelectItemClause = &xast.QueryStmt_SQLSelect_SQLSelectItem_FieldNested{FieldNested: fieldNested}
	case *sqlast.UnaryExpr:
		fieldUnary, err := xunaryTo(s)
		if err != nil { return err }
		selectItem.SelectItemClause = &xast.QueryStmt_SQLSelect_SQLSelectItem_FieldUnary{FieldUnary: fieldUnary}
	case *sqlast.BinaryExpr:
		fieldBinary, err := xbinaryexprTo(s)
		if err != nil { return err }
		selectItem.SelectItemClause = &xast.QueryStmt_SQLSelect_SQLSelectItem_FieldBinary{FieldBinary: fieldBinary}
	default:
		return fmt.Errorf("unknown select item %T: %#v", t, t)
	}
	return nil
}

func xselectitemTo(item sqlast.SQLSelectItem) (*xast.QueryStmt_SQLSelect_SQLSelectItem, error) {
	output := &xast.QueryStmt_SQLSelect_SQLSelectItem{}
	var err error

	switch t := item.(type) {
	case *sqlast.UnnamedSelectItem:
		err = xitemToXsql(output, t.Node)
	case *sqlast.AliasSelectItem:
		err = xitemToXsql(output, t.Expr)
		if err == nil {
			output.AliasName = xidentTo(t.Alias)
		}
	case *sqlast.QualifiedWildcardSelectItem:
		output.SelectItemClause = &xast.QueryStmt_SQLSelect_SQLSelectItem_FieldIdents{FieldIdents: xwildcarditemTo(t)}
	default:
		return nil, fmt.Errorf("top select item %#v", t)
	}
	return output, err
}

func selectitemTo(item *xast.QueryStmt_SQLSelect_SQLSelectItem) sqlast.SQLSelectItem {
	if item == nil { return nil }

	var node sqlast.Node
	if item.GetFieldFunction() != nil {
		node = functionTo(item.GetFieldFunction())
	} else if item.GetFieldCase() != nil {
		node = caseExprTo(item.GetFieldCase())
	} else if item.GetFieldIdents() != nil {
		node = compoundTo(item.GetFieldIdents())
	} else if item.GetFieldNested() != nil {
		node = nestedTo(item.GetFieldNested())
	} else if item.GetFieldCase() != nil {
		node = caseExprTo(item.GetFieldCase())
	} else if item.GetFieldUnary() != nil {
		node = unaryTo(item.GetFieldUnary())
	} else if item.GetFieldBinary() != nil {
		node = binaryexprTo(item.GetFieldBinary())
	} else {
		return nil
	}

	if item.AliasName != nil {
		return &sqlast.AliasSelectItem{
			Alias: identTo(item.AliasName).(*sqlast.Ident),
			Expr: node}
	}
	return &sqlast.UnnamedSelectItem{Node: node}
}

func sqlastplusRights(right1 sqlast.SQLSetExpr, all bool, op sqlast.SQLSetOperator, right2 sqlast.SQLSetExpr) sqlast.SQLSetExpr {
	return &sqlast.SetOperationExpr{
		Op: op,
		All: all,
		Left: right1,
		Right: right2}
}

func xsetoperationTo(body *sqlast.SetOperationExpr) (*xast.QueryStmt_SetOperationExpr, error) {
	output := &xast.QueryStmt_SetOperationExpr{
		AllBool: body.All}
	op, err := xsetoperatorTo(body.Op)
	if err != nil { return nil, err }
	output.Op = op

	// body.Left is never nil, sqlast.SQLSetExpr
	switch t := body.Left.(type) {
	case *sqlast.SQLSelect:
		left, err := xselectTo(t)
		if err != nil { return nil, err }
		output.LeftSide = left
		right, err := xsetexprTo(body.Right)
		if err != nil { return nil, err }
		output.RightSide = right
	case *sqlast.SetOperationExpr:
		left, err := xsetoperationTo(t)
		if err != nil { return nil, err }

		if left.RightSide == nil {
			bodyRight, err := xsetexprTo(body.Right)
			if err != nil { return nil, err }
			leftOp, err := xsetoperatorTo(t.Op)
			if err != nil { return nil, err }
			output.RightSide = &xast.QueryStmt_SetOperationExpr{
				AllBool: left.AllBool,
				Op: leftOp,
				LeftSide: left.LeftSide,
				RightSide: bodyRight}
		} else {
			newBody := &sqlast.SetOperationExpr{
				Op: t.Op,
				All: t.All,
				Left: t.Left,
				Right: sqlastplusRights(t.Right, body.All, body.Op, body.Right)}
			return xsetoperationTo(newBody)
		}
	default:
		return nil, fmt.Errorf("setexpr left is %#v", body.Left)
	}

	return output, nil
}

func setoperationTo(body *xast.QueryStmt_SetOperationExpr) sqlast.SQLSetExpr {
	if body == nil { return nil }

	if body.RightSide == nil {
		return selectTo(body.LeftSide)
	}

	return &sqlast.SetOperationExpr{
		All: body.AllBool,
		Op: setoperatorTo(body.Op),
		Left: selectTo(body.LeftSide),
		Right: setoperationTo(body.RightSide)}
}

func xsetexprTo(body sqlast.SQLSetExpr) (*xast.QueryStmt_SetOperationExpr, error) {
	switch t := body.(type) {
    case *sqlast.SQLSelect:
        xbody, err := xselectTo(t)
        if err != nil { return nil, err }
        return &xast.QueryStmt_SetOperationExpr{LeftSide: xbody}, nil
    case *sqlast.SetOperationExpr:
        return xsetoperationTo(t)
    default:
	}
	return nil, fmt.Errorf("body is %#v", body)
}
