package stable

import (
	"encoding/json"
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ SecurityAlertEvidence = SecurityAmazonResourceEvidence{}

type SecurityAmazonResourceEvidence struct {
	// The unique identifier for the Amazon account.
	AmazonAccountId nullable.Type[string] `json:"amazonAccountId,omitempty"`

	// The Amazon resource identifier (ARN) for the cloud resource.
	AmazonResourceId nullable.Type[string] `json:"amazonResourceId,omitempty"`

	// The name of the resource.
	ResourceName nullable.Type[string] `json:"resourceName,omitempty"`

	// The type of the resource.
	ResourceType nullable.Type[string] `json:"resourceType,omitempty"`

	// Fields inherited from SecurityAlertEvidence

	// The date and time when the evidence was created and added to the alert. The Timestamp type represents date and time
	// information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is
	// 2014-01-01T00:00:00Z.
	CreatedDateTime *string `json:"createdDateTime,omitempty"`

	// Detailed description of the entity role/s in an alert. Values are free-form.
	DetailedRoles *[]string `json:"detailedRoles,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	RemediationStatus *SecurityEvidenceRemediationStatus `json:"remediationStatus,omitempty"`

	// Details about the remediation status.
	RemediationStatusDetails nullable.Type[string] `json:"remediationStatusDetails,omitempty"`

	// The role/s that an evidence entity represents in an alert, for example, an IP address that is associated with an
	// attacker has the evidence role Attacker.
	Roles *[]SecurityEvidenceRole `json:"roles,omitempty"`

	// Array of custom tags associated with an evidence instance, for example, to denote a group of devices, high-value
	// assets, etc.
	Tags *[]string `json:"tags,omitempty"`

	Verdict *SecurityEvidenceVerdict `json:"verdict,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s SecurityAmazonResourceEvidence) SecurityAlertEvidence() BaseSecurityAlertEvidenceImpl {
	return BaseSecurityAlertEvidenceImpl{
		CreatedDateTime:          s.CreatedDateTime,
		DetailedRoles:            s.DetailedRoles,
		ODataId:                  s.ODataId,
		ODataType:                s.ODataType,
		RemediationStatus:        s.RemediationStatus,
		RemediationStatusDetails: s.RemediationStatusDetails,
		Roles:                    s.Roles,
		Tags:                     s.Tags,
		Verdict:                  s.Verdict,
	}
}

var _ json.Marshaler = SecurityAmazonResourceEvidence{}

func (s SecurityAmazonResourceEvidence) MarshalJSON() ([]byte, error) {
	type wrapper SecurityAmazonResourceEvidence
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SecurityAmazonResourceEvidence: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SecurityAmazonResourceEvidence: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.security.amazonResourceEvidence"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SecurityAmazonResourceEvidence: %+v", err)
	}

	return encoded, nil
}
