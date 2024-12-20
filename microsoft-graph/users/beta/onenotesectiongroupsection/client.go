package onenotesectiongroupsection

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OnenoteSectionGroupSectionClient struct {
	Client *msgraph.Client
}

func NewOnenoteSectionGroupSectionClientWithBaseURI(sdkApi sdkEnv.Api) (*OnenoteSectionGroupSectionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "onenotesectiongroupsection", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating OnenoteSectionGroupSectionClient: %+v", err)
	}

	return &OnenoteSectionGroupSectionClient{
		Client: client,
	}, nil
}
