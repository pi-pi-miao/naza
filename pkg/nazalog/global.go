// Copyright 2019, Chef.  All rights reserved.
// https://github.com/q191201771/naza
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

package nazalog

import (
	"fmt"

	"github.com/q191201771/naza/pkg/fake"
)

var global Logger

func Outputf(level Level, calldepth int, format string, v ...interface{}) {
	global.Out(level, 3, fmt.Sprintf(format, v...))
}

func Debugf(format string, v ...interface{}) {
	global.Out(LevelDebug, 3, fmt.Sprintf(format, v...))
}

func Infof(format string, v ...interface{}) {
	global.Out(LevelInfo, 3, fmt.Sprintf(format, v...))
}

func Warnf(format string, v ...interface{}) {
	global.Out(LevelWarn, 3, fmt.Sprintf(format, v...))
}

func Errorf(format string, v ...interface{}) {
	global.Out(LevelError, 3, fmt.Sprintf(format, v...))
}

func Fatalf(format string, v ...interface{}) {
	global.Out(LevelFatal, 3, fmt.Sprintf(format, v...))
	fake.Exit(1)
}

func Panicf(format string, v ...interface{}) {
	global.Out(LevelPanic, 3, fmt.Sprintf(format, v...))
}

func Output(level Level, calldepth int, v ...interface{}) {
	global.Out(level, 3, fmt.Sprint(v...))
}

func Debug(v ...interface{}) {
	global.Out(LevelDebug, 3, fmt.Sprint(v...))
}

func Info(v ...interface{}) {
	global.Out(LevelInfo, 3, fmt.Sprint(v...))
}

func Warn(v ...interface{}) {
	global.Out(LevelWarn, 3, fmt.Sprint(v...))
}

func Error(v ...interface{}) {
	global.Out(LevelError, 3, fmt.Sprint(v...))
}

func Fatal(v ...interface{}) {
	global.Out(LevelFatal, 3, fmt.Sprint(v...))
	fake.Exit(1)
}

func Panic(v ...interface{}) {
	global.Out(LevelPanic, 3, fmt.Sprint(v...))
}

func FatalIfErrorNotNil(err error) {
	if err != nil {
		global.Out(LevelError, 3, fmt.Sprintf("fatal since error not nil. err=%+v", err))
		fake.Exit(1)
	}
}

func PanicIfErrorNotNil(err error) {
	if err != nil {
		global.Out(LevelPanic, 3, fmt.Sprintf("fatal since error not nil. err=%+v", err))
		panic(err)
	}
}

func Out(level Level, calldepth int, s string) {
	global.Out(level, calldepth, s)
}

// 这里不加锁保护，如果要调用Init函数初始化全局的Logger，那么由调用方保证调用Init函数时不会并发调用全局Logger的其他方法
func Init(modOptions ...ModOption) error {
	var err error
	global, err = New(modOptions...)
	return err
}

func init() {
	_ = Init()
}
