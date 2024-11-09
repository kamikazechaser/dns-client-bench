package client

import (
	"context"
	"net"
	"testing"
)

func BenchmarkSTDLibDNHostResolver(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		client := net.Resolver{
			PreferGo: false,
		}
		for pb.Next() {
			client.LookupNetIP(context.Background(), "ip4", "store.steampowered.com")
		}
	})
}
