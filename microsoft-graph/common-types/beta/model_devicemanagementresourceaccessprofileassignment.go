package beta

import (
	"encoding/json"
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = DeviceManagementResourceAccessProfileAssignment{}

type DeviceManagementResourceAccessProfileAssignment struct {
	// The administrator intent for the assignment of the profile.
	Intent *DeviceManagementResourceAccessProfileIntent `json:"intent,omitempty"`

	// The identifier of the source of the assignment.
	SourceId nullable.Type[string] `json:"sourceId,omitempty"`

	// Base type for assignment targets.
	Target DeviceAndAppManagementAssignmentTarget `json:"target"`

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

func (s DeviceManagementResourceAccessProfileAssignment) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = DeviceManagementResourceAccessProfileAssignment{}

func (s DeviceManagementResourceAccessProfileAssignment) MarshalJSON() ([]byte, error) {
	type wrapper DeviceManagementResourceAccessProfileAssignment
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DeviceManagementResourceAccessProfileAssignment: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DeviceManagementResourceAccessProfileAssignment: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.deviceManagementResourceAccessProfileAssignment"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DeviceManagementResourceAccessProfileAssignment: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &DeviceManagementResourceAccessProfileAssignment{}

func (s *DeviceManagementResourceAccessProfileAssignment) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		Intent    *DeviceManagementResourceAccessProfileIntent `json:"intent,omitempty"`
		SourceId  nullable.Type[string]                        `json:"sourceId,omitempty"`
		Id        *string                                      `json:"id,omitempty"`
		ODataId   *string                                      `json:"@odata.id,omitempty"`
		ODataType *string                                      `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.Intent = decoded.Intent
	s.SourceId = decoded.SourceId
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling DeviceManagementResourceAccessProfileAssignment into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["target"]; ok {
		impl, err := UnmarshalDeviceAndAppManagementAssignmentTargetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Target' for 'DeviceManagementResourceAccessProfileAssignment': %+v", err)
		}
		s.Target = impl
	}

	return nil
}
