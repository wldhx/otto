package hashicorp

import (
	"github.com/hashicorp/otto/foundation"
	"github.com/hashicorp/otto/plan"
)

type Foundation struct{}

func (f *Foundation) Compile(*foundation.Context) (*foundation.CompileResult, error) {
	return nil, nil
}

func (f *Foundation) Plan(*foundation.Context) ([]*plan.Plan, error) {
	return nil, nil
}
