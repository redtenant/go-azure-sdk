package tokenlifetimepolicy

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import "github.com/redtenant/go-azure-sdk/microsoft-graph/common-types/beta"

type TokenLifetimePolicyOperationPredicate struct {
}

func (p TokenLifetimePolicyOperationPredicate) Matches(input beta.TokenLifetimePolicy) bool {

	return true
}
