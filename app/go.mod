module github.com/apernet/hysteria/app

go 1.20

require (
	github.com/LiamHaworth/go-tproxy v0.0.0-20190726054950-ef7efd7f24ed
	github.com/antonfisher/nested-logrus-formatter v1.3.1
	github.com/apernet/hysteria/core v1.3.3
	github.com/caddyserver/certmagic v0.17.2
	github.com/docker/go-units v0.5.0
	github.com/elazarl/goproxy v0.0.0-20221015165544-a0805db90819
	github.com/elazarl/goproxy/ext v0.0.0-20221015165544-a0805db90819
	github.com/folbricht/routedns v0.1.21-0.20230220022436-4ae86ce30d53
	github.com/fsnotify/fsnotify v1.7.0
	github.com/oschwald/geoip2-golang v1.9.0
	github.com/prometheus/client_golang v1.14.0
	github.com/quic-go/quic-go v0.38.1
	github.com/sirupsen/logrus v1.9.3
	github.com/spf13/cobra v1.8.0
	github.com/spf13/viper v1.18.2
	github.com/txthinking/socks5 v0.0.0-20230325130024-4230056ae301
	github.com/xjasonlyu/tun2socks/v2 v2.5.2
	github.com/yosuke-furukawa/json5 v0.1.1
	go.uber.org/zap v1.26.0
	gvisor.dev/gvisor v0.0.0-20230603040744-5c9219dedd33
)

require (
	github.com/RackSec/srslog v0.0.0-20180709174129-a4725f04ec91 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/coreos/go-iptables v0.7.0 // indirect
	github.com/go-task/slim-sprig v0.0.0-20230315185526-52ccab3ef572 // indirect
	github.com/golang/mock v1.6.0 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/google/btree v1.1.2 // indirect
	github.com/google/gopacket v1.1.19 // indirect
	github.com/google/pprof v0.0.0-20230502171905-255e3b9b56de // indirect
	github.com/hashicorp/golang-lru/v2 v2.0.2 // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/jtacoma/uritemplates v1.0.0 // indirect
	github.com/klauspost/cpuid/v2 v2.2.4 // indirect
	github.com/libdns/libdns v0.2.1 // indirect
	github.com/lunixbochs/struc v0.0.0-20200707160740-784aaebc1d40 // indirect
	github.com/magiconair/properties v1.8.7 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.4 // indirect
	github.com/mholt/acmez v1.1.0 // indirect
	github.com/miekg/dns v1.1.57 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/onsi/ginkgo/v2 v2.9.5 // indirect
	github.com/oschwald/maxminddb-golang v1.12.0 // indirect
	github.com/patrickmn/go-cache v2.1.0+incompatible // indirect
	github.com/pelletier/go-toml/v2 v2.1.0 // indirect
	github.com/pion/dtls/v2 v2.2.6 // indirect
	github.com/pion/logging v0.2.2 // indirect
	github.com/pion/transport/v2 v2.0.2 // indirect
	github.com/pion/udp/v2 v2.0.1 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/prometheus/client_model v0.3.0 // indirect
	github.com/prometheus/common v0.42.0 // indirect
	github.com/prometheus/procfs v0.9.0 // indirect
	github.com/quic-go/qpack v0.4.0 // indirect
	github.com/quic-go/qtls-go1-19 v0.3.2 // indirect
	github.com/quic-go/qtls-go1-20 v0.2.3 // indirect
	github.com/sagikazarmark/locafero v0.4.0 // indirect
	github.com/sagikazarmark/slog-shim v0.1.0 // indirect
	github.com/sourcegraph/conc v0.3.0 // indirect
	github.com/spf13/afero v1.11.0 // indirect
	github.com/spf13/cast v1.6.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/subosito/gotenv v1.6.0 // indirect
	github.com/txthinking/runnergroup v0.0.0-20230325130830-408dc5853f86 // indirect
	github.com/valyala/bytebufferpool v0.0.0-20201104193830-18533face0df // indirect
	github.com/zeebo/blake3 v0.2.3 // indirect
	go.uber.org/multierr v1.10.0 // indirect
	golang.org/x/crypto v0.17.0 // indirect
	golang.org/x/exp v0.0.0-20231226003508-02704c960a9b // indirect
	golang.org/x/mod v0.14.0 // indirect
	golang.org/x/net v0.19.0 // indirect
	golang.org/x/sys v0.15.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	golang.org/x/time v0.5.0 // indirect
	golang.org/x/tools v0.16.1 // indirect
	golang.zx2c4.com/wintun v0.0.0-20230126152724-0fa3db229ce2 // indirect
	golang.zx2c4.com/wireguard v0.0.0-20230325221338-052af4a8072b // indirect
	google.golang.org/protobuf v1.32.0 // indirect
	gopkg.in/ini.v1 v1.67.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/apernet/hysteria/core => ../core/

replace github.com/quic-go/quic-go => github.com/apernet/quic-go v0.34.1-0.20230507231629-ec008b7e8473

replace github.com/LiamHaworth/go-tproxy => github.com/apernet/go-tproxy v0.0.0-20221025153553-ed04a2935f88

replace github.com/elazarl/goproxy => github.com/apernet/goproxy v0.0.0-20221124043924-155acfaf278f
