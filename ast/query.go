package ast

import (
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

	body, err := xsetexprTo(stmt.Body)
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

	output.Body = setoperationTo(stmt.Body)

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
