// Code generated by smithy-go-codegen DO NOT EDIT.

package ssmincidents

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes a timeline event from an incident.
func (c *Client) DeleteTimelineEvent(ctx context.Context, params *DeleteTimelineEventInput, optFns ...func(*Options)) (*DeleteTimelineEventOutput, error) {
	if params == nil {
		params = &DeleteTimelineEventInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteTimelineEvent", params, optFns, c.addOperationDeleteTimelineEventMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteTimelineEventOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteTimelineEventInput struct {

	// The ID of the event to update. You can use ListTimelineEvents to find an
	// event's ID.
	//
	// This member is required.
	EventId *string

	// The Amazon Resource Name (ARN) of the incident that includes the timeline event.
	//
	// This member is required.
	IncidentRecordArn *string

	noSmithyDocumentSerde
}

type DeleteTimelineEventOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteTimelineEventMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDeleteTimelineEvent{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteTimelineEvent{}, middleware.After)
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
	if err = addOpDeleteTimelineEventValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteTimelineEvent(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteTimelineEvent(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ssm-incidents",
		OperationName: "DeleteTimelineEvent",
	}
}
