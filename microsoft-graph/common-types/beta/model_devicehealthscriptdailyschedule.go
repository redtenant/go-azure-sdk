package beta

import (
	"encoding/json"
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DeviceHealthScriptTimeSchedule = DeviceHealthScriptDailySchedule{}

type DeviceHealthScriptDailySchedule struct {

	// Fields inherited from DeviceHealthScriptTimeSchedule

	// At what time the script is scheduled to run. This collection can contain a maximum of 20 elements.
	Time nullable.Type[string] `json:"time,omitempty"`

	// Indicate if the time is Utc or client local time.
	UseUtc *bool `json:"useUtc,omitempty"`

	// Fields inherited from DeviceHealthScriptRunSchedule

	// The x value of every x hours for hourly schedule, every x days for Daily Schedule, every x weeks for weekly schedule,
	// every x months for Monthly Schedule. Valid values 1 to 23
	Interval *int64 `json:"interval,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s DeviceHealthScriptDailySchedule) DeviceHealthScriptTimeSchedule() BaseDeviceHealthScriptTimeScheduleImpl {
	return BaseDeviceHealthScriptTimeScheduleImpl{
		Time:      s.Time,
		UseUtc:    s.UseUtc,
		Interval:  s.Interval,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

func (s DeviceHealthScriptDailySchedule) DeviceHealthScriptRunSchedule() BaseDeviceHealthScriptRunScheduleImpl {
	return BaseDeviceHealthScriptRunScheduleImpl{
		Interval:  s.Interval,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = DeviceHealthScriptDailySchedule{}

func (s DeviceHealthScriptDailySchedule) MarshalJSON() ([]byte, error) {
	type wrapper DeviceHealthScriptDailySchedule
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DeviceHealthScriptDailySchedule: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DeviceHealthScriptDailySchedule: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.deviceHealthScriptDailySchedule"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DeviceHealthScriptDailySchedule: %+v", err)
	}

	return encoded, nil
}
