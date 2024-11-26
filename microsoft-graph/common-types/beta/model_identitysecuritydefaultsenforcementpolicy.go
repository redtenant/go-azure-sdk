package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IdentitySecurityDefaultsEnforcementPolicy struct {
	// If set to true, Microsoft Entra security defaults are enabled for the tenant.
	IsEnabled *bool `json:"isEnabled,omitempty"`
}
