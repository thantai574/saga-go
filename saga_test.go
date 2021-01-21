package saga_go

import (
	"context"
	"fmt"
	"github.com/thantai574/saga-go/entities"
	"testing"
)

func TestSaga(t *testing.T) {
	tests := []struct {
		name      string
		want      string
		errorStep []error
	}{
		{
			name:      "HAPPY",
			errorStep: []error{nil, nil},
		},
		{
			name:      "error_step1",
			errorStep: []error{fmt.Errorf("test"), nil},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sg := NewSaga(Options{
				TypeSaga: TypeMemory,
			})

			c := context.TODO()
			sg.AddStep(entities.Step{
				Func: func(ctx context.Context) error {
					c = context.WithValue(c, "test", "ttes")
					return tt.errorStep[0]
				},
				Compensate: func(ctx context.Context) error {
					fmt.Println("Cancel Step 1 ")
					return nil
				},
			})

			sg.AddStep(entities.Step{
				Func: func(ctx context.Context) error {
					fmt.Println(c.Value("test"))
					return tt.errorStep[1]
				},
				Compensate: func(ctx context.Context) error {
					fmt.Println("Cancel Step 2 ")
					return nil
				},
			})
			coor := NewCoordinator(c, sg)

			coor.Start()
		})
	}
}
