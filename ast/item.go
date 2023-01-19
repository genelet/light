package ast

import (
	"fmt"
	"github.com/genelet/sqlproto/xast"
	"github.com/akito0107/xsqlparser/sqlast"
)

func xnestedTo(body *sqlast.Nested) (*xast.Nested, error) {
	output := &xast.SQLSelectItem{}
	err := xitemToXsql(output, body.AST)
	return &xast.Nested{
		AST: output,
		LParen: xposTo(body.LParen),
		RParen: xposTo(body.RParen)}, err
}

func nestedTo(body *xast.Nested) *sqlast.Nested {
	return &sqlast.Nested{
		AST: selectitemTo(body.AST),
		LParen: posTo(body.LParen),
		RParen: posTo(body.RParen)}
}

func xitemToXsql(selectItem *xast.SQLSelectItem, t sqlast.Node) error {
	switch s := t.(type) {
	case *sqlast.Ident:
		selectItem.SelectItemClause = &xast.SQLSelectItem_FieldIdents{FieldIdents: xidentsTo(s)}
	case *sqlast.CompoundIdent:
		selectItem.SelectItemClause = &xast.SQLSelectItem_FieldIdents{FieldIdents: xcompoundTo(s)}
	case *sqlast.Function: // single function name
		fieldFunction, err := xfunctionTo(s)
		if err != nil { return err }
		selectItem.SelectItemClause = &xast.SQLSelectItem_FieldFunction{FieldFunction: fieldFunction}
	case *sqlast.Wildcard:
		selectItem.SelectItemClause = &xast.SQLSelectItem_FieldIdents{FieldIdents: xwildcardsTo(s)}
	case *sqlast.CaseExpr:
		fieldCase, err := xcaseExprTo(s)
		if err != nil { return err }
		selectItem.SelectItemClause = &xast.SQLSelectItem_FieldCase{FieldCase: fieldCase}
	case *sqlast.Nested:
		fieldNested, err := xnestedTo(s)
		if err != nil { return err }
		selectItem.SelectItemClause = &xast.SQLSelectItem_FieldNested{FieldNested: fieldNested}
	case *sqlast.UnaryExpr:
		fieldUnary, err := xunaryExprTo(s)
		if err != nil { return err }
		selectItem.SelectItemClause = &xast.SQLSelectItem_FieldUnary{FieldUnary: fieldUnary}
	case *sqlast.BinaryExpr:
		fieldBinary, err := xbinaryExprTo(s)
		if err != nil { return err }
		selectItem.SelectItemClause = &xast.SQLSelectItem_FieldBinary{FieldBinary: fieldBinary}
	default:
		return fmt.Errorf("unknown select item %T: %#v", t, t)
	}
	return nil
}

func xselectitemTo(item sqlast.SQLSelectItem) (*xast.SQLSelectItem, error) {
	output := &xast.SQLSelectItem{}
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
		output.SelectItemClause = &xast.SQLSelectItem_FieldIdents{FieldIdents: xwildcarditemTo(t)}
	default:
		return nil, fmt.Errorf("top select item %#v", t)
	}
	return output, err
}

func selectitemTo(item *xast.SQLSelectItem) sqlast.SQLSelectItem {
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
		node = unaryExprTo(item.GetFieldUnary())
	} else if item.GetFieldBinary() != nil {
		node = binaryExprTo(item.GetFieldBinary())
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

func xinsubqueryTo(sq *sqlast.InSubQuery) (*xast.InSubQuery, error) {
	query, err := XQueryTo(sq.SubQuery)
	if err != nil { return nil, err }

	output := &xast.InSubQuery{
		SubQuery: query,
		Negated: sq.Negated,
		RParen: xposTo(sq.RParen)}

	if sq.Expr == nil { return output, nil }

	switch t := sq.Expr.(type) {
	case *sqlast.Ident:
		output.Expr = xidentsTo(t)
	case *sqlast.CompoundIdent:
		output.Expr = xcompoundTo(t)
	default:
		return nil, fmt.Errorf("expr is %#v", sq.Expr)
	}

	return output, nil
}

func insubqueryTo(sq *xast.InSubQuery) *sqlast.InSubQuery {
	query := QueryTo(sq.SubQuery)

	return &sqlast.InSubQuery{
		Expr: compoundTo(sq.Expr),
		SubQuery: query,
		Negated: sq.Negated,
		RParen: posTo(sq.RParen)}
}

func xselectTo(body *sqlast.SQLSelect) (*xast.SQLSelect, error) {
	query := &xast.SQLSelect{
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
			inQuery := &xast.SQLSelect_InQuery{InQuery: where}
			query.WhereClause = inQuery
		case *sqlast.BinaryExpr:
			where, err := xbinaryExprTo(t)
			if err != nil { return nil, err }
			binExpr := &xast.SQLSelect_BinExpr{BinExpr: where}
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
		having, err := xbinaryExprTo(havingBin)
		if err != nil { return nil, err }
		query.HavingClause = having
	}

	return query, nil
}

func selectTo(body *xast.SQLSelect) *sqlast.SQLSelect {
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
		query.WhereClause = binaryExprTo(v)
	}

	for _, item := range body.GroupByClause {
		query.GroupByClause = append(query.GroupByClause, compoundTo(item))
	}
	if body.HavingClause != nil {
		query.HavingClause = binaryExprTo(body.HavingClause)
	}

	return query
}

func sqlastplusRights(right1 sqlast.SQLSetExpr, all bool, op sqlast.SQLSetOperator, right2 sqlast.SQLSetExpr) sqlast.SQLSetExpr {
	return &sqlast.SetOperationExpr{
		Op: op,
		All: all,
		Left: right1,
		Right: right2}
}

func xsetoperationTo(body *sqlast.SetOperationExpr) (*xast.SetOperationExpr, error) {
	output := &xast.SetOperationExpr{
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
			output.RightSide = &xast.SetOperationExpr{
				Op: leftOp,
				AllBool: left.AllBool,
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

func setoperationTo(body *xast.SetOperationExpr) sqlast.SQLSetExpr {
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

func xsetexprTo(body sqlast.SQLSetExpr) (*xast.SetOperationExpr, error) {
	switch t := body.(type) {
    case *sqlast.SQLSelect:
        xbody, err := xselectTo(t)
        if err != nil { return nil, err }
        return &xast.SetOperationExpr{LeftSide: xbody}, nil
    case *sqlast.SetOperationExpr:
        return xsetoperationTo(t)
    default:
	}
	return nil, fmt.Errorf("body is %#v", body)
}