package onenotesectiongroupparentnotebook

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OnenoteSectionGroupParentNotebookClient struct {
	Client *msgraph.Client
}

func NewOnenoteSectionGroupParentNotebookClientWithBaseURI(sdkApi sdkEnv.Api) (*OnenoteSectionGroupParentNotebookClient, error) {
	client, err := msgraph.NewClient(sdkApi, "onenotesectiongroupparentnotebook", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating OnenoteSectionGroupParentNotebookClient: %+v", err)
	}

	return &OnenoteSectionGroupParentNotebookClient{
		Client: client,
	}, nil
}
