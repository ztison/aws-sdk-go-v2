// Code generated by smithy-go-codegen DO NOT EDIT.

package macie2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/macie2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Retrieves the criteria and other settings for a custom data identifier.
func (c *Client) GetCustomDataIdentifier(ctx context.Context, params *GetCustomDataIdentifierInput, optFns ...func(*Options)) (*GetCustomDataIdentifierOutput, error) {
	if params == nil {
		params = &GetCustomDataIdentifierInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetCustomDataIdentifier", params, optFns, c.addOperationGetCustomDataIdentifierMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetCustomDataIdentifierOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetCustomDataIdentifierInput struct {

	// The unique identifier for the Amazon Macie resource that the request applies to.
	//
	// This member is required.
	Id *string

	noSmithyDocumentSerde
}

type GetCustomDataIdentifierOutput struct {

	// The Amazon Resource Name (ARN) of the custom data identifier.
	Arn *string

	// The date and time, in UTC and extended ISO 8601 format, when the custom data
	// identifier was created.
	CreatedAt *time.Time

	// Specifies whether the custom data identifier was deleted. If you delete a
	// custom data identifier, Amazon Macie doesn't delete it permanently. Instead, it
	// soft deletes the identifier.
	Deleted bool

	// The custom description of the custom data identifier.
	Description *string

	// The unique identifier for the custom data identifier.
	Id *string

	// An array that lists specific character sequences (ignore words) to exclude from
	// the results. If the text matched by the regular expression contains any string
	// in this array, Amazon Macie ignores it. Ignore words are case sensitive.
	IgnoreWords []string

	// An array that lists specific character sequences (keywords), one of which must
	// precede and be within proximity (maximumMatchDistance) of the regular expression
	// to match. Keywords aren't case sensitive.
	Keywords []string

	// The maximum number of characters that can exist between the end of at least one
	// complete character sequence specified by the keywords array and the end of the
	// text that matches the regex pattern. If a complete keyword precedes all the text
	// that matches the pattern and the keyword is within the specified distance,
	// Amazon Macie includes the result. Otherwise, Macie excludes the result.
	MaximumMatchDistance int32

	// The custom name of the custom data identifier.
	Name *string

	// The regular expression (regex) that defines the pattern to match.
	Regex *string

	// Specifies the severity that's assigned to findings that the custom data
	// identifier produces, based on the number of occurrences of text that matches the
	// custom data identifier's detection criteria. By default, Amazon Macie creates
	// findings for S3 objects that contain at least one occurrence of text that
	// matches the detection criteria, and Macie assigns the MEDIUM severity to those
	// findings.
	SeverityLevels []types.SeverityLevel

	// A map of key-value pairs that identifies the tags (keys and values) that are
	// associated with the custom data identifier.
	Tags map[string]string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetCustomDataIdentifierMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetCustomDataIdentifier{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetCustomDataIdentifier{}, middleware.After)
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
	if err = addOpGetCustomDataIdentifierValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetCustomDataIdentifier(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetCustomDataIdentifier(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "macie2",
		OperationName: "GetCustomDataIdentifier",
	}
}
