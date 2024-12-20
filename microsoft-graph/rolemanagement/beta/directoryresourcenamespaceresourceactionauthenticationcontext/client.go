package directoryresourcenamespaceresourceactionauthenticationcontext

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DirectoryResourceNamespaceResourceActionAuthenticationContextClient struct {
	Client *msgraph.Client
}

func NewDirectoryResourceNamespaceResourceActionAuthenticationContextClientWithBaseURI(sdkApi sdkEnv.Api) (*DirectoryResourceNamespaceResourceActionAuthenticationContextClient, error) {
	client, err := msgraph.NewClient(sdkApi, "directoryresourcenamespaceresourceactionauthenticationcontext", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DirectoryResourceNamespaceResourceActionAuthenticationContextClient: %+v", err)
	}

	return &DirectoryResourceNamespaceResourceActionAuthenticationContextClient{
		Client: client,
	}, nil
}
