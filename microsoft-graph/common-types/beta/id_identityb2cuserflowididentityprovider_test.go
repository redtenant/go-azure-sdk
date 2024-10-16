package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityB2cUserFlowIdIdentityProviderId{}

func TestNewIdentityB2cUserFlowIdIdentityProviderID(t *testing.T) {
	id := NewIdentityB2cUserFlowIdIdentityProviderID("b2cIdentityUserFlowId", "identityProviderId")

	if id.B2cIdentityUserFlowId != "b2cIdentityUserFlowId" {
		t.Fatalf("Expected %q but got %q for Segment 'B2cIdentityUserFlowId'", id.B2cIdentityUserFlowId, "b2cIdentityUserFlowId")
	}

	if id.IdentityProviderId != "identityProviderId" {
		t.Fatalf("Expected %q but got %q for Segment 'IdentityProviderId'", id.IdentityProviderId, "identityProviderId")
	}
}

func TestFormatIdentityB2cUserFlowIdIdentityProviderID(t *testing.T) {
	actual := NewIdentityB2cUserFlowIdIdentityProviderID("b2cIdentityUserFlowId", "identityProviderId").ID()
	expected := "/identity/b2cUserFlows/b2cIdentityUserFlowId/identityProviders/identityProviderId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseIdentityB2cUserFlowIdIdentityProviderID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityB2cUserFlowIdIdentityProviderId
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
			Input: "/identity/b2cUserFlows/b2cIdentityUserFlowId/identityProviders",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identity/b2cUserFlows/b2cIdentityUserFlowId/identityProviders/identityProviderId",
			Expected: &IdentityB2cUserFlowIdIdentityProviderId{
				B2cIdentityUserFlowId: "b2cIdentityUserFlowId",
				IdentityProviderId:    "identityProviderId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identity/b2cUserFlows/b2cIdentityUserFlowId/identityProviders/identityProviderId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityB2cUserFlowIdIdentityProviderID(v.Input)
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

		if actual.IdentityProviderId != v.Expected.IdentityProviderId {
			t.Fatalf("Expected %q but got %q for IdentityProviderId", v.Expected.IdentityProviderId, actual.IdentityProviderId)
		}

	}
}

func TestParseIdentityB2cUserFlowIdIdentityProviderIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityB2cUserFlowIdIdentityProviderId
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
			Input: "/identity/b2cUserFlows/b2cIdentityUserFlowId/identityProviders",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItY/b2cUsErFlOwS/b2cIdEnTiTyUsErFlOwId/iDeNtItYpRoViDeRs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identity/b2cUserFlows/b2cIdentityUserFlowId/identityProviders/identityProviderId",
			Expected: &IdentityB2cUserFlowIdIdentityProviderId{
				B2cIdentityUserFlowId: "b2cIdentityUserFlowId",
				IdentityProviderId:    "identityProviderId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identity/b2cUserFlows/b2cIdentityUserFlowId/identityProviders/identityProviderId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItY/b2cUsErFlOwS/b2cIdEnTiTyUsErFlOwId/iDeNtItYpRoViDeRs/iDeNtItYpRoViDeRiD",
			Expected: &IdentityB2cUserFlowIdIdentityProviderId{
				B2cIdentityUserFlowId: "b2cIdEnTiTyUsErFlOwId",
				IdentityProviderId:    "iDeNtItYpRoViDeRiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItY/b2cUsErFlOwS/b2cIdEnTiTyUsErFlOwId/iDeNtItYpRoViDeRs/iDeNtItYpRoViDeRiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityB2cUserFlowIdIdentityProviderIDInsensitively(v.Input)
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

		if actual.IdentityProviderId != v.Expected.IdentityProviderId {
			t.Fatalf("Expected %q but got %q for IdentityProviderId", v.Expected.IdentityProviderId, actual.IdentityProviderId)
		}

	}
}

func TestSegmentsForIdentityB2cUserFlowIdIdentityProviderId(t *testing.T) {
	segments := IdentityB2cUserFlowIdIdentityProviderId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("IdentityB2cUserFlowIdIdentityProviderId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
