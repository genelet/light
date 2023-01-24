package ast

import (
	"fmt"
	"github.com/genelet/sqlproto/xast"
	"github.com/akito0107/xsqlparser/sqlast"
)

func xjoinSpecTo(c sqlast.JoinSpec) (*xast.JoinSpec, error) {
	output := &xast.JoinSpec{}

	switch t := c.(type) {
	case *sqlast.JoinCondition:
		x, err := xjoinConditionTo(t)
		if err != nil { return nil, err }
		output.JoinSpecClause = &xast.JoinSpec_JoinItem{JoinItem: x}
	case *sqlast.NamedColumnsJoin:
		x, err := xnamedColumnsJoinTo(t)
		if err != nil { return nil, err }
		output.JoinSpecClause = &xast.JoinSpec_NameItem{NameItem: x}
	default:
		return nil, fmt.Errorf("missing join spec type %T", t)
	}
	return output, nil
}

func joinSpecTo(c *xast.JoinSpec) sqlast.JoinSpec {
	if c == nil { return nil }

	if x := c.GetJoinItem(); x != nil {
		return joinConditionTo(x)
	} else if x := c.GetNameItem(); x != nil {
		return namedColumnsJoinTo(x)
	}
	return nil
}

func xnamedColumnsJoinTo(item *sqlast.NamedColumnsJoin) (*xast.NamedColumnsJoin, error) {
	output := &xast.NamedColumnsJoin{
		Using: xposTo(item.Using),
		RParen: xposTo(item.RParen)}
	for _, ident := range item.ColumnList {
		output.ColumnList = append(output.ColumnList, xidentTo(ident))
	}
	return output, nil
}

func namedColumnsJoinTo(item *xast.NamedColumnsJoin) *sqlast.NamedColumnsJoin {
	if item == nil { return nil }

	output := &sqlast.NamedColumnsJoin{
		Using: posTo(item.Using),
		RParen: posTo(item.RParen)}
	for _, ident := range item.ColumnList {
		output.ColumnList = append(output.ColumnList, identTo(ident).(*sqlast.Ident))
	}
	return output
}

func xjoinConditionTo(c *sqlast.JoinCondition) (*xast.JoinCondition, error) {
	x, err := xbinaryExprTo(c.SearchCondition.(*sqlast.BinaryExpr))
	return &xast.JoinCondition{
		SearchCondition: x,
		On: xposTo(c.On)}, err
}

func joinConditionTo(c *xast.JoinCondition) *sqlast.JoinCondition {
	if c == nil { return nil }
	return &sqlast.JoinCondition{
		SearchCondition: binaryExprTo(c.SearchCondition),
		On: posTo(c.On)}
}

/*
func xtableTo(t *sqlast.Table) *xast.QualifiedJoin {
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
*/

func xtableJoinElementTo(item *sqlast.TableJoinElement) (*xast.TableJoinElement, error) {
	x, err := xtableReferenceTo(item.Ref)
	return &xast.TableJoinElement{Ref: x}, err
}

func tableJoinElementTo(item *xast.TableJoinElement) *sqlast.TableJoinElement {
	return &sqlast.TableJoinElement{Ref: tableReferenceTo(item.Ref)}
}

func xnaturalJoinTo(item *sqlast.NaturalJoin) (*xast.NaturalJoin, error) {
	if item == nil { return nil, nil }

	left, err := xtableJoinElementTo(item.LeftElement)
	if err != nil { return nil, err }
	right, err := xtableJoinElementTo(item.RightElement)
	if err != nil { return nil, err }

	return &xast.NaturalJoin{
		LeftElement: left,
		Type: xjoinTypeTo(item.Type),
		RightElement: right}, nil
}

func naturalJoinTo(item *xast.NaturalJoin) *sqlast.NaturalJoin {
	if item == nil { return nil }
	// thisLeft is never nil

	return &sqlast.NaturalJoin{
		LeftElement: tableJoinElementTo(item.LeftElement),
		Type: joinTypeTo(item.Type),
		RightElement: tableJoinElementTo(item.RightElement)}
}

func xqualifiedJoinTo(item *sqlast.QualifiedJoin) (*xast.QualifiedJoin, error) {
	if item == nil { return nil, nil }

	left, err := xtableJoinElementTo(item.LeftElement)
	if err != nil { return nil, err }
	right, err := xtableJoinElementTo(item.RightElement)
	if err != nil { return nil, err }
	x, err := xjoinSpecTo(item.Spec)
	if err != nil { return nil, err }

	return &xast.QualifiedJoin{
		LeftElement: left,
		Type: xjoinTypeTo(item.Type),
		RightElement: right,
		Spec: x}, nil
}

func qualifiedJoinTo(item *xast.QualifiedJoin) *sqlast.QualifiedJoin {
	if item == nil { return nil }
	// thisLeft is never nil

	return &sqlast.QualifiedJoin{
		LeftElement: tableJoinElementTo(item.LeftElement),
		Type: joinTypeTo(item.Type),
		RightElement: tableJoinElementTo(item.RightElement),
		Spec: joinSpecTo(item.Spec)}
}

func xtableReferenceTo(item sqlast.TableReference) (*xast.TableReference, error) {
	if item == nil { return nil, nil }

	output := &xast.TableReference{}
	switch t := item.(type) {
	case *sqlast.Table:
		x, err := xtableTo(t)
		if err != nil { return nil, err }
		output.TableReferenceClause = &xast.TableReference_TableItem{TableItem: x}
	case *sqlast.NaturalJoin:
		x, err := xnaturalJoinTo(t)
		if err != nil { return nil, err }
		output.TableReferenceClause = &xast.TableReference_NaturalItem{NaturalItem: x}
	case *sqlast.QualifiedJoin:
		x, err := xqualifiedJoinTo(t)
		if err != nil { return nil, err }
		output.TableReferenceClause = &xast.TableReference_QualifiedItem{QualifiedItem: x}
	default:
		return nil, fmt.Errorf("missing table reference type %T", t)
	}
	return output, nil
}

func tableReferenceTo(item *xast.TableReference) sqlast.TableReference {
	if item == nil { return nil }

	if x := item.GetQualifiedItem(); x != nil {
		return qualifiedJoinTo(x)
	} else if x := item.GetTableItem(); x != nil {
		return tableTo(x)
	} else if x := item.GetNaturalItem(); x != nil {
		return naturalJoinTo(x)
	}
	return nil
}

func xtableTo(item *sqlast.Table) (*xast.Table, error) {
	output := &xast.Table{
		Name: xobjectnameTo(item.Name),
		Alias: xidentTo(item.Alias),
		ArgsRParen: xposTo(item.ArgsRParen),
		WithHintsRParen: xposTo(item.WithHintsRParen)}
	for _, arg := range item.Args {
		x, err := xargsNodeTo(arg)
		if err != nil { return nil, err }
		output.Args = append(output.Args, x)
	}
	for _, arg := range item.WithHints {
		x, err := xargsNodeTo(arg)
		if err != nil { return nil, err }
		output.WithHints = append(output.WithHints, x)
	}
	return output, nil
}

func tableTo(item *xast.Table) *sqlast.Table {
	output := &sqlast.Table{
		Name: objectnameTo(item.Name),
		ArgsRParen: posTo(item.ArgsRParen),
		WithHintsRParen: posTo(item.WithHintsRParen)}
	if item.Alias != nil {
		output.Alias = identTo(item.Alias).(*sqlast.Ident)
	}
	for _, arg := range item.Args {
		x := argsNodeTo(arg)
		output.Args = append(output.Args, x)
	}
	for _, arg := range item.WithHints {
		x := argsNodeTo(arg)
		output.WithHints = append(output.WithHints, x)
	}
	return output
}
