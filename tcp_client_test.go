package client

import (
	"context"
	"net"
	"testing"

	"github.com/miekg/dns"
	"github.com/phuslu/fastdns"
)

func BenchmarkSTDLibDNS(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		d := &net.Dialer{}
		client := net.Resolver{
			PreferGo: true,
			Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
				return d.DialContext(ctx, "tcp", upstreamDNSServer)
			},
		}
		for pb.Next() {
			client.LookupNetIP(context.Background(), "ip4", "store.steampowered.com")
		}
	})
}

func BenchmarkPhusluFastDNS(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		client := &fastdns.Client{
			Addr: upstreamDNSServer,
		}
		for pb.Next() {
			client.LookupNetIP(context.Background(), "ip4", "store.steampowered.com")
		}
	})
}

func BenchmarkMiekgDNS(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		client := &dns.Client{
			Net: "tcp",
		}

		msg := &dns.Msg{}
		msg.SetQuestion("store.steampowered.com.", dns.TypeA)

		for pb.Next() {
			client.ExchangeContext(context.Background(), msg, upstreamDNSServer)
		}
	})
}
