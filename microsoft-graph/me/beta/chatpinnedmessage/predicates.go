package chatpinnedmessage

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import "github.com/redtenant/go-azure-sdk/microsoft-graph/common-types/beta"

type PinnedChatMessageInfoOperationPredicate struct {
}

func (p PinnedChatMessageInfoOperationPredicate) Matches(input beta.PinnedChatMessageInfo) bool {

	return true
}
