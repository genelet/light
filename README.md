# sqlproto

sqlproto parses standard SQL query into protocol buffer, and vice versa.

[![GoDoc](https://godoc.org/github.com/genelet/sqlproto?status.svg)](https://godoc.org/github.com/genelet/sqlproto)


<br /><br />
## 1. Introdution

This project defines a protocol buffer message for SQL query statement in the [ANSI/ISO SQL standard](https://en.wikipedia.org/wiki/ISO/IEC_9075).
It uses [xsqlparser](https://github.com/akito0107/xsqlparser), which is ported of [sqlparser-rs](https://github.com/andygrove/sqlparser-rs) in Go, to translate SQL query into protobuf, and protobuf into SQL query.

Why is the project? Protocol buffer provides an easier, cleaner and better graph to represent parsed SQL structure than native GO or RUST objects. There are three usages:

- build SQL parser and new SQL engine
- contruct complex SQL query on-the-fly, to search noSQL database, time-series database and other loosely-structured data systems, as if they were RDBs.
- provide a machine learning framework on database meta. For example, the [text-to-SQL](https://yale-lily.github.io/spider) semantic parsing requires a meta standard. Solving text-to-SQL graph problems will let one to treat neutral language questions as SQL queries, and to answer them as SQL searches on knowledge base.

<br /><br />
## 2. Protocol Buffer

The definition of [the SQL meta protocol buffer](https://github.com/genelet/sqlproto/blob/main/proto/sqlight.proto) is

<details>
	<summary>Click to read the proto</summary>
	
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
</details>

which covers most, if not all, SQL query cases. For example

##### simple query: 
```sql
SELECT a from test_table
```

##### join and aggregate: 
```sql
SELECT orders.product as prod, SUM(orders.quantity) AS product_units, accounts.* FROM orders LEFT JOIN accounts ON orders.account_id = accounts.id INNER JOIN accounts_type ON accounts_type.type_id = accounts.type_id WHERE orders.region IN (SELECT region FROM top_regions) ORDER BY product_units ASC LIMIT 100
```

##### union set: 
```sql
SELECT x FROM a UNION SELECT x FROM b EXCEPT select x FROM c
```

##### sub queries: 
```sql
WITH regional_sales AS (SELECT region, SUM(amount) AS total_sales FROM orders GROUP BY region) SELECT product, SUM(quantity) AS product_units FROM orders WHERE region IN (SELECT region FROM top_regions) GROUP BY region, product
```

<br /><br />
## 3. Usage

Two main functions are implmented:
- _SQL2Proto_, to parse SQL query into protobuf
- _Proto2SQL_, to construct SQL query from protobuf

#### 3.1) Example, to parse SQL query

```go
package main

import (
    "fmt"
    "github.com/genelet/sqlproto/light"
)

func main() {
    pb, err := light.SQL2Proto(`SELECT * FROM test_table`)
    if err != nil { panic(err) }
    fmt.Printf("%s\n", pb.String())
}
```

The output:

```bash
body:{leftSide:{projection:{fieldIdents:{idents:"*"}} fromClause:{name:{idents:"test_table"}}}}.
```



#### 3.2) Example, to construct SQL query

```go
package main

import (
    "fmt"
    "github.com/genelet/sqlproto/light"
    "github.com/genelet/sqlproto/xlight"
)

func main() {
    project := &xlight.QueryStmt_SQLSelect_SQLSelectItem{
        FieldIdents: &xlight.CompoundIdent{Idents: []string{`*`}},
    }
    fromClause:= &xlight.QueryStmt_SQLSelect_QualifiedJoin{
        Name: &xlight.CompoundIdent{Idents: []string{`test_table`}},
    }
    left := &xlight.QueryStmt_SQLSelect{
        Projection: []*xlight.QueryStmt_SQLSelect_SQLSelectItem{project},
        FromClause: []*xlight.QueryStmt_SQLSelect_QualifiedJoin{fromClause},
    }
    pb := &xlight.QueryStmt{
        Body: &xlight.QueryStmt_SetOperationExpr{LeftSide: left},
    }
    str := light.Proto2SQL(pb)
    fmt.Printf("%s\n", str)
}
```

The output:

```bash
SELECT * FROM test_table
```

Please check [https://godoc.org/github.com/genelet/sqlproto](https://godoc.org/github.com/genelet/sqlproto) for detailed document.
