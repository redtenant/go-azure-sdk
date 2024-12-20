package stable

import (
	"github.com/redtenant/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SynchronizationError struct {
	// The error code. For example, AzureDirectoryB2BManagementPolicyCheckFailure.
	Code nullable.Type[string] `json:"code,omitempty"`

	// The error message. For example, Policy permitting auto-redemption of invitations not configured.
	Message nullable.Type[string] `json:"message,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The action to take to resolve the error. For example, false.
	TenantActionable *bool `json:"tenantActionable,omitempty"`
}
