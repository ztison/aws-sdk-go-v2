// Code generated by smithy-go-codegen DO NOT EDIT.

package alexaforbusiness

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/alexaforbusiness/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the details of the schedules that a user configured. A download URL of
// the report associated with each schedule is returned every time this action is
// called. A new download URL is returned each time, and is valid for 24 hours.
func (c *Client) ListBusinessReportSchedules(ctx context.Context, params *ListBusinessReportSchedulesInput, optFns ...func(*Options)) (*ListBusinessReportSchedulesOutput, error) {
	if params == nil {
		params = &ListBusinessReportSchedulesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListBusinessReportSchedules", params, optFns, c.addOperationListBusinessReportSchedulesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListBusinessReportSchedulesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListBusinessReportSchedulesInput struct {

	// The maximum number of schedules listed in the call.
	MaxResults *int32

	// The token used to list the remaining schedules from the previous API call.
	NextToken *string

	noSmithyDocumentSerde
}

type ListBusinessReportSchedulesOutput struct {

	// The schedule of the reports.
	BusinessReportSchedules []types.BusinessReportSchedule

	// The token used to list the remaining schedules from the previous API call.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListBusinessReportSchedulesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListBusinessReportSchedules{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListBusinessReportSchedules{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListBusinessReportSchedules(options.Region), middleware.Before); err != nil {
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

// ListBusinessReportSchedulesAPIClient is a client that implements the
// ListBusinessReportSchedules operation.
type ListBusinessReportSchedulesAPIClient interface {
	ListBusinessReportSchedules(context.Context, *ListBusinessReportSchedulesInput, ...func(*Options)) (*ListBusinessReportSchedulesOutput, error)
}

var _ ListBusinessReportSchedulesAPIClient = (*Client)(nil)

// ListBusinessReportSchedulesPaginatorOptions is the paginator options for
// ListBusinessReportSchedules
type ListBusinessReportSchedulesPaginatorOptions struct {
	// The maximum number of schedules listed in the call.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListBusinessReportSchedulesPaginator is a paginator for
// ListBusinessReportSchedules
type ListBusinessReportSchedulesPaginator struct {
	options   ListBusinessReportSchedulesPaginatorOptions
	client    ListBusinessReportSchedulesAPIClient
	params    *ListBusinessReportSchedulesInput
	nextToken *string
	firstPage bool
}

// NewListBusinessReportSchedulesPaginator returns a new
// ListBusinessReportSchedulesPaginator
func NewListBusinessReportSchedulesPaginator(client ListBusinessReportSchedulesAPIClient, params *ListBusinessReportSchedulesInput, optFns ...func(*ListBusinessReportSchedulesPaginatorOptions)) *ListBusinessReportSchedulesPaginator {
	if params == nil {
		params = &ListBusinessReportSchedulesInput{}
	}

	options := ListBusinessReportSchedulesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListBusinessReportSchedulesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListBusinessReportSchedulesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListBusinessReportSchedules page.
func (p *ListBusinessReportSchedulesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListBusinessReportSchedulesOutput, error) {
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

	result, err := p.client.ListBusinessReportSchedules(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListBusinessReportSchedules(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "a4b",
		OperationName: "ListBusinessReportSchedules",
	}
}
