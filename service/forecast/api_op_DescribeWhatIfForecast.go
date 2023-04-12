// Code generated by smithy-go-codegen DO NOT EDIT.

package forecast

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/forecast/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Describes the what-if forecast created using the CreateWhatIfForecast
// operation. In addition to listing the properties provided in the
// CreateWhatIfForecast request, this operation lists the following properties:
//   - CreationTime
//   - LastModificationTime
//   - Message - If an error occurred, information about the error.
//   - Status
func (c *Client) DescribeWhatIfForecast(ctx context.Context, params *DescribeWhatIfForecastInput, optFns ...func(*Options)) (*DescribeWhatIfForecastOutput, error) {
	if params == nil {
		params = &DescribeWhatIfForecastInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeWhatIfForecast", params, optFns, c.addOperationDescribeWhatIfForecastMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeWhatIfForecastOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeWhatIfForecastInput struct {

	// The Amazon Resource Name (ARN) of the what-if forecast that you are interested
	// in.
	//
	// This member is required.
	WhatIfForecastArn *string

	noSmithyDocumentSerde
}

type DescribeWhatIfForecastOutput struct {

	// When the what-if forecast was created.
	CreationTime *time.Time

	// The approximate time remaining to complete the what-if forecast, in minutes.
	EstimatedTimeRemainingInMinutes *int64

	// The quantiles at which probabilistic forecasts are generated. You can specify
	// up to five quantiles per what-if forecast in the CreateWhatIfForecast
	// operation. If you didn't specify quantiles, the default values are ["0.1",
	// "0.5", "0.9"] .
	ForecastTypes []string

	// The last time the resource was modified. The timestamp depends on the status of
	// the job:
	//   - CREATE_PENDING - The CreationTime .
	//   - CREATE_IN_PROGRESS - The current timestamp.
	//   - CREATE_STOPPING - The current timestamp.
	//   - CREATE_STOPPED - When the job stopped.
	//   - ACTIVE or CREATE_FAILED - When the job finished or failed.
	LastModificationTime *time.Time

	// If an error occurred, an informational message about the error.
	Message *string

	// The status of the what-if forecast. States include:
	//   - ACTIVE
	//   - CREATE_PENDING , CREATE_IN_PROGRESS , CREATE_FAILED
	//   - CREATE_STOPPING , CREATE_STOPPED
	//   - DELETE_PENDING , DELETE_IN_PROGRESS , DELETE_FAILED
	// The Status of the what-if forecast must be ACTIVE before you can access the
	// forecast.
	Status *string

	// An array of S3Config , Schema , and Format elements that describe the
	// replacement time series.
	TimeSeriesReplacementsDataSource *types.TimeSeriesReplacementsDataSource

	// An array of Action and TimeSeriesConditions elements that describe what
	// transformations were applied to which time series.
	TimeSeriesTransformations []types.TimeSeriesTransformation

	// The Amazon Resource Name (ARN) of the what-if analysis that contains this
	// forecast.
	WhatIfAnalysisArn *string

	// The Amazon Resource Name (ARN) of the what-if forecast.
	WhatIfForecastArn *string

	// The name of the what-if forecast.
	WhatIfForecastName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeWhatIfForecastMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeWhatIfForecast{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeWhatIfForecast{}, middleware.After)
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
	if err = addOpDescribeWhatIfForecastValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeWhatIfForecast(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeWhatIfForecast(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "forecast",
		OperationName: "DescribeWhatIfForecast",
	}
}
