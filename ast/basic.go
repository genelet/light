package ast

import (
	"fmt"
	"strings"
	"github.com/genelet/sqlproto/xast"

	"github.com/akito0107/xsqlparser/sqlast"
	"github.com/akito0107/xsqlparser/sqltoken"
)

func xposTo(pos sqltoken.Pos) *xast.Pos {
	return &xast.Pos{
		Line:int32(pos.Line),
		Col:int32(pos.Col)}
}

func posTo(pos *xast.Pos) sqltoken.Pos {
	return sqltoken.Pos{
		Line:int(pos.Line),
		Col:int(pos.Col)}
}

func xposplusTo(pos sqltoken.Pos) *xast.Pos {
	return &xast.Pos{
		Line:int32(pos.Line),
		Col:int32(pos.Col)+1}
}

func xidentTo(ident *sqlast.Ident) *xast.Ident {
	if ident == nil { return nil }

	return &xast.Ident{
		Value: ident.Value,
		From: xposTo(ident.From),
		To: xposTo(ident.To)}
}

func xwildcardTo(card *sqlast.Wildcard) *xast.Ident {
	if card == nil { return nil }

	return &xast.Ident{
		Value: "*",
		From: xposTo(card.Wildcard),
		To: xposplusTo(card.Wildcard)}
}

func identTo(ident *xast.Ident) sqlast.Node {
	if ident == nil { return nil }

	if ident.Value== "*" {
		return &sqlast.Wildcard{Wildcard: posTo(ident.From)}
	}
	return &sqlast.Ident{
		Value: ident.Value,
		From: posTo(ident.From),
		To: posTo(ident.To)}
}

func xidentsTo(ident *sqlast.Ident) *xast.CompoundIdent {
	if ident == nil { return nil }

	return &xast.CompoundIdent{Idents:[]*xast.Ident{xidentTo(ident)}}
}

func xwildcardsTo(card *sqlast.Wildcard) *xast.CompoundIdent {
	if card == nil { return nil }

	return &xast.CompoundIdent{Idents:[]*xast.Ident{xwildcardTo(card)}}
}

func xcompoundTo(idents *sqlast.CompoundIdent) *xast.CompoundIdent {
	if idents == nil { return nil }

	var xs []*xast.Ident
	for _, item := range idents.Idents {
		xs = append(xs, xidentTo(item))
	}
	return &xast.CompoundIdent{Idents:xs}
}

func xwildcarditemTo(t *sqlast.QualifiedWildcardSelectItem) *xast.CompoundIdent {
	if t == nil { return nil }

	comp := xobjectnameTo(t.Prefix)
	comp.Idents = append(comp.Idents, &xast.Ident{
		Value: "*",
		From: xposTo(t.Pos()),
		To: xposplusTo(t.Pos())})
	return comp
}

func compoundTo(idents *xast.CompoundIdent) sqlast.Node {
	if idents == nil { return nil }

	if len(idents.Idents) == 1 {
		return identTo(idents.Idents[0])
	}
	var xs []*sqlast.Ident
	for _, item := range idents.Idents {
		switch t := identTo(item).(type) {
		case *sqlast.Wildcard:
			return &sqlast.QualifiedWildcardSelectItem{
				Prefix: &sqlast.ObjectName{Idents:xs}}
		case *sqlast.Ident:
			xs = append(xs, t)
		default:
		}
	}
	return &sqlast.CompoundIdent{Idents:xs}
}

func xobjectnameTo(idents *sqlast.ObjectName) *xast.CompoundIdent {
	if idents == nil { return nil }

	var xs []*xast.Ident
	for _, item := range idents.Idents {
		xs = append(xs, xidentTo(item))
	}
	return &xast.CompoundIdent{Idents:xs}
}

func compoundToObjectname(idents *xast.CompoundIdent) *sqlast.ObjectName {
	if idents == nil { return nil }

	var xs []*sqlast.Ident
	for _, item := range idents.Idents {
		xs = append(xs, identTo(item).(*sqlast.Ident))
	}
	return &sqlast.ObjectName{Idents:xs}
}

func xoperatorTo(op *sqlast.Operator) *xast.Operator {
	if op == nil { return nil }

	return &xast.Operator{
		Type: xast.OperatorType(op.Type),
		From: xposTo(op.From),
		To: xposTo(op.To)}
}

func operatorTo(op *xast.Operator) *sqlast.Operator {
	if op == nil { return nil }

	return &sqlast.Operator{
		Type: sqlast.OperatorType(op.Type),
		From: posTo(op.From),
		To: posTo(op.To)}
}

func xjointypeTo(t *sqlast.JoinType) *xast.JoinType {
	if t == nil { return nil }

	return &xast.JoinType{
		Condition: xast.JoinTypeCondition(t.Condition),
		From: xposTo(t.From),
		To: xposTo(t.To)}
}

func jointypeTo(t *xast.JoinType) *sqlast.JoinType {
	if t == nil { return nil }

	return &sqlast.JoinType{
		Condition: sqlast.JoinTypeCondition(t.Condition),
		From: posTo(t.From),
		To: posTo(t.To)}
}

func xstringTo(t *sqlast.SingleQuotedString) *xast.StringUnit {
    if t == nil { return nil }

    return &xast.StringUnit{
        Value: t.String,
        From: xposTo(t.From),
        To: xposTo(t.To)}
}

func stringTo(t *xast.StringUnit) *sqlast.SingleQuotedString {
    if t == nil { return nil }

    return &sqlast.SingleQuotedString{
        String: t.Value,
        From: posTo(t.From),
        To: posTo(t.To)}
}

func xdoubleTo(t *sqlast.DoubleValue) *xast.DoubleUnit {
    if t == nil { return nil }

    return &xast.DoubleUnit{
        Value: t.Double,
        From: xposTo(t.From),
        To: xposTo(t.To)}
}

func doubleTo(t *xast.DoubleUnit) *sqlast.DoubleValue {
    if t == nil { return nil }

    return &sqlast.DoubleValue{
        Double: t.Value,
        From: posTo(t.From),
        To: posTo(t.To)}
}

func xlongTo(t *sqlast.LongValue) *xast.LongUnit {
    if t == nil { return nil }

    return &xast.LongUnit{
        Value: t.Long,
        From: xposTo(t.From),
        To: xposTo(t.To)}
}

func longTo(t *xast.LongUnit) *sqlast.LongValue {
    if t == nil { return nil }

    return &sqlast.LongValue{
        Long: t.Value,
        From: posTo(t.From),
        To: posTo(t.To)}
}

func xfunctionTo(s *sqlast.Function) *xast.AggFunction {
    if s == nil { return nil }

	name := s.Name.Idents[0]
	aggType := xast.AggType(xast.AggType_value[strings.ToUpper(name.Value)])
	var args []*xast.CompoundIdent
	for _, item := range s.Args {
		switch t := item.(type) {
		case *sqlast.Ident:
			args = append(args, xidentsTo(t))
		case *sqlast.CompoundIdent:
			args = append(args, xcompoundTo(t))
		case *sqlast.Wildcard:
			args = append(args, xwildcardsTo(t))
		default:
			args = append(args, nil)
		}
	}
	return &xast.AggFunction{
		TypeName: aggType,
		RestArgs: args,
		From: xposTo(name.From),
		To: xposTo(name.To)}
}

func functionTo(f *xast.AggFunction) *sqlast.Function {
    if f == nil { return nil }

	aggname := xast.AggType_name[int32(f.TypeName)]
	on := &sqlast.ObjectName{Idents:[]*sqlast.Ident{&sqlast.Ident{
		Value: aggname,
		From: posTo(f.From),
		To: posTo(f.To)}}}

	var c *xast.CompoundIdent
	if f.RestArgs != nil {
		c = f.RestArgs[0]
	}
	if c == nil { return &sqlast.Function{Name: on} }

	var args []sqlast.Node
	for _, item := range f.RestArgs {
		args = append(args, compoundTo(item))
	}
	return &sqlast.Function{
		Name: on,
		Args: args}
}

func xsetoperatorTo(op sqlast.SQLSetOperator) (*xast.SetOperator, error) {
    xop := &xast.SetOperator{}
    switch t := op.(type) {
    case *sqlast.UnionOperator:
        xop.Type = xast.SetOperatorType_Union
        xop.From = xposTo(t.From)
        xop.To = xposTo(t.To)
    case *sqlast.IntersectOperator:
        xop.Type = xast.SetOperatorType_Intersect
        xop.From = xposTo(t.From)
        xop.To = xposTo(t.To)
    case *sqlast.ExceptOperator:
        xop.Type = xast.SetOperatorType_Except
        xop.From = xposTo(t.From)
        xop.To = xposTo(t.To)
    default:
        return nil, fmt.Errorf("unknow set operation %#v", op)
    }
	return xop, nil
}

func setoperatorTo(op *xast.SetOperator) sqlast.SQLSetOperator {
    switch op.Type {
    case xast.SetOperatorType_Union:
        return &sqlast.UnionOperator{
            From: posTo(op.From),
            To: posTo(op.To)}
    case xast.SetOperatorType_Intersect:
        return &sqlast.IntersectOperator{
            From: posTo(op.From),
            To: posTo(op.To)}
    case xast.SetOperatorType_Except:
        return &sqlast.ExceptOperator{
            From: posTo(op.From),
            To: posTo(op.To)}
    default:
    }
	return nil
}

func xunaryTo(body *sqlast.UnaryExpr) (*xast.QueryStmt_UnaryExpr, error) {
	expr, err := xbinaryexprTo(body.Expr.(*sqlast.BinaryExpr))
	return &xast.QueryStmt_UnaryExpr{
		From: xposTo(body.From),
		Op: xoperatorTo(body.Op),
		Expr: expr}, err
}

func unaryTo(body *xast.QueryStmt_UnaryExpr) *sqlast.UnaryExpr {
	return &sqlast.UnaryExpr{
		From: posTo(body.From),
		Op: operatorTo(body.Op),
		Expr: binaryexprTo(body.Expr)}
}

func xnestedTo(body *sqlast.Nested) (*xast.QueryStmt_SQLSelect_Nested, error) {
	output := &xast.QueryStmt_SQLSelect_SQLSelectItem{}
	err := xitemToXsql(output, body.AST)
	return &xast.QueryStmt_SQLSelect_Nested{
		AST: output,
		LParen: xposTo(body.LParen),
		RParen: xposTo(body.RParen)}, err
}

func nestedTo(body *xast.QueryStmt_SQLSelect_Nested) *sqlast.Nested {
	return &sqlast.Nested{
		AST: selectitemTo(body.AST),
		LParen: posTo(body.LParen),
		RParen: posTo(body.RParen)}
}

func xcaseExprTo(body *sqlast.CaseExpr) (*xast.QueryStmt_SQLSelect_CaseExpr, error) {
	output := &xast.QueryStmt_SQLSelect_CaseExpr{
		Case: xposTo(body.Case),
		CaseEnd: xposTo(body.CaseEnd)}
	if body.Operand != nil {
		output.Operand = xoperatorTo(body.Operand.(*sqlast.Operator))
	}
	if body.ElseResult != nil {
		output.ElseResult = xidentTo(body.ElseResult.(*sqlast.Ident))
	}
	for i, condition := range body.Conditions {
		item, err := xbinaryexprTo(condition.(*sqlast.BinaryExpr))
		if err != nil { return nil, err }

		resultMessage := new(xast.QueryStmt_SQLSelect_ResultMessage)
		switch t := body.Results[i].(type) {
        case *sqlast.Ident:
			resultMessage.ResultClause = &xast.QueryStmt_SQLSelect_ResultMessage_ResultIdent{ResultIdent: xidentTo(t)}
		case *sqlast.UnaryExpr:
			result, err := xunaryTo(t)
			if err != nil { return nil, err }
			resultMessage.ResultClause = &xast.QueryStmt_SQLSelect_ResultMessage_ResultUnary{ResultUnary: result}
		default:	
            return nil, fmt.Errorf("missing result type in CaseExpr %T", t)
        }

		output.Conditions = append(output.Conditions, item)
		output.Results = append(output.Results, resultMessage)
	}
	return output, nil
}

func caseExprTo(body *xast.QueryStmt_SQLSelect_CaseExpr) *sqlast.CaseExpr {
	output := &sqlast.CaseExpr{
		Case: posTo(body.Case),
		CaseEnd: posTo(body.CaseEnd)}
	if body.Operand != nil {
		output.Operand = operatorTo(body.Operand)
	}
	if body.ElseResult != nil {
		output.ElseResult = identTo(body.ElseResult)
	}
	for i, condition := range body.Conditions {
		output.Conditions = append(output.Conditions, binaryexprTo(condition))
		result := body.Results[i]
		if ident := result.GetResultIdent(); ident != nil {
			output.Results = append(output.Results, identTo(ident))
		} else {
			output.Results = append(output.Results, unaryTo(result.GetResultUnary()))
		}
	}
	return output
}
