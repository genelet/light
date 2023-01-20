package ast

import (
	"fmt"
	"github.com/genelet/sqlproto/xast"
	"github.com/akito0107/xsqlparser/sqlast"
)

func xnestedTo(body *sqlast.Nested) (*xast.Nested, error) {
	x, err := xargsNodeTo(body.AST)
	return &xast.Nested{
		AST: x,
		LParen: xposTo(body.LParen),
		RParen: xposTo(body.RParen)}, err
}

func nestedTo(body *xast.Nested) *sqlast.Nested {
	return &sqlast.Nested{
		AST: argsNodeTo(body.AST),
		LParen: posTo(body.LParen),
		RParen: posTo(body.RParen)}
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
		selectItem, err := xsqlSelectItemTo(item)
		if err != nil { return nil, err }
		query.Projection = append(query.Projection, selectItem)
	}

	for _, item := range body.FromClause {
		from, err := xtablereferenceTo(item)
		if err != nil { return nil, err }
		query.FromClause = append(query.FromClause, from)
	}

	x, err := xwhereNodeTo(body.WhereClause)
	if err!= nil { return nil, err }
	query.WhereClause = x

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
		query.Projection = append(query.Projection, sqlSelectItemTo(item))
	}
	for _, item := range body.FromClause {
		query.FromClause = append(query.FromClause, tablereferenceTo(item))
	}

	query.WhereClause = whereNodeTo(body.WhereClause)

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
