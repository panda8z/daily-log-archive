package main

import (
	"fmt"
	"github.com/pingcap/tidb/parser"
	"github.com/pingcap/tidb/ast"
)

type visitor struct {}

func (v *visitor) Enter(in ast.Node) (out ast.Node, skipChildren bool)  {
	fmt.Printf("%T\n", in)
	return in, false
}

func (v *visitor) Leave(in ast.Node) (out ast.Node, ok bool) {
	return in, true
}

func main()  {
	sql := "SELECT /*+ TIDB_SMJ(employees) */ emp_no, first_name, last_name " +
		   "FROM employees USE INDEX (last_name) " +
		   "where last_name='Aamodt' and gender='F' and birth_date > '1960-01-01'" 
	sqlParser := paser.New()
	stmtNodes, err := sqlParser.Parse(sql, "", "")
	if err != nil {
		fmt.Printf("parse error:\n%v\n%s", err, sql)
		return
	}
	for _, stmtNode := range stmtNodes {
		v := visitor{}
		stmtNode.Accept(&v)
	}
}

/*
% go build
# github.com/pingcap/tidb/parser
/Users/apple/go/pkg/mod/github.com/pingcap/tidb@v2.0.11+incompatible/parser/lexer.go:132:26: undefined: yySymType
/Users/apple/go/pkg/mod/github.com/pingcap/tidb@v2.0.11+incompatible/parser/misc.go:595:24: undefined: yySymType
/Users/apple/go/pkg/mod/github.com/pingcap/tidb@v2.0.11+incompatible/parser/yy_parser.go:70:11: undefined: yySymType
/Users/apple/go/pkg/mod/github.com/pingcap/tidb@v2.0.11+incompatible/parser/yy_parser.go:71:9: undefined: yySymType
/Users/apple/go/pkg/mod/github.com/pingcap/tidb@v2.0.11+incompatible/parser/yy_parser.go:72:9: undefined: yySymType
/Users/apple/go/pkg/mod/github.com/pingcap/tidb@v2.0.11+incompatible/parser/yy_parser.go:151:38: undefined: yySymType
/Users/apple/go/pkg/mod/github.com/pingcap/tidb@v2.0.11+incompatible/parser/yy_parser.go:155:36: undefined: yySymType
/Users/apple/go/pkg/mod/github.com/pingcap/tidb@v2.0.11+incompatible/parser/yy_parser.go:163:14: undefined: yyLexer
/Users/apple/go/pkg/mod/github.com/pingcap/tidb@v2.0.11+incompatible/parser/yy_parser.go:163:29: undefined: yySymType
/Users/apple/go/pkg/mod/github.com/pingcap/tidb@v2.0.11+incompatible/parser/yy_parser.go:191:18: undefined: yyLexer
/Users/apple/go/pkg/mod/github.com/pingcap/tidb@v2.0.11+incompatible/parser/yy_parser.go:191:18: too many errors
*/