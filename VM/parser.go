package vm

const (
	NT_VALUE      = "Value"
	NT_PLUS       = "Plus"
	NT_MINUS      = "Minus"
	NT_MUL        = "Mul"
	NT_DIV        = "Div"
	NT_VAR        = "Var"
	NT_DEFINE_VAR = "Def Var"
	NT_SET        = "Set"
	NT_SET2       = "Set2"
	NT_SCOPE      = "Scope"
	NT_FUNCTION   = "Function"
	NT_ARGS       = "Args"
	NT_CALL       = "Call"
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
		this.tok = 0
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
		case TT_VAR:
			return &Node{NodeType: NT_VAR, Value: []Value{oldt.Value}}
		case TT_DEFINE_VAR:
			return &Node{NodeType: NT_DEFINE_VAR, Value: []Value{this.Main()}}
		case TT_FUNCTION:
			if this.curTok.Type == TT_VAR {
				n := &Node{NullType{}, []Value{this.Name(), this.GetArgs(), nil}, NT_FUNCTION}
				if this.curTok.Type == TT_DOUBLE_POINT {
					this.NextTok()
					n.Value[len(n.Value)-1] = this.Type()
				}
				n.Value = append(n.Value, this.GetScope())
				return n
			} else {
				n := &Node{NullType{}, []Value{this.GetArgs(), nil}, NT_FUNCTION}
				if this.curTok.Type == TT_DOUBLE_POINT {
					this.NextTok()
					n.Value[len(n.Value)-1] = this.Type()
				}
				if this.curTok.Type == TT_LKEY {
					n.Value = append(n.Value, this.GetScope())
				}
				return n
			}

		}
	}
	return &Node{NodeType: NT_VALUE, Value: []Value{&NullType{}}}
}
func (this *Parser) Name() *Node {
	node := this.value()
	return node
}
func (this *Parser) Type() *Node {
	node := this.Name()
	return node
}
func (this *Parser) GetScope() *Node {
	this.NextTok()
	ret := []Value{}
	toks := [][]Token{}
	v := uint(1)
	block := []Token{}
	for v > 0 && this.curTok != nil {
		if this.curTok.Type == TT_LKEY {
			v++
		} else if this.curTok.Type == TT_RKEY {
			v--
			if v == 0 {
				toks = append(toks, block)
				this.NextTok()
				break
			}
		}
		block = append(block, *this.curTok)
		this.NextTok()
		if this.curTok == nil {
			this.NextBlock()
			toks = append(toks, block)
			block = []Token{}
			if this.curBlock == nil {
				break
			}
		}
	}
	val := ParseBlocks(toks)
	for _, vv := range val {
		ret = append(ret, vv)
	}
	return &Node{NodeType: NT_SCOPE, Value: ret}
}
func (this *Parser) GetArgs() *Node {
	ret := []Value{}
	toks := []Token{}
	v := uint(1)
	this.NextTok()
	for v > 0 && this.curTok != nil {
		if this.curTok.Type == TT_LPAREN {
			v++
		} else if this.curTok.Type == TT_RPAREN {
			v--
			if v == 0 {
				this.NextTok()
				break
			}
		}
		toks = append(toks, *this.curTok)
		this.NextTok()
		if this.curTok == nil {
			this.NextBlock()
			if this.curBlock == nil {
				break
			}
		}
	}
	val := parseTypes(splitWithSeparators(toks, []string{TT_VIRGULA}), true)
	for _, vv := range val {
		ret = append(ret, vv)
	}
	return &Node{NodeType: NT_ARGS, Value: ret}
}
func (this *Parser) GetType() *Node {
	node := this.tier2()
	for this.curTok != nil && containsString([]string{TT_DOUBLE_POINT}, this.curTok.Type) {
		oldt := *this.curTok
		this.NextTok()
		switch oldt.Type {
		case TT_DOUBLE_POINT:
			node = &Node{NodeType: NT_SET2, Value: []Value{node, this.tier2()}}
		}
	}
	return node
}
func (this *Parser) tier2() *Node {
	node := this.tier3()
	for this.curTok != nil && containsString([]string{TT_PLUS, TT_MINUS}, this.curTok.Type) {
		oldt := *this.curTok
		this.NextTok()
		switch oldt.Type {
		case TT_PLUS:
			node = &Node{NodeType: NT_PLUS, Value: []Value{node, this.tier2()}}
		case TT_MINUS:
			node = &Node{NodeType: NT_MINUS, Value: []Value{node, this.tier2()}}
		}
	}
	return node
}
func (this *Parser) tier3() *Node {
	node := this.Type()
	for this.curTok != nil && containsString([]string{TT_MUL, TT_DIV}, this.curTok.Type) {
		oldt := *this.curTok
		this.NextTok()
		switch oldt.Type {
		case TT_MUL:
			node = &Node{NodeType: NT_MUL, Value: []Value{node, this.tier3()}}
		case TT_DIV:
			node = &Node{NodeType: NT_DIV, Value: []Value{node, this.tier3()}}
		}
	}
	return node
}
func (this *Parser) Main() *Node {
	node := this.GetType()
	for this.curTok != nil && containsString([]string{TT_EQUAL}, this.curTok.Type) {
		oldt := *this.curTok
		this.NextTok()
		switch oldt.Type {
		case TT_EQUAL:
			node = &Node{NodeType: NT_SET, Value: []Value{node, this.GetType()}}
		}
	}
	return node
}
func Parse(toks []Token) []*Node {
	return ParseBlocks(splitWithSeparators(toks, []string{TT_BREAKPOINT, TT_NEWLINE}))
}
func ParseBlocks(toks [][]Token) []*Node {
	p := Parser{}
	ret := []*Node{}
	p.Blocks = toks
	p.NextBlock()
	for p.curBlock != nil {
		ret = append(ret, p.Main())
		p.NextBlock()
	}
	return ret
}
func parseTypes(toks [][]Token, ok bool) []*Node {
	p := Parser{}
	ret := []*Node{}
	p.Blocks = toks
	p.NextBlock()
	oldt := "int"
	for p.curBlock != nil {
		ret = append(ret, p.GetType())
		if ret[len(ret)-1].NodeType == NT_VAR && ok {
			ret[len(ret)-1] = &Node{NullType{}, []Value{ret[len(ret)-1], &Node{NullType{}, []Value{&String{value: oldt}}, NT_VAR}}, NT_SET2}
		}
		p.NextBlock()
	}
	return ret
}
