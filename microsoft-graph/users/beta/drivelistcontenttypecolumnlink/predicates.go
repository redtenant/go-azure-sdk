package drivelistcontenttypecolumnlink

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import "github.com/redtenant/go-azure-sdk/microsoft-graph/common-types/beta"

type ColumnLinkOperationPredicate struct {
}

func (p ColumnLinkOperationPredicate) Matches(input beta.ColumnLink) bool {

	return true
}
