// Code generated by smithy-go-codegen DO NOT EDIT.

package mediatailor

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/mediatailor/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the prefetch schedules for a playback configuration.
func (c *Client) ListPrefetchSchedules(ctx context.Context, params *ListPrefetchSchedulesInput, optFns ...func(*Options)) (*ListPrefetchSchedulesOutput, error) {
	if params == nil {
		params = &ListPrefetchSchedulesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListPrefetchSchedules", params, optFns, c.addOperationListPrefetchSchedulesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListPrefetchSchedulesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListPrefetchSchedulesInput struct {

	// Retrieves the prefetch schedule(s) for a specific playback configuration.
	//
	// This member is required.
	PlaybackConfigurationName *string

	// The maximum number of prefetch schedules that you want MediaTailor to return in
	// response to the current request. If there are more than MaxResults prefetch
	// schedules, use the value of NextToken in the response to get the next page of
	// results.
	MaxResults int32

	// (Optional) If the playback configuration has more than MaxResults prefetch
	// schedules, use NextToken to get the second and subsequent pages of results. For
	// the first ListPrefetchSchedulesRequest request, omit this value. For the second
	// and subsequent requests, get the value of NextToken from the previous response
	// and specify that value for NextToken in the request. If the previous response
	// didn't include a NextToken element, there are no more prefetch schedules to get.
	NextToken *string

	// An optional filtering parameter whereby MediaTailor filters the prefetch
	// schedules to include only specific streams.
	StreamId *string

	noSmithyDocumentSerde
}

type ListPrefetchSchedulesOutput struct {

	// Lists the prefetch schedules. An empty Items list doesn't mean there aren't
	// more items to fetch, just that that page was empty.
	Items []types.PrefetchSchedule

	// Pagination token returned by the list request when results exceed the maximum
	// allowed. Use the token to fetch the next page of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListPrefetchSchedulesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListPrefetchSchedules{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListPrefetchSchedules{}, middleware.After)
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
	if err = addOpListPrefetchSchedulesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListPrefetchSchedules(options.Region), middleware.Before); err != nil {
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

// ListPrefetchSchedulesAPIClient is a client that implements the
// ListPrefetchSchedules operation.
type ListPrefetchSchedulesAPIClient interface {
	ListPrefetchSchedules(context.Context, *ListPrefetchSchedulesInput, ...func(*Options)) (*ListPrefetchSchedulesOutput, error)
}

var _ ListPrefetchSchedulesAPIClient = (*Client)(nil)

// ListPrefetchSchedulesPaginatorOptions is the paginator options for
// ListPrefetchSchedules
type ListPrefetchSchedulesPaginatorOptions struct {
	// The maximum number of prefetch schedules that you want MediaTailor to return in
	// response to the current request. If there are more than MaxResults prefetch
	// schedules, use the value of NextToken in the response to get the next page of
	// results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListPrefetchSchedulesPaginator is a paginator for ListPrefetchSchedules
type ListPrefetchSchedulesPaginator struct {
	options   ListPrefetchSchedulesPaginatorOptions
	client    ListPrefetchSchedulesAPIClient
	params    *ListPrefetchSchedulesInput
	nextToken *string
	firstPage bool
}

// NewListPrefetchSchedulesPaginator returns a new ListPrefetchSchedulesPaginator
func NewListPrefetchSchedulesPaginator(client ListPrefetchSchedulesAPIClient, params *ListPrefetchSchedulesInput, optFns ...func(*ListPrefetchSchedulesPaginatorOptions)) *ListPrefetchSchedulesPaginator {
	if params == nil {
		params = &ListPrefetchSchedulesInput{}
	}

	options := ListPrefetchSchedulesPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListPrefetchSchedulesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListPrefetchSchedulesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListPrefetchSchedules page.
func (p *ListPrefetchSchedulesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListPrefetchSchedulesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.ListPrefetchSchedules(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListPrefetchSchedules(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mediatailor",
		OperationName: "ListPrefetchSchedules",
	}
}
