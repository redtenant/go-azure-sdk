package beta

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/redtenant/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServiceAnnouncementBase interface {
	Entity
	ServiceAnnouncementBase() BaseServiceAnnouncementBaseImpl
}

var _ ServiceAnnouncementBase = BaseServiceAnnouncementBaseImpl{}

type BaseServiceAnnouncementBaseImpl struct {
	// Extra details about service event. This property doesn't support filters.
	Details *[]KeyValuePair `json:"details,omitempty"`

	// The end time of the service event.
	EndDateTime nullable.Type[string] `json:"endDateTime,omitempty"`

	// The last modified time of the service event.
	LastModifiedDateTime *string `json:"lastModifiedDateTime,omitempty"`

	// The start time of the service event.
	StartDateTime *string `json:"startDateTime,omitempty"`

	// The title of the service event.
	Title *string `json:"title,omitempty"`

	// Fields inherited from Entity

	// The unique identifier for an entity. Read-only.
	Id *string `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseServiceAnnouncementBaseImpl) ServiceAnnouncementBase() BaseServiceAnnouncementBaseImpl {
	return s
}

func (s BaseServiceAnnouncementBaseImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ ServiceAnnouncementBase = RawServiceAnnouncementBaseImpl{}

// RawServiceAnnouncementBaseImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawServiceAnnouncementBaseImpl struct {
	serviceAnnouncementBase BaseServiceAnnouncementBaseImpl
	Type                    string
	Values                  map[string]interface{}
}

func (s RawServiceAnnouncementBaseImpl) ServiceAnnouncementBase() BaseServiceAnnouncementBaseImpl {
	return s.serviceAnnouncementBase
}

func (s RawServiceAnnouncementBaseImpl) Entity() BaseEntityImpl {
	return s.serviceAnnouncementBase.Entity()
}

var _ json.Marshaler = BaseServiceAnnouncementBaseImpl{}

func (s BaseServiceAnnouncementBaseImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseServiceAnnouncementBaseImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseServiceAnnouncementBaseImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseServiceAnnouncementBaseImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.serviceAnnouncementBase"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseServiceAnnouncementBaseImpl: %+v", err)
	}

	return encoded, nil
}

func UnmarshalServiceAnnouncementBaseImplementation(input []byte) (ServiceAnnouncementBase, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling ServiceAnnouncementBase into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.serviceHealthIssue") {
		var out ServiceHealthIssue
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ServiceHealthIssue: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.serviceUpdateMessage") {
		var out ServiceUpdateMessage
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ServiceUpdateMessage: %+v", err)
		}
		return out, nil
	}

	var parent BaseServiceAnnouncementBaseImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseServiceAnnouncementBaseImpl: %+v", err)
	}

	return RawServiceAnnouncementBaseImpl{
		serviceAnnouncementBase: parent,
		Type:                    value,
		Values:                  temp,
	}, nil

}
