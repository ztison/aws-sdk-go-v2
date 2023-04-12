// Code generated by smithy-go-codegen DO NOT EDIT.

// Package support provides the API client, operations, and parameter types for
// AWS Support.
//
// Amazon Web Services Support The Amazon Web Services Support API Reference is
// intended for programmers who need detailed information about the Amazon Web
// Services Support operations and data types. You can use the API to manage your
// support cases programmatically. The Amazon Web Services Support API uses HTTP
// methods that return results in JSON format.
//   - You must have a Business, Enterprise On-Ramp, or Enterprise Support plan to
//     use the Amazon Web Services Support API.
//   - If you call the Amazon Web Services Support API from an account that
//     doesn't have a Business, Enterprise On-Ramp, or Enterprise Support plan, the
//     SubscriptionRequiredException error message appears. For information about
//     changing your support plan, see Amazon Web Services Support (http://aws.amazon.com/premiumsupport/)
//     .
//
// You can also use the Amazon Web Services Support API to access features for
// Trusted Advisor (http://aws.amazon.com/premiumsupport/trustedadvisor/) . You can
// return a list of checks and their descriptions, get check results, specify
// checks to refresh, and get the refresh status of checks. You can manage your
// support cases with the following Amazon Web Services Support API operations:
//   - The CreateCase , DescribeCases , DescribeAttachment , and ResolveCase
//     operations create Amazon Web Services Support cases, retrieve information about
//     cases, and resolve cases.
//   - The DescribeCommunications , AddCommunicationToCase , and
//     AddAttachmentsToSet operations retrieve and add communications and attachments
//     to Amazon Web Services Support cases.
//   - The DescribeServices and DescribeSeverityLevels operations return Amazon Web
//     Service names, service codes, service categories, and problem severity levels.
//     You use these values when you call the CreateCase operation.
//
// You can also use the Amazon Web Services Support API to call the Trusted
// Advisor operations. For more information, see Trusted Advisor (https://docs.aws.amazon.com/)
// in the Amazon Web Services Support User Guide. For authentication of requests,
// Amazon Web Services Support uses Signature Version 4 Signing Process (https://docs.aws.amazon.com/general/latest/gr/signature-version-4.html)
// . For more information about this service and the endpoints to use, see About
// the Amazon Web Services Support API (https://docs.aws.amazon.com/awssupport/latest/user/about-support-api.html)
// in the Amazon Web Services Support User Guide.
package support
