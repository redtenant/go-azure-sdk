package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityB2cUserFlowId{}

func TestNewIdentityB2cUserFlowID(t *testing.T) {
	id := NewIdentityB2cUserFlowID("b2cIdentityUserFlowId")

	if id.B2cIdentityUserFlowId != "b2cIdentityUserFlowId" {
		t.Fatalf("Expected %q but got %q for Segment 'B2cIdentityUserFlowId'", id.B2cIdentityUserFlowId, "b2cIdentityUserFlowId")
	}
}

func TestFormatIdentityB2cUserFlowID(t *testing.T) {
	actual := NewIdentityB2cUserFlowID("b2cIdentityUserFlowId").ID()
	expected := "/identity/b2cUserFlows/b2cIdentityUserFlowId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseIdentityB2cUserFlowID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityB2cUserFlowId
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
			// Valid URI
			Input: "/identity/b2cUserFlows/b2cIdentityUserFlowId",
			Expected: &IdentityB2cUserFlowId{
				B2cIdentityUserFlowId: "b2cIdentityUserFlowId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identity/b2cUserFlows/b2cIdentityUserFlowId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityB2cUserFlowID(v.Input)
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

	}
}

func TestParseIdentityB2cUserFlowIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityB2cUserFlowId
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
			// Valid URI
			Input: "/identity/b2cUserFlows/b2cIdentityUserFlowId",
			Expected: &IdentityB2cUserFlowId{
				B2cIdentityUserFlowId: "b2cIdentityUserFlowId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identity/b2cUserFlows/b2cIdentityUserFlowId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItY/b2cUsErFlOwS/b2cIdEnTiTyUsErFlOwId",
			Expected: &IdentityB2cUserFlowId{
				B2cIdentityUserFlowId: "b2cIdEnTiTyUsErFlOwId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItY/b2cUsErFlOwS/b2cIdEnTiTyUsErFlOwId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityB2cUserFlowIDInsensitively(v.Input)
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

	}
}

func TestSegmentsForIdentityB2cUserFlowId(t *testing.T) {
	segments := IdentityB2cUserFlowId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("IdentityB2cUserFlowId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
