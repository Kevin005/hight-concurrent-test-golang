package main

import (
	"testing"
	"time"
)

//测试过程为，打开terminal top和编辑器，让内存保持稳定进行测试
//goroutine测试，用top查看mem内存后得出20万个goroutine占用500MB内存空间，占用512000KB，每个goroutine = 512000/200000 =  2.56KB
//相当于一个goroutine内存开销为2.5KB
func Test_goroutine(t *testing.T) {
	for i := 0; i < 200000; i++ {
		go func() {
			time.Sleep(5 * time.Second)
		}()
	}
	time.Sleep(10 * time.Second)
}
