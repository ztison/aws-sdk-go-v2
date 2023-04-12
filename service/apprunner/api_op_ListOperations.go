// Code generated by smithy-go-codegen DO NOT EDIT.

package apprunner

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/apprunner/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Return a list of operations that occurred on an App Runner service. The
// resulting list of OperationSummary objects is sorted in reverse chronological
// order. The first object on the list represents the last started operation.
func (c *Client) ListOperations(ctx context.Context, params *ListOperationsInput, optFns ...func(*Options)) (*ListOperationsOutput, error) {
	if params == nil {
		params = &ListOperationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListOperations", params, optFns, c.addOperationListOperationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListOperationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListOperationsInput struct {

	// The Amazon Resource Name (ARN) of the App Runner service that you want a list
	// of operations for.
	//
	// This member is required.
	ServiceArn *string

	// The maximum number of results to include in each response (result page). It's
	// used for a paginated request. If you don't specify MaxResults , the request
	// retrieves all available results in a single response.
	MaxResults *int32

	// A token from a previous result page. It's used for a paginated request. The
	// request retrieves the next result page. All other parameter values must be
	// identical to the ones specified in the initial request. If you don't specify
	// NextToken , the request retrieves the first result page.
	NextToken *string

	noSmithyDocumentSerde
}

type ListOperationsOutput struct {

	// The token that you can pass in a subsequent request to get the next result
	// page. It's returned in a paginated request.
	NextToken *string

	// A list of operation summary information records. In a paginated request, the
	// request returns up to MaxResults records for each call.
	OperationSummaryList []types.OperationSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListOperationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpListOperations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpListOperations{}, middleware.After)
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
	if err = addOpListOperationsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListOperations(options.Region), middleware.Before); err != nil {
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

// ListOperationsAPIClient is a client that implements the ListOperations
// operation.
type ListOperationsAPIClient interface {
	ListOperations(context.Context, *ListOperationsInput, ...func(*Options)) (*ListOperationsOutput, error)
}

var _ ListOperationsAPIClient = (*Client)(nil)

// ListOperationsPaginatorOptions is the paginator options for ListOperations
type ListOperationsPaginatorOptions struct {
	// The maximum number of results to include in each response (result page). It's
	// used for a paginated request. If you don't specify MaxResults , the request
	// retrieves all available results in a single response.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListOperationsPaginator is a paginator for ListOperations
type ListOperationsPaginator struct {
	options   ListOperationsPaginatorOptions
	client    ListOperationsAPIClient
	params    *ListOperationsInput
	nextToken *string
	firstPage bool
}

// NewListOperationsPaginator returns a new ListOperationsPaginator
func NewListOperationsPaginator(client ListOperationsAPIClient, params *ListOperationsInput, optFns ...func(*ListOperationsPaginatorOptions)) *ListOperationsPaginator {
	if params == nil {
		params = &ListOperationsInput{}
	}

	options := ListOperationsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListOperationsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListOperationsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListOperations page.
func (p *ListOperationsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListOperationsOutput, error) {
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

	result, err := p.client.ListOperations(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListOperations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "apprunner",
		OperationName: "ListOperations",
	}
}
