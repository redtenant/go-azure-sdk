package beta

import (
	"github.com/redtenant/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCWindowsSetting struct {
	// The Windows language or region tag to use for language pack configuration and localization of the Cloud PC. The
	// default value is en-US, which corresponds to English (United States).
	Locale nullable.Type[string] `json:"locale,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
