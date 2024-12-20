package v2023_09_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/resource-manager/mobilenetwork/2023-09-01/attacheddatanetwork"
	"github.com/redtenant/go-azure-sdk/resource-manager/mobilenetwork/2023-09-01/attacheddatanetworks"
	"github.com/redtenant/go-azure-sdk/resource-manager/mobilenetwork/2023-09-01/datanetwork"
	"github.com/redtenant/go-azure-sdk/resource-manager/mobilenetwork/2023-09-01/datanetworks"
	"github.com/redtenant/go-azure-sdk/resource-manager/mobilenetwork/2023-09-01/diagnosticspackages"
	"github.com/redtenant/go-azure-sdk/resource-manager/mobilenetwork/2023-09-01/mobilenetwork"
	"github.com/redtenant/go-azure-sdk/resource-manager/mobilenetwork/2023-09-01/mobilenetworks"
	"github.com/redtenant/go-azure-sdk/resource-manager/mobilenetwork/2023-09-01/packetcaptures"
	"github.com/redtenant/go-azure-sdk/resource-manager/mobilenetwork/2023-09-01/packetcorecontrolplane"
	"github.com/redtenant/go-azure-sdk/resource-manager/mobilenetwork/2023-09-01/packetcorecontrolplanecollectdiagnosticspackage"
	"github.com/redtenant/go-azure-sdk/resource-manager/mobilenetwork/2023-09-01/packetcorecontrolplanereinstall"
	"github.com/redtenant/go-azure-sdk/resource-manager/mobilenetwork/2023-09-01/packetcorecontrolplanerollback"
	"github.com/redtenant/go-azure-sdk/resource-manager/mobilenetwork/2023-09-01/packetcorecontrolplanes"
	"github.com/redtenant/go-azure-sdk/resource-manager/mobilenetwork/2023-09-01/packetcorecontrolplaneversion"
	"github.com/redtenant/go-azure-sdk/resource-manager/mobilenetwork/2023-09-01/packetcoredataplane"
	"github.com/redtenant/go-azure-sdk/resource-manager/mobilenetwork/2023-09-01/packetcoredataplanes"
	"github.com/redtenant/go-azure-sdk/resource-manager/mobilenetwork/2023-09-01/service"
	"github.com/redtenant/go-azure-sdk/resource-manager/mobilenetwork/2023-09-01/services"
	"github.com/redtenant/go-azure-sdk/resource-manager/mobilenetwork/2023-09-01/sim"
	"github.com/redtenant/go-azure-sdk/resource-manager/mobilenetwork/2023-09-01/simgroup"
	"github.com/redtenant/go-azure-sdk/resource-manager/mobilenetwork/2023-09-01/simgroups"
	"github.com/redtenant/go-azure-sdk/resource-manager/mobilenetwork/2023-09-01/simpolicies"
	"github.com/redtenant/go-azure-sdk/resource-manager/mobilenetwork/2023-09-01/simpolicy"
	"github.com/redtenant/go-azure-sdk/resource-manager/mobilenetwork/2023-09-01/sims"
	"github.com/redtenant/go-azure-sdk/resource-manager/mobilenetwork/2023-09-01/site"
	"github.com/redtenant/go-azure-sdk/resource-manager/mobilenetwork/2023-09-01/sites"
	"github.com/redtenant/go-azure-sdk/resource-manager/mobilenetwork/2023-09-01/slice"
	"github.com/redtenant/go-azure-sdk/resource-manager/mobilenetwork/2023-09-01/slices"
	"github.com/redtenant/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

type Client struct {
	AttachedDataNetwork                             *attacheddatanetwork.AttachedDataNetworkClient
	AttachedDataNetworks                            *attacheddatanetworks.AttachedDataNetworksClient
	DataNetwork                                     *datanetwork.DataNetworkClient
	DataNetworks                                    *datanetworks.DataNetworksClient
	DiagnosticsPackages                             *diagnosticspackages.DiagnosticsPackagesClient
	MobileNetwork                                   *mobilenetwork.MobileNetworkClient
	MobileNetworks                                  *mobilenetworks.MobileNetworksClient
	PacketCaptures                                  *packetcaptures.PacketCapturesClient
	PacketCoreControlPlane                          *packetcorecontrolplane.PacketCoreControlPlaneClient
	PacketCoreControlPlaneCollectDiagnosticsPackage *packetcorecontrolplanecollectdiagnosticspackage.PacketCoreControlPlaneCollectDiagnosticsPackageClient
	PacketCoreControlPlaneReinstall                 *packetcorecontrolplanereinstall.PacketCoreControlPlaneReinstallClient
	PacketCoreControlPlaneRollback                  *packetcorecontrolplanerollback.PacketCoreControlPlaneRollbackClient
	PacketCoreControlPlaneVersion                   *packetcorecontrolplaneversion.PacketCoreControlPlaneVersionClient
	PacketCoreControlPlanes                         *packetcorecontrolplanes.PacketCoreControlPlanesClient
	PacketCoreDataPlane                             *packetcoredataplane.PacketCoreDataPlaneClient
	PacketCoreDataPlanes                            *packetcoredataplanes.PacketCoreDataPlanesClient
	SIM                                             *sim.SIMClient
	SIMGroup                                        *simgroup.SIMGroupClient
	SIMGroups                                       *simgroups.SIMGroupsClient
	SIMPolicies                                     *simpolicies.SIMPoliciesClient
	SIMPolicy                                       *simpolicy.SIMPolicyClient
	SIMs                                            *sims.SIMsClient
	Service                                         *service.ServiceClient
	Services                                        *services.ServicesClient
	Site                                            *site.SiteClient
	Sites                                           *sites.SitesClient
	Slice                                           *slice.SliceClient
	Slices                                          *slices.SlicesClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	attachedDataNetworkClient, err := attacheddatanetwork.NewAttachedDataNetworkClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AttachedDataNetwork client: %+v", err)
	}
	configureFunc(attachedDataNetworkClient.Client)

	attachedDataNetworksClient, err := attacheddatanetworks.NewAttachedDataNetworksClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AttachedDataNetworks client: %+v", err)
	}
	configureFunc(attachedDataNetworksClient.Client)

	dataNetworkClient, err := datanetwork.NewDataNetworkClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DataNetwork client: %+v", err)
	}
	configureFunc(dataNetworkClient.Client)

	dataNetworksClient, err := datanetworks.NewDataNetworksClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DataNetworks client: %+v", err)
	}
	configureFunc(dataNetworksClient.Client)

	diagnosticsPackagesClient, err := diagnosticspackages.NewDiagnosticsPackagesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DiagnosticsPackages client: %+v", err)
	}
	configureFunc(diagnosticsPackagesClient.Client)

	mobileNetworkClient, err := mobilenetwork.NewMobileNetworkClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MobileNetwork client: %+v", err)
	}
	configureFunc(mobileNetworkClient.Client)

	mobileNetworksClient, err := mobilenetworks.NewMobileNetworksClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MobileNetworks client: %+v", err)
	}
	configureFunc(mobileNetworksClient.Client)

	packetCapturesClient, err := packetcaptures.NewPacketCapturesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PacketCaptures client: %+v", err)
	}
	configureFunc(packetCapturesClient.Client)

	packetCoreControlPlaneClient, err := packetcorecontrolplane.NewPacketCoreControlPlaneClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PacketCoreControlPlane client: %+v", err)
	}
	configureFunc(packetCoreControlPlaneClient.Client)

	packetCoreControlPlaneCollectDiagnosticsPackageClient, err := packetcorecontrolplanecollectdiagnosticspackage.NewPacketCoreControlPlaneCollectDiagnosticsPackageClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PacketCoreControlPlaneCollectDiagnosticsPackage client: %+v", err)
	}
	configureFunc(packetCoreControlPlaneCollectDiagnosticsPackageClient.Client)

	packetCoreControlPlaneReinstallClient, err := packetcorecontrolplanereinstall.NewPacketCoreControlPlaneReinstallClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PacketCoreControlPlaneReinstall client: %+v", err)
	}
	configureFunc(packetCoreControlPlaneReinstallClient.Client)

	packetCoreControlPlaneRollbackClient, err := packetcorecontrolplanerollback.NewPacketCoreControlPlaneRollbackClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PacketCoreControlPlaneRollback client: %+v", err)
	}
	configureFunc(packetCoreControlPlaneRollbackClient.Client)

	packetCoreControlPlaneVersionClient, err := packetcorecontrolplaneversion.NewPacketCoreControlPlaneVersionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PacketCoreControlPlaneVersion client: %+v", err)
	}
	configureFunc(packetCoreControlPlaneVersionClient.Client)

	packetCoreControlPlanesClient, err := packetcorecontrolplanes.NewPacketCoreControlPlanesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PacketCoreControlPlanes client: %+v", err)
	}
	configureFunc(packetCoreControlPlanesClient.Client)

	packetCoreDataPlaneClient, err := packetcoredataplane.NewPacketCoreDataPlaneClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PacketCoreDataPlane client: %+v", err)
	}
	configureFunc(packetCoreDataPlaneClient.Client)

	packetCoreDataPlanesClient, err := packetcoredataplanes.NewPacketCoreDataPlanesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PacketCoreDataPlanes client: %+v", err)
	}
	configureFunc(packetCoreDataPlanesClient.Client)

	sIMClient, err := sim.NewSIMClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SIM client: %+v", err)
	}
	configureFunc(sIMClient.Client)

	sIMGroupClient, err := simgroup.NewSIMGroupClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SIMGroup client: %+v", err)
	}
	configureFunc(sIMGroupClient.Client)

	sIMGroupsClient, err := simgroups.NewSIMGroupsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SIMGroups client: %+v", err)
	}
	configureFunc(sIMGroupsClient.Client)

	sIMPoliciesClient, err := simpolicies.NewSIMPoliciesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SIMPolicies client: %+v", err)
	}
	configureFunc(sIMPoliciesClient.Client)

	sIMPolicyClient, err := simpolicy.NewSIMPolicyClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SIMPolicy client: %+v", err)
	}
	configureFunc(sIMPolicyClient.Client)

	sIMsClient, err := sims.NewSIMsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SIMs client: %+v", err)
	}
	configureFunc(sIMsClient.Client)

	serviceClient, err := service.NewServiceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Service client: %+v", err)
	}
	configureFunc(serviceClient.Client)

	servicesClient, err := services.NewServicesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Services client: %+v", err)
	}
	configureFunc(servicesClient.Client)

	siteClient, err := site.NewSiteClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Site client: %+v", err)
	}
	configureFunc(siteClient.Client)

	sitesClient, err := sites.NewSitesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Sites client: %+v", err)
	}
	configureFunc(sitesClient.Client)

	sliceClient, err := slice.NewSliceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Slice client: %+v", err)
	}
	configureFunc(sliceClient.Client)

	slicesClient, err := slices.NewSlicesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Slices client: %+v", err)
	}
	configureFunc(slicesClient.Client)

	return &Client{
		AttachedDataNetwork:    attachedDataNetworkClient,
		AttachedDataNetworks:   attachedDataNetworksClient,
		DataNetwork:            dataNetworkClient,
		DataNetworks:           dataNetworksClient,
		DiagnosticsPackages:    diagnosticsPackagesClient,
		MobileNetwork:          mobileNetworkClient,
		MobileNetworks:         mobileNetworksClient,
		PacketCaptures:         packetCapturesClient,
		PacketCoreControlPlane: packetCoreControlPlaneClient,
		PacketCoreControlPlaneCollectDiagnosticsPackage: packetCoreControlPlaneCollectDiagnosticsPackageClient,
		PacketCoreControlPlaneReinstall:                 packetCoreControlPlaneReinstallClient,
		PacketCoreControlPlaneRollback:                  packetCoreControlPlaneRollbackClient,
		PacketCoreControlPlaneVersion:                   packetCoreControlPlaneVersionClient,
		PacketCoreControlPlanes:                         packetCoreControlPlanesClient,
		PacketCoreDataPlane:                             packetCoreDataPlaneClient,
		PacketCoreDataPlanes:                            packetCoreDataPlanesClient,
		SIM:                                             sIMClient,
		SIMGroup:                                        sIMGroupClient,
		SIMGroups:                                       sIMGroupsClient,
		SIMPolicies:                                     sIMPoliciesClient,
		SIMPolicy:                                       sIMPolicyClient,
		SIMs:                                            sIMsClient,
		Service:                                         serviceClient,
		Services:                                        servicesClient,
		Site:                                            siteClient,
		Sites:                                           sitesClient,
		Slice:                                           sliceClient,
		Slices:                                          slicesClient,
	}, nil
}
