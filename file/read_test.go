package file

import (
	"testing"
)

func TestRead0(t *testing.T) {
	if _, err := Read0("/home/zl/learn/go-mefs-v2/mefs-user"); err != nil {
		t.Error("read0(\"/home/zl/learn/go-mefs-v2/mefs-user\") error")
	}
}

func BenchmarkRead0(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Read0("/home/zl/learn/go-mefs-v2/mefs-user")
	}
}
