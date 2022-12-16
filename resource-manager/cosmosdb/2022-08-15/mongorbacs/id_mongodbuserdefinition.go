package mongorbacs

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = MongodbUserDefinitionId{}

// MongodbUserDefinitionId is a struct representing the Resource ID for a Mongodb User Definition
type MongodbUserDefinitionId struct {
	SubscriptionId        string
	ResourceGroupName     string
	AccountName           string
	MongoUserDefinitionId string
}

// NewMongodbUserDefinitionID returns a new MongodbUserDefinitionId struct
func NewMongodbUserDefinitionID(subscriptionId string, resourceGroupName string, accountName string, mongoUserDefinitionId string) MongodbUserDefinitionId {
	return MongodbUserDefinitionId{
		SubscriptionId:        subscriptionId,
		ResourceGroupName:     resourceGroupName,
		AccountName:           accountName,
		MongoUserDefinitionId: mongoUserDefinitionId,
	}
}

// ParseMongodbUserDefinitionID parses 'input' into a MongodbUserDefinitionId
func ParseMongodbUserDefinitionID(input string) (*MongodbUserDefinitionId, error) {
	parser := resourceids.NewParserFromResourceIdType(MongodbUserDefinitionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MongodbUserDefinitionId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.AccountName, ok = parsed.Parsed["accountName"]; !ok {
		return nil, fmt.Errorf("the segment 'accountName' was not found in the resource id %q", input)
	}

	if id.MongoUserDefinitionId, ok = parsed.Parsed["mongoUserDefinitionId"]; !ok {
		return nil, fmt.Errorf("the segment 'mongoUserDefinitionId' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParseMongodbUserDefinitionIDInsensitively parses 'input' case-insensitively into a MongodbUserDefinitionId
// note: this method should only be used for API response data and not user input
func ParseMongodbUserDefinitionIDInsensitively(input string) (*MongodbUserDefinitionId, error) {
	parser := resourceids.NewParserFromResourceIdType(MongodbUserDefinitionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MongodbUserDefinitionId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.AccountName, ok = parsed.Parsed["accountName"]; !ok {
		return nil, fmt.Errorf("the segment 'accountName' was not found in the resource id %q", input)
	}

	if id.MongoUserDefinitionId, ok = parsed.Parsed["mongoUserDefinitionId"]; !ok {
		return nil, fmt.Errorf("the segment 'mongoUserDefinitionId' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidateMongodbUserDefinitionID checks that 'input' can be parsed as a Mongodb User Definition ID
func ValidateMongodbUserDefinitionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMongodbUserDefinitionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Mongodb User Definition ID
func (id MongodbUserDefinitionId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.DocumentDB/databaseAccounts/%s/mongodbUserDefinitions/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.AccountName, id.MongoUserDefinitionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Mongodb User Definition ID
func (id MongodbUserDefinitionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftDocumentDB", "Microsoft.DocumentDB", "Microsoft.DocumentDB"),
		resourceids.StaticSegment("staticDatabaseAccounts", "databaseAccounts", "databaseAccounts"),
		resourceids.UserSpecifiedSegment("accountName", "accountValue"),
		resourceids.StaticSegment("staticMongodbUserDefinitions", "mongodbUserDefinitions", "mongodbUserDefinitions"),
		resourceids.UserSpecifiedSegment("mongoUserDefinitionId", "mongoUserDefinitionIdValue"),
	}
}

// String returns a human-readable description of this Mongodb User Definition ID
func (id MongodbUserDefinitionId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Account Name: %q", id.AccountName),
		fmt.Sprintf("Mongo User Definition: %q", id.MongoUserDefinitionId),
	}
	return fmt.Sprintf("Mongodb User Definition (%s)", strings.Join(components, "\n"))
}