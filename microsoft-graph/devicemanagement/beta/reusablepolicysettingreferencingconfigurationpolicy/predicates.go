package reusablepolicysettingreferencingconfigurationpolicy

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import "github.com/redtenant/go-azure-sdk/microsoft-graph/common-types/beta"

type DeviceManagementConfigurationPolicyOperationPredicate struct {
}

func (p DeviceManagementConfigurationPolicyOperationPredicate) Matches(input beta.DeviceManagementConfigurationPolicy) bool {

	return true
}

type DeviceManagementConfigurationPolicyAssignmentOperationPredicate struct {
}

func (p DeviceManagementConfigurationPolicyAssignmentOperationPredicate) Matches(input beta.DeviceManagementConfigurationPolicyAssignment) bool {

	return true
}
