package beta

import (
	"github.com/redtenant/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MetaDataKeyStringPair struct {
	// Key of the meta data.
	Key nullable.Type[string] `json:"key,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Value of the meta data.
	Value nullable.Type[string] `json:"value,omitempty"`
}
