package privilegedaccessgroupeligibilityschedule

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import "github.com/redtenant/go-azure-sdk/microsoft-graph/common-types/beta"

type PrivilegedAccessGroupEligibilityScheduleOperationPredicate struct {
}

func (p PrivilegedAccessGroupEligibilityScheduleOperationPredicate) Matches(input beta.PrivilegedAccessGroupEligibilitySchedule) bool {

	return true
}
