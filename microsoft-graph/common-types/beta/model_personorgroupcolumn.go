package beta

import (
	"github.com/redtenant/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PersonOrGroupColumn struct {
	// Indicates whether multiple values can be selected from the source.
	AllowMultipleSelection nullable.Type[bool] `json:"allowMultipleSelection,omitempty"`

	// Whether to allow selection of people only, or people and groups. Must be one of peopleAndGroups or peopleOnly.
	ChooseFromType nullable.Type[string] `json:"chooseFromType,omitempty"`

	// How to display the information about the person or group chosen. See below.
	DisplayAs nullable.Type[string] `json:"displayAs,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
