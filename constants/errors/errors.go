/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package errors

import (
	"github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/persistenceOne/persistenceSDK/constants"
)

var DeletionNotAllowed = errors.Register(constants.ProjectRoute, 101, "DeletionNotAllowed")
var EntityAlreadyExists = errors.Register(constants.ProjectRoute, 102, "EntityAlreadyExists")
var EntityNotFound = errors.Register(constants.ProjectRoute, 103, "EntityNotFound")
var IncorrectMessage = errors.Register(constants.ProjectRoute, 105, "IncorrectMessage")
var InsufficientBalance = errors.Register(constants.ProjectRoute, 106, "InsufficientBalance")
var MetaDataError = errors.Register(constants.ProjectRoute, 107, "MetaDataError")
var NotAuthorized = errors.Register(constants.ProjectRoute, 108, "NotAuthorized")
