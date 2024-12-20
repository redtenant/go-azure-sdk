package v2019_05_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/resource-manager/logic/2019-05-01/integrationaccountagreements"
	"github.com/redtenant/go-azure-sdk/resource-manager/logic/2019-05-01/integrationaccountassemblies"
	"github.com/redtenant/go-azure-sdk/resource-manager/logic/2019-05-01/integrationaccountbatchconfigurations"
	"github.com/redtenant/go-azure-sdk/resource-manager/logic/2019-05-01/integrationaccountcertificates"
	"github.com/redtenant/go-azure-sdk/resource-manager/logic/2019-05-01/integrationaccountmaps"
	"github.com/redtenant/go-azure-sdk/resource-manager/logic/2019-05-01/integrationaccountpartners"
	"github.com/redtenant/go-azure-sdk/resource-manager/logic/2019-05-01/integrationaccounts"
	"github.com/redtenant/go-azure-sdk/resource-manager/logic/2019-05-01/integrationaccountschemas"
	"github.com/redtenant/go-azure-sdk/resource-manager/logic/2019-05-01/integrationaccountsessions"
	"github.com/redtenant/go-azure-sdk/resource-manager/logic/2019-05-01/integrationserviceenvironmentmanagedapi"
	"github.com/redtenant/go-azure-sdk/resource-manager/logic/2019-05-01/integrationserviceenvironmentmanagedapis"
	"github.com/redtenant/go-azure-sdk/resource-manager/logic/2019-05-01/integrationserviceenvironmentnetworkhealth"
	"github.com/redtenant/go-azure-sdk/resource-manager/logic/2019-05-01/integrationserviceenvironmentrestart"
	"github.com/redtenant/go-azure-sdk/resource-manager/logic/2019-05-01/integrationserviceenvironments"
	"github.com/redtenant/go-azure-sdk/resource-manager/logic/2019-05-01/integrationserviceenvironmentskus"
	"github.com/redtenant/go-azure-sdk/resource-manager/logic/2019-05-01/workflowrunactions"
	"github.com/redtenant/go-azure-sdk/resource-manager/logic/2019-05-01/workflowrunoperations"
	"github.com/redtenant/go-azure-sdk/resource-manager/logic/2019-05-01/workflowruns"
	"github.com/redtenant/go-azure-sdk/resource-manager/logic/2019-05-01/workflows"
	"github.com/redtenant/go-azure-sdk/resource-manager/logic/2019-05-01/workflowtriggerhistories"
	"github.com/redtenant/go-azure-sdk/resource-manager/logic/2019-05-01/workflowtriggers"
	"github.com/redtenant/go-azure-sdk/resource-manager/logic/2019-05-01/workflowversions"
	"github.com/redtenant/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

type Client struct {
	IntegrationAccountAgreements               *integrationaccountagreements.IntegrationAccountAgreementsClient
	IntegrationAccountAssemblies               *integrationaccountassemblies.IntegrationAccountAssembliesClient
	IntegrationAccountBatchConfigurations      *integrationaccountbatchconfigurations.IntegrationAccountBatchConfigurationsClient
	IntegrationAccountCertificates             *integrationaccountcertificates.IntegrationAccountCertificatesClient
	IntegrationAccountMaps                     *integrationaccountmaps.IntegrationAccountMapsClient
	IntegrationAccountPartners                 *integrationaccountpartners.IntegrationAccountPartnersClient
	IntegrationAccountSchemas                  *integrationaccountschemas.IntegrationAccountSchemasClient
	IntegrationAccountSessions                 *integrationaccountsessions.IntegrationAccountSessionsClient
	IntegrationAccounts                        *integrationaccounts.IntegrationAccountsClient
	IntegrationServiceEnvironmentManagedApi    *integrationserviceenvironmentmanagedapi.IntegrationServiceEnvironmentManagedApiClient
	IntegrationServiceEnvironmentManagedApis   *integrationserviceenvironmentmanagedapis.IntegrationServiceEnvironmentManagedApisClient
	IntegrationServiceEnvironmentNetworkHealth *integrationserviceenvironmentnetworkhealth.IntegrationServiceEnvironmentNetworkHealthClient
	IntegrationServiceEnvironmentRestart       *integrationserviceenvironmentrestart.IntegrationServiceEnvironmentRestartClient
	IntegrationServiceEnvironmentSkus          *integrationserviceenvironmentskus.IntegrationServiceEnvironmentSkusClient
	IntegrationServiceEnvironments             *integrationserviceenvironments.IntegrationServiceEnvironmentsClient
	WorkflowRunActions                         *workflowrunactions.WorkflowRunActionsClient
	WorkflowRunOperations                      *workflowrunoperations.WorkflowRunOperationsClient
	WorkflowRuns                               *workflowruns.WorkflowRunsClient
	WorkflowTriggerHistories                   *workflowtriggerhistories.WorkflowTriggerHistoriesClient
	WorkflowTriggers                           *workflowtriggers.WorkflowTriggersClient
	WorkflowVersions                           *workflowversions.WorkflowVersionsClient
	Workflows                                  *workflows.WorkflowsClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	integrationAccountAgreementsClient, err := integrationaccountagreements.NewIntegrationAccountAgreementsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building IntegrationAccountAgreements client: %+v", err)
	}
	configureFunc(integrationAccountAgreementsClient.Client)

	integrationAccountAssembliesClient, err := integrationaccountassemblies.NewIntegrationAccountAssembliesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building IntegrationAccountAssemblies client: %+v", err)
	}
	configureFunc(integrationAccountAssembliesClient.Client)

	integrationAccountBatchConfigurationsClient, err := integrationaccountbatchconfigurations.NewIntegrationAccountBatchConfigurationsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building IntegrationAccountBatchConfigurations client: %+v", err)
	}
	configureFunc(integrationAccountBatchConfigurationsClient.Client)

	integrationAccountCertificatesClient, err := integrationaccountcertificates.NewIntegrationAccountCertificatesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building IntegrationAccountCertificates client: %+v", err)
	}
	configureFunc(integrationAccountCertificatesClient.Client)

	integrationAccountMapsClient, err := integrationaccountmaps.NewIntegrationAccountMapsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building IntegrationAccountMaps client: %+v", err)
	}
	configureFunc(integrationAccountMapsClient.Client)

	integrationAccountPartnersClient, err := integrationaccountpartners.NewIntegrationAccountPartnersClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building IntegrationAccountPartners client: %+v", err)
	}
	configureFunc(integrationAccountPartnersClient.Client)

	integrationAccountSchemasClient, err := integrationaccountschemas.NewIntegrationAccountSchemasClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building IntegrationAccountSchemas client: %+v", err)
	}
	configureFunc(integrationAccountSchemasClient.Client)

	integrationAccountSessionsClient, err := integrationaccountsessions.NewIntegrationAccountSessionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building IntegrationAccountSessions client: %+v", err)
	}
	configureFunc(integrationAccountSessionsClient.Client)

	integrationAccountsClient, err := integrationaccounts.NewIntegrationAccountsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building IntegrationAccounts client: %+v", err)
	}
	configureFunc(integrationAccountsClient.Client)

	integrationServiceEnvironmentManagedApiClient, err := integrationserviceenvironmentmanagedapi.NewIntegrationServiceEnvironmentManagedApiClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building IntegrationServiceEnvironmentManagedApi client: %+v", err)
	}
	configureFunc(integrationServiceEnvironmentManagedApiClient.Client)

	integrationServiceEnvironmentManagedApisClient, err := integrationserviceenvironmentmanagedapis.NewIntegrationServiceEnvironmentManagedApisClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building IntegrationServiceEnvironmentManagedApis client: %+v", err)
	}
	configureFunc(integrationServiceEnvironmentManagedApisClient.Client)

	integrationServiceEnvironmentNetworkHealthClient, err := integrationserviceenvironmentnetworkhealth.NewIntegrationServiceEnvironmentNetworkHealthClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building IntegrationServiceEnvironmentNetworkHealth client: %+v", err)
	}
	configureFunc(integrationServiceEnvironmentNetworkHealthClient.Client)

	integrationServiceEnvironmentRestartClient, err := integrationserviceenvironmentrestart.NewIntegrationServiceEnvironmentRestartClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building IntegrationServiceEnvironmentRestart client: %+v", err)
	}
	configureFunc(integrationServiceEnvironmentRestartClient.Client)

	integrationServiceEnvironmentSkusClient, err := integrationserviceenvironmentskus.NewIntegrationServiceEnvironmentSkusClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building IntegrationServiceEnvironmentSkus client: %+v", err)
	}
	configureFunc(integrationServiceEnvironmentSkusClient.Client)

	integrationServiceEnvironmentsClient, err := integrationserviceenvironments.NewIntegrationServiceEnvironmentsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building IntegrationServiceEnvironments client: %+v", err)
	}
	configureFunc(integrationServiceEnvironmentsClient.Client)

	workflowRunActionsClient, err := workflowrunactions.NewWorkflowRunActionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building WorkflowRunActions client: %+v", err)
	}
	configureFunc(workflowRunActionsClient.Client)

	workflowRunOperationsClient, err := workflowrunoperations.NewWorkflowRunOperationsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building WorkflowRunOperations client: %+v", err)
	}
	configureFunc(workflowRunOperationsClient.Client)

	workflowRunsClient, err := workflowruns.NewWorkflowRunsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building WorkflowRuns client: %+v", err)
	}
	configureFunc(workflowRunsClient.Client)

	workflowTriggerHistoriesClient, err := workflowtriggerhistories.NewWorkflowTriggerHistoriesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building WorkflowTriggerHistories client: %+v", err)
	}
	configureFunc(workflowTriggerHistoriesClient.Client)

	workflowTriggersClient, err := workflowtriggers.NewWorkflowTriggersClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building WorkflowTriggers client: %+v", err)
	}
	configureFunc(workflowTriggersClient.Client)

	workflowVersionsClient, err := workflowversions.NewWorkflowVersionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building WorkflowVersions client: %+v", err)
	}
	configureFunc(workflowVersionsClient.Client)

	workflowsClient, err := workflows.NewWorkflowsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Workflows client: %+v", err)
	}
	configureFunc(workflowsClient.Client)

	return &Client{
		IntegrationAccountAgreements:               integrationAccountAgreementsClient,
		IntegrationAccountAssemblies:               integrationAccountAssembliesClient,
		IntegrationAccountBatchConfigurations:      integrationAccountBatchConfigurationsClient,
		IntegrationAccountCertificates:             integrationAccountCertificatesClient,
		IntegrationAccountMaps:                     integrationAccountMapsClient,
		IntegrationAccountPartners:                 integrationAccountPartnersClient,
		IntegrationAccountSchemas:                  integrationAccountSchemasClient,
		IntegrationAccountSessions:                 integrationAccountSessionsClient,
		IntegrationAccounts:                        integrationAccountsClient,
		IntegrationServiceEnvironmentManagedApi:    integrationServiceEnvironmentManagedApiClient,
		IntegrationServiceEnvironmentManagedApis:   integrationServiceEnvironmentManagedApisClient,
		IntegrationServiceEnvironmentNetworkHealth: integrationServiceEnvironmentNetworkHealthClient,
		IntegrationServiceEnvironmentRestart:       integrationServiceEnvironmentRestartClient,
		IntegrationServiceEnvironmentSkus:          integrationServiceEnvironmentSkusClient,
		IntegrationServiceEnvironments:             integrationServiceEnvironmentsClient,
		WorkflowRunActions:                         workflowRunActionsClient,
		WorkflowRunOperations:                      workflowRunOperationsClient,
		WorkflowRuns:                               workflowRunsClient,
		WorkflowTriggerHistories:                   workflowTriggerHistoriesClient,
		WorkflowTriggers:                           workflowTriggersClient,
		WorkflowVersions:                           workflowVersionsClient,
		Workflows:                                  workflowsClient,
	}, nil
}
