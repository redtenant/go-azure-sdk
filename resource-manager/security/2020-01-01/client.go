package v2020_01_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/resource-manager/security/2020-01-01/alerts"
	"github.com/redtenant/go-azure-sdk/resource-manager/security/2020-01-01/allowedconnections"
	"github.com/redtenant/go-azure-sdk/resource-manager/security/2020-01-01/assessments"
	"github.com/redtenant/go-azure-sdk/resource-manager/security/2020-01-01/assessmentsmetadata"
	"github.com/redtenant/go-azure-sdk/resource-manager/security/2020-01-01/discoveredsecuritysolutions"
	"github.com/redtenant/go-azure-sdk/resource-manager/security/2020-01-01/externalsecuritysolutions"
	"github.com/redtenant/go-azure-sdk/resource-manager/security/2020-01-01/jitnetworkaccesspolicies"
	"github.com/redtenant/go-azure-sdk/resource-manager/security/2020-01-01/securescore"
	"github.com/redtenant/go-azure-sdk/resource-manager/security/2020-01-01/securescorecontroldefinitions"
	"github.com/redtenant/go-azure-sdk/resource-manager/security/2020-01-01/securescorecontrols"
	"github.com/redtenant/go-azure-sdk/resource-manager/security/2020-01-01/securitysolutions"
	"github.com/redtenant/go-azure-sdk/resource-manager/security/2020-01-01/securitysolutionsreferencedata"
	"github.com/redtenant/go-azure-sdk/resource-manager/security/2020-01-01/servervulnerabilityassessment"
	"github.com/redtenant/go-azure-sdk/resource-manager/security/2020-01-01/servervulnerabilityassessments"
	"github.com/redtenant/go-azure-sdk/resource-manager/security/2020-01-01/topology"
	"github.com/redtenant/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

type Client struct {
	Alerts                         *alerts.AlertsClient
	AllowedConnections             *allowedconnections.AllowedConnectionsClient
	Assessments                    *assessments.AssessmentsClient
	AssessmentsMetadata            *assessmentsmetadata.AssessmentsMetadataClient
	DiscoveredSecuritySolutions    *discoveredsecuritysolutions.DiscoveredSecuritySolutionsClient
	ExternalSecuritySolutions      *externalsecuritysolutions.ExternalSecuritySolutionsClient
	JitNetworkAccessPolicies       *jitnetworkaccesspolicies.JitNetworkAccessPoliciesClient
	SecureScore                    *securescore.SecureScoreClient
	SecureScoreControlDefinitions  *securescorecontroldefinitions.SecureScoreControlDefinitionsClient
	SecureScoreControls            *securescorecontrols.SecureScoreControlsClient
	SecuritySolutions              *securitysolutions.SecuritySolutionsClient
	SecuritySolutionsReferenceData *securitysolutionsreferencedata.SecuritySolutionsReferenceDataClient
	ServerVulnerabilityAssessment  *servervulnerabilityassessment.ServerVulnerabilityAssessmentClient
	ServerVulnerabilityAssessments *servervulnerabilityassessments.ServerVulnerabilityAssessmentsClient
	Topology                       *topology.TopologyClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	alertsClient, err := alerts.NewAlertsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Alerts client: %+v", err)
	}
	configureFunc(alertsClient.Client)

	allowedConnectionsClient, err := allowedconnections.NewAllowedConnectionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AllowedConnections client: %+v", err)
	}
	configureFunc(allowedConnectionsClient.Client)

	assessmentsClient, err := assessments.NewAssessmentsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Assessments client: %+v", err)
	}
	configureFunc(assessmentsClient.Client)

	assessmentsMetadataClient, err := assessmentsmetadata.NewAssessmentsMetadataClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AssessmentsMetadata client: %+v", err)
	}
	configureFunc(assessmentsMetadataClient.Client)

	discoveredSecuritySolutionsClient, err := discoveredsecuritysolutions.NewDiscoveredSecuritySolutionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DiscoveredSecuritySolutions client: %+v", err)
	}
	configureFunc(discoveredSecuritySolutionsClient.Client)

	externalSecuritySolutionsClient, err := externalsecuritysolutions.NewExternalSecuritySolutionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ExternalSecuritySolutions client: %+v", err)
	}
	configureFunc(externalSecuritySolutionsClient.Client)

	jitNetworkAccessPoliciesClient, err := jitnetworkaccesspolicies.NewJitNetworkAccessPoliciesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building JitNetworkAccessPolicies client: %+v", err)
	}
	configureFunc(jitNetworkAccessPoliciesClient.Client)

	secureScoreClient, err := securescore.NewSecureScoreClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SecureScore client: %+v", err)
	}
	configureFunc(secureScoreClient.Client)

	secureScoreControlDefinitionsClient, err := securescorecontroldefinitions.NewSecureScoreControlDefinitionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SecureScoreControlDefinitions client: %+v", err)
	}
	configureFunc(secureScoreControlDefinitionsClient.Client)

	secureScoreControlsClient, err := securescorecontrols.NewSecureScoreControlsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SecureScoreControls client: %+v", err)
	}
	configureFunc(secureScoreControlsClient.Client)

	securitySolutionsClient, err := securitysolutions.NewSecuritySolutionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SecuritySolutions client: %+v", err)
	}
	configureFunc(securitySolutionsClient.Client)

	securitySolutionsReferenceDataClient, err := securitysolutionsreferencedata.NewSecuritySolutionsReferenceDataClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SecuritySolutionsReferenceData client: %+v", err)
	}
	configureFunc(securitySolutionsReferenceDataClient.Client)

	serverVulnerabilityAssessmentClient, err := servervulnerabilityassessment.NewServerVulnerabilityAssessmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ServerVulnerabilityAssessment client: %+v", err)
	}
	configureFunc(serverVulnerabilityAssessmentClient.Client)

	serverVulnerabilityAssessmentsClient, err := servervulnerabilityassessments.NewServerVulnerabilityAssessmentsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ServerVulnerabilityAssessments client: %+v", err)
	}
	configureFunc(serverVulnerabilityAssessmentsClient.Client)

	topologyClient, err := topology.NewTopologyClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Topology client: %+v", err)
	}
	configureFunc(topologyClient.Client)

	return &Client{
		Alerts:                         alertsClient,
		AllowedConnections:             allowedConnectionsClient,
		Assessments:                    assessmentsClient,
		AssessmentsMetadata:            assessmentsMetadataClient,
		DiscoveredSecuritySolutions:    discoveredSecuritySolutionsClient,
		ExternalSecuritySolutions:      externalSecuritySolutionsClient,
		JitNetworkAccessPolicies:       jitNetworkAccessPoliciesClient,
		SecureScore:                    secureScoreClient,
		SecureScoreControlDefinitions:  secureScoreControlDefinitionsClient,
		SecureScoreControls:            secureScoreControlsClient,
		SecuritySolutions:              securitySolutionsClient,
		SecuritySolutionsReferenceData: securitySolutionsReferenceDataClient,
		ServerVulnerabilityAssessment:  serverVulnerabilityAssessmentClient,
		ServerVulnerabilityAssessments: serverVulnerabilityAssessmentsClient,
		Topology:                       topologyClient,
	}, nil
}
