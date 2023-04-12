// Code generated by smithy-go-codegen DO NOT EDIT.

package ecs

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ecs/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes one or more task definitions. You must deregister a task definition
// revision before you delete it. For more information, see
// DeregisterTaskDefinition (https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_DeregisterTaskDefinition.html)
// . When you delete a task definition revision, it is immediately transitions from
// the INACTIVE to DELETE_IN_PROGRESS . Existing tasks and services that reference
// a DELETE_IN_PROGRESS task definition revision continue to run without
// disruption. Existing services that reference a DELETE_IN_PROGRESS task
// definition revision can still scale up or down by modifying the service's
// desired count. You can't use a DELETE_IN_PROGRESS task definition revision to
// run new tasks or create new services. You also can't update an existing service
// to reference a DELETE_IN_PROGRESS task definition revision. A task definition
// revision will stay in DELETE_IN_PROGRESS status until all the associated tasks
// and services have been terminated.
func (c *Client) DeleteTaskDefinitions(ctx context.Context, params *DeleteTaskDefinitionsInput, optFns ...func(*Options)) (*DeleteTaskDefinitionsOutput, error) {
	if params == nil {
		params = &DeleteTaskDefinitionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteTaskDefinitions", params, optFns, c.addOperationDeleteTaskDefinitionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteTaskDefinitionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteTaskDefinitionsInput struct {

	// The family and revision ( family:revision ) or full Amazon Resource Name (ARN)
	// of the task definition to delete. You must specify a revision . You can specify
	// up to 10 task definitions as a comma separated list.
	//
	// This member is required.
	TaskDefinitions []string

	noSmithyDocumentSerde
}

type DeleteTaskDefinitionsOutput struct {

	// Any failures associated with the call.
	Failures []types.Failure

	// The list of deleted task definitions.
	TaskDefinitions []types.TaskDefinition

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteTaskDefinitionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteTaskDefinitions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteTaskDefinitions{}, middleware.After)
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
	if err = addOpDeleteTaskDefinitionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteTaskDefinitions(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteTaskDefinitions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ecs",
		OperationName: "DeleteTaskDefinitions",
	}
}
