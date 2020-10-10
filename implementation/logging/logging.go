package logging

import (
	"context"
	"github.com/StephanHCB/go-autumn-logging-log/implementation/contextawarelogging"
	auloggingapi "github.com/StephanHCB/go-autumn-logging/api"
)

type PlainLoggingImplementation struct{}

func (l *PlainLoggingImplementation) Ctx(ctx context.Context) auloggingapi.ContextAwareLoggingImplementation {
	// this implementation completely ignores the context for now
	return &contextawarelogging.PlainContextAwareLoggingImplementation{}
}

func (l *PlainLoggingImplementation) NoCtx() auloggingapi.ContextAwareLoggingImplementation {
	return &contextawarelogging.PlainContextAwareLoggingImplementation{}
}
