package light

import (
	"fmt"
	"github.com/genelet/sqlproto/xast"
	"github.com/akito0107/xsqlparser/sqlast"
)

func xjoinconditionTo(c *xlight.JoinCondition) *xast.JoinCondition {
	if c == nil { return nil }	
	
	output := &xast.JoinCondition{
			On: xposTo()}
	if x := c.GetSearchCondition(); x != nil {
		output.SearchCondition = xbinaryExprTo(x)
	} else {
		return nil
	}
	return output
}

func joinconditionTo(c *xast.JoinCondition) *xlight.JoinCondition {
	if c == nil { return nil }	
	
	output := &xlight.JoinCondition{}
	if x := c.GetSearchCondition(); x != nil {
		output.SearchCondition = binaryExprTo(x)
	} else {
		return nil
	}
	return output
}

func xtableTo(t *xlight.Table) *xast.QualifiedJoin {
	return &xast.QualifiedJoin {
		Name: xobjectnameTo(t.Name),
		AliasName: xidentTo(t.Alias)}
}

func tableTo(t *xast.QualifiedJoin) *sqlast.Table {
	table := &sqlast.Table{
		Name: objectnameTo(t.Name)}
	if t.AliasName != nil {
		table.Alias = identTo(t.AliasName).(*sqlast.Ident)
	}
	return table
}

func xqualifiedjoinTo(item *xlight.QualifiedJoin) (*xast.QualifiedJoin, error) {
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

func qualifiedjoinTo(item *xast.QualifiedJoin) *xlight.QualifiedJoin {
	// thisLeft is never nil
	thisLeft := item.LeftElement
	var ref xlight.TableReference
	if thisLeft.LeftElement != nil {
		ref = qualifiedjoinTo(thisLeft)
	} else {
		ref = tableTo(thisLeft)
	}

	return &sqlast.QualifiedJoin{
		LeftElement: &xlight.TableJoinElement{Ref: ref},
		Type: jointypeTo(item.TypeCondition),
		RightElement: &xlight.TableJoinElement{Ref: tableTo(item)},
		Spec: joinconditionTo(item.Spec)}
}

func xtablereferenceTo(item xlight.TableReference) (*xast.QualifiedJoin, error) {
	switch t := item.(type) {
	case *xlight.Table:
		return xtableTo(t), nil
	case *xlight.QualifiedJoin:
		return xqualifiedjoinTo(t)
	default:
	}
	return nil, fmt.Errorf("join type %#v", item)
}

func tablereferenceTo(item *xast.QualifiedJoin) xlight.TableReference {
	if item == nil { return nil }
	if item.LeftElement != nil {
		return qualifiedjoinTo(item)
	}
	return tableTo(item)
}
