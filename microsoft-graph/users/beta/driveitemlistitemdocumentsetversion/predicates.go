package driveitemlistitemdocumentsetversion

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import "github.com/redtenant/go-azure-sdk/microsoft-graph/common-types/beta"

type DocumentSetVersionOperationPredicate struct {
}

func (p DocumentSetVersionOperationPredicate) Matches(input beta.DocumentSetVersion) bool {

	return true
}
