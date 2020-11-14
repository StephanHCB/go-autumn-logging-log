package contextawarelogging

import (
	"github.com/StephanHCB/go-autumn-logging-log/implementation/leveledlogging"
	"github.com/StephanHCB/go-autumn-logging/api"
)

type PlainContextAwareLoggingImplementation struct{
	RequestId string
}

func (ca *PlainContextAwareLoggingImplementation) Trace() auloggingapi.LeveledLoggingImplementation {
	return &leveledlogging.PlainLeveledLoggingImplementation{LogLevel: "TRACE", RequestId: ca.RequestId}
}

func (ca *PlainContextAwareLoggingImplementation) Debug() auloggingapi.LeveledLoggingImplementation {
	return &leveledlogging.PlainLeveledLoggingImplementation{LogLevel: "DEBUG", RequestId: ca.RequestId}
}

func (ca *PlainContextAwareLoggingImplementation) Info() auloggingapi.LeveledLoggingImplementation {
	return &leveledlogging.PlainLeveledLoggingImplementation{LogLevel: "INFO ", RequestId: ca.RequestId}
}

func (ca *PlainContextAwareLoggingImplementation) Warn() auloggingapi.LeveledLoggingImplementation {
	return &leveledlogging.PlainLeveledLoggingImplementation{LogLevel: "WARN ", RequestId: ca.RequestId}
}

func (ca *PlainContextAwareLoggingImplementation) Error() auloggingapi.LeveledLoggingImplementation {
	return &leveledlogging.PlainLeveledLoggingImplementation{LogLevel: "ERROR", RequestId: ca.RequestId}
}

func (ca *PlainContextAwareLoggingImplementation) Fatal() auloggingapi.LeveledLoggingImplementation {
	return &leveledlogging.PlainLeveledLoggingImplementation{LogLevel: "FATAL", AbortAfter: true, RequestId: ca.RequestId}
}

func (ca *PlainContextAwareLoggingImplementation) Panic() auloggingapi.LeveledLoggingImplementation {
	return &leveledlogging.PlainLeveledLoggingImplementation{LogLevel: "PANIC", AbortAfter: true, RequestId: ca.RequestId}
}
