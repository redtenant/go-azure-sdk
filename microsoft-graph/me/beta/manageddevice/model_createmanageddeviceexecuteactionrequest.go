package manageddevice

import (
	"github.com/redtenant/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/redtenant/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateManagedDeviceExecuteActionRequest struct {
	ActionName             *beta.ManagedDeviceRemoteAction `json:"actionName,omitempty"`
	CarrierUrl             nullable.Type[string]           `json:"carrierUrl,omitempty"`
	DeprovisionReason      nullable.Type[string]           `json:"deprovisionReason,omitempty"`
	DeviceIds              *[]string                       `json:"deviceIds,omitempty"`
	DeviceName             nullable.Type[string]           `json:"deviceName,omitempty"`
	KeepEnrollmentData     nullable.Type[bool]             `json:"keepEnrollmentData,omitempty"`
	KeepUserData           nullable.Type[bool]             `json:"keepUserData,omitempty"`
	NotificationBody       nullable.Type[string]           `json:"notificationBody,omitempty"`
	NotificationTitle      nullable.Type[string]           `json:"notificationTitle,omitempty"`
	OrganizationalUnitPath nullable.Type[string]           `json:"organizationalUnitPath,omitempty"`
	PersistEsimDataPlan    nullable.Type[bool]             `json:"persistEsimDataPlan,omitempty"`
}
