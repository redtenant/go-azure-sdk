package beta

import (
	"encoding/json"
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ MeetingInfo = OrganizerMeetingInfo{}

type OrganizerMeetingInfo struct {
	Organizer IdentitySet `json:"organizer"`

	// Fields inherited from MeetingInfo

	AllowConversationWithoutHost nullable.Type[bool] `json:"allowConversationWithoutHost,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s OrganizerMeetingInfo) MeetingInfo() BaseMeetingInfoImpl {
	return BaseMeetingInfoImpl{
		AllowConversationWithoutHost: s.AllowConversationWithoutHost,
		ODataId:                      s.ODataId,
		ODataType:                    s.ODataType,
	}
}

var _ json.Marshaler = OrganizerMeetingInfo{}

func (s OrganizerMeetingInfo) MarshalJSON() ([]byte, error) {
	type wrapper OrganizerMeetingInfo
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling OrganizerMeetingInfo: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling OrganizerMeetingInfo: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.organizerMeetingInfo"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling OrganizerMeetingInfo: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &OrganizerMeetingInfo{}

func (s *OrganizerMeetingInfo) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		AllowConversationWithoutHost nullable.Type[bool] `json:"allowConversationWithoutHost,omitempty"`
		ODataId                      *string             `json:"@odata.id,omitempty"`
		ODataType                    *string             `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.AllowConversationWithoutHost = decoded.AllowConversationWithoutHost
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling OrganizerMeetingInfo into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["organizer"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Organizer' for 'OrganizerMeetingInfo': %+v", err)
		}
		s.Organizer = impl
	}

	return nil
}
