package beta

import (
	"github.com/redtenant/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkbookEmailIdentity struct {
	// Display name of the user.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// Email address of the user.
	Email nullable.Type[string] `json:"email,omitempty"`

	// The unique identifier of the user.
	Id nullable.Type[string] `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
