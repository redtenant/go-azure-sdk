package entitlementmanagementresourcerequestcatalogaccesspackage

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementResourceRequestCatalogAccessPackageClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementResourceRequestCatalogAccessPackageClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementResourceRequestCatalogAccessPackageClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementresourcerequestcatalogaccesspackage", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementResourceRequestCatalogAccessPackageClient: %+v", err)
	}

	return &EntitlementManagementResourceRequestCatalogAccessPackageClient{
		Client: client,
	}, nil
}
