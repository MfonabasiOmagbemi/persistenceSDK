/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package mappables

import (
	"github.com/persistenceOne/persistenceSDK/schema/traits"
	"github.com/persistenceOne/persistenceSDK/schema/types"
)

type InterNFT interface {
	types.NFT
	traits.InterChain
	traits.HasImmutables
	traits.HasMutables
	traits.Burnable
	traits.Lockable
	traits.Mappable
}
