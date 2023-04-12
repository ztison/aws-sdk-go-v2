// Code generated by smithy-go-codegen DO NOT EDIT.

package cognitosync

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cognitosync/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Posts updates to records and adds and deletes records for a dataset and user.
// The sync count in the record patch is your last known sync count for that
// record. The server will reject an UpdateRecords request with a
// ResourceConflictException if you try to patch a record with a new value but a
// stale sync count.For example, if the sync count on the server is 5 for a key
// called highScore and you try and submit a new highScore with sync count of 4,
// the request will be rejected. To obtain the current sync count for a record,
// call ListRecords. On a successful update of the record, the response returns the
// new sync count for that record. You should present that sync count the next time
// you try to update that same record. When the record does not exist, specify the
// sync count as 0. This API can be called with temporary user credentials provided
// by Cognito Identity or with developer credentials.
func (c *Client) UpdateRecords(ctx context.Context, params *UpdateRecordsInput, optFns ...func(*Options)) (*UpdateRecordsOutput, error) {
	if params == nil {
		params = &UpdateRecordsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateRecords", params, optFns, c.addOperationUpdateRecordsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateRecordsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A request to post updates to records or add and delete records for a dataset
// and user.
type UpdateRecordsInput struct {

	// A string of up to 128 characters. Allowed characters are a-z, A-Z, 0-9, '_'
	// (underscore), '-' (dash), and '.' (dot).
	//
	// This member is required.
	DatasetName *string

	// A name-spaced GUID (for example,
	// us-east-1:23EC4050-6AEA-7089-A2DD-08002EXAMPLE) created by Amazon Cognito. GUID
	// generation is unique within a region.
	//
	// This member is required.
	IdentityId *string

	// A name-spaced GUID (for example,
	// us-east-1:23EC4050-6AEA-7089-A2DD-08002EXAMPLE) created by Amazon Cognito. GUID
	// generation is unique within a region.
	//
	// This member is required.
	IdentityPoolId *string

	// The SyncSessionToken returned by a previous call to ListRecords for this
	// dataset and identity.
	//
	// This member is required.
	SyncSessionToken *string

	// Intended to supply a device ID that will populate the lastModifiedBy field
	// referenced in other methods. The ClientContext field is not yet implemented.
	ClientContext *string

	// The unique ID generated for this device by Cognito.
	DeviceId *string

	// A list of patch operations.
	RecordPatches []types.RecordPatch

	noSmithyDocumentSerde
}

// Returned for a successful UpdateRecordsRequest.
type UpdateRecordsOutput struct {

	// A list of records that have been updated.
	Records []types.Record

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateRecordsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateRecords{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateRecords{}, middleware.After)
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
	if err = addOpUpdateRecordsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateRecords(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateRecords(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cognito-sync",
		OperationName: "UpdateRecords",
	}
}
