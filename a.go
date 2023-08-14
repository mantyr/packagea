package packagea

import (
	"github.com/mantyr/packageb"
)

type A struct {
	B packageb.B
}

func (a *A) GetTest() error {
	return a.B.Check()
}