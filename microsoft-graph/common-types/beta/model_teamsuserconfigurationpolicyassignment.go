package beta

import (
	"github.com/redtenant/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamsUserConfigurationPolicyAssignment struct {
	AssignmentType *TeamsUserConfigurationAssignmentType `json:"assignmentType,omitempty"`
	DisplayName    *string                               `json:"displayName,omitempty"`
	GroupId        nullable.Type[string]                 `json:"groupId,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	PolicyId *string `json:"policyId,omitempty"`
}
