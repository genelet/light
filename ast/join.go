package ast

import (
	"fmt"
	"github.com/genelet/sqlproto/xast"
	"github.com/akito0107/xsqlparser/sqlast"
)

func xjoinconditionTo(c *sqlast.JoinCondition) (*xast.JoinCondition, error) {
	switch t := c.SearchCondition.(type) {
	case *sqlast.BinaryExpr:
		sc, err := xbinaryExprTo(t)
		if err != nil { return nil, err }
		return &xast.JoinCondition{
			SearchCondition: sc,
			On: xposTo(c.On)}, nil
	default:
	}
	return nil, fmt.Errorf("search condition %#v", c.SearchCondition)
}

func joinconditionTo(c *xast.JoinCondition) *sqlast.JoinCondition {
	if c == nil || c.SearchCondition == nil { return nil }
	return &sqlast.JoinCondition{
		SearchCondition: binaryExprTo(c.SearchCondition),
		On: posTo(c.On)}
}

func xtableTo(t *sqlast.Table) *xast.QualifiedJoin {
	return &xast.QualifiedJoin {
		Name: xobjectnameTo(t.Name),
		AliasName: xidentTo(t.Alias)}
}

func tableTo(t *xast.QualifiedJoin) *sqlast.Table {
	table := &sqlast.Table{
		Name: compoundToObjectname(t.Name)}
	if t.AliasName != nil {
		table.Alias = identTo(t.AliasName).(*sqlast.Ident)
	}
	return table
}

func xqualifiedjoinTo(item *sqlast.QualifiedJoin) (*xast.QualifiedJoin, error) {
	spec, err := xjoinconditionTo(item.Spec.(*sqlast.JoinCondition))
	if err != nil { return nil, err }

	table := item.RightElement.Ref.(*sqlast.Table)
	output := &xast.QualifiedJoin{
		Name: xobjectnameTo(table.Name),
		AliasName: xidentTo(table.Alias),
		TypeCondition: xjointypeTo(item.Type),
		Spec: spec}

	switch t := item.LeftElement.Ref.(type) {
	case *sqlast.Table:
		output.LeftElement = xtableTo(t)
	case *sqlast.QualifiedJoin:
		output.LeftElement, err = xqualifiedjoinTo(t)
	default:
		return nil, fmt.Errorf("left type %#v", t)
	}
	return output, err
}

func qualifiedjoinTo(item *xast.QualifiedJoin) *sqlast.QualifiedJoin {
	// thisLeft is never nil
	thisLeft := item.LeftElement
	var ref sqlast.TableReference
	if thisLeft.LeftElement != nil {
		ref = qualifiedjoinTo(thisLeft)
	} else {
		ref = tableTo(thisLeft)
	}

	return &sqlast.QualifiedJoin{
		LeftElement: &sqlast.TableJoinElement{Ref: ref},
		Type: jointypeTo(item.TypeCondition),
		RightElement: &sqlast.TableJoinElement{Ref: tableTo(item)},
		Spec: joinconditionTo(item.Spec)}
}

func xtablereferenceTo(item sqlast.TableReference) (*xast.QualifiedJoin, error) {
	switch t := item.(type) {
	case *sqlast.Table:
		return xtableTo(t), nil
	case *sqlast.QualifiedJoin:
		return xqualifiedjoinTo(t)
	default:
	}
	return nil, fmt.Errorf("join type %#v", item)
}

func tablereferenceTo(item *xast.QualifiedJoin) sqlast.TableReference {
	if item == nil { return nil }
	if item.LeftElement != nil {
		return qualifiedjoinTo(item)
	}
	return tableTo(item)
}
