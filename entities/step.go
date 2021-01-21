package entities

type Step struct {
	StepIndex
	Func       FuncSaga
	Compensate FuncSaga
	StateStep
	StepChannel chan Step
	Error       error
}
