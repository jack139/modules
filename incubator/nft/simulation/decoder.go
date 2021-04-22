package simulation

import (
	"bytes"
	"fmt"

	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/modules/incubator/nft/types"
)

// DecodeStore unmarshals the KVPair's Value to the corresponding gov type
func DecodeStore(cdc *codec.LegacyAmino, kvA, kvB abci.EventAttribute) string {
	switch {
	case bytes.Equal(kvA.Key[:1], types.CollectionsKeyPrefix):
		var collectionA, collectionB types.Collection
		cdc.MustUnmarshalBinaryLengthPrefixed(kvA.Value, &collectionA)
		cdc.MustUnmarshalBinaryLengthPrefixed(kvB.Value, &collectionB)
		return fmt.Sprintf("%v\n%v", collectionA, collectionB)

	case bytes.Equal(kvA.Key[:1], types.OwnersKeyPrefix):
		var idCollectionA, idCollectionB types.IDCollection
		cdc.MustUnmarshalBinaryLengthPrefixed(kvA.Value, &idCollectionA)
		cdc.MustUnmarshalBinaryLengthPrefixed(kvB.Value, &idCollectionB)
		return fmt.Sprintf("%v\n%v", idCollectionA, idCollectionB)

	default:
		panic(fmt.Sprintf("invalid %s key prefix %X", types.ModuleName, kvA.Key[:1]))
	}
}
