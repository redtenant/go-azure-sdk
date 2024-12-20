package siteonenotenotebooksectiongroupsectionpagecontent

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SiteOnenoteNotebookSectionGroupSectionPageContentClient struct {
	Client *msgraph.Client
}

func NewSiteOnenoteNotebookSectionGroupSectionPageContentClientWithBaseURI(sdkApi sdkEnv.Api) (*SiteOnenoteNotebookSectionGroupSectionPageContentClient, error) {
	client, err := msgraph.NewClient(sdkApi, "siteonenotenotebooksectiongroupsectionpagecontent", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SiteOnenoteNotebookSectionGroupSectionPageContentClient: %+v", err)
	}

	return &SiteOnenoteNotebookSectionGroupSectionPageContentClient{
		Client: client,
	}, nil
}
