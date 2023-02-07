module github.com/apernet/hysteria/core

go 1.19

require (
	github.com/coreos/go-iptables v0.6.0
	github.com/google/gopacket v1.1.19
	github.com/hashicorp/golang-lru/v2 v2.0.1
	github.com/lunixbochs/struc v0.0.0-20200707160740-784aaebc1d40
	github.com/oschwald/geoip2-golang v1.8.0
	github.com/quic-go/quic-go v0.32.0
	github.com/txthinking/socks5 v0.0.0-20220615051428-39268faee3e6
	github.com/valyala/bytebufferpool v0.0.0-20201104193830-18533face0df
	github.com/zeebo/blake3 v0.2.3
	golang.org/x/sys v0.4.0
)

require (
	github.com/go-task/slim-sprig v0.0.0-20210107165309-348f09dbbbc0 // indirect
	github.com/golang/mock v1.6.0 // indirect
	github.com/google/go-cmp v0.5.9 // indirect
	github.com/google/pprof v0.0.0-20210720184732-4bb14d4b1be1 // indirect
	github.com/klauspost/cpuid/v2 v2.2.3 // indirect
	github.com/onsi/ginkgo/v2 v2.2.0 // indirect
	github.com/oschwald/maxminddb-golang v1.10.0 // indirect
	github.com/patrickmn/go-cache v2.1.0+incompatible // indirect
	github.com/quic-go/qtls-go1-18 v0.2.0 // indirect
	github.com/quic-go/qtls-go1-19 v0.2.0 // indirect
	github.com/quic-go/qtls-go1-20 v0.1.0 // indirect
	github.com/stretchr/testify v1.8.1 // indirect
	github.com/txthinking/runnergroup v0.0.0-20220212043759-8da8edb7dae8 // indirect
	github.com/txthinking/x v0.0.0-20220929041811-1b4d914e9133 // indirect
	golang.org/x/crypto v0.4.0 // indirect
	golang.org/x/exp v0.0.0-20221217163422-3c43f8badb15 // indirect
	golang.org/x/mod v0.7.0 // indirect
	golang.org/x/net v0.5.0 // indirect
	golang.org/x/tools v0.4.0 // indirect
	google.golang.org/protobuf v1.28.1 // indirect
)

replace github.com/quic-go/quic-go => github.com/apernet/quic-go v0.31.2-0.20230202062024-7418480ea9b5
