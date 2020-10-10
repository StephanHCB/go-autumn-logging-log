package contextawarelogging

import (
	"github.com/StephanHCB/go-autumn-logging-log/implementation/leveledlogging"
	"github.com/StephanHCB/go-autumn-logging/api"
)

type PlainContextAwareLoggingImplementation struct{
}

func (ca *PlainContextAwareLoggingImplementation) Trace() auloggingapi.LeveledLoggingImplementation {
	return &leveledlogging.PlainLeveledLoggingImplementation{LogLevel: "TRACE"}
}

func (ca *PlainContextAwareLoggingImplementation) Debug() auloggingapi.LeveledLoggingImplementation {
	return &leveledlogging.PlainLeveledLoggingImplementation{LogLevel: "DEBUG"}
}

func (ca *PlainContextAwareLoggingImplementation) Info() auloggingapi.LeveledLoggingImplementation {
	return &leveledlogging.PlainLeveledLoggingImplementation{LogLevel: "INFO"}
}

func (ca *PlainContextAwareLoggingImplementation) Warn() auloggingapi.LeveledLoggingImplementation {
	return &leveledlogging.PlainLeveledLoggingImplementation{LogLevel: "WARN"}
}

func (ca *PlainContextAwareLoggingImplementation) Error() auloggingapi.LeveledLoggingImplementation {
	return &leveledlogging.PlainLeveledLoggingImplementation{LogLevel: "ERROR"}
}

func (ca *PlainContextAwareLoggingImplementation) Fatal() auloggingapi.LeveledLoggingImplementation {
	return &leveledlogging.PlainLeveledLoggingImplementation{LogLevel: "FATAL", AbortAfter: true}
}

func (ca *PlainContextAwareLoggingImplementation) Panic() auloggingapi.LeveledLoggingImplementation {
	return &leveledlogging.PlainLeveledLoggingImplementation{LogLevel: "PANIC", AbortAfter: true}
}
