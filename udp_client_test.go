package client

import (
	"context"
	"net"
	"testing"

	"github.com/miekg/dns"
	"github.com/phuslu/fastdns"
)

const upstreamDNSServer = "1.1.1.1:53"

func BenchmarkUDPSTDLibDNS(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		d := &net.Dialer{}
		client := net.Resolver{
			PreferGo: true,
			Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
				return d.DialContext(ctx, "udp", upstreamDNSServer)
			},
		}
		for pb.Next() {
			client.LookupNetIP(context.Background(), "ip4", "store.steampowered.com")
		}
	})
}

func BenchmarkUDPPhusluFastDNS(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		client := &fastdns.Client{
			Addr: upstreamDNSServer,
			Dialer: &fastdns.UDPDialer{
				Addr:     func() (u *net.UDPAddr) { u, _ = net.ResolveUDPAddr("udp", upstreamDNSServer); return }(),
				MaxConns: 1024,
			},
		}
		for pb.Next() {
			client.LookupNetIP(context.Background(), "ip4", "store.steampowered.com")
		}
	})
}

func BenchmarkUDPMiekgDNS(b *testing.B) {
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
