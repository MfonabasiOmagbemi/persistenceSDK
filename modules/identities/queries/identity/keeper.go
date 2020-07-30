package identity

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/modules/identities/mapper"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
)

type queryKeeper struct {
	mapper helpers.Mapper
}

var _ helpers.QueryKeeper = (*queryKeeper)(nil)

func (queryKeeper queryKeeper) Enquire(context sdkTypes.Context, queryRequest helpers.QueryRequest) helpers.QueryResponse {
	return newQueryResponse(mapper.NewIdentities(queryKeeper.mapper, context).Fetch(queryRequestFromInterface(queryRequest).IdentityID))
}

func initializeQueryKeeper(mapper helpers.Mapper, _ []interface{}) helpers.QueryKeeper {
	return queryKeeper{mapper: mapper}
}
