// Copyright 2019, Chef.  All rights reserved.
// https://github.com/q191201771/naza
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

package taskpool

var global Pool

func Go(task Task) {
	global.Go(task)
}

func Status() (idleWorkerNum int, busyWorkerNum int) {
	return global.Status()
}

func KillIdleWorkers() {
	global.KillIdleWorkers()
}

func Init(modOptions ...ModOption) error {
	var err error
	global, err = NewPool(modOptions...)
	return err
}

func init() {
	_ = Init()
}
