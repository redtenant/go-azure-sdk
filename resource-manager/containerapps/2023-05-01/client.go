package v2023_05_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/resource-manager/containerapps/2023-05-01/availableworkloadprofiles"
	"github.com/redtenant/go-azure-sdk/resource-manager/containerapps/2023-05-01/billingmeters"
	"github.com/redtenant/go-azure-sdk/resource-manager/containerapps/2023-05-01/certificates"
	"github.com/redtenant/go-azure-sdk/resource-manager/containerapps/2023-05-01/connectedenvironments"
	"github.com/redtenant/go-azure-sdk/resource-manager/containerapps/2023-05-01/connectedenvironmentsstorages"
	"github.com/redtenant/go-azure-sdk/resource-manager/containerapps/2023-05-01/containerapps"
	"github.com/redtenant/go-azure-sdk/resource-manager/containerapps/2023-05-01/containerappsauthconfigs"
	"github.com/redtenant/go-azure-sdk/resource-manager/containerapps/2023-05-01/containerappsrevisionreplicas"
	"github.com/redtenant/go-azure-sdk/resource-manager/containerapps/2023-05-01/containerappsrevisions"
	"github.com/redtenant/go-azure-sdk/resource-manager/containerapps/2023-05-01/containerappssourcecontrols"
	"github.com/redtenant/go-azure-sdk/resource-manager/containerapps/2023-05-01/daprcomponents"
	"github.com/redtenant/go-azure-sdk/resource-manager/containerapps/2023-05-01/diagnostics"
	"github.com/redtenant/go-azure-sdk/resource-manager/containerapps/2023-05-01/jobs"
	"github.com/redtenant/go-azure-sdk/resource-manager/containerapps/2023-05-01/managedcertificates"
	"github.com/redtenant/go-azure-sdk/resource-manager/containerapps/2023-05-01/managedenvironments"
	"github.com/redtenant/go-azure-sdk/resource-manager/containerapps/2023-05-01/managedenvironmentsstorages"
	"github.com/redtenant/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

type Client struct {
	AvailableWorkloadProfiles     *availableworkloadprofiles.AvailableWorkloadProfilesClient
	BillingMeters                 *billingmeters.BillingMetersClient
	Certificates                  *certificates.CertificatesClient
	ConnectedEnvironments         *connectedenvironments.ConnectedEnvironmentsClient
	ConnectedEnvironmentsStorages *connectedenvironmentsstorages.ConnectedEnvironmentsStoragesClient
	ContainerApps                 *containerapps.ContainerAppsClient
	ContainerAppsAuthConfigs      *containerappsauthconfigs.ContainerAppsAuthConfigsClient
	ContainerAppsRevisionReplicas *containerappsrevisionreplicas.ContainerAppsRevisionReplicasClient
	ContainerAppsRevisions        *containerappsrevisions.ContainerAppsRevisionsClient
	ContainerAppsSourceControls   *containerappssourcecontrols.ContainerAppsSourceControlsClient
	DaprComponents                *daprcomponents.DaprComponentsClient
	Diagnostics                   *diagnostics.DiagnosticsClient
	Jobs                          *jobs.JobsClient
	ManagedCertificates           *managedcertificates.ManagedCertificatesClient
	ManagedEnvironments           *managedenvironments.ManagedEnvironmentsClient
	ManagedEnvironmentsStorages   *managedenvironmentsstorages.ManagedEnvironmentsStoragesClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	availableWorkloadProfilesClient, err := availableworkloadprofiles.NewAvailableWorkloadProfilesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AvailableWorkloadProfiles client: %+v", err)
	}
	configureFunc(availableWorkloadProfilesClient.Client)

	billingMetersClient, err := billingmeters.NewBillingMetersClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building BillingMeters client: %+v", err)
	}
	configureFunc(billingMetersClient.Client)

	certificatesClient, err := certificates.NewCertificatesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Certificates client: %+v", err)
	}
	configureFunc(certificatesClient.Client)

	connectedEnvironmentsClient, err := connectedenvironments.NewConnectedEnvironmentsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ConnectedEnvironments client: %+v", err)
	}
	configureFunc(connectedEnvironmentsClient.Client)

	connectedEnvironmentsStoragesClient, err := connectedenvironmentsstorages.NewConnectedEnvironmentsStoragesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ConnectedEnvironmentsStorages client: %+v", err)
	}
	configureFunc(connectedEnvironmentsStoragesClient.Client)

	containerAppsAuthConfigsClient, err := containerappsauthconfigs.NewContainerAppsAuthConfigsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ContainerAppsAuthConfigs client: %+v", err)
	}
	configureFunc(containerAppsAuthConfigsClient.Client)

	containerAppsClient, err := containerapps.NewContainerAppsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ContainerApps client: %+v", err)
	}
	configureFunc(containerAppsClient.Client)

	containerAppsRevisionReplicasClient, err := containerappsrevisionreplicas.NewContainerAppsRevisionReplicasClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ContainerAppsRevisionReplicas client: %+v", err)
	}
	configureFunc(containerAppsRevisionReplicasClient.Client)

	containerAppsRevisionsClient, err := containerappsrevisions.NewContainerAppsRevisionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ContainerAppsRevisions client: %+v", err)
	}
	configureFunc(containerAppsRevisionsClient.Client)

	containerAppsSourceControlsClient, err := containerappssourcecontrols.NewContainerAppsSourceControlsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ContainerAppsSourceControls client: %+v", err)
	}
	configureFunc(containerAppsSourceControlsClient.Client)

	daprComponentsClient, err := daprcomponents.NewDaprComponentsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DaprComponents client: %+v", err)
	}
	configureFunc(daprComponentsClient.Client)

	diagnosticsClient, err := diagnostics.NewDiagnosticsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Diagnostics client: %+v", err)
	}
	configureFunc(diagnosticsClient.Client)

	jobsClient, err := jobs.NewJobsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Jobs client: %+v", err)
	}
	configureFunc(jobsClient.Client)

	managedCertificatesClient, err := managedcertificates.NewManagedCertificatesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ManagedCertificates client: %+v", err)
	}
	configureFunc(managedCertificatesClient.Client)

	managedEnvironmentsClient, err := managedenvironments.NewManagedEnvironmentsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ManagedEnvironments client: %+v", err)
	}
	configureFunc(managedEnvironmentsClient.Client)

	managedEnvironmentsStoragesClient, err := managedenvironmentsstorages.NewManagedEnvironmentsStoragesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ManagedEnvironmentsStorages client: %+v", err)
	}
	configureFunc(managedEnvironmentsStoragesClient.Client)

	return &Client{
		AvailableWorkloadProfiles:     availableWorkloadProfilesClient,
		BillingMeters:                 billingMetersClient,
		Certificates:                  certificatesClient,
		ConnectedEnvironments:         connectedEnvironmentsClient,
		ConnectedEnvironmentsStorages: connectedEnvironmentsStoragesClient,
		ContainerApps:                 containerAppsClient,
		ContainerAppsAuthConfigs:      containerAppsAuthConfigsClient,
		ContainerAppsRevisionReplicas: containerAppsRevisionReplicasClient,
		ContainerAppsRevisions:        containerAppsRevisionsClient,
		ContainerAppsSourceControls:   containerAppsSourceControlsClient,
		DaprComponents:                daprComponentsClient,
		Diagnostics:                   diagnosticsClient,
		Jobs:                          jobsClient,
		ManagedCertificates:           managedCertificatesClient,
		ManagedEnvironments:           managedEnvironmentsClient,
		ManagedEnvironmentsStorages:   managedEnvironmentsStoragesClient,
	}, nil
}
