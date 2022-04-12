package file

import "testing"

func BenchmarkFileSize1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fileSize1("/home/zl/learn/go-mefs-v2/mefs-user")
	}
}

func BenchmarkFileSize2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fileSize2("/home/zl/learn/go-mefs-v2/mefs-user")
	}
}

func BenchmarkFileSize3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fileSize3("/home/zl/learn/go-mefs-v2/mefs-user")
	}
}

func BenchmarkFileSize4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fileSize4("/home/zl/learn/go-mefs-v2/mefs-user")
	}
}
