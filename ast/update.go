package ast

import (
	"github.com/genelet/sqlproto/xast"
	"github.com/akito0107/xsqlparser/sqlast"
)

func XUpdateTo(stmt *sqlast.UpdateStmt) (*xast.UpdateStmt, error) {
	selection, err := xwhereStmtTo(stmt.Selection)
	if err != nil { return nil, err }

	output := &xast.UpdateStmt{
		Update: xposTo(stmt.Update),
		TableName: xobjectnameTo(stmt.TableName),
		Selection: selection}
	for _, assignment := range stmt.Assignments {
		x, err := xassignmentTo(assignment)
		if err != nil { return nil, err }	
		output.Assignments = append(output.Assignments, x)
	}
	return output, nil
}

func UpdateTo(stmt *xast.UpdateStmt) *sqlast.UpdateStmt {
	output := &sqlast.UpdateStmt{
		Update: posTo(stmt.Update),
		TableName: objectnameTo(stmt.TableName),
		Selection: whereStmtTo(stmt.Selection)}
	for _, assignment := range stmt.Assignments {
		output.Assignments = append(output.Assignments, assignmentTo(assignment))
	}
	return output
}
