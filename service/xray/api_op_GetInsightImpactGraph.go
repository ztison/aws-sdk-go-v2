// Code generated by smithy-go-codegen DO NOT EDIT.

package xray

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/xray/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Retrieves a service graph structure filtered by the specified insight. The
// service graph is limited to only structural information. For a complete service
// graph, use this API with the GetServiceGraph API.
func (c *Client) GetInsightImpactGraph(ctx context.Context, params *GetInsightImpactGraphInput, optFns ...func(*Options)) (*GetInsightImpactGraphOutput, error) {
	if params == nil {
		params = &GetInsightImpactGraphInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetInsightImpactGraph", params, optFns, c.addOperationGetInsightImpactGraphMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetInsightImpactGraphOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetInsightImpactGraphInput struct {

	// The estimated end time of the insight, in Unix time seconds. The EndTime is
	// exclusive of the value provided. The time range between the start time and end
	// time can't be more than six hours.
	//
	// This member is required.
	EndTime *time.Time

	// The insight's unique identifier. Use the GetInsightSummaries action to retrieve
	// an InsightId.
	//
	// This member is required.
	InsightId *string

	// The estimated start time of the insight, in Unix time seconds. The StartTime is
	// inclusive of the value provided and can't be more than 30 days old.
	//
	// This member is required.
	StartTime *time.Time

	// Specify the pagination token returned by a previous request to retrieve the
	// next page of results.
	NextToken *string

	noSmithyDocumentSerde
}

type GetInsightImpactGraphOutput struct {

	// The provided end time.
	EndTime *time.Time

	// The insight's unique identifier.
	InsightId *string

	// Pagination token.
	NextToken *string

	// The time, in Unix seconds, at which the service graph ended.
	ServiceGraphEndTime *time.Time

	// The time, in Unix seconds, at which the service graph started.
	ServiceGraphStartTime *time.Time

	// The Amazon Web Services instrumented services related to the insight.
	Services []types.InsightImpactGraphService

	// The provided start time.
	StartTime *time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetInsightImpactGraphMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetInsightImpactGraph{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetInsightImpactGraph{}, middleware.After)
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
	if err = addOpGetInsightImpactGraphValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetInsightImpactGraph(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetInsightImpactGraph(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "xray",
		OperationName: "GetInsightImpactGraph",
	}
}
