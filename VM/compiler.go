package vm

type Compiler struct {
}

func (this *Compiler) CompileNode(node *Node) string {
	switch node.NodeType {
	case NT_VALUE:
		return node.Value[0].Compile()
	case NT_PLUS:
		return this.CompileNode(node.Value[0].(*Node)) + "+" + this.CompileNode(node.Value[1].(*Node))
	case NT_MINUS:
		return this.CompileNode(node.Value[0].(*Node)) + "-" + this.CompileNode(node.Value[1].(*Node))
	case NT_MUL:
		return this.CompileNode(node.Value[0].(*Node)) + "*" + this.CompileNode(node.Value[1].(*Node))
	case NT_DIV:
		return this.CompileNode(node.Value[0].(*Node)) + "/" + this.CompileNode(node.Value[1].(*Node))
	default:
		return ""
	}
}
func Compile(nodes []*Node) string {
	ret := ""
	c := Compiler{}
	for i := range nodes {
		if i > 0 {
			ret += "\n"
		}
		ret += c.CompileNode(nodes[i])
	}
	return ret
}
func NewCompiler() *Compiler {
	return &Compiler{}
}
