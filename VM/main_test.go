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
	fmt.Println(txt)
	if txt != "10+10;" {
		t.Fail()
	}
}

func TestVar(t *testing.T) {
	txt := vm.Compile(vm.Parse(vm.GerateTokens("test")))
	fmt.Println(txt)
	if txt != "test;" {
		t.Fail()
	}
}
func TestDefVar(t *testing.T) {
	txt := vm.Compile(vm.Parse(vm.GerateTokens("var test:uint16=0")))
	fmt.Println(txt)
	if txt != "unsigned short int test = 0;" {
		t.Fail()
	}
}

func TestDefVar2(t *testing.T) {
	txt := vm.Compile(vm.Parse(vm.GerateTokens("var test:uint16")))
	fmt.Println(txt)
	if txt != "unsigned short int test;" {
		t.Fail()
	}
}
func TestFunction(t *testing.T) {
	txt := vm.Compile(vm.Parse(vm.GerateTokens("fn main(a,b:int):int{a+b}")))
	fmt.Println(txt)
	if txt != "long int main(long int a,long int b){\n    a+b;\n};" {
		t.Fail()
	}
}
func TestFunction2(t *testing.T) {
	txt := vm.Compile(vm.Parse(vm.GerateTokens("fn main(a,b:int){a+b}")))
	fmt.Println(txt)
	if txt != "void main(long int a,long int b){\n    a+b;\n};" {
		t.Fail()
	}
}
