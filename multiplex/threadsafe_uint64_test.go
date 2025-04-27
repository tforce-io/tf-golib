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
)

func TestUint64ThreadSafe_SetAndValue(t *testing.T) {
	counter := &Uint64ThreadSafe{}

	counter.Set(10)
	if counter.Value() != 10 {
		t.Errorf("Expected value to be 10, got %d", counter.Value())
	}
}

func TestUint64ThreadSafe_Add(t *testing.T) {
	counter := &Uint64ThreadSafe{}

	counter.Set(10)
	counter.Add(5)
	if counter.Value() != 15 {
		t.Errorf("Expected value to be 15, got %d", counter.Value())
	}
}

func TestUint64ThreadSafe_Sub(t *testing.T) {
	counter := &Uint64ThreadSafe{}

	counter.Set(15)
	counter.Sub(3)
	if counter.Value() != 12 {
		t.Errorf("Expected value to be 12, got %d", counter.Value())
	}
}

func TestUint64ThreadSafe_Mul(t *testing.T) {
	counter := &Uint64ThreadSafe{}

	counter.Set(12)
	counter.Mul(2)
	if counter.Value() != 24 {
		t.Errorf("Expected value to be 24, got %d", counter.Value())
	}
}

func TestUint64ThreadSafe_Div(t *testing.T) {
	counter := &Uint64ThreadSafe{}

	counter.Set(24)
	counter.Div(4)
	if counter.Value() != 6 {
		t.Errorf("Expected value to be 6, got %d", counter.Value())
	}
}

func TestUint64ThreadSafe_ConcurrentAdd(t *testing.T) {
	var wg sync.WaitGroup
	counter := &Uint64ThreadSafe{}

	counter.Set(6)
	wg.Add(2)
	go func() {
		defer wg.Done()
		counter.Add(10)
	}()
	go func() {
		defer wg.Done()
		counter.Add(20)
	}()
	wg.Wait()
	if counter.Value() != 36 {
		t.Errorf("Expected value to be 36 after concurrent Add, got %d", counter.Value())
	}
}

func TestUint64ThreadSafe_LockAndUnlock(t *testing.T) {
	counter := &Uint64ThreadSafe{}

	counter.Lock()
	go func() {
		defer counter.Unlock()
		counter.SetNoLock(50)
	}()
	counter.Lock()
	if counter.ValueNoLock() != 50 {
		t.Errorf("Expected value to be 50 after manual lock/unlock, got %d", counter.Value())
	}
	counter.Unlock()
}
