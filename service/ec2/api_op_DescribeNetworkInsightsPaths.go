// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Describes one or more of your paths.
func (c *Client) DescribeNetworkInsightsPaths(ctx context.Context, params *DescribeNetworkInsightsPathsInput, optFns ...func(*Options)) (*DescribeNetworkInsightsPathsOutput, error) {
	if params == nil {
		params = &DescribeNetworkInsightsPathsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeNetworkInsightsPaths", params, optFns, c.addOperationDescribeNetworkInsightsPathsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeNetworkInsightsPathsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeNetworkInsightsPathsInput struct {

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation . Otherwise, it is
	// UnauthorizedOperation .
	DryRun *bool

	// The filters. The following are the possible values:
	//   - destination - The ID of the resource.
	//   - filter-at-source.source-address - The source IPv4 address at the source.
	//   - filter-at-source.source-port-range - The source port range at the source.
	//   - filter-at-source.destination-address - The destination IPv4 address at the
	//   source.
	//   - filter-at-source.destination-port-range - The destination port range at the
	//   source.
	//   - filter-at-destination.source-address - The source IPv4 address at the
	//   destination.
	//   - filter-at-destination.source-port-range - The source port range at the
	//   destination.
	//   - filter-at-destination.destination-address - The destination IPv4 address at
	//   the destination.
	//   - filter-at-destination.destination-port-range - The destination port range
	//   at the destination.
	//   - protocol - The protocol.
	//   - source - The ID of the resource.
	Filters []types.Filter

	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	MaxResults *int32

	// The IDs of the paths.
	NetworkInsightsPathIds []string

	// The token for the next page of results.
	NextToken *string

	noSmithyDocumentSerde
}

type DescribeNetworkInsightsPathsOutput struct {

	// Information about the paths.
	NetworkInsightsPaths []types.NetworkInsightsPath

	// The token to use to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeNetworkInsightsPathsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpDescribeNetworkInsightsPaths{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpDescribeNetworkInsightsPaths{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeNetworkInsightsPaths(options.Region), middleware.Before); err != nil {
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

// DescribeNetworkInsightsPathsAPIClient is a client that implements the
// DescribeNetworkInsightsPaths operation.
type DescribeNetworkInsightsPathsAPIClient interface {
	DescribeNetworkInsightsPaths(context.Context, *DescribeNetworkInsightsPathsInput, ...func(*Options)) (*DescribeNetworkInsightsPathsOutput, error)
}

var _ DescribeNetworkInsightsPathsAPIClient = (*Client)(nil)

// DescribeNetworkInsightsPathsPaginatorOptions is the paginator options for
// DescribeNetworkInsightsPaths
type DescribeNetworkInsightsPathsPaginatorOptions struct {
	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeNetworkInsightsPathsPaginator is a paginator for
// DescribeNetworkInsightsPaths
type DescribeNetworkInsightsPathsPaginator struct {
	options   DescribeNetworkInsightsPathsPaginatorOptions
	client    DescribeNetworkInsightsPathsAPIClient
	params    *DescribeNetworkInsightsPathsInput
	nextToken *string
	firstPage bool
}

// NewDescribeNetworkInsightsPathsPaginator returns a new
// DescribeNetworkInsightsPathsPaginator
func NewDescribeNetworkInsightsPathsPaginator(client DescribeNetworkInsightsPathsAPIClient, params *DescribeNetworkInsightsPathsInput, optFns ...func(*DescribeNetworkInsightsPathsPaginatorOptions)) *DescribeNetworkInsightsPathsPaginator {
	if params == nil {
		params = &DescribeNetworkInsightsPathsInput{}
	}

	options := DescribeNetworkInsightsPathsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeNetworkInsightsPathsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeNetworkInsightsPathsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeNetworkInsightsPaths page.
func (p *DescribeNetworkInsightsPathsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeNetworkInsightsPathsOutput, error) {
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

	result, err := p.client.DescribeNetworkInsightsPaths(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeNetworkInsightsPaths(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DescribeNetworkInsightsPaths",
	}
}
