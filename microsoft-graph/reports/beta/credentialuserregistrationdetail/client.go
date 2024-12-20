package credentialuserregistrationdetail

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CredentialUserRegistrationDetailClient struct {
	Client *msgraph.Client
}

func NewCredentialUserRegistrationDetailClientWithBaseURI(sdkApi sdkEnv.Api) (*CredentialUserRegistrationDetailClient, error) {
	client, err := msgraph.NewClient(sdkApi, "credentialuserregistrationdetail", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CredentialUserRegistrationDetailClient: %+v", err)
	}

	return &CredentialUserRegistrationDetailClient{
		Client: client,
	}, nil
}
