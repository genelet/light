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

	output.Body = xsqlSetExprTo(stmt.Body)

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
	output := &xlight.QueryStmt{
		Body: sqlSetExprTo(stmt.Body)}

	for _, item := range stmt.CTEs {
		output.CTEs = append(output.CTEs, cteTo(item))
	}

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

func xnestedTo(body *xlight.Nested) *xast.Nested {
    x := xargsNodeTo(body.AST)
    return &xast.Nested{
        AST: x,
        LParen: xposTo(),
        RParen: xposTo(body)}
}

func nestedTo(body *xast.Nested) *xlight.Nested {
    return &xlight.Nested{
        AST: argsNodeTo(body.AST)}
}

func xinsubqueryTo(sq *xlight.InSubQuery) *xast.InSubQuery {
	query := XQueryTo(sq.SubQuery)

	return &xast.InSubQuery{
		Expr: xcompoundTo(sq.Expr),
		SubQuery: query,
		Negated: sq.Negated,
		RParen: xposTo()}
}

func insubqueryTo(sq *xast.InSubQuery) *xlight.InSubQuery {
	query := QueryTo(sq.SubQuery)

	return &xlight.InSubQuery{
		Expr: compoundTo(sq.Expr),
		SubQuery: query,
		Negated: sq.Negated}
}

func xselectTo(body *xlight.SQLSelect) *xast.SQLSelect {
    if body == nil { return nil }
    query := &xast.SQLSelect{
        DistinctBool: body.DistinctBool,
        Select: xposTo(),
		WhereClause: xwhereNodeTo(body.WhereClause),
		HavingClause: xbinaryExprTo(body.HavingClause)}

    for _, item := range body.Projection {
        query.Projection = append(query.Projection, xsqlSelectItemTo(item))
    }
    for _, item := range body.FromClause {
        query.FromClause = append(query.FromClause, xtablereferenceTo(item))
    }
    for _, item := range body.GroupByClause {
        query.GroupByClause = append(query.GroupByClause, xcompoundTo(item))
    }

    return query
}

func selectTo(body *xast.SQLSelect) *xlight.SQLSelect {
    query := &xlight.SQLSelect{
        DistinctBool: body.DistinctBool,
		WhereClause: whereNodeTo(body.WhereClause),
        HavingClause: binaryExprTo(body.HavingClause)}

    for _, item := range body.Projection {
        query.Projection = append(query.Projection, selectItemTo(item))
    }
    for _, item := range body.FromClause {
        query.FromClause = append(query.FromClause, tablereferenceTo(item))
    }
    for _, item := range body.GroupByClause {
        query.GroupByClause = append(query.GroupByClause, compoundTo(item))
    }

    return query
}

func xsetOperationExprTo(item *xlight.SetOperationExpr) *xast.SetOperationExpr {
    output := &xast.SetOperationExpr{
		Op: xsetoperatorTo(item.Op),
		All: item.All}
    output.Left = xsqlSetExprTo(item.Left)
    if item.Right != nil {
        output.Right = xsqlSetExprTo(item.Right)
    }
    return output
}

func setOperationExprTo(item *xast.SetOperationExpr) *xlight.SetOperationExpr {
    return &xlight.SetOperationExpr{
        Op: setoperatorTo(item.Op),
        All: item.All,
        Left: sqlSetExprTo(item.Left),
        Right: sqlSetExprTo(item.Right)}
}

func xunaryExprTo(body *xlight.UnaryExpr) *xast.UnaryExpr {
    return &xast.UnaryExpr{
        From: xposTo(),
        Op: xoperatorTo(body.Op),
        Expr: xbinaryExprTo(body.Expr)}
}

func unaryExprTo(body *xast.UnaryExpr) *xlight.UnaryExpr {
    return &xlight.UnaryExpr{
        Op: operatorTo(body.Op),
        Expr: binaryExprTo(body.Expr)}
}

func xcaseExprTo(body *xlight.CaseExpr) *xast.CaseExpr {
    output := &xast.CaseExpr{
        Case: xposTo(),
        CaseEnd: xposplusTo(body),
        Operand: xoperatorTo(body.Operand),
		ElseResult: xargsNodeTo(body.ElseResult)}
    for _, condition := range body.Conditions {
        output.Conditions = append(output.Conditions, xconditionNodeTo(condition))
    }
    for _, result := range body.Results {
        output.Results = append(output.Results, xargsNodeTo(result))
    }
    return output
}

func caseExprTo(body *xast.CaseExpr) *xlight.CaseExpr {
    output := &xlight.CaseExpr{
        ElseResult: argsNodeTo(body.ElseResult),
        Operand: operatorTo(body.Operand)}
    for _, condition := range body.Conditions {
        output.Conditions = append(output.Conditions, conditionNodeTo(condition))
    }
    for _, result := range body.Results {
        output.Results = append(output.Results, argsNodeTo(result))
    }

    return output
}

func xbinaryExprTo(item *xlight.BinaryExpr) *xast.BinaryExpr {
    if item == nil { return nil }

    return &xast.BinaryExpr{
		Op: xoperatorTo(item.Op),
		Left: xargsNodeTo(item.Left),
		Right: xargsNodeTo(item.Right)}
}

func binaryExprTo(item *xast.BinaryExpr) *xlight.BinaryExpr {
    if item == nil { return nil }

    return &xlight.BinaryExpr{
        Op: operatorTo(item.Op),
        Left: argsNodeTo(item.Left),
        Right: argsNodeTo(item.Right)}
}
