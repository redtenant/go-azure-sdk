package beta

import (
	"github.com/redtenant/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamGuestSettings struct {
	// If set to true, guests can add and update channels.
	AllowCreateUpdateChannels nullable.Type[bool] `json:"allowCreateUpdateChannels,omitempty"`

	// If set to true, guests can delete channels.
	AllowDeleteChannels nullable.Type[bool] `json:"allowDeleteChannels,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
