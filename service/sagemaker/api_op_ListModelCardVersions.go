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
	"time"
)

// List existing versions of an Amazon SageMaker Model Card.
func (c *Client) ListModelCardVersions(ctx context.Context, params *ListModelCardVersionsInput, optFns ...func(*Options)) (*ListModelCardVersionsOutput, error) {
	if params == nil {
		params = &ListModelCardVersionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListModelCardVersions", params, optFns, c.addOperationListModelCardVersionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListModelCardVersionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListModelCardVersionsInput struct {

	// List model card versions for the model card with the specified name.
	//
	// This member is required.
	ModelCardName *string

	// Only list model card versions that were created after the time specified.
	CreationTimeAfter *time.Time

	// Only list model card versions that were created before the time specified.
	CreationTimeBefore *time.Time

	// The maximum number of model card versions to list.
	MaxResults *int32

	// Only list model card versions with the specified approval status.
	ModelCardStatus types.ModelCardStatus

	// If the response to a previous ListModelCardVersions request was truncated, the
	// response includes a NextToken . To retrieve the next set of model card versions,
	// use the token in the next request.
	NextToken *string

	// Sort listed model card versions by version. Sorts by version by default.
	SortBy types.ModelCardVersionSortBy

	// Sort model card versions by ascending or descending order.
	SortOrder types.ModelCardSortOrder

	noSmithyDocumentSerde
}

type ListModelCardVersionsOutput struct {

	// The summaries of the listed versions of the model card.
	//
	// This member is required.
	ModelCardVersionSummaryList []types.ModelCardVersionSummary

	// If the response is truncated, SageMaker returns this token. To retrieve the
	// next set of model card versions, use it in the subsequent request.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListModelCardVersionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListModelCardVersions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListModelCardVersions{}, middleware.After)
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
	if err = addOpListModelCardVersionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListModelCardVersions(options.Region), middleware.Before); err != nil {
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

// ListModelCardVersionsAPIClient is a client that implements the
// ListModelCardVersions operation.
type ListModelCardVersionsAPIClient interface {
	ListModelCardVersions(context.Context, *ListModelCardVersionsInput, ...func(*Options)) (*ListModelCardVersionsOutput, error)
}

var _ ListModelCardVersionsAPIClient = (*Client)(nil)

// ListModelCardVersionsPaginatorOptions is the paginator options for
// ListModelCardVersions
type ListModelCardVersionsPaginatorOptions struct {
	// The maximum number of model card versions to list.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListModelCardVersionsPaginator is a paginator for ListModelCardVersions
type ListModelCardVersionsPaginator struct {
	options   ListModelCardVersionsPaginatorOptions
	client    ListModelCardVersionsAPIClient
	params    *ListModelCardVersionsInput
	nextToken *string
	firstPage bool
}

// NewListModelCardVersionsPaginator returns a new ListModelCardVersionsPaginator
func NewListModelCardVersionsPaginator(client ListModelCardVersionsAPIClient, params *ListModelCardVersionsInput, optFns ...func(*ListModelCardVersionsPaginatorOptions)) *ListModelCardVersionsPaginator {
	if params == nil {
		params = &ListModelCardVersionsInput{}
	}

	options := ListModelCardVersionsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListModelCardVersionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListModelCardVersionsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListModelCardVersions page.
func (p *ListModelCardVersionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListModelCardVersionsOutput, error) {
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

	result, err := p.client.ListModelCardVersions(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListModelCardVersions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "ListModelCardVersions",
	}
}
