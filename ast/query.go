package ast

import (
	"fmt"
	"github.com/genelet/protodb/xast"
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

func xinsubqueryTo(sq *sqlast.InSubQuery) (*xast.QueryStmt_InSubQuery, error) {
	query, err := XQueryTo(sq.SubQuery)
	if err != nil { return nil, err }

	output := &xast.QueryStmt_InSubQuery{
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

func insubqueryTo(sq *xast.QueryStmt_InSubQuery) *sqlast.InSubQuery {
	query := QueryTo(sq.SubQuery)

	return &sqlast.InSubQuery{
		Expr: compoundTo(sq.Expr),
		SubQuery: query,
		Negated: sq.Negated,
		RParen: posTo(sq.RParen)}
}

func xbinaryexprTo(binary *sqlast.BinaryExpr) (*xast.QueryStmt_BinaryExpr, error) {
	if binary == nil { return nil, nil }

	item := &xast.QueryStmt_BinaryExpr{Op: xoperatorTo(binary.Op)}

	switch left := binary.Left.(type) {
	case *sqlast.Ident:
		item.LeftOneOf = &xast.QueryStmt_BinaryExpr_LeftIdents{LeftIdents:xidentsTo(left)}
	case *sqlast.CompoundIdent:
		item.LeftOneOf = &xast.QueryStmt_BinaryExpr_LeftIdents{LeftIdents:xcompoundTo(left)}
	case *sqlast.BinaryExpr:
		middle, err := xbinaryexprTo(left)
		if err != nil { return nil, err }
		item.LeftOneOf = &xast.QueryStmt_BinaryExpr_LeftBinary{LeftBinary:middle}
	default:
		return nil, fmt.Errorf("left type %#v", left)
	}

	switch right := binary.Right.(type) {
	case *sqlast.Ident:
		item.RightOneOf = &xast.QueryStmt_BinaryExpr_RightIdents{RightIdents:xidentsTo(right)}
	case *sqlast.CompoundIdent:
		item.RightOneOf = &xast.QueryStmt_BinaryExpr_RightIdents{RightIdents:xcompoundTo(right)}
	case *sqlast.BinaryExpr:
		middle, err := xbinaryexprTo(right)
		if err != nil { return nil, err }
		item.RightOneOf = &xast.QueryStmt_BinaryExpr_RightBinary{RightBinary:middle}
	case *sqlast.InSubQuery:
		insub, err := xinsubqueryTo(right)
		if err != nil { return nil, err }
		item.RightOneOf = &xast.QueryStmt_BinaryExpr_QueryValue{QueryValue: insub}
	case *sqlast.LongValue:
		item.RightOneOf = &xast.QueryStmt_BinaryExpr_LongValue{LongValue:xlongTo(right)}
	case *sqlast.SingleQuotedString:
		item.RightOneOf = &xast.QueryStmt_BinaryExpr_SingleQuotedString{SingleQuotedString:xstringTo(right)}
	case *sqlast.DoubleValue:
		item.RightOneOf = &xast.QueryStmt_BinaryExpr_DoubleValue{DoubleValue:xdoubleTo(right)}
	default:
		return nil, fmt.Errorf("right type %#v", right)
	}

	return item, nil
}

func binaryexprTo(binary *xast.QueryStmt_BinaryExpr) *sqlast.BinaryExpr {
	if binary == nil { return nil }

	item := &sqlast.BinaryExpr{Op: operatorTo(binary.Op)}

	if v := binary.GetLeftIdents(); v != nil {
		item.Left = compoundTo(v)
	} else if v := binary.GetLeftBinary(); v != nil {
		item.Left = binaryexprTo(v)
	}

	if v := binary.GetRightIdents(); v != nil {
		item.Right = compoundTo(v)
	} else if v := binary.GetSingleQuotedString(); v != nil {
		item.Right = stringTo(v)
	} else if v := binary.GetDoubleValue(); v != nil {
		item.Right = doubleTo(v)
	} else if v := binary.GetLongValue(); v != nil {
		item.Right = longTo(v)
	} else if v := binary.GetQueryValue(); v != nil {
		item.Right = insubqueryTo(v)
	} else if v := binary.GetRightBinary(); v != nil {
		item.Right = binaryexprTo(v)
	}

	return item
}

func xorderbyTo(orderby *sqlast.OrderByExpr) (*xast.QueryStmt_OrderByExpr, error) {
	if orderby == nil { return nil, nil }
	output := &xast.QueryStmt_OrderByExpr{
		OrderingPos: xposTo(orderby.OrderingPos)}
	if orderby.ASC == nil {
		output.ASCBool = true
	} else {
		output.ASCBool = *orderby.ASC
	}

	switch t := orderby.Expr.(type) {
	case *sqlast.Ident:
		output.Expr = xidentsTo(t)
	case *sqlast.CompoundIdent:
		output.Expr = xcompoundTo(t)
	default:
		return nil, fmt.Errorf("order by is %#v", orderby.Expr)
	}

	return output, nil
}

func orderbyTo(orderby *xast.QueryStmt_OrderByExpr) *sqlast.OrderByExpr {
	if orderby == nil { return nil }
	return &sqlast.OrderByExpr{
		OrderingPos: posTo(orderby.OrderingPos),
		ASC: &orderby.ASCBool,
		Expr: compoundTo(orderby.Expr)}
}

func xlimitTo(limit *sqlast.LimitExpr) *xast.QueryStmt_LimitExpr {
	if limit == nil { return nil }
	return &xast.QueryStmt_LimitExpr{
		AllBool: limit.All,
		AllPos: xposTo(limit.AllPos),
		Limit: xposTo(limit.Limit),
		LimitValue: xlongTo(limit.LimitValue),
		OffsetValue: xlongTo(limit.OffsetValue)}
}

func limitTo(limit *xast.QueryStmt_LimitExpr) *sqlast.LimitExpr {
	if limit == nil { return nil }
	return &sqlast.LimitExpr{
		All: limit.AllBool,
		AllPos: posTo(limit.AllPos),
		Limit: posTo(limit.Limit),
		LimitValue: longTo(limit.LimitValue),
		OffsetValue: longTo(limit.OffsetValue)}
}
