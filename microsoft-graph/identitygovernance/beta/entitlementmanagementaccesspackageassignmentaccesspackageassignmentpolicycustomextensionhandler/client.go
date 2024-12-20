package entitlementmanagementaccesspackageassignmentaccesspackageassignmentpolicycustomextensionhandler

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementAccessPackageAssignmentAccessPackageAssignmentPolicyCustomExtensionHandlerClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementAccessPackageAssignmentAccessPackageAssignmentPolicyCustomExtensionHandlerClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementAccessPackageAssignmentAccessPackageAssignmentPolicyCustomExtensionHandlerClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementaccesspackageassignmentaccesspackageassignmentpolicycustomextensionhandler", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementAccessPackageAssignmentAccessPackageAssignmentPolicyCustomExtensionHandlerClient: %+v", err)
	}

	return &EntitlementManagementAccessPackageAssignmentAccessPackageAssignmentPolicyCustomExtensionHandlerClient{
		Client: client,
	}, nil
}
