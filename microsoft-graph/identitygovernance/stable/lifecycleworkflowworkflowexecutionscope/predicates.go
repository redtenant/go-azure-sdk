package lifecycleworkflowworkflowexecutionscope

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import "github.com/redtenant/go-azure-sdk/microsoft-graph/common-types/stable"

type IdentityGovernanceUserProcessingResultOperationPredicate struct {
}

func (p IdentityGovernanceUserProcessingResultOperationPredicate) Matches(input stable.IdentityGovernanceUserProcessingResult) bool {

	return true
}
