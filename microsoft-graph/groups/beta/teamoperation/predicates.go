package teamoperation

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import "github.com/redtenant/go-azure-sdk/microsoft-graph/common-types/beta"

type TeamsAsyncOperationOperationPredicate struct {
}

func (p TeamsAsyncOperationOperationPredicate) Matches(input beta.TeamsAsyncOperation) bool {

	return true
}
