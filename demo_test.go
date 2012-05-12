// run: go test demo_test.go
// test filename end of "_test.go"
// or just under that directory: go test
// intro: http://tiaozhanshu.com/computer/golang/go-unit-test/
package main

import (
 
	"testing"
)

func Hi() string{
 return "hi"
}


 
func TestHi(t *testing.T){
	if Hi()!="hi"	{
	 t.Error("err...")
	}
}

