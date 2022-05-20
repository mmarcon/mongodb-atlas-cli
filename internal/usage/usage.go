// Copyright 2021 MongoDB Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package usage

const (
	ProjectID                                = "Project ID to use. Overrides the settings in the configuration file or environment variable."
	OrgID                                    = "Organization ID to use. Overrides the settings in the configuration file or environment variable."
	Profile                                  = "Profile to use from your configuration file."
	Members                                  = "Number of members in the replica set."
	Shards                                   = "Number of shards in the cluster."
	ProcessName                              = "Unique identifier for the host of a MongoDB process in the following format: {hostname}:{port}."
	Since                                    = "Point in time, specified as milliseconds since the Unix Epoch, from which you want to receive results."
	HostID                                   = "Unique identifier for the host of a MongoDB process."
	Duration                                 = "Length of time from the since parameter, in milliseconds, for which you want to receive results."
	Tier                                     = "Tier for each data-bearing server in the cluster."
	NLog                                     = "Maximum number of log lines to return."
	SlowQueryNamespaces                      = "Namespaces from which to retrieve suggested slow query logs."
	DiskSizeGB                               = "Capacity, in gigabytes, of the host's root volume."
	Backup                                   = "If true, enables Continuous Cloud Backup for your deployment."
	BIConnector                              = "If true, enables BI Connector for Atlas on the deployment."
	SuggestedIndexNamespaces                 = "Namespaces from which to retrieve suggested indexes."
	NExamples                                = "Maximum number of example queries to provide that will be improved by a suggested index."
	NIndexes                                 = "Maximum number of indexes to suggest."
	MDBVersion                               = "MongoDB version of the cluster to deploy."
	AuthDB                                   = "Authentication database name."
	Granularity                              = "Duration in ISO 8601 notation that specifies the interval between measurement data points. Only the following subset of ISO 8601-formatted time periods are supported: PT10S, PT1M, PT5M, PT1H, P1D."
	Page                                     = "Page number."
	Forever                                  = "Acknowledges an alert 'forever'."
	Status                                   = "Alert's status. The values are TRACKING, OPEN, CLOSED, and CANCELED."
	Until                                    = "Time, in ISO 8601 format, until which the alert has been acknowledged. mongocli returns this value after the alert has been acknowledged."
	ConnectionStringType                     = "When set to 'private', retrieves the connection string for the network peering endpoint."
	Limit                                    = "Number of items per page."
	Username                                 = "Username of the user."
	BackupStatus                             = "Current (or desired) status of the backup configuration."
	StorageEngine                            = "Storage engine used for the backup."
	AuthMechanism                            = "Authentication mechanism needed to connect to the sync source database."
	Provisioned                              = "Flag that indicates if Ops Manager has provisioned the resources needed to store a backup."
	Encryption                               = "Flag that indicates if encryption is enabled for the backup configuration."
	SSL                                      = "Flag that indicates if TLS is enabled for the sync source database."
	OplogSSL                                 = "Flag indicating whether this oplog store only accepts connections encrypted using TLS."
	SyncSource                               = "mongod instance from which you retrieve backup data."
	ExcludedNamespace                        = "List of database names and collection names to omit from the backup."
	IncludedNamespace                        = "List of database names and collection names to include in the backup."
	TeamUsername                             = "List of usernames to add to the new team."
	DBUsername                               = "Username for authenticating to MongoDB."
	TeamName                                 = "Name of the team."
	UserID                                   = "Unique identifier of the user."
	LDAPHostname                             = "Hostname or IP address of the LDAP server."
	LDAPPort                                 = "Port to which the LDAP server listens to for client connections."
	Hostname                                 = "The fully qualified domain name (FQDN) of the target node that should receive the authentication attempt."
	BindUsername                             = "The user distinguished name (DN) that Atlas uses to connect to the LDAP server."
	BindPassword                             = "Password used to authenticate the bindUsername."
	CaCertificate                            = "CA certificate used to verify the identify of the LDAP server."
	AuthzQueryTemplate                       = "An LDAP query template that Atlas executes to obtain the LDAP groups to which the authenticated user belongs."
	MappingMatch                             = "An ECMAScript-formatted regular expression (regex) to match against a provided username."
	MappingLdapQuery                         = "An LDAP query formatting template that inserts the authentication name matched by the match regex into an LDAP query URI encoded respecting RFC4515 and RFC4516."
	MappingSubstitution                      = "An LDAP distinguished name (DN) formatting template that converts the authentication name matched by the match regex into a LDAP DN."
	AuthenticationEnabled                    = "Specifies whether user authentication with LDAP is enabled."
	AuthorizationEnabled                     = "Specifies whether user authorization with LDAP is enabled."
	TeamID                                   = "Unique identification of the team."
	Password                                 = "User's password."
	Country                                  = "The ISO 3166-1 alpha two-letter country code of the user's country of residence."
	Mobile                                   = "The user's mobile or cell phone number."
	Period                                   = "Duration in ISO 8601 that specifies how far back in the past to retrieve measurements."
	Roles                                    = "User's roles and the databases or collections on which the roles apply."
	Scopes                                   = "Array of clusters and Atlas Data Lakes that this user has access to."
	DataLakeRole                             = "Amazon Resource Name (ARN) of the role which Atlas Data Lake uses for accessing the data stores."
	DataLakeRegion                           = "Name of the region to which Atlas Data Lake routes client connections for data processing."
	DataLakeTestBucket                       = "Name of an Amazon S3 data bucket which Atlas Data Lake uses to validate the provided role."
	PrivateEndpointRegion                    = "Cloud provider region in which you want to create the private endpoint connection."
	PrivateEndpointProvider                  = "Name of the cloud provider you want to create the private endpoint connection for."
	Comment                                  = "Optional description or comment for the entry."
	AccessListsDeleteAfter                   = "ISO-8601-formatted UTC date after which Atlas removes the value from the entry."
	BDUsersDeleteAfter                       = "Timestamp in ISO 8601 in UTC after which Atlas deletes the user."
	ForceVersionManifest                     = "If specified, skips Ops Manager version check."
	Force                                    = "If specified, skips asking for confirmation before proceeding with a requested action."
	ForceFile                                = "Overwrites the destination file."
	Email                                    = "User's email address."
	LogOut                                   = "Optional output file name. Uses the log name if the output file's name is not specified."
	DiagnoseOut                              = "Optional output file name. Uses diagnose-archive.tar.gz if the output file's name is not specified."
	LogStart                                 = "Beginning of the date and time, in UNIX Epoch format, from which to retrieve logs."
	LogEnd                                   = "End of the period, in UNIX Epoch format, until which to retrieve logs."
	ArchiveLimit                             = "Max number of entries for the diagnose archive."
	ArchiveMinutes                           = "Beginning of the period, in UNIX Epoch format, from which to start retrieving the diagnose archive. Ops Manager takes out minutes from the current time."
	MeasurementStart                         = "Beginning of the period, in UNIX Epoch format, from which to start retrieving measurements."
	MeasurementEnd                           = "End of the period, in UNIX Epoch format, until which to retrieve measurements."
	MeasurementType                          = "Measurements to return. If it is not specified, all measurements are returned."
	FirstName                                = "User's first name."
	LastName                                 = "User's last name."
	OrgRole                                  = "User's roles  for the associated organization."
	ProjectRole                              = "User's roles  for the associated project."
	TeamRole                                 = "Project role you want to assign to the team."
	MaxDate                                  = "Returns events whose created date is less than or equal to it."
	MinDate                                  = "Returns events whose created date is greater than or equal to it."
	ClusterFilename                          = "File name to use, optional file with a json cluster configuration."
	PoliciesFilename                         = "File name to use, optional file with a json policy configuration."
	SearchFilename                           = "File name to use, file with a json index configuration."
	AccessListIps                            = "IP addresses to add to the new user's access list."
	StartDate                                = "Timestamp in ISO 8601 date and time format in UTC when the maintenance window starts."
	EndDate                                  = "Timestamp in ISO 8601 date and time format in UTC when the maintenance window ends."
	AuthResult                               = "Flag that indicates whether to return either successful or failed authentication attempts. When set to success, Atlas filters the log to return only successful authentication attempts. When set to fail, Atlas filters the log to return only failed authentication attempts."
	AccessLogDate                            = "Timestamp in the number of milliseconds that have elapsed, since the UNIX Epoch, for the first entry that Atlas returns from the database access logs."
	AccessLogIP                              = "IP address that attempted to authenticate with the database. Atlas filters the returned logs to include documents with only this IP address."
	ServerUsageStartDate                     = "Timestamp, in ISO 8601 date format, when the list of host assignments starts."
	ServerUsageEndDate                       = "Timestamp, in ISO 8601 date format, when the list of host assignments ends."
	AlertType                                = "Alert types to silence during maintenance window. Valid values are HOST, REPLICA_SET, CLUSTER, AGENT, or BACKUP."
	MaintenanceDescription                   = "Description of the maintenance window."
	Event                                    = "Type of the event that triggered the alert."
	Enabled                                  = "If set to true, the alert configuration is enabled."
	MatcherFieldName                         = "Name of the field in the target object to match on."
	MatcherOperator                          = "Operator to test the field's value."
	MatcherValue                             = "Value to test with the specified operator."
	MetricName                               = "Name of the metric against which Atlas checks the configured alert."
	MetricOperator                           = "Operator to apply when checking the current metric value against the threshold value."
	MetricThreshold                          = "Threshold value outside of which an alert will be triggered."
	MetricUnits                              = "Units for the threshold value."
	MetricMode                               = "If specified, Atlas computes the current metric value as an average."
	NotificationToken                        = "Slack API token, Bot token, or Flowdock personal API token." //nolint:gosec // This is just a message not a password
	NotificationsChannelName                 = "Slack channel name. Required for the SLACK notifications type."
	AlertConfigAPIKey                        = "Datadog API Key, Opsgenie API Key, VictorOps API key." //nolint:gosec // This is just a message not a credential
	APIKey                                   = "API Key."
	RoutingKey                               = "An optional field for your Routing Key."
	IntegrationAPIToken                      = "Your API Token." //nolint:gosec // This is just a message not a credential
	OrgName                                  = "Your Flowdock organization's name."
	OrgNameFilter                            = "Performs a case-insensitive search for organizations which exactly match the specified name."
	OrgIncludeDeleted                        = "If specified, Atlas includes the deleted organizations."
	FlowName                                 = "Flowdock Flow name."
	BlockstoreAssignment                     = "Flag indicating whether this blockstore can be assigned backup jobs."
	OplogAssignment                          = "Flag indicating whether this oplog can be assigned backup jobs."
	FileSystemAssignment                     = "Flag indicating whether this file system store can be assigned backup jobs."
	SyncAssignment                           = "Flag indicating whether this sync store can be assigned backup jobs."
	EncryptedCredentials                     = "Flag indicating whether the username and password were encrypted using the credentials tool."
	MMAPV1CompressionSetting                 = "The compression setting for the MMAPv1 storage engine snaphots."
	WTCompressionSetting                     = "The compression setting for the WiredTiger storage engine snaphots."
	StorePath                                = "The location where file system-based backups are stored on the file system store host."
	Label                                    = "Array of tags to manage which backup jobs Ops Manager can assign to which blockstores."
	LoadFactor                               = "A positive, non-zero integer that expresses how much backup work this snapshot store should perform compared to another snapshot store."
	MaxCapacityGB                            = "The maximum amount of data in GB this blockstore can store."
	BlockstoreURI                            = "A comma-separated list of hosts in the <hostname:port> format that can be used to access this blockstore."
	BlockstoreSSL                            = "Flag indicating whether this blockstore only accepts connections encrypted using TLS."
	BlockstoreName                           = "Unique name that labels this blockstore."
	OplogName                                = "Unique name that labels this oplog store."
	FileSystemName                           = "Unique name that labels this file system store configuration."
	SyncName                                 = "Unique name that labels this sync store."
	WriteConcern                             = "The write concern used for this blockstore."
	AWSAccessKey                             = "AWS Access Key ID that can access the Amazon S3 bucket specified in s3BucketName."
	AWSSecretKey                             = "AWS Secret Access Key that can access the Amazon S3 bucket specified in s3BucketName." //nolint:gosec // This is just a message not a credential
	DisableProxyS3                           = "Flag indicating whether the HTTP proxy should be used when connecting to Amazon S3."
	S3BucketEndpoint                         = "URL that Ops Manager uses to access this Amazon S3 or S3-compatible bucket."
	S3BucketName                             = "Name of the Amazon S3 bucket that hosts the S3 blockstore."
	S3MaxConnections                         = "Positive integer indicating the maximum number of connections to this Amazon S3 blockstore."
	AcceptedTos                              = "Flag indicating whether or not you accepted the terms of service for using Amazon S3-compatible stores with Ops Manager."
	SSEEnabled                               = "Flag indicating whether this Amazon S3 blockstore enables server-side encryption."
	PathStyleAccessEnabled                   = "Flag indicating the style of this endpoint."
	APIKeyDescription                        = "Description of the API key."
	APIKeyRoles                              = "List of roles for the API key." //nolint:gosec // This is just a message not a credential
	NotificationRegion                       = "Region that indicates which API URL to use."
	NotificationDelayMin                     = "Number of minutes to wait after an alert condition is detected before sending out the first notification."
	NotificationEmailAddress                 = "Email address to which alert notifications are sent."
	NotificationEmailEnabled                 = "Flag indicating whether email notifications should be sent."
	NotificationFlowName                     = "Flowdock Flow name in lower-case letters for sending alert notifications."
	NotificationIntervalMin                  = "Number of minutes to wait between successive notifications for unacknowledged alerts that are not resolved."
	NotificationMobileNumber                 = "Mobile number to which alert notifications are sent."
	NotificationOrgName                      = "Flowdock organization's name in lower-case letters."
	NotificationServiceKey                   = "PagerDuty service key."
	NotificationSmsEnabled                   = "Flag indicating whether text message notifications should be sent."
	NotificationTeamID                       = "Unique identifier of a team."
	NotificationType                         = "Type of alert notification. The values are DATADOG, EMAIL, FLOWDOCK, GROUP (Project), ORG, OPS_GENIE, PAGER_DUTY, SLACK, SMS, USER, or VICTOR_OPS."
	NotificationUsername                     = "Name of the Atlas user to which to send notifications."
	NotificationVictorOpsRoutingKey          = "VictorOps routing key."
	SnapshotID                               = "Unique identifier of the snapshot to restore."
	SnapshotDescription                      = "Description of the on-demand snapshot."
	Database                                 = "Database name."
	DatabaseUser                             = "Username of a database user."
	MonthsUntilExpiration                    = "Number of months that the certificate is valid for."
	Collection                               = "Collection name."
	Append                                   = "The input action and inheritedRoles that will be appended to the existing role."
	PrivilegeAction                          = "List of actions per database and collection. If no database or collections are provided, cluster scope is assumed."
	InheritedRoles                           = "List of inherited roles and the database on which the role is granted."
	Analyzer                                 = "Analyzer to use when creating the index."
	SearchAnalyzer                           = "Analyzer to use when searching the index."
	Dynamic                                  = "Indicates whether the index uses dynamic or static mappings."
	SearchFields                             = "Static field specifications."
	RSName                                   = "The replica set that the index is built on."
	Key                                      = "Field to be indexed and the type of index in the following format: field:type."
	LogTypes                                 = "Array of strings specifying the types of logs to collect."
	SizeRequestedPerFileBytes                = "Size for each log file in bytes."
	LogRedacted                              = "If set to true, emails, hostnames, IP addresses, and namespaces in API responses involving this job are replaced with random string values."
	Sparse                                   = "Flag that indicates whether Atlas should create a sparse index."
	Locale                                   = "Locale that the ICU defines."
	CaseLevel                                = "Flag that indicates whether the index uses case comparison. This flag applies only if the strength level is set to 1 or 2."
	CaseFirst                                = "Flag that, if specified, determines the sort order of case differences during tertiary level comparisons."
	Strength                                 = "Level of comparison to perform."
	Alternate                                = "Flag that, if specified, determines whether collation should consider whitespace and punctuation as base characters during comparisons."
	MaxVariable                              = "Flag that, if specified, determines which characters to ignore. This flag applies only if indexConfigs.collation.alternate is set to shifted."
	NumericOrdering                          = "Flag that, if set to true, indicates that collation compares numeric strings as numbers. If set to false, collation compares numeric strings as strings."
	Normalization                            = "Flag that, if set to true, collation checks if text requires normalization and performs normalization to compare text."
	Backwards                                = "Flag that, if set to true, strings with diacritics sort from the back to the front of the string."
	ClusterName                              = "Name of the cluster."
	IndexName                                = "Name of the index."
	CASFilePath                              = "Path to a PEM file containing one or more CAs for database user authentication."
	Verbose                                  = "Flag that, if set to true, returns all child jobs in the response."
	ClusterID                                = "Unique identifier of the cluster."
	ReferenceTimeZoneOffset                  = "The ISO-8601-formatted timezone offset where the Ops Manager host resides."
	DailySnapshotRetentionDays               = "Number of days to retain daily snapshots. Accepted values are between 1 and 365 inclusive."
	SnapshotRetentionDays                    = "Number of days to keep recent snapshots. Accepted values are between 2 and 5 inclusive."
	WeeklySnapshotRetentionWeeks             = "Number of weeks to retain weekly snapshots. Accepted values are between 1 and 52 inclusive."
	PointInTimeWindowHours                   = "Number of hours in the past for which MongoDB should create a point-in-time snapshot."
	ReferenceHourOfDay                       = "Hour of the day to schedule snapshots using a 24-hour clock. Accepted values are between 0 and 23 inclusive."
	ReferenceMinuteOfHour                    = "Minute of the hour to schedule snapshots. Accepted values are between 0 and 59 inclusive."
	MonthlySnapshotRetentionMonths           = "Number of months to retain monthly snapshots. Accepted values are between 1 and 36 inclusive."
	Policy                                   = "List of policies that the external system applies to this Ops Manager Project."
	SystemID                                 = "Unique identifier of the external system that manages this Ops Manager Project."
	ExternalSystemName                       = "Identifying label for the external system that manages this Ops Manager Project."
	DateField                                = "Name of an already indexed date field from the documents."
	PartitionFields                          = "Fields to use to partition data. You can specify up to two frequently queried fields to use for partitioning data."
	ArchiveAfter                             = "Number of days that specifies the age limit for the data in the live Atlas cluster."
	TargetProjectID                          = "Unique identifier of the project that contains the destination cluster for the restore job."
	APIAccessListIPEntry                     = "IP address to be allowed for a given API key."
	NetworkAccessListIPEntry                 = "IP address to be allowed to access the deployment."
	AccessListCIDREntry                      = "Access list entry in CIDR notation to be added for a given API key."
	LinkTokenAccessListCIDREntries           = "IP address access list entries that are associated with the link-token." //nolint:gosec // This is just a message not a credential
	LinkToken                                = "Link-token generated by Atlas."                                          //nolint:gosec // This is just a message not a credential
	LiveMigrationID                          = "Unique 24-hexadecimal digit string that identifies the live migration job."
	PrivateEndpointID                        = "Unique identifier of the AWS PrivateLink connection."
	EndpointServiceID                        = "Unique identifier of the private endpoint service for which you want to retrieve a private endpoint."
	PrivateEndpointIDAzure                   = "Unique identifier of the Azure private endpoint resource."
	PrivateEndpointIPAddressAzure            = "Private IP address of the private endpoint network interface you created in your Azure VNet."
	Endpoint                                 = "List of GCP endpoints in the group separated by commas, such as: endpointName1@ipAddress1,...,endpointNameN@ipAddressN"
	AccountID                                = "Account ID of the owner of the peer VPC."
	NewRelicAccountID                        = "Unique identifier of your New Relic account."
	LicenceKey                               = "Your License Key."
	ServiceKey                               = "Your Service Key."
	URL                                      = "Your webhook URL."
	Secret                                   = "An optional field for your webhook secret." //nolint:gosec // This is just a message not a credential
	WriteToken                               = "Your Insights Insert Key."
	DayOfWeek                                = "Day of the week that you want the maintenance window to start, as a 1-based integer."
	HourOfDay                                = "Hour of the day that you want the maintenance window to start. This parameter uses the 24-hour clock, where midnight is 0 and noon is 12."
	StartASAP                                = "Flag that, if specified, indicates to start maintenance immediately upon receiving this request."
	ReadToken                                = "Your Insights Query Key."
	RouteTableCidrBlock                      = "Peer VPC CIDR block or subnet."
	VpcID                                    = "Unique identifier of the peer VPC."
	AtlasCIDRBlock                           = "CIDR block that Atlas uses for your clusters."
	VNet                                     = "Name of your Azure VNet."
	ResourceGroup                            = "Name of your Azure resource group."
	IAMAssumedRoleARN                        = "Role ARN that Atlas assumes to access your AWS account."
	DirectoryID                              = "Unique identifier for an Azure AD directory."
	SubscriptionID                           = "Unique identifier of the Azure subscription in which the VNet resides."
	GCPProjectID                             = "Unique identifier of the GCP project in which the network peer resides."
	Network                                  = "Unique identifier of the Network Peering connection in the Atlas project."
	APIRegion                                = "Indicates which API URL to use, either US or EU. The integration service will use US by default."
	SkipMongosh                              = "Indicates whether to skip accessing your deployment with MongoDB Shell."
	SkipSampleData                           = "Indicates whether to skip loading sample data into your Atlas cluster."
	ContainerRegion                          = "Atlas region where the container resides."
	ContainerRegions                         = "List of Atlas regions where the container resides."
	ProjectOwnerID                           = "Unique 24-hexadecimal digit string that identifies the Atlas user account to be granted the Project Owner role on the specified project."
	GovCloudRegionsOnly                      = "Only for AtlasGov projects. If specified, designates that the project uses only the AWS GovCloud region. If unspecified, the project uses only the AWS Standard region. You can't deploy clusters across AWS GovCloud and AWS Standard regions in the same project."
	ReclaimFreeSpaceTimestamp                = "Timestamp in ISO 8601 format when the service reclaims the space. If not set, defaults to the current timestamp."
	QuickstartDefault                        = "Flag that indicates whether to run the Quickstart command with all the auto-generated values to deploy and access an Atlas cluster."
	ServerlessProvider                       = "Cloud service provider that applies to the provisioned serverless instance."
	ServerlessRegion                         = "Human-readable label that identifies the physical location of your MongoDB serverless instance. The region you choose can affect network latency for clients accessing your databases."
	WithoutDefaultAlertSettings              = "Flag that indicates whether to create the new project with the default alert settings enabled. This flag defaults to false. In this case, Atlas creates the project without the default alert settings. This is useful if you create projects programmatically and want to create your own alerts instead of using the default alert settings. To create a project that uses the default alert settings, set this flag to true."
	FormatOut                                = "Output format. Valid values are json, json-path, go-template, or go-template-file."
	TargetClusterID                          = "Unique identifier of the target cluster. For use only with automated restore jobs."
	TargetClusterName                        = "Name of the target cluster. For use only with automated restore jobs."
	CheckpointID                             = "Unique identifier for the sharded cluster checkpoint that represents the point in time to which your data will be restored. If you set checkpointId, you cannot set oplogInc, oplogTs, snapshotId, or pointInTimeUTCMillis."
	OplogTS                                  = "Oplog timestamp given as a timestamp in the number of seconds that have elapsed since the UNIX Epoch. When paired with oplogInc, they represent the point in time to which your data will be restored."
	OplogInc                                 = "32-bit incrementing ordinal that represents operations within a given second. When paired with oplogTs, they represent the point in time to which your data will be restored."
	PointInTimeUTCMillis                     = "Timestamp in the number of milliseconds that have elapsed since the UNIX epoch that represents the point in time to which your data will be restored. This timestamp must be within the last 24 hours of the current time."
	Expires                                  = "Timestamp in ISO 8601 date and time format after which the URL is no longer available. For use only with download restore jobs."
	ExpirationHours                          = "Number of hours the download URL is valid once the restore job is complete. For use only with download restore jobs."
	MaxDownloads                             = "Number of times the download URL can be used. This value must be 1 or greater. For use only with download restore jobs."
	Mechanisms                               = "Authentication mechanism. Valid values are SCRAM-SHA-1 or SCRAM-SHA-256."
	AccessListType                           = "Type of access list entry. Valid values are cidrBlock, ipAddress, or awsSecurityGroup."
	Service                                  = "Type of MongoDB service. Valid values are cloud, cloudgov, cloud-manager, or ops-manager."
	Provider                                 = "Name of your cloud service provider. Valid values are AWS, AZURE, or GCP."
	ClusterTypes                             = "Type of the cluster that you want to create. Valid values are REPLICASET or SHARDED."
	Region                                   = "Physical location of your MongoDB cluster. For a complete list of supported AWS regions, see: https://docs.atlas.mongodb.com/reference/amazon-aws/#amazon-aws. For a complete list of supported Azure regions, see: https://docs.atlas.mongodb.com/reference/microsoft-azure/#microsoft-azure. For a complete list of supported GCP regions, see: https://docs.atlas.mongodb.com/reference/google-gcp/#google-gcp."
	AWSIAMType                               = "AWS IAM method by which the provided username is authenticated. Valid values are NONE, USER, or ROLE."
	X509Type                                 = "X.509 method by which the provided username is authenticated.  Valid values are NONE, MANAGED, or CUSTOMER."
	LDAPType                                 = "LDAP method by which the provided username is authenticated. Valid values are NONE, USER, or GROUP."
	DateFormat                               = "Date format for the date field. Valid values are ISODATE, EPOCH_SECONDS, EPOCH_MILLIS, or EPOCH_NANOSECONDS."
	ServerUsageFormat                        = "Compression format of the resulting report. Valid values are ZIP, TAR, or .GZ."
	S3AuthMethod                             = "Method used to authorize access to the Amazon S3 bucket specified in s3BucketName. Valid values are KEYS or IAM_ROLE."
	ClusterCheckpointIntervalMin             = "Number of minutes between successive cluster checkpoints. Valid values are 15, 30, or 60."
	SnapshotIntervalHours                    = "Number of hours between snapshots. Valid values are 6, 8, 12, or 24."
	LiveMigrationDestinationClusterName      = "Human-readable label that identifies the Atlas destination cluster."
	LiveMigrationHostEntries                 = "List of hosts running the MongoDB Agent that can transfer your MongoDB data from the source (Cloud Manager or Ops Manager) to destination (Atlas) deployments. Each live migration process uses its own dedicated migration host."
	LiveMigrationSourceClusterName           = "Human-readable label that identifies the source Cloud Manager or Ops Manager cluster."
	LiveMigrationSourceProjectID             = "Unique 24-hexadecimal digit string that identifies the source project."
	LiveMigrationSourceSSL                   = "Flag that indicates whether data source has TLS enabled."
	LiveMigrationSourceCACertificatePath     = "Path to the CA certificate that signed TLS certificates use to authenticate to the source Cloud Manager or Ops Manager cluster. Omit this value if --sourceSSL is not passed."
	LiveMigrationSourceManagedAuthentication = "Flag that indicates whether MongoDB Automation manages authentication to the source Cloud Manager or Ops Manager cluster. If you set this to true, don't provide values for --sourceUsername and --sourcePassword."
	LiveMigrationSourceUsername              = "Human-readable label that identifies the SCRAM-SHA user that connects to the source Cloud Manager or Ops Manager cluster. Omit this value if --sourceManagedAuthentication is set."
	LiveMigrationSourcePassword              = "Password that authenticates the username to the source Cloud Manager or Ops Manager cluster. Omit this value if --sourceManagedAuthentication is passed."
	LiveMigrationDropCollections             = "Flag that indicates whether this process should drop existing collections from the destination (Atlas) cluster given in --destinationClusterName before starting the migration of data from the source cluster."
	LiveMigrationValidationID                = "Unique 24-hexadecimal digit string that identifies the validation job."
	CurrentIP                                = "Flag that indicates whether to use the IP Address from the host that is currently executing the command. Only applicable for type ipAddress entries. To learn more, see: https://docs.mongodb.com/mongocli/master/command/mongocli-atlas-accessLists-create/."
	CurrentIPSimplified                      = "Flag that indicates whether to use the IP Address from the host that is currently executing the command."
	Gov                                      = "Create a default profile for atlas for gov"
	EncryptedLogFile                         = "Path to the file that contains encrypted audit logs."
	OutputLogFile                            = "Path to the file where MongoCLI will save the contents of the decrypted audit log. If not specified, MongoCLI writes the contents of the decrypted audit log to stdout."
	LocalKeyFile                             = "Path to the file that contains the Key Encryption Key (KEK) that is used to encrypt the Log Encryption Key (LEK)."
	KMIPServerCAFile                         = "Path to the CA file used to connect to the server that supports KMIP."
	KMIPClientCertificateFile                = "Path to the Client Certificate file used to connect to the server that supports Key Management Interoperability Protocol (KMIP)."
	KMIPClientCertificatePassword            = "Password to decrypt the Private Key of the Client Certificate used to connect to the server that supports KMIP."
	KMIPUsername                             = "Username for authenticating to the server that supports KMIP."
	KMIPPassword                             = "Password that authenticates the username to the server that supports KMIP."
	GCPServiceAccountKey                     = "GCP service account key file."
	AzureClientID                            = "Application (client) ID assigned by Azure portal - App registrations experience."
	AzureTenantID                            = "Tenant value in the path of the request can be used to control who can sign into the application."
	AzureSecret                              = "Application secret created in the Azure app registration portal."
	DecryptAWSAccessKey                      = "AWS Access Key ID that is part of long-term credentials."
	DecryptAWSSecretKey                      = "AWS Secret Access Key  that is part of long-term credentials." //nolint:gosec // This is just a message not a credential
	AWSSessionToken                          = "AWS session token used with temporary security credentials."   //nolint:gosec // This is just a message not a credential
	Keep                                     = "If specified, skips deleting config file."
)
