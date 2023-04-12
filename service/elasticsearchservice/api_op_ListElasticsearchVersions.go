// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticsearchservice

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// List all supported Elasticsearch versions
func (c *Client) ListElasticsearchVersions(ctx context.Context, params *ListElasticsearchVersionsInput, optFns ...func(*Options)) (*ListElasticsearchVersionsOutput, error) {
	if params == nil {
		params = &ListElasticsearchVersionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListElasticsearchVersions", params, optFns, c.addOperationListElasticsearchVersionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListElasticsearchVersionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Container for the parameters to the ListElasticsearchVersions operation. Use
// MaxResults to control the maximum number of results to retrieve in a single
// call. Use NextToken in response to retrieve more results. If the received
// response does not contain a NextToken, then there are no more results to
// retrieve.
type ListElasticsearchVersionsInput struct {

	// Set this value to limit the number of results returned. Value provided must be
	// greater than 10 else it wont be honored.
	MaxResults int32

	// Paginated APIs accepts NextToken input to returns next page results and
	// provides a NextToken output in the response which can be used by the client to
	// retrieve more results.
	NextToken *string

	noSmithyDocumentSerde
}

// Container for the parameters for response received from
// ListElasticsearchVersions operation.
type ListElasticsearchVersionsOutput struct {

	// List of supported elastic search versions.
	ElasticsearchVersions []string

	// Paginated APIs accepts NextToken input to returns next page results and
	// provides a NextToken output in the response which can be used by the client to
	// retrieve more results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListElasticsearchVersionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListElasticsearchVersions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListElasticsearchVersions{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListElasticsearchVersions(options.Region), middleware.Before); err != nil {
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

// ListElasticsearchVersionsAPIClient is a client that implements the
// ListElasticsearchVersions operation.
type ListElasticsearchVersionsAPIClient interface {
	ListElasticsearchVersions(context.Context, *ListElasticsearchVersionsInput, ...func(*Options)) (*ListElasticsearchVersionsOutput, error)
}

var _ ListElasticsearchVersionsAPIClient = (*Client)(nil)

// ListElasticsearchVersionsPaginatorOptions is the paginator options for
// ListElasticsearchVersions
type ListElasticsearchVersionsPaginatorOptions struct {
	// Set this value to limit the number of results returned. Value provided must be
	// greater than 10 else it wont be honored.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListElasticsearchVersionsPaginator is a paginator for ListElasticsearchVersions
type ListElasticsearchVersionsPaginator struct {
	options   ListElasticsearchVersionsPaginatorOptions
	client    ListElasticsearchVersionsAPIClient
	params    *ListElasticsearchVersionsInput
	nextToken *string
	firstPage bool
}

// NewListElasticsearchVersionsPaginator returns a new
// ListElasticsearchVersionsPaginator
func NewListElasticsearchVersionsPaginator(client ListElasticsearchVersionsAPIClient, params *ListElasticsearchVersionsInput, optFns ...func(*ListElasticsearchVersionsPaginatorOptions)) *ListElasticsearchVersionsPaginator {
	if params == nil {
		params = &ListElasticsearchVersionsInput{}
	}

	options := ListElasticsearchVersionsPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListElasticsearchVersionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListElasticsearchVersionsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListElasticsearchVersions page.
func (p *ListElasticsearchVersionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListElasticsearchVersionsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.ListElasticsearchVersions(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListElasticsearchVersions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "es",
		OperationName: "ListElasticsearchVersions",
	}
}
