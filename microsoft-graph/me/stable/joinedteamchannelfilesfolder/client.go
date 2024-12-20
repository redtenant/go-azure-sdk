package joinedteamchannelfilesfolder

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type JoinedTeamChannelFilesFolderClient struct {
	Client *msgraph.Client
}

func NewJoinedTeamChannelFilesFolderClientWithBaseURI(sdkApi sdkEnv.Api) (*JoinedTeamChannelFilesFolderClient, error) {
	client, err := msgraph.NewClient(sdkApi, "joinedteamchannelfilesfolder", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating JoinedTeamChannelFilesFolderClient: %+v", err)
	}

	return &JoinedTeamChannelFilesFolderClient{
		Client: client,
	}, nil
}
