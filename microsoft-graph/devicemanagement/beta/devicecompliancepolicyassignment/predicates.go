package devicecompliancepolicyassignment

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import "github.com/redtenant/go-azure-sdk/microsoft-graph/common-types/beta"

type DeviceCompliancePolicyAssignmentOperationPredicate struct {
}

func (p DeviceCompliancePolicyAssignmentOperationPredicate) Matches(input beta.DeviceCompliancePolicyAssignment) bool {

	return true
}
