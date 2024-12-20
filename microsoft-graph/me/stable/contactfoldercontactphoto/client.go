package contactfoldercontactphoto

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ContactFolderContactPhotoClient struct {
	Client *msgraph.Client
}

func NewContactFolderContactPhotoClientWithBaseURI(sdkApi sdkEnv.Api) (*ContactFolderContactPhotoClient, error) {
	client, err := msgraph.NewClient(sdkApi, "contactfoldercontactphoto", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ContactFolderContactPhotoClient: %+v", err)
	}

	return &ContactFolderContactPhotoClient{
		Client: client,
	}, nil
}
