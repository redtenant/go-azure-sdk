package template

import (
	"github.com/redtenant/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/redtenant/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateTemplateInstanceRequest struct {
	Description     nullable.Type[string]                   `json:"description,omitempty"`
	DisplayName     nullable.Type[string]                   `json:"displayName,omitempty"`
	RoleScopeTagIds *[]string                               `json:"roleScopeTagIds,omitempty"`
	SettingsDelta   *[]beta.DeviceManagementSettingInstance `json:"settingsDelta,omitempty"`
}
