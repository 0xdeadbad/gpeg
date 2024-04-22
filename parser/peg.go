package parser

type PEG struct {
	stk []*Rule
}

func NewPEG() *PEG {
	return &PEG{
		stk: make([]*Rule, 0),
	}
}

func (p *PEG) Push(r *Rule) {
	p.stk = append(p.stk, r)
}

func (p *PEG) Pop() *Rule {
	if len(p.stk) == 0 {
		return nil
	}

	defer func() { p.stk = p.stk[:len(p.stk)-1] }()
	r := p.stk[len(p.stk)-1]

	return r
}

func (p *PEG) Peek() *Rule {
	if len(p.stk) == 0 {
		return nil
	}

	return p.stk[len(p.stk)-1]
}
