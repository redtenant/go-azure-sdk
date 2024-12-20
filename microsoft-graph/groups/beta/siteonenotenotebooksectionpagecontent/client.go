package siteonenotenotebooksectionpagecontent

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SiteOnenoteNotebookSectionPageContentClient struct {
	Client *msgraph.Client
}

func NewSiteOnenoteNotebookSectionPageContentClientWithBaseURI(sdkApi sdkEnv.Api) (*SiteOnenoteNotebookSectionPageContentClient, error) {
	client, err := msgraph.NewClient(sdkApi, "siteonenotenotebooksectionpagecontent", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SiteOnenoteNotebookSectionPageContentClient: %+v", err)
	}

	return &SiteOnenoteNotebookSectionPageContentClient{
		Client: client,
	}, nil
}
