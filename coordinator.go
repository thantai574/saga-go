package saga_go

import (
	"context"
	"saga_go/entities"
)

type coordinator struct {
	*saga
	ctx          context.Context
	replyChannel chan entities.Step
}

func (c *coordinator) replyConsumer() (err error) {
consumerLoop:
	for v := range c.replyChannel {
		if len(c.GetAllSteps())-1 == v.StepIndex.GetInt() && v.StateStep.IsSuccess() {
			err = v.Error
			break consumerLoop
		}

		if v.StepIndex.GetInt() == 0 && v.StateStep.IsAbort() {
			break consumerLoop
		}

		if v.Error == nil {
			c.saga.Exec(c.ctx, v.StepIndex.IncrementStep())
		} else {
			err = v.Error
			c.saga.Draw(c.ctx, v.StepIndex.DecrementStep())
		}

	}

	return
}

func (c *coordinator) setAddressReplySteps() {
	for _, step := range c.GetAllSteps() {
		step.StepChannel = c.replyChannel
	}
}

func (c *coordinator) Start() (err error) {
	c.setAddressReplySteps()
	c.Exec(c.ctx, 0)
	err = c.replyConsumer()
	return
}

func NewCoordinator(ctx context.Context, s *saga) *coordinator {
	return &coordinator{
		saga:         s,
		ctx:          ctx,
		replyChannel: make(chan entities.Step, len(s.GetAllSteps())),
	}
}
