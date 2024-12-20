package comanageddevice

import (
	"github.com/redtenant/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/redtenant/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateComanagedDeviceOverrideComplianceStateRequest struct {
	// Administrator configured device compliance state Enum
	ComplianceState *beta.AdministratorConfiguredDeviceComplianceState `json:"complianceState,omitempty"`

	RemediationUrl nullable.Type[string] `json:"remediationUrl,omitempty"`
}
