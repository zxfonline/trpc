package arith

type Arith interface {
	Add(a, b int) (result int, err error)
	Mul(a, b int) (result int, err error)
	Div(a, b int) (result int, err error)
}
