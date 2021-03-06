/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package helpers

import (
	"github.com/cosmos/cosmos-sdk/client/context"
)

type QueryRequest interface {
	Request
	FromCLI(CLICommand, context.CLIContext) QueryRequest
	FromMap(map[string]string) QueryRequest
}
