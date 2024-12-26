package p2p

import "testing"

func BenchmarkListenerAcceptAndHandle(b *testing.B) {
	b.StartTimer()
	b.StopTimer()
}
