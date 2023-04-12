// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticache

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/elasticache/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a new cache subnet group. Use this parameter only when you are creating
// a cluster in an Amazon Virtual Private Cloud (Amazon VPC).
func (c *Client) CreateCacheSubnetGroup(ctx context.Context, params *CreateCacheSubnetGroupInput, optFns ...func(*Options)) (*CreateCacheSubnetGroupOutput, error) {
	if params == nil {
		params = &CreateCacheSubnetGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateCacheSubnetGroup", params, optFns, c.addOperationCreateCacheSubnetGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateCacheSubnetGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input of a CreateCacheSubnetGroup operation.
type CreateCacheSubnetGroupInput struct {

	// A description for the cache subnet group.
	//
	// This member is required.
	CacheSubnetGroupDescription *string

	// A name for the cache subnet group. This value is stored as a lowercase string.
	// Constraints: Must contain no more than 255 alphanumeric characters or hyphens.
	// Example: mysubnetgroup
	//
	// This member is required.
	CacheSubnetGroupName *string

	// A list of VPC subnet IDs for the cache subnet group.
	//
	// This member is required.
	SubnetIds []string

	// A list of tags to be added to this resource. A tag is a key-value pair. A tag
	// key must be accompanied by a tag value, although null is accepted.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateCacheSubnetGroupOutput struct {

	// Represents the output of one of the following operations:
	//   - CreateCacheSubnetGroup
	//   - ModifyCacheSubnetGroup
	CacheSubnetGroup *types.CacheSubnetGroup

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateCacheSubnetGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpCreateCacheSubnetGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpCreateCacheSubnetGroup{}, middleware.After)
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
	if err = addOpCreateCacheSubnetGroupValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateCacheSubnetGroup(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateCacheSubnetGroup(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticache",
		OperationName: "CreateCacheSubnetGroup",
	}
}
