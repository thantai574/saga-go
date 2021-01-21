package entities

type StepIndex int

func (s StepIndex) SetInt(step int) StepIndex {
	return StepIndex(step)
}

func (s StepIndex) GetInt() int {
	return int(s)
}

func (s StepIndex) IncrementStep() int {
	return s.GetInt() + 1
}

func (s StepIndex) DecrementStep() int {
	if s.GetInt()-1 < 0 {
		return 0
	}
	return s.GetInt() - 1
}
