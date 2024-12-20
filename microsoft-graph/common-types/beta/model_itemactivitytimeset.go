package beta

import (
	"github.com/redtenant/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ItemActivityTimeSet struct {
	LastRecordedDateTime nullable.Type[string] `json:"lastRecordedDateTime,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// When the activity was observed to take place.
	ObservedDateTime nullable.Type[string] `json:"observedDateTime,omitempty"`

	// When the observation was recorded on the service.
	RecordedDateTime nullable.Type[string] `json:"recordedDateTime,omitempty"`
}
