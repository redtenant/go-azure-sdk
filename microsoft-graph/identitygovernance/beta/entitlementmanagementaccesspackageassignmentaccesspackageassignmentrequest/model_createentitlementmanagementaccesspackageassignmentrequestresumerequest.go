package entitlementmanagementaccesspackageassignmentaccesspackageassignmentrequest

import (
	"github.com/redtenant/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/redtenant/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateEntitlementManagementAccessPackageAssignmentRequestResumeRequest struct {
	Data   *beta.CustomExtensionData `json:"data,omitempty"`
	Source nullable.Type[string]     `json:"source,omitempty"`
	Type   nullable.Type[string]     `json:"type,omitempty"`
}
