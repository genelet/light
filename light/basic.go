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

func xcompoundTo(idents *xlight.CompoundIdent, ns ...int) *xast.CompoundIdent {
	if idents == nil { return nil }

	output := &xast.CompoundIdent{}
	n := 0
	if ns != nil {
		n = ns[0]
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

func xstringTo(t string) *xast.StringUnit {
    return &xast.StringUnit{
        Value: t,
        From: xposTo(),
        To: xposplusTo(t)}
}

func stringTo(t *xast.StringUnit) string {
	return t.Value
}

func xdoubleTo(t float64) *xast.DoubleUnit {
    return &xast.DoubleUnit{
        Value: t,
        From: xposTo(),
        To: xposplusTo(t)}
}

func doubleTo(t *xast.DoubleUnit) float64 {
	return t.Value
}

func xlongTo(t int64) *xast.LongUnit {
    return &xast.LongUnit{
        Value: t,
        From: xposTo(),
        To: xposplusTo(t)}
}

func longTo(t *xast.LongUnit) int64 {
	return t.Value
}

func xfunctionTo(f *xlight.AggFunction) *xast.AggFunction {
	aggType := xast.AggType(f.TypeName)
	output := &xast.AggFunction{
		TypeName: aggType,
		From: xposTo(),
		To: xposplusTo(aggType)}

	n := 0
	for _, item := range f.RestArgs {
		y := xcompoundTo(item, n)
		output.RestArgs = append(output.RestArgs, y)
		n += len(y.Idents[len(y.Idents)-1].Value) + 1
	}

	return output
}

func functionTo(f *xast.AggFunction) *xlight.AggFunction {
    if f == nil { return nil }

	fl := &xlight.AggFunction{TypeName: xlight.AggType(f.TypeName)}
	for _, item := range f.RestArgs {
		fl.RestArgs = append(fl.RestArgs, compoundTo(item))
	}

	return fl
}
