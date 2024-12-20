package beta

import (
	"github.com/redtenant/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OnPremisesWritebackConfiguration struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The distinguished name of the on-premises container that the customer is using to store unified groups which are
	// created in the cloud.
	UnifiedGroupContainer nullable.Type[string] `json:"unifiedGroupContainer,omitempty"`

	// The distinguished name of the on-premises container that the customer is using to store users which are created in
	// the cloud.
	UserContainer nullable.Type[string] `json:"userContainer,omitempty"`
}
