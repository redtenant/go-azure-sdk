package beta

import (
	"github.com/redtenant/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TermStoreLocalizedDescription struct {
	// The description in the localized language.
	Description nullable.Type[string] `json:"description,omitempty"`

	// The language tag for the label.
	LanguageTag nullable.Type[string] `json:"languageTag,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
