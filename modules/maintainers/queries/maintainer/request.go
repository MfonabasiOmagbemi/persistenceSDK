/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package maintainer

import (
	"github.com/asaskevich/govalidator"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/persistenceOne/persistenceSDK/constants/flags"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
)

type queryRequest struct {
	MaintainerID types.ID `json:"maintainerID" valid:"required~required field maintainerID missing"`
}

var _ helpers.QueryRequest = (*queryRequest)(nil)

func (queryRequest queryRequest) Validate() error {
	_, Error := govalidator.ValidateStruct(queryRequest)
	return Error
}

func (queryRequest queryRequest) FromCLI(cliCommand helpers.CLICommand, _ context.CLIContext) helpers.QueryRequest {
	return newQueryRequest(base.NewID(cliCommand.ReadString(flags.MaintainerID)))
}

func (queryRequest queryRequest) FromMap(vars map[string]string) helpers.QueryRequest {
	return newQueryRequest(base.NewID(vars[flags.MaintainerID.GetName()]))
}

func queryRequestPrototype() helpers.QueryRequest {
	return queryRequest{}
}

func queryRequestFromInterface(QueryRequest helpers.QueryRequest) queryRequest {
	switch value := QueryRequest.(type) {
	case queryRequest:
		return value
	default:
		return queryRequest{}
	}
}

func newQueryRequest(maintainerID types.ID) helpers.QueryRequest {
	return queryRequest{MaintainerID: maintainerID}
}
