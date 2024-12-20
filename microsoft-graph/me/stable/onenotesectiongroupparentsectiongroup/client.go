package onenotesectiongroupparentsectiongroup

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OnenoteSectionGroupParentSectionGroupClient struct {
	Client *msgraph.Client
}

func NewOnenoteSectionGroupParentSectionGroupClientWithBaseURI(sdkApi sdkEnv.Api) (*OnenoteSectionGroupParentSectionGroupClient, error) {
	client, err := msgraph.NewClient(sdkApi, "onenotesectiongroupparentsectiongroup", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating OnenoteSectionGroupParentSectionGroupClient: %+v", err)
	}

	return &OnenoteSectionGroupParentSectionGroupClient{
		Client: client,
	}, nil
}
