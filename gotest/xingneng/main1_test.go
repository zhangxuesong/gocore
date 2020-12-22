package xingneng

import (
	"testing"
	"time"
)

//func BenchmarkGetPrimes(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		GetPrimes(1000)
//	}
//}

func BenchmarkGetPrimes(b *testing.B) {
	b.StopTimer()
	time.Sleep(time.Millisecond * 500) // 模拟某个耗时但与被测程序关系不大的操作。
	max := 10000
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		GetPrimes(max)
	}
}
