package devicelocalcredential

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import "github.com/redtenant/go-azure-sdk/microsoft-graph/common-types/stable"

type DeviceLocalCredentialInfoOperationPredicate struct {
}

func (p DeviceLocalCredentialInfoOperationPredicate) Matches(input stable.DeviceLocalCredentialInfo) bool {

	return true
}
