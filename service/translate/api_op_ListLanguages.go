// Code generated by smithy-go-codegen DO NOT EDIT.

package translate

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/translate/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Provides a list of languages (RFC-5646 codes and names) that Amazon Translate
// supports.
func (c *Client) ListLanguages(ctx context.Context, params *ListLanguagesInput, optFns ...func(*Options)) (*ListLanguagesOutput, error) {
	if params == nil {
		params = &ListLanguagesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListLanguages", params, optFns, c.addOperationListLanguagesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListLanguagesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListLanguagesInput struct {

	// The language code for the language to use to display the language names in the
	// response. The language code is en by default.
	DisplayLanguageCode types.DisplayLanguageCode

	// The maximum number of results to return in each response.
	MaxResults *int32

	// Include the NextToken value to fetch the next group of supported languages.
	NextToken *string

	noSmithyDocumentSerde
}

type ListLanguagesOutput struct {

	// The language code passed in with the request.
	DisplayLanguageCode types.DisplayLanguageCode

	// The list of supported languages.
	Languages []types.Language

	// If the response does not include all remaining results, use the NextToken in
	// the next request to fetch the next group of supported languages.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListLanguagesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListLanguages{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListLanguages{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListLanguages(options.Region), middleware.Before); err != nil {
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

// ListLanguagesAPIClient is a client that implements the ListLanguages operation.
type ListLanguagesAPIClient interface {
	ListLanguages(context.Context, *ListLanguagesInput, ...func(*Options)) (*ListLanguagesOutput, error)
}

var _ ListLanguagesAPIClient = (*Client)(nil)

// ListLanguagesPaginatorOptions is the paginator options for ListLanguages
type ListLanguagesPaginatorOptions struct {
	// The maximum number of results to return in each response.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListLanguagesPaginator is a paginator for ListLanguages
type ListLanguagesPaginator struct {
	options   ListLanguagesPaginatorOptions
	client    ListLanguagesAPIClient
	params    *ListLanguagesInput
	nextToken *string
	firstPage bool
}

// NewListLanguagesPaginator returns a new ListLanguagesPaginator
func NewListLanguagesPaginator(client ListLanguagesAPIClient, params *ListLanguagesInput, optFns ...func(*ListLanguagesPaginatorOptions)) *ListLanguagesPaginator {
	if params == nil {
		params = &ListLanguagesInput{}
	}

	options := ListLanguagesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListLanguagesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListLanguagesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListLanguages page.
func (p *ListLanguagesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListLanguagesOutput, error) {
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

	result, err := p.client.ListLanguages(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListLanguages(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "translate",
		OperationName: "ListLanguages",
	}
}
