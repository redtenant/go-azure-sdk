package userflowattribute

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import "github.com/redtenant/go-azure-sdk/microsoft-graph/common-types/stable"

type IdentityUserFlowAttributeOperationPredicate struct {
}

func (p IdentityUserFlowAttributeOperationPredicate) Matches(input stable.IdentityUserFlowAttribute) bool {

	return true
}
