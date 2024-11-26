package entitlementmanagementaccesspackageaccesspackageassignmentpolicycustomextensionhandlercustomextension

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementAccessPackageAccessPackageAssignmentPolicyCustomExtensionHandlerCustomExtensionClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementAccessPackageAccessPackageAssignmentPolicyCustomExtensionHandlerCustomExtensionClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementAccessPackageAccessPackageAssignmentPolicyCustomExtensionHandlerCustomExtensionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementaccesspackageaccesspackageassignmentpolicycustomextensionhandlercustomextension", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementAccessPackageAccessPackageAssignmentPolicyCustomExtensionHandlerCustomExtensionClient: %+v", err)
	}

	return &EntitlementManagementAccessPackageAccessPackageAssignmentPolicyCustomExtensionHandlerCustomExtensionClient{
		Client: client,
	}, nil
}
