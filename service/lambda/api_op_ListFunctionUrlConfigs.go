// Code generated by smithy-go-codegen DO NOT EDIT.

package lambda

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lambda/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of Lambda function URLs for the specified function.
func (c *Client) ListFunctionUrlConfigs(ctx context.Context, params *ListFunctionUrlConfigsInput, optFns ...func(*Options)) (*ListFunctionUrlConfigsOutput, error) {
	if params == nil {
		params = &ListFunctionUrlConfigsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListFunctionUrlConfigs", params, optFns, c.addOperationListFunctionUrlConfigsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListFunctionUrlConfigsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListFunctionUrlConfigsInput struct {

	// The name of the Lambda function. Name formats
	//   - Function name – my-function .
	//   - Function ARN – arn:aws:lambda:us-west-2:123456789012:function:my-function .
	//   - Partial ARN – 123456789012:function:my-function .
	// The length constraint applies only to the full ARN. If you specify only the
	// function name, it is limited to 64 characters in length.
	//
	// This member is required.
	FunctionName *string

	// Specify the pagination token that's returned by a previous request to retrieve
	// the next page of results.
	Marker *string

	// The maximum number of function URLs to return in the response. Note that
	// ListFunctionUrlConfigs returns a maximum of 50 items in each response, even if
	// you set the number higher.
	MaxItems *int32

	noSmithyDocumentSerde
}

type ListFunctionUrlConfigsOutput struct {

	// A list of function URL configurations.
	//
	// This member is required.
	FunctionUrlConfigs []types.FunctionUrlConfig

	// The pagination token that's included if more results are available.
	NextMarker *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListFunctionUrlConfigsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListFunctionUrlConfigs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListFunctionUrlConfigs{}, middleware.After)
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
	if err = addOpListFunctionUrlConfigsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListFunctionUrlConfigs(options.Region), middleware.Before); err != nil {
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

// ListFunctionUrlConfigsAPIClient is a client that implements the
// ListFunctionUrlConfigs operation.
type ListFunctionUrlConfigsAPIClient interface {
	ListFunctionUrlConfigs(context.Context, *ListFunctionUrlConfigsInput, ...func(*Options)) (*ListFunctionUrlConfigsOutput, error)
}

var _ ListFunctionUrlConfigsAPIClient = (*Client)(nil)

// ListFunctionUrlConfigsPaginatorOptions is the paginator options for
// ListFunctionUrlConfigs
type ListFunctionUrlConfigsPaginatorOptions struct {
	// The maximum number of function URLs to return in the response. Note that
	// ListFunctionUrlConfigs returns a maximum of 50 items in each response, even if
	// you set the number higher.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListFunctionUrlConfigsPaginator is a paginator for ListFunctionUrlConfigs
type ListFunctionUrlConfigsPaginator struct {
	options   ListFunctionUrlConfigsPaginatorOptions
	client    ListFunctionUrlConfigsAPIClient
	params    *ListFunctionUrlConfigsInput
	nextToken *string
	firstPage bool
}

// NewListFunctionUrlConfigsPaginator returns a new ListFunctionUrlConfigsPaginator
func NewListFunctionUrlConfigsPaginator(client ListFunctionUrlConfigsAPIClient, params *ListFunctionUrlConfigsInput, optFns ...func(*ListFunctionUrlConfigsPaginatorOptions)) *ListFunctionUrlConfigsPaginator {
	if params == nil {
		params = &ListFunctionUrlConfigsInput{}
	}

	options := ListFunctionUrlConfigsPaginatorOptions{}
	if params.MaxItems != nil {
		options.Limit = *params.MaxItems
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListFunctionUrlConfigsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.Marker,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListFunctionUrlConfigsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListFunctionUrlConfigs page.
func (p *ListFunctionUrlConfigsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListFunctionUrlConfigsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.Marker = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxItems = limit

	result, err := p.client.ListFunctionUrlConfigs(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextMarker

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opListFunctionUrlConfigs(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lambda",
		OperationName: "ListFunctionUrlConfigs",
	}
}
