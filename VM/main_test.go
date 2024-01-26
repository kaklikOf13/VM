package vm_test

import (
	"fmt"
	vm "main/VM"
	"testing"
)

func TestLexer(t *testing.T) {
	toks := vm.GerateTokens("10+100")
	fmt.Println(toks)
	if len(toks) < 3 {
		t.Fail()
	}
}
func TestParser(t *testing.T) {
	n := vm.Parse(vm.GerateTokens("10+10"))
	fmt.Println(n)
}

func TestCompiler(t *testing.T) {
	txt := vm.Compile(vm.Parse(vm.GerateTokens("10 + 10")))
	if txt != "10+10" {
		t.Fail()
	}
}
