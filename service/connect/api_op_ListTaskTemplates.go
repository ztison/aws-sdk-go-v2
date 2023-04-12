// Code generated by smithy-go-codegen DO NOT EDIT.

package connect

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/connect/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists task templates for the specified Amazon Connect instance.
func (c *Client) ListTaskTemplates(ctx context.Context, params *ListTaskTemplatesInput, optFns ...func(*Options)) (*ListTaskTemplatesOutput, error) {
	if params == nil {
		params = &ListTaskTemplatesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListTaskTemplates", params, optFns, c.addOperationListTaskTemplatesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListTaskTemplatesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListTaskTemplatesInput struct {

	// The identifier of the Amazon Connect instance. You can find the instance ID (https://docs.aws.amazon.com/connect/latest/adminguide/find-instance-arn.html)
	// in the Amazon Resource Name (ARN) of the instance.
	//
	// This member is required.
	InstanceId *string

	// The maximum number of results to return per page. It is not expected that you
	// set this.
	MaxResults *int32

	// The name of the task template.
	Name *string

	// The token for the next set of results. Use the value returned in the previous
	// response in the next request to retrieve the next set of results. It is not
	// expected that you set this because the value returned in the previous response
	// is always null.
	NextToken *string

	// Marks a template as ACTIVE or INACTIVE for a task to refer to it. Tasks can
	// only be created from ACTIVE templates. If a template is marked as INACTIVE ,
	// then a task that refers to this template cannot be created.
	Status types.TaskTemplateStatus

	noSmithyDocumentSerde
}

type ListTaskTemplatesOutput struct {

	// If there are additional results, this is the token for the next set of results.
	// This is always returned as a null in the response.
	NextToken *string

	// Provides details about a list of task templates belonging to an instance.
	TaskTemplates []types.TaskTemplateMetadata

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListTaskTemplatesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListTaskTemplates{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListTaskTemplates{}, middleware.After)
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
	if err = addOpListTaskTemplatesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListTaskTemplates(options.Region), middleware.Before); err != nil {
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

// ListTaskTemplatesAPIClient is a client that implements the ListTaskTemplates
// operation.
type ListTaskTemplatesAPIClient interface {
	ListTaskTemplates(context.Context, *ListTaskTemplatesInput, ...func(*Options)) (*ListTaskTemplatesOutput, error)
}

var _ ListTaskTemplatesAPIClient = (*Client)(nil)

// ListTaskTemplatesPaginatorOptions is the paginator options for ListTaskTemplates
type ListTaskTemplatesPaginatorOptions struct {
	// The maximum number of results to return per page. It is not expected that you
	// set this.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListTaskTemplatesPaginator is a paginator for ListTaskTemplates
type ListTaskTemplatesPaginator struct {
	options   ListTaskTemplatesPaginatorOptions
	client    ListTaskTemplatesAPIClient
	params    *ListTaskTemplatesInput
	nextToken *string
	firstPage bool
}

// NewListTaskTemplatesPaginator returns a new ListTaskTemplatesPaginator
func NewListTaskTemplatesPaginator(client ListTaskTemplatesAPIClient, params *ListTaskTemplatesInput, optFns ...func(*ListTaskTemplatesPaginatorOptions)) *ListTaskTemplatesPaginator {
	if params == nil {
		params = &ListTaskTemplatesInput{}
	}

	options := ListTaskTemplatesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListTaskTemplatesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListTaskTemplatesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListTaskTemplates page.
func (p *ListTaskTemplatesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListTaskTemplatesOutput, error) {
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

	result, err := p.client.ListTaskTemplates(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListTaskTemplates(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "connect",
		OperationName: "ListTaskTemplates",
	}
}
