package mapper

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

func RegisterCodec(codec *codec.Codec) {
	codec.RegisterConcrete(&baseAsset{}, "assetFactory/baseAsset", nil)
	codec.RegisterConcrete(&baseAssetAddress{}, "assetFactory/baseAssetAddress", nil)
}