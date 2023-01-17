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
	if body.ElseResult != nil {
		output.ElseResult = xidentTo(body.ElseResult.(*sqlast.Ident))
	}
	for i, condition := range body.Conditions {
		item, err := xbinaryExprTo(condition.(*sqlast.BinaryExpr))
		if err != nil { return nil, err }

		resultMessage := new(xast.ResultMessage)
		switch t := body.Results[i].(type) {
        case *sqlast.Ident:
			resultMessage.ResultClause = &xast.ResultMessage_ResultIdent{ResultIdent: xidentTo(t)}
		case *sqlast.UnaryExpr:
			result, err := xunaryExprTo(t)
			if err != nil { return nil, err }
			resultMessage.ResultClause = &xast.ResultMessage_ResultUnary{ResultUnary: result}
		default:	
            return nil, fmt.Errorf("missing result type in CaseExpr %T", t)
        }

		output.Conditions = append(output.Conditions, item)
		output.Results = append(output.Results, resultMessage)
	}
	return output, nil
}

func caseExprTo(body *xast.CaseExpr) *sqlast.CaseExpr {
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
		output.Conditions = append(output.Conditions, binaryExprTo(condition))
		result := body.Results[i]
		if ident := result.GetResultIdent(); ident != nil {
			output.Results = append(output.Results, identTo(ident))
		} else {
			output.Results = append(output.Results, unaryExprTo(result.GetResultUnary()))
		}
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
