package light

import (
	"github.com/genelet/sqlproto/xast"
	"github.com/genelet/sqlproto/xlight"
)

func xjoinSpecTo(c *xlight.JoinSpec) *xast.JoinSpec {
    if c == nil { return nil }

	output := &xast.JoinSpec{}
	if x := c.GetJoinItem(); x != nil {
		output.JoinSpecClause = &xast.JoinSpec_JoinItem{JoinItem: xjoinConditionTo(x)}
	} else if x := c.GetNameItem(); x != nil {
		output.JoinSpecClause = &xast.JoinSpec_NameItem{NameItem: xnamedColumnsJoinTo(x)}
	}
	return output
}

func joinSpecTo(c *xast.JoinSpec) *xlight.JoinSpec {
    if c == nil { return nil }

	output := &xlight.JoinSpec{}
    if x := c.GetJoinItem(); x != nil {
        output.JoinSpecClause = &xlight.JoinSpec_JoinItem{JoinItem: joinConditionTo(x)}
    } else if x := c.GetNameItem(); x != nil {
        output.JoinSpecClause = &xlight.JoinSpec_NameItem{NameItem:  namedColumnsJoinTo(x)}
    }
    return output
}

func xnamedColumnsJoinTo(item *xlight.NamedColumnsJoin) *xast.NamedColumnsJoin {
    output := &xast.NamedColumnsJoin{
        Using: xposTo(),
        RParen: xposTo(item)}
    for _, ident := range item.ColumnList {
        output.ColumnList = append(output.ColumnList, xidentTo(ident))
    }
    return output
}

func namedColumnsJoinTo(item *xast.NamedColumnsJoin) *xlight.NamedColumnsJoin {
    if item == nil { return nil }

    output := &xlight.NamedColumnsJoin{}
    for _, ident := range item.ColumnList {
        output.ColumnList = append(output.ColumnList, ident.Value)
    }
    return output
}

func xjoinConditionTo(c *xlight.JoinCondition) *xast.JoinCondition {
    x := xbinaryExprTo(c.SearchCondition)
    return &xast.JoinCondition{
        SearchCondition: x,
		On: xposplusTo(c)}
}

func joinConditionTo(c *xast.JoinCondition) *xlight.JoinCondition {
    if c == nil { return nil }
    return &xlight.JoinCondition{
        SearchCondition: binaryExprTo(c.SearchCondition)}
}

func xtableJoinElementTo(item *xlight.TableJoinElement) *xast.TableJoinElement {
    x := xtableReferenceTo(item.Ref)
    return &xast.TableJoinElement{Ref: x}
}

func tableJoinElementTo(item *xast.TableJoinElement) *xlight.TableJoinElement {
    return &xlight.TableJoinElement{Ref: tableReferenceTo(item.Ref)}
}

func xnaturalJoinTo(item *xlight.NaturalJoin) *xast.NaturalJoin {
    if item == nil { return nil }

    left := xtableJoinElementTo(item.LeftElement)
    right := xtableJoinElementTo(item.RightElement)

	joinType := &xast.JoinType{
			From: xposTo(),
			To: xposplusTo(item.Type),
			Condition: xast.JoinTypeCondition(item.Type)}

    return &xast.NaturalJoin{
        LeftElement: left,
        Type: joinType,
        RightElement: right}
}

func naturalJoinTo(item *xast.NaturalJoin) *xlight.NaturalJoin {
    if item == nil { return nil }

    return &xlight.NaturalJoin{
        LeftElement: tableJoinElementTo(item.LeftElement),
        Type: xlight.JoinTypeCondition(item.Type.Condition),
        RightElement: tableJoinElementTo(item.RightElement)}
}

func xqualifiedJoinTo(item *xlight.QualifiedJoin) *xast.QualifiedJoin {
    if item == nil { return nil }

    left := xtableJoinElementTo(item.LeftElement)
    right:= xtableJoinElementTo(item.RightElement)
    x := xjoinSpecTo(item.Spec)

	joinType := &xast.JoinType{
			From: xposTo(),
			To: xposplusTo(item.Type),
			Condition: xast.JoinTypeCondition(item.Type)}

    return &xast.QualifiedJoin{
        LeftElement: left,
        Type: joinType,
        RightElement: right,
        Spec: x}
}

func qualifiedJoinTo(item *xast.QualifiedJoin) *xlight.QualifiedJoin {
    if item == nil { return nil }

    return &xlight.QualifiedJoin{
        LeftElement: tableJoinElementTo(item.LeftElement),
        Type: xlight.JoinTypeCondition(item.Type.Condition),
        RightElement: tableJoinElementTo(item.RightElement),
        Spec: joinSpecTo(item.Spec)}
}

func xtableReferenceTo(item *xlight.TableReference) *xast.TableReference {
    if item == nil { return nil }

	output := &xast.TableReference{}
	if x := item.GetTableItem(); x != nil {
		output.TableReferenceClause = &xast.TableReference_TableItem{TableItem: xtableTo(x)}
	} else if x := item.GetQualifiedItem(); x != nil {
		output.TableReferenceClause = &xast.TableReference_QualifiedItem{QualifiedItem: xqualifiedJoinTo(x)}
	} else if x := item.GetNaturalItem(); x != nil {
		output.TableReferenceClause = &xast.TableReference_NaturalItem{NaturalItem: xnaturalJoinTo(x)}
	}
	return output
}

func tableReferenceTo(item *xast.TableReference) *xlight.TableReference {
    if item == nil { return nil }

	output := &xlight.TableReference{}
    if x := item.GetQualifiedItem(); x != nil {
        output.TableReferenceClause = &xlight.TableReference_QualifiedItem{QualifiedItem: qualifiedJoinTo(x)}
    } else if x := item.GetTableItem(); x != nil {
        output.TableReferenceClause = &xlight.TableReference_TableItem{TableItem: tableTo(x)}
    } else if x := item.GetNaturalItem(); x != nil {
        output.TableReferenceClause = &xlight.TableReference_NaturalItem{NaturalItem: naturalJoinTo(x)}
    }
    return output
}

func xtableTo(item *xlight.Table) *xast.Table {
    output := &xast.Table{
        Name: xobjectnameTo(item.Name),
        Alias: xidentTo(item.Alias),
        ArgsRParen: xposTo(item.Args),
        WithHintsRParen: xposTo(item.WithHints)}
    for _, arg := range item.Args {
        x := xargsNodeTo(arg)
        output.Args = append(output.Args, x)
    }
    for _, arg := range item.WithHints {
        x := xargsNodeTo(arg)
        output.WithHints = append(output.WithHints, x)
    }
    return output
}

func tableTo(item *xast.Table) *xlight.Table {
    output := &xlight.Table{
        Name: objectnameTo(item.Name)}
    if item.Alias != nil {
        output.Alias = identTo(item.Alias)
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
