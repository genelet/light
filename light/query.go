package light

import (
	"bytes"
	"fmt"

	"github.com/genelet/sqlproto/ast"
	"github.com/genelet/sqlproto/xast"
	"github.com/genelet/sqlproto/xlight"

	"github.com/akito0107/xsqlparser"
	"github.com/akito0107/xsqlparser/sqlast"
	"github.com/akito0107/xsqlparser/dialect"
)

// SQL2Proto returns protobuf message from SQL query string
//
func SQL2Proto(str string) (*xlight.QueryStmt, error) {
	parser, err := xsqlparser.NewParser(bytes.NewBufferString(str), &dialect.GenericSQLDialect{})
	if err != nil { return nil, err }
	istmt, err := parser.ParseStatement()
	if err != nil { return nil, err }
	stmt, ok := istmt.(*sqlast.QueryStmt)
	if !ok { return nil, fmt.Errorf("not ast QueryStmt") }

	xquery, err := ast.XQueryTo(stmt)
	if err != nil { return nil, err }

	return QueryTo(xquery), nil
}

// Proto2SQL returns SQL query string from protobuf message
//
func Proto2SQL(pb *xlight.QueryStmt) string {
	astPB := XQueryTo(pb)
	sqlPB := ast.QueryTo(astPB)
	return sqlPB.ToSQLString()
}

// QueryTo translates a xlight query statement into xast statement
//
func XQueryTo(stmt *xlight.QueryStmt) *xast.QueryStmt {
	output := &xast.QueryStmt{}

	for _, item := range stmt.CTEs {
		output.CTEs = append(output.CTEs, xcteTo(item))
	}

	output.Body = xsetoperationTo(stmt.Body)

	for _, item := range stmt.OrderBy {
		output.OrderBy = append(output.OrderBy, xorderbyTo(item))
	}
	if stmt.LimitExpression != nil {
		output.LimitExpression = xlimitTo(stmt.LimitExpression)
	}

	return output
}

// QueryTo translates a xast query statement into xlight statement
//
func QueryTo(stmt *xast.QueryStmt) *xlight.QueryStmt {
	output := &xlight.QueryStmt{}

	for _, item := range stmt.CTEs {
		output.CTEs = append(output.CTEs, cteTo(item))
	}

	output.Body = setoperationTo(stmt.Body)

	for _, item := range stmt.OrderBy {
		output.OrderBy = append(output.OrderBy, orderbyTo(item))
	}
	if stmt.LimitExpression != nil {
		output.LimitExpression = limitTo(stmt.LimitExpression)
	}

	return output
}

func xcteTo(cte *xlight.QueryStmt_CTE) *xast.QueryStmt_CTE {
	query := XQueryTo(cte.Query)
	return &xast.QueryStmt_CTE{
		AliasName: xidentTo(cte.AliasName),
		Query: query,
		RParen: xposTo()}
}

func cteTo(cte *xast.QueryStmt_CTE) *xlight.QueryStmt_CTE {
	output := &xlight.QueryStmt_CTE{
		Query: QueryTo(cte.Query)}
	if cte.AliasName != nil {
		output.AliasName = identTo(cte.AliasName)
	}
	return output
}

func xinsubqueryTo(sq *xlight.QueryStmt_InSubQuery) *xast.QueryStmt_InSubQuery {
	query := XQueryTo(sq.SubQuery)

	return &xast.QueryStmt_InSubQuery{
		Expr: xcompoundTo(sq.Expr),
		SubQuery: query,
		Negated: sq.Negated,
		RParen: xposTo()}
}

func insubqueryTo(sq *xast.QueryStmt_InSubQuery) *xlight.QueryStmt_InSubQuery {
	query := QueryTo(sq.SubQuery)

	return &xlight.QueryStmt_InSubQuery{
		Expr: compoundTo(sq.Expr),
		SubQuery: query,
		Negated: sq.Negated}
}

func xbinaryexprTo(binary *xlight.QueryStmt_BinaryExpr) *xast.QueryStmt_BinaryExpr {
	if binary == nil { return nil }

	item := &xast.QueryStmt_BinaryExpr{Op: xoperatorTo(binary.Op)}

	if v := binary.GetLeftIdents(); v != nil {
		item.LeftOneOf = &xast.QueryStmt_BinaryExpr_LeftIdents{LeftIdents:xcompoundTo(v)}
	} else if v := binary.GetLeftBinary(); v != nil {
		item.LeftOneOf = &xast.QueryStmt_BinaryExpr_LeftBinary{LeftBinary:xbinaryexprTo(v)}
	}

	if v := binary.GetRightIdents(); v != nil {
		item.RightOneOf = &xast.QueryStmt_BinaryExpr_RightIdents{RightIdents:xcompoundTo(v)}
	} else if v := binary.GetSingleQuotedString(); v != "" {
		item.RightOneOf = &xast.QueryStmt_BinaryExpr_SingleQuotedString{SingleQuotedString:xstringTo(v)}
	} else if v := binary.GetDoubleValue(); v != 0 {
		item.RightOneOf = &xast.QueryStmt_BinaryExpr_DoubleValue{DoubleValue:xdoubleTo(v)}
	} else if v := binary.GetLongValue(); v != 0 {
		item.RightOneOf = &xast.QueryStmt_BinaryExpr_LongValue{LongValue:xlongTo(v)}
	} else if v := binary.GetQueryValue(); v != nil {
		item.RightOneOf = &xast.QueryStmt_BinaryExpr_QueryValue{QueryValue:xinsubqueryTo(v)}
	} else if v := binary.GetRightBinary(); v != nil {
		item.RightOneOf = &xast.QueryStmt_BinaryExpr_RightBinary{RightBinary:xbinaryexprTo(v)}
	}

	return item
}

func binaryexprTo(binary *xast.QueryStmt_BinaryExpr) *xlight.QueryStmt_BinaryExpr {
	if binary == nil { return nil }

	item := &xlight.QueryStmt_BinaryExpr{Op: operatorTo(binary.Op)}

	if v := binary.GetLeftIdents(); v != nil {
		item.LeftOneOf = &xlight.QueryStmt_BinaryExpr_LeftIdents{LeftIdents:compoundTo(v)}
	} else if v := binary.GetLeftBinary(); v != nil {
		item.LeftOneOf = &xlight.QueryStmt_BinaryExpr_LeftBinary{LeftBinary:binaryexprTo(v)}
	}

	if v := binary.GetRightIdents(); v != nil {
		item.RightOneOf = &xlight.QueryStmt_BinaryExpr_RightIdents{RightIdents:compoundTo(v)}
	} else if v := binary.GetSingleQuotedString(); v != nil {
		item.RightOneOf = &xlight.QueryStmt_BinaryExpr_SingleQuotedString{SingleQuotedString:stringTo(v)}
	} else if v := binary.GetDoubleValue(); v != nil {
		item.RightOneOf = &xlight.QueryStmt_BinaryExpr_DoubleValue{DoubleValue:doubleTo(v)}
	} else if v := binary.GetLongValue(); v != nil {
		item.RightOneOf = &xlight.QueryStmt_BinaryExpr_LongValue{LongValue:longTo(v)}
	} else if v := binary.GetQueryValue(); v != nil {
		item.RightOneOf = &xlight.QueryStmt_BinaryExpr_QueryValue{QueryValue:insubqueryTo(v)}
	} else if v := binary.GetRightBinary(); v != nil {
		item.RightOneOf = &xlight.QueryStmt_BinaryExpr_RightBinary{RightBinary:binaryexprTo(v)}
	}

	return item
}

func xorderbyTo(orderby *xlight.QueryStmt_OrderByExpr) *xast.QueryStmt_OrderByExpr {
	if orderby == nil { return nil }

	return &xast.QueryStmt_OrderByExpr{
		OrderingPos: xposTo(),
		ASCBool: orderby.ASCBool,
		Expr: xcompoundTo(orderby.Expr)}
}

func orderbyTo(orderby *xast.QueryStmt_OrderByExpr) *xlight.QueryStmt_OrderByExpr {
	if orderby == nil { return nil }

	return &xlight.QueryStmt_OrderByExpr{
		ASCBool: orderby.ASCBool,
		Expr: compoundTo(orderby.Expr)}
}

func xlimitTo(limit *xlight.QueryStmt_LimitExpr) *xast.QueryStmt_LimitExpr {
	if limit == nil { return nil }

	output := &xast.QueryStmt_LimitExpr{
		AllBool: limit.AllBool,
		AllPos: xposTo(),
		Limit: xposTo(),
		LimitValue: xlongTo(limit.LimitValue)}
	if limit.OffsetValue != 0 {
		output.OffsetValue = xlongTo(limit.OffsetValue)
	}

	return output
}

func limitTo(limit *xast.QueryStmt_LimitExpr) *xlight.QueryStmt_LimitExpr {
	if limit == nil { return nil }

	output := &xlight.QueryStmt_LimitExpr{
		AllBool: limit.AllBool,
		LimitValue: longTo(limit.LimitValue)}
	if limit.OffsetValue != nil {
		output.OffsetValue = longTo(limit.OffsetValue)
	}

	return output
}
