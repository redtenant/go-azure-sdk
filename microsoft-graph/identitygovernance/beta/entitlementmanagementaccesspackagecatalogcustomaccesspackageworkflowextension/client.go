package entitlementmanagementaccesspackagecatalogcustomaccesspackageworkflowextension

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementAccessPackageCatalogCustomAccessPackageWorkflowExtensionClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementAccessPackageCatalogCustomAccessPackageWorkflowExtensionClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementAccessPackageCatalogCustomAccessPackageWorkflowExtensionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementaccesspackagecatalogcustomaccesspackageworkflowextension", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementAccessPackageCatalogCustomAccessPackageWorkflowExtensionClient: %+v", err)
	}

	return &EntitlementManagementAccessPackageCatalogCustomAccessPackageWorkflowExtensionClient{
		Client: client,
	}, nil
}
