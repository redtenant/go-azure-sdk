package onlinemeetingregistrationcustomquestion

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OnlineMeetingRegistrationCustomQuestionClient struct {
	Client *msgraph.Client
}

func NewOnlineMeetingRegistrationCustomQuestionClientWithBaseURI(sdkApi sdkEnv.Api) (*OnlineMeetingRegistrationCustomQuestionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "onlinemeetingregistrationcustomquestion", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating OnlineMeetingRegistrationCustomQuestionClient: %+v", err)
	}

	return &OnlineMeetingRegistrationCustomQuestionClient{
		Client: client,
	}, nil
}
