// Copyright 2019, Chef.  All rights reserved.
// https://github.com/q191201771/naza
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

package mockwriter

import (
	"testing"

	"github.com/q191201771/naza/pkg/assert"
)

func TestNewMockWriter(t *testing.T) {
	_ = NewMockWriter(WriterTypeDoNothing)
}

func TestMockWriter_Write(t *testing.T) {
	var (
		w   *MockWriter
		n   int
		err error
		b   = []byte("hello")
	)

	w = NewMockWriter(WriterTypeDoNothing)
	n, err = w.Write(b)
	assert.Equal(t, 5, n)
	assert.Equal(t, nil, err)

	w = NewMockWriter(WriterTypeReturnError)
	n, err = w.Write(b)
	assert.Equal(t, 0, n)
	assert.Equal(t, mockWriterErr, err)

	w = NewMockWriter(WriterTypeIntoBuffer)
	n, err = w.Write(b)
	assert.Equal(t, 5, n)
	assert.Equal(t, nil, err)
}

func TestMockWriter_SetSpecificType(t *testing.T) {
	var (
		w   *MockWriter
		n   int
		err error
		b   = []byte("hello")
	)
	w = NewMockWriter(WriterTypeDoNothing)
	w.SetSpecificType(map[uint32]WriterType{
		0: WriterTypeReturnError,
		2: WriterTypeReturnError,
		4: WriterTypeDoNothing,
	})

	expectedLen := map[int]int{
		0: 0,
		1: 5,
		2: 0,
		3: 5,
		4: 5,
		5: 5,
	}
	expectedErr := map[int]error{
		0: mockWriterErr,
		1: nil,
		2: mockWriterErr,
		3: nil,
		4: nil,
		5: nil,
	}

	for i := 0; i < 6; i++ {
		n, err = w.Write(b)
		assert.Equal(t, expectedLen[i], n)
		assert.Equal(t, expectedErr[i], err)
	}
}

func BenchmarkNewMockWriter(b *testing.B) {
	b.ReportAllocs()
	var tmp uint32
	for i := 0; i < b.N; i++ {
		mw := NewMockWriter(WriterTypeDoNothing)
		tmp = tmp + mw.count
	}
}

func newMockWriter2(t WriterType) MockWriter {
	return MockWriter{
		t: t,
	}
}

func BenchmarkNewMockWriter2(b *testing.B) {
	b.ReportAllocs()
	var tmp uint32
	for i := 0; i < b.N; i++ {
		mw := newMockWriter2(WriterTypeDoNothing)
		tmp = tmp + mw.count
	}
}
