package types

type ServiceAction string

const (
	ActionContinue ServiceAction = "continue"
	ActionStop     ServiceAction = "stop"
)
