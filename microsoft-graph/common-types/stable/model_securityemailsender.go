package stable

import (
	"github.com/redtenant/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityEmailSender struct {
	// The name of the sender.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// Sender domain.
	DomainName nullable.Type[string] `json:"domainName,omitempty"`

	// Sender email address.
	EmailAddress nullable.Type[string] `json:"emailAddress,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
