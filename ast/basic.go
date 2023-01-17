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

func xintTypeTo(t *sqlast.Int) *xast.IntType {
    if t == nil { return nil }

    return &xast.IntType{
        From: xposTo(t.From),
        To: xposTo(t.To),
        IsUnsigned: t.IsUnsigned,
		Unsigned: xposTo(t.Unsigned)}
}

func intTypeTo(t *xast.IntType) *sqlast.Int {
    if t == nil { return nil }

    return &sqlast.Int{
        From: posTo(t.From),
        To: posTo(t.To),
        IsUnsigned: t.IsUnsigned,
		Unsigned: posTo(t.Unsigned)}
}

func xvarcharTypeTo(t *sqlast.VarcharType) *xast.VarcharType {
    if t == nil { return nil }

    return &xast.VarcharType{
        Size: uint32(*t.Size),
        Character: xposTo(t.Character),
        Varying: xposTo(t.Varying),
        RParen: xposTo(t.RParen)}
}

func varcharTypeTo(t *xast.VarcharType) *sqlast.VarcharType {
    if t == nil { return nil }

	size := uint(t.Size)
    return &sqlast.VarcharType{
        Size: &size,
        Character: posTo(t.Character),
        Varying: posTo(t.Varying),
        RParen: posTo(t.RParen)}
}

func xfunctionTo(s *sqlast.Function) (*xast.AggFunction, error) {
	name := s.Name.Idents[0]
	aggType := xast.AggType(xast.AggType_value[strings.ToUpper(name.Value)])
	var args []*xast.AggFunction_ArgsMessage
	for _, item := range s.Args {
		arg := new(xast.AggFunction_ArgsMessage)
		switch t := item.(type) {
		case *sqlast.Ident:
			arg.ArgsClause = &xast.AggFunction_ArgsMessage_FieldIdents{FieldIdents: xidentsTo(t)}
		case *sqlast.CompoundIdent:
			arg.ArgsClause = &xast.AggFunction_ArgsMessage_FieldIdents{FieldIdents: xcompoundTo(t)}
		case *sqlast.Wildcard:
			arg.ArgsClause = &xast.AggFunction_ArgsMessage_FieldIdents{FieldIdents: xwildcardsTo(t)}
		case *sqlast.Function:
			fieldFunction, err := xfunctionTo(t)
			if err != nil { return nil, err }
			arg.ArgsClause = &xast.AggFunction_ArgsMessage_FieldFunction{FieldFunction: fieldFunction}
		case *sqlast.CaseExpr:
			fieldCase, err := xcaseExprTo(t)
			if err != nil { return nil, err }
			arg.ArgsClause = &xast.AggFunction_ArgsMessage_FieldCase{FieldCase: fieldCase}
		default:
			return nil, fmt.Errorf("args type not found: %T", t)	
		}
		args = append(args, arg)
	}
	return &xast.AggFunction{
		TypeName: aggType,
		RestArgs: args,
		From: xposTo(name.From),
		To: xposTo(name.To)}, nil
}

func functionTo(f *xast.AggFunction) *sqlast.Function {
    if f == nil { return nil }

	aggname := xast.AggType_name[int32(f.TypeName)]
	on := &sqlast.ObjectName{Idents:[]*sqlast.Ident{&sqlast.Ident{
		Value: aggname,
		From: posTo(f.From),
		To: posTo(f.To)}}}

	var args []sqlast.Node
	for _, item := range f.RestArgs {
		var arg sqlast.Node
		if item.GetFieldIdents() != nil {
			arg = compoundTo(item.GetFieldIdents())
		} else if item.GetFieldFunction() != nil {
			arg = functionTo(item.GetFieldFunction())	
		} else if item.GetFieldCase() != nil {
			arg = caseExprTo(item.GetFieldCase())	
		}
		args = append(args, arg)
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

func xorderbyTo(orderby *sqlast.OrderByExpr) (*xast.OrderByExpr, error) {
	if orderby == nil { return nil, nil }
	output := &xast.OrderByExpr{
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

func orderbyTo(orderby *xast.OrderByExpr) *sqlast.OrderByExpr {
	if orderby == nil { return nil }
	return &sqlast.OrderByExpr{
		OrderingPos: posTo(orderby.OrderingPos),
		ASC: &orderby.ASCBool,
		Expr: compoundTo(orderby.Expr)}
}

func xlimitTo(limit *sqlast.LimitExpr) *xast.LimitExpr {
	if limit == nil { return nil }
	return &xast.LimitExpr{
		AllBool: limit.All,
		AllPos: xposTo(limit.AllPos),
		Limit: xposTo(limit.Limit),
		LimitValue: xlongTo(limit.LimitValue),
		OffsetValue: xlongTo(limit.OffsetValue)}
}

func limitTo(limit *xast.LimitExpr) *sqlast.LimitExpr {
	if limit == nil { return nil }
	return &sqlast.LimitExpr{
		All: limit.AllBool,
		AllPos: posTo(limit.AllPos),
		Limit: posTo(limit.Limit),
		LimitValue: longTo(limit.LimitValue),
		OffsetValue: longTo(limit.OffsetValue)}
}
