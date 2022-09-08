package gotterfly

import "testing"

func TestLogs(t *testing.T) {
	var logger = WithPrefix("TEST")
	logger.Timer("TIMER TEST")
	logger.Debugf("Debug output")
	logger.Infof("Info output")
	logger.Warnf("Warning output")
	logger.Trace("Trace output")
	logger.Errorf("Error output")
	logger.Timer("TIMER TEST")
}
