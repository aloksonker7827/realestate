package logger

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"gitlab.com/umi7/DezHab_user_backend/utils/constants"
	"os"
	"runtime"
	"strings"
)

var logger = logrus.New()

const (
	repoPath = "gitlab.com/umi7/DezHab_user_backend/"
	gormLog  = "gorm"
)

func init() {
	logger.SetReportCaller(true)
	logger.Formatter = getFormatter(8)
}

// Info is used to log the information
//
// This should be called for tracking the events happened in the application.
func Info(ctx context.Context, args ...interface{}) {
	logger.Info(messageWithContext(ctx, args))
}

// Error is used to log the error.
//
// This should be called from the place where the error is originating
// Do not call this at all the places when err != nil
// As we wil be logging too many duplicate information.
//
// We don't need to add noticeError of new_relic here
// as that is automatically handled by transaction.
// i.e : When the response status code is not success status code
// noticeError function will be called by new_relic transaction.
func Error(ctx context.Context, args ...interface{}) {
	logger.Error(messageWithContext(ctx, args))
}

// WithAlert is used to log and monitor the error.
//
// This should be called when there is a panic in the application.
//
// Using new_relic transaction, we are adding a custom attribute with
// which alerts will be created.
func WithAlert(ctx context.Context, err interface{}, args ...interface{}) {
	if err, ok := err.(error); ok {
		logger.Error(err)
	}
	Error(ctx, err, args)
}

// Fatal should be used when there is a possibility for
// un-recoverable panics.
func Fatal(ctx context.Context, args ...interface{}) {
	logger.Fatal(messageWithContext(ctx, args))
}

// Debug should be used to print the entire stack trace
// of the application.
//
// If some issues is non-reproducible in lower level environments,
// this can be used.
func Debug(ctx context.Context, args ...interface{}) {
	logger.Debug(messageWithContext(ctx, args))
}

// messageWithContext adds the requestId to arguments at the first,
// which will be fetched from context.
func messageWithContext(ctx context.Context, args ...interface{}) interface{} {
	if ctx.Value(constants.Provider) != nil {
		args = append([]interface{}{ctx.Value(constants.Provider)}, args...)
	}
	if ctx.Value(constants.RequestID) != nil {
		args = append([]interface{}{ctx.Value(constants.RequestID)}, args...)
	}
	return args
}

var getFormatter = func(skip int) *logrus.JSONFormatter {
	return &logrus.JSONFormatter{
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			defer func() { recover() }()
			//Note : This is exclusively used for dezhab user.
			//If anyone wants to re-use this, they need to make sure that there are
			//at least 8 calls in stack trace.
			//We are skipping 8 calls, as all of those calls are made in the vendor folder.
			_, fn, line, _ := runtime.Caller(skip)
			repoPath := fmt.Sprintf("%s/src/%s", os.Getenv("GOPATH"), repoPath)
			fileName := strings.Replace(fn, repoPath, "", -1)
			return "", fmt.Sprintf("%s:%d", fileName, line)
		},
	}
}

var GetDBLogger = func() *logrus.Logger {
	dbLogger := logrus.New()
	dbLogger.SetReportCaller(true)
	dbLogger.Formatter = &logrus.JSONFormatter{
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			return "", gormLog
		},
	}
	return dbLogger
}
