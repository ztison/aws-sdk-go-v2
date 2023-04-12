// Code generated by smithy-go-codegen DO NOT EDIT.

package backup

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/backup/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Returns details about your report jobs.
func (c *Client) ListReportJobs(ctx context.Context, params *ListReportJobsInput, optFns ...func(*Options)) (*ListReportJobsOutput, error) {
	if params == nil {
		params = &ListReportJobsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListReportJobs", params, optFns, c.addOperationListReportJobsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListReportJobsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListReportJobsInput struct {

	// Returns only report jobs that were created after the date and time specified in
	// Unix format and Coordinated Universal Time (UTC). For example, the value
	// 1516925490 represents Friday, January 26, 2018 12:11:30 AM.
	ByCreationAfter *time.Time

	// Returns only report jobs that were created before the date and time specified
	// in Unix format and Coordinated Universal Time (UTC). For example, the value
	// 1516925490 represents Friday, January 26, 2018 12:11:30 AM.
	ByCreationBefore *time.Time

	// Returns only report jobs with the specified report plan name.
	ByReportPlanName *string

	// Returns only report jobs that are in the specified status. The statuses are:
	// CREATED | RUNNING | COMPLETED | FAILED
	ByStatus *string

	// The number of desired results from 1 to 1000. Optional. If unspecified, the
	// query will return 1 MB of data.
	MaxResults *int32

	// An identifier that was returned from the previous call to this operation, which
	// can be used to return the next set of items in the list.
	NextToken *string

	noSmithyDocumentSerde
}

type ListReportJobsOutput struct {

	// An identifier that was returned from the previous call to this operation, which
	// can be used to return the next set of items in the list.
	NextToken *string

	// Details about your report jobs in JSON format.
	ReportJobs []types.ReportJob

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListReportJobsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListReportJobs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListReportJobs{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListReportJobs(options.Region), middleware.Before); err != nil {
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

// ListReportJobsAPIClient is a client that implements the ListReportJobs
// operation.
type ListReportJobsAPIClient interface {
	ListReportJobs(context.Context, *ListReportJobsInput, ...func(*Options)) (*ListReportJobsOutput, error)
}

var _ ListReportJobsAPIClient = (*Client)(nil)

// ListReportJobsPaginatorOptions is the paginator options for ListReportJobs
type ListReportJobsPaginatorOptions struct {
	// The number of desired results from 1 to 1000. Optional. If unspecified, the
	// query will return 1 MB of data.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListReportJobsPaginator is a paginator for ListReportJobs
type ListReportJobsPaginator struct {
	options   ListReportJobsPaginatorOptions
	client    ListReportJobsAPIClient
	params    *ListReportJobsInput
	nextToken *string
	firstPage bool
}

// NewListReportJobsPaginator returns a new ListReportJobsPaginator
func NewListReportJobsPaginator(client ListReportJobsAPIClient, params *ListReportJobsInput, optFns ...func(*ListReportJobsPaginatorOptions)) *ListReportJobsPaginator {
	if params == nil {
		params = &ListReportJobsInput{}
	}

	options := ListReportJobsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListReportJobsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListReportJobsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListReportJobs page.
func (p *ListReportJobsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListReportJobsOutput, error) {
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

	result, err := p.client.ListReportJobs(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListReportJobs(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "backup",
		OperationName: "ListReportJobs",
	}
}
