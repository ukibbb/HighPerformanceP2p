package tcp

import "testing"

func BenchmarkListenerAcceptAndHandle(b *testing.B) {
	b.StartTimer()
	b.StopTimer()
}
