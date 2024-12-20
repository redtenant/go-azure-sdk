package contactfolderchildfoldercontactextension

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ContactFolderChildFolderContactExtensionClient struct {
	Client *msgraph.Client
}

func NewContactFolderChildFolderContactExtensionClientWithBaseURI(sdkApi sdkEnv.Api) (*ContactFolderChildFolderContactExtensionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "contactfolderchildfoldercontactextension", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ContactFolderChildFolderContactExtensionClient: %+v", err)
	}

	return &ContactFolderChildFolderContactExtensionClient{
		Client: client,
	}, nil
}
