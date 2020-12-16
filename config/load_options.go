package config

import (
	"context"
	"io"

	"github.com/aws/aws-sdk-go-v2/feature/ec2/imds"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/credentials/ec2rolecreds"
	"github.com/aws/aws-sdk-go-v2/credentials/endpointcreds"
	"github.com/aws/aws-sdk-go-v2/credentials/processcreds"
	"github.com/aws/aws-sdk-go-v2/credentials/stscreds"
	"github.com/awslabs/smithy-go/logging"
	"github.com/awslabs/smithy-go/middleware"
)

// LoadOptionsFunc is a type alias for LoadOptions functional option
type LoadOptionsFunc func(*LoadOptions) error

// LoadOptions are discrete set of options that are valid for loading the configuration
type LoadOptions struct {

	// Region is the region to send requests to.
	Region string

	// Credentials object to use when signing requests.
	Credentials aws.CredentialsProvider

	// HTTPClient the SDK's API clients will use to invoke HTTP requests.
	HTTPClient HTTPClient

	// EndpointResolver that can be used to provide or override an endpoint for the given
	// service and region Please see the `aws.EndpointResolver` documentation on usage.
	EndpointResolver aws.EndpointResolver

	// Retryer guides how HTTP requests should be retried in case of
	// recoverable failures.
	Retryer aws.Retryer

	// APIOptions provides the set of middleware mutations modify how the API
	// client requests will be handled. This is useful for adding additional
	// tracing data to a request, or changing behavior of the SDK's client.
	APIOptions []func(*middleware.Stack) error

	// Logger writer interface to write logging messages to.
	Logger logging.Logger

	// ClientLogMode is used to configure the events that will be sent to the configured logger.
	// This can be used to configure the logging of signing, retries, request, and responses
	// of the SDK clients.
	//
	// See the ClientLogMode type documentation for the complete set of logging modes and available
	// configuration.
	ClientLogMode *aws.ClientLogMode

	// SharedConfigProfile is the profile to be used when loading the SharedConfig
	SharedConfigProfile string

	// SharedConfigFiles is the slice of custom shared config files to use when loading the SharedConfig
	SharedConfigFiles []string

	// CustomCABundle is CA bundle PEM bytes reader
	CustomCABundle io.Reader

	// DefaultRegion is the fall back region, used if a region was not resolved from other sources
	DefaultRegion string

	// UseEC2IMDSRegion indicates if SDK should retrieve the region
	// from the EC2 Metadata service
	UseEC2IMDSRegion *UseEC2IMDSRegion

	// ProcessCredentialOptions is a function for setting
	// the processcreds.Options
	ProcessCredentialOptions func(*processcreds.Options)

	// EC2RoleCredentialOptions is a function for setting
	// the ec2rolecreds.Options
	EC2RoleCredentialOptions func(*ec2rolecreds.Options)

	// EndpointCredentialOptions is a function for setting
	// the endpointcreds.Options
	EndpointCredentialOptions func(*endpointcreds.Options)

	// WebIdentityRoleCredentialOptions is a function for setting
	// the stscreds.WebIdentityRoleOptions
	WebIdentityRoleCredentialOptions func(*stscreds.WebIdentityRoleOptions)

	// AssumeRoleCredentialOptions is a function for setting the
	// stscreds.AssumeRoleOptions
	AssumeRoleCredentialOptions func(*stscreds.AssumeRoleOptions)

	// LogConfigurationWarnings when set to true, enables logging
	// configuration warnings
	LogConfigurationWarnings *bool

	// S3UseARNRegion specifies if the S3 service should allow ARNs to direct
	// the region, the client's requests are sent to.
	S3UseARNRegion *bool
}

// getRegion returns Region from config's LoadOptions
func (o LoadOptions) getRegion(ctx context.Context) (string, bool, error) {
	if len(o.Region) == 0 {
		return "", false, nil
	}

	return o.Region, true, nil
}

// WithRegion is a helper function to construct functional options
// that sets Region on config's LoadOptions. Setting the region to
// an empty string, will result in the region value being ignored.
// If multiple WithRegion calls are made, the last call overrides
// the previous call values.
func WithRegion(v string) LoadOptionsFunc {
	return func(o *LoadOptions) error {
		o.Region = v
		return nil
	}
}

// getDefaultRegion returns DefaultRegion from config's LoadOptions
func (o LoadOptions) getDefaultRegion(ctx context.Context) (string, bool, error) {
	if len(o.DefaultRegion) == 0 {
		return "", false, nil
	}

	return o.DefaultRegion, true, nil
}

// WithDefaultRegion is a helper function to construct functional options
// that sets a DefaultRegion on config's LoadOptions. Setting the default
// region to an empty string, will result in the default region value
// being ignored. If multiple WithDefaultRegion calls are made, the last
// call overrides the previous call values. Note that both WithRegion and
// WithEC2IMDSRegion call takes precedence over WithDefaultRegion call
// when resolving region.
func WithDefaultRegion(v string) LoadOptionsFunc {
	return func(o *LoadOptions) error {
		o.DefaultRegion = v
		return nil
	}
}

// getSharedConfigProfile returns SharedConfigProfile from config's LoadOptions
func (o LoadOptions) getSharedConfigProfile(ctx context.Context) (string, bool, error) {
	if len(o.SharedConfigProfile) == 0 {
		return "", false, nil
	}

	return o.SharedConfigProfile, true, nil
}

// WithSharedConfigProfile is a helper function to construct functional options
// that sets SharedConfigProfile on config's LoadOptions. Setting the shared
// config profile to an empty string, will result in the shared config profile
// value being ignored.
// If multiple WithSharedConfigProfile calls are made, the last call overrides
// the previous call values.
func WithSharedConfigProfile(v string) LoadOptionsFunc {
	return func(o *LoadOptions) error {
		o.SharedConfigProfile = v
		return nil
	}
}

// getSharedConfigFiles returns SharedConfigFiles set on config's LoadOptions
func (o LoadOptions) getSharedConfigFiles(ctx context.Context) ([]string, bool, error) {
	if o.SharedConfigFiles == nil {
		return nil, false, nil
	}

	return o.SharedConfigFiles, true, nil
}

// WithSharedConfigFiles is a helper function to construct functional options
// that sets slice of SharedConfigFiles on config's LoadOptions.
// Setting the shared config files to an nil string slice, will result in the
// shared config files value being ignored.
// If multiple WithSharedConfigFiles calls are made, the last call overrides
// the previous call values.
func WithSharedConfigFiles(v []string) LoadOptionsFunc {
	return func(o *LoadOptions) error {
		o.SharedConfigFiles = v
		return nil
	}
}

// getCustomCABundle returns CustomCABundle from LoadOptions
func (o LoadOptions) getCustomCABundle(ctx context.Context) (io.Reader, bool, error) {
	if o.CustomCABundle == nil {
		return nil, false, nil
	}

	return o.CustomCABundle, true, nil
}

// WithCustomCABundle is a helper function to construct functional options
// that sets CustomCABundle on config's LoadOptions. Setting the custom CA Bundle
// to nil will result in custom CA Bundle value being ignored.
// If multiple WithCustomCABundle calls are made, the last call overrides the
// previous call values.
func WithCustomCABundle(v io.Reader) LoadOptionsFunc {
	return func(o *LoadOptions) error {
		o.CustomCABundle = v
		return nil
	}
}

// UseEC2IMDSRegion provides a regionProvider that retrieves the region
// from the EC2 Metadata service.
type UseEC2IMDSRegion struct {
	// If unset will default to generic EC2 IMDS client.
	Client *imds.Client
}

// getRegion attempts to retrieve the region from EC2 Metadata service.
func (p *UseEC2IMDSRegion) getRegion(ctx context.Context) (string, bool, error) {
	if ctx == nil {
		ctx = context.Background()
	}

	client := p.Client
	if client == nil {
		client = imds.New(imds.Options{})
	}

	result, err := p.Client.GetRegion(ctx, nil)
	if err != nil {
		return "", false, err
	}
	if len(result.Region) != 0 {
		return result.Region, true, nil
	}
	return "", false, nil
}

// getEC2IMDSRegion returns the value of EC2 IMDS region.
func (o LoadOptions) getEC2IMDSRegion(ctx context.Context) (string, bool, error) {
	if o.UseEC2IMDSRegion == nil {
		return "", false, nil
	}

	return o.UseEC2IMDSRegion.getRegion(ctx)
}

// WithEC2IMDSRegion is a helper function to construct functional options
// that enables resolving EC2IMDS region. The function takes
// in a UseEC2IMDSRegion functional option, and can be used to set the
// EC2IMDS client which will be used to resolve EC2IMDSRegion.
// If no functional option is provided, an EC2IMDS client is built and used
// by the resolver. If multiple WithEC2IMDSRegion calls are made, the last
// call overrides the previous call values. Note that the WithRegion calls takes
// precedence over WithEC2IMDSRegion when resolving region.
func WithEC2IMDSRegion(fnOpts ...func(o *UseEC2IMDSRegion)) LoadOptionsFunc {
	return func(o *LoadOptions) error {
		o.UseEC2IMDSRegion = &UseEC2IMDSRegion{}

		for _, fn := range fnOpts {
			fn(o.UseEC2IMDSRegion)
		}
		return nil
	}
}

// getCredentialsProvider returns the credentials value
func (o LoadOptions) getCredentialsProvider(ctx context.Context) (aws.CredentialsProvider, bool, error) {
	if o.Credentials == nil {
		return nil, false, nil
	}

	return o.Credentials, true, nil
}

// WithCredentialsProvider is a helper function to construct functional options
// that sets Credential provider value on config's LoadOptions. If credentials
// provider is set to nil, the credentials provider value will be ignored.
// If multiple WithCredentialsProvider calls are made, the last call overrides
// the previous call values.
func WithCredentialsProvider(v aws.CredentialsProvider) LoadOptionsFunc {
	return func(o *LoadOptions) error {
		o.Credentials = v
		return nil
	}
}

// getProcessCredentialOptions returns the wrapped function to set processcreds.Options
func (o LoadOptions) getProcessCredentialOptions(ctx context.Context) (func(*processcreds.Options), bool, error) {
	if o.ProcessCredentialOptions == nil {
		return nil, false, nil
	}

	return o.ProcessCredentialOptions, true, nil
}

// WithProcessCredentialOptions is a helper function to construct functional options
// that sets a function to use processcreds.Options on config's LoadOptions.
// If process credential options is set to nil, the process credential value will
// be ignored. If multiple WithProcessCredentialOptions calls are made, the last call
// overrides the previous call values.
func WithProcessCredentialOptions(v func(*processcreds.Options)) LoadOptionsFunc {
	return func(o *LoadOptions) error {
		o.ProcessCredentialOptions = v
		return nil
	}
}

// getEC2RoleCredentialOptions returns the wrapped function to set the ec2rolecreds.Options
func (o LoadOptions) getEC2RoleCredentialOptions(ctx context.Context) (func(*ec2rolecreds.Options), bool, error) {
	if o.EC2RoleCredentialOptions == nil {
		return nil, false, nil
	}

	return o.EC2RoleCredentialOptions, true, nil
}

// WithEC2RoleCredentialOptions is a helper function to construct functional options
// that sets a function to use ec2rolecreds.Options on config's LoadOptions. If
// EC2 role credential options is set to nil, the EC2 role credential options value
// will be ignored. If multiple WithEC2RoleCredentialOptions calls are made,
// the last call overrides the previous call values.
func WithEC2RoleCredentialOptions(v func(*ec2rolecreds.Options)) LoadOptionsFunc {
	return func(o *LoadOptions) error {
		o.EC2RoleCredentialOptions = v
		return nil
	}
}

// getEndpointCredentialOptions returns the wrapped function to set endpointcreds.Options
func (o LoadOptions) getEndpointCredentialOptions(context.Context) (func(*endpointcreds.Options), bool, error) {
	if o.EndpointCredentialOptions == nil {
		return nil, false, nil
	}

	return o.EndpointCredentialOptions, true, nil
}

// WithEndpointCredentialOptions is a helper function to construct functional options
// that sets a function to use endpointcreds.Options on config's LoadOptions. If
// endpoint credential options is set to nil, the endpoint credential options
// value will be ignored. If multiple WithEndpointCredentialOptions calls are made,
// the last call overrides the previous call values.
func WithEndpointCredentialOptions(v func(*endpointcreds.Options)) LoadOptionsFunc {
	return func(o *LoadOptions) error {
		o.EndpointCredentialOptions = v
		return nil
	}
}

// getWebIdentityRoleCredentialOptions returns the wrapped function
func (o LoadOptions) getWebIdentityRoleCredentialOptions(context.Context) (func(*stscreds.WebIdentityRoleOptions), bool, error) {
	if o.WebIdentityRoleCredentialOptions == nil {
		return nil, false, nil
	}

	return o.WebIdentityRoleCredentialOptions, true, nil
}

// WithWebIdentityRoleCredentialOptions is a helper function to construct
// functional options that sets a function to use stscreds.WebIdentityRoleOptions
// on config's LoadOptions. If web identity role credentials options is set to nil,
// the web identity role credentials value will be ignored. If multiple
// WithWebIdentityRoleCredentialOptions calls are made, the last call
// overrides the previous call values.
func WithWebIdentityRoleCredentialOptions(v func(*stscreds.WebIdentityRoleOptions)) LoadOptionsFunc {
	return func(o *LoadOptions) error {
		o.WebIdentityRoleCredentialOptions = v
		return nil
	}
}

// getAssumeRoleCredentialOptions returns AssumeRoleCredentialOptions from LoadOptions
func (o LoadOptions) getAssumeRoleCredentialOptions(context.Context) (func(options *stscreds.AssumeRoleOptions), bool, error) {
	if o.AssumeRoleCredentialOptions == nil {
		return nil, false, nil
	}

	return o.AssumeRoleCredentialOptions, true, nil
}

// WithAssumeRoleCredentialOptions  is a helper function to construct
// functional options that sets a function to use stscreds.AssumeRoleOptions
// on config's LoadOptions. If assume role credentials options is set to nil,
// the assume role credentials value will be ignored. If multiple
// WithAssumeRoleCredentialOptions calls are made, the last call overrides
// the previous call values.
func WithAssumeRoleCredentialOptions(v func(*stscreds.AssumeRoleOptions)) LoadOptionsFunc {
	return func(o *LoadOptions) error {
		o.AssumeRoleCredentialOptions = v
		return nil
	}
}

func (o LoadOptions) getHTTPClient(ctx context.Context) (HTTPClient, bool, error) {
	if o.HTTPClient == nil {
		return nil, false, nil
	}

	return o.HTTPClient, true, nil
}

// WithHTTPClient is a helper function to construct functional options
// that sets HTTPClient on LoadOptions. If HTTPClient is set to nil,
// the HTTPClient value will be ignored.
// If multiple WithHTTPClient calls are made, the last call overrides
// the previous call values.
func WithHTTPClient(v HTTPClient) LoadOptionsFunc {
	return func(o *LoadOptions) error {
		o.HTTPClient = v
		return nil
	}
}

func (o LoadOptions) getAPIOptions(ctx context.Context) ([]func(*middleware.Stack) error, bool, error) {
	if o.APIOptions == nil {
		return nil, false, nil
	}

	return o.APIOptions, true, nil
}

// WithAPIOptions is a helper function to construct functional options
// that sets APIOptions on LoadOptions. If APIOptions is set to nil, the
// APIOptions value is ignored. If multiple WithAPIOptions calls are
// made, the last call overrides the previous call values.
func WithAPIOptions(v []func(*middleware.Stack) error) LoadOptionsFunc {
	return func(o *LoadOptions) error {
		if v == nil {
			return nil
		}

		o.APIOptions = append(o.APIOptions, v...)
		return nil
	}
}

func (o LoadOptions) getRetryer(ctx context.Context) (aws.Retryer, bool, error) {
	if o.Retryer == nil {
		return nil, false, nil
	}

	return o.Retryer, true, nil
}

// WithRetryer is a helper function to construct functional options
// that sets Retryer on LoadOptions. If Retryer is set to nil, the
// Retryer value is ignored. If multiple WithRetryer calls are
// made, the last call overrides the previous call values.
func WithRetryer(v aws.Retryer) LoadOptionsFunc {
	return func(o *LoadOptions) error {
		o.Retryer = v
		return nil
	}
}

func (o LoadOptions) getEndpointResolver(ctx context.Context) (aws.EndpointResolver, bool, error) {
	if o.EndpointResolver == nil {
		return nil, false, nil
	}

	return o.EndpointResolver, true, nil
}

// WithEndpointResolver is a helper function to construct functional options
// that sets endpoint resolver on LoadOptions. The EndpointResolver is set to nil,
// the EndpointResolver value is ignored. If multiple WithEndpointResolver calls
// are made, the last call overrides the previous call values.
func WithEndpointResolver(v aws.EndpointResolver) LoadOptionsFunc {
	return func(o *LoadOptions) error {
		o.EndpointResolver = v
		return nil
	}
}

func (o LoadOptions) getLogger(ctx context.Context) (logging.Logger, bool, error) {
	if o.Logger == nil {
		return nil, false, nil
	}

	return o.Logger, true, nil
}

// WithLogger is a helper function to construct functional options
// that sets Logger on LoadOptions. If Logger is set to nil, the
// Logger value will be ignored. If multiple WithLogger calls are made,
// the last call overrides the previous call values.
func WithLogger(v logging.Logger) LoadOptionsFunc {
	return func(o *LoadOptions) error {
		o.Logger = v
		return nil
	}
}

func (o LoadOptions) getClientLogMode(ctx context.Context) (aws.ClientLogMode, bool, error) {
	if o.ClientLogMode == nil {
		return 0, false, nil
	}

	return *o.ClientLogMode, true, nil
}

// WithClientLogMode is a helper function to construct functional options
// that sets client log mode on LoadOptions. If client log mode is set to nil,
// the client log mode value will be ignored. If multiple WithClientLogMode calls are made,
// the last call overrides the previous call values.
func WithClientLogMode(v aws.ClientLogMode) LoadOptionsFunc {
	return func(o *LoadOptions) error {
		o.ClientLogMode = &v
		return nil
	}
}

func (o LoadOptions) getLogConfigurationWarnings(ctx context.Context) (v bool, found bool, err error) {
	if o.LogConfigurationWarnings == nil {
		return false, false, nil
	}
	return *o.LogConfigurationWarnings, true, nil
}

// WithLogConfigurationWarnings is a helper function to construct
// functional options that can be used to set LogConfigurationWarnings
// on LoadOptions.
//
// If multiple WithLogConfigurationWarnings calls are made, the last call
// overrides the previous call values.
func WithLogConfigurationWarnings(v bool) LoadOptionsFunc {
	return func(o *LoadOptions) error {
		o.LogConfigurationWarnings = &v
		return nil
	}
}

// GetS3UseARNRegion returns whether to allow ARNs to direct the region
// the S3 client's requests are sent to.
func (o LoadOptions) GetS3UseARNRegion(ctx context.Context) (v bool, found bool, err error) {
	if o.S3UseARNRegion == nil {
		return false, false, nil
	}
	return *o.S3UseARNRegion, true, nil
}

// WithS3UseARNRegion is a helper function to construct functional options
// that can be used to set S3UseARNRegion on LoadOptions.
// If multiple WithS3UseARNRegion calls are made, the last call overrides
// the previous call values.
func WithS3UseARNRegion(v bool) LoadOptionsFunc {
	return func(o *LoadOptions) error {
		o.S3UseARNRegion = &v
		return nil
	}
}
