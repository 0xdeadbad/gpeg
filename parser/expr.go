package parser

import "golang.org/x/exp/constraints"

type Expr interface {
}

type ExprBase struct {
	name string
}

type RangeExpr[T constraints.Ordered] struct {
	ExprBase
	min T
	max T
}
