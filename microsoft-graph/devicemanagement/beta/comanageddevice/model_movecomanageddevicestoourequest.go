package comanageddevice

import (
	"github.com/redtenant/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MoveComanagedDevicesToOURequest struct {
	DeviceIds              *[]string             `json:"deviceIds,omitempty"`
	OrganizationalUnitPath nullable.Type[string] `json:"organizationalUnitPath,omitempty"`
}
