package b2xuserflowapiconnectorconfigurationpostfederationsignup

import (
	"github.com/redtenant/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateB2xUserFlowApiConnectorConfigurationPostFederationSignupUploadClientCertificateRequest struct {
	Password    nullable.Type[string] `json:"password,omitempty"`
	Pkcs12Value nullable.Type[string] `json:"pkcs12Value,omitempty"`
}
