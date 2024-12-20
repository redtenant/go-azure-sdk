package onenotepageparentsection

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OnenotePageParentSectionClient struct {
	Client *msgraph.Client
}

func NewOnenotePageParentSectionClientWithBaseURI(sdkApi sdkEnv.Api) (*OnenotePageParentSectionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "onenotepageparentsection", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating OnenotePageParentSectionClient: %+v", err)
	}

	return &OnenotePageParentSectionClient{
		Client: client,
	}, nil
}
