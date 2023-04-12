// Code generated by smithy-go-codegen DO NOT EDIT.

package marketplacecommerceanalytics

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/marketplacecommerceanalytics/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Given a data set type and data set publication date, asynchronously publishes
// the requested data set to the specified S3 bucket and notifies the specified SNS
// topic once the data is available. Returns a unique request identifier that can
// be used to correlate requests with notifications from the SNS topic. Data sets
// will be published in comma-separated values (CSV) format with the file name
// {data_set_type}_YYYY-MM-DD.csv. If a file with the same name already exists
// (e.g. if the same data set is requested twice), the original file will be
// overwritten by the new file. Requires a Role with an attached permissions policy
// providing Allow permissions for the following actions: s3:PutObject,
// s3:GetBucketLocation, sns:GetTopicAttributes, sns:Publish, iam:GetRolePolicy.
func (c *Client) GenerateDataSet(ctx context.Context, params *GenerateDataSetInput, optFns ...func(*Options)) (*GenerateDataSetOutput, error) {
	if params == nil {
		params = &GenerateDataSetInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GenerateDataSet", params, optFns, c.addOperationGenerateDataSetMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GenerateDataSetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Container for the parameters to the GenerateDataSet operation.
type GenerateDataSetInput struct {

	// The date a data set was published. For daily data sets, provide a date with
	// day-level granularity for the desired day. For monthly data sets except those
	// with prefix disbursed_amount, provide a date with month-level granularity for
	// the desired month (the day value will be ignored). For data sets with prefix
	// disbursed_amount, provide a date with day-level granularity for the desired day.
	// For these data sets we will look backwards in time over the range of 31 days
	// until the first data set is found (the latest one).
	//
	// This member is required.
	DataSetPublicationDate *time.Time

	// The desired data set type.
	//   - customer_subscriber_hourly_monthly_subscriptions From 2017-09-15 to
	//   present: Available daily by 24:00 UTC.
	//   - customer_subscriber_annual_subscriptions From 2017-09-15 to present:
	//   Available daily by 24:00 UTC.
	//   - daily_business_usage_by_instance_type From 2017-09-15 to present: Available
	//   daily by 24:00 UTC.
	//   - daily_business_fees From 2017-09-15 to present: Available daily by 24:00
	//   UTC.
	//   - daily_business_free_trial_conversions From 2017-09-15 to present: Available
	//   daily by 24:00 UTC.
	//   - daily_business_new_instances From 2017-09-15 to present: Available daily by
	//   24:00 UTC.
	//   - daily_business_new_product_subscribers From 2017-09-15 to present:
	//   Available daily by 24:00 UTC.
	//   - daily_business_canceled_product_subscribers From 2017-09-15 to present:
	//   Available daily by 24:00 UTC.
	//   - monthly_revenue_billing_and_revenue_data From 2017-09-15 to present:
	//   Available monthly on the 15th day of the month by 24:00 UTC. Data includes
	//   metered transactions (e.g. hourly) from one month prior.
	//   - monthly_revenue_annual_subscriptions From 2017-09-15 to present: Available
	//   monthly on the 15th day of the month by 24:00 UTC. Data includes up-front
	//   software charges (e.g. annual) from one month prior.
	//   - monthly_revenue_field_demonstration_usage From 2018-03-15 to present:
	//   Available monthly on the 15th day of the month by 24:00 UTC.
	//   - monthly_revenue_flexible_payment_schedule From 2018-11-15 to present:
	//   Available monthly on the 15th day of the month by 24:00 UTC.
	//   - disbursed_amount_by_product From 2017-09-15 to present: Available every 30
	//   days by 24:00 UTC.
	//   - disbursed_amount_by_instance_hours From 2017-09-15 to present: Available
	//   every 30 days by 24:00 UTC.
	//   - disbursed_amount_by_customer_geo From 2017-09-15 to present: Available
	//   every 30 days by 24:00 UTC.
	//   - disbursed_amount_by_age_of_uncollected_funds From 2017-09-15 to present:
	//   Available every 30 days by 24:00 UTC.
	//   - disbursed_amount_by_age_of_disbursed_funds From 2017-09-15 to present:
	//   Available every 30 days by 24:00 UTC.
	//   - disbursed_amount_by_age_of_past_due_funds From 2018-04-07 to present:
	//   Available every 30 days by 24:00 UTC.
	//   - disbursed_amount_by_uncollected_funds_breakdown From 2019-10-04 to present:
	//   Available every 30 days by 24:00 UTC.
	//   - sales_compensation_billed_revenue From 2017-09-15 to present: Available
	//   monthly on the 15th day of the month by 24:00 UTC. Data includes metered
	//   transactions (e.g. hourly) from one month prior, and up-front software charges
	//   (e.g. annual) from one month prior.
	//   - us_sales_and_use_tax_records From 2017-09-15 to present: Available monthly
	//   on the 15th day of the month by 24:00 UTC.
	//   - disbursed_amount_by_product_with_uncollected_funds This data set is
	//   deprecated. Download related reports from AMMP instead!
	//   - customer_profile_by_industry This data set is deprecated. Download related
	//   reports from AMMP instead!
	//   - customer_profile_by_revenue This data set is deprecated. Download related
	//   reports from AMMP instead!
	//   - customer_profile_by_geography This data set is deprecated. Download related
	//   reports from AMMP instead!
	//
	// This member is required.
	DataSetType types.DataSetType

	// The name (friendly name, not ARN) of the destination S3 bucket.
	//
	// This member is required.
	DestinationS3BucketName *string

	// The Amazon Resource Name (ARN) of the Role with an attached permissions policy
	// to interact with the provided AWS services.
	//
	// This member is required.
	RoleNameArn *string

	// Amazon Resource Name (ARN) for the SNS Topic that will be notified when the
	// data set has been published or if an error has occurred.
	//
	// This member is required.
	SnsTopicArn *string

	// (Optional) Key-value pairs which will be returned, unmodified, in the Amazon
	// SNS notification message and the data set metadata file. These key-value pairs
	// can be used to correlated responses with tracking information from other
	// systems.
	CustomerDefinedValues map[string]string

	// (Optional) The desired S3 prefix for the published data set, similar to a
	// directory path in standard file systems. For example, if given the bucket name
	// "mybucket" and the prefix "myprefix/mydatasets", the output file "outputfile"
	// would be published to "s3://mybucket/myprefix/mydatasets/outputfile". If the
	// prefix directory structure does not exist, it will be created. If no prefix is
	// provided, the data set will be published to the S3 bucket root.
	DestinationS3Prefix *string

	noSmithyDocumentSerde
}

// Container for the result of the GenerateDataSet operation.
type GenerateDataSetOutput struct {

	// A unique identifier representing a specific request to the GenerateDataSet
	// operation. This identifier can be used to correlate a request with notifications
	// from the SNS topic.
	DataSetRequestId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGenerateDataSetMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGenerateDataSet{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGenerateDataSet{}, middleware.After)
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
	if err = addOpGenerateDataSetValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGenerateDataSet(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGenerateDataSet(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "marketplacecommerceanalytics",
		OperationName: "GenerateDataSet",
	}
}
