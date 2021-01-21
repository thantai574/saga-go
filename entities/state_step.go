package entities

const (
	StatePending StateStep = iota
	StateStart
	StateAbort
	StateSuccess
)

type StateStep int

func (s StateStep) IsSuccess() bool {
	return s == StateSuccess
}

func (s StateStep) IsAbort() bool {
	return s == StateAbort
}
