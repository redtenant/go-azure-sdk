package me

import (
	"github.com/redtenant/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WipeManagedAppRegistrationByDeviceTagRequest struct {
	DeviceTag nullable.Type[string] `json:"deviceTag,omitempty"`
}
