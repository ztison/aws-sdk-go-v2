// Code generated by smithy-go-codegen DO NOT EDIT.

package sfn

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sfn/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Starts a Synchronous Express state machine execution. StartSyncExecution is not
// available for STANDARD workflows. StartSyncExecution will return a 200 OK
// response, even if your execution fails, because the status code in the API
// response doesn't reflect function errors. Error codes are reserved for errors
// that prevent your execution from running, such as permissions errors, limit
// errors, or issues with your state machine code and configuration. This API
// action isn't logged in CloudTrail.
func (c *Client) StartSyncExecution(ctx context.Context, params *StartSyncExecutionInput, optFns ...func(*Options)) (*StartSyncExecutionOutput, error) {
	if params == nil {
		params = &StartSyncExecutionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartSyncExecution", params, optFns, c.addOperationStartSyncExecutionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartSyncExecutionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartSyncExecutionInput struct {

	// The Amazon Resource Name (ARN) of the state machine to execute.
	//
	// This member is required.
	StateMachineArn *string

	// The string that contains the JSON input data for the execution, for example:
	// "input": "{\"first_name\" : \"test\"}" If you don't include any JSON input data,
	// you still must include the two braces, for example: "input": "{}" Length
	// constraints apply to the payload size, and are expressed as bytes in UTF-8
	// encoding.
	Input *string

	// The name of the execution.
	Name *string

	// Passes the X-Ray trace header. The trace header can also be passed in the
	// request payload.
	TraceHeader *string

	noSmithyDocumentSerde
}

type StartSyncExecutionOutput struct {

	// The Amazon Resource Name (ARN) that identifies the execution.
	//
	// This member is required.
	ExecutionArn *string

	// The date the execution is started.
	//
	// This member is required.
	StartDate *time.Time

	// The current status of the execution.
	//
	// This member is required.
	Status types.SyncExecutionStatus

	// If the execution has already ended, the date the execution stopped.
	//
	// This member is required.
	StopDate *time.Time

	// An object that describes workflow billing details, including billed duration
	// and memory use.
	BillingDetails *types.BillingDetails

	// A more detailed explanation of the cause of the failure.
	Cause *string

	// The error code of the failure.
	Error *string

	// The string that contains the JSON input data of the execution. Length
	// constraints apply to the payload size, and are expressed as bytes in UTF-8
	// encoding.
	Input *string

	// Provides details about execution input or output.
	InputDetails *types.CloudWatchEventsExecutionDataDetails

	// The name of the execution.
	Name *string

	// The JSON output data of the execution. Length constraints apply to the payload
	// size, and are expressed as bytes in UTF-8 encoding. This field is set only if
	// the execution succeeds. If the execution fails, this field is null.
	Output *string

	// Provides details about execution input or output.
	OutputDetails *types.CloudWatchEventsExecutionDataDetails

	// The Amazon Resource Name (ARN) that identifies the state machine.
	StateMachineArn *string

	// The X-Ray trace header that was passed to the execution.
	TraceHeader *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartSyncExecutionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpStartSyncExecution{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpStartSyncExecution{}, middleware.After)
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
	if err = addEndpointPrefix_opStartSyncExecutionMiddleware(stack); err != nil {
		return err
	}
	if err = addOpStartSyncExecutionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartSyncExecution(options.Region), middleware.Before); err != nil {
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

type endpointPrefix_opStartSyncExecutionMiddleware struct {
}

func (*endpointPrefix_opStartSyncExecutionMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opStartSyncExecutionMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "sync-" + req.URL.Host

	return next.HandleSerialize(ctx, in)
}
func addEndpointPrefix_opStartSyncExecutionMiddleware(stack *middleware.Stack) error {
	return stack.Serialize.Insert(&endpointPrefix_opStartSyncExecutionMiddleware{}, `OperationSerializer`, middleware.After)
}

func newServiceMetadataMiddleware_opStartSyncExecution(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "states",
		OperationName: "StartSyncExecution",
	}
}
