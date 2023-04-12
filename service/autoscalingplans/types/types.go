// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	smithydocument "github.com/aws/smithy-go/document"
	"time"
)

// Represents an application source.
type ApplicationSource struct {

	// The Amazon Resource Name (ARN) of a AWS CloudFormation stack.
	CloudFormationStackARN *string

	// A set of tags (up to 50).
	TagFilters []TagFilter

	noSmithyDocumentSerde
}

// Represents a CloudWatch metric of your choosing that can be used for predictive
// scaling. For predictive scaling to work with a customized load metric
// specification, AWS Auto Scaling needs access to the Sum and Average statistics
// that CloudWatch computes from metric data. When you choose a load metric, make
// sure that the required Sum and Average statistics for your metric are available
// in CloudWatch and that they provide relevant data for predictive scaling. The
// Sum statistic must represent the total load on the resource, and the Average
// statistic must represent the average load per capacity unit of the resource. For
// example, there is a metric that counts the number of requests processed by your
// Auto Scaling group. If the Sum statistic represents the total request count
// processed by the group, then the Average statistic for the specified metric
// must represent the average request count processed by each instance of the
// group. If you publish your own metrics, you can aggregate the data points at a
// given interval and then publish the aggregated data points to CloudWatch. Before
// AWS Auto Scaling generates the forecast, it sums up all the metric data points
// that occurred within each hour to match the granularity period that is used in
// the forecast (60 minutes). For information about terminology, available metrics,
// or how to publish new metrics, see Amazon CloudWatch Concepts (https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/cloudwatch_concepts.html)
// in the Amazon CloudWatch User Guide. After creating your scaling plan, you can
// use the AWS Auto Scaling console to visualize forecasts for the specified
// metric. For more information, see View Scaling Information for a Resource (https://docs.aws.amazon.com/autoscaling/plans/userguide/gs-create-scaling-plan.html#gs-view-resource)
// in the AWS Auto Scaling User Guide.
type CustomizedLoadMetricSpecification struct {

	// The name of the metric.
	//
	// This member is required.
	MetricName *string

	// The namespace of the metric.
	//
	// This member is required.
	Namespace *string

	// The statistic of the metric. The only valid value is Sum .
	//
	// This member is required.
	Statistic MetricStatistic

	// The dimensions of the metric. Conditional: If you published your metric with
	// dimensions, you must specify the same dimensions in your customized load metric
	// specification.
	Dimensions []MetricDimension

	// The unit of the metric.
	Unit *string

	noSmithyDocumentSerde
}

// Represents a CloudWatch metric of your choosing that can be used for dynamic
// scaling as part of a target tracking scaling policy. To create your customized
// scaling metric specification:
//   - Add values for each required parameter from CloudWatch. You can use an
//     existing metric, or a new metric that you create. To use your own metric, you
//     must first publish the metric to CloudWatch. For more information, see
//     Publish Custom Metrics (https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/publishingMetrics.html)
//     in the Amazon CloudWatch User Guide.
//   - Choose a metric that changes proportionally with capacity. The value of the
//     metric should increase or decrease in inverse proportion to the number of
//     capacity units. That is, the value of the metric should decrease when capacity
//     increases.
//
// For information about terminology, available metrics, or how to publish new
// metrics, see Amazon CloudWatch Concepts (https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/cloudwatch_concepts.html)
// in the Amazon CloudWatch User Guide.
type CustomizedScalingMetricSpecification struct {

	// The name of the metric.
	//
	// This member is required.
	MetricName *string

	// The namespace of the metric.
	//
	// This member is required.
	Namespace *string

	// The statistic of the metric.
	//
	// This member is required.
	Statistic MetricStatistic

	// The dimensions of the metric. Conditional: If you published your metric with
	// dimensions, you must specify the same dimensions in your customized scaling
	// metric specification.
	Dimensions []MetricDimension

	// The unit of the metric.
	Unit *string

	noSmithyDocumentSerde
}

// Represents a single value in the forecast data used for predictive scaling.
type Datapoint struct {

	// The time stamp for the data point in UTC format.
	Timestamp *time.Time

	// The value of the data point.
	Value *float64

	noSmithyDocumentSerde
}

// Represents a dimension for a customized metric.
type MetricDimension struct {

	// The name of the dimension.
	//
	// This member is required.
	Name *string

	// The value of the dimension.
	//
	// This member is required.
	Value *string

	noSmithyDocumentSerde
}

// Represents a predefined metric that can be used for predictive scaling. After
// creating your scaling plan, you can use the AWS Auto Scaling console to
// visualize forecasts for the specified metric. For more information, see View
// Scaling Information for a Resource (https://docs.aws.amazon.com/autoscaling/plans/userguide/gs-create-scaling-plan.html#gs-view-resource)
// in the AWS Auto Scaling User Guide.
type PredefinedLoadMetricSpecification struct {

	// The metric type.
	//
	// This member is required.
	PredefinedLoadMetricType LoadMetricType

	// Identifies the resource associated with the metric type. You can't specify a
	// resource label unless the metric type is ALBTargetGroupRequestCount and there
	// is a target group for an Application Load Balancer attached to the Auto Scaling
	// group. You create the resource label by appending the final portion of the load
	// balancer ARN and the final portion of the target group ARN into a single value,
	// separated by a forward slash (/). The format is app///targetgroup//, where:
	//   - app// is the final portion of the load balancer ARN
	//   - targetgroup// is the final portion of the target group ARN.
	// This is an example:
	// app/EC2Co-EcsEl-1TKLTMITMM0EO/f37c06a68c1748aa/targetgroup/EC2Co-Defau-LDNM7Q3ZH1ZN/6d4ea56ca2d6a18d.
	// To find the ARN for an Application Load Balancer, use the DescribeLoadBalancers (https://docs.aws.amazon.com/elasticloadbalancing/latest/APIReference/API_DescribeLoadBalancers.html)
	// API operation. To find the ARN for the target group, use the
	// DescribeTargetGroups (https://docs.aws.amazon.com/elasticloadbalancing/latest/APIReference/API_DescribeTargetGroups.html)
	// API operation.
	ResourceLabel *string

	noSmithyDocumentSerde
}

// Represents a predefined metric that can be used for dynamic scaling as part of
// a target tracking scaling policy.
type PredefinedScalingMetricSpecification struct {

	// The metric type. The ALBRequestCountPerTarget metric type applies only to Auto
	// Scaling groups, Spot Fleet requests, and ECS services.
	//
	// This member is required.
	PredefinedScalingMetricType ScalingMetricType

	// Identifies the resource associated with the metric type. You can't specify a
	// resource label unless the metric type is ALBRequestCountPerTarget and there is
	// a target group for an Application Load Balancer attached to the Auto Scaling
	// group, Spot Fleet request, or ECS service. You create the resource label by
	// appending the final portion of the load balancer ARN and the final portion of
	// the target group ARN into a single value, separated by a forward slash (/). The
	// format is app///targetgroup//, where:
	//   - app// is the final portion of the load balancer ARN
	//   - targetgroup// is the final portion of the target group ARN.
	// This is an example:
	// app/EC2Co-EcsEl-1TKLTMITMM0EO/f37c06a68c1748aa/targetgroup/EC2Co-Defau-LDNM7Q3ZH1ZN/6d4ea56ca2d6a18d.
	// To find the ARN for an Application Load Balancer, use the DescribeLoadBalancers (https://docs.aws.amazon.com/elasticloadbalancing/latest/APIReference/API_DescribeLoadBalancers.html)
	// API operation. To find the ARN for the target group, use the
	// DescribeTargetGroups (https://docs.aws.amazon.com/elasticloadbalancing/latest/APIReference/API_DescribeTargetGroups.html)
	// API operation.
	ResourceLabel *string

	noSmithyDocumentSerde
}

// Describes a scaling instruction for a scalable resource in a scaling plan. Each
// scaling instruction applies to one resource. AWS Auto Scaling creates target
// tracking scaling policies based on the scaling instructions. Target tracking
// scaling policies adjust the capacity of your scalable resource as required to
// maintain resource utilization at the target value that you specified. AWS Auto
// Scaling also configures predictive scaling for your Amazon EC2 Auto Scaling
// groups using a subset of parameters, including the load metric, the scaling
// metric, the target value for the scaling metric, the predictive scaling mode
// (forecast and scale or forecast only), and the desired behavior when the
// forecast capacity exceeds the maximum capacity of the resource. With predictive
// scaling, AWS Auto Scaling generates forecasts with traffic predictions for the
// two days ahead and schedules scaling actions that proactively add and remove
// resource capacity to match the forecast. We recommend waiting a minimum of 24
// hours after creating an Auto Scaling group to configure predictive scaling. At
// minimum, there must be 24 hours of historical data to generate a forecast. For
// more information, see Best Practices for AWS Auto Scaling (https://docs.aws.amazon.com/autoscaling/plans/userguide/gs-best-practices.html)
// in the AWS Auto Scaling User Guide.
type ScalingInstruction struct {

	// The maximum capacity of the resource. The exception to this upper limit is if
	// you specify a non-default setting for PredictiveScalingMaxCapacityBehavior.
	//
	// This member is required.
	MaxCapacity *int32

	// The minimum capacity of the resource.
	//
	// This member is required.
	MinCapacity *int32

	// The ID of the resource. This string consists of the resource type and unique
	// identifier.
	//   - Auto Scaling group - The resource type is autoScalingGroup and the unique
	//   identifier is the name of the Auto Scaling group. Example:
	//   autoScalingGroup/my-asg .
	//   - ECS service - The resource type is service and the unique identifier is the
	//   cluster name and service name. Example: service/default/sample-webapp .
	//   - Spot Fleet request - The resource type is spot-fleet-request and the unique
	//   identifier is the Spot Fleet request ID. Example:
	//   spot-fleet-request/sfr-73fbd2ce-aa30-494c-8788-1cee4EXAMPLE .
	//   - DynamoDB table - The resource type is table and the unique identifier is the
	//   resource ID. Example: table/my-table .
	//   - DynamoDB global secondary index - The resource type is index and the unique
	//   identifier is the resource ID. Example: table/my-table/index/my-table-index .
	//   - Aurora DB cluster - The resource type is cluster and the unique identifier
	//   is the cluster name. Example: cluster:my-db-cluster .
	//
	// This member is required.
	ResourceId *string

	// The scalable dimension associated with the resource.
	//   - autoscaling:autoScalingGroup:DesiredCapacity - The desired capacity of an
	//   Auto Scaling group.
	//   - ecs:service:DesiredCount - The desired task count of an ECS service.
	//   - ec2:spot-fleet-request:TargetCapacity - The target capacity of a Spot Fleet
	//   request.
	//   - dynamodb:table:ReadCapacityUnits - The provisioned read capacity for a
	//   DynamoDB table.
	//   - dynamodb:table:WriteCapacityUnits - The provisioned write capacity for a
	//   DynamoDB table.
	//   - dynamodb:index:ReadCapacityUnits - The provisioned read capacity for a
	//   DynamoDB global secondary index.
	//   - dynamodb:index:WriteCapacityUnits - The provisioned write capacity for a
	//   DynamoDB global secondary index.
	//   - rds:cluster:ReadReplicaCount - The count of Aurora Replicas in an Aurora DB
	//   cluster. Available for Aurora MySQL-compatible edition and Aurora
	//   PostgreSQL-compatible edition.
	//
	// This member is required.
	ScalableDimension ScalableDimension

	// The namespace of the AWS service.
	//
	// This member is required.
	ServiceNamespace ServiceNamespace

	// The target tracking configurations (up to 10). Each of these structures must
	// specify a unique scaling metric and a target value for the metric.
	//
	// This member is required.
	TargetTrackingConfigurations []TargetTrackingConfiguration

	// The customized load metric to use for predictive scaling. This parameter or a
	// PredefinedLoadMetricSpecification is required when configuring predictive
	// scaling, and cannot be used otherwise.
	CustomizedLoadMetricSpecification *CustomizedLoadMetricSpecification

	// Controls whether dynamic scaling by AWS Auto Scaling is disabled. When dynamic
	// scaling is enabled, AWS Auto Scaling creates target tracking scaling policies
	// based on the specified target tracking configurations. The default is enabled (
	// false ).
	DisableDynamicScaling *bool

	// The predefined load metric to use for predictive scaling. This parameter or a
	// CustomizedLoadMetricSpecification is required when configuring predictive
	// scaling, and cannot be used otherwise.
	PredefinedLoadMetricSpecification *PredefinedLoadMetricSpecification

	// Defines the behavior that should be applied if the forecast capacity approaches
	// or exceeds the maximum capacity specified for the resource. The default value is
	// SetForecastCapacityToMaxCapacity . The following are possible values:
	//   - SetForecastCapacityToMaxCapacity - AWS Auto Scaling cannot scale resource
	//   capacity higher than the maximum capacity. The maximum capacity is enforced as a
	//   hard limit.
	//   - SetMaxCapacityToForecastCapacity - AWS Auto Scaling may scale resource
	//   capacity higher than the maximum capacity to equal but not exceed forecast
	//   capacity.
	//   - SetMaxCapacityAboveForecastCapacity - AWS Auto Scaling may scale resource
	//   capacity higher than the maximum capacity by a specified buffer value. The
	//   intention is to give the target tracking scaling policy extra capacity if
	//   unexpected traffic occurs.
	// Only valid when configuring predictive scaling.
	PredictiveScalingMaxCapacityBehavior PredictiveScalingMaxCapacityBehavior

	// The size of the capacity buffer to use when the forecast capacity is close to
	// or exceeds the maximum capacity. The value is specified as a percentage relative
	// to the forecast capacity. For example, if the buffer is 10, this means a 10
	// percent buffer, such that if the forecast capacity is 50, and the maximum
	// capacity is 40, then the effective maximum capacity is 55. Only valid when
	// configuring predictive scaling. Required if the
	// PredictiveScalingMaxCapacityBehavior is set to
	// SetMaxCapacityAboveForecastCapacity , and cannot be used otherwise. The range is
	// 1-100.
	PredictiveScalingMaxCapacityBuffer *int32

	// The predictive scaling mode. The default value is ForecastAndScale . Otherwise,
	// AWS Auto Scaling forecasts capacity but does not create any scheduled scaling
	// actions based on the capacity forecast.
	PredictiveScalingMode PredictiveScalingMode

	// Controls whether a resource's externally created scaling policies are kept or
	// replaced. The default value is KeepExternalPolicies . If the parameter is set to
	// ReplaceExternalPolicies , any scaling policies that are external to AWS Auto
	// Scaling are deleted and new target tracking scaling policies created. Only valid
	// when configuring dynamic scaling. Condition: The number of existing policies to
	// be replaced must be less than or equal to 50. If there are more than 50 policies
	// to be replaced, AWS Auto Scaling keeps all existing policies and does not create
	// new ones.
	ScalingPolicyUpdateBehavior ScalingPolicyUpdateBehavior

	// The amount of time, in seconds, to buffer the run time of scheduled scaling
	// actions when scaling out. For example, if the forecast says to add capacity at
	// 10:00 AM, and the buffer time is 5 minutes, then the run time of the
	// corresponding scheduled scaling action will be 9:55 AM. The intention is to give
	// resources time to be provisioned. For example, it can take a few minutes to
	// launch an EC2 instance. The actual amount of time required depends on several
	// factors, such as the size of the instance and whether there are startup scripts
	// to complete. The value must be less than the forecast interval duration of 3600
	// seconds (60 minutes). The default is 300 seconds. Only valid when configuring
	// predictive scaling.
	ScheduledActionBufferTime *int32

	noSmithyDocumentSerde
}

// Represents a scaling plan.
type ScalingPlan struct {

	// A CloudFormation stack or a set of tags. You can create one scaling plan per
	// application source.
	//
	// This member is required.
	ApplicationSource *ApplicationSource

	// The scaling instructions.
	//
	// This member is required.
	ScalingInstructions []ScalingInstruction

	// The name of the scaling plan.
	//
	// This member is required.
	ScalingPlanName *string

	// The version number of the scaling plan.
	//
	// This member is required.
	ScalingPlanVersion *int64

	// The status of the scaling plan.
	//   - Active - The scaling plan is active.
	//   - ActiveWithProblems - The scaling plan is active, but the scaling
	//   configuration for one or more resources could not be applied.
	//   - CreationInProgress - The scaling plan is being created.
	//   - CreationFailed - The scaling plan could not be created.
	//   - DeletionInProgress - The scaling plan is being deleted.
	//   - DeletionFailed - The scaling plan could not be deleted.
	//   - UpdateInProgress - The scaling plan is being updated.
	//   - UpdateFailed - The scaling plan could not be updated.
	//
	// This member is required.
	StatusCode ScalingPlanStatusCode

	// The Unix time stamp when the scaling plan was created.
	CreationTime *time.Time

	// A simple message about the current status of the scaling plan.
	StatusMessage *string

	// The Unix time stamp when the scaling plan entered the current status.
	StatusStartTime *time.Time

	noSmithyDocumentSerde
}

// Represents a scalable resource.
type ScalingPlanResource struct {

	// The ID of the resource. This string consists of the resource type and unique
	// identifier.
	//   - Auto Scaling group - The resource type is autoScalingGroup and the unique
	//   identifier is the name of the Auto Scaling group. Example:
	//   autoScalingGroup/my-asg .
	//   - ECS service - The resource type is service and the unique identifier is the
	//   cluster name and service name. Example: service/default/sample-webapp .
	//   - Spot Fleet request - The resource type is spot-fleet-request and the unique
	//   identifier is the Spot Fleet request ID. Example:
	//   spot-fleet-request/sfr-73fbd2ce-aa30-494c-8788-1cee4EXAMPLE .
	//   - DynamoDB table - The resource type is table and the unique identifier is the
	//   resource ID. Example: table/my-table .
	//   - DynamoDB global secondary index - The resource type is index and the unique
	//   identifier is the resource ID. Example: table/my-table/index/my-table-index .
	//   - Aurora DB cluster - The resource type is cluster and the unique identifier
	//   is the cluster name. Example: cluster:my-db-cluster .
	//
	// This member is required.
	ResourceId *string

	// The scalable dimension for the resource.
	//   - autoscaling:autoScalingGroup:DesiredCapacity - The desired capacity of an
	//   Auto Scaling group.
	//   - ecs:service:DesiredCount - The desired task count of an ECS service.
	//   - ec2:spot-fleet-request:TargetCapacity - The target capacity of a Spot Fleet
	//   request.
	//   - dynamodb:table:ReadCapacityUnits - The provisioned read capacity for a
	//   DynamoDB table.
	//   - dynamodb:table:WriteCapacityUnits - The provisioned write capacity for a
	//   DynamoDB table.
	//   - dynamodb:index:ReadCapacityUnits - The provisioned read capacity for a
	//   DynamoDB global secondary index.
	//   - dynamodb:index:WriteCapacityUnits - The provisioned write capacity for a
	//   DynamoDB global secondary index.
	//   - rds:cluster:ReadReplicaCount - The count of Aurora Replicas in an Aurora DB
	//   cluster. Available for Aurora MySQL-compatible edition and Aurora
	//   PostgreSQL-compatible edition.
	//
	// This member is required.
	ScalableDimension ScalableDimension

	// The name of the scaling plan.
	//
	// This member is required.
	ScalingPlanName *string

	// The version number of the scaling plan.
	//
	// This member is required.
	ScalingPlanVersion *int64

	// The scaling status of the resource.
	//   - Active - The scaling configuration is active.
	//   - Inactive - The scaling configuration is not active because the scaling plan
	//   is being created or the scaling configuration could not be applied. Check the
	//   status message for more information.
	//   - PartiallyActive - The scaling configuration is partially active because the
	//   scaling plan is being created or deleted or the scaling configuration could not
	//   be fully applied. Check the status message for more information.
	//
	// This member is required.
	ScalingStatusCode ScalingStatusCode

	// The namespace of the AWS service.
	//
	// This member is required.
	ServiceNamespace ServiceNamespace

	// The scaling policies.
	ScalingPolicies []ScalingPolicy

	// A simple message about the current scaling status of the resource.
	ScalingStatusMessage *string

	noSmithyDocumentSerde
}

// Represents a scaling policy.
type ScalingPolicy struct {

	// The name of the scaling policy.
	//
	// This member is required.
	PolicyName *string

	// The type of scaling policy.
	//
	// This member is required.
	PolicyType PolicyType

	// The target tracking scaling policy. Includes support for predefined or
	// customized metrics.
	TargetTrackingConfiguration *TargetTrackingConfiguration

	noSmithyDocumentSerde
}

// Represents a tag.
type TagFilter struct {

	// The tag key.
	Key *string

	// The tag values (0 to 20).
	Values []string

	noSmithyDocumentSerde
}

// Describes a target tracking configuration to use with AWS Auto Scaling. Used
// with ScalingInstruction and ScalingPolicy .
type TargetTrackingConfiguration struct {

	// The target value for the metric. Although this property accepts numbers of type
	// Double, it won't accept values that are either too small or too large. Values
	// must be in the range of -2^360 to 2^360.
	//
	// This member is required.
	TargetValue *float64

	// A customized metric. You can specify either a predefined metric or a customized
	// metric.
	CustomizedScalingMetricSpecification *CustomizedScalingMetricSpecification

	// Indicates whether scale in by the target tracking scaling policy is disabled.
	// If the value is true , scale in is disabled and the target tracking scaling
	// policy doesn't remove capacity from the scalable resource. Otherwise, scale in
	// is enabled and the target tracking scaling policy can remove capacity from the
	// scalable resource. The default value is false .
	DisableScaleIn *bool

	// The estimated time, in seconds, until a newly launched instance can contribute
	// to the CloudWatch metrics. This value is used only if the resource is an Auto
	// Scaling group.
	EstimatedInstanceWarmup *int32

	// A predefined metric. You can specify either a predefined metric or a customized
	// metric.
	PredefinedScalingMetricSpecification *PredefinedScalingMetricSpecification

	// The amount of time, in seconds, after a scale-in activity completes before
	// another scale-in activity can start. This property is not used if the scalable
	// resource is an Auto Scaling group. With the scale-in cooldown period, the
	// intention is to scale in conservatively to protect your application’s
	// availability, so scale-in activities are blocked until the cooldown period has
	// expired. However, if another alarm triggers a scale-out activity during the
	// scale-in cooldown period, Auto Scaling scales out the target immediately. In
	// this case, the scale-in cooldown period stops and doesn't complete.
	ScaleInCooldown *int32

	// The amount of time, in seconds, to wait for a previous scale-out activity to
	// take effect. This property is not used if the scalable resource is an Auto
	// Scaling group. With the scale-out cooldown period, the intention is to
	// continuously (but not excessively) scale out. After Auto Scaling successfully
	// scales out using a target tracking scaling policy, it starts to calculate the
	// cooldown time. The scaling policy won't increase the desired capacity again
	// unless either a larger scale out is triggered or the cooldown period ends.
	ScaleOutCooldown *int32

	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde
