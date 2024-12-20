package onenotenotebooksection

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OnenoteNotebookSectionClient struct {
	Client *msgraph.Client
}

func NewOnenoteNotebookSectionClientWithBaseURI(sdkApi sdkEnv.Api) (*OnenoteNotebookSectionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "onenotenotebooksection", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating OnenoteNotebookSectionClient: %+v", err)
	}

	return &OnenoteNotebookSectionClient{
		Client: client,
	}, nil
}
