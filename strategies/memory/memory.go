package memory

import (
	"context"
	"saga_go/entities"
)

type SagaMemory struct {
	Steps       []*entities.Step
	currentStep int
}

func (s *SagaMemory) AddStep(step entities.Step) error {
	step.StepIndex = step.StepIndex.SetInt(len(s.Steps))
	s.Steps = append(s.Steps, &step)
	return nil
}

func (s *SagaMemory) reply() {
	cur := s.Steps[s.currentStep]
	if cur != nil {
		s.Steps[s.currentStep].StepChannel <- *cur
	}

}
func (s *SagaMemory) GetAllSteps() []*entities.Step {
	return s.Steps
}

func (s *SagaMemory) Exec(ctx context.Context, step int) *entities.Step {
	s.currentStep = step
	cur := s.Steps[step]

	err := cur.Func(ctx)

	cur.Error = err
	cur.StateStep = entities.StateSuccess
	s.reply()
	return cur
}

func (s *SagaMemory) Draw(ctx context.Context, step int) *entities.Step {
	s.currentStep = step
	cur := s.Steps[step]

	err := cur.Compensate(ctx)
	cur.Error = err
	cur.StateStep = entities.StateAbort
	s.reply()
	return cur
}

func NewMemory() *SagaMemory {
	return &SagaMemory{
		Steps: []*entities.Step{},
	}
}
