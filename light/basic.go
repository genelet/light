package light

import (
	"fmt"
	"github.com/genelet/sqlproto/xast"

	"github.com/genelet/sqlproto/xlight"
)

func xposTo(x ...interface{}) *xast.Pos {
	end := 1
	if x != nil {
		end += len(fmt.Sprintf("%v", x[0]))
	}
	return &xast.Pos{
		Line: 1,
		Col: int32(end)}
}

func xposplusTo(x ...interface{}) *xast.Pos {
	y := xposTo(x...)
	y.Col += int32(1)
	return y
}

func xidentTo(ident string, n ...int) *xast.Ident {
	y := &xast.Ident{
		Value: ident,
		From: xposTo(),
		To: xposplusTo(ident)}

	if n != nil {
		y.To.Col += int32(n[0])
	}

	return y
}

func identTo(ident *xast.Ident) string {
	return ident.Value
}

func xcompoundTo(idents *xlight.CompoundIdent, init ...int) *xast.CompoundIdent {
	if idents == nil { return nil }

	output := &xast.CompoundIdent{}
	n := 0
	if init != nil {
		n = init[0]
	}
	for _, item := range idents.Idents {
		output.Idents = append(output.Idents, xidentTo(item, n))
		n += len(item) + 1
	}
	return output
}

func compoundTo(idents *xast.CompoundIdent) *xlight.CompoundIdent {
	if idents == nil { return nil }

	output := &xlight.CompoundIdent{}
	for _, item := range idents.Idents {
		output.Idents = append(output.Idents, identTo(item))
	}
	return output
}

func xobjectnameTo(idents *xlight.CompoundIdent, init ...int) *xast.ObjectName {
    if idents == nil { return nil }

	ci := xcompoundTo(idents, init...)
	return &xast.ObjectName{Idents: ci.Idents}
}

func objectnameTo(idents *xast.ObjectName) *xlight.CompoundIdent {
    if idents == nil { return nil }

	return compoundTo(&xast.CompoundIdent{Idents: idents.Idents})
}

func xoperatorTo(op xlight.OperatorType) *xast.Operator {
	return &xast.Operator{
		Type: xast.OperatorType(op),
		From: xposTo(),
		To: xposplusTo(op)}
}

func operatorTo(op *xast.Operator) xlight.OperatorType {
	return xlight.OperatorType(op.Type)
}

func xjointypeTo(t xlight.JoinTypeCondition) *xast.JoinType {
	return &xast.JoinType{
		Condition: xast.JoinTypeCondition(t),
		From: xposTo(),
		To: xposplusTo(t)}
}

func jointypeTo(t *xast.JoinType) xlight.JoinTypeCondition {
	return xlight.JoinTypeCondition(t.Condition)
}

func xstringTo(t string) *xast.SingleQuotedString {
    return &xast.SingleQuotedString{
        Value: t,
        From: xposTo(),
        To: xposplusTo(t)}
}

func stringTo(t *xast.SingleQuotedString) string {
	return t.Value
}

func xdoubleTo(t float64) *xast.DoubleValue {
    return &xast.DoubleValue{
        Value: t,
        From: xposTo(),
        To: xposplusTo(t)}
}

func doubleTo(t *xast.DoubleValue) float64 {
	return t.Value
}

func xlongTo(t int64) *xast.LongValue {
    return &xast.LongValue{
        Value: t,
        From: xposTo(),
        To: xposplusTo(t)}
}

func longTo(t *xast.LongValue) int64 {
	return t.Value
}

/*
func xintTo(t int) *xast.Int {
    return &xast.Int{
        Value: t,
        From: xposTo(),
        To: xposplusTo(t)}
}

func intTo(t *xast.Int) int64 {
	return t.Value
}

func xsmallIntTo(t int16) *xast.SmallInt {
    if t == nil { return nil }

    return &xast.SmallInt{
        From: xposTo(),
        To: xposplusTo(t),
        IsUnsigned: t.IsUnsigned,
        Unsigned: xposTo(t.Unsigned)}
}

func smallIntTo(t *xast.SmallInt) *light.SmallInt {
    if t == nil { return nil }

    return &light.SmallInt{
        From: posTo(t.From),
        To: posTo(t.To),
        IsUnsigned: t.IsUnsigned,
        Unsigned: posTo(t.Unsigned)}
}
*/

func xfunctionTo(f *xlight.AggFunction) *xast.AggFunction {
	output := &xast.AggFunction{
		TypeName: xast.AggType(f.TypeName),
		From: xposTo(),
		To: xposplusTo(f)}

	for _, item := range f.RestArgs {
		output.RestArgs = append(output.RestArgs, xargsNodeTo(item))
	}

	return output
}

func functionTo(f *xast.AggFunction) *xlight.AggFunction {
    if f == nil { return nil }

	fl := &xlight.AggFunction{
		TypeName: xlight.AggType(f.TypeName)}
	for _, item := range f.RestArgs {
		fl.RestArgs = append(fl.RestArgs, argsNodeTo(item))
	}

	return fl
}

func xsetoperatorTo(op xlight.SetOperatorType) *xast.SetOperator {
    return &xast.SetOperator{
        Type: xast.SetOperatorType(op),
        From: xposTo(),
        To: xposplusTo(op)}
}

func setoperatorTo(op *xast.SetOperator) xlight.SetOperatorType {
    return xlight.SetOperatorType(op.Type)
}

func xorderbyTo(orderby *xlight.OrderByExpr) *xast.OrderByExpr {
    if orderby == nil { return nil }

    return &xast.OrderByExpr{
		ASCBool: orderby.ASCBool,
        OrderingPos: xposTo(),
		Expr: xcompoundTo(orderby.Expr)}
}

func orderbyTo(item *xast.OrderByExpr) *xlight.OrderByExpr {
    if item == nil { return nil }

    return &xlight.OrderByExpr{
        ASCBool: item.ASCBool,
        Expr: compoundTo(item.Expr)}
}

func xlimitTo(item *xlight.LimitExpr) *xast.LimitExpr {
    if item == nil { return nil }

	v := item.LimitValue
    output := &xast.LimitExpr{
        AllBool: item.AllBool,
        AllPos: xposTo(),
        Limit: xposTo(),
        LimitValue: xlongTo(v)}
	if item.OffsetValue != 0 {
        output.OffsetValue = xlongTo(item.OffsetValue)
	}

	return output
}

func limitTo(item *xast.LimitExpr) *xlight.LimitExpr {
    if item == nil { return nil }

    output := &xlight.LimitExpr{
        AllBool: item.AllBool,
        LimitValue: longTo(item.LimitValue)}
	if item.OffsetValue != nil {
        output.OffsetValue = longTo(item.OffsetValue)
	}
	return output
}
