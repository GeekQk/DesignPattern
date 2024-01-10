package main

import (
	"sync"
	"testing"
)

func TestGetInstance5(t *testing.T) {
	// 正常流程测试
	instance1 := GetInstance5()
	instance2 := GetInstance5()

	if instance1 != instance2 {
		t.Errorf("the Expect instance is %v, but the actual instance is %v", instance1, instance2)
	}

	// 异常流程测试
	var wg sync.WaitGroup
	count := 50
	wg.Add(count)

	for i := 0; i < count; i++ {
		go func() {
			defer wg.Done()
			GetInstance5()
		}()
	}
	wg.Wait()
}
