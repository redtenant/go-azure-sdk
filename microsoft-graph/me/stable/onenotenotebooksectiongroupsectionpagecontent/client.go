package onenotenotebooksectiongroupsectionpagecontent

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OnenoteNotebookSectionGroupSectionPageContentClient struct {
	Client *msgraph.Client
}

func NewOnenoteNotebookSectionGroupSectionPageContentClientWithBaseURI(sdkApi sdkEnv.Api) (*OnenoteNotebookSectionGroupSectionPageContentClient, error) {
	client, err := msgraph.NewClient(sdkApi, "onenotenotebooksectiongroupsectionpagecontent", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating OnenoteNotebookSectionGroupSectionPageContentClient: %+v", err)
	}

	return &OnenoteNotebookSectionGroupSectionPageContentClient{
		Client: client,
	}, nil
}
