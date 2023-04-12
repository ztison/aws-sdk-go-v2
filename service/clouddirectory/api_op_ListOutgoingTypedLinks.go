// Code generated by smithy-go-codegen DO NOT EDIT.

package clouddirectory

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/clouddirectory/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a paginated list of all the outgoing TypedLinkSpecifier information for
// an object. It also supports filtering by typed link facet and identity
// attributes. For more information, see Typed Links (https://docs.aws.amazon.com/clouddirectory/latest/developerguide/directory_objects_links.html#directory_objects_links_typedlink)
// .
func (c *Client) ListOutgoingTypedLinks(ctx context.Context, params *ListOutgoingTypedLinksInput, optFns ...func(*Options)) (*ListOutgoingTypedLinksOutput, error) {
	if params == nil {
		params = &ListOutgoingTypedLinksInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListOutgoingTypedLinks", params, optFns, c.addOperationListOutgoingTypedLinksMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListOutgoingTypedLinksOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListOutgoingTypedLinksInput struct {

	// The Amazon Resource Name (ARN) of the directory where you want to list the
	// typed links.
	//
	// This member is required.
	DirectoryArn *string

	// A reference that identifies the object whose attributes will be listed.
	//
	// This member is required.
	ObjectReference *types.ObjectReference

	// The consistency level to execute the request at.
	ConsistencyLevel types.ConsistencyLevel

	// Provides range filters for multiple attributes. When providing ranges to typed
	// link selection, any inexact ranges must be specified at the end. Any attributes
	// that do not have a range specified are presumed to match the entire range.
	FilterAttributeRanges []types.TypedLinkAttributeRange

	// Filters are interpreted in the order of the attributes defined on the typed
	// link facet, not the order they are supplied to any API calls.
	FilterTypedLink *types.TypedLinkSchemaAndFacetName

	// The maximum number of results to retrieve.
	MaxResults *int32

	// The pagination token.
	NextToken *string

	noSmithyDocumentSerde
}

type ListOutgoingTypedLinksOutput struct {

	// The pagination token.
	NextToken *string

	// Returns a typed link specifier as output.
	TypedLinkSpecifiers []types.TypedLinkSpecifier

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListOutgoingTypedLinksMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListOutgoingTypedLinks{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListOutgoingTypedLinks{}, middleware.After)
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
	if err = addOpListOutgoingTypedLinksValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListOutgoingTypedLinks(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListOutgoingTypedLinks(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "clouddirectory",
		OperationName: "ListOutgoingTypedLinks",
	}
}
