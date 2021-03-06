/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package swap

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/modules/exchanges/mapper"
	"github.com/persistenceOne/persistenceSDK/modules/splits/auxiliaries/burn"
	"github.com/persistenceOne/persistenceSDK/modules/splits/auxiliaries/mint"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
)

type auxiliaryKeeper struct {
	mapper              helpers.Mapper
	splitsMintAuxiliary helpers.Auxiliary
	splitsBurnAuxiliary helpers.Auxiliary
}

var _ helpers.AuxiliaryKeeper = (*auxiliaryKeeper)(nil)

func (auxiliaryKeeper auxiliaryKeeper) Help(context sdkTypes.Context, AuxiliaryRequest helpers.AuxiliaryRequest) helpers.AuxiliaryResponse {
	auxiliaryRequest := auxiliaryRequestFromInterface(AuxiliaryRequest)

	if auxiliaryResponse := auxiliaryKeeper.splitsBurnAuxiliary.GetKeeper().Help(context, burn.NewAuxiliaryRequest(base.NewID(mapper.ModuleName), auxiliaryRequest.MakerSplitID, auxiliaryRequest.MakerSplit)); !auxiliaryResponse.IsSuccessful() {
		return newAuxiliaryResponse(auxiliaryResponse.GetError())
	}
	if auxiliaryResponse := auxiliaryKeeper.splitsBurnAuxiliary.GetKeeper().Help(context, burn.NewAuxiliaryRequest(auxiliaryRequest.TakerID, auxiliaryRequest.TakerSplitID, auxiliaryRequest.TakerSplit)); !auxiliaryResponse.IsSuccessful() {
		return newAuxiliaryResponse(auxiliaryResponse.GetError())
	}
	if auxiliaryResponse := auxiliaryKeeper.splitsMintAuxiliary.GetKeeper().Help(context, mint.NewAuxiliaryRequest(auxiliaryRequest.MakerID, auxiliaryRequest.TakerSplitID, auxiliaryRequest.TakerSplit)); !auxiliaryResponse.IsSuccessful() {
		return newAuxiliaryResponse(auxiliaryResponse.GetError())
	}
	if auxiliaryResponse := auxiliaryKeeper.splitsMintAuxiliary.GetKeeper().Help(context, mint.NewAuxiliaryRequest(auxiliaryRequest.TakerID, auxiliaryRequest.MakerSplitID, auxiliaryRequest.MakerSplit)); !auxiliaryResponse.IsSuccessful() {
		return newAuxiliaryResponse(auxiliaryResponse.GetError())
	}
	return newAuxiliaryResponse(nil)
}

func initializeAuxiliaryKeeper(mapper helpers.Mapper, externalKeepers []interface{}) helpers.AuxiliaryKeeper {
	auxiliaryKeeper := auxiliaryKeeper{mapper: mapper}
	for _, externalKeeper := range externalKeepers {
		switch value := externalKeeper.(type) {
		case helpers.Auxiliary:
			switch value.GetName() {
			case mint.Auxiliary.GetName():
				auxiliaryKeeper.splitsMintAuxiliary = value
			case burn.Auxiliary.GetName():
				auxiliaryKeeper.splitsBurnAuxiliary = value
			}
		}
	}
	return auxiliaryKeeper
}
