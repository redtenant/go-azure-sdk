package prerules

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = PreRuleId{}

func TestNewPreRuleID(t *testing.T) {
	id := NewPreRuleID("globalRuleStackValue", "preRuleValue")

	if id.GlobalRuleStackName != "globalRuleStackValue" {
		t.Fatalf("Expected %q but got %q for Segment 'GlobalRuleStackName'", id.GlobalRuleStackName, "globalRuleStackValue")
	}

	if id.PreRuleName != "preRuleValue" {
		t.Fatalf("Expected %q but got %q for Segment 'PreRuleName'", id.PreRuleName, "preRuleValue")
	}
}

func TestFormatPreRuleID(t *testing.T) {
	actual := NewPreRuleID("globalRuleStackValue", "preRuleValue").ID()
	expected := "/providers/PaloAltoNetworks.CloudNGFW/globalRuleStacks/globalRuleStackValue/preRules/preRuleValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParsePreRuleID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *PreRuleId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/PaloAltoNetworks.CloudNGFW",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/PaloAltoNetworks.CloudNGFW/globalRuleStacks",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/PaloAltoNetworks.CloudNGFW/globalRuleStacks/globalRuleStackValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/PaloAltoNetworks.CloudNGFW/globalRuleStacks/globalRuleStackValue/preRules",
			Error: true,
		},
		{
			// Valid URI
			Input: "/providers/PaloAltoNetworks.CloudNGFW/globalRuleStacks/globalRuleStackValue/preRules/preRuleValue",
			Expected: &PreRuleId{
				GlobalRuleStackName: "globalRuleStackValue",
				PreRuleName:         "preRuleValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/providers/PaloAltoNetworks.CloudNGFW/globalRuleStacks/globalRuleStackValue/preRules/preRuleValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParsePreRuleID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.GlobalRuleStackName != v.Expected.GlobalRuleStackName {
			t.Fatalf("Expected %q but got %q for GlobalRuleStackName", v.Expected.GlobalRuleStackName, actual.GlobalRuleStackName)
		}

		if actual.PreRuleName != v.Expected.PreRuleName {
			t.Fatalf("Expected %q but got %q for PreRuleName", v.Expected.PreRuleName, actual.PreRuleName)
		}

	}
}

func TestParsePreRuleIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *PreRuleId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/PaloAltoNetworks.CloudNGFW",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/pAlOaLtOnEtWoRkS.ClOuDnGfW",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/PaloAltoNetworks.CloudNGFW/globalRuleStacks",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/pAlOaLtOnEtWoRkS.ClOuDnGfW/gLoBaLrUlEsTaCkS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/PaloAltoNetworks.CloudNGFW/globalRuleStacks/globalRuleStackValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/pAlOaLtOnEtWoRkS.ClOuDnGfW/gLoBaLrUlEsTaCkS/gLoBaLrUlEsTaCkVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/PaloAltoNetworks.CloudNGFW/globalRuleStacks/globalRuleStackValue/preRules",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/pAlOaLtOnEtWoRkS.ClOuDnGfW/gLoBaLrUlEsTaCkS/gLoBaLrUlEsTaCkVaLuE/pReRuLeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/providers/PaloAltoNetworks.CloudNGFW/globalRuleStacks/globalRuleStackValue/preRules/preRuleValue",
			Expected: &PreRuleId{
				GlobalRuleStackName: "globalRuleStackValue",
				PreRuleName:         "preRuleValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/providers/PaloAltoNetworks.CloudNGFW/globalRuleStacks/globalRuleStackValue/preRules/preRuleValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/pAlOaLtOnEtWoRkS.ClOuDnGfW/gLoBaLrUlEsTaCkS/gLoBaLrUlEsTaCkVaLuE/pReRuLeS/pReRuLeVaLuE",
			Expected: &PreRuleId{
				GlobalRuleStackName: "gLoBaLrUlEsTaCkVaLuE",
				PreRuleName:         "pReRuLeVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/pAlOaLtOnEtWoRkS.ClOuDnGfW/gLoBaLrUlEsTaCkS/gLoBaLrUlEsTaCkVaLuE/pReRuLeS/pReRuLeVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParsePreRuleIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.GlobalRuleStackName != v.Expected.GlobalRuleStackName {
			t.Fatalf("Expected %q but got %q for GlobalRuleStackName", v.Expected.GlobalRuleStackName, actual.GlobalRuleStackName)
		}

		if actual.PreRuleName != v.Expected.PreRuleName {
			t.Fatalf("Expected %q but got %q for PreRuleName", v.Expected.PreRuleName, actual.PreRuleName)
		}

	}
}

func TestSegmentsForPreRuleId(t *testing.T) {
	segments := PreRuleId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("PreRuleId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
