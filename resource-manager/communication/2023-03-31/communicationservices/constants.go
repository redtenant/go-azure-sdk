package communicationservices

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CheckNameAvailabilityReason string

const (
	CheckNameAvailabilityReasonAlreadyExists CheckNameAvailabilityReason = "AlreadyExists"
	CheckNameAvailabilityReasonInvalid       CheckNameAvailabilityReason = "Invalid"
)

func PossibleValuesForCheckNameAvailabilityReason() []string {
	return []string{
		string(CheckNameAvailabilityReasonAlreadyExists),
		string(CheckNameAvailabilityReasonInvalid),
	}
}

func parseCheckNameAvailabilityReason(input string) (*CheckNameAvailabilityReason, error) {
	vals := map[string]CheckNameAvailabilityReason{
		"alreadyexists": CheckNameAvailabilityReasonAlreadyExists,
		"invalid":       CheckNameAvailabilityReasonInvalid,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CheckNameAvailabilityReason(input)
	return &out, nil
}

type CommunicationServicesProvisioningState string

const (
	CommunicationServicesProvisioningStateCanceled  CommunicationServicesProvisioningState = "Canceled"
	CommunicationServicesProvisioningStateCreating  CommunicationServicesProvisioningState = "Creating"
	CommunicationServicesProvisioningStateDeleting  CommunicationServicesProvisioningState = "Deleting"
	CommunicationServicesProvisioningStateFailed    CommunicationServicesProvisioningState = "Failed"
	CommunicationServicesProvisioningStateMoving    CommunicationServicesProvisioningState = "Moving"
	CommunicationServicesProvisioningStateRunning   CommunicationServicesProvisioningState = "Running"
	CommunicationServicesProvisioningStateSucceeded CommunicationServicesProvisioningState = "Succeeded"
	CommunicationServicesProvisioningStateUnknown   CommunicationServicesProvisioningState = "Unknown"
	CommunicationServicesProvisioningStateUpdating  CommunicationServicesProvisioningState = "Updating"
)

func PossibleValuesForCommunicationServicesProvisioningState() []string {
	return []string{
		string(CommunicationServicesProvisioningStateCanceled),
		string(CommunicationServicesProvisioningStateCreating),
		string(CommunicationServicesProvisioningStateDeleting),
		string(CommunicationServicesProvisioningStateFailed),
		string(CommunicationServicesProvisioningStateMoving),
		string(CommunicationServicesProvisioningStateRunning),
		string(CommunicationServicesProvisioningStateSucceeded),
		string(CommunicationServicesProvisioningStateUnknown),
		string(CommunicationServicesProvisioningStateUpdating),
	}
}

func parseCommunicationServicesProvisioningState(input string) (*CommunicationServicesProvisioningState, error) {
	vals := map[string]CommunicationServicesProvisioningState{
		"canceled":  CommunicationServicesProvisioningStateCanceled,
		"creating":  CommunicationServicesProvisioningStateCreating,
		"deleting":  CommunicationServicesProvisioningStateDeleting,
		"failed":    CommunicationServicesProvisioningStateFailed,
		"moving":    CommunicationServicesProvisioningStateMoving,
		"running":   CommunicationServicesProvisioningStateRunning,
		"succeeded": CommunicationServicesProvisioningStateSucceeded,
		"unknown":   CommunicationServicesProvisioningStateUnknown,
		"updating":  CommunicationServicesProvisioningStateUpdating,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CommunicationServicesProvisioningState(input)
	return &out, nil
}

type KeyType string

const (
	KeyTypePrimary   KeyType = "Primary"
	KeyTypeSecondary KeyType = "Secondary"
)

func PossibleValuesForKeyType() []string {
	return []string{
		string(KeyTypePrimary),
		string(KeyTypeSecondary),
	}
}

func parseKeyType(input string) (*KeyType, error) {
	vals := map[string]KeyType{
		"primary":   KeyTypePrimary,
		"secondary": KeyTypeSecondary,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := KeyType(input)
	return &out, nil
}
