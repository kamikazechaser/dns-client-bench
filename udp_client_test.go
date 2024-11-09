package client

import (
	"context"
	"net"
	"testing"

	"github.com/miekg/dns"
	"github.com/phuslu/fastdns"
)

const upstreamDNSServer = "1.1.1.1:53"

func BenchmarkSTDLibDNS(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		client := net.Resolver{
			PreferGo: true,
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
			Dialer: &fastdns.UDPDialer{
				Addr: func() (u *net.UDPAddr) { u, _ = net.ResolveUDPAddr("udp", upstreamDNSServer); return }(),
			},
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
			Net: "",
		}

		msg := &dns.Msg{}
		msg.SetQuestion("store.steampowered.com.", dns.TypeA)

		for pb.Next() {
			client.ExchangeContext(context.Background(), msg, upstreamDNSServer)
		}
	})
}
