// Code generated by smithy-go-codegen DO NOT EDIT.

package securitylake

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/securitylake/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Adds a third-party custom source in Amazon Security Lake, from the Amazon Web
// Services Region where you want to create a custom source. Security Lake can
// collect logs and events from third-party custom sources. After creating the
// appropriate IAM role to invoke Glue crawler, use this API to add a custom source
// name in Security Lake. This operation creates a partition in the Amazon S3
// bucket for Security Lake as the target location for log files from the custom
// source in addition to an associated Glue table and an Glue crawler.
func (c *Client) CreateCustomLogSource(ctx context.Context, params *CreateCustomLogSourceInput, optFns ...func(*Options)) (*CreateCustomLogSourceOutput, error) {
	if params == nil {
		params = &CreateCustomLogSourceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateCustomLogSource", params, optFns, c.addOperationCreateCustomLogSourceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateCustomLogSourceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateCustomLogSourceInput struct {

	// The name for a third-party custom source. This must be a Regionally unique
	// value.
	//
	// This member is required.
	CustomSourceName *string

	// The Open Cybersecurity Schema Framework (OCSF) event class which describes the
	// type of data that the custom source will send to Security Lake.
	//
	// This member is required.
	EventClass types.OcsfEventClass

	// The Amazon Resource Name (ARN) of the Identity and Access Management (IAM) role
	// to be used by the Glue crawler. The recommended IAM policies are:
	//   - The managed policy AWSGlueServiceRole
	//   - A custom policy granting access to your Amazon S3 Data Lake
	//
	// This member is required.
	GlueInvocationRoleArn *string

	// The Amazon Web Services account ID of the custom source that will write logs
	// and events into the Amazon S3 Data Lake.
	//
	// This member is required.
	LogProviderAccountId *string

	noSmithyDocumentSerde
}

type CreateCustomLogSourceOutput struct {

	// The location of the partition in the Amazon S3 bucket for Security Lake.
	//
	// This member is required.
	CustomDataLocation *string

	// The name of the Glue crawler.
	//
	// This member is required.
	GlueCrawlerName *string

	// The Glue database where results are written, such as:
	// arn:aws:daylight:us-east-1::database/sometable/* .
	//
	// This member is required.
	GlueDatabaseName *string

	// The table name of the Glue crawler.
	//
	// This member is required.
	GlueTableName *string

	// The ARN of the IAM role to be used by the entity putting logs into your custom
	// source partition. Security Lake will apply the correct access policies to this
	// role, but you must first manually create the trust policy for this role. The IAM
	// role name must start with the text 'Security Lake'. The IAM role must trust the
	// logProviderAccountId to assume the role.
	//
	// This member is required.
	LogProviderAccessRoleArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateCustomLogSourceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateCustomLogSource{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateCustomLogSource{}, middleware.After)
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
	if err = addOpCreateCustomLogSourceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateCustomLogSource(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateCustomLogSource(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "securitylake",
		OperationName: "CreateCustomLogSource",
	}
}
