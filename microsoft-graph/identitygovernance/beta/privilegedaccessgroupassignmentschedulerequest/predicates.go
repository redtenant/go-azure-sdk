package privilegedaccessgroupassignmentschedulerequest

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import "github.com/redtenant/go-azure-sdk/microsoft-graph/common-types/beta"

type PrivilegedAccessGroupAssignmentScheduleRequestOperationPredicate struct {
}

func (p PrivilegedAccessGroupAssignmentScheduleRequestOperationPredicate) Matches(input beta.PrivilegedAccessGroupAssignmentScheduleRequest) bool {

	return true
}
