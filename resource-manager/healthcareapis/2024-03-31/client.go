package v2024_03_31

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/resource-manager/healthcareapis/2024-03-31/collection"
	"github.com/redtenant/go-azure-sdk/resource-manager/healthcareapis/2024-03-31/dicomservices"
	"github.com/redtenant/go-azure-sdk/resource-manager/healthcareapis/2024-03-31/fhirservices"
	"github.com/redtenant/go-azure-sdk/resource-manager/healthcareapis/2024-03-31/iotconnectors"
	"github.com/redtenant/go-azure-sdk/resource-manager/healthcareapis/2024-03-31/privateendpointconnections"
	"github.com/redtenant/go-azure-sdk/resource-manager/healthcareapis/2024-03-31/privatelinkresources"
	"github.com/redtenant/go-azure-sdk/resource-manager/healthcareapis/2024-03-31/proxy"
	"github.com/redtenant/go-azure-sdk/resource-manager/healthcareapis/2024-03-31/resource"
	"github.com/redtenant/go-azure-sdk/resource-manager/healthcareapis/2024-03-31/workspaceprivateendpointconnections"
	"github.com/redtenant/go-azure-sdk/resource-manager/healthcareapis/2024-03-31/workspaceprivatelinkresources"
	"github.com/redtenant/go-azure-sdk/resource-manager/healthcareapis/2024-03-31/workspaces"
	"github.com/redtenant/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

type Client struct {
	Collection                          *collection.CollectionClient
	DicomServices                       *dicomservices.DicomServicesClient
	FhirServices                        *fhirservices.FhirServicesClient
	IotConnectors                       *iotconnectors.IotConnectorsClient
	PrivateEndpointConnections          *privateendpointconnections.PrivateEndpointConnectionsClient
	PrivateLinkResources                *privatelinkresources.PrivateLinkResourcesClient
	Proxy                               *proxy.ProxyClient
	Resource                            *resource.ResourceClient
	WorkspacePrivateEndpointConnections *workspaceprivateendpointconnections.WorkspacePrivateEndpointConnectionsClient
	WorkspacePrivateLinkResources       *workspaceprivatelinkresources.WorkspacePrivateLinkResourcesClient
	Workspaces                          *workspaces.WorkspacesClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	collectionClient, err := collection.NewCollectionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Collection client: %+v", err)
	}
	configureFunc(collectionClient.Client)

	dicomServicesClient, err := dicomservices.NewDicomServicesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DicomServices client: %+v", err)
	}
	configureFunc(dicomServicesClient.Client)

	fhirServicesClient, err := fhirservices.NewFhirServicesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building FhirServices client: %+v", err)
	}
	configureFunc(fhirServicesClient.Client)

	iotConnectorsClient, err := iotconnectors.NewIotConnectorsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building IotConnectors client: %+v", err)
	}
	configureFunc(iotConnectorsClient.Client)

	privateEndpointConnectionsClient, err := privateendpointconnections.NewPrivateEndpointConnectionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PrivateEndpointConnections client: %+v", err)
	}
	configureFunc(privateEndpointConnectionsClient.Client)

	privateLinkResourcesClient, err := privatelinkresources.NewPrivateLinkResourcesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PrivateLinkResources client: %+v", err)
	}
	configureFunc(privateLinkResourcesClient.Client)

	proxyClient, err := proxy.NewProxyClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Proxy client: %+v", err)
	}
	configureFunc(proxyClient.Client)

	resourceClient, err := resource.NewResourceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Resource client: %+v", err)
	}
	configureFunc(resourceClient.Client)

	workspacePrivateEndpointConnectionsClient, err := workspaceprivateendpointconnections.NewWorkspacePrivateEndpointConnectionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building WorkspacePrivateEndpointConnections client: %+v", err)
	}
	configureFunc(workspacePrivateEndpointConnectionsClient.Client)

	workspacePrivateLinkResourcesClient, err := workspaceprivatelinkresources.NewWorkspacePrivateLinkResourcesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building WorkspacePrivateLinkResources client: %+v", err)
	}
	configureFunc(workspacePrivateLinkResourcesClient.Client)

	workspacesClient, err := workspaces.NewWorkspacesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Workspaces client: %+v", err)
	}
	configureFunc(workspacesClient.Client)

	return &Client{
		Collection:                          collectionClient,
		DicomServices:                       dicomServicesClient,
		FhirServices:                        fhirServicesClient,
		IotConnectors:                       iotConnectorsClient,
		PrivateEndpointConnections:          privateEndpointConnectionsClient,
		PrivateLinkResources:                privateLinkResourcesClient,
		Proxy:                               proxyClient,
		Resource:                            resourceClient,
		WorkspacePrivateEndpointConnections: workspacePrivateEndpointConnectionsClient,
		WorkspacePrivateLinkResources:       workspacePrivateLinkResourcesClient,
		Workspaces:                          workspacesClient,
	}, nil
}
