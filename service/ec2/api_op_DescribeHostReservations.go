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

// Describes reservations that are associated with Dedicated Hosts in your account.
func (c *Client) DescribeHostReservations(ctx context.Context, params *DescribeHostReservationsInput, optFns ...func(*Options)) (*DescribeHostReservationsOutput, error) {
	if params == nil {
		params = &DescribeHostReservationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeHostReservations", params, optFns, c.addOperationDescribeHostReservationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeHostReservationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeHostReservationsInput struct {

	// The filters.
	//   - instance-family - The instance family (for example, m4 ).
	//   - payment-option - The payment option ( NoUpfront | PartialUpfront |
	//   AllUpfront ).
	//   - state - The state of the reservation ( payment-pending | payment-failed |
	//   active | retired ).
	//   - tag: - The key/value combination of a tag assigned to the resource. Use the
	//   tag key in the filter name and the tag value as the filter value. For example,
	//   to find all resources that have a tag with the key Owner and the value TeamA ,
	//   specify tag:Owner for the filter name and TeamA for the filter value.
	//   - tag-key - The key of a tag assigned to the resource. Use this filter to find
	//   all resources assigned a tag with a specific key, regardless of the tag value.
	Filter []types.Filter

	// The host reservation IDs.
	HostReservationIdSet []string

	// The maximum number of results to return for the request in a single page. The
	// remaining results can be seen by sending another request with the returned
	// nextToken value. This value can be between 5 and 500. If maxResults is given a
	// larger value than 500, you receive an error.
	MaxResults *int32

	// The token to use to retrieve the next page of results.
	NextToken *string

	noSmithyDocumentSerde
}

type DescribeHostReservationsOutput struct {

	// Details about the reservation's configuration.
	HostReservationSet []types.HostReservation

	// The token to use to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeHostReservationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpDescribeHostReservations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpDescribeHostReservations{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeHostReservations(options.Region), middleware.Before); err != nil {
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

// DescribeHostReservationsAPIClient is a client that implements the
// DescribeHostReservations operation.
type DescribeHostReservationsAPIClient interface {
	DescribeHostReservations(context.Context, *DescribeHostReservationsInput, ...func(*Options)) (*DescribeHostReservationsOutput, error)
}

var _ DescribeHostReservationsAPIClient = (*Client)(nil)

// DescribeHostReservationsPaginatorOptions is the paginator options for
// DescribeHostReservations
type DescribeHostReservationsPaginatorOptions struct {
	// The maximum number of results to return for the request in a single page. The
	// remaining results can be seen by sending another request with the returned
	// nextToken value. This value can be between 5 and 500. If maxResults is given a
	// larger value than 500, you receive an error.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeHostReservationsPaginator is a paginator for DescribeHostReservations
type DescribeHostReservationsPaginator struct {
	options   DescribeHostReservationsPaginatorOptions
	client    DescribeHostReservationsAPIClient
	params    *DescribeHostReservationsInput
	nextToken *string
	firstPage bool
}

// NewDescribeHostReservationsPaginator returns a new
// DescribeHostReservationsPaginator
func NewDescribeHostReservationsPaginator(client DescribeHostReservationsAPIClient, params *DescribeHostReservationsInput, optFns ...func(*DescribeHostReservationsPaginatorOptions)) *DescribeHostReservationsPaginator {
	if params == nil {
		params = &DescribeHostReservationsInput{}
	}

	options := DescribeHostReservationsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeHostReservationsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeHostReservationsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeHostReservations page.
func (p *DescribeHostReservationsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeHostReservationsOutput, error) {
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

	result, err := p.client.DescribeHostReservations(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeHostReservations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DescribeHostReservations",
	}
}
