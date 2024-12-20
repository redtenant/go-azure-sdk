package entitlementmanagementaccesspackageaccesspackageassignmentpolicycustomextensionhandler

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementAccessPackageAccessPackageAssignmentPolicyCustomExtensionHandlerClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementAccessPackageAccessPackageAssignmentPolicyCustomExtensionHandlerClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementAccessPackageAccessPackageAssignmentPolicyCustomExtensionHandlerClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementaccesspackageaccesspackageassignmentpolicycustomextensionhandler", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementAccessPackageAccessPackageAssignmentPolicyCustomExtensionHandlerClient: %+v", err)
	}

	return &EntitlementManagementAccessPackageAccessPackageAssignmentPolicyCustomExtensionHandlerClient{
		Client: client,
	}, nil
}
