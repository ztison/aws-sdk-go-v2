// Code generated by smithy-go-codegen DO NOT EDIT.

package wafregional

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// This is AWS WAF Classic documentation. For more information, see AWS WAF Classic (https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html)
// in the developer guide. For the latest version of AWS WAF, use the AWS WAFV2 API
// and see the AWS WAF Developer Guide (https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html)
// . With the latest version, AWS WAF has a single set of endpoints for regional
// and global use. Returns an array of IP addresses currently being blocked by the
// RateBasedRule that is specified by the RuleId . The maximum number of managed
// keys that will be blocked is 10,000. If more than 10,000 addresses exceed the
// rate limit, the 10,000 addresses with the highest rates will be blocked.
func (c *Client) GetRateBasedRuleManagedKeys(ctx context.Context, params *GetRateBasedRuleManagedKeysInput, optFns ...func(*Options)) (*GetRateBasedRuleManagedKeysOutput, error) {
	if params == nil {
		params = &GetRateBasedRuleManagedKeysInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetRateBasedRuleManagedKeys", params, optFns, c.addOperationGetRateBasedRuleManagedKeysMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetRateBasedRuleManagedKeysOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetRateBasedRuleManagedKeysInput struct {

	// The RuleId of the RateBasedRule for which you want to get a list of ManagedKeys
	// . RuleId is returned by CreateRateBasedRule and by ListRateBasedRules .
	//
	// This member is required.
	RuleId *string

	// A null value and not currently used. Do not include this in your request.
	NextMarker *string

	noSmithyDocumentSerde
}

type GetRateBasedRuleManagedKeysOutput struct {

	// An array of IP addresses that currently are blocked by the specified
	// RateBasedRule .
	ManagedKeys []string

	// A null value and not currently used.
	NextMarker *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetRateBasedRuleManagedKeysMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetRateBasedRuleManagedKeys{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetRateBasedRuleManagedKeys{}, middleware.After)
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
	if err = addOpGetRateBasedRuleManagedKeysValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetRateBasedRuleManagedKeys(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetRateBasedRuleManagedKeys(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "waf-regional",
		OperationName: "GetRateBasedRuleManagedKeys",
	}
}
