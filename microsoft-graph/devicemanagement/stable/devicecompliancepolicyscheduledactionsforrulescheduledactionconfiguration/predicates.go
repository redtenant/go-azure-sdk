package devicecompliancepolicyscheduledactionsforrulescheduledactionconfiguration

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import "github.com/redtenant/go-azure-sdk/microsoft-graph/common-types/stable"

type DeviceComplianceActionItemOperationPredicate struct {
}

func (p DeviceComplianceActionItemOperationPredicate) Matches(input stable.DeviceComplianceActionItem) bool {

	return true
}
