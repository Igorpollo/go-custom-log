package logo

import "testing"

func TestLog(t * testing.T) {
    Info("testing...");
    Warning("testing...");
    Error("testing...");
    Success("testing...");
	DebugMode(true)
	Debug("testing...");
}