package onenotenotebooksectiongroupsectionparentnotebook

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OnenoteNotebookSectionGroupSectionParentNotebookClient struct {
	Client *msgraph.Client
}

func NewOnenoteNotebookSectionGroupSectionParentNotebookClientWithBaseURI(sdkApi sdkEnv.Api) (*OnenoteNotebookSectionGroupSectionParentNotebookClient, error) {
	client, err := msgraph.NewClient(sdkApi, "onenotenotebooksectiongroupsectionparentnotebook", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating OnenoteNotebookSectionGroupSectionParentNotebookClient: %+v", err)
	}

	return &OnenoteNotebookSectionGroupSectionParentNotebookClient{
		Client: client,
	}, nil
}
