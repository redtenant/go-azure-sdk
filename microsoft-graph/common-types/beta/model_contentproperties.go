package beta

import (
	"github.com/redtenant/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ContentProperties struct {
	Extensions           *[]string             `json:"extensions,omitempty"`
	LastModifiedBy       nullable.Type[string] `json:"lastModifiedBy,omitempty"`
	LastModifiedDateTime nullable.Type[string] `json:"lastModifiedDateTime,omitempty"`
	Metadata             *ContentMetadata      `json:"metadata,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
