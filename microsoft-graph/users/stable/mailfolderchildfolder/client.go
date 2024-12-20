package mailfolderchildfolder

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MailFolderChildFolderClient struct {
	Client *msgraph.Client
}

func NewMailFolderChildFolderClientWithBaseURI(sdkApi sdkEnv.Api) (*MailFolderChildFolderClient, error) {
	client, err := msgraph.NewClient(sdkApi, "mailfolderchildfolder", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MailFolderChildFolderClient: %+v", err)
	}

	return &MailFolderChildFolderClient{
		Client: client,
	}, nil
}
