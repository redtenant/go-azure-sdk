package onenotenotebooksectionpage

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OnenoteNotebookSectionPageClient struct {
	Client *msgraph.Client
}

func NewOnenoteNotebookSectionPageClientWithBaseURI(sdkApi sdkEnv.Api) (*OnenoteNotebookSectionPageClient, error) {
	client, err := msgraph.NewClient(sdkApi, "onenotenotebooksectionpage", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating OnenoteNotebookSectionPageClient: %+v", err)
	}

	return &OnenoteNotebookSectionPageClient{
		Client: client,
	}, nil
}
