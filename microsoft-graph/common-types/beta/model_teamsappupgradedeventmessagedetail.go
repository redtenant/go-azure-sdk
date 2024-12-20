package beta

import (
	"encoding/json"
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ EventMessageDetail = TeamsAppUpgradedEventMessageDetail{}

type TeamsAppUpgradedEventMessageDetail struct {
	// Initiator of the event.
	Initiator IdentitySet `json:"initiator"`

	// Display name of the teamsApp.
	TeamsAppDisplayName nullable.Type[string] `json:"teamsAppDisplayName,omitempty"`

	// Unique identifier of the teamsApp.
	TeamsAppId nullable.Type[string] `json:"teamsAppId,omitempty"`

	// Fields inherited from EventMessageDetail

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s TeamsAppUpgradedEventMessageDetail) EventMessageDetail() BaseEventMessageDetailImpl {
	return BaseEventMessageDetailImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = TeamsAppUpgradedEventMessageDetail{}

func (s TeamsAppUpgradedEventMessageDetail) MarshalJSON() ([]byte, error) {
	type wrapper TeamsAppUpgradedEventMessageDetail
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling TeamsAppUpgradedEventMessageDetail: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling TeamsAppUpgradedEventMessageDetail: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.teamsAppUpgradedEventMessageDetail"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling TeamsAppUpgradedEventMessageDetail: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &TeamsAppUpgradedEventMessageDetail{}

func (s *TeamsAppUpgradedEventMessageDetail) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		TeamsAppDisplayName nullable.Type[string] `json:"teamsAppDisplayName,omitempty"`
		TeamsAppId          nullable.Type[string] `json:"teamsAppId,omitempty"`
		ODataId             *string               `json:"@odata.id,omitempty"`
		ODataType           *string               `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.TeamsAppDisplayName = decoded.TeamsAppDisplayName
	s.TeamsAppId = decoded.TeamsAppId
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling TeamsAppUpgradedEventMessageDetail into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["initiator"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Initiator' for 'TeamsAppUpgradedEventMessageDetail': %+v", err)
		}
		s.Initiator = impl
	}

	return nil
}
