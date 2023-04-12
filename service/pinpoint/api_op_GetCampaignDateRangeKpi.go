// Code generated by smithy-go-codegen DO NOT EDIT.

package pinpoint

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/pinpoint/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Retrieves (queries) pre-aggregated data for a standard metric that applies to a
// campaign.
func (c *Client) GetCampaignDateRangeKpi(ctx context.Context, params *GetCampaignDateRangeKpiInput, optFns ...func(*Options)) (*GetCampaignDateRangeKpiOutput, error) {
	if params == nil {
		params = &GetCampaignDateRangeKpiInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetCampaignDateRangeKpi", params, optFns, c.addOperationGetCampaignDateRangeKpiMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetCampaignDateRangeKpiOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetCampaignDateRangeKpiInput struct {

	// The unique identifier for the application. This identifier is displayed as the
	// Project ID on the Amazon Pinpoint console.
	//
	// This member is required.
	ApplicationId *string

	// The unique identifier for the campaign.
	//
	// This member is required.
	CampaignId *string

	// The name of the metric, also referred to as a key performance indicator (KPI),
	// to retrieve data for. This value describes the associated metric and consists of
	// two or more terms, which are comprised of lowercase alphanumeric characters,
	// separated by a hyphen. Examples are email-open-rate and
	// successful-delivery-rate. For a list of valid values, see the Amazon Pinpoint
	// Developer Guide (https://docs.aws.amazon.com/pinpoint/latest/developerguide/analytics-standard-metrics.html)
	// .
	//
	// This member is required.
	KpiName *string

	// The last date and time to retrieve data for, as part of an inclusive date range
	// that filters the query results. This value should be in extended ISO 8601 format
	// and use Coordinated Universal Time (UTC), for example: 2019-07-26T20:00:00Z for
	// 8:00 PM UTC July 26, 2019.
	EndTime *time.Time

	// The string that specifies which page of results to return in a paginated
	// response. This parameter is not supported for application, campaign, and journey
	// metrics.
	NextToken *string

	// The maximum number of items to include in each page of a paginated response.
	// This parameter is not supported for application, campaign, and journey metrics.
	PageSize *string

	// The first date and time to retrieve data for, as part of an inclusive date
	// range that filters the query results. This value should be in extended ISO 8601
	// format and use Coordinated Universal Time (UTC), for example:
	// 2019-07-19T20:00:00Z for 8:00 PM UTC July 19, 2019. This value should also be
	// fewer than 90 days from the current day.
	StartTime *time.Time

	noSmithyDocumentSerde
}

type GetCampaignDateRangeKpiOutput struct {

	// Provides the results of a query that retrieved the data for a standard metric
	// that applies to a campaign, and provides information about that query.
	//
	// This member is required.
	CampaignDateRangeKpiResponse *types.CampaignDateRangeKpiResponse

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetCampaignDateRangeKpiMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetCampaignDateRangeKpi{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetCampaignDateRangeKpi{}, middleware.After)
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
	if err = addOpGetCampaignDateRangeKpiValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetCampaignDateRangeKpi(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetCampaignDateRangeKpi(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mobiletargeting",
		OperationName: "GetCampaignDateRangeKpi",
	}
}
