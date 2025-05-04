// Copyright (C) 2025 T-Force I/O
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package multiplex

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestServiceMessageInitialization(t *testing.T) {
	msg := ServiceMessage{
		Command: "testCommand",
		Params:  ExecParams{"key": "value"},
		Extra:   "extraData",
	}

	assert.Equal(t, "testCommand", msg.Command, "Command should match")
	assert.Equal(t, "value", msg.Params["key"], "Params['key'] should match")
	assert.Equal(t, "extraData", msg.Extra, "Extra should match")
}

func TestExecParams_Get(t *testing.T) {
	params := ExecParams{"key": "value"}
	assert.Equal(t, "value", params.Get("key", "default"), "Get should return the correct value")
	assert.Equal(t, "default", params.Get("missing", "default"), "Get should return the default value for missing keys")
}

func TestExecParams_Set(t *testing.T) {
	params := ExecParams{}
	params.Set("key", "value")
	assert.Equal(t, "value", params["key"], "Set should correctly set the key-value pair")
}

func TestExecParams_Delete(t *testing.T) {
	params := ExecParams{"key": "value"}
	params.Delete("key")
	_, exists := params["key"]
	assert.False(t, exists, "Delete should remove the key")
}

func TestExecParams_ExpectReturn(t *testing.T) {
	params := ExecParams{}
	params.ExpectReturn()
	_, exists := params["return"]
	assert.True(t, exists, "ExpectReturn should set the 'return' key")
}

func TestExecParams_Return(t *testing.T) {
	params := ExecParams{}
	params.ExpectReturn()
	go func() {
		params.Return("result")
	}()
	assert.Equal(t, "result", params.WaitForReturn(), "Return should set the result and signal completion")
}

func TestExecParams_WaitForReturn_NoReturn(t *testing.T) {
	params := ExecParams{}
	assert.Nil(t, params.WaitForReturn(), "WaitForReturn should return nil if no return is expected")
}

func TestReturnParams_Signal(t *testing.T) {
	signal := &sync.WaitGroup{}
	retParams := ReturnParams{signal: signal}
	assert.Equal(t, signal, retParams.Signal(), "Signal should return the correct signal")
}

func TestReturnParams_Result(t *testing.T) {
	retParams := ReturnParams{result: "testResult"}
	assert.Equal(t, "testResult", retParams.Result(), "Result should return the correct result")
}
