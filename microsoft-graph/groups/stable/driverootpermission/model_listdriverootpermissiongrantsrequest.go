package driverootpermission

import (
	"github.com/redtenant/go-azure-sdk/microsoft-graph/common-types/stable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListDriveRootPermissionGrantsRequest struct {
	Recipients *[]stable.DriveRecipient `json:"recipients,omitempty"`
	Roles      *[]string                `json:"roles,omitempty"`
}
