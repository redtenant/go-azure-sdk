package v2024_07_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/resource-manager/batch/2024-07-01/application"
	"github.com/redtenant/go-azure-sdk/resource-manager/batch/2024-07-01/applicationpackage"
	"github.com/redtenant/go-azure-sdk/resource-manager/batch/2024-07-01/batchaccount"
	"github.com/redtenant/go-azure-sdk/resource-manager/batch/2024-07-01/batchmanagements"
	"github.com/redtenant/go-azure-sdk/resource-manager/batch/2024-07-01/certificate"
	"github.com/redtenant/go-azure-sdk/resource-manager/batch/2024-07-01/location"
	"github.com/redtenant/go-azure-sdk/resource-manager/batch/2024-07-01/networksecurityperimeter"
	"github.com/redtenant/go-azure-sdk/resource-manager/batch/2024-07-01/pool"
	"github.com/redtenant/go-azure-sdk/resource-manager/batch/2024-07-01/privateendpointconnection"
	"github.com/redtenant/go-azure-sdk/resource-manager/batch/2024-07-01/privatelinkresource"
	"github.com/redtenant/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

type Client struct {
	Application               *application.ApplicationClient
	ApplicationPackage        *applicationpackage.ApplicationPackageClient
	BatchAccount              *batchaccount.BatchAccountClient
	BatchManagements          *batchmanagements.BatchManagementsClient
	Certificate               *certificate.CertificateClient
	Location                  *location.LocationClient
	NetworkSecurityPerimeter  *networksecurityperimeter.NetworkSecurityPerimeterClient
	Pool                      *pool.PoolClient
	PrivateEndpointConnection *privateendpointconnection.PrivateEndpointConnectionClient
	PrivateLinkResource       *privatelinkresource.PrivateLinkResourceClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	applicationClient, err := application.NewApplicationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Application client: %+v", err)
	}
	configureFunc(applicationClient.Client)

	applicationPackageClient, err := applicationpackage.NewApplicationPackageClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ApplicationPackage client: %+v", err)
	}
	configureFunc(applicationPackageClient.Client)

	batchAccountClient, err := batchaccount.NewBatchAccountClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building BatchAccount client: %+v", err)
	}
	configureFunc(batchAccountClient.Client)

	batchManagementsClient, err := batchmanagements.NewBatchManagementsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building BatchManagements client: %+v", err)
	}
	configureFunc(batchManagementsClient.Client)

	certificateClient, err := certificate.NewCertificateClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Certificate client: %+v", err)
	}
	configureFunc(certificateClient.Client)

	locationClient, err := location.NewLocationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Location client: %+v", err)
	}
	configureFunc(locationClient.Client)

	networkSecurityPerimeterClient, err := networksecurityperimeter.NewNetworkSecurityPerimeterClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building NetworkSecurityPerimeter client: %+v", err)
	}
	configureFunc(networkSecurityPerimeterClient.Client)

	poolClient, err := pool.NewPoolClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Pool client: %+v", err)
	}
	configureFunc(poolClient.Client)

	privateEndpointConnectionClient, err := privateendpointconnection.NewPrivateEndpointConnectionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PrivateEndpointConnection client: %+v", err)
	}
	configureFunc(privateEndpointConnectionClient.Client)

	privateLinkResourceClient, err := privatelinkresource.NewPrivateLinkResourceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PrivateLinkResource client: %+v", err)
	}
	configureFunc(privateLinkResourceClient.Client)

	return &Client{
		Application:               applicationClient,
		ApplicationPackage:        applicationPackageClient,
		BatchAccount:              batchAccountClient,
		BatchManagements:          batchManagementsClient,
		Certificate:               certificateClient,
		Location:                  locationClient,
		NetworkSecurityPerimeter:  networkSecurityPerimeterClient,
		Pool:                      poolClient,
		PrivateEndpointConnection: privateEndpointConnectionClient,
		PrivateLinkResource:       privateLinkResourceClient,
	}, nil
}
