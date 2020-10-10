package go_autumn_logging_log

import (
	aulogging "github.com/StephanHCB/go-autumn-logging"
	"github.com/StephanHCB/go-autumn-logging-log/implementation/logging"
)

func init() {
	aulogging.Logger = &logging.PlainLoggingImplementation{}
}