package ast

import (
	"fmt"
	"github.com/genelet/sqlproto/xast"
	"github.com/akito0107/xsqlparser/sqlast"
)

// XQueryTo translates a xsqlparser query statement into xast statement
//
// see https://github.com/akito0107/xsqlparser for xsqlparser
//
func XQueryTo(stmt *sqlast.QueryStmt) (*xast.QueryStmt, error) {
	output := &xast.QueryStmt{
		With: xposTo(stmt.With)}
	for _, item := range stmt.CTEs {
		v, err := xcteTo(item)
		if err != nil { return nil, err }
		output.CTEs = append(output.CTEs, v)
	}

	body, err := xsqlSetExprTo(stmt.Body)
	if err != nil { return nil, err }
	output.Body = body

	for _, item := range stmt.OrderBy {
		v, err := xorderbyTo(item)
		if err != nil { return nil, err }
		output.OrderBy = append(output.OrderBy, v)
	}

	if stmt.Limit != nil {
		output.LimitExpression = xlimitTo(stmt.Limit)
	}

	return output, nil
}

// QueryTo translates a xast query statement into xsqlparser statement
//
func QueryTo(stmt *xast.QueryStmt) *sqlast.QueryStmt {
	output := &sqlast.QueryStmt{}

	if stmt.With != nil {
		output.With = posTo(stmt.With)
	}
	for _, item := range stmt.CTEs {
		output.CTEs = append(output.CTEs, cteTo(item))
	}

	output.Body = sqlSetExprTo(stmt.Body)

	for _, item := range stmt.OrderBy {
		output.OrderBy = append(output.OrderBy, orderbyTo(item))
	}
	if stmt.LimitExpression != nil {
		output.Limit = limitTo(stmt.LimitExpression)
	}

	return output
}

func xcteTo(cte *sqlast.CTE) (*xast.QueryStmt_CTE, error) {
	query, err := XQueryTo(cte.Query)
	if err != nil { return nil, err }
	return &xast.QueryStmt_CTE{
		AliasName: xidentTo(cte.Alias),
		Query: query,
		RParen: xposTo(cte.RParen)}, nil
}

func cteTo(cte *xast.QueryStmt_CTE) *sqlast.CTE {
	output := &sqlast.CTE{
		Query: QueryTo(cte.Query),
		RParen: posTo(cte.RParen)}
	if cte.AliasName != nil {
		output.Alias = identTo(cte.AliasName).(*sqlast.Ident)
	}
	return output
}

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
		from, err := xtableReferenceTo(item)
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
		query.FromClause = append(query.FromClause, tableReferenceTo(item))
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

func xunaryExprTo(body *sqlast.UnaryExpr) (*xast.UnaryExpr, error) {
	expr, err := xbinaryExprTo(body.Expr.(*sqlast.BinaryExpr))
	return &xast.UnaryExpr{
		From: xposTo(body.From),
		Op: xoperatorTo(body.Op),
		Expr: expr}, err
}

func unaryExprTo(body *xast.UnaryExpr) *sqlast.UnaryExpr {
	return &sqlast.UnaryExpr{
		From: posTo(body.From),
		Op: operatorTo(body.Op),
		Expr: binaryExprTo(body.Expr)}
}

func xcaseExprTo(body *sqlast.CaseExpr) (*xast.CaseExpr, error) {
	output := &xast.CaseExpr{
		Case: xposTo(body.Case),
		CaseEnd: xposTo(body.CaseEnd)}
	if body.Operand != nil {
		output.Operand = xoperatorTo(body.Operand.(*sqlast.Operator))
	}
	x, err := xargsNodeTo(body.ElseResult)
	if err != nil { return nil, err }
	output.ElseResult = x
	for _, condition := range body.Conditions {
		x, err := xconditionNodeTo(condition)
		if err != nil { return nil, err }
		output.Conditions = append(output.Conditions, x)
	}
	for _, result := range body.Results {
		x, err := xargsNodeTo(result)
		if err != nil { return nil, err }
		output.Results = append(output.Results, x)
	}
	return output, nil
}

func caseExprTo(body *xast.CaseExpr) *sqlast.CaseExpr {
	output := &sqlast.CaseExpr{
		Case: posTo(body.Case),
		CaseEnd: posTo(body.CaseEnd),
		ElseResult: argsNodeTo(body.ElseResult)}
	if body.Operand != nil {
		output.Operand = operatorTo(body.Operand)
	}
	for _, condition := range body.Conditions {
		output.Conditions = append(output.Conditions, conditionNodeTo(condition))
	}
	for _, result := range body.Results {
		output.Results = append(output.Results, argsNodeTo(result))
	}

	return output
}

func xbinaryExprTo(item *sqlast.BinaryExpr) (*xast.BinaryExpr, error) {
	if item == nil { return nil, nil }

	output := &xast.BinaryExpr{Op: xoperatorTo(item.Op)}
	x, err := xargsNodeTo(item.Left)
	if err != nil { return nil, err }
	output.Left = x

	if item.Right != nil {
		x, err := xargsNodeTo(item.Right)
		if err != nil { return nil, err }
		output.Right = x
	}
	return output, nil
}

func binaryExprTo(item *xast.BinaryExpr) *sqlast.BinaryExpr {
	if item == nil { return nil }

	return &sqlast.BinaryExpr{
		Op: operatorTo(item.Op),
		Left: argsNodeTo(item.Left),
		Right: argsNodeTo(item.Right)}
}
