// Code generated by smithy-go-codegen DO NOT EDIT.

package identitystore

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/identitystore/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists all users in the identity store. Returns a paginated list of complete User
// objects. Filtering for a User by the UserName attribute is deprecated. Instead,
// use the GetUserId API action.
func (c *Client) ListUsers(ctx context.Context, params *ListUsersInput, optFns ...func(*Options)) (*ListUsersOutput, error) {
	if params == nil {
		params = &ListUsersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListUsers", params, optFns, c.addOperationListUsersMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListUsersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListUsersInput struct {

	// The globally unique identifier for the identity store, such as d-1234567890 . In
	// this example, d- is a fixed prefix, and 1234567890 is a randomly generated
	// string that contains numbers and lower case letters. This value is generated at
	// the time that a new identity store is created.
	//
	// This member is required.
	IdentityStoreId *string

	// A list of Filter objects, which is used in the ListUsers and ListGroups
	// requests.
	//
	// Deprecated: Using filters with ListUsers API is deprecated, please use
	// GetGroupId API instead.
	Filters []types.Filter

	// The maximum number of results to be returned per request. This parameter is
	// used in the ListUsers and ListGroups requests to specify how many results to
	// return in one page. The length limit is 50 characters.
	MaxResults *int32

	// The pagination token used for the ListUsers and ListGroups API operations. This
	// value is generated by the identity store service. It is returned in the API
	// response if the total results are more than the size of one page. This token is
	// also returned when it is used in the API request to search for the next page.
	NextToken *string

	noSmithyDocumentSerde
}

type ListUsersOutput struct {

	// A list of User objects in the identity store.
	//
	// This member is required.
	Users []types.User

	// The pagination token used for the ListUsers and ListGroups API operations. This
	// value is generated by the identity store service. It is returned in the API
	// response if the total results are more than the size of one page. This token is
	// also returned when it is used in the API request to search for the next page.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListUsersMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListUsers{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListUsers{}, middleware.After)
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
	if err = addOpListUsersValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListUsers(options.Region), middleware.Before); err != nil {
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

// ListUsersAPIClient is a client that implements the ListUsers operation.
type ListUsersAPIClient interface {
	ListUsers(context.Context, *ListUsersInput, ...func(*Options)) (*ListUsersOutput, error)
}

var _ ListUsersAPIClient = (*Client)(nil)

// ListUsersPaginatorOptions is the paginator options for ListUsers
type ListUsersPaginatorOptions struct {
	// The maximum number of results to be returned per request. This parameter is
	// used in the ListUsers and ListGroups requests to specify how many results to
	// return in one page. The length limit is 50 characters.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListUsersPaginator is a paginator for ListUsers
type ListUsersPaginator struct {
	options   ListUsersPaginatorOptions
	client    ListUsersAPIClient
	params    *ListUsersInput
	nextToken *string
	firstPage bool
}

// NewListUsersPaginator returns a new ListUsersPaginator
func NewListUsersPaginator(client ListUsersAPIClient, params *ListUsersInput, optFns ...func(*ListUsersPaginatorOptions)) *ListUsersPaginator {
	if params == nil {
		params = &ListUsersInput{}
	}

	options := ListUsersPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListUsersPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListUsersPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListUsers page.
func (p *ListUsersPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListUsersOutput, error) {
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

	result, err := p.client.ListUsers(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListUsers(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "identitystore",
		OperationName: "ListUsers",
	}
}
