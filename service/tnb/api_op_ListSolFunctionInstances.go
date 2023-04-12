// Code generated by smithy-go-codegen DO NOT EDIT.

package tnb

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/tnb/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists network function instances. A network function instance is a function in
// a function package .
func (c *Client) ListSolFunctionInstances(ctx context.Context, params *ListSolFunctionInstancesInput, optFns ...func(*Options)) (*ListSolFunctionInstancesOutput, error) {
	if params == nil {
		params = &ListSolFunctionInstancesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListSolFunctionInstances", params, optFns, c.addOperationListSolFunctionInstancesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListSolFunctionInstancesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListSolFunctionInstancesInput struct {

	// The maximum number of results to include in the response.
	MaxResults *int32

	// The token for the next page of results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListSolFunctionInstancesOutput struct {

	// Network function instances.
	FunctionInstances []types.ListSolFunctionInstanceInfo

	// The token to use to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListSolFunctionInstancesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListSolFunctionInstances{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListSolFunctionInstances{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListSolFunctionInstances(options.Region), middleware.Before); err != nil {
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

// ListSolFunctionInstancesAPIClient is a client that implements the
// ListSolFunctionInstances operation.
type ListSolFunctionInstancesAPIClient interface {
	ListSolFunctionInstances(context.Context, *ListSolFunctionInstancesInput, ...func(*Options)) (*ListSolFunctionInstancesOutput, error)
}

var _ ListSolFunctionInstancesAPIClient = (*Client)(nil)

// ListSolFunctionInstancesPaginatorOptions is the paginator options for
// ListSolFunctionInstances
type ListSolFunctionInstancesPaginatorOptions struct {
	// The maximum number of results to include in the response.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListSolFunctionInstancesPaginator is a paginator for ListSolFunctionInstances
type ListSolFunctionInstancesPaginator struct {
	options   ListSolFunctionInstancesPaginatorOptions
	client    ListSolFunctionInstancesAPIClient
	params    *ListSolFunctionInstancesInput
	nextToken *string
	firstPage bool
}

// NewListSolFunctionInstancesPaginator returns a new
// ListSolFunctionInstancesPaginator
func NewListSolFunctionInstancesPaginator(client ListSolFunctionInstancesAPIClient, params *ListSolFunctionInstancesInput, optFns ...func(*ListSolFunctionInstancesPaginatorOptions)) *ListSolFunctionInstancesPaginator {
	if params == nil {
		params = &ListSolFunctionInstancesInput{}
	}

	options := ListSolFunctionInstancesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListSolFunctionInstancesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListSolFunctionInstancesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListSolFunctionInstances page.
func (p *ListSolFunctionInstancesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListSolFunctionInstancesOutput, error) {
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

	result, err := p.client.ListSolFunctionInstances(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListSolFunctionInstances(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "tnb",
		OperationName: "ListSolFunctionInstances",
	}
}
