package test

import "testing"

/**
 * @Author: Alan Yin
 * @Date: 2022/5/20
 * @Desc:
 */

func BenchmarkRandInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RandInt()
	}
}

//
//func BenchmarkBigLen(b *testing.B) {
//	big := NewBig() // 耗时的准备工作
//	b.ResetTimer()  // 重置性能测试时间计数
//	for i := 0; i < b.N; i++ {
//		big.Len()
//	}
//}
//
//func BenchmarkBigLen2(b *testing.B) {
//	b.StopTimer()
//	big := NewBig() // 耗时的准备工作
//	b.StartTimer()  // 重新开始时间
//	for i := 0; i < b.N; i++ {
//		big.Len()
//	}
//}
