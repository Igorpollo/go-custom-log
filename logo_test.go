package logo

import "testing"

func TestLog(t *testing.T) {
	Info("Testing Info message...")
	Warning("Testing Warning message...")
	Error("Testing Error message...")
	Success("Testing Success message...")
	DebugMode(true)
	Debug("Testing Debug message...")
	ShowTime(true)
	Info("Testing Info message with time...")
	Debug("Testing Info message with time...")
}
