package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityB2cUserFlowIdLanguageIdDefaultPageId{}

func TestNewIdentityB2cUserFlowIdLanguageIdDefaultPageID(t *testing.T) {
	id := NewIdentityB2cUserFlowIdLanguageIdDefaultPageID("b2cIdentityUserFlowId", "userFlowLanguageConfigurationId", "userFlowLanguagePageId")

	if id.B2cIdentityUserFlowId != "b2cIdentityUserFlowId" {
		t.Fatalf("Expected %q but got %q for Segment 'B2cIdentityUserFlowId'", id.B2cIdentityUserFlowId, "b2cIdentityUserFlowId")
	}

	if id.UserFlowLanguageConfigurationId != "userFlowLanguageConfigurationId" {
		t.Fatalf("Expected %q but got %q for Segment 'UserFlowLanguageConfigurationId'", id.UserFlowLanguageConfigurationId, "userFlowLanguageConfigurationId")
	}

	if id.UserFlowLanguagePageId != "userFlowLanguagePageId" {
		t.Fatalf("Expected %q but got %q for Segment 'UserFlowLanguagePageId'", id.UserFlowLanguagePageId, "userFlowLanguagePageId")
	}
}

func TestFormatIdentityB2cUserFlowIdLanguageIdDefaultPageID(t *testing.T) {
	actual := NewIdentityB2cUserFlowIdLanguageIdDefaultPageID("b2cIdentityUserFlowId", "userFlowLanguageConfigurationId", "userFlowLanguagePageId").ID()
	expected := "/identity/b2cUserFlows/b2cIdentityUserFlowId/languages/userFlowLanguageConfigurationId/defaultPages/userFlowLanguagePageId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseIdentityB2cUserFlowIdLanguageIdDefaultPageID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityB2cUserFlowIdLanguageIdDefaultPageId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identity",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identity/b2cUserFlows",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identity/b2cUserFlows/b2cIdentityUserFlowId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identity/b2cUserFlows/b2cIdentityUserFlowId/languages",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identity/b2cUserFlows/b2cIdentityUserFlowId/languages/userFlowLanguageConfigurationId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identity/b2cUserFlows/b2cIdentityUserFlowId/languages/userFlowLanguageConfigurationId/defaultPages",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identity/b2cUserFlows/b2cIdentityUserFlowId/languages/userFlowLanguageConfigurationId/defaultPages/userFlowLanguagePageId",
			Expected: &IdentityB2cUserFlowIdLanguageIdDefaultPageId{
				B2cIdentityUserFlowId:           "b2cIdentityUserFlowId",
				UserFlowLanguageConfigurationId: "userFlowLanguageConfigurationId",
				UserFlowLanguagePageId:          "userFlowLanguagePageId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identity/b2cUserFlows/b2cIdentityUserFlowId/languages/userFlowLanguageConfigurationId/defaultPages/userFlowLanguagePageId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityB2cUserFlowIdLanguageIdDefaultPageID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.B2cIdentityUserFlowId != v.Expected.B2cIdentityUserFlowId {
			t.Fatalf("Expected %q but got %q for B2cIdentityUserFlowId", v.Expected.B2cIdentityUserFlowId, actual.B2cIdentityUserFlowId)
		}

		if actual.UserFlowLanguageConfigurationId != v.Expected.UserFlowLanguageConfigurationId {
			t.Fatalf("Expected %q but got %q for UserFlowLanguageConfigurationId", v.Expected.UserFlowLanguageConfigurationId, actual.UserFlowLanguageConfigurationId)
		}

		if actual.UserFlowLanguagePageId != v.Expected.UserFlowLanguagePageId {
			t.Fatalf("Expected %q but got %q for UserFlowLanguagePageId", v.Expected.UserFlowLanguagePageId, actual.UserFlowLanguagePageId)
		}

	}
}

func TestParseIdentityB2cUserFlowIdLanguageIdDefaultPageIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityB2cUserFlowIdLanguageIdDefaultPageId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identity",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItY",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identity/b2cUserFlows",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItY/b2cUsErFlOwS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identity/b2cUserFlows/b2cIdentityUserFlowId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItY/b2cUsErFlOwS/b2cIdEnTiTyUsErFlOwId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identity/b2cUserFlows/b2cIdentityUserFlowId/languages",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItY/b2cUsErFlOwS/b2cIdEnTiTyUsErFlOwId/lAnGuAgEs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identity/b2cUserFlows/b2cIdentityUserFlowId/languages/userFlowLanguageConfigurationId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItY/b2cUsErFlOwS/b2cIdEnTiTyUsErFlOwId/lAnGuAgEs/uSeRfLoWlAnGuAgEcOnFiGuRaTiOnId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identity/b2cUserFlows/b2cIdentityUserFlowId/languages/userFlowLanguageConfigurationId/defaultPages",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItY/b2cUsErFlOwS/b2cIdEnTiTyUsErFlOwId/lAnGuAgEs/uSeRfLoWlAnGuAgEcOnFiGuRaTiOnId/dEfAuLtPaGeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identity/b2cUserFlows/b2cIdentityUserFlowId/languages/userFlowLanguageConfigurationId/defaultPages/userFlowLanguagePageId",
			Expected: &IdentityB2cUserFlowIdLanguageIdDefaultPageId{
				B2cIdentityUserFlowId:           "b2cIdentityUserFlowId",
				UserFlowLanguageConfigurationId: "userFlowLanguageConfigurationId",
				UserFlowLanguagePageId:          "userFlowLanguagePageId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identity/b2cUserFlows/b2cIdentityUserFlowId/languages/userFlowLanguageConfigurationId/defaultPages/userFlowLanguagePageId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItY/b2cUsErFlOwS/b2cIdEnTiTyUsErFlOwId/lAnGuAgEs/uSeRfLoWlAnGuAgEcOnFiGuRaTiOnId/dEfAuLtPaGeS/uSeRfLoWlAnGuAgEpAgEiD",
			Expected: &IdentityB2cUserFlowIdLanguageIdDefaultPageId{
				B2cIdentityUserFlowId:           "b2cIdEnTiTyUsErFlOwId",
				UserFlowLanguageConfigurationId: "uSeRfLoWlAnGuAgEcOnFiGuRaTiOnId",
				UserFlowLanguagePageId:          "uSeRfLoWlAnGuAgEpAgEiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItY/b2cUsErFlOwS/b2cIdEnTiTyUsErFlOwId/lAnGuAgEs/uSeRfLoWlAnGuAgEcOnFiGuRaTiOnId/dEfAuLtPaGeS/uSeRfLoWlAnGuAgEpAgEiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityB2cUserFlowIdLanguageIdDefaultPageIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.B2cIdentityUserFlowId != v.Expected.B2cIdentityUserFlowId {
			t.Fatalf("Expected %q but got %q for B2cIdentityUserFlowId", v.Expected.B2cIdentityUserFlowId, actual.B2cIdentityUserFlowId)
		}

		if actual.UserFlowLanguageConfigurationId != v.Expected.UserFlowLanguageConfigurationId {
			t.Fatalf("Expected %q but got %q for UserFlowLanguageConfigurationId", v.Expected.UserFlowLanguageConfigurationId, actual.UserFlowLanguageConfigurationId)
		}

		if actual.UserFlowLanguagePageId != v.Expected.UserFlowLanguagePageId {
			t.Fatalf("Expected %q but got %q for UserFlowLanguagePageId", v.Expected.UserFlowLanguagePageId, actual.UserFlowLanguagePageId)
		}

	}
}

func TestSegmentsForIdentityB2cUserFlowIdLanguageIdDefaultPageId(t *testing.T) {
	segments := IdentityB2cUserFlowIdLanguageIdDefaultPageId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("IdentityB2cUserFlowIdLanguageIdDefaultPageId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
