package siteinformationprotectiondatalosspreventionpolicy

import (
	"github.com/redtenant/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/redtenant/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EvaluateSiteInformationProtectionDataLossPreventionPoliciesRequest struct {
	EvaluationInput  *beta.DlpEvaluationInput `json:"evaluationInput,omitempty"`
	NotificationInfo *beta.DlpNotification    `json:"notificationInfo,omitempty"`
	Target           nullable.Type[string]    `json:"target,omitempty"`
}
