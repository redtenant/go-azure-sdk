package devicecompliancepolicyuserstatus

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import "github.com/redtenant/go-azure-sdk/microsoft-graph/common-types/beta"

type DeviceComplianceUserStatusOperationPredicate struct {
}

func (p DeviceComplianceUserStatusOperationPredicate) Matches(input beta.DeviceComplianceUserStatus) bool {

	return true
}
