package logging

import (
	"context"
	aulogging "github.com/StephanHCB/go-autumn-logging"
	"github.com/StephanHCB/go-autumn-logging-log/implementation/contextawarelogging"
	auloggingapi "github.com/StephanHCB/go-autumn-logging/api"
)

type PlainLoggingImplementation struct{}

func (l *PlainLoggingImplementation) Ctx(ctx context.Context) auloggingapi.ContextAwareLoggingImplementation {
	requestId := ""
	if aulogging.RequestIdRetriever != nil {
		requestId = "[" + aulogging.RequestIdRetriever(ctx) + "] "
	}
	return &contextawarelogging.PlainContextAwareLoggingImplementation{RequestId: requestId}
}

func (l *PlainLoggingImplementation) NoCtx() auloggingapi.ContextAwareLoggingImplementation {
	requestId := ""
	if aulogging.RequestIdRetriever != nil {
		requestId = "[" + aulogging.RequestIdRetriever(context.TODO()) + "] "
	}
	return &contextawarelogging.PlainContextAwareLoggingImplementation{RequestId: requestId}
}
