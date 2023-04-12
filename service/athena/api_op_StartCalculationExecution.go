// Code generated by smithy-go-codegen DO NOT EDIT.

package athena

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/athena/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Submits calculations for execution within a session. You can supply the code to
// run as an inline code block within the request.
func (c *Client) StartCalculationExecution(ctx context.Context, params *StartCalculationExecutionInput, optFns ...func(*Options)) (*StartCalculationExecutionOutput, error) {
	if params == nil {
		params = &StartCalculationExecutionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartCalculationExecution", params, optFns, c.addOperationStartCalculationExecutionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartCalculationExecutionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartCalculationExecutionInput struct {

	// The session ID.
	//
	// This member is required.
	SessionId *string

	// Contains configuration information for the calculation.
	//
	// Deprecated: Kepler Post GA Tasks : https://sim.amazon.com/issues/ATHENA-39828
	CalculationConfiguration *types.CalculationConfiguration

	// A unique case-sensitive string used to ensure the request to create the
	// calculation is idempotent (executes only once). If another
	// StartCalculationExecutionRequest is received, the same response is returned and
	// another calculation is not created. If a parameter has changed, an error is
	// returned. This token is listed as not required because Amazon Web Services SDKs
	// (for example the Amazon Web Services SDK for Java) auto-generate the token for
	// users. If you are not using the Amazon Web Services SDK or the Amazon Web
	// Services CLI, you must provide this token or the action will fail.
	ClientRequestToken *string

	// A string that contains the code of the calculation.
	CodeBlock *string

	// A description of the calculation.
	Description *string

	noSmithyDocumentSerde
}

type StartCalculationExecutionOutput struct {

	// The calculation execution UUID.
	CalculationExecutionId *string

	// CREATING - The calculation is in the process of being created. CREATED - The
	// calculation has been created and is ready to run. QUEUED - The calculation has
	// been queued for processing. RUNNING - The calculation is running. CANCELING - A
	// request to cancel the calculation has been received and the system is working to
	// stop it. CANCELED - The calculation is no longer running as the result of a
	// cancel request. COMPLETED - The calculation has completed without error. FAILED
	// - The calculation failed and is no longer running.
	State types.CalculationExecutionState

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartCalculationExecutionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpStartCalculationExecution{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpStartCalculationExecution{}, middleware.After)
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
	if err = addOpStartCalculationExecutionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartCalculationExecution(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStartCalculationExecution(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "athena",
		OperationName: "StartCalculationExecution",
	}
}
