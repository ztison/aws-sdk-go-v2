// Code generated by smithy-go-codegen DO NOT EDIT.

package emr

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/emr/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Fetches mapping details for the specified Amazon EMR Studio and identity (user
// or group).
func (c *Client) GetStudioSessionMapping(ctx context.Context, params *GetStudioSessionMappingInput, optFns ...func(*Options)) (*GetStudioSessionMappingOutput, error) {
	if params == nil {
		params = &GetStudioSessionMappingInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetStudioSessionMapping", params, optFns, c.addOperationGetStudioSessionMappingMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetStudioSessionMappingOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetStudioSessionMappingInput struct {

	// Specifies whether the identity to fetch is a user or a group.
	//
	// This member is required.
	IdentityType types.IdentityType

	// The ID of the Amazon EMR Studio.
	//
	// This member is required.
	StudioId *string

	// The globally unique identifier (GUID) of the user or group. For more
	// information, see UserId (https://docs.aws.amazon.com/singlesignon/latest/IdentityStoreAPIReference/API_User.html#singlesignon-Type-User-UserId)
	// and GroupId (https://docs.aws.amazon.com/singlesignon/latest/IdentityStoreAPIReference/API_Group.html#singlesignon-Type-Group-GroupId)
	// in the IAM Identity Center Identity Store API Reference. Either IdentityName or
	// IdentityId must be specified.
	IdentityId *string

	// The name of the user or group to fetch. For more information, see UserName (https://docs.aws.amazon.com/singlesignon/latest/IdentityStoreAPIReference/API_User.html#singlesignon-Type-User-UserName)
	// and DisplayName (https://docs.aws.amazon.com/singlesignon/latest/IdentityStoreAPIReference/API_Group.html#singlesignon-Type-Group-DisplayName)
	// in the IAM Identity Center Identity Store API Reference. Either IdentityName or
	// IdentityId must be specified.
	IdentityName *string

	noSmithyDocumentSerde
}

type GetStudioSessionMappingOutput struct {

	// The session mapping details for the specified Amazon EMR Studio and identity,
	// including session policy ARN and creation time.
	SessionMapping *types.SessionMappingDetail

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetStudioSessionMappingMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetStudioSessionMapping{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetStudioSessionMapping{}, middleware.After)
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
	if err = addOpGetStudioSessionMappingValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetStudioSessionMapping(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetStudioSessionMapping(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticmapreduce",
		OperationName: "GetStudioSessionMapping",
	}
}
