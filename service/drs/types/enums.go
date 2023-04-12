// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type DataReplicationErrorString string

// Enum values for DataReplicationErrorString
const (
	DataReplicationErrorStringAgentNotSeen                            DataReplicationErrorString = "AGENT_NOT_SEEN"
	DataReplicationErrorStringSnapshotsFailure                        DataReplicationErrorString = "SNAPSHOTS_FAILURE"
	DataReplicationErrorStringNotConverging                           DataReplicationErrorString = "NOT_CONVERGING"
	DataReplicationErrorStringUnstableNetwork                         DataReplicationErrorString = "UNSTABLE_NETWORK"
	DataReplicationErrorStringFailedToCreateSecurityGroup             DataReplicationErrorString = "FAILED_TO_CREATE_SECURITY_GROUP"
	DataReplicationErrorStringFailedToLaunchReplicationServer         DataReplicationErrorString = "FAILED_TO_LAUNCH_REPLICATION_SERVER"
	DataReplicationErrorStringFailedToBootReplicationServer           DataReplicationErrorString = "FAILED_TO_BOOT_REPLICATION_SERVER"
	DataReplicationErrorStringFailedToAuthenticateWithService         DataReplicationErrorString = "FAILED_TO_AUTHENTICATE_WITH_SERVICE"
	DataReplicationErrorStringFailedToDownloadReplicationSoftware     DataReplicationErrorString = "FAILED_TO_DOWNLOAD_REPLICATION_SOFTWARE"
	DataReplicationErrorStringFailedToCreateStagingDisks              DataReplicationErrorString = "FAILED_TO_CREATE_STAGING_DISKS"
	DataReplicationErrorStringFailedToAttachStagingDisks              DataReplicationErrorString = "FAILED_TO_ATTACH_STAGING_DISKS"
	DataReplicationErrorStringFailedToPairReplicationServerWithAgent  DataReplicationErrorString = "FAILED_TO_PAIR_REPLICATION_SERVER_WITH_AGENT"
	DataReplicationErrorStringFailedToConnectAgentToReplicationServer DataReplicationErrorString = "FAILED_TO_CONNECT_AGENT_TO_REPLICATION_SERVER"
	DataReplicationErrorStringFailedToStartDataTransfer               DataReplicationErrorString = "FAILED_TO_START_DATA_TRANSFER"
)

// Values returns all known values for DataReplicationErrorString. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (DataReplicationErrorString) Values() []DataReplicationErrorString {
	return []DataReplicationErrorString{
		"AGENT_NOT_SEEN",
		"SNAPSHOTS_FAILURE",
		"NOT_CONVERGING",
		"UNSTABLE_NETWORK",
		"FAILED_TO_CREATE_SECURITY_GROUP",
		"FAILED_TO_LAUNCH_REPLICATION_SERVER",
		"FAILED_TO_BOOT_REPLICATION_SERVER",
		"FAILED_TO_AUTHENTICATE_WITH_SERVICE",
		"FAILED_TO_DOWNLOAD_REPLICATION_SOFTWARE",
		"FAILED_TO_CREATE_STAGING_DISKS",
		"FAILED_TO_ATTACH_STAGING_DISKS",
		"FAILED_TO_PAIR_REPLICATION_SERVER_WITH_AGENT",
		"FAILED_TO_CONNECT_AGENT_TO_REPLICATION_SERVER",
		"FAILED_TO_START_DATA_TRANSFER",
	}
}

type DataReplicationInitiationStepName string

// Enum values for DataReplicationInitiationStepName
const (
	DataReplicationInitiationStepNameWait                            DataReplicationInitiationStepName = "WAIT"
	DataReplicationInitiationStepNameCreateSecurityGroup             DataReplicationInitiationStepName = "CREATE_SECURITY_GROUP"
	DataReplicationInitiationStepNameLaunchReplicationServer         DataReplicationInitiationStepName = "LAUNCH_REPLICATION_SERVER"
	DataReplicationInitiationStepNameBootReplicationServer           DataReplicationInitiationStepName = "BOOT_REPLICATION_SERVER"
	DataReplicationInitiationStepNameAuthenticateWithService         DataReplicationInitiationStepName = "AUTHENTICATE_WITH_SERVICE"
	DataReplicationInitiationStepNameDownloadReplicationSoftware     DataReplicationInitiationStepName = "DOWNLOAD_REPLICATION_SOFTWARE"
	DataReplicationInitiationStepNameCreateStagingDisks              DataReplicationInitiationStepName = "CREATE_STAGING_DISKS"
	DataReplicationInitiationStepNameAttachStagingDisks              DataReplicationInitiationStepName = "ATTACH_STAGING_DISKS"
	DataReplicationInitiationStepNamePairReplicationServerWithAgent  DataReplicationInitiationStepName = "PAIR_REPLICATION_SERVER_WITH_AGENT"
	DataReplicationInitiationStepNameConnectAgentToReplicationServer DataReplicationInitiationStepName = "CONNECT_AGENT_TO_REPLICATION_SERVER"
	DataReplicationInitiationStepNameStartDataTransfer               DataReplicationInitiationStepName = "START_DATA_TRANSFER"
)

// Values returns all known values for DataReplicationInitiationStepName. Note
// that this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (DataReplicationInitiationStepName) Values() []DataReplicationInitiationStepName {
	return []DataReplicationInitiationStepName{
		"WAIT",
		"CREATE_SECURITY_GROUP",
		"LAUNCH_REPLICATION_SERVER",
		"BOOT_REPLICATION_SERVER",
		"AUTHENTICATE_WITH_SERVICE",
		"DOWNLOAD_REPLICATION_SOFTWARE",
		"CREATE_STAGING_DISKS",
		"ATTACH_STAGING_DISKS",
		"PAIR_REPLICATION_SERVER_WITH_AGENT",
		"CONNECT_AGENT_TO_REPLICATION_SERVER",
		"START_DATA_TRANSFER",
	}
}

type DataReplicationInitiationStepStatus string

// Enum values for DataReplicationInitiationStepStatus
const (
	DataReplicationInitiationStepStatusNotStarted DataReplicationInitiationStepStatus = "NOT_STARTED"
	DataReplicationInitiationStepStatusInProgress DataReplicationInitiationStepStatus = "IN_PROGRESS"
	DataReplicationInitiationStepStatusSucceeded  DataReplicationInitiationStepStatus = "SUCCEEDED"
	DataReplicationInitiationStepStatusFailed     DataReplicationInitiationStepStatus = "FAILED"
	DataReplicationInitiationStepStatusSkipped    DataReplicationInitiationStepStatus = "SKIPPED"
)

// Values returns all known values for DataReplicationInitiationStepStatus. Note
// that this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (DataReplicationInitiationStepStatus) Values() []DataReplicationInitiationStepStatus {
	return []DataReplicationInitiationStepStatus{
		"NOT_STARTED",
		"IN_PROGRESS",
		"SUCCEEDED",
		"FAILED",
		"SKIPPED",
	}
}

type DataReplicationState string

// Enum values for DataReplicationState
const (
	DataReplicationStateStopped          DataReplicationState = "STOPPED"
	DataReplicationStateInitiating       DataReplicationState = "INITIATING"
	DataReplicationStateInitialSync      DataReplicationState = "INITIAL_SYNC"
	DataReplicationStateBacklog          DataReplicationState = "BACKLOG"
	DataReplicationStateCreatingSnapshot DataReplicationState = "CREATING_SNAPSHOT"
	DataReplicationStateContinuous       DataReplicationState = "CONTINUOUS"
	DataReplicationStatePaused           DataReplicationState = "PAUSED"
	DataReplicationStateRescan           DataReplicationState = "RESCAN"
	DataReplicationStateStalled          DataReplicationState = "STALLED"
	DataReplicationStateDisconnected     DataReplicationState = "DISCONNECTED"
)

// Values returns all known values for DataReplicationState. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (DataReplicationState) Values() []DataReplicationState {
	return []DataReplicationState{
		"STOPPED",
		"INITIATING",
		"INITIAL_SYNC",
		"BACKLOG",
		"CREATING_SNAPSHOT",
		"CONTINUOUS",
		"PAUSED",
		"RESCAN",
		"STALLED",
		"DISCONNECTED",
	}
}

type EC2InstanceState string

// Enum values for EC2InstanceState
const (
	EC2InstanceStatePending      EC2InstanceState = "PENDING"
	EC2InstanceStateRunning      EC2InstanceState = "RUNNING"
	EC2InstanceStateStopping     EC2InstanceState = "STOPPING"
	EC2InstanceStateStopped      EC2InstanceState = "STOPPED"
	EC2InstanceStateShuttingDown EC2InstanceState = "SHUTTING-DOWN"
	EC2InstanceStateTerminated   EC2InstanceState = "TERMINATED"
	EC2InstanceStateNotFound     EC2InstanceState = "NOT_FOUND"
)

// Values returns all known values for EC2InstanceState. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (EC2InstanceState) Values() []EC2InstanceState {
	return []EC2InstanceState{
		"PENDING",
		"RUNNING",
		"STOPPING",
		"STOPPED",
		"SHUTTING-DOWN",
		"TERMINATED",
		"NOT_FOUND",
	}
}

type ExtensionStatus string

// Enum values for ExtensionStatus
const (
	ExtensionStatusExtended       ExtensionStatus = "EXTENDED"
	ExtensionStatusExtensionError ExtensionStatus = "EXTENSION_ERROR"
	ExtensionStatusNotExtended    ExtensionStatus = "NOT_EXTENDED"
)

// Values returns all known values for ExtensionStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ExtensionStatus) Values() []ExtensionStatus {
	return []ExtensionStatus{
		"EXTENDED",
		"EXTENSION_ERROR",
		"NOT_EXTENDED",
	}
}

type FailbackLaunchType string

// Enum values for FailbackLaunchType
const (
	FailbackLaunchTypeRecovery FailbackLaunchType = "RECOVERY"
	FailbackLaunchTypeDrill    FailbackLaunchType = "DRILL"
)

// Values returns all known values for FailbackLaunchType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (FailbackLaunchType) Values() []FailbackLaunchType {
	return []FailbackLaunchType{
		"RECOVERY",
		"DRILL",
	}
}

type FailbackReplicationError string

// Enum values for FailbackReplicationError
const (
	FailbackReplicationErrorAgentNotSeen                                          FailbackReplicationError = "AGENT_NOT_SEEN"
	FailbackReplicationErrorFailbackClientNotSeen                                 FailbackReplicationError = "FAILBACK_CLIENT_NOT_SEEN"
	FailbackReplicationErrorNotConverging                                         FailbackReplicationError = "NOT_CONVERGING"
	FailbackReplicationErrorUnstableNetwork                                       FailbackReplicationError = "UNSTABLE_NETWORK"
	FailbackReplicationErrorFailedToEstablishRecoveryInstanceCommunication        FailbackReplicationError = "FAILED_TO_ESTABLISH_RECOVERY_INSTANCE_COMMUNICATION"
	FailbackReplicationErrorFailedToDownloadReplicationSoftwareToFailbackClient   FailbackReplicationError = "FAILED_TO_DOWNLOAD_REPLICATION_SOFTWARE_TO_FAILBACK_CLIENT"
	FailbackReplicationErrorFailedToConfigureReplicationSoftware                  FailbackReplicationError = "FAILED_TO_CONFIGURE_REPLICATION_SOFTWARE"
	FailbackReplicationErrorFailedToPairAgentWithReplicationSoftware              FailbackReplicationError = "FAILED_TO_PAIR_AGENT_WITH_REPLICATION_SOFTWARE"
	FailbackReplicationErrorFailedToEstablishAgentReplicatorSoftwareCommunication FailbackReplicationError = "FAILED_TO_ESTABLISH_AGENT_REPLICATOR_SOFTWARE_COMMUNICATION"
	FailbackReplicationErrorFailedGettingReplicationState                         FailbackReplicationError = "FAILED_GETTING_REPLICATION_STATE"
	FailbackReplicationErrorSnapshotsFailure                                      FailbackReplicationError = "SNAPSHOTS_FAILURE"
	FailbackReplicationErrorFailedToCreateSecurityGroup                           FailbackReplicationError = "FAILED_TO_CREATE_SECURITY_GROUP"
	FailbackReplicationErrorFailedToLaunchReplicationServer                       FailbackReplicationError = "FAILED_TO_LAUNCH_REPLICATION_SERVER"
	FailbackReplicationErrorFailedToBootReplicationServer                         FailbackReplicationError = "FAILED_TO_BOOT_REPLICATION_SERVER"
	FailbackReplicationErrorFailedToAuthenticateWithService                       FailbackReplicationError = "FAILED_TO_AUTHENTICATE_WITH_SERVICE"
	FailbackReplicationErrorFailedToDownloadReplicationSoftware                   FailbackReplicationError = "FAILED_TO_DOWNLOAD_REPLICATION_SOFTWARE"
	FailbackReplicationErrorFailedToCreateStagingDisks                            FailbackReplicationError = "FAILED_TO_CREATE_STAGING_DISKS"
	FailbackReplicationErrorFailedToAttachStagingDisks                            FailbackReplicationError = "FAILED_TO_ATTACH_STAGING_DISKS"
	FailbackReplicationErrorFailedToPairReplicationServerWithAgent                FailbackReplicationError = "FAILED_TO_PAIR_REPLICATION_SERVER_WITH_AGENT"
	FailbackReplicationErrorFailedToConnectAgentToReplicationServer               FailbackReplicationError = "FAILED_TO_CONNECT_AGENT_TO_REPLICATION_SERVER"
	FailbackReplicationErrorFailedToStartDataTransfer                             FailbackReplicationError = "FAILED_TO_START_DATA_TRANSFER"
)

// Values returns all known values for FailbackReplicationError. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (FailbackReplicationError) Values() []FailbackReplicationError {
	return []FailbackReplicationError{
		"AGENT_NOT_SEEN",
		"FAILBACK_CLIENT_NOT_SEEN",
		"NOT_CONVERGING",
		"UNSTABLE_NETWORK",
		"FAILED_TO_ESTABLISH_RECOVERY_INSTANCE_COMMUNICATION",
		"FAILED_TO_DOWNLOAD_REPLICATION_SOFTWARE_TO_FAILBACK_CLIENT",
		"FAILED_TO_CONFIGURE_REPLICATION_SOFTWARE",
		"FAILED_TO_PAIR_AGENT_WITH_REPLICATION_SOFTWARE",
		"FAILED_TO_ESTABLISH_AGENT_REPLICATOR_SOFTWARE_COMMUNICATION",
		"FAILED_GETTING_REPLICATION_STATE",
		"SNAPSHOTS_FAILURE",
		"FAILED_TO_CREATE_SECURITY_GROUP",
		"FAILED_TO_LAUNCH_REPLICATION_SERVER",
		"FAILED_TO_BOOT_REPLICATION_SERVER",
		"FAILED_TO_AUTHENTICATE_WITH_SERVICE",
		"FAILED_TO_DOWNLOAD_REPLICATION_SOFTWARE",
		"FAILED_TO_CREATE_STAGING_DISKS",
		"FAILED_TO_ATTACH_STAGING_DISKS",
		"FAILED_TO_PAIR_REPLICATION_SERVER_WITH_AGENT",
		"FAILED_TO_CONNECT_AGENT_TO_REPLICATION_SERVER",
		"FAILED_TO_START_DATA_TRANSFER",
	}
}

type FailbackState string

// Enum values for FailbackState
const (
	FailbackStateFailbackNotStarted              FailbackState = "FAILBACK_NOT_STARTED"
	FailbackStateFailbackInProgress              FailbackState = "FAILBACK_IN_PROGRESS"
	FailbackStateFailbackReadyForLaunch          FailbackState = "FAILBACK_READY_FOR_LAUNCH"
	FailbackStateFailbackCompleted               FailbackState = "FAILBACK_COMPLETED"
	FailbackStateFailbackError                   FailbackState = "FAILBACK_ERROR"
	FailbackStateFailbackNotReadyForLaunch       FailbackState = "FAILBACK_NOT_READY_FOR_LAUNCH"
	FailbackStateFailbackLaunchStateNotAvailable FailbackState = "FAILBACK_LAUNCH_STATE_NOT_AVAILABLE"
)

// Values returns all known values for FailbackState. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (FailbackState) Values() []FailbackState {
	return []FailbackState{
		"FAILBACK_NOT_STARTED",
		"FAILBACK_IN_PROGRESS",
		"FAILBACK_READY_FOR_LAUNCH",
		"FAILBACK_COMPLETED",
		"FAILBACK_ERROR",
		"FAILBACK_NOT_READY_FOR_LAUNCH",
		"FAILBACK_LAUNCH_STATE_NOT_AVAILABLE",
	}
}

type InitiatedBy string

// Enum values for InitiatedBy
const (
	InitiatedByStartRecovery              InitiatedBy = "START_RECOVERY"
	InitiatedByStartDrill                 InitiatedBy = "START_DRILL"
	InitiatedByFailback                   InitiatedBy = "FAILBACK"
	InitiatedByDiagnostic                 InitiatedBy = "DIAGNOSTIC"
	InitiatedByTerminateRecoveryInstances InitiatedBy = "TERMINATE_RECOVERY_INSTANCES"
	InitiatedByTargetAccount              InitiatedBy = "TARGET_ACCOUNT"
)

// Values returns all known values for InitiatedBy. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (InitiatedBy) Values() []InitiatedBy {
	return []InitiatedBy{
		"START_RECOVERY",
		"START_DRILL",
		"FAILBACK",
		"DIAGNOSTIC",
		"TERMINATE_RECOVERY_INSTANCES",
		"TARGET_ACCOUNT",
	}
}

type JobLogEvent string

// Enum values for JobLogEvent
const (
	JobLogEventJobStart                    JobLogEvent = "JOB_START"
	JobLogEventServerSkipped               JobLogEvent = "SERVER_SKIPPED"
	JobLogEventCleanupStart                JobLogEvent = "CLEANUP_START"
	JobLogEventCleanupEnd                  JobLogEvent = "CLEANUP_END"
	JobLogEventCleanupFail                 JobLogEvent = "CLEANUP_FAIL"
	JobLogEventSnapshotStart               JobLogEvent = "SNAPSHOT_START"
	JobLogEventSnapshotEnd                 JobLogEvent = "SNAPSHOT_END"
	JobLogEventSnapshotFail                JobLogEvent = "SNAPSHOT_FAIL"
	JobLogEventUsingPreviousSnapshot       JobLogEvent = "USING_PREVIOUS_SNAPSHOT"
	JobLogEventUsingPreviousSnapshotFailed JobLogEvent = "USING_PREVIOUS_SNAPSHOT_FAILED"
	JobLogEventConversionStart             JobLogEvent = "CONVERSION_START"
	JobLogEventConversionEnd               JobLogEvent = "CONVERSION_END"
	JobLogEventConversionFail              JobLogEvent = "CONVERSION_FAIL"
	JobLogEventLaunchStart                 JobLogEvent = "LAUNCH_START"
	JobLogEventLaunchFailed                JobLogEvent = "LAUNCH_FAILED"
	JobLogEventJobCancel                   JobLogEvent = "JOB_CANCEL"
	JobLogEventJobEnd                      JobLogEvent = "JOB_END"
)

// Values returns all known values for JobLogEvent. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (JobLogEvent) Values() []JobLogEvent {
	return []JobLogEvent{
		"JOB_START",
		"SERVER_SKIPPED",
		"CLEANUP_START",
		"CLEANUP_END",
		"CLEANUP_FAIL",
		"SNAPSHOT_START",
		"SNAPSHOT_END",
		"SNAPSHOT_FAIL",
		"USING_PREVIOUS_SNAPSHOT",
		"USING_PREVIOUS_SNAPSHOT_FAILED",
		"CONVERSION_START",
		"CONVERSION_END",
		"CONVERSION_FAIL",
		"LAUNCH_START",
		"LAUNCH_FAILED",
		"JOB_CANCEL",
		"JOB_END",
	}
}

type JobStatus string

// Enum values for JobStatus
const (
	JobStatusPending   JobStatus = "PENDING"
	JobStatusStarted   JobStatus = "STARTED"
	JobStatusCompleted JobStatus = "COMPLETED"
)

// Values returns all known values for JobStatus. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (JobStatus) Values() []JobStatus {
	return []JobStatus{
		"PENDING",
		"STARTED",
		"COMPLETED",
	}
}

type JobType string

// Enum values for JobType
const (
	JobTypeLaunch                  JobType = "LAUNCH"
	JobTypeTerminate               JobType = "TERMINATE"
	JobTypeCreateConvertedSnapshot JobType = "CREATE_CONVERTED_SNAPSHOT"
)

// Values returns all known values for JobType. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (JobType) Values() []JobType {
	return []JobType{
		"LAUNCH",
		"TERMINATE",
		"CREATE_CONVERTED_SNAPSHOT",
	}
}

type LastLaunchResult string

// Enum values for LastLaunchResult
const (
	LastLaunchResultNotStarted LastLaunchResult = "NOT_STARTED"
	LastLaunchResultPending    LastLaunchResult = "PENDING"
	LastLaunchResultSucceeded  LastLaunchResult = "SUCCEEDED"
	LastLaunchResultFailed     LastLaunchResult = "FAILED"
)

// Values returns all known values for LastLaunchResult. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (LastLaunchResult) Values() []LastLaunchResult {
	return []LastLaunchResult{
		"NOT_STARTED",
		"PENDING",
		"SUCCEEDED",
		"FAILED",
	}
}

type LastLaunchType string

// Enum values for LastLaunchType
const (
	LastLaunchTypeRecovery LastLaunchType = "RECOVERY"
	LastLaunchTypeDrill    LastLaunchType = "DRILL"
)

// Values returns all known values for LastLaunchType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (LastLaunchType) Values() []LastLaunchType {
	return []LastLaunchType{
		"RECOVERY",
		"DRILL",
	}
}

type LaunchDisposition string

// Enum values for LaunchDisposition
const (
	LaunchDispositionStopped LaunchDisposition = "STOPPED"
	LaunchDispositionStarted LaunchDisposition = "STARTED"
)

// Values returns all known values for LaunchDisposition. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (LaunchDisposition) Values() []LaunchDisposition {
	return []LaunchDisposition{
		"STOPPED",
		"STARTED",
	}
}

type LaunchStatus string

// Enum values for LaunchStatus
const (
	LaunchStatusPending    LaunchStatus = "PENDING"
	LaunchStatusInProgress LaunchStatus = "IN_PROGRESS"
	LaunchStatusLaunched   LaunchStatus = "LAUNCHED"
	LaunchStatusFailed     LaunchStatus = "FAILED"
	LaunchStatusTerminated LaunchStatus = "TERMINATED"
)

// Values returns all known values for LaunchStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (LaunchStatus) Values() []LaunchStatus {
	return []LaunchStatus{
		"PENDING",
		"IN_PROGRESS",
		"LAUNCHED",
		"FAILED",
		"TERMINATED",
	}
}

type OriginEnvironment string

// Enum values for OriginEnvironment
const (
	OriginEnvironmentOnPremises OriginEnvironment = "ON_PREMISES"
	OriginEnvironmentAws        OriginEnvironment = "AWS"
)

// Values returns all known values for OriginEnvironment. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (OriginEnvironment) Values() []OriginEnvironment {
	return []OriginEnvironment{
		"ON_PREMISES",
		"AWS",
	}
}

type PITPolicyRuleUnits string

// Enum values for PITPolicyRuleUnits
const (
	PITPolicyRuleUnitsMinute PITPolicyRuleUnits = "MINUTE"
	PITPolicyRuleUnitsHour   PITPolicyRuleUnits = "HOUR"
	PITPolicyRuleUnitsDay    PITPolicyRuleUnits = "DAY"
)

// Values returns all known values for PITPolicyRuleUnits. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (PITPolicyRuleUnits) Values() []PITPolicyRuleUnits {
	return []PITPolicyRuleUnits{
		"MINUTE",
		"HOUR",
		"DAY",
	}
}

type RecoveryInstanceDataReplicationInitiationStepName string

// Enum values for RecoveryInstanceDataReplicationInitiationStepName
const (
	RecoveryInstanceDataReplicationInitiationStepNameLinkFailbackClientWithRecoveryInstance        RecoveryInstanceDataReplicationInitiationStepName = "LINK_FAILBACK_CLIENT_WITH_RECOVERY_INSTANCE"
	RecoveryInstanceDataReplicationInitiationStepNameCompleteVolumeMapping                         RecoveryInstanceDataReplicationInitiationStepName = "COMPLETE_VOLUME_MAPPING"
	RecoveryInstanceDataReplicationInitiationStepNameEstablishRecoveryInstanceCommunication        RecoveryInstanceDataReplicationInitiationStepName = "ESTABLISH_RECOVERY_INSTANCE_COMMUNICATION"
	RecoveryInstanceDataReplicationInitiationStepNameDownloadReplicationSoftwareToFailbackClient   RecoveryInstanceDataReplicationInitiationStepName = "DOWNLOAD_REPLICATION_SOFTWARE_TO_FAILBACK_CLIENT"
	RecoveryInstanceDataReplicationInitiationStepNameConfigureReplicationSoftware                  RecoveryInstanceDataReplicationInitiationStepName = "CONFIGURE_REPLICATION_SOFTWARE"
	RecoveryInstanceDataReplicationInitiationStepNamePairAgentWithReplicationSoftware              RecoveryInstanceDataReplicationInitiationStepName = "PAIR_AGENT_WITH_REPLICATION_SOFTWARE"
	RecoveryInstanceDataReplicationInitiationStepNameEstablishAgentReplicatorSoftwareCommunication RecoveryInstanceDataReplicationInitiationStepName = "ESTABLISH_AGENT_REPLICATOR_SOFTWARE_COMMUNICATION"
	RecoveryInstanceDataReplicationInitiationStepNameWait                                          RecoveryInstanceDataReplicationInitiationStepName = "WAIT"
	RecoveryInstanceDataReplicationInitiationStepNameCreateSecurityGroup                           RecoveryInstanceDataReplicationInitiationStepName = "CREATE_SECURITY_GROUP"
	RecoveryInstanceDataReplicationInitiationStepNameLaunchReplicationServer                       RecoveryInstanceDataReplicationInitiationStepName = "LAUNCH_REPLICATION_SERVER"
	RecoveryInstanceDataReplicationInitiationStepNameBootReplicationServer                         RecoveryInstanceDataReplicationInitiationStepName = "BOOT_REPLICATION_SERVER"
	RecoveryInstanceDataReplicationInitiationStepNameAuthenticateWithService                       RecoveryInstanceDataReplicationInitiationStepName = "AUTHENTICATE_WITH_SERVICE"
	RecoveryInstanceDataReplicationInitiationStepNameDownloadReplicationSoftware                   RecoveryInstanceDataReplicationInitiationStepName = "DOWNLOAD_REPLICATION_SOFTWARE"
	RecoveryInstanceDataReplicationInitiationStepNameCreateStagingDisks                            RecoveryInstanceDataReplicationInitiationStepName = "CREATE_STAGING_DISKS"
	RecoveryInstanceDataReplicationInitiationStepNameAttachStagingDisks                            RecoveryInstanceDataReplicationInitiationStepName = "ATTACH_STAGING_DISKS"
	RecoveryInstanceDataReplicationInitiationStepNamePairReplicationServerWithAgent                RecoveryInstanceDataReplicationInitiationStepName = "PAIR_REPLICATION_SERVER_WITH_AGENT"
	RecoveryInstanceDataReplicationInitiationStepNameConnectAgentToReplicationServer               RecoveryInstanceDataReplicationInitiationStepName = "CONNECT_AGENT_TO_REPLICATION_SERVER"
	RecoveryInstanceDataReplicationInitiationStepNameStartDataTransfer                             RecoveryInstanceDataReplicationInitiationStepName = "START_DATA_TRANSFER"
)

// Values returns all known values for
// RecoveryInstanceDataReplicationInitiationStepName. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (RecoveryInstanceDataReplicationInitiationStepName) Values() []RecoveryInstanceDataReplicationInitiationStepName {
	return []RecoveryInstanceDataReplicationInitiationStepName{
		"LINK_FAILBACK_CLIENT_WITH_RECOVERY_INSTANCE",
		"COMPLETE_VOLUME_MAPPING",
		"ESTABLISH_RECOVERY_INSTANCE_COMMUNICATION",
		"DOWNLOAD_REPLICATION_SOFTWARE_TO_FAILBACK_CLIENT",
		"CONFIGURE_REPLICATION_SOFTWARE",
		"PAIR_AGENT_WITH_REPLICATION_SOFTWARE",
		"ESTABLISH_AGENT_REPLICATOR_SOFTWARE_COMMUNICATION",
		"WAIT",
		"CREATE_SECURITY_GROUP",
		"LAUNCH_REPLICATION_SERVER",
		"BOOT_REPLICATION_SERVER",
		"AUTHENTICATE_WITH_SERVICE",
		"DOWNLOAD_REPLICATION_SOFTWARE",
		"CREATE_STAGING_DISKS",
		"ATTACH_STAGING_DISKS",
		"PAIR_REPLICATION_SERVER_WITH_AGENT",
		"CONNECT_AGENT_TO_REPLICATION_SERVER",
		"START_DATA_TRANSFER",
	}
}

type RecoveryInstanceDataReplicationInitiationStepStatus string

// Enum values for RecoveryInstanceDataReplicationInitiationStepStatus
const (
	RecoveryInstanceDataReplicationInitiationStepStatusNotStarted RecoveryInstanceDataReplicationInitiationStepStatus = "NOT_STARTED"
	RecoveryInstanceDataReplicationInitiationStepStatusInProgress RecoveryInstanceDataReplicationInitiationStepStatus = "IN_PROGRESS"
	RecoveryInstanceDataReplicationInitiationStepStatusSucceeded  RecoveryInstanceDataReplicationInitiationStepStatus = "SUCCEEDED"
	RecoveryInstanceDataReplicationInitiationStepStatusFailed     RecoveryInstanceDataReplicationInitiationStepStatus = "FAILED"
	RecoveryInstanceDataReplicationInitiationStepStatusSkipped    RecoveryInstanceDataReplicationInitiationStepStatus = "SKIPPED"
)

// Values returns all known values for
// RecoveryInstanceDataReplicationInitiationStepStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (RecoveryInstanceDataReplicationInitiationStepStatus) Values() []RecoveryInstanceDataReplicationInitiationStepStatus {
	return []RecoveryInstanceDataReplicationInitiationStepStatus{
		"NOT_STARTED",
		"IN_PROGRESS",
		"SUCCEEDED",
		"FAILED",
		"SKIPPED",
	}
}

type RecoveryInstanceDataReplicationState string

// Enum values for RecoveryInstanceDataReplicationState
const (
	RecoveryInstanceDataReplicationStateStopped                      RecoveryInstanceDataReplicationState = "STOPPED"
	RecoveryInstanceDataReplicationStateInitiating                   RecoveryInstanceDataReplicationState = "INITIATING"
	RecoveryInstanceDataReplicationStateInitialSync                  RecoveryInstanceDataReplicationState = "INITIAL_SYNC"
	RecoveryInstanceDataReplicationStateBacklog                      RecoveryInstanceDataReplicationState = "BACKLOG"
	RecoveryInstanceDataReplicationStateCreatingSnapshot             RecoveryInstanceDataReplicationState = "CREATING_SNAPSHOT"
	RecoveryInstanceDataReplicationStateContinuous                   RecoveryInstanceDataReplicationState = "CONTINUOUS"
	RecoveryInstanceDataReplicationStatePaused                       RecoveryInstanceDataReplicationState = "PAUSED"
	RecoveryInstanceDataReplicationStateRescan                       RecoveryInstanceDataReplicationState = "RESCAN"
	RecoveryInstanceDataReplicationStateStalled                      RecoveryInstanceDataReplicationState = "STALLED"
	RecoveryInstanceDataReplicationStateDisconnected                 RecoveryInstanceDataReplicationState = "DISCONNECTED"
	RecoveryInstanceDataReplicationStateReplicationStateNotAvailable RecoveryInstanceDataReplicationState = "REPLICATION_STATE_NOT_AVAILABLE"
	RecoveryInstanceDataReplicationStateNotStarted                   RecoveryInstanceDataReplicationState = "NOT_STARTED"
)

// Values returns all known values for RecoveryInstanceDataReplicationState. Note
// that this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (RecoveryInstanceDataReplicationState) Values() []RecoveryInstanceDataReplicationState {
	return []RecoveryInstanceDataReplicationState{
		"STOPPED",
		"INITIATING",
		"INITIAL_SYNC",
		"BACKLOG",
		"CREATING_SNAPSHOT",
		"CONTINUOUS",
		"PAUSED",
		"RESCAN",
		"STALLED",
		"DISCONNECTED",
		"REPLICATION_STATE_NOT_AVAILABLE",
		"NOT_STARTED",
	}
}

type RecoverySnapshotsOrder string

// Enum values for RecoverySnapshotsOrder
const (
	RecoverySnapshotsOrderAsc  RecoverySnapshotsOrder = "ASC"
	RecoverySnapshotsOrderDesc RecoverySnapshotsOrder = "DESC"
)

// Values returns all known values for RecoverySnapshotsOrder. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (RecoverySnapshotsOrder) Values() []RecoverySnapshotsOrder {
	return []RecoverySnapshotsOrder{
		"ASC",
		"DESC",
	}
}

type ReplicationConfigurationDataPlaneRouting string

// Enum values for ReplicationConfigurationDataPlaneRouting
const (
	ReplicationConfigurationDataPlaneRoutingPrivateIp ReplicationConfigurationDataPlaneRouting = "PRIVATE_IP"
	ReplicationConfigurationDataPlaneRoutingPublicIp  ReplicationConfigurationDataPlaneRouting = "PUBLIC_IP"
)

// Values returns all known values for ReplicationConfigurationDataPlaneRouting.
// Note that this can be expanded in the future, and so it is only as up to date as
// the client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (ReplicationConfigurationDataPlaneRouting) Values() []ReplicationConfigurationDataPlaneRouting {
	return []ReplicationConfigurationDataPlaneRouting{
		"PRIVATE_IP",
		"PUBLIC_IP",
	}
}

type ReplicationConfigurationDefaultLargeStagingDiskType string

// Enum values for ReplicationConfigurationDefaultLargeStagingDiskType
const (
	ReplicationConfigurationDefaultLargeStagingDiskTypeGp2  ReplicationConfigurationDefaultLargeStagingDiskType = "GP2"
	ReplicationConfigurationDefaultLargeStagingDiskTypeGp3  ReplicationConfigurationDefaultLargeStagingDiskType = "GP3"
	ReplicationConfigurationDefaultLargeStagingDiskTypeSt1  ReplicationConfigurationDefaultLargeStagingDiskType = "ST1"
	ReplicationConfigurationDefaultLargeStagingDiskTypeAuto ReplicationConfigurationDefaultLargeStagingDiskType = "AUTO"
)

// Values returns all known values for
// ReplicationConfigurationDefaultLargeStagingDiskType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ReplicationConfigurationDefaultLargeStagingDiskType) Values() []ReplicationConfigurationDefaultLargeStagingDiskType {
	return []ReplicationConfigurationDefaultLargeStagingDiskType{
		"GP2",
		"GP3",
		"ST1",
		"AUTO",
	}
}

type ReplicationConfigurationEbsEncryption string

// Enum values for ReplicationConfigurationEbsEncryption
const (
	ReplicationConfigurationEbsEncryptionDefault ReplicationConfigurationEbsEncryption = "DEFAULT"
	ReplicationConfigurationEbsEncryptionCustom  ReplicationConfigurationEbsEncryption = "CUSTOM"
)

// Values returns all known values for ReplicationConfigurationEbsEncryption. Note
// that this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (ReplicationConfigurationEbsEncryption) Values() []ReplicationConfigurationEbsEncryption {
	return []ReplicationConfigurationEbsEncryption{
		"DEFAULT",
		"CUSTOM",
	}
}

type ReplicationConfigurationReplicatedDiskStagingDiskType string

// Enum values for ReplicationConfigurationReplicatedDiskStagingDiskType
const (
	ReplicationConfigurationReplicatedDiskStagingDiskTypeAuto     ReplicationConfigurationReplicatedDiskStagingDiskType = "AUTO"
	ReplicationConfigurationReplicatedDiskStagingDiskTypeGp2      ReplicationConfigurationReplicatedDiskStagingDiskType = "GP2"
	ReplicationConfigurationReplicatedDiskStagingDiskTypeGp3      ReplicationConfigurationReplicatedDiskStagingDiskType = "GP3"
	ReplicationConfigurationReplicatedDiskStagingDiskTypeIo1      ReplicationConfigurationReplicatedDiskStagingDiskType = "IO1"
	ReplicationConfigurationReplicatedDiskStagingDiskTypeSc1      ReplicationConfigurationReplicatedDiskStagingDiskType = "SC1"
	ReplicationConfigurationReplicatedDiskStagingDiskTypeSt1      ReplicationConfigurationReplicatedDiskStagingDiskType = "ST1"
	ReplicationConfigurationReplicatedDiskStagingDiskTypeStandard ReplicationConfigurationReplicatedDiskStagingDiskType = "STANDARD"
)

// Values returns all known values for
// ReplicationConfigurationReplicatedDiskStagingDiskType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ReplicationConfigurationReplicatedDiskStagingDiskType) Values() []ReplicationConfigurationReplicatedDiskStagingDiskType {
	return []ReplicationConfigurationReplicatedDiskStagingDiskType{
		"AUTO",
		"GP2",
		"GP3",
		"IO1",
		"SC1",
		"ST1",
		"STANDARD",
	}
}

type ReplicationDirection string

// Enum values for ReplicationDirection
const (
	ReplicationDirectionFailover ReplicationDirection = "FAILOVER"
	ReplicationDirectionFailback ReplicationDirection = "FAILBACK"
)

// Values returns all known values for ReplicationDirection. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ReplicationDirection) Values() []ReplicationDirection {
	return []ReplicationDirection{
		"FAILOVER",
		"FAILBACK",
	}
}

type TargetInstanceTypeRightSizingMethod string

// Enum values for TargetInstanceTypeRightSizingMethod
const (
	TargetInstanceTypeRightSizingMethodNone  TargetInstanceTypeRightSizingMethod = "NONE"
	TargetInstanceTypeRightSizingMethodBasic TargetInstanceTypeRightSizingMethod = "BASIC"
)

// Values returns all known values for TargetInstanceTypeRightSizingMethod. Note
// that this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (TargetInstanceTypeRightSizingMethod) Values() []TargetInstanceTypeRightSizingMethod {
	return []TargetInstanceTypeRightSizingMethod{
		"NONE",
		"BASIC",
	}
}

type ValidationExceptionReason string

// Enum values for ValidationExceptionReason
const (
	ValidationExceptionReasonUnknownOperation      ValidationExceptionReason = "unknownOperation"
	ValidationExceptionReasonCannotParse           ValidationExceptionReason = "cannotParse"
	ValidationExceptionReasonFieldValidationFailed ValidationExceptionReason = "fieldValidationFailed"
	ValidationExceptionReasonOther                 ValidationExceptionReason = "other"
)

// Values returns all known values for ValidationExceptionReason. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (ValidationExceptionReason) Values() []ValidationExceptionReason {
	return []ValidationExceptionReason{
		"unknownOperation",
		"cannotParse",
		"fieldValidationFailed",
		"other",
	}
}
