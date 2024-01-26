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
