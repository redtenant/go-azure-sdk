package beta

import (
	"encoding/json"
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ MacOSCertificateProfileBase = MacOSScepCertificateProfile{}

type MacOSScepCertificateProfile struct {
	// AllowAllAppsAccess setting
	AllowAllAppsAccess nullable.Type[bool] `json:"allowAllAppsAccess,omitempty"`

	// Target store certificate. Possible values are: user, machine.
	CertificateStore *CertificateStore `json:"certificateStore,omitempty"`

	// Custom Subject Alternative Name Settings. This collection can contain a maximum of 500 elements.
	CustomSubjectAlternativeNames *[]CustomSubjectAlternativeName `json:"customSubjectAlternativeNames,omitempty"`

	// Indicates the deployment channel type used to deploy the configuration profile. Possible values are deviceChannel,
	// userChannel. Possible values are: deviceChannel, userChannel, unknownFutureValue.
	DeploymentChannel *AppleDeploymentChannel `json:"deploymentChannel,omitempty"`

	// Extended Key Usage (EKU) settings. This collection can contain a maximum of 500 elements.
	ExtendedKeyUsages *[]ExtendedKeyUsage `json:"extendedKeyUsages,omitempty"`

	// SCEP Hash Algorithm. Possible values are: sha1, sha2.
	HashAlgorithm *HashAlgorithms `json:"hashAlgorithm,omitempty"`

	// Key Size Options.
	KeySize *KeySize `json:"keySize,omitempty"`

	// Key Usage Options.
	KeyUsage *KeyUsages `json:"keyUsage,omitempty"`

	// Certificate state for devices. This collection can contain a maximum of 2147483647 elements.
	ManagedDeviceCertificateStates *[]ManagedDeviceCertificateState `json:"managedDeviceCertificateStates,omitempty"`

	// Trusted Root Certificate.
	RootCertificate *MacOSTrustedRootCertificate `json:"rootCertificate,omitempty"`

	// SCEP Server Url(s).
	ScepServerUrls *[]string `json:"scepServerUrls,omitempty"`

	// Custom String that defines the AAD Attribute.
	SubjectAlternativeNameFormatString nullable.Type[string] `json:"subjectAlternativeNameFormatString,omitempty"`

	// Custom format to use with SubjectNameFormat = Custom. Example: CN={{EmailAddress}},E={{EmailAddress}},OU=Enterprise
	// Users,O=Contoso Corporation,L=Redmond,ST=WA,C=US
	SubjectNameFormatString nullable.Type[string] `json:"subjectNameFormatString,omitempty"`

	// Fields inherited from MacOSCertificateProfileBase

	// Certificate Validity Period Options.
	CertificateValidityPeriodScale *CertificateValidityPeriodScale `json:"certificateValidityPeriodScale,omitempty"`

	// Value for the Certificate Validity Period.
	CertificateValidityPeriodValue *int64 `json:"certificateValidityPeriodValue,omitempty"`

	// Certificate renewal threshold percentage.
	RenewalThresholdPercentage *int64 `json:"renewalThresholdPercentage,omitempty"`

	// Certificate Subject Alternative Name Type. Possible values are: none, emailAddress, userPrincipalName,
	// customAzureADAttribute, domainNameService, universalResourceIdentifier.
	SubjectAlternativeNameType *SubjectAlternativeNameType `json:"subjectAlternativeNameType,omitempty"`

	// Subject Name Format Options for Apple devices.
	SubjectNameFormat *AppleSubjectNameFormat `json:"subjectNameFormat,omitempty"`

	// Fields inherited from DeviceConfiguration

	// The list of assignments for the device configuration profile.
	Assignments *[]DeviceConfigurationAssignment `json:"assignments,omitempty"`

	// DateTime the object was created.
	CreatedDateTime *string `json:"createdDateTime,omitempty"`

	// Admin provided description of the Device Configuration.
	Description nullable.Type[string] `json:"description,omitempty"`

	// The device mode applicability rule for this Policy.
	DeviceManagementApplicabilityRuleDeviceMode *DeviceManagementApplicabilityRuleDeviceMode `json:"deviceManagementApplicabilityRuleDeviceMode,omitempty"`

	// The OS edition applicability for this Policy.
	DeviceManagementApplicabilityRuleOsEdition *DeviceManagementApplicabilityRuleOsEdition `json:"deviceManagementApplicabilityRuleOsEdition,omitempty"`

	// The OS version applicability rule for this Policy.
	DeviceManagementApplicabilityRuleOsVersion *DeviceManagementApplicabilityRuleOsVersion `json:"deviceManagementApplicabilityRuleOsVersion,omitempty"`

	// Device Configuration Setting State Device Summary
	DeviceSettingStateSummaries *[]SettingStateDeviceSummary `json:"deviceSettingStateSummaries,omitempty"`

	// Device Configuration devices status overview
	DeviceStatusOverview *DeviceConfigurationDeviceOverview `json:"deviceStatusOverview,omitempty"`

	// Device configuration installation status by device.
	DeviceStatuses *[]DeviceConfigurationDeviceStatus `json:"deviceStatuses,omitempty"`

	// Admin provided name of the device configuration.
	DisplayName *string `json:"displayName,omitempty"`

	// The list of group assignments for the device configuration profile.
	GroupAssignments *[]DeviceConfigurationGroupAssignment `json:"groupAssignments,omitempty"`

	// DateTime the object was last modified.
	LastModifiedDateTime *string `json:"lastModifiedDateTime,omitempty"`

	// List of Scope Tags for this Entity instance.
	RoleScopeTagIds *[]string `json:"roleScopeTagIds,omitempty"`

	// Indicates whether or not the underlying Device Configuration supports the assignment of scope tags. Assigning to the
	// ScopeTags property is not allowed when this value is false and entities will not be visible to scoped users. This
	// occurs for Legacy policies created in Silverlight and can be resolved by deleting and recreating the policy in the
	// Azure Portal. This property is read-only.
	SupportsScopeTags *bool `json:"supportsScopeTags,omitempty"`

	// Device Configuration users status overview
	UserStatusOverview *DeviceConfigurationUserOverview `json:"userStatusOverview,omitempty"`

	// Device configuration installation status by user.
	UserStatuses *[]DeviceConfigurationUserStatus `json:"userStatuses,omitempty"`

	// Version of the device configuration.
	Version *int64 `json:"version,omitempty"`

	// Fields inherited from Entity

	// The unique identifier for an entity. Read-only.
	Id *string `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s MacOSScepCertificateProfile) MacOSCertificateProfileBase() BaseMacOSCertificateProfileBaseImpl {
	return BaseMacOSCertificateProfileBaseImpl{
		CertificateValidityPeriodScale:              s.CertificateValidityPeriodScale,
		CertificateValidityPeriodValue:              s.CertificateValidityPeriodValue,
		RenewalThresholdPercentage:                  s.RenewalThresholdPercentage,
		SubjectAlternativeNameType:                  s.SubjectAlternativeNameType,
		SubjectNameFormat:                           s.SubjectNameFormat,
		Assignments:                                 s.Assignments,
		CreatedDateTime:                             s.CreatedDateTime,
		Description:                                 s.Description,
		DeviceManagementApplicabilityRuleDeviceMode: s.DeviceManagementApplicabilityRuleDeviceMode,
		DeviceManagementApplicabilityRuleOsEdition:  s.DeviceManagementApplicabilityRuleOsEdition,
		DeviceManagementApplicabilityRuleOsVersion:  s.DeviceManagementApplicabilityRuleOsVersion,
		DeviceSettingStateSummaries:                 s.DeviceSettingStateSummaries,
		DeviceStatusOverview:                        s.DeviceStatusOverview,
		DeviceStatuses:                              s.DeviceStatuses,
		DisplayName:                                 s.DisplayName,
		GroupAssignments:                            s.GroupAssignments,
		LastModifiedDateTime:                        s.LastModifiedDateTime,
		RoleScopeTagIds:                             s.RoleScopeTagIds,
		SupportsScopeTags:                           s.SupportsScopeTags,
		UserStatusOverview:                          s.UserStatusOverview,
		UserStatuses:                                s.UserStatuses,
		Version:                                     s.Version,
		Id:                                          s.Id,
		ODataId:                                     s.ODataId,
		ODataType:                                   s.ODataType,
	}
}

func (s MacOSScepCertificateProfile) DeviceConfiguration() BaseDeviceConfigurationImpl {
	return BaseDeviceConfigurationImpl{
		Assignments:     s.Assignments,
		CreatedDateTime: s.CreatedDateTime,
		Description:     s.Description,
		DeviceManagementApplicabilityRuleDeviceMode: s.DeviceManagementApplicabilityRuleDeviceMode,
		DeviceManagementApplicabilityRuleOsEdition:  s.DeviceManagementApplicabilityRuleOsEdition,
		DeviceManagementApplicabilityRuleOsVersion:  s.DeviceManagementApplicabilityRuleOsVersion,
		DeviceSettingStateSummaries:                 s.DeviceSettingStateSummaries,
		DeviceStatusOverview:                        s.DeviceStatusOverview,
		DeviceStatuses:                              s.DeviceStatuses,
		DisplayName:                                 s.DisplayName,
		GroupAssignments:                            s.GroupAssignments,
		LastModifiedDateTime:                        s.LastModifiedDateTime,
		RoleScopeTagIds:                             s.RoleScopeTagIds,
		SupportsScopeTags:                           s.SupportsScopeTags,
		UserStatusOverview:                          s.UserStatusOverview,
		UserStatuses:                                s.UserStatuses,
		Version:                                     s.Version,
		Id:                                          s.Id,
		ODataId:                                     s.ODataId,
		ODataType:                                   s.ODataType,
	}
}

func (s MacOSScepCertificateProfile) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = MacOSScepCertificateProfile{}

func (s MacOSScepCertificateProfile) MarshalJSON() ([]byte, error) {
	type wrapper MacOSScepCertificateProfile
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling MacOSScepCertificateProfile: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling MacOSScepCertificateProfile: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.macOSScepCertificateProfile"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling MacOSScepCertificateProfile: %+v", err)
	}

	return encoded, nil
}
