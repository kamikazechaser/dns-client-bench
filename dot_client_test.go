package client

import (
	"context"
	"crypto/tls"
	"net"
	"testing"

	"github.com/miekg/dns"
	"github.com/phuslu/fastdns"
)

const upstreamDOTServer = "1.1.1.1:853"

func BenchmarkDOTSTDLibDNS(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		var dialer net.Dialer
		client := net.Resolver{
			PreferGo: true,
			Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
				conn, err := dialer.DialContext(ctx, "tcp", upstreamDOTServer)
				if err != nil {
					return nil, err
				}
				return tls.Client(conn, &tls.Config{
					InsecureSkipVerify: true,
					ServerName:         upstreamDOTServer,
					ClientSessionCache: tls.NewLRUClientSessionCache(1024),
				}), nil
			},
		}
		for pb.Next() {
			client.LookupNetIP(context.Background(), "ip4", "store.steampowered.com")
		}
	})
}

func BenchmarkDOTPhusluFastDNS(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		client := &fastdns.Client{
			Addr: upstreamDOTServer,
			Dialer: &fastdns.TLSDialer{
				Addr: func() (u *net.TCPAddr) { u, _ = net.ResolveTCPAddr("tcp", upstreamDOTServer); return }(),
				TLSConfig: &tls.Config{
					InsecureSkipVerify: true,
					ServerName:         upstreamDOTServer,
					ClientSessionCache: tls.NewLRUClientSessionCache(1024),
				},
			},
		}
		for pb.Next() {
			client.LookupNetIP(context.Background(), "ip4", "store.steampowered.com")
		}
	})
}

func BenchmarkDOTMiekgDNS(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		client := &dns.Client{
			Net: "tcp-tls",
			TLSConfig: &tls.Config{
				InsecureSkipVerify: true,
				ServerName:         upstreamDOTServer,
				ClientSessionCache: tls.NewLRUClientSessionCache(1024),
			},
		}

		msg := &dns.Msg{}
		msg.SetQuestion("store.steampowered.com.", dns.TypeA)

		for pb.Next() {
			client.ExchangeContext(context.Background(), msg, upstreamDOTServer)
		}
	})
}
