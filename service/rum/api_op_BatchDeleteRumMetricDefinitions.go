// Code generated by smithy-go-codegen DO NOT EDIT.

package rum

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/rum/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Removes the specified metrics from being sent to an extended metrics
// destination. If some metric definition IDs specified in a
// BatchDeleteRumMetricDefinitions operations are not valid, those metric
// definitions fail and return errors, but all valid metric definition IDs in the
// same operation are still deleted. The maximum number of metric definitions that
// you can specify in one BatchDeleteRumMetricDefinitions operation is 200.
func (c *Client) BatchDeleteRumMetricDefinitions(ctx context.Context, params *BatchDeleteRumMetricDefinitionsInput, optFns ...func(*Options)) (*BatchDeleteRumMetricDefinitionsOutput, error) {
	if params == nil {
		params = &BatchDeleteRumMetricDefinitionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "BatchDeleteRumMetricDefinitions", params, optFns, c.addOperationBatchDeleteRumMetricDefinitionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*BatchDeleteRumMetricDefinitionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type BatchDeleteRumMetricDefinitionsInput struct {

	// The name of the CloudWatch RUM app monitor that is sending these metrics.
	//
	// This member is required.
	AppMonitorName *string

	// Defines the destination where you want to stop sending the specified metrics.
	// Valid values are CloudWatch and Evidently . If you specify Evidently , you must
	// also specify the ARN of the CloudWatchEvidently experiment that is to be the
	// destination and an IAM role that has permission to write to the experiment.
	//
	// This member is required.
	Destination types.MetricDestination

	// An array of structures which define the metrics that you want to stop sending.
	//
	// This member is required.
	MetricDefinitionIds []string

	// This parameter is required if Destination is Evidently . If Destination is
	// CloudWatch , do not use this parameter. This parameter specifies the ARN of the
	// Evidently experiment that was receiving the metrics that are being deleted.
	DestinationArn *string

	noSmithyDocumentSerde
}

type BatchDeleteRumMetricDefinitionsOutput struct {

	// An array of error objects, if the operation caused any errors.
	//
	// This member is required.
	Errors []types.BatchDeleteRumMetricDefinitionsError

	// The IDs of the metric definitions that were deleted.
	MetricDefinitionIds []string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationBatchDeleteRumMetricDefinitionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpBatchDeleteRumMetricDefinitions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpBatchDeleteRumMetricDefinitions{}, middleware.After)
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
	if err = addOpBatchDeleteRumMetricDefinitionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opBatchDeleteRumMetricDefinitions(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opBatchDeleteRumMetricDefinitions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rum",
		OperationName: "BatchDeleteRumMetricDefinitions",
	}
}
