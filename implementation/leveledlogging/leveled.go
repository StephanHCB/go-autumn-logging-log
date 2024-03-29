package leveledlogging

import (
	"context"
	"fmt"
	aulogging "github.com/StephanHCB/go-autumn-logging"
	auloggingapi "github.com/StephanHCB/go-autumn-logging/api"
	"log"
	"os"
)

type PlainLeveledLoggingImplementation struct {
	LogLevel         string
	RequestId        string
	AbortAfter       bool
	Err              error
	AdditionalFields map[string]string
	Ctx              context.Context
}

func (lv *PlainLeveledLoggingImplementation) WithErr(err error) auloggingapi.LeveledLoggingImplementation {
	lv.Err = err
	return lv
}

func (lv *PlainLeveledLoggingImplementation) With(key string, value string) auloggingapi.LeveledLoggingImplementation {
	if lv.AdditionalFields == nil {
		lv.AdditionalFields = make(map[string]string)
	}
	lv.AdditionalFields[key] = value
	return lv
}

func (lv *PlainLeveledLoggingImplementation) addFieldsAndErrOutput() []interface{} {
	result := []interface{}{}
	if lv.Err != nil {
		result = append(result, fmt.Sprintf("error=%s", lv.Err.Error()))
	}
	if lv.AdditionalFields != nil {
		for k, v := range lv.AdditionalFields {
			result = append(result, fmt.Sprintf("%s=%s", k, v))
		}
	}
	return result
}

func (lv *PlainLeveledLoggingImplementation) Print(v ...interface{}) {
	message := fmt.Sprint(v...)
	if aulogging.LogEventCallback != nil {
		aulogging.LogEventCallback(lv.Ctx, lv.LogLevel, message, lv.Err, lv.AdditionalFields)
	}
	if lv.Err != nil || lv.AdditionalFields != nil {
		log.Print(lv.RequestId + lv.LogLevel + " " + message + " " + fmt.Sprint(lv.addFieldsAndErrOutput()...))
	} else {
		log.Print(lv.RequestId + lv.LogLevel + " " + message)
	}
	if lv.AbortAfter {
		os.Exit(1)
	}
}

func (lv *PlainLeveledLoggingImplementation) Printf(format string, v ...interface{}) {
	lv.Print(fmt.Sprintf(format, v...))
}
