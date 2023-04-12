// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/service/tnb/document"
	smithydocument "github.com/aws/smithy-go/document"
	"time"
)

// Provides error information.
type ErrorInfo struct {

	// Error cause.
	Cause *string

	// Error details.
	Details *string

	noSmithyDocumentSerde
}

// Metadata for function package artifacts. Artifacts are the contents of the
// package descriptor file and the state of the package.
type FunctionArtifactMeta struct {

	// Lists of function package overrides.
	Overrides []ToscaOverride

	noSmithyDocumentSerde
}

// The metadata of a network function instance. A network function instance is a
// function in a function package .
type GetSolFunctionInstanceMetadata struct {

	// The date that the resource was created.
	//
	// This member is required.
	CreatedAt *time.Time

	// The date that the resource was last modified.
	//
	// This member is required.
	LastModified *time.Time

	noSmithyDocumentSerde
}

// Metadata related to the function package. A function package is a .zip file in
// CSAR (Cloud Service Archive) format that contains a network function (an ETSI
// standard telecommunication application) and function package descriptor that
// uses the TOSCA standard to describe how the network functions should run on your
// network.
type GetSolFunctionPackageMetadata struct {

	// The date that the resource was created.
	//
	// This member is required.
	CreatedAt *time.Time

	// The date that the resource was last modified.
	//
	// This member is required.
	LastModified *time.Time

	// Metadata related to the function package descriptor of the function package.
	Vnfd *FunctionArtifactMeta

	noSmithyDocumentSerde
}

// Information about a network function. A network instance is a single network
// created in Amazon Web Services TNB that can be deployed and on which life-cycle
// operations (like terminate, update, and delete) can be performed.
type GetSolInstantiatedVnfInfo struct {

	// State of the network function.
	VnfState VnfOperationalState

	noSmithyDocumentSerde
}

// The metadata of a network instance. A network instance is a single network
// created in Amazon Web Services TNB that can be deployed and on which life-cycle
// operations (like terminate, update, and delete) can be performed.
type GetSolNetworkInstanceMetadata struct {

	// The date that the resource was created.
	//
	// This member is required.
	CreatedAt *time.Time

	// The date that the resource was last modified.
	//
	// This member is required.
	LastModified *time.Time

	noSmithyDocumentSerde
}

// Metadata related to a network operation occurrence. A network operation is any
// operation that is done to your network, such as network instance instantiation
// or termination.
type GetSolNetworkOperationMetadata struct {

	// The date that the resource was created.
	//
	// This member is required.
	CreatedAt *time.Time

	// The date that the resource was last modified.
	//
	// This member is required.
	LastModified *time.Time

	noSmithyDocumentSerde
}

// Gets the details of a network operation. A network operation is any operation
// that is done to your network, such as network instance instantiation or
// termination.
type GetSolNetworkOperationTaskDetails struct {

	// Context for the network operation task.
	TaskContext map[string]string

	// Task end time.
	TaskEndTime *time.Time

	// Task error details.
	TaskErrorDetails *ErrorInfo

	// Task name.
	TaskName *string

	// Task start time.
	TaskStartTime *time.Time

	// Task status.
	TaskStatus TaskStatus

	noSmithyDocumentSerde
}

// Metadata associated with a network package. A network package is a .zip file in
// CSAR (Cloud Service Archive) format defines the function packages you want to
// deploy and the Amazon Web Services infrastructure you want to deploy them on.
type GetSolNetworkPackageMetadata struct {

	// The date that the resource was created.
	//
	// This member is required.
	CreatedAt *time.Time

	// The date that the resource was last modified.
	//
	// This member is required.
	LastModified *time.Time

	// Metadata related to the onboarded network service descriptor in the network
	// package.
	Nsd *NetworkArtifactMeta

	noSmithyDocumentSerde
}

// Details of resource associated with a network function. A network instance is a
// single network created in Amazon Web Services TNB that can be deployed and on
// which life-cycle operations (like terminate, update, and delete) can be
// performed.
type GetSolVnfcResourceInfo struct {

	// The metadata of the network function compute.
	Metadata *GetSolVnfcResourceInfoMetadata

	noSmithyDocumentSerde
}

// The metadata of a network function. A network instance is a single network
// created in Amazon Web Services TNB that can be deployed and on which life-cycle
// operations (like terminate, update, and delete) can be performed.
type GetSolVnfcResourceInfoMetadata struct {

	// Information about the cluster.
	Cluster *string

	// Information about the helm chart.
	HelmChart *string

	// Information about the node group.
	NodeGroup *string

	noSmithyDocumentSerde
}

// Information about the network function. A network function instance is a
// function in a function package .
type GetSolVnfInfo struct {

	// State of the network function instance.
	VnfState VnfOperationalState

	// Compute info used by the network function instance.
	VnfcResourceInfo []GetSolVnfcResourceInfo

	noSmithyDocumentSerde
}

// Lifecycle management operation details on the network instance. Lifecycle
// management operations are deploy, update, or delete operations.
type LcmOperationInfo struct {

	// The identifier of the network operation.
	//
	// This member is required.
	NsLcmOpOccId *string

	noSmithyDocumentSerde
}

// Lists information about a network function instance. A network function
// instance is a function in a function package .
type ListSolFunctionInstanceInfo struct {

	// Network function instance ARN.
	//
	// This member is required.
	Arn *string

	// Network function instance ID.
	//
	// This member is required.
	Id *string

	// Network function instance instantiation state.
	//
	// This member is required.
	InstantiationState VnfInstantiationState

	// Network function instance metadata.
	//
	// This member is required.
	Metadata *ListSolFunctionInstanceMetadata

	// Network instance ID.
	//
	// This member is required.
	NsInstanceId *string

	// Function package ID.
	//
	// This member is required.
	VnfPkgId *string

	// Information about a network function. A network instance is a single network
	// created in Amazon Web Services TNB that can be deployed and on which life-cycle
	// operations (like terminate, update, and delete) can be performed.
	InstantiatedVnfInfo *GetSolInstantiatedVnfInfo

	// Function package name.
	VnfPkgName *string

	noSmithyDocumentSerde
}

// Lists network function instance metadata. A network function instance is a
// function in a function package .
type ListSolFunctionInstanceMetadata struct {

	// When the network function instance was created.
	//
	// This member is required.
	CreatedAt *time.Time

	// When the network function instance was last modified.
	//
	// This member is required.
	LastModified *time.Time

	noSmithyDocumentSerde
}

// Information about a function package. A function package is a .zip file in CSAR
// (Cloud Service Archive) format that contains a network function (an ETSI
// standard telecommunication application) and function package descriptor that
// uses the TOSCA standard to describe how the network functions should run on your
// network.
type ListSolFunctionPackageInfo struct {

	// Function package ARN.
	//
	// This member is required.
	Arn *string

	// ID of the function package.
	//
	// This member is required.
	Id *string

	// Onboarding state of the function package.
	//
	// This member is required.
	OnboardingState OnboardingState

	// Operational state of the function package.
	//
	// This member is required.
	OperationalState OperationalState

	// Usage state of the function package.
	//
	// This member is required.
	UsageState UsageState

	// The metadata of the function package.
	Metadata *ListSolFunctionPackageMetadata

	// The product name for the network function.
	VnfProductName *string

	// Provider of the function package and the function package descriptor.
	VnfProvider *string

	// Identifies the function package and the function package descriptor.
	VnfdId *string

	// Identifies the version of the function package descriptor.
	VnfdVersion *string

	noSmithyDocumentSerde
}

// Details for the function package metadata. A function package is a .zip file in
// CSAR (Cloud Service Archive) format that contains a network function (an ETSI
// standard telecommunication application) and function package descriptor that
// uses the TOSCA standard to describe how the network functions should run on your
// network.
type ListSolFunctionPackageMetadata struct {

	// The date that the resource was created.
	//
	// This member is required.
	CreatedAt *time.Time

	// The date that the resource was last modified.
	//
	// This member is required.
	LastModified *time.Time

	noSmithyDocumentSerde
}

// Info about the specific network instance. A network instance is a single
// network created in Amazon Web Services TNB that can be deployed and on which
// life-cycle operations (like terminate, update, and delete) can be performed.
type ListSolNetworkInstanceInfo struct {

	// Network instance ARN.
	//
	// This member is required.
	Arn *string

	// ID of the network instance.
	//
	// This member is required.
	Id *string

	// The metadata of the network instance.
	//
	// This member is required.
	Metadata *ListSolNetworkInstanceMetadata

	// Human-readable description of the network instance.
	//
	// This member is required.
	NsInstanceDescription *string

	// Human-readable name of the network instance.
	//
	// This member is required.
	NsInstanceName *string

	// The state of the network instance.
	//
	// This member is required.
	NsState NsState

	// ID of the network service descriptor in the network package.
	//
	// This member is required.
	NsdId *string

	// ID of the network service descriptor in the network package.
	//
	// This member is required.
	NsdInfoId *string

	noSmithyDocumentSerde
}

// Metadata details for a network instance. A network instance is a single network
// created in Amazon Web Services TNB that can be deployed and on which life-cycle
// operations (like terminate, update, and delete) can be performed.
type ListSolNetworkInstanceMetadata struct {

	// The date that the resource was created.
	//
	// This member is required.
	CreatedAt *time.Time

	// The date that the resource was last modified.
	//
	// This member is required.
	LastModified *time.Time

	noSmithyDocumentSerde
}

// Information parameters for a network operation.
type ListSolNetworkOperationsInfo struct {

	// Network operation ARN.
	//
	// This member is required.
	Arn *string

	// ID of this network operation.
	//
	// This member is required.
	Id *string

	// Type of lifecycle management network operation.
	//
	// This member is required.
	LcmOperationType LcmOperationType

	// ID of the network instance related to this operation.
	//
	// This member is required.
	NsInstanceId *string

	// The state of the network operation.
	//
	// This member is required.
	OperationState NsLcmOperationState

	// Error related to this specific network operation.
	Error *ProblemDetails

	// Metadata related to this network operation.
	Metadata *ListSolNetworkOperationsMetadata

	noSmithyDocumentSerde
}

// Metadata related to a network operation. A network operation is any operation
// that is done to your network, such as network instance instantiation or
// termination.
type ListSolNetworkOperationsMetadata struct {

	// The date that the resource was created.
	//
	// This member is required.
	CreatedAt *time.Time

	// The date that the resource was last modified.
	//
	// This member is required.
	LastModified *time.Time

	noSmithyDocumentSerde
}

// Details of a network package. A network package is a .zip file in CSAR (Cloud
// Service Archive) format defines the function packages you want to deploy and the
// Amazon Web Services infrastructure you want to deploy them on.
type ListSolNetworkPackageInfo struct {

	// Network package ARN.
	//
	// This member is required.
	Arn *string

	// ID of the individual network package.
	//
	// This member is required.
	Id *string

	// The metadata of the network package.
	//
	// This member is required.
	Metadata *ListSolNetworkPackageMetadata

	// Onboarding state of the network service descriptor in the network package.
	//
	// This member is required.
	NsdOnboardingState NsdOnboardingState

	// Operational state of the network service descriptor in the network package.
	//
	// This member is required.
	NsdOperationalState NsdOperationalState

	// Usage state of the network service descriptor in the network package.
	//
	// This member is required.
	NsdUsageState NsdUsageState

	// Designer of the onboarded network service descriptor in the network package.
	NsdDesigner *string

	// ID of the network service descriptor on which the network package is based.
	NsdId *string

	// Identifies a network service descriptor in a version independent manner.
	NsdInvariantId *string

	// Name of the onboarded network service descriptor in the network package.
	NsdName *string

	// Version of the onboarded network service descriptor in the network package.
	NsdVersion *string

	// Identifies the function package for the function package descriptor referenced
	// by the onboarded network package.
	VnfPkgIds []string

	noSmithyDocumentSerde
}

// Metadata related to a network package. A network package is a .zip file in CSAR
// (Cloud Service Archive) format defines the function packages you want to deploy
// and the Amazon Web Services infrastructure you want to deploy them on.
type ListSolNetworkPackageMetadata struct {

	// The date that the resource was created.
	//
	// This member is required.
	CreatedAt *time.Time

	// The date that the resource was last modified.
	//
	// This member is required.
	LastModified *time.Time

	noSmithyDocumentSerde
}

// Metadata for network package artifacts. Artifacts are the contents of the
// package descriptor file and the state of the package.
type NetworkArtifactMeta struct {

	// Lists network package overrides.
	Overrides []ToscaOverride

	noSmithyDocumentSerde
}

// Details related to problems with AWS TNB resources.
type ProblemDetails struct {

	// A human-readable explanation specific to this occurrence of the problem.
	//
	// This member is required.
	Detail *string

	// A human-readable title of the problem type.
	Title *string

	noSmithyDocumentSerde
}

// Update metadata in a function package. A function package is a .zip file in
// CSAR (Cloud Service Archive) format that contains a network function (an ETSI
// standard telecommunication application) and function package descriptor that
// uses the TOSCA standard to describe how the network functions should run on your
// network.
type PutSolFunctionPackageContentMetadata struct {

	// Metadata for function package artifacts. Artifacts are the contents of the
	// package descriptor file and the state of the package.
	Vnfd *FunctionArtifactMeta

	noSmithyDocumentSerde
}

// Update metadata in a network package. A network package is a .zip file in CSAR
// (Cloud Service Archive) format defines the function packages you want to deploy
// and the Amazon Web Services infrastructure you want to deploy them on.
type PutSolNetworkPackageContentMetadata struct {

	// Metadata for network package artifacts. Artifacts are the contents of the
	// package descriptor file and the state of the package.
	Nsd *NetworkArtifactMeta

	noSmithyDocumentSerde
}

// Overrides of the TOSCA node.
type ToscaOverride struct {

	// Default value for the override.
	DefaultValue *string

	// Name of the TOSCA override.
	Name *string

	noSmithyDocumentSerde
}

// Information parameters and/or the configurable properties for a network
// function. A network function instance is a function in a function package .
type UpdateSolNetworkModify struct {

	// Provides values for the configurable properties declared in the function
	// package descriptor.
	//
	// This member is required.
	VnfConfigurableProperties document.Interface

	// ID of the network function instance. A network function instance is a function
	// in a function package .
	//
	// This member is required.
	VnfInstanceId *string

	noSmithyDocumentSerde
}

// Validates function package content metadata. A function package is a .zip file
// in CSAR (Cloud Service Archive) format that contains a network function (an ETSI
// standard telecommunication application) and function package descriptor that
// uses the TOSCA standard to describe how the network functions should run on your
// network.
type ValidateSolFunctionPackageContentMetadata struct {

	// Metadata for function package artifacts. Artifacts are the contents of the
	// package descriptor file and the state of the package.
	Vnfd *FunctionArtifactMeta

	noSmithyDocumentSerde
}

// Validates network package content metadata. A network package is a .zip file in
// CSAR (Cloud Service Archive) format defines the function packages you want to
// deploy and the Amazon Web Services infrastructure you want to deploy them on.
type ValidateSolNetworkPackageContentMetadata struct {

	// Metadata for network package artifacts. Artifacts are the contents of the
	// package descriptor file and the state of the package.
	Nsd *NetworkArtifactMeta

	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde
