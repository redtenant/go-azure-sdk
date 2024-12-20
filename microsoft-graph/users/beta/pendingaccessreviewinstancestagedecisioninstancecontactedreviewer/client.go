package pendingaccessreviewinstancestagedecisioninstancecontactedreviewer

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PendingAccessReviewInstanceStageDecisionInstanceContactedReviewerClient struct {
	Client *msgraph.Client
}

func NewPendingAccessReviewInstanceStageDecisionInstanceContactedReviewerClientWithBaseURI(sdkApi sdkEnv.Api) (*PendingAccessReviewInstanceStageDecisionInstanceContactedReviewerClient, error) {
	client, err := msgraph.NewClient(sdkApi, "pendingaccessreviewinstancestagedecisioninstancecontactedreviewer", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PendingAccessReviewInstanceStageDecisionInstanceContactedReviewerClient: %+v", err)
	}

	return &PendingAccessReviewInstanceStageDecisionInstanceContactedReviewerClient{
		Client: client,
	}, nil
}
