package stable

import (
	"github.com/redtenant/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthenticationMethodsRegistrationCampaignIncludeTarget struct {
	// The object identifier of a Microsoft Entra user or group.
	Id *string `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	TargetType *AuthenticationMethodTargetType `json:"targetType,omitempty"`

	// The authentication method that the user is prompted to register. The value must be microsoftAuthenticator.
	TargetedAuthenticationMethod nullable.Type[string] `json:"targetedAuthenticationMethod,omitempty"`
}
