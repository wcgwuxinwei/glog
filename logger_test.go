package glog

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGlog_SetLevel(t *testing.T) {
	l := New(false)

	l.SetLevel(LogLevelError)
	assert.Nilf(t, l.SetLevel(LogLevelDebug), "Want: nil, Got: %s", l.SetLevel(LogLevelDebug))
	assert.Nilf(t, l.SetLevel(LogLevelInfo), "Want: nil, Got: %s", l.SetLevel(LogLevelInfo))
	assert.Nilf(t, l.SetLevel(LogLevelWarn), "Want: nil, Got: %s", l.SetLevel(LogLevelWarn))
	assert.Nilf(t, l.SetLevel(LogLevelError), "Want: nil, Got: %s", l.SetLevel(LogLevelError))
}

func TestGlog_Debugf(t *testing.T) {
	l := New(false)

	// Test debug level
	l.SetLevel(LogLevelDebug)
	assert.Equal(t, "[LOG][DEBUG]: test", l.Debugf("test"), "them should be equal")
	assert.Equal(t, "[LOG][INFO]: test", l.Infof("test"), "them should be equal")
	assert.Equal(t, "[LOG][WARN]: test", l.Warnf("test"), "them should be equal")
	assert.Equal(t, "[LOG][ERROR]: test", l.Errorf("test"), "them should be equal")
}

func TestGlog_Infof(t *testing.T) {
	l := New(false)

	// Test info level
	l.SetLevel(LogLevelInfo)
	assert.Equal(t, "", l.Debugf("test"), "them should be equal")
	assert.Equal(t, "[LOG][INFO]: test", l.Infof("test"), "them should be equal")
	assert.Equal(t, "[LOG][WARN]: test", l.Warnf("test"), "them should be equal")
	assert.Equal(t, "[LOG][ERROR]: test", l.Errorf("test"), "them should be equal")
}

func TestGlog_Warnf(t *testing.T) {
	l := New(false)

	// Test warn level
	l.SetLevel(LogLevelWarn)
	assert.Equal(t, "", l.Debugf("test"), "them should be equal")
	assert.Equal(t, "", l.Infof("test"), "them should be equal")
	assert.Equal(t, "[LOG][WARN]: test", l.Warnf("test"), "them should be equal")
	assert.Equal(t, "[LOG][ERROR]: test", l.Errorf("test"), "them should be equal")
}

func TestGlog_Errorf(t *testing.T) {
	l := New(false)

	// Test error level
	l.SetLevel(LogLevelError)
	assert.Equal(t, "", l.Debugf("test"), "them should be equal")
	assert.Equal(t, "", l.Infof("test"), "them should be equal")
	assert.Equal(t, "", l.Warnf("test"), "them should be equal")
	assert.Equal(t, "[LOG][ERROR]: test", l.Errorf("test"), "them should be equal")
}

func TestGlog_ReloadLogLevel(t *testing.T) {
	l := New(true)

	// Test error level
	l.SetLevel(LogLevelDebug)
	assert.Equal(t, "[LOG][DEBUG]: test", l.Debugf("test"), "them should be equal")

	// Test change the log level when running
	os.Setenv("LOG_LEVEL", "ERROR")
	time.Sleep(11 * time.Second)
	assert.Equal(t, "", l.Debugf("test"), "them should be equal")
	assert.Equal(t, "[LOG][ERROR]: test", l.Errorf("test"), "them should be equal")

	os.Setenv("LOG_LEVEL", "DEBUG")
	time.Sleep(11 * time.Second)
	assert.Equal(t, "[LOG][DEBUG]: test", l.Debugf("test"), "them should be equal")
	assert.Equal(t, "[LOG][ERROR]: test", l.Errorf("test"), "them should be equal")
}
