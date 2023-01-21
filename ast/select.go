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

func xsetOperationExprTo(item *sqlast.SetOperationExpr) (*xast.SetOperationExpr, error) {
	op, err := xsetoperatorTo(item.Op)
	if err != nil { return nil, err }
	output := &xast.SetOperationExpr{Op: op, All: item.All}
	//left is never nil
	x, err := xsqlSetExprTo(item.Left)
	if err != nil { return nil, err }
	output.Left = x
	if item.Right != nil {
		x, err := xsqlSetExprTo(item.Right)	
		if err != nil { return nil, err }
		output.Right = x
	}
	return output, nil
}

func setOperationExprTo(item *xast.SetOperationExpr) *sqlast.SetOperationExpr {
	return &sqlast.SetOperationExpr{
		Op: setoperatorTo(item.Op),
		All: item.All,
		Left: sqlSetExprTo(item.Left),
		Right: sqlSetExprTo(item.Right)}
}
