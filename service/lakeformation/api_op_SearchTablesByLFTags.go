// Code generated by smithy-go-codegen DO NOT EDIT.

package lakeformation

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lakeformation/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// This operation allows a search on TABLE resources by LFTag s. This will be used
// by admins who want to grant user permissions on certain LF-tags. Before making a
// grant, the admin can use SearchTablesByLFTags to find all resources where the
// given LFTag s are valid to verify whether the returned resources can be shared.
func (c *Client) SearchTablesByLFTags(ctx context.Context, params *SearchTablesByLFTagsInput, optFns ...func(*Options)) (*SearchTablesByLFTagsOutput, error) {
	if params == nil {
		params = &SearchTablesByLFTagsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SearchTablesByLFTags", params, optFns, c.addOperationSearchTablesByLFTagsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SearchTablesByLFTagsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SearchTablesByLFTagsInput struct {

	// A list of conditions ( LFTag structures) to search for in table resources.
	//
	// This member is required.
	Expression []types.LFTag

	// The identifier for the Data Catalog. By default, the account ID. The Data
	// Catalog is the persistent metadata store. It contains database definitions,
	// table definitions, and other control information to manage your Lake Formation
	// environment.
	CatalogId *string

	// The maximum number of results to return.
	MaxResults *int32

	// A continuation token, if this is not the first call to retrieve this list.
	NextToken *string

	noSmithyDocumentSerde
}

type SearchTablesByLFTagsOutput struct {

	// A continuation token, present if the current list segment is not the last.
	NextToken *string

	// A list of tables that meet the LF-tag conditions.
	TableList []types.TaggedTable

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationSearchTablesByLFTagsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpSearchTablesByLFTags{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpSearchTablesByLFTags{}, middleware.After)
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
	if err = addOpSearchTablesByLFTagsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSearchTablesByLFTags(options.Region), middleware.Before); err != nil {
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

// SearchTablesByLFTagsAPIClient is a client that implements the
// SearchTablesByLFTags operation.
type SearchTablesByLFTagsAPIClient interface {
	SearchTablesByLFTags(context.Context, *SearchTablesByLFTagsInput, ...func(*Options)) (*SearchTablesByLFTagsOutput, error)
}

var _ SearchTablesByLFTagsAPIClient = (*Client)(nil)

// SearchTablesByLFTagsPaginatorOptions is the paginator options for
// SearchTablesByLFTags
type SearchTablesByLFTagsPaginatorOptions struct {
	// The maximum number of results to return.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// SearchTablesByLFTagsPaginator is a paginator for SearchTablesByLFTags
type SearchTablesByLFTagsPaginator struct {
	options   SearchTablesByLFTagsPaginatorOptions
	client    SearchTablesByLFTagsAPIClient
	params    *SearchTablesByLFTagsInput
	nextToken *string
	firstPage bool
}

// NewSearchTablesByLFTagsPaginator returns a new SearchTablesByLFTagsPaginator
func NewSearchTablesByLFTagsPaginator(client SearchTablesByLFTagsAPIClient, params *SearchTablesByLFTagsInput, optFns ...func(*SearchTablesByLFTagsPaginatorOptions)) *SearchTablesByLFTagsPaginator {
	if params == nil {
		params = &SearchTablesByLFTagsInput{}
	}

	options := SearchTablesByLFTagsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &SearchTablesByLFTagsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *SearchTablesByLFTagsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next SearchTablesByLFTags page.
func (p *SearchTablesByLFTagsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*SearchTablesByLFTagsOutput, error) {
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

	result, err := p.client.SearchTablesByLFTags(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opSearchTablesByLFTags(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lakeformation",
		OperationName: "SearchTablesByLFTags",
	}
}
