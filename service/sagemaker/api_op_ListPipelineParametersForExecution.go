// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets a list of parameters for a pipeline execution.
func (c *Client) ListPipelineParametersForExecution(ctx context.Context, params *ListPipelineParametersForExecutionInput, optFns ...func(*Options)) (*ListPipelineParametersForExecutionOutput, error) {
	if params == nil {
		params = &ListPipelineParametersForExecutionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListPipelineParametersForExecution", params, optFns, c.addOperationListPipelineParametersForExecutionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListPipelineParametersForExecutionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListPipelineParametersForExecutionInput struct {

	// The Amazon Resource Name (ARN) of the pipeline execution.
	//
	// This member is required.
	PipelineExecutionArn *string

	// The maximum number of parameters to return in the response.
	MaxResults *int32

	// If the result of the previous ListPipelineParametersForExecution request was
	// truncated, the response includes a NextToken . To retrieve the next set of
	// parameters, use the token in the next request.
	NextToken *string

	noSmithyDocumentSerde
}

type ListPipelineParametersForExecutionOutput struct {

	// If the result of the previous ListPipelineParametersForExecution request was
	// truncated, the response includes a NextToken . To retrieve the next set of
	// parameters, use the token in the next request.
	NextToken *string

	// Contains a list of pipeline parameters. This list can be empty.
	PipelineParameters []types.Parameter

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListPipelineParametersForExecutionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListPipelineParametersForExecution{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListPipelineParametersForExecution{}, middleware.After)
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
	if err = addOpListPipelineParametersForExecutionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListPipelineParametersForExecution(options.Region), middleware.Before); err != nil {
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

// ListPipelineParametersForExecutionAPIClient is a client that implements the
// ListPipelineParametersForExecution operation.
type ListPipelineParametersForExecutionAPIClient interface {
	ListPipelineParametersForExecution(context.Context, *ListPipelineParametersForExecutionInput, ...func(*Options)) (*ListPipelineParametersForExecutionOutput, error)
}

var _ ListPipelineParametersForExecutionAPIClient = (*Client)(nil)

// ListPipelineParametersForExecutionPaginatorOptions is the paginator options for
// ListPipelineParametersForExecution
type ListPipelineParametersForExecutionPaginatorOptions struct {
	// The maximum number of parameters to return in the response.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListPipelineParametersForExecutionPaginator is a paginator for
// ListPipelineParametersForExecution
type ListPipelineParametersForExecutionPaginator struct {
	options   ListPipelineParametersForExecutionPaginatorOptions
	client    ListPipelineParametersForExecutionAPIClient
	params    *ListPipelineParametersForExecutionInput
	nextToken *string
	firstPage bool
}

// NewListPipelineParametersForExecutionPaginator returns a new
// ListPipelineParametersForExecutionPaginator
func NewListPipelineParametersForExecutionPaginator(client ListPipelineParametersForExecutionAPIClient, params *ListPipelineParametersForExecutionInput, optFns ...func(*ListPipelineParametersForExecutionPaginatorOptions)) *ListPipelineParametersForExecutionPaginator {
	if params == nil {
		params = &ListPipelineParametersForExecutionInput{}
	}

	options := ListPipelineParametersForExecutionPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListPipelineParametersForExecutionPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListPipelineParametersForExecutionPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListPipelineParametersForExecution page.
func (p *ListPipelineParametersForExecutionPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListPipelineParametersForExecutionOutput, error) {
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

	result, err := p.client.ListPipelineParametersForExecution(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListPipelineParametersForExecution(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "ListPipelineParametersForExecution",
	}
}
