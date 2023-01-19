package ast

import (
	"fmt"
	"github.com/genelet/sqlproto/xast"
	"github.com/akito0107/xsqlparser/sqlast"
)

func xwhereStmtTo(stmt sqlast.Node) (*xast.WhereStmt, error ) {
	if stmt == nil { return nil, nil }

	output := &xast.WhereStmt{}
    switch t := stmt.(type) {
    case *sqlast.InSubQuery:
        where, err := xinsubqueryTo(t)
        if err != nil { return nil, err }
        output.WhereStmtClause = &xast.WhereStmt_InQuery{InQuery: where}
    case *sqlast.BinaryExpr:
        where, err := xbinaryExprTo(t)
        if err != nil { return nil, err }
        output.WhereStmtClause = &xast.WhereStmt_BinExpr{BinExpr: where}
    default:
        return nil, fmt.Errorf("missing where type %T", t)
    }

	return output, nil
}

func whereStmtTo(stmt *xast.WhereStmt) sqlast.Node {
	if stmt == nil { return nil }

	if x := stmt.GetInQuery(); x != nil {
        return insubqueryTo(x)
    } else if x := stmt.GetBinExpr(); x != nil {
        return binaryExprTo(x)
    }
	return nil
}

func xvalueStmtTo(stmt sqlast.Node) (*xast.ValueStmt, error ) {
	if stmt == nil { return nil, nil }

	output := &xast.ValueStmt{}
    switch t := stmt.(type) {
    case *sqlast.SingleQuotedString:
        output.ValueStmtClause = &xast.ValueStmt_StringStmtValue{StringStmtValue: xstringTo(t)}
    case *sqlast.LongValue:
        output.ValueStmtClause = &xast.ValueStmt_LongStmtValue{LongStmtValue: xlongTo(t)}
    case *sqlast.Ident:
        output.ValueStmtClause = &xast.ValueStmt_IdentStmtValue{IdentStmtValue: xidentTo(t)}
    default:
        return nil, fmt.Errorf("missing value stmt type %T", t)
    }

	return output, nil
}

func valueStmtTo(stmt *xast.ValueStmt) sqlast.Node {
	if stmt == nil { return nil }

	if x := stmt.GetLongStmtValue(); x != nil {
		return longTo(x)
	} else if x := stmt.GetIdentStmtValue(); x != nil {
		return identTo(x).(*sqlast.Ident)
	} else if x := stmt.GetStringStmtValue(); x != nil {
		return stringTo(x)
	}

	return nil
}

func xsourceStmtTo(stmt sqlast.InsertSource) (*xast.SourceStmt, error ) {
	if stmt == nil { return nil, nil }

	output := &xast.SourceStmt{}
    switch t := stmt.(type) {
    case *sqlast.SubQuerySource:
		// definition sqlast.SubQuerySource{SubQuery: q}
		source, err := XQueryTo(t.SubQuery)
		if err != nil { return nil, err }
		output.SourceStmtClause = &xast.SourceStmt_SubItem{SubItem: &xast.SubQuerySource{SubQuery: source}}
    case *sqlast.ConstructorSource:
        source, err := xconstructorSourceTo(t)
        if err != nil { return nil, err }
        output.SourceStmtClause = &xast.SourceStmt_StructorItem{StructorItem: source}
    default:
        return nil, fmt.Errorf("missing source type %T", t)
    }

	return output, nil
}

func sourceStmtTo(stmt *xast.SourceStmt) sqlast.InsertSource {
	if stmt == nil { return nil }

	if x := stmt.GetSubItem(); x != nil {
        return &sqlast.SubQuerySource{SubQuery: QueryTo(x.SubQuery)}
    } else if x := stmt.GetStructorItem(); x != nil {
        return constructorSourceTo(x)
    }
	return nil
}

func xalterTableActionTo(stmt sqlast.AlterTableAction) (*xast.AlterTableAction, error) {
	if stmt == nil { return nil, nil }

	output := &xast.AlterTableAction{}
    switch t := stmt.(type) {
    case *sqlast.AddColumnTableAction:
        x, err := xaddColumnTableActionTo(t)
        if err != nil { return nil, err }
        output.AlterTableActionClause = &xast.AlterTableAction_AddColumnItem{AddColumnItem: x}
    case *sqlast.AddConstraintTableAction:
        x, err := xaddConstraintTableActionTo(t)
        if err != nil { return nil, err }
        output.AlterTableActionClause = &xast.AlterTableAction_AddConstraintItem{AddConstraintItem: x}
    case *sqlast.DropConstraintTableAction:
        x, err := xdropConstraintTableActionTo(t)
        if err != nil { return nil, err }
        output.AlterTableActionClause = &xast.AlterTableAction_DropConstraintItem{DropConstraintItem: x}
    case *sqlast.RemoveColumnTableAction:
        x, err := xremoveColumnTableActionTo(t)
        if err != nil { return nil, err }
        output.AlterTableActionClause = &xast.AlterTableAction_RemoveColumnItem{RemoveColumnItem: x}
    default:
        return nil, fmt.Errorf("missing actio node type %T", t)
    }

	return output, nil
}

func alterTableActionTo(stmt *xast.AlterTableAction) sqlast.AlterTableAction {
	if stmt == nil { return nil }

	if x := stmt.GetAddColumnItem(); x != nil {
        return addColumnTableActionTo(x)
    } else if x := stmt.GetAddConstraintItem(); x != nil {
        return addConstraintTableActionTo(x)
    } else if x := stmt.GetDropConstraintItem(); x != nil {
        return dropConstraintTableActionTo(x)
    } else if x := stmt.GetRemoveColumnItem(); x != nil {
        return removeColumnTableActionTo(x)
    }
	return nil
}
