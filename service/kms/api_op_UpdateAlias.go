// Code generated by smithy-go-codegen DO NOT EDIT.

package kms

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Associates an existing KMS alias with a different KMS key. Each alias is
// associated with only one KMS key at a time, although a KMS key can have multiple
// aliases. The alias and the KMS key must be in the same Amazon Web Services
// account and Region. Adding, deleting, or updating an alias can allow or deny
// permission to the KMS key. For details, see ABAC for KMS (https://docs.aws.amazon.com/kms/latest/developerguide/abac.html)
// in the Key Management Service Developer Guide. The current and new KMS key must
// be the same type (both symmetric or both asymmetric or both HMAC), and they must
// have the same key usage. This restriction prevents errors in code that uses
// aliases. If you must assign an alias to a different type of KMS key, use
// DeleteAlias to delete the old alias and CreateAlias to create a new alias. You
// cannot use UpdateAlias to change an alias name. To change an alias name, use
// DeleteAlias to delete the old alias and CreateAlias to create a new alias.
// Because an alias is not a property of a KMS key, you can create, update, and
// delete the aliases of a KMS key without affecting the KMS key. Also, aliases do
// not appear in the response from the DescribeKey operation. To get the aliases
// of all KMS keys in the account, use the ListAliases operation. The KMS key that
// you use for this operation must be in a compatible key state. For details, see
// Key states of KMS keys (https://docs.aws.amazon.com/kms/latest/developerguide/key-state.html)
// in the Key Management Service Developer Guide. Cross-account use: No. You cannot
// perform this operation on a KMS key in a different Amazon Web Services account.
// Required permissions
//   - kms:UpdateAlias (https://docs.aws.amazon.com/kms/latest/developerguide/kms-api-permissions-reference.html)
//     on the alias (IAM policy).
//   - kms:UpdateAlias (https://docs.aws.amazon.com/kms/latest/developerguide/kms-api-permissions-reference.html)
//     on the current KMS key (key policy).
//   - kms:UpdateAlias (https://docs.aws.amazon.com/kms/latest/developerguide/kms-api-permissions-reference.html)
//     on the new KMS key (key policy).
//
// For details, see Controlling access to aliases (https://docs.aws.amazon.com/kms/latest/developerguide/kms-alias.html#alias-access)
// in the Key Management Service Developer Guide. Related operations:
//   - CreateAlias
//   - DeleteAlias
//   - ListAliases
func (c *Client) UpdateAlias(ctx context.Context, params *UpdateAliasInput, optFns ...func(*Options)) (*UpdateAliasOutput, error) {
	if params == nil {
		params = &UpdateAliasInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateAlias", params, optFns, c.addOperationUpdateAliasMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateAliasOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateAliasInput struct {

	// Identifies the alias that is changing its KMS key. This value must begin with
	// alias/ followed by the alias name, such as alias/ExampleAlias . You cannot use
	// UpdateAlias to change the alias name.
	//
	// This member is required.
	AliasName *string

	// Identifies the customer managed key (https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#customer-cmk)
	// to associate with the alias. You don't have permission to associate an alias
	// with an Amazon Web Services managed key (https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#aws-managed-cmk)
	// . The KMS key must be in the same Amazon Web Services account and Region as the
	// alias. Also, the new target KMS key must be the same type as the current target
	// KMS key (both symmetric or both asymmetric or both HMAC) and they must have the
	// same key usage. Specify the key ID or key ARN of the KMS key. For example:
	//   - Key ID: 1234abcd-12ab-34cd-56ef-1234567890ab
	//   - Key ARN:
	//   arn:aws:kms:us-east-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab
	// To get the key ID and key ARN for a KMS key, use ListKeys or DescribeKey . To
	// verify that the alias is mapped to the correct KMS key, use ListAliases .
	//
	// This member is required.
	TargetKeyId *string

	noSmithyDocumentSerde
}

type UpdateAliasOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateAliasMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateAlias{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateAlias{}, middleware.After)
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
	if err = addOpUpdateAliasValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateAlias(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateAlias(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kms",
		OperationName: "UpdateAlias",
	}
}
