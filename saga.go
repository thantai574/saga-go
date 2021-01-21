package saga

import (
	"context"
	"saga/entities"
	"saga/strategies/memory"
)

type (
	TypeSaga int

	Options struct {
		ISaga
		TypeSaga
	}

	Func       entities.FuncSaga
	Compensate func(ctx context.Context) error
)

const (
	TypeKafka TypeSaga = iota
	TypeMemory
)

type ISaga interface {
	AddStep(step entities.Step) error
	GetAllSteps() []*entities.Step
	Exec(context.Context, int) *entities.Step
	Draw(context.Context, int) *entities.Step
}

type saga struct {
	ISaga
	Options
}

func NewSaga(o Options) *saga {
	sg := new(saga)
	switch o.TypeSaga {
	case TypeKafka:

	case TypeMemory:
		sg.ISaga = memory.NewMemory()
	}
	return sg
}
