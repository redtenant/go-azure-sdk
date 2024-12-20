package windowsautopilotdeploymentprofileassigneddevice

import (
	"github.com/redtenant/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AssignWindowsAutopilotDeploymentProfileAssignedDeviceResourceAccountToDeviceRequest struct {
	AddressableUserName nullable.Type[string] `json:"addressableUserName,omitempty"`
	ResourceAccountName nullable.Type[string] `json:"resourceAccountName,omitempty"`
	UserPrincipalName   nullable.Type[string] `json:"userPrincipalName,omitempty"`
}
