package onlinemeeting

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import "github.com/redtenant/go-azure-sdk/microsoft-graph/common-types/beta"

type OnlineMeetingOperationPredicate struct {
}

func (p OnlineMeetingOperationPredicate) Matches(input beta.OnlineMeeting) bool {

	return true
}
