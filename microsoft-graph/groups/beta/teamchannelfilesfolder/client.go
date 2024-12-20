package teamchannelfilesfolder

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamChannelFilesFolderClient struct {
	Client *msgraph.Client
}

func NewTeamChannelFilesFolderClientWithBaseURI(sdkApi sdkEnv.Api) (*TeamChannelFilesFolderClient, error) {
	client, err := msgraph.NewClient(sdkApi, "teamchannelfilesfolder", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating TeamChannelFilesFolderClient: %+v", err)
	}

	return &TeamChannelFilesFolderClient{
		Client: client,
	}, nil
}
