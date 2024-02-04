package vm

type Compiler struct {
}

func GetTP(t string) string {
	if _, ok := mainTypes[t]; ok {
		return mainTypes[t]
	}
	return t
}
func (this *Compiler) CompileNode(node *Node, offset string) string {
	switch node.NodeType {
	case NT_VALUE:
		return node.Value[0].Compile()
	case NT_PLUS:
		return this.CompileNode(node.Value[0].(*Node), offset) + "+" + this.CompileNode(node.Value[1].(*Node), offset)
	case NT_MINUS:
		return this.CompileNode(node.Value[0].(*Node), offset) + "-" + this.CompileNode(node.Value[1].(*Node), offset)
	case NT_MUL:
		return this.CompileNode(node.Value[0].(*Node), offset) + "*" + this.CompileNode(node.Value[1].(*Node), offset)
	case NT_DIV:
		return this.CompileNode(node.Value[0].(*Node), offset) + "/" + this.CompileNode(node.Value[1].(*Node), offset)
	case NT_VAR:
		return node.Value[0].String()
	case NT_DEFINE_VAR:
		if node.Value[0].(*Node).NodeType == NT_SET {
			return GetTP(this.CompileNode(node.Value[0].(*Node).Value[0].(*Node).Value[1].(*Node), offset)) + " " + this.CompileNode(node.Value[0].(*Node).Value[0].(*Node).Value[0].(*Node), offset) + " = " + this.CompileNode(node.Value[0].(*Node).Value[1].(*Node), offset)
		} else if node.Value[0].(*Node).NodeType == NT_SET2 {
			return GetTP(this.CompileNode(node.Value[0].(*Node).Value[1].(*Node), offset)) + " " + this.CompileNode(node.Value[0].(*Node).Value[0].(*Node), offset)
		}
	case NT_SCOPE:
		v := []*Node{}
		for _, vv := range node.Value {
			v = append(v, vv.(*Node))
		}
		return "{\n" + compile(v, offset+"    ") + "\n" + offset + "}"
	default:
		return ""
	}
	return ""
}
func compile(nodes []*Node, offset string) string {
	ret := ""
	c := Compiler{}
	for i := range nodes {
		if i > 0 {
			ret += "\n"
		}
		ret += offset + c.CompileNode(nodes[i], offset) + ";"
	}
	return ret
}
func Compile(nodes []*Node) string {
	return compile(nodes, "")
}
func NewCompiler() *Compiler {
	return &Compiler{}
}
