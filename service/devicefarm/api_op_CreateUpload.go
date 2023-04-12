// Code generated by smithy-go-codegen DO NOT EDIT.

package devicefarm

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/devicefarm/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Uploads an app or test scripts.
func (c *Client) CreateUpload(ctx context.Context, params *CreateUploadInput, optFns ...func(*Options)) (*CreateUploadOutput, error) {
	if params == nil {
		params = &CreateUploadInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateUpload", params, optFns, c.addOperationCreateUploadMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateUploadOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents a request to the create upload operation.
type CreateUploadInput struct {

	// The upload's file name. The name should not contain any forward slashes ( / ).
	// If you are uploading an iOS app, the file name must end with the .ipa
	// extension. If you are uploading an Android app, the file name must end with the
	// .apk extension. For all others, the file name must end with the .zip file
	// extension.
	//
	// This member is required.
	Name *string

	// The ARN of the project for the upload.
	//
	// This member is required.
	ProjectArn *string

	// The upload's upload type. Must be one of the following values:
	//   - ANDROID_APP
	//   - IOS_APP
	//   - WEB_APP
	//   - EXTERNAL_DATA
	//   - APPIUM_JAVA_JUNIT_TEST_PACKAGE
	//   - APPIUM_JAVA_TESTNG_TEST_PACKAGE
	//   - APPIUM_PYTHON_TEST_PACKAGE
	//   - APPIUM_NODE_TEST_PACKAGE
	//   - APPIUM_RUBY_TEST_PACKAGE
	//   - APPIUM_WEB_JAVA_JUNIT_TEST_PACKAGE
	//   - APPIUM_WEB_JAVA_TESTNG_TEST_PACKAGE
	//   - APPIUM_WEB_PYTHON_TEST_PACKAGE
	//   - APPIUM_WEB_NODE_TEST_PACKAGE
	//   - APPIUM_WEB_RUBY_TEST_PACKAGE
	//   - CALABASH_TEST_PACKAGE
	//   - INSTRUMENTATION_TEST_PACKAGE
	//   - UIAUTOMATION_TEST_PACKAGE
	//   - UIAUTOMATOR_TEST_PACKAGE
	//   - XCTEST_TEST_PACKAGE
	//   - XCTEST_UI_TEST_PACKAGE
	//   - APPIUM_JAVA_JUNIT_TEST_SPEC
	//   - APPIUM_JAVA_TESTNG_TEST_SPEC
	//   - APPIUM_PYTHON_TEST_SPEC
	//   - APPIUM_NODE_TEST_SPEC
	//   - APPIUM_RUBY_TEST_SPEC
	//   - APPIUM_WEB_JAVA_JUNIT_TEST_SPEC
	//   - APPIUM_WEB_JAVA_TESTNG_TEST_SPEC
	//   - APPIUM_WEB_PYTHON_TEST_SPEC
	//   - APPIUM_WEB_NODE_TEST_SPEC
	//   - APPIUM_WEB_RUBY_TEST_SPEC
	//   - INSTRUMENTATION_TEST_SPEC
	//   - XCTEST_UI_TEST_SPEC
	// If you call CreateUpload with WEB_APP specified, AWS Device Farm throws an
	// ArgumentException error.
	//
	// This member is required.
	Type types.UploadType

	// The upload's content type (for example, application/octet-stream ).
	ContentType *string

	noSmithyDocumentSerde
}

// Represents the result of a create upload request.
type CreateUploadOutput struct {

	// The newly created upload.
	Upload *types.Upload

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateUploadMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateUpload{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateUpload{}, middleware.After)
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
	if err = addOpCreateUploadValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateUpload(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateUpload(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "devicefarm",
		OperationName: "CreateUpload",
	}
}
