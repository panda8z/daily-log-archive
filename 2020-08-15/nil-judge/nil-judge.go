/*
Interface in the go source code has two parts.
one is type , anothe is data .
when pass a nil as a value to the interface ,
the interface`s type and data is all nil value.

On this time  nil == the-interface is true.

On the other hand ,if a type pass to the nil interface .
the interface`s type is not nil, and the value is the nil meanwhile.
then  nil == the-interface is false,
*/
package main

import "fmt"

type MyImplement struct{}

// String implements fmt.Stringer
func (m *MyImplement) String() string {
	return "hi"
}

// GetStringer after this func return value`s type is *MyImplement and data is nil
func GetStringer() fmt.Stringer {
	var s *MyImplement = nil
	// if s == nil {
	// 	return nil
	// }
	return s
}

func main() {
	a := GetStringer()
	fmt.Printf("type of a is %T\n", a)
	fmt.Printf("value of a is %v\n", a)

	if a == nil {
		fmt.Println("GetStringer() == nil")
	} else {
		fmt.Println("GetStringer() != nil")
	}
}
