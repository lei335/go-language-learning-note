package file

import "testing"

func BenchmarkFileSize1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fileSize1("/home/zl/go/bin/tendermint")
	}
}

func BenchmarkFileSize2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fileSize2("/home/zl/go/bin/tendermint")
	}
}

func BenchmarkFileSize3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fileSize3("/home/zl/go/bin/tendermint")
	}
}

func BenchmarkFileSize4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fileSize4("/home/zl/go/bin/tendermint")
	}
}
