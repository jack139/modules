module github.com/cosmos/modules/incubator/nft

go 1.13

require (
	github.com/cosmos/cosmos-sdk v0.42.4
	github.com/gogo/protobuf v1.3.3
	github.com/gorilla/mux v1.8.0
	github.com/spf13/cobra v1.1.1
	github.com/spf13/viper v1.7.1
	github.com/stretchr/testify v1.7.0
	github.com/tendermint/tendermint v0.34.9
	github.com/tendermint/tm-db v0.6.4
)

replace github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1
