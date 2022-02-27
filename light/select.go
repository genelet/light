package light

import (
	"github.com/genelet/sqlproto/xast"
	"github.com/genelet/sqlproto/xlight"
)

func xselectTo(body *xlight.QueryStmt_SQLSelect) *xast.QueryStmt_SQLSelect {
	if body == nil { return nil }
	query := &xast.QueryStmt_SQLSelect{
		DistinctBool: body.DistinctBool,
		Select: xposTo()}

	for _, item := range body.Projection {
		selectItem := xselectitemTo(item)
		query.Projection = append(query.Projection, selectItem)
	}

	for _, item := range body.FromClause {
		from := xtablereferenceTo(item)
		query.FromClause = append(query.FromClause, from)
	}

	if v := body.GetInQuery(); v != nil {
		query.WhereClause = &xast.QueryStmt_SQLSelect_InQuery{InQuery: xinsubqueryTo(v)}
	} else if v := body.GetBinExpr(); v != nil {
		query.WhereClause = &xast.QueryStmt_SQLSelect_BinExpr{BinExpr: xbinaryexprTo(v)}
	}

	for _, item := range body.GroupByClause {
		query.GroupByClause = append(query.GroupByClause, xcompoundTo(item))
	}
	if body.HavingClause != nil {
		query.HavingClause = xbinaryexprTo(body.HavingClause)
	}

	return query
}

func selectTo(body *xast.QueryStmt_SQLSelect) *xlight.QueryStmt_SQLSelect {
	query := &xlight.QueryStmt_SQLSelect{
		DistinctBool: body.DistinctBool}

	for _, item := range body.Projection {
		query.Projection = append(query.Projection, selectitemTo(item))
	}
	for _, item := range body.FromClause {
		query.FromClause = append(query.FromClause, tablereferenceTo(item))
	}

	if v := body.GetInQuery(); v != nil {
		query.WhereClause = &xlight.QueryStmt_SQLSelect_InQuery{InQuery: insubqueryTo(v)}
	} else if v := body.GetBinExpr(); v != nil {
		query.WhereClause = &xlight.QueryStmt_SQLSelect_BinExpr{BinExpr: binaryexprTo(v)}
	}

	for _, item := range body.GroupByClause {
		query.GroupByClause = append(query.GroupByClause, compoundTo(item))
	}
	if body.HavingClause != nil {
		query.HavingClause = binaryexprTo(body.HavingClause)
	}

	return query
}

func xselectitemTo(item *xlight.QueryStmt_SQLSelect_SQLSelectItem) *xast.QueryStmt_SQLSelect_SQLSelectItem {
	if item == nil { return nil }

	output := &xast.QueryStmt_SQLSelect_SQLSelectItem{
		FieldIdents: xcompoundTo(item.FieldIdents)}
	if item.AliasName != "" {
		output.AliasName = xidentTo(item.AliasName)
	}
	if item.FieldFunction != nil {
		output.FieldFunction = xfunctionTo(item.FieldFunction)
	}
	return output
}

func selectitemTo(item *xast.QueryStmt_SQLSelect_SQLSelectItem) *xlight.QueryStmt_SQLSelect_SQLSelectItem {
	if item == nil { return nil }

	output := &xlight.QueryStmt_SQLSelect_SQLSelectItem{
		FieldIdents: compoundTo(item.FieldIdents)}
	if item.AliasName != nil {
		output.AliasName = identTo(item.AliasName)
	}
	if item.FieldFunction != nil {
		output.FieldFunction = functionTo(item.FieldFunction)
	}
	return output
}

func xsetoperationTo(body *xlight.QueryStmt_SetOperationExpr) *xast.QueryStmt_SetOperationExpr {
    if body == nil { return nil }

	output := &xast.QueryStmt_SetOperationExpr{
		LeftSide: xselectTo(body.LeftSide)}
	if body.RightSide != nil {
		output.AllBool = body.AllBool
		output.Op = &xast.SetOperator{
			Type: xast.SetOperatorType(body.Op),
			From: xposTo(),
			To: xposplusTo(body.Op)}
		output.RightSide = xsetoperationTo(body.RightSide)
	}

    return output
}

func setoperationTo(body *xast.QueryStmt_SetOperationExpr) *xlight.QueryStmt_SetOperationExpr {
	if body == nil { return nil }

	output := &xlight.QueryStmt_SetOperationExpr{
		LeftSide: selectTo(body.LeftSide)}
	if body.RightSide != nil {
		output.AllBool = body.AllBool
		output.Op = xlight.SetOperatorType(body.Op.Type)
		output.RightSide = setoperationTo(body.RightSide)
	}

    return output
}
