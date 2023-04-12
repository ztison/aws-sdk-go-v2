// Code generated by smithy-go-codegen DO NOT EDIT.

package customerprofiles

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/customerprofiles/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Searches for profiles within a specific domain using one or more predefined
// search keys (e.g., _fullName, _phone, _email, _account, etc.) and/or
// custom-defined search keys. A search key is a data type pair that consists of a
// KeyName and Values list. This operation supports searching for profiles with a
// minimum of 1 key-value(s) pair and up to 5 key-value(s) pairs using either AND
// or OR logic.
func (c *Client) SearchProfiles(ctx context.Context, params *SearchProfilesInput, optFns ...func(*Options)) (*SearchProfilesOutput, error) {
	if params == nil {
		params = &SearchProfilesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SearchProfiles", params, optFns, c.addOperationSearchProfilesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SearchProfilesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SearchProfilesInput struct {

	// The unique name of the domain.
	//
	// This member is required.
	DomainName *string

	// A searchable identifier of a customer profile. The predefined keys you can use
	// to search include: _account, _profileId, _assetId, _caseId, _orderId, _fullName,
	// _phone, _email, _ctrContactId, _marketoLeadId, _salesforceAccountId,
	// _salesforceContactId, _salesforceAssetId, _zendeskUserId, _zendeskExternalId,
	// _zendeskTicketId, _serviceNowSystemId, _serviceNowIncidentId, _segmentUserId,
	// _shopifyCustomerId, _shopifyOrderId.
	//
	// This member is required.
	KeyName *string

	// A list of key values.
	//
	// This member is required.
	Values []string

	// A list of AdditionalSearchKey objects that are each searchable identifiers of a
	// profile. Each AdditionalSearchKey object contains a KeyName and a list of Values
	// associated with that specific key (i.e., a key-value(s) pair). These additional
	// search keys will be used in conjunction with the LogicalOperator and the
	// required KeyName and Values parameters to search for profiles that satisfy the
	// search criteria.
	AdditionalSearchKeys []types.AdditionalSearchKey

	// Relationship between all specified search keys that will be used to search for
	// profiles. This includes the required KeyName and Values parameters as well as
	// any key-value(s) pairs specified in the AdditionalSearchKeys list. This
	// parameter influences which profiles will be returned in the response in the
	// following manner:
	//   - AND - The response only includes profiles that match all of the search keys.
	//   - OR - The response includes profiles that match at least one of the search
	//   keys.
	// The OR relationship is the default behavior if this parameter is not included
	// in the request.
	LogicalOperator types.LogicalOperator

	// The maximum number of objects returned per page. The default is 20 if this
	// parameter is not included in the request.
	MaxResults *int32

	// The pagination token from the previous SearchProfiles API call.
	NextToken *string

	noSmithyDocumentSerde
}

type SearchProfilesOutput struct {

	// The list of Profiles matching the search criteria.
	Items []types.Profile

	// The pagination token from the previous SearchProfiles API call.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationSearchProfilesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpSearchProfiles{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpSearchProfiles{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpSearchProfilesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSearchProfiles(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opSearchProfiles(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "profile",
		OperationName: "SearchProfiles",
	}
}
