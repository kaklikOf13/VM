package vm

const (
	NT_VALUE = "Value"
	NT_PLUS  = "Plus"
	NT_MINUS = "Minus"
	NT_MUL   = "Mul"
	NT_DIV   = "Div"
)

type Parser struct {
	curTok   *Token
	curBlock []Token
	tok      int
	block    int
	Blocks   [][]Token
}

func (this *Parser) NextTok() {
	if this.tok >= len(this.curBlock) {
		this.curTok = nil
	} else {
		this.curTok = &this.curBlock[this.tok]
		this.tok++
	}
}
func (this *Parser) NextBlock() {
	if this.block >= len(this.Blocks) {
		this.curBlock = nil
	} else {
		this.curBlock = this.Blocks[this.block]
		this.NextTok()
		this.block++
	}
}
func (this *Parser) value() *Node {
	oldt := this.curTok
	if oldt != nil {
		this.NextTok()
		switch oldt.Type {
		case TT_VALUE:
			return &Node{NodeType: NT_VALUE, Value: []Value{oldt.Value}}
		}
	}
	return &Node{NodeType: NT_VALUE, Value: []Value{&NullType{}}}
}
func (this *Parser) tier1() *Node {
	node := this.value()
	for this.curTok != nil && containsString([]string{TT_MUL, TT_DIV}, this.curTok.Type) {
		oldt := *this.curTok
		this.NextTok()
		switch oldt.Type {
		case TT_MUL:
			node = &Node{NodeType: NT_MUL, Value: []Value{node, this.value()}}
		case TT_DIV:
			node = &Node{NodeType: NT_DIV, Value: []Value{node, this.value()}}
		}
	}
	return node
}
func (this *Parser) Main() *Node {
	node := this.tier1()
	for this.curTok != nil && containsString([]string{TT_PLUS, TT_MINUS}, this.curTok.Type) {
		oldt := *this.curTok
		this.NextTok()
		switch oldt.Type {
		case TT_PLUS:
			node = &Node{NodeType: NT_PLUS, Value: []Value{node, this.value()}}
		case TT_MINUS:
			node = &Node{NodeType: NT_MINUS, Value: []Value{node, this.value()}}
		}
	}
	return node
}
func Parse(toks []Token) []*Node {
	p := Parser{}
	ret := []*Node{}
	p.Blocks = splitWithSeparators(toks, []string{TT_BREAKPOINT, TT_NEWLINE})
	p.NextBlock()
	for p.curBlock != nil {
		ret = append(ret, p.Main())
		p.NextBlock()
	}
	return ret
}
