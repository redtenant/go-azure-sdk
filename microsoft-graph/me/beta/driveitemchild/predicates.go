package driveitemchild

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import "github.com/redtenant/go-azure-sdk/microsoft-graph/common-types/beta"

type DriveItemOperationPredicate struct {
}

func (p DriveItemOperationPredicate) Matches(input beta.DriveItem) bool {

	return true
}
