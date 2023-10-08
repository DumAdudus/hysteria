module github.com/apernet/hysteria/core

go 1.20

require (
	github.com/coreos/go-iptables v0.7.0
	github.com/google/gopacket v1.1.19
	github.com/hashicorp/golang-lru/v2 v2.0.2
	github.com/lunixbochs/struc v0.0.0-20200707160740-784aaebc1d40
	github.com/oschwald/geoip2-golang v1.9.0
	github.com/quic-go/quic-go v0.34.0
	github.com/txthinking/socks5 v0.0.0-20230325130024-4230056ae301
	github.com/valyala/bytebufferpool v0.0.0-20201104193830-18533face0df
	github.com/zeebo/blake3 v0.2.3
	golang.org/x/sys v0.13.0
)

require (
	github.com/go-task/slim-sprig v0.0.0-20230315185526-52ccab3ef572 // indirect
	github.com/golang/mock v1.6.0 // indirect
	github.com/google/pprof v0.0.0-20230502171905-255e3b9b56de // indirect
	github.com/klauspost/cpuid/v2 v2.2.4 // indirect
	github.com/onsi/ginkgo/v2 v2.9.4 // indirect
	github.com/oschwald/maxminddb-golang v1.11.0 // indirect
	github.com/patrickmn/go-cache v2.1.0+incompatible // indirect
	github.com/quic-go/qtls-go1-19 v0.3.2 // indirect
	github.com/quic-go/qtls-go1-20 v0.2.3 // indirect
	github.com/txthinking/runnergroup v0.0.0-20230325130830-408dc5853f86 // indirect
	golang.org/x/crypto v0.14.0 // indirect
	golang.org/x/exp v0.0.0-20231006140011-7918f672742d // indirect
	golang.org/x/mod v0.13.0 // indirect
	golang.org/x/net v0.16.0 // indirect
	golang.org/x/tools v0.14.0 // indirect
	google.golang.org/protobuf v1.31.0 // indirect
)

replace github.com/quic-go/quic-go => github.com/apernet/quic-go v0.34.1-0.20230507231629-ec008b7e8473
