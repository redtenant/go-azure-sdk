package entitlementmanagementroleassignmentapprovalstep

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import "github.com/redtenant/go-azure-sdk/microsoft-graph/common-types/beta"

type ApprovalStepOperationPredicate struct {
}

func (p ApprovalStepOperationPredicate) Matches(input beta.ApprovalStep) bool {

	return true
}
