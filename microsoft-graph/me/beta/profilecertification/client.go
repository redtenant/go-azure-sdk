package profilecertification

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProfileCertificationClient struct {
	Client *msgraph.Client
}

func NewProfileCertificationClientWithBaseURI(sdkApi sdkEnv.Api) (*ProfileCertificationClient, error) {
	client, err := msgraph.NewClient(sdkApi, "profilecertification", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ProfileCertificationClient: %+v", err)
	}

	return &ProfileCertificationClient{
		Client: client,
	}, nil
}
