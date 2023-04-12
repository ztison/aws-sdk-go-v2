// Code generated by smithy-go-codegen DO NOT EDIT.

package oam

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates or updates the resource policy that grants permissions to source
// accounts to link to the monitoring account sink. When you create a sink policy,
// you can grant permissions to all accounts in an organization or to individual
// accounts. You can also use a sink policy to limit the types of data that is
// shared. The three types that you can allow or deny are:
//   - Metrics - Specify with AWS::CloudWatch::Metric
//   - Log groups - Specify with AWS::Logs::LogGroup
//   - Traces - Specify with AWS::XRay::Trace
//
// See the examples in this section to see how to specify permitted source
// accounts and data types.
func (c *Client) PutSinkPolicy(ctx context.Context, params *PutSinkPolicyInput, optFns ...func(*Options)) (*PutSinkPolicyOutput, error) {
	if params == nil {
		params = &PutSinkPolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutSinkPolicy", params, optFns, c.addOperationPutSinkPolicyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutSinkPolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutSinkPolicyInput struct {

	// The JSON policy to use. If you are updating an existing policy, the entire
	// existing policy is replaced by what you specify here. The policy must be in JSON
	// string format with quotation marks escaped and no newlines. For examples of
	// different types of policies, see the Examples section on this page.
	//
	// This member is required.
	Policy *string

	// The ARN of the sink to attach this policy to.
	//
	// This member is required.
	SinkIdentifier *string

	noSmithyDocumentSerde
}

type PutSinkPolicyOutput struct {

	// The policy that you specified.
	Policy *string

	// The ARN of the sink.
	SinkArn *string

	// The random ID string that Amazon Web Services generated as part of the sink ARN.
	SinkId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutSinkPolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpPutSinkPolicy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpPutSinkPolicy{}, middleware.After)
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
	if err = addOpPutSinkPolicyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutSinkPolicy(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutSinkPolicy(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "oam",
		OperationName: "PutSinkPolicy",
	}
}
