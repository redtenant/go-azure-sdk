package siteonenotenotebooksectiongroupsectionpageparentnotebook

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SiteOnenoteNotebookSectionGroupSectionPageParentNotebookClient struct {
	Client *msgraph.Client
}

func NewSiteOnenoteNotebookSectionGroupSectionPageParentNotebookClientWithBaseURI(sdkApi sdkEnv.Api) (*SiteOnenoteNotebookSectionGroupSectionPageParentNotebookClient, error) {
	client, err := msgraph.NewClient(sdkApi, "siteonenotenotebooksectiongroupsectionpageparentnotebook", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SiteOnenoteNotebookSectionGroupSectionPageParentNotebookClient: %+v", err)
	}

	return &SiteOnenoteNotebookSectionGroupSectionPageParentNotebookClient{
		Client: client,
	}, nil
}
