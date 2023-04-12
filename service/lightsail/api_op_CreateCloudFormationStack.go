// Code generated by smithy-go-codegen DO NOT EDIT.

package lightsail

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lightsail/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates an AWS CloudFormation stack, which creates a new Amazon EC2 instance
// from an exported Amazon Lightsail snapshot. This operation results in a
// CloudFormation stack record that can be used to track the AWS CloudFormation
// stack created. Use the get cloud formation stack records operation to get a
// list of the CloudFormation stacks created. Wait until after your new Amazon EC2
// instance is created before running the create cloud formation stack operation
// again with the same export snapshot record.
func (c *Client) CreateCloudFormationStack(ctx context.Context, params *CreateCloudFormationStackInput, optFns ...func(*Options)) (*CreateCloudFormationStackOutput, error) {
	if params == nil {
		params = &CreateCloudFormationStackInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateCloudFormationStack", params, optFns, c.addOperationCreateCloudFormationStackMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateCloudFormationStackOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateCloudFormationStackInput struct {

	// An array of parameters that will be used to create the new Amazon EC2 instance.
	// You can only pass one instance entry at a time in this array. You will get an
	// invalid parameter error if you pass more than one instance entry in this array.
	//
	// This member is required.
	Instances []types.InstanceEntry

	noSmithyDocumentSerde
}

type CreateCloudFormationStackOutput struct {

	// An array of objects that describe the result of the action, such as the status
	// of the request, the timestamp of the request, and the resources affected by the
	// request.
	Operations []types.Operation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateCloudFormationStackMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateCloudFormationStack{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateCloudFormationStack{}, middleware.After)
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
	if err = addOpCreateCloudFormationStackValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateCloudFormationStack(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateCloudFormationStack(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lightsail",
		OperationName: "CreateCloudFormationStack",
	}
}
