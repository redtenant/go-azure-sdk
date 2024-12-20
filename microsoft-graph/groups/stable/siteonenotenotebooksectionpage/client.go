package siteonenotenotebooksectionpage

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SiteOnenoteNotebookSectionPageClient struct {
	Client *msgraph.Client
}

func NewSiteOnenoteNotebookSectionPageClientWithBaseURI(sdkApi sdkEnv.Api) (*SiteOnenoteNotebookSectionPageClient, error) {
	client, err := msgraph.NewClient(sdkApi, "siteonenotenotebooksectionpage", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SiteOnenoteNotebookSectionPageClient: %+v", err)
	}

	return &SiteOnenoteNotebookSectionPageClient{
		Client: client,
	}, nil
}
