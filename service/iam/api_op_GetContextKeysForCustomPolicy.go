// Code generated by smithy-go-codegen DO NOT EDIT.

package iam

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets a list of all of the context keys referenced in the input policies. The
// policies are supplied as a list of one or more strings. To get the context keys
// from policies associated with an IAM user, group, or role, use
// GetContextKeysForPrincipalPolicy . Context keys are variables maintained by
// Amazon Web Services and its services that provide details about the context of
// an API query request. Context keys can be evaluated by testing against a value
// specified in an IAM policy. Use GetContextKeysForCustomPolicy to understand
// what key names and values you must supply when you call SimulateCustomPolicy .
// Note that all parameters are shown in unencoded form here for clarity but must
// be URL encoded to be included as a part of a real HTML request.
func (c *Client) GetContextKeysForCustomPolicy(ctx context.Context, params *GetContextKeysForCustomPolicyInput, optFns ...func(*Options)) (*GetContextKeysForCustomPolicyOutput, error) {
	if params == nil {
		params = &GetContextKeysForCustomPolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetContextKeysForCustomPolicy", params, optFns, c.addOperationGetContextKeysForCustomPolicyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetContextKeysForCustomPolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetContextKeysForCustomPolicyInput struct {

	// A list of policies for which you want the list of context keys referenced in
	// those policies. Each document is specified as a string containing the complete,
	// valid JSON text of an IAM policy. The regex pattern (http://wikipedia.org/wiki/regex)
	// used to validate this parameter is a string of characters consisting of the
	// following:
	//   - Any printable ASCII character ranging from the space character ( \u0020 )
	//   through the end of the ASCII character range
	//   - The printable characters in the Basic Latin and Latin-1 Supplement
	//   character set (through \u00FF )
	//   - The special characters tab ( \u0009 ), line feed ( \u000A ), and carriage
	//   return ( \u000D )
	//
	// This member is required.
	PolicyInputList []string

	noSmithyDocumentSerde
}

// Contains the response to a successful GetContextKeysForPrincipalPolicy or
// GetContextKeysForCustomPolicy request.
type GetContextKeysForCustomPolicyOutput struct {

	// The list of context keys that are referenced in the input policies.
	ContextKeyNames []string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetContextKeysForCustomPolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpGetContextKeysForCustomPolicy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpGetContextKeysForCustomPolicy{}, middleware.After)
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
	if err = addOpGetContextKeysForCustomPolicyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetContextKeysForCustomPolicy(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetContextKeysForCustomPolicy(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iam",
		OperationName: "GetContextKeysForCustomPolicy",
	}
}
