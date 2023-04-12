// Code generated by smithy-go-codegen DO NOT EDIT.

package ram

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves the resource policies for the specified resources that you own and
// have shared.
func (c *Client) GetResourcePolicies(ctx context.Context, params *GetResourcePoliciesInput, optFns ...func(*Options)) (*GetResourcePoliciesOutput, error) {
	if params == nil {
		params = &GetResourcePoliciesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetResourcePolicies", params, optFns, c.addOperationGetResourcePoliciesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetResourcePoliciesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetResourcePoliciesInput struct {

	// Specifies the Amazon Resource Names (ARNs) (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html)
	// of the resources whose policies you want to retrieve.
	//
	// This member is required.
	ResourceArns []string

	// Specifies the total number of results that you want included on each page of
	// the response. If you do not include this parameter, it defaults to a value that
	// is specific to the operation. If additional items exist beyond the number you
	// specify, the NextToken response element is returned with a value (not null).
	// Include the specified value as the NextToken request parameter in the next call
	// to the operation to get the next part of the results. Note that the service
	// might return fewer results than the maximum even when there are more results
	// available. You should check NextToken after every operation to ensure that you
	// receive all of the results.
	MaxResults *int32

	// Specifies that you want to receive the next page of results. Valid only if you
	// received a NextToken response in the previous request. If you did, it indicates
	// that more output is available. Set this parameter to the value provided by the
	// previous call's NextToken response to request the next page of results.
	NextToken *string

	// Specifies the principal.
	Principal *string

	noSmithyDocumentSerde
}

type GetResourcePoliciesOutput struct {

	// If present, this value indicates that more output is available than is included
	// in the current response. Use this value in the NextToken request parameter in a
	// subsequent call to the operation to get the next part of the output. You should
	// repeat this until the NextToken response element comes back as null . This
	// indicates that this is the last page of results.
	NextToken *string

	// An array of resource policy documents in JSON format.
	Policies []string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetResourcePoliciesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetResourcePolicies{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetResourcePolicies{}, middleware.After)
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
	if err = addOpGetResourcePoliciesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetResourcePolicies(options.Region), middleware.Before); err != nil {
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

// GetResourcePoliciesAPIClient is a client that implements the
// GetResourcePolicies operation.
type GetResourcePoliciesAPIClient interface {
	GetResourcePolicies(context.Context, *GetResourcePoliciesInput, ...func(*Options)) (*GetResourcePoliciesOutput, error)
}

var _ GetResourcePoliciesAPIClient = (*Client)(nil)

// GetResourcePoliciesPaginatorOptions is the paginator options for
// GetResourcePolicies
type GetResourcePoliciesPaginatorOptions struct {
	// Specifies the total number of results that you want included on each page of
	// the response. If you do not include this parameter, it defaults to a value that
	// is specific to the operation. If additional items exist beyond the number you
	// specify, the NextToken response element is returned with a value (not null).
	// Include the specified value as the NextToken request parameter in the next call
	// to the operation to get the next part of the results. Note that the service
	// might return fewer results than the maximum even when there are more results
	// available. You should check NextToken after every operation to ensure that you
	// receive all of the results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetResourcePoliciesPaginator is a paginator for GetResourcePolicies
type GetResourcePoliciesPaginator struct {
	options   GetResourcePoliciesPaginatorOptions
	client    GetResourcePoliciesAPIClient
	params    *GetResourcePoliciesInput
	nextToken *string
	firstPage bool
}

// NewGetResourcePoliciesPaginator returns a new GetResourcePoliciesPaginator
func NewGetResourcePoliciesPaginator(client GetResourcePoliciesAPIClient, params *GetResourcePoliciesInput, optFns ...func(*GetResourcePoliciesPaginatorOptions)) *GetResourcePoliciesPaginator {
	if params == nil {
		params = &GetResourcePoliciesInput{}
	}

	options := GetResourcePoliciesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetResourcePoliciesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetResourcePoliciesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next GetResourcePolicies page.
func (p *GetResourcePoliciesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetResourcePoliciesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxResults = limit

	result, err := p.client.GetResourcePolicies(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opGetResourcePolicies(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ram",
		OperationName: "GetResourcePolicies",
	}
}
