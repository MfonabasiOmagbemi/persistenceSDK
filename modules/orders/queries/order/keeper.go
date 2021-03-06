/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package order

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/modules/orders/mapper"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
)

type queryKeeper struct {
	mapper helpers.Mapper
}

var _ helpers.QueryKeeper = (*queryKeeper)(nil)

func (queryKeeper queryKeeper) Enquire(context sdkTypes.Context, queryRequest helpers.QueryRequest) helpers.QueryResponse {
	return newQueryResponse(mapper.NewOrders(queryKeeper.mapper, context).Fetch(queryRequestFromInterface(queryRequest).OrderID), nil)
}

func initializeQueryKeeper(mapper helpers.Mapper, _ []interface{}) helpers.QueryKeeper {
	return queryKeeper{mapper: mapper}
}
