package contextawarelogging

import (
	"context"
	"github.com/StephanHCB/go-autumn-logging-log/implementation/leveledlogging"
	"github.com/StephanHCB/go-autumn-logging/api"
)

type PlainContextAwareLoggingImplementation struct {
	RequestId string
	Ctx       context.Context
}

func (ca *PlainContextAwareLoggingImplementation) Trace() auloggingapi.LeveledLoggingImplementation {
	return &leveledlogging.PlainLeveledLoggingImplementation{LogLevel: "TRACE", RequestId: ca.RequestId, Ctx: ca.Ctx}
}

func (ca *PlainContextAwareLoggingImplementation) Debug() auloggingapi.LeveledLoggingImplementation {
	return &leveledlogging.PlainLeveledLoggingImplementation{LogLevel: "DEBUG", RequestId: ca.RequestId, Ctx: ca.Ctx}
}

func (ca *PlainContextAwareLoggingImplementation) Info() auloggingapi.LeveledLoggingImplementation {
	return &leveledlogging.PlainLeveledLoggingImplementation{LogLevel: "INFO ", RequestId: ca.RequestId, Ctx: ca.Ctx}
}

func (ca *PlainContextAwareLoggingImplementation) Warn() auloggingapi.LeveledLoggingImplementation {
	return &leveledlogging.PlainLeveledLoggingImplementation{LogLevel: "WARN ", RequestId: ca.RequestId, Ctx: ca.Ctx}
}

func (ca *PlainContextAwareLoggingImplementation) Error() auloggingapi.LeveledLoggingImplementation {
	return &leveledlogging.PlainLeveledLoggingImplementation{LogLevel: "ERROR", RequestId: ca.RequestId, Ctx: ca.Ctx}
}

func (ca *PlainContextAwareLoggingImplementation) Fatal() auloggingapi.LeveledLoggingImplementation {
	return &leveledlogging.PlainLeveledLoggingImplementation{LogLevel: "FATAL", AbortAfter: true, RequestId: ca.RequestId, Ctx: ca.Ctx}
}

func (ca *PlainContextAwareLoggingImplementation) Panic() auloggingapi.LeveledLoggingImplementation {
	return &leveledlogging.PlainLeveledLoggingImplementation{LogLevel: "PANIC", AbortAfter: true, RequestId: ca.RequestId, Ctx: ca.Ctx}
}
