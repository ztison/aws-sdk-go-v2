// Code generated by smithy-go-codegen DO NOT EDIT.

package kendra

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/kendra/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Gets information about a block list used for query suggestions for an index.
// This is used to check the current settings that are applied to a block list.
// DescribeQuerySuggestionsBlockList is currently not supported in the Amazon Web
// Services GovCloud (US-West) region.
func (c *Client) DescribeQuerySuggestionsBlockList(ctx context.Context, params *DescribeQuerySuggestionsBlockListInput, optFns ...func(*Options)) (*DescribeQuerySuggestionsBlockListOutput, error) {
	if params == nil {
		params = &DescribeQuerySuggestionsBlockListInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeQuerySuggestionsBlockList", params, optFns, c.addOperationDescribeQuerySuggestionsBlockListMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeQuerySuggestionsBlockListOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeQuerySuggestionsBlockListInput struct {

	// The identifier of the block list you want to get information on.
	//
	// This member is required.
	Id *string

	// The identifier of the index for the block list.
	//
	// This member is required.
	IndexId *string

	noSmithyDocumentSerde
}

type DescribeQuerySuggestionsBlockListOutput struct {

	// The Unix timestamp when a block list for query suggestions was created.
	CreatedAt *time.Time

	// The description for the block list.
	Description *string

	// The error message containing details if there are issues processing the block
	// list.
	ErrorMessage *string

	// The current size of the block list text file in S3.
	FileSizeBytes *int64

	// The identifier of the block list.
	Id *string

	// The identifier of the index for the block list.
	IndexId *string

	// The current number of valid, non-empty words or phrases in the block list text
	// file.
	ItemCount *int32

	// The name of the block list.
	Name *string

	// The IAM (Identity and Access Management) role used by Amazon Kendra to access
	// the block list text file in S3. The role needs S3 read permissions to your file
	// in S3 and needs to give STS (Security Token Service) assume role permissions to
	// Amazon Kendra.
	RoleArn *string

	// Shows the current S3 path to your block list text file in your S3 bucket. Each
	// block word or phrase should be on a separate line in a text file. For
	// information on the current quota limits for block lists, see Quotas for Amazon
	// Kendra (https://docs.aws.amazon.com/kendra/latest/dg/quotas.html) .
	SourceS3Path *types.S3Path

	// The current status of the block list. When the value is ACTIVE , the block list
	// is ready for use.
	Status types.QuerySuggestionsBlockListStatus

	// The Unix timestamp when a block list for query suggestions was last updated.
	UpdatedAt *time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeQuerySuggestionsBlockListMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeQuerySuggestionsBlockList{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeQuerySuggestionsBlockList{}, middleware.After)
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
	if err = addOpDescribeQuerySuggestionsBlockListValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeQuerySuggestionsBlockList(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeQuerySuggestionsBlockList(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kendra",
		OperationName: "DescribeQuerySuggestionsBlockList",
	}
}
