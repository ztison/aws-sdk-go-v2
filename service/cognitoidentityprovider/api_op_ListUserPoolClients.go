// Code generated by smithy-go-codegen DO NOT EDIT.

package cognitoidentityprovider

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the clients that have been created for the specified user pool.
func (c *Client) ListUserPoolClients(ctx context.Context, params *ListUserPoolClientsInput, optFns ...func(*Options)) (*ListUserPoolClientsOutput, error) {
	if params == nil {
		params = &ListUserPoolClientsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListUserPoolClients", params, optFns, c.addOperationListUserPoolClientsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListUserPoolClientsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the request to list the user pool clients.
type ListUserPoolClientsInput struct {

	// The user pool ID for the user pool where you want to list user pool clients.
	//
	// This member is required.
	UserPoolId *string

	// The maximum number of results you want the request to return when listing the
	// user pool clients.
	MaxResults int32

	// An identifier that was returned from the previous call to this operation, which
	// can be used to return the next set of items in the list.
	NextToken *string

	noSmithyDocumentSerde
}

// Represents the response from the server that lists user pool clients.
type ListUserPoolClientsOutput struct {

	// An identifier that was returned from the previous call to this operation, which
	// can be used to return the next set of items in the list.
	NextToken *string

	// The user pool clients in the response that lists user pool clients.
	UserPoolClients []types.UserPoolClientDescription

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListUserPoolClientsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListUserPoolClients{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListUserPoolClients{}, middleware.After)
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
	if err = addOpListUserPoolClientsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListUserPoolClients(options.Region), middleware.Before); err != nil {
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

// ListUserPoolClientsAPIClient is a client that implements the
// ListUserPoolClients operation.
type ListUserPoolClientsAPIClient interface {
	ListUserPoolClients(context.Context, *ListUserPoolClientsInput, ...func(*Options)) (*ListUserPoolClientsOutput, error)
}

var _ ListUserPoolClientsAPIClient = (*Client)(nil)

// ListUserPoolClientsPaginatorOptions is the paginator options for
// ListUserPoolClients
type ListUserPoolClientsPaginatorOptions struct {
	// The maximum number of results you want the request to return when listing the
	// user pool clients.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListUserPoolClientsPaginator is a paginator for ListUserPoolClients
type ListUserPoolClientsPaginator struct {
	options   ListUserPoolClientsPaginatorOptions
	client    ListUserPoolClientsAPIClient
	params    *ListUserPoolClientsInput
	nextToken *string
	firstPage bool
}

// NewListUserPoolClientsPaginator returns a new ListUserPoolClientsPaginator
func NewListUserPoolClientsPaginator(client ListUserPoolClientsAPIClient, params *ListUserPoolClientsInput, optFns ...func(*ListUserPoolClientsPaginatorOptions)) *ListUserPoolClientsPaginator {
	if params == nil {
		params = &ListUserPoolClientsInput{}
	}

	options := ListUserPoolClientsPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListUserPoolClientsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListUserPoolClientsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListUserPoolClients page.
func (p *ListUserPoolClientsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListUserPoolClientsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.ListUserPoolClients(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListUserPoolClients(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cognito-idp",
		OperationName: "ListUserPoolClients",
	}
}
