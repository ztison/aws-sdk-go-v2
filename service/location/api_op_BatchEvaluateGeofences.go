// Code generated by smithy-go-codegen DO NOT EDIT.

package location

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/location/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Evaluates device positions against the geofence geometries from a given
// geofence collection. This operation always returns an empty response because
// geofences are asynchronously evaluated. The evaluation determines if the device
// has entered or exited a geofenced area, and then publishes one of the following
// events to Amazon EventBridge:
//   - ENTER if Amazon Location determines that the tracked device has entered a
//     geofenced area.
//   - EXIT if Amazon Location determines that the tracked device has exited a
//     geofenced area.
//
// The last geofence that a device was observed within is tracked for 30 days
// after the most recent device position update. Geofence evaluation uses the given
// device position. It does not account for the optional Accuracy of a
// DevicePositionUpdate . The DeviceID is used as a string to represent the
// device. You do not need to have a Tracker associated with the DeviceID .
func (c *Client) BatchEvaluateGeofences(ctx context.Context, params *BatchEvaluateGeofencesInput, optFns ...func(*Options)) (*BatchEvaluateGeofencesOutput, error) {
	if params == nil {
		params = &BatchEvaluateGeofencesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "BatchEvaluateGeofences", params, optFns, c.addOperationBatchEvaluateGeofencesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*BatchEvaluateGeofencesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type BatchEvaluateGeofencesInput struct {

	// The geofence collection used in evaluating the position of devices against its
	// geofences.
	//
	// This member is required.
	CollectionName *string

	// Contains device details for each device to be evaluated against the given
	// geofence collection.
	//
	// This member is required.
	DevicePositionUpdates []types.DevicePositionUpdate

	noSmithyDocumentSerde
}

type BatchEvaluateGeofencesOutput struct {

	// Contains error details for each device that failed to evaluate its position
	// against the given geofence collection.
	//
	// This member is required.
	Errors []types.BatchEvaluateGeofencesError

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationBatchEvaluateGeofencesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpBatchEvaluateGeofences{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpBatchEvaluateGeofences{}, middleware.After)
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
	if err = addEndpointPrefix_opBatchEvaluateGeofencesMiddleware(stack); err != nil {
		return err
	}
	if err = addOpBatchEvaluateGeofencesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opBatchEvaluateGeofences(options.Region), middleware.Before); err != nil {
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

type endpointPrefix_opBatchEvaluateGeofencesMiddleware struct {
}

func (*endpointPrefix_opBatchEvaluateGeofencesMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opBatchEvaluateGeofencesMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "geofencing." + req.URL.Host

	return next.HandleSerialize(ctx, in)
}
func addEndpointPrefix_opBatchEvaluateGeofencesMiddleware(stack *middleware.Stack) error {
	return stack.Serialize.Insert(&endpointPrefix_opBatchEvaluateGeofencesMiddleware{}, `OperationSerializer`, middleware.After)
}

func newServiceMetadataMiddleware_opBatchEvaluateGeofences(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "geo",
		OperationName: "BatchEvaluateGeofences",
	}
}
