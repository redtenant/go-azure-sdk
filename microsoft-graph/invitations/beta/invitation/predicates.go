package invitation

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import "github.com/redtenant/go-azure-sdk/microsoft-graph/common-types/beta"

type InvitationOperationPredicate struct {
}

func (p InvitationOperationPredicate) Matches(input beta.Invitation) bool {

	return true
}
