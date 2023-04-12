// Code generated by smithy-go-codegen DO NOT EDIT.

package efs

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/efs/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns the description of a specific Amazon EFS file system if either the file
// system CreationToken or the FileSystemId is provided. Otherwise, it returns
// descriptions of all file systems owned by the caller's Amazon Web Services
// account in the Amazon Web Services Region of the endpoint that you're calling.
// When retrieving all file system descriptions, you can optionally specify the
// MaxItems parameter to limit the number of descriptions in a response. This
// number is automatically set to 100. If more file system descriptions remain,
// Amazon EFS returns a NextMarker , an opaque token, in the response. In this
// case, you should send a subsequent request with the Marker request parameter
// set to the value of NextMarker . To retrieve a list of your file system
// descriptions, this operation is used in an iterative process, where
// DescribeFileSystems is called first without the Marker and then the operation
// continues to call it with the Marker parameter set to the value of the
// NextMarker from the previous response until the response has no NextMarker . The
// order of file systems returned in the response of one DescribeFileSystems call
// and the order of file systems returned across the responses of a multi-call
// iteration is unspecified. This operation requires permissions for the
// elasticfilesystem:DescribeFileSystems action.
func (c *Client) DescribeFileSystems(ctx context.Context, params *DescribeFileSystemsInput, optFns ...func(*Options)) (*DescribeFileSystemsOutput, error) {
	if params == nil {
		params = &DescribeFileSystemsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeFileSystems", params, optFns, c.addOperationDescribeFileSystemsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeFileSystemsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeFileSystemsInput struct {

	// (Optional) Restricts the list to the file system with this creation token
	// (String). You specify a creation token when you create an Amazon EFS file
	// system.
	CreationToken *string

	// (Optional) ID of the file system whose description you want to retrieve
	// (String).
	FileSystemId *string

	// (Optional) Opaque pagination token returned from a previous DescribeFileSystems
	// operation (String). If present, specifies to continue the list from where the
	// returning call had left off.
	Marker *string

	// (Optional) Specifies the maximum number of file systems to return in the
	// response (integer). This number is automatically set to 100. The response is
	// paginated at 100 per page if you have more than 100 file systems.
	MaxItems *int32

	noSmithyDocumentSerde
}

type DescribeFileSystemsOutput struct {

	// An array of file system descriptions.
	FileSystems []types.FileSystemDescription

	// Present if provided by caller in the request (String).
	Marker *string

	// Present if there are more file systems than returned in the response (String).
	// You can use the NextMarker in the subsequent request to fetch the descriptions.
	NextMarker *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeFileSystemsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeFileSystems{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeFileSystems{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeFileSystems(options.Region), middleware.Before); err != nil {
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

// DescribeFileSystemsAPIClient is a client that implements the
// DescribeFileSystems operation.
type DescribeFileSystemsAPIClient interface {
	DescribeFileSystems(context.Context, *DescribeFileSystemsInput, ...func(*Options)) (*DescribeFileSystemsOutput, error)
}

var _ DescribeFileSystemsAPIClient = (*Client)(nil)

// DescribeFileSystemsPaginatorOptions is the paginator options for
// DescribeFileSystems
type DescribeFileSystemsPaginatorOptions struct {
	// (Optional) Specifies the maximum number of file systems to return in the
	// response (integer). This number is automatically set to 100. The response is
	// paginated at 100 per page if you have more than 100 file systems.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeFileSystemsPaginator is a paginator for DescribeFileSystems
type DescribeFileSystemsPaginator struct {
	options   DescribeFileSystemsPaginatorOptions
	client    DescribeFileSystemsAPIClient
	params    *DescribeFileSystemsInput
	nextToken *string
	firstPage bool
}

// NewDescribeFileSystemsPaginator returns a new DescribeFileSystemsPaginator
func NewDescribeFileSystemsPaginator(client DescribeFileSystemsAPIClient, params *DescribeFileSystemsInput, optFns ...func(*DescribeFileSystemsPaginatorOptions)) *DescribeFileSystemsPaginator {
	if params == nil {
		params = &DescribeFileSystemsInput{}
	}

	options := DescribeFileSystemsPaginatorOptions{}
	if params.MaxItems != nil {
		options.Limit = *params.MaxItems
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeFileSystemsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.Marker,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeFileSystemsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeFileSystems page.
func (p *DescribeFileSystemsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeFileSystemsOutput, error) {
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

	result, err := p.client.DescribeFileSystems(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeFileSystems(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticfilesystem",
		OperationName: "DescribeFileSystems",
	}
}
