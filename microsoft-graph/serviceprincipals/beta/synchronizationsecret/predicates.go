package synchronizationsecret

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import "github.com/redtenant/go-azure-sdk/microsoft-graph/common-types/beta"

type SynchronizationSecretKeyStringValuePairOperationPredicate struct {
}

func (p SynchronizationSecretKeyStringValuePairOperationPredicate) Matches(input beta.SynchronizationSecretKeyStringValuePair) bool {

	return true
}
