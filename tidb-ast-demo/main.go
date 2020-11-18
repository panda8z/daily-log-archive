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