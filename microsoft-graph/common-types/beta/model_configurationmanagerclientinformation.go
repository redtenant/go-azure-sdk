package beta

import (
	"github.com/redtenant/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConfigurationManagerClientInformation struct {
	// Configuration Manager Client Id from SCCM
	ClientIdentifier nullable.Type[string] `json:"clientIdentifier,omitempty"`

	// Configuration Manager Client version from SCCM
	ClientVersion nullable.Type[string] `json:"clientVersion,omitempty"`

	// Configuration Manager Client blocked status from SCCM
	IsBlocked *bool `json:"isBlocked,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
