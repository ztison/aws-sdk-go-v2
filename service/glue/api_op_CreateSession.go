// Code generated by smithy-go-codegen DO NOT EDIT.

package glue

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a new session.
func (c *Client) CreateSession(ctx context.Context, params *CreateSessionInput, optFns ...func(*Options)) (*CreateSessionOutput, error) {
	if params == nil {
		params = &CreateSessionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateSession", params, optFns, c.addOperationCreateSessionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateSessionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Request to create a new session.
type CreateSessionInput struct {

	// The SessionCommand that runs the job.
	//
	// This member is required.
	Command *types.SessionCommand

	// The ID of the session request.
	//
	// This member is required.
	Id *string

	// The IAM Role ARN
	//
	// This member is required.
	Role *string

	// The number of connections to use for the session.
	Connections *types.ConnectionsList

	// A map array of key-value pairs. Max is 75 pairs.
	DefaultArguments map[string]string

	// The description of the session.
	Description *string

	// The Glue version determines the versions of Apache Spark and Python that Glue
	// supports. The GlueVersion must be greater than 2.0.
	GlueVersion *string

	// The number of minutes when idle before session times out. Default for Spark ETL
	// jobs is value of Timeout. Consult the documentation for other job types.
	IdleTimeout *int32

	// The number of Glue data processing units (DPUs) that can be allocated when the
	// job runs. A DPU is a relative measure of processing power that consists of 4
	// vCPUs of compute capacity and 16 GB memory.
	MaxCapacity *float64

	// The number of workers of a defined WorkerType to use for the session.
	NumberOfWorkers *int32

	// The origin of the request.
	RequestOrigin *string

	// The name of the SecurityConfiguration structure to be used with the session
	SecurityConfiguration *string

	// The map of key value pairs (tags) belonging to the session.
	Tags map[string]string

	// The number of minutes before session times out. Default for Spark ETL jobs is
	// 48 hours (2880 minutes), the maximum session lifetime for this job type. Consult
	// the documentation for other job types.
	Timeout *int32

	// The type of predefined worker that is allocated to use for the session. Accepts
	// a value of Standard, G.1X, G.2X, or G.025X.
	//   - For the Standard worker type, each worker provides 4 vCPU, 16 GB of memory
	//   and a 50GB disk, and 2 executors per worker.
	//   - For the G.1X worker type, each worker maps to 1 DPU (4 vCPU, 16 GB of
	//   memory, 64 GB disk), and provides 1 executor per worker. We recommend this
	//   worker type for memory-intensive jobs.
	//   - For the G.2X worker type, each worker maps to 2 DPU (8 vCPU, 32 GB of
	//   memory, 128 GB disk), and provides 1 executor per worker. We recommend this
	//   worker type for memory-intensive jobs.
	//   - For the G.025X worker type, each worker maps to 0.25 DPU (2 vCPU, 4 GB of
	//   memory, 64 GB disk), and provides 1 executor per worker. We recommend this
	//   worker type for low volume streaming jobs. This worker type is only available
	//   for Glue version 3.0 streaming jobs.
	WorkerType types.WorkerType

	noSmithyDocumentSerde
}

type CreateSessionOutput struct {

	// Returns the session object in the response.
	Session *types.Session

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateSessionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateSession{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateSession{}, middleware.After)
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
	if err = addOpCreateSessionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateSession(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateSession(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "glue",
		OperationName: "CreateSession",
	}
}
