package beta

import (
	"encoding/json"
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ BaseCollectionPaginationCountResponse = AospDeviceOwnerTrustedRootCertificateCollectionResponse{}

type AospDeviceOwnerTrustedRootCertificateCollectionResponse struct {
	Value *[]AospDeviceOwnerTrustedRootCertificate `json:"value,omitempty"`

	// Fields inherited from BaseCollectionPaginationCountResponse

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	ODataNextLink nullable.Type[string] `json:"@odata.nextLink,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s AospDeviceOwnerTrustedRootCertificateCollectionResponse) BaseCollectionPaginationCountResponse() BaseBaseCollectionPaginationCountResponseImpl {
	return BaseBaseCollectionPaginationCountResponseImpl{
		ODataId:       s.ODataId,
		ODataNextLink: s.ODataNextLink,
		ODataType:     s.ODataType,
	}
}

var _ json.Marshaler = AospDeviceOwnerTrustedRootCertificateCollectionResponse{}

func (s AospDeviceOwnerTrustedRootCertificateCollectionResponse) MarshalJSON() ([]byte, error) {
	type wrapper AospDeviceOwnerTrustedRootCertificateCollectionResponse
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AospDeviceOwnerTrustedRootCertificateCollectionResponse: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AospDeviceOwnerTrustedRootCertificateCollectionResponse: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.aospDeviceOwnerTrustedRootCertificateCollectionResponse"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AospDeviceOwnerTrustedRootCertificateCollectionResponse: %+v", err)
	}

	return encoded, nil
}
