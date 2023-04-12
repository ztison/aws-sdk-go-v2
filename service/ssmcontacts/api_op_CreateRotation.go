// Code generated by smithy-go-codegen DO NOT EDIT.

package ssmcontacts

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ssmcontacts/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Creates a rotation in an on-call schedule.
func (c *Client) CreateRotation(ctx context.Context, params *CreateRotationInput, optFns ...func(*Options)) (*CreateRotationOutput, error) {
	if params == nil {
		params = &CreateRotationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateRotation", params, optFns, c.addOperationCreateRotationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateRotationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateRotationInput struct {

	// The Amazon Resource Names (ARNs) of the contacts to add to the rotation. The
	// order that you list the contacts in is their shift order in the rotation
	// schedule. To change the order of the contact's shifts, use the UpdateRotation
	// operation.
	//
	// This member is required.
	ContactIds []string

	// The name of the rotation.
	//
	// This member is required.
	Name *string

	// Information about the rule that specifies when a shift's team members rotate.
	//
	// This member is required.
	Recurrence *types.RecurrenceSettings

	// The time zone to base the rotation’s activity on in Internet Assigned Numbers
	// Authority (IANA) format. For example: "America/Los_Angeles", "UTC", or
	// "Asia/Seoul". For more information, see the Time Zone Database (https://www.iana.org/time-zones)
	// on the IANA website. Designators for time zones that don’t support Daylight
	// Savings Time rules, such as Pacific Standard Time (PST) and Pacific Daylight
	// Time (PDT), are not supported.
	//
	// This member is required.
	TimeZoneId *string

	// A token that ensures that the operation is called only once with the specified
	// details.
	IdempotencyToken *string

	// The date and time that the rotation goes into effect.
	StartTime *time.Time

	// Optional metadata to assign to the rotation. Tags enable you to categorize a
	// resource in different ways, such as by purpose, owner, or environment. For more
	// information, see Tagging Incident Manager resources (https://docs.aws.amazon.com/incident-manager/latest/userguide/tagging.html)
	// in the Incident Manager User Guide.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateRotationOutput struct {

	// The Amazon Resource Name (ARN) of the created rotation.
	//
	// This member is required.
	RotationArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateRotationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateRotation{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateRotation{}, middleware.After)
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
	if err = addOpCreateRotationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateRotation(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateRotation(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ssm-contacts",
		OperationName: "CreateRotation",
	}
}
