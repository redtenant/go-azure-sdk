package joinedteaminstalledapp

import (
	"github.com/redtenant/go-azure-sdk/microsoft-graph/common-types/stable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateJoinedTeamInstalledAppUpgradeRequest struct {
	ConsentedPermissionSet *stable.TeamsAppPermissionSet `json:"consentedPermissionSet,omitempty"`
}
