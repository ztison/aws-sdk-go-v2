// Code generated by smithy-go-codegen DO NOT EDIT.

package comprehend

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/comprehend/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Information about the history of a flywheel iteration. For more information
// about flywheels, see Flywheel overview (https://docs.aws.amazon.com/comprehend/latest/dg/flywheels-about.html)
// in the Amazon Comprehend Developer Guide.
func (c *Client) ListFlywheelIterationHistory(ctx context.Context, params *ListFlywheelIterationHistoryInput, optFns ...func(*Options)) (*ListFlywheelIterationHistoryOutput, error) {
	if params == nil {
		params = &ListFlywheelIterationHistoryInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListFlywheelIterationHistory", params, optFns, c.addOperationListFlywheelIterationHistoryMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListFlywheelIterationHistoryOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListFlywheelIterationHistoryInput struct {

	// The ARN of the flywheel.
	//
	// This member is required.
	FlywheelArn *string

	// Filter the flywheel iteration history based on creation time.
	Filter *types.FlywheelIterationFilter

	// Maximum number of iteration history results to return
	MaxResults *int32

	// Next token
	NextToken *string

	noSmithyDocumentSerde
}

type ListFlywheelIterationHistoryOutput struct {

	// List of flywheel iteration properties
	FlywheelIterationPropertiesList []types.FlywheelIterationProperties

	// Next token
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListFlywheelIterationHistoryMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListFlywheelIterationHistory{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListFlywheelIterationHistory{}, middleware.After)
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
	if err = addOpListFlywheelIterationHistoryValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListFlywheelIterationHistory(options.Region), middleware.Before); err != nil {
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

// ListFlywheelIterationHistoryAPIClient is a client that implements the
// ListFlywheelIterationHistory operation.
type ListFlywheelIterationHistoryAPIClient interface {
	ListFlywheelIterationHistory(context.Context, *ListFlywheelIterationHistoryInput, ...func(*Options)) (*ListFlywheelIterationHistoryOutput, error)
}

var _ ListFlywheelIterationHistoryAPIClient = (*Client)(nil)

// ListFlywheelIterationHistoryPaginatorOptions is the paginator options for
// ListFlywheelIterationHistory
type ListFlywheelIterationHistoryPaginatorOptions struct {
	// Maximum number of iteration history results to return
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListFlywheelIterationHistoryPaginator is a paginator for
// ListFlywheelIterationHistory
type ListFlywheelIterationHistoryPaginator struct {
	options   ListFlywheelIterationHistoryPaginatorOptions
	client    ListFlywheelIterationHistoryAPIClient
	params    *ListFlywheelIterationHistoryInput
	nextToken *string
	firstPage bool
}

// NewListFlywheelIterationHistoryPaginator returns a new
// ListFlywheelIterationHistoryPaginator
func NewListFlywheelIterationHistoryPaginator(client ListFlywheelIterationHistoryAPIClient, params *ListFlywheelIterationHistoryInput, optFns ...func(*ListFlywheelIterationHistoryPaginatorOptions)) *ListFlywheelIterationHistoryPaginator {
	if params == nil {
		params = &ListFlywheelIterationHistoryInput{}
	}

	options := ListFlywheelIterationHistoryPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListFlywheelIterationHistoryPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListFlywheelIterationHistoryPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListFlywheelIterationHistory page.
func (p *ListFlywheelIterationHistoryPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListFlywheelIterationHistoryOutput, error) {
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

	result, err := p.client.ListFlywheelIterationHistory(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListFlywheelIterationHistory(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "comprehend",
		OperationName: "ListFlywheelIterationHistory",
	}
}
