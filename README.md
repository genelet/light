# sqlproto

sqlproto parses standard SQL query into protocol buffer, and vice versa.

<br /><br />

## Introdution

This project defines a protobuf message for query statement in the [ANSI/ISO SQL standard](https://en.wikipedia.org/wiki/ISO/IEC_9075).
It uses [xsqlparser](https://github.com/akito0107/xsqlparser), which is ported of [sqlparser-rs](https://github.com/andygrove/sqlparser-rs) in Go, to translate SQL query into protocol buffer, and/or protocol buffer into SQL query.

Why the project? Protobuf provides a simple, clean and better graph to represent parsed SQL structure than native GO or RUST objects. There are three usages:

- build SQL parser and so new SQL engine
- contruct complex SQL query on-the-fly, to search noSQL database, time-series database and other loosely-structured data systems.
- provide a machine learning framework on database schema, or meta. For example, the [text-to-SQL](https://yale-lily.github.io/spider) semantic parsing requires a meta standard, so neutral language questions can be machine-learned, expressed as SQLs, and searched from knowledge bases.

<br /><br />

## Protocol Buffer

The definition of [the SQL protocol buffer](https://github.com/genelet/sqlproto/blob/main/proto/sqlight.proto) is

```protobuf
syntax = "proto3";
package sqlight;

option go_package = "./xlight";

enum OperatorType {
	Plus = 0;
	Minus = 1;
	Multiply = 2;
	Divide = 3;
	Modulus = 4;
	Gt = 5;
	Lt = 6;
	GtEq = 7;
	LtEq = 8;
	Eq = 9;
	NotEq = 10;
	And = 11;
	Or = 12;
	Not = 13;
	Like = 14;
	NotLike = 15;
	None = 16;
}

enum AggType {
	UnknownAgg = 0;
	MAX    = 1;
	MIN    = 2;
	COUNT  = 3;
	SUM    = 4;
	AVG    = 5;
}

enum SetOperatorType {
	Union = 0;
	Intersect = 1;
	Except = 2;
}

enum JoinTypeCondition {
	INNER = 0;
	LEFT = 1;
	RIGHT = 2;
	FULL = 3;
	LEFTOUTER = 4;
	RIGHTOUTER = 5;
	FULLOUTER = 6;
	IMPLICIT = 7;
}

message CompoundIdent {
	repeated string idents = 1;
}

message AggFunction {
	AggType typeName = 1;
	repeated CompoundIdent restArgs = 2;
}

message QueryStmt {
	message CTE {
		string aliasName = 1;
		QueryStmt query = 2;
    }
	repeated CTE CTEs = 2;

	message InSubQuery {
		CompoundIdent expr = 1;
		QueryStmt subQuery = 2;
		bool negated = 3;
	}

	message BinaryExpr {
		oneof LeftOneOf {
			CompoundIdent leftIdents = 1;
			BinaryExpr leftBinary = 2;
		}
		OperatorType op = 3;
		oneof RightOneOf {
			CompoundIdent rightIdents = 4;
			BinaryExpr rightBinary = 5;
			InSubQuery queryValue = 6;
			string singleQuotedString = 7;
			double doubleValue = 8;
			int64 longValue = 9;
		}
	}

	message SQLSelect {
		bool distinctBool = 1;

		message SQLSelectItem {
			AggFunction fieldFunction = 1;
			CompoundIdent fieldIdents = 2;
			string aliasName = 3;
		}
		repeated SQLSelectItem projection = 2;

		message QualifiedJoin {
			CompoundIdent name = 1;
			string aliasName = 2;
			QualifiedJoin leftElement = 3;
			JoinTypeCondition typeCondition = 4;
			BinaryExpr spec = 5;
		}
		repeated QualifiedJoin fromClause = 3;

		oneof WhereClause {
			InSubQuery inQuery = 4;
			BinaryExpr binExpr = 5;
		}

		repeated CompoundIdent groupByClause = 8;
		BinaryExpr havingClause = 9;
	}
	message SetOperationExpr {
		SQLSelect leftSide = 1;
		bool allBool = 2;
		SetOperatorType op = 3;
		SetOperationExpr rightSide = 4;
	}
	SetOperationExpr body = 4;

	message OrderByExpr {
		CompoundIdent expr = 1;
		bool aSCBool = 3;
	}
	repeated OrderByExpr orderBy = 5;

	message LimitExpr {
		bool allBool = 1;
		int64 limitValue = 4;
		int64 offsetValue = 5;
	}
	LimitExpr limitExpression = 6;
}
```

which covers most, if not all, SQL query cases. For example

- simple query: *SELECT a from test_table*
- join and aggregate: _SELECT orders.product as prod, SUM(orders.quantity) AS product_units, accounts.* FROM orders LEFT JOIN accounts ON orders.account_id = accounts.id INNER JOIN accounts_type ON accounts_type.type_id = accounts.type_id WHERE orders.region IN (SELECT region FROM top_regions) ORDER BY product_units ASC LIMIT 100_
- union set: *SELECT x FROM a UNION SELECT x FROM b EXCEPT select x FROM c*
- sub queries: *WITH regional_sales AS (SELECT region, SUM(amount) AS total_sales FROM orders GROUP BY region) SELECT product, SUM(quantity) AS product_units FROM orders WHERE region IN (SELECT region FROM top_regions) GROUP BY region, product*

<br /><br />

## Usage

