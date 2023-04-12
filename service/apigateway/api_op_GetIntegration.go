// Code generated by smithy-go-codegen DO NOT EDIT.

package apigateway

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/apigateway/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Get the integration settings.
func (c *Client) GetIntegration(ctx context.Context, params *GetIntegrationInput, optFns ...func(*Options)) (*GetIntegrationOutput, error) {
	if params == nil {
		params = &GetIntegrationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetIntegration", params, optFns, c.addOperationGetIntegrationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetIntegrationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents a request to get the integration configuration.
type GetIntegrationInput struct {

	// Specifies a get integration request's HTTP method.
	//
	// This member is required.
	HttpMethod *string

	// Specifies a get integration request's resource identifier
	//
	// This member is required.
	ResourceId *string

	// The string identifier of the associated RestApi.
	//
	// This member is required.
	RestApiId *string

	noSmithyDocumentSerde
}

// Represents an HTTP, HTTP_PROXY, AWS, AWS_PROXY, or Mock integration.
type GetIntegrationOutput struct {

	// A list of request parameters whose values API Gateway caches. To be valid
	// values for cacheKeyParameters , these parameters must also be specified for
	// Method requestParameters .
	CacheKeyParameters []string

	// Specifies a group of related cached parameters. By default, API Gateway uses
	// the resource ID as the cacheNamespace . You can specify the same cacheNamespace
	// across resources to return the same cached data for requests to different
	// resources.
	CacheNamespace *string

	// The ID of the VpcLink used for the integration when connectionType=VPC_LINK and
	// undefined, otherwise.
	ConnectionId *string

	// The type of the network connection to the integration endpoint. The valid value
	// is INTERNET for connections through the public routable internet or VPC_LINK
	// for private connections between API Gateway and a network load balancer in a
	// VPC. The default value is INTERNET .
	ConnectionType types.ConnectionType

	// Specifies how to handle request payload content type conversions. Supported
	// values are CONVERT_TO_BINARY and CONVERT_TO_TEXT , with the following behaviors:
	// If this property is not defined, the request payload will be passed through from
	// the method request to integration request without modification, provided that
	// the passthroughBehavior is configured to support payload pass-through.
	ContentHandling types.ContentHandlingStrategy

	// Specifies the credentials required for the integration, if any. For AWS
	// integrations, three options are available. To specify an IAM Role for API
	// Gateway to assume, use the role's Amazon Resource Name (ARN). To require that
	// the caller's identity be passed through from the request, specify the string
	// arn:aws:iam::\*:user/\* . To use resource-based permissions on supported AWS
	// services, specify null.
	Credentials *string

	// Specifies the integration's HTTP method type.
	HttpMethod *string

	// Specifies the integration's responses.
	IntegrationResponses map[string]types.IntegrationResponse

	// Specifies how the method request body of an unmapped content type will be
	// passed through the integration request to the back end without transformation. A
	// content type is unmapped if no mapping template is defined in the integration or
	// the content type does not match any of the mapped content types, as specified in
	// requestTemplates . The valid value is one of the following: WHEN_NO_MATCH :
	// passes the method request body through the integration request to the back end
	// without transformation when the method request content type does not match any
	// content type associated with the mapping templates defined in the integration
	// request. WHEN_NO_TEMPLATES : passes the method request body through the
	// integration request to the back end without transformation when no mapping
	// template is defined in the integration request. If a template is defined when
	// this option is selected, the method request of an unmapped content-type will be
	// rejected with an HTTP 415 Unsupported Media Type response. NEVER : rejects the
	// method request with an HTTP 415 Unsupported Media Type response when either the
	// method request content type does not match any content type associated with the
	// mapping templates defined in the integration request or no mapping template is
	// defined in the integration request.
	PassthroughBehavior *string

	// A key-value map specifying request parameters that are passed from the method
	// request to the back end. The key is an integration request parameter name and
	// the associated value is a method request parameter value or static value that
	// must be enclosed within single quotes and pre-encoded as required by the back
	// end. The method request parameter value must match the pattern of
	// method.request.{location}.{name} , where location is querystring , path , or
	// header and name must be a valid and unique method request parameter name.
	RequestParameters map[string]string

	// Represents a map of Velocity templates that are applied on the request payload
	// based on the value of the Content-Type header sent by the client. The content
	// type value is the key in this map, and the template (as a String) is the value.
	RequestTemplates map[string]string

	// Custom timeout between 50 and 29,000 milliseconds. The default value is 29,000
	// milliseconds or 29 seconds.
	TimeoutInMillis int32

	// Specifies the TLS configuration for an integration.
	TlsConfig *types.TlsConfig

	// Specifies an API method integration type. The valid value is one of the
	// following: For the HTTP and HTTP proxy integrations, each integration can
	// specify a protocol ( http/https ), port and path. Standard 80 and 443 ports are
	// supported as well as custom ports above 1024. An HTTP or HTTP proxy integration
	// with a connectionType of VPC_LINK is referred to as a private integration and
	// uses a VpcLink to connect API Gateway to a network load balancer of a VPC.
	Type types.IntegrationType

	// Specifies Uniform Resource Identifier (URI) of the integration endpoint. For
	// HTTP or HTTP_PROXY integrations, the URI must be a fully formed, encoded
	// HTTP(S) URL according to the RFC-3986 specification, for either standard
	// integration, where connectionType is not VPC_LINK , or private integration,
	// where connectionType is VPC_LINK . For a private HTTP integration, the URI is
	// not used for routing. For AWS or AWS_PROXY integrations, the URI is of the form
	// arn:aws:apigateway:{region}:{subdomain.service|service}:path|action/{service_api}
	// . Here, {Region} is the API Gateway region (e.g., us-east-1); {service} is the
	// name of the integrated Amazon Web Services service (e.g., s3); and {subdomain}
	// is a designated subdomain supported by certain Amazon Web Services service for
	// fast host-name lookup. action can be used for an Amazon Web Services service
	// action-based API, using an Action={name}&{p1}={v1}&p2={v2}... query string. The
	// ensuing {service_api} refers to a supported action {name} plus any required
	// input parameters. Alternatively, path can be used for an AWS service path-based
	// API. The ensuing service_api refers to the path to an Amazon Web Services
	// service resource, including the region of the integrated Amazon Web Services
	// service, if applicable. For example, for integration with the S3 API of
	// GetObject, the uri can be either
	// arn:aws:apigateway:us-west-2:s3:action/GetObject&Bucket={bucket}&Key={key} or
	// arn:aws:apigateway:us-west-2:s3:path/{bucket}/{key}
	Uri *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetIntegrationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetIntegration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetIntegration{}, middleware.After)
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
	if err = addOpGetIntegrationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetIntegration(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addAcceptHeader(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opGetIntegration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "apigateway",
		OperationName: "GetIntegration",
	}
}
