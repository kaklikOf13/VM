package main

import (
	vm "main/VM"
)

func main() {
	vm := vm.NewVM()
	vm.Init()
	vm.Compile("")
}
