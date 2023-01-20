package ast

import (
	"fmt"
	"github.com/genelet/sqlproto/xast"
	"github.com/akito0107/xsqlparser/sqlast"
)

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

func xbinaryExprTo(binary *sqlast.BinaryExpr) (*xast.BinaryExpr, error) {
	if binary == nil { return nil, nil }

	item := &xast.BinaryExpr{Op: xoperatorTo(binary.Op)}

	switch left := binary.Left.(type) {
	case *sqlast.Ident:
		item.LeftOneOf = &xast.BinaryExpr_LeftIdents{LeftIdents:xidentsTo(left)}
	case *sqlast.CompoundIdent:
		item.LeftOneOf = &xast.BinaryExpr_LeftIdents{LeftIdents:xcompoundTo(left)}
	case *sqlast.BinaryExpr:
		middle, err := xbinaryExprTo(left)
		if err != nil { return nil, err }
		item.LeftOneOf = &xast.BinaryExpr_LeftBinary{LeftBinary:middle}
	case *sqlast.LongValue:
		item.LeftOneOf = &xast.BinaryExpr_LeftLong{LeftLong:xlongTo(left)}
	case *sqlast.SingleQuotedString:
		item.LeftOneOf = &xast.BinaryExpr_LeftSingleQuoted{LeftSingleQuoted:xstringTo(left)}
	case *sqlast.DoubleValue:
		item.LeftOneOf = &xast.BinaryExpr_LeftDouble{LeftDouble:xdoubleTo(left)}
	default:
		return nil, fmt.Errorf("left type %#v", left)
	}

	switch right := binary.Right.(type) {
	case *sqlast.Ident:
		item.RightOneOf = &xast.BinaryExpr_RightIdents{RightIdents:xidentsTo(right)}
	case *sqlast.CompoundIdent:
		item.RightOneOf = &xast.BinaryExpr_RightIdents{RightIdents:xcompoundTo(right)}
	case *sqlast.BinaryExpr:
		middle, err := xbinaryExprTo(right)
		if err != nil { return nil, err }
		item.RightOneOf = &xast.BinaryExpr_RightBinary{RightBinary:middle}
	case *sqlast.InSubQuery:
		insub, err := xinsubqueryTo(right)
		if err != nil { return nil, err }
		item.RightOneOf = &xast.BinaryExpr_QueryValue{QueryValue: insub}
	case *sqlast.LongValue:
		item.RightOneOf = &xast.BinaryExpr_LongValue{LongValue:xlongTo(right)}
	case *sqlast.SingleQuotedString:
		item.RightOneOf = &xast.BinaryExpr_SingleQuotedString{SingleQuotedString:xstringTo(right)}
	case *sqlast.DoubleValue:
		item.RightOneOf = &xast.BinaryExpr_DoubleValue{DoubleValue:xdoubleTo(right)}
	default:
		return nil, fmt.Errorf("right type %#v", right)
	}

	return item, nil
}

func binaryExprTo(binary *xast.BinaryExpr) *sqlast.BinaryExpr {
	if binary == nil { return nil }

	item := &sqlast.BinaryExpr{Op: operatorTo(binary.Op)}

	if v := binary.GetLeftIdents(); v != nil {
		item.Left = compoundTo(v)
	} else if v := binary.GetLeftBinary(); v != nil {
		item.Left = binaryExprTo(v)
	} else if v := binary.GetLeftSingleQuoted(); v != nil {
		item.Left = stringTo(v)
	} else if v := binary.GetLeftDouble(); v != nil {
		item.Left = doubleTo(v)
	} else if v := binary.GetLeftLong(); v != nil {
		item.Left = longTo(v)
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
		item.Right = binaryExprTo(v)
	}

	return item
}
