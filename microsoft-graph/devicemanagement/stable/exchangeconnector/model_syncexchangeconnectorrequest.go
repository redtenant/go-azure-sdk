package exchangeconnector

import (
	"github.com/redtenant/go-azure-sdk/microsoft-graph/common-types/stable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SyncExchangeConnectorRequest struct {
	// The type of Exchange Connector sync requested.
	SyncType *stable.DeviceManagementExchangeConnectorSyncType `json:"syncType,omitempty"`
}
