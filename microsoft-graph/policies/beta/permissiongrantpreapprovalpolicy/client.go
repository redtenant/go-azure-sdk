package permissiongrantpreapprovalpolicy

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PermissionGrantPreApprovalPolicyClient struct {
	Client *msgraph.Client
}

func NewPermissionGrantPreApprovalPolicyClientWithBaseURI(sdkApi sdkEnv.Api) (*PermissionGrantPreApprovalPolicyClient, error) {
	client, err := msgraph.NewClient(sdkApi, "permissiongrantpreapprovalpolicy", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PermissionGrantPreApprovalPolicyClient: %+v", err)
	}

	return &PermissionGrantPreApprovalPolicyClient{
		Client: client,
	}, nil
}
