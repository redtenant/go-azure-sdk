package windowsautopilotdeploymentprofileassigneddevice

import (
	"github.com/redtenant/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AssignWindowsAutopilotDeploymentProfileAssignedDeviceUserToDeviceRequest struct {
	AddressableUserName nullable.Type[string] `json:"addressableUserName,omitempty"`
	UserPrincipalName   nullable.Type[string] `json:"userPrincipalName,omitempty"`
}
