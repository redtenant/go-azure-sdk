package stable

import (
	"github.com/redtenant/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SynchronizationMetadataEntry struct {
	// Possible values are: GalleryApplicationIdentifier, GalleryApplicationKey, IsOAuthEnabled,
	// IsSynchronizationAgentAssignmentRequired, IsSynchronizationAgentRequired, IsSynchronizationInPreview, OAuthSettings,
	// SynchronizationLearnMoreIbizaFwLink, ConfigurationFields.
	Key *SynchronizationMetadata `json:"key,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Value of the metadata property.
	Value nullable.Type[string] `json:"value,omitempty"`
}
