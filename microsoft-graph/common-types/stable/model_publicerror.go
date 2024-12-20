package stable

import (
	"github.com/redtenant/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PublicError struct {
	// Represents the error code.
	Code nullable.Type[string] `json:"code,omitempty"`

	// Details of the error.
	Details *[]PublicErrorDetail `json:"details,omitempty"`

	// Details of the inner error.
	InnerError *PublicInnerError `json:"innerError,omitempty"`

	// A non-localized message for the developer.
	Message nullable.Type[string] `json:"message,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The target of the error.
	Target nullable.Type[string] `json:"target,omitempty"`
}
