package light

import (
//	"fmt"
	"github.com/genelet/protodb/xast"
	"github.com/genelet/protodb/xlight"
)

func xtableTo(t *xlight.QueryStmt_SQLSelect_QualifiedJoin) *xast.QueryStmt_SQLSelect_QualifiedJoin {
	output := &xast.QueryStmt_SQLSelect_QualifiedJoin {
		Name: xcompoundTo(t.Name)}
	if t.AliasName != "" {
		output.AliasName = xidentTo(t.AliasName)
	}
	return output
}

func tableTo(t *xast.QueryStmt_SQLSelect_QualifiedJoin) *xlight.QueryStmt_SQLSelect_QualifiedJoin {
	output := &xlight.QueryStmt_SQLSelect_QualifiedJoin{
		Name: compoundTo(t.Name)}
	if t.AliasName != nil {
		output.AliasName = identTo(t.AliasName)
	}

	return output
}

func xqualifiedjoinTo(item *xlight.QueryStmt_SQLSelect_QualifiedJoin) *xast.QueryStmt_SQLSelect_QualifiedJoin {
	// thisLeft is never nil
	thisLeft := item.LeftElement
	var ref *xast.QueryStmt_SQLSelect_QualifiedJoin
	if thisLeft.LeftElement != nil {
		ref = xqualifiedjoinTo(thisLeft)
	} else {
		ref = xtableTo(thisLeft)
	}

	output := xtableTo(item)
	output.LeftElement = ref
	output.TypeCondition = &xast.JoinType{
		Condition: xast.JoinTypeCondition(item.TypeCondition),
		From: xposTo(),
		To: xposplusTo(item.TypeCondition)}
	output.Spec = &xast.QueryStmt_SQLSelect_QualifiedJoin_JoinCondition{
		SearchCondition: xbinaryexprTo(item.Spec),
		On: xposTo()}

	return output
}

func qualifiedjoinTo(item *xast.QueryStmt_SQLSelect_QualifiedJoin) *xlight.QueryStmt_SQLSelect_QualifiedJoin {
	// thisLeft is never nil
	thisLeft := item.LeftElement
	var ref *xlight.QueryStmt_SQLSelect_QualifiedJoin
	if thisLeft.LeftElement != nil {
		ref = qualifiedjoinTo(thisLeft)
	} else {
		ref = tableTo(thisLeft)
	}

	output := tableTo(item)
	output.LeftElement = ref
	output.TypeCondition = jointypeTo(item.TypeCondition)
	output.Spec = binaryexprTo(item.Spec.SearchCondition)

	return output
}

func xtablereferenceTo(item *xlight.QueryStmt_SQLSelect_QualifiedJoin) *xast.QueryStmt_SQLSelect_QualifiedJoin {
	if item == nil { return nil }

	if item.LeftElement != nil {
		return xqualifiedjoinTo(item)
	}
	return xtableTo(item)
}

func tablereferenceTo(item *xast.QueryStmt_SQLSelect_QualifiedJoin) *xlight.QueryStmt_SQLSelect_QualifiedJoin {
	if item == nil { return nil }

	if item.LeftElement != nil {
		return qualifiedjoinTo(item)
	}
	return tableTo(item)
}
