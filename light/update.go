package light

import (
	"github.com/genelet/sqlproto/xast"
	"github.com/genelet/sqlproto/xlight"
)

func XUpdateTo(stmt *xlight.UpdateStmt) *xast.UpdateStmt {
	if stmt == nil { return nil }

	output := &xast.UpdateStmt{
		Update: xposTo(),
		TableName: xobjectnameTo(stmt.TableName),
		Selection: xwhereNodeTo(stmt.Selection)}
	for _, assignment := range stmt.Assignments {
		output.Assignments = append(output.Assignments, xassignmentTo(assignment))
	}
	return output
}

func UpdateTo(stmt *xast.UpdateStmt) *xlight.UpdateStmt {
	if stmt == nil { return nil }

	output := &xlight.UpdateStmt{
		TableName: objectnameTo(stmt.TableName),
		Selection: whereNodeTo(stmt.Selection)}
	for _, assignment := range stmt.Assignments {
		output.Assignments = append(output.Assignments, assignmentTo(assignment))
	}
	return output
}
