package tests

import "testing"

func BenchmarkTest1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Test1()
	}

}

func BenchmarkTest2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Test2()
	}
}
