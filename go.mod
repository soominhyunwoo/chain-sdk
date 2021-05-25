go 1.15

module github.com/soominhyunwoo/chain-sdk

require (
	github.com/99designs/keyring v1.1.6
	github.com/armon/go-metrics v0.3.8
	github.com/bgentry/speakeasy v0.1.0
	github.com/btcsuite/btcd v0.21.0-beta
	github.com/btcsuite/btcutil v1.0.2
	github.com/coinbase/rosetta-sdk-go v0.6.10
	github.com/confio/ics23/go v0.6.6
	github.com/desertbit/timer v0.0.0-20180107155436-c41aec40b27f // indirect
	github.com/enigmampc/btcutil v1.0.3-0.20200723161021-e2fb6adb2a25
	github.com/gogo/gateway v1.1.0
	github.com/gogo/protobuf v1.3.2
	github.com/golang/mock v1.4.4
	github.com/golang/protobuf v1.5.2
	github.com/gorilla/handlers v1.5.1
	github.com/gorilla/mux v1.8.0
	github.com/grpc-ecosystem/go-grpc-middleware v1.3.0
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/hashicorp/golang-lru v0.5.4
	github.com/hdevalence/ed25519consensus v0.0.0-20210204194344-59a8610d2b87
	github.com/improbable-eng/grpc-web v0.14.0
	github.com/jhump/protoreflect v1.8.2
	github.com/magiconair/properties v1.8.5
	github.com/mattn/go-isatty v0.0.12
	github.com/otiai10/copy v1.6.0
	github.com/pelletier/go-toml v1.8.1 // indirect
	github.com/pkg/errors v0.9.1
	github.com/prometheus/client_golang v1.10.0
	github.com/prometheus/common v0.23.0
	github.com/rakyll/statik v0.1.7
	github.com/rs/zerolog v1.21.0
	github.com/soominhyunwoo/chain-proto v0.0.1
	github.com/soominhyunwoo/go-bip39 v0.0.1
	github.com/soominhyunwoo/iavl v0.0.6
	github.com/soominhyunwoo/ledger-chain-go v0.0.2
	github.com/spf13/afero v1.3.4 // indirect
	github.com/spf13/cast v1.3.1
	github.com/spf13/cobra v1.1.3
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5
	github.com/spf13/viper v1.7.1
	github.com/stretchr/testify v1.7.0
	github.com/soominhyunwoo/btcd v0.0.1
	github.com/soominhyunwoo/crypto v0.0.1
	github.com/soominhyunwoo/chain-amino v0.0.1
	github.com/soominhyunwoo/tendermint v0.0.6
	github.com/soominhyunwoo/tm-db v0.0.2
	golang.org/x/crypto v0.0.0-20201221181555-eec23a3978ad
	golang.org/x/sys v0.0.0-20210510120138-977fb7262007 // indirect
	google.golang.org/genproto v0.0.0-20210114201628-6edceaf6022f
	google.golang.org/grpc v1.37.1
	google.golang.org/protobuf v1.26.0
	gopkg.in/ini.v1 v1.61.0 // indirect
	gopkg.in/yaml.v2 v2.4.0
	nhooyr.io/websocket v1.8.6 // indirect
)

replace google.golang.org/grpc => google.golang.org/grpc v1.33.2
