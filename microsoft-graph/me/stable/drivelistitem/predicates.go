package drivelistitem

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import "github.com/redtenant/go-azure-sdk/microsoft-graph/common-types/stable"

type ListItemOperationPredicate struct {
}

func (p ListItemOperationPredicate) Matches(input stable.ListItem) bool {

	return true
}
