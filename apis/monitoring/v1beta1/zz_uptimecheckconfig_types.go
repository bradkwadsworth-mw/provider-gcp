/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type AcceptedResponseStatusCodesInitParameters struct {

	// A class of status codes to accept.
	// Possible values are: STATUS_CLASS_1XX, STATUS_CLASS_2XX, STATUS_CLASS_3XX, STATUS_CLASS_4XX, STATUS_CLASS_5XX, STATUS_CLASS_ANY.
	StatusClass *string `json:"statusClass,omitempty" tf:"status_class,omitempty"`

	// A status code to accept.
	StatusValue *float64 `json:"statusValue,omitempty" tf:"status_value,omitempty"`
}

type AcceptedResponseStatusCodesObservation struct {

	// A class of status codes to accept.
	// Possible values are: STATUS_CLASS_1XX, STATUS_CLASS_2XX, STATUS_CLASS_3XX, STATUS_CLASS_4XX, STATUS_CLASS_5XX, STATUS_CLASS_ANY.
	StatusClass *string `json:"statusClass,omitempty" tf:"status_class,omitempty"`

	// A status code to accept.
	StatusValue *float64 `json:"statusValue,omitempty" tf:"status_value,omitempty"`
}

type AcceptedResponseStatusCodesParameters struct {

	// A class of status codes to accept.
	// Possible values are: STATUS_CLASS_1XX, STATUS_CLASS_2XX, STATUS_CLASS_3XX, STATUS_CLASS_4XX, STATUS_CLASS_5XX, STATUS_CLASS_ANY.
	// +kubebuilder:validation:Optional
	StatusClass *string `json:"statusClass,omitempty" tf:"status_class,omitempty"`

	// A status code to accept.
	// +kubebuilder:validation:Optional
	StatusValue *float64 `json:"statusValue,omitempty" tf:"status_value,omitempty"`
}

type AuthInfoInitParameters struct {

	// The username to authenticate.
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

type AuthInfoObservation struct {

	// The username to authenticate.
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

type AuthInfoParameters struct {

	// The password to authenticate.
	// Note: This property is sensitive and will not be displayed in the plan.
	// +kubebuilder:validation:Required
	PasswordSecretRef v1.SecretKeySelector `json:"passwordSecretRef" tf:"-"`

	// The username to authenticate.
	// +kubebuilder:validation:Optional
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

type ContentMatchersInitParameters struct {

	// String or regex content to match (max 1024 bytes)
	Content *string `json:"content,omitempty" tf:"content,omitempty"`

	// Information needed to perform a JSONPath content match. Used for ContentMatcherOption::MATCHES_JSON_PATH and ContentMatcherOption::NOT_MATCHES_JSON_PATH.
	// Structure is documented below.
	JSONPathMatcher []JSONPathMatcherInitParameters `json:"jsonPathMatcher,omitempty" tf:"json_path_matcher,omitempty"`

	// The type of content matcher that will be applied to the server output, compared to the content string when the check is run.
	// Default value is CONTAINS_STRING.
	// Possible values are: CONTAINS_STRING, NOT_CONTAINS_STRING, MATCHES_REGEX, NOT_MATCHES_REGEX, MATCHES_JSON_PATH, NOT_MATCHES_JSON_PATH.
	Matcher *string `json:"matcher,omitempty" tf:"matcher,omitempty"`
}

type ContentMatchersObservation struct {

	// String or regex content to match (max 1024 bytes)
	Content *string `json:"content,omitempty" tf:"content,omitempty"`

	// Information needed to perform a JSONPath content match. Used for ContentMatcherOption::MATCHES_JSON_PATH and ContentMatcherOption::NOT_MATCHES_JSON_PATH.
	// Structure is documented below.
	JSONPathMatcher []JSONPathMatcherObservation `json:"jsonPathMatcher,omitempty" tf:"json_path_matcher,omitempty"`

	// The type of content matcher that will be applied to the server output, compared to the content string when the check is run.
	// Default value is CONTAINS_STRING.
	// Possible values are: CONTAINS_STRING, NOT_CONTAINS_STRING, MATCHES_REGEX, NOT_MATCHES_REGEX, MATCHES_JSON_PATH, NOT_MATCHES_JSON_PATH.
	Matcher *string `json:"matcher,omitempty" tf:"matcher,omitempty"`
}

type ContentMatchersParameters struct {

	// String or regex content to match (max 1024 bytes)
	// +kubebuilder:validation:Optional
	Content *string `json:"content,omitempty" tf:"content,omitempty"`

	// Information needed to perform a JSONPath content match. Used for ContentMatcherOption::MATCHES_JSON_PATH and ContentMatcherOption::NOT_MATCHES_JSON_PATH.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	JSONPathMatcher []JSONPathMatcherParameters `json:"jsonPathMatcher,omitempty" tf:"json_path_matcher,omitempty"`

	// The type of content matcher that will be applied to the server output, compared to the content string when the check is run.
	// Default value is CONTAINS_STRING.
	// Possible values are: CONTAINS_STRING, NOT_CONTAINS_STRING, MATCHES_REGEX, NOT_MATCHES_REGEX, MATCHES_JSON_PATH, NOT_MATCHES_JSON_PATH.
	// +kubebuilder:validation:Optional
	Matcher *string `json:"matcher,omitempty" tf:"matcher,omitempty"`
}

type HTTPCheckInitParameters struct {

	// If present, the check will only pass if the HTTP response status code is in this set of status codes. If empty, the HTTP status code will only pass if the HTTP status code is 200-299.
	// Structure is documented below.
	AcceptedResponseStatusCodes []AcceptedResponseStatusCodesInitParameters `json:"acceptedResponseStatusCodes,omitempty" tf:"accepted_response_status_codes,omitempty"`

	// The authentication information. Optional when creating an HTTP check; defaults to empty.
	// Structure is documented below.
	AuthInfo []AuthInfoInitParameters `json:"authInfo,omitempty" tf:"auth_info,omitempty"`

	// The request body associated with the HTTP POST request. If contentType is URL_ENCODED, the body passed in must be URL-encoded. Users can provide a Content-Length header via the headers field or the API will do so. If the requestMethod is GET and body is not empty, the API will return an error. The maximum byte size is 1 megabyte. Note - As with all bytes fields JSON representations are base64 encoded. e.g. "foo=bar" in URL-encoded form is "foo%3Dbar" and in base64 encoding is "Zm9vJTI1M0RiYXI=".
	Body *string `json:"body,omitempty" tf:"body,omitempty"`

	// The content type to use for the check.
	// Possible values are: TYPE_UNSPECIFIED, URL_ENCODED.
	ContentType *string `json:"contentType,omitempty" tf:"content_type,omitempty"`

	// The list of headers to send as part of the uptime check request. If two headers have the same key and different values, they should be entered as a single header, with the value being a comma-separated list of all the desired values as described at https://www.w3.org/Protocols/rfc2616/rfc2616.txt (page 31). Entering two separate headers with the same key in a Create call will cause the first to be overwritten by the second. The maximum number of headers allowed is 100.
	Headers map[string]*string `json:"headers,omitempty" tf:"headers,omitempty"`

	// Boolean specifying whether to encrypt the header information. Encryption should be specified for any headers related to authentication that you do not wish to be seen when retrieving the configuration. The server will be responsible for encrypting the headers. On Get/List calls, if mask_headers is set to True then the headers will be obscured with ******.
	MaskHeaders *bool `json:"maskHeaders,omitempty" tf:"mask_headers,omitempty"`

	// The path to the page to run the check against. Will be combined with the host (specified within the MonitoredResource) and port to construct the full URL. If the provided path does not begin with "/", a "/" will be prepended automatically. Optional (defaults to "/").
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// The port to the page to run the check against. Will be combined with host (specified within the MonitoredResource) and path to construct the full URL. Optional (defaults to 80 without SSL, or 443 with SSL).
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// The HTTP request method to use for the check. If set to METHOD_UNSPECIFIED then requestMethod defaults to GET.
	// Default value is GET.
	// Possible values are: METHOD_UNSPECIFIED, GET, POST.
	RequestMethod *string `json:"requestMethod,omitempty" tf:"request_method,omitempty"`

	// If true, use HTTPS instead of HTTP to run the check.
	UseSSL *bool `json:"useSsl,omitempty" tf:"use_ssl,omitempty"`

	// Boolean specifying whether to include SSL certificate validation as a part of the Uptime check. Only applies to checks where monitoredResource is set to uptime_url. If useSsl is false, setting validateSsl to true has no effect.
	ValidateSSL *bool `json:"validateSsl,omitempty" tf:"validate_ssl,omitempty"`
}

type HTTPCheckObservation struct {

	// If present, the check will only pass if the HTTP response status code is in this set of status codes. If empty, the HTTP status code will only pass if the HTTP status code is 200-299.
	// Structure is documented below.
	AcceptedResponseStatusCodes []AcceptedResponseStatusCodesObservation `json:"acceptedResponseStatusCodes,omitempty" tf:"accepted_response_status_codes,omitempty"`

	// The authentication information. Optional when creating an HTTP check; defaults to empty.
	// Structure is documented below.
	AuthInfo []AuthInfoObservation `json:"authInfo,omitempty" tf:"auth_info,omitempty"`

	// The request body associated with the HTTP POST request. If contentType is URL_ENCODED, the body passed in must be URL-encoded. Users can provide a Content-Length header via the headers field or the API will do so. If the requestMethod is GET and body is not empty, the API will return an error. The maximum byte size is 1 megabyte. Note - As with all bytes fields JSON representations are base64 encoded. e.g. "foo=bar" in URL-encoded form is "foo%3Dbar" and in base64 encoding is "Zm9vJTI1M0RiYXI=".
	Body *string `json:"body,omitempty" tf:"body,omitempty"`

	// The content type to use for the check.
	// Possible values are: TYPE_UNSPECIFIED, URL_ENCODED.
	ContentType *string `json:"contentType,omitempty" tf:"content_type,omitempty"`

	// The list of headers to send as part of the uptime check request. If two headers have the same key and different values, they should be entered as a single header, with the value being a comma-separated list of all the desired values as described at https://www.w3.org/Protocols/rfc2616/rfc2616.txt (page 31). Entering two separate headers with the same key in a Create call will cause the first to be overwritten by the second. The maximum number of headers allowed is 100.
	Headers map[string]*string `json:"headers,omitempty" tf:"headers,omitempty"`

	// Boolean specifying whether to encrypt the header information. Encryption should be specified for any headers related to authentication that you do not wish to be seen when retrieving the configuration. The server will be responsible for encrypting the headers. On Get/List calls, if mask_headers is set to True then the headers will be obscured with ******.
	MaskHeaders *bool `json:"maskHeaders,omitempty" tf:"mask_headers,omitempty"`

	// The path to the page to run the check against. Will be combined with the host (specified within the MonitoredResource) and port to construct the full URL. If the provided path does not begin with "/", a "/" will be prepended automatically. Optional (defaults to "/").
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// The port to the page to run the check against. Will be combined with host (specified within the MonitoredResource) and path to construct the full URL. Optional (defaults to 80 without SSL, or 443 with SSL).
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// The HTTP request method to use for the check. If set to METHOD_UNSPECIFIED then requestMethod defaults to GET.
	// Default value is GET.
	// Possible values are: METHOD_UNSPECIFIED, GET, POST.
	RequestMethod *string `json:"requestMethod,omitempty" tf:"request_method,omitempty"`

	// If true, use HTTPS instead of HTTP to run the check.
	UseSSL *bool `json:"useSsl,omitempty" tf:"use_ssl,omitempty"`

	// Boolean specifying whether to include SSL certificate validation as a part of the Uptime check. Only applies to checks where monitoredResource is set to uptime_url. If useSsl is false, setting validateSsl to true has no effect.
	ValidateSSL *bool `json:"validateSsl,omitempty" tf:"validate_ssl,omitempty"`
}

type HTTPCheckParameters struct {

	// If present, the check will only pass if the HTTP response status code is in this set of status codes. If empty, the HTTP status code will only pass if the HTTP status code is 200-299.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	AcceptedResponseStatusCodes []AcceptedResponseStatusCodesParameters `json:"acceptedResponseStatusCodes,omitempty" tf:"accepted_response_status_codes,omitempty"`

	// The authentication information. Optional when creating an HTTP check; defaults to empty.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	AuthInfo []AuthInfoParameters `json:"authInfo,omitempty" tf:"auth_info,omitempty"`

	// The request body associated with the HTTP POST request. If contentType is URL_ENCODED, the body passed in must be URL-encoded. Users can provide a Content-Length header via the headers field or the API will do so. If the requestMethod is GET and body is not empty, the API will return an error. The maximum byte size is 1 megabyte. Note - As with all bytes fields JSON representations are base64 encoded. e.g. "foo=bar" in URL-encoded form is "foo%3Dbar" and in base64 encoding is "Zm9vJTI1M0RiYXI=".
	// +kubebuilder:validation:Optional
	Body *string `json:"body,omitempty" tf:"body,omitempty"`

	// The content type to use for the check.
	// Possible values are: TYPE_UNSPECIFIED, URL_ENCODED.
	// +kubebuilder:validation:Optional
	ContentType *string `json:"contentType,omitempty" tf:"content_type,omitempty"`

	// The list of headers to send as part of the uptime check request. If two headers have the same key and different values, they should be entered as a single header, with the value being a comma-separated list of all the desired values as described at https://www.w3.org/Protocols/rfc2616/rfc2616.txt (page 31). Entering two separate headers with the same key in a Create call will cause the first to be overwritten by the second. The maximum number of headers allowed is 100.
	// +kubebuilder:validation:Optional
	Headers map[string]*string `json:"headers,omitempty" tf:"headers,omitempty"`

	// Boolean specifying whether to encrypt the header information. Encryption should be specified for any headers related to authentication that you do not wish to be seen when retrieving the configuration. The server will be responsible for encrypting the headers. On Get/List calls, if mask_headers is set to True then the headers will be obscured with ******.
	// +kubebuilder:validation:Optional
	MaskHeaders *bool `json:"maskHeaders,omitempty" tf:"mask_headers,omitempty"`

	// The path to the page to run the check against. Will be combined with the host (specified within the MonitoredResource) and port to construct the full URL. If the provided path does not begin with "/", a "/" will be prepended automatically. Optional (defaults to "/").
	// +kubebuilder:validation:Optional
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// The port to the page to run the check against. Will be combined with host (specified within the MonitoredResource) and path to construct the full URL. Optional (defaults to 80 without SSL, or 443 with SSL).
	// +kubebuilder:validation:Optional
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// The HTTP request method to use for the check. If set to METHOD_UNSPECIFIED then requestMethod defaults to GET.
	// Default value is GET.
	// Possible values are: METHOD_UNSPECIFIED, GET, POST.
	// +kubebuilder:validation:Optional
	RequestMethod *string `json:"requestMethod,omitempty" tf:"request_method,omitempty"`

	// If true, use HTTPS instead of HTTP to run the check.
	// +kubebuilder:validation:Optional
	UseSSL *bool `json:"useSsl,omitempty" tf:"use_ssl,omitempty"`

	// Boolean specifying whether to include SSL certificate validation as a part of the Uptime check. Only applies to checks where monitoredResource is set to uptime_url. If useSsl is false, setting validateSsl to true has no effect.
	// +kubebuilder:validation:Optional
	ValidateSSL *bool `json:"validateSsl,omitempty" tf:"validate_ssl,omitempty"`
}

type JSONPathMatcherInitParameters struct {

	// Options to perform JSONPath content matching.
	// Default value is EXACT_MATCH.
	// Possible values are: EXACT_MATCH, REGEX_MATCH.
	JSONMatcher *string `json:"jsonMatcher,omitempty" tf:"json_matcher,omitempty"`

	// JSONPath within the response output pointing to the expected ContentMatcher::content to match against.
	JSONPath *string `json:"jsonPath,omitempty" tf:"json_path,omitempty"`
}

type JSONPathMatcherObservation struct {

	// Options to perform JSONPath content matching.
	// Default value is EXACT_MATCH.
	// Possible values are: EXACT_MATCH, REGEX_MATCH.
	JSONMatcher *string `json:"jsonMatcher,omitempty" tf:"json_matcher,omitempty"`

	// JSONPath within the response output pointing to the expected ContentMatcher::content to match against.
	JSONPath *string `json:"jsonPath,omitempty" tf:"json_path,omitempty"`
}

type JSONPathMatcherParameters struct {

	// Options to perform JSONPath content matching.
	// Default value is EXACT_MATCH.
	// Possible values are: EXACT_MATCH, REGEX_MATCH.
	// +kubebuilder:validation:Optional
	JSONMatcher *string `json:"jsonMatcher,omitempty" tf:"json_matcher,omitempty"`

	// JSONPath within the response output pointing to the expected ContentMatcher::content to match against.
	// +kubebuilder:validation:Optional
	JSONPath *string `json:"jsonPath,omitempty" tf:"json_path,omitempty"`
}

type MonitoredResourceInitParameters struct {

	// Values for all of the labels listed in the associated monitored resource descriptor. For example, Compute Engine VM instances use the labels "project_id", "instance_id", and "zone".
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// The monitored resource type. This field must match the type field of a MonitoredResourceDescriptor (https://cloud.google.com/monitoring/api/ref_v3/rest/v3/projects.monitoredResourceDescriptors#MonitoredResourceDescriptor) object. For example, the type of a Compute Engine VM instance is gce_instance. For a list of types, see Monitoring resource types (https://cloud.google.com/monitoring/api/resources) and Logging resource types (https://cloud.google.com/logging/docs/api/v2/resource-list).
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type MonitoredResourceObservation struct {

	// Values for all of the labels listed in the associated monitored resource descriptor. For example, Compute Engine VM instances use the labels "project_id", "instance_id", and "zone".
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// The monitored resource type. This field must match the type field of a MonitoredResourceDescriptor (https://cloud.google.com/monitoring/api/ref_v3/rest/v3/projects.monitoredResourceDescriptors#MonitoredResourceDescriptor) object. For example, the type of a Compute Engine VM instance is gce_instance. For a list of types, see Monitoring resource types (https://cloud.google.com/monitoring/api/resources) and Logging resource types (https://cloud.google.com/logging/docs/api/v2/resource-list).
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type MonitoredResourceParameters struct {

	// Values for all of the labels listed in the associated monitored resource descriptor. For example, Compute Engine VM instances use the labels "project_id", "instance_id", and "zone".
	// +kubebuilder:validation:Optional
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// The monitored resource type. This field must match the type field of a MonitoredResourceDescriptor (https://cloud.google.com/monitoring/api/ref_v3/rest/v3/projects.monitoredResourceDescriptors#MonitoredResourceDescriptor) object. For example, the type of a Compute Engine VM instance is gce_instance. For a list of types, see Monitoring resource types (https://cloud.google.com/monitoring/api/resources) and Logging resource types (https://cloud.google.com/logging/docs/api/v2/resource-list).
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type ResourceGroupInitParameters struct {

	// The resource type of the group members.
	// Possible values are: RESOURCE_TYPE_UNSPECIFIED, INSTANCE, AWS_ELB_LOAD_BALANCER.
	ResourceType *string `json:"resourceType,omitempty" tf:"resource_type,omitempty"`
}

type ResourceGroupObservation struct {

	// The group of resources being monitored. Should be the name of a group
	GroupID *string `json:"groupId,omitempty" tf:"group_id,omitempty"`

	// The resource type of the group members.
	// Possible values are: RESOURCE_TYPE_UNSPECIFIED, INSTANCE, AWS_ELB_LOAD_BALANCER.
	ResourceType *string `json:"resourceType,omitempty" tf:"resource_type,omitempty"`
}

type ResourceGroupParameters struct {

	// The group of resources being monitored. Should be the name of a group
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/monitoring/v1beta1.Group
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("name",true)
	// +kubebuilder:validation:Optional
	GroupID *string `json:"groupId,omitempty" tf:"group_id,omitempty"`

	// Reference to a Group in monitoring to populate groupId.
	// +kubebuilder:validation:Optional
	GroupIDRef *v1.Reference `json:"groupIdRef,omitempty" tf:"-"`

	// Selector for a Group in monitoring to populate groupId.
	// +kubebuilder:validation:Optional
	GroupIDSelector *v1.Selector `json:"groupIdSelector,omitempty" tf:"-"`

	// The resource type of the group members.
	// Possible values are: RESOURCE_TYPE_UNSPECIFIED, INSTANCE, AWS_ELB_LOAD_BALANCER.
	// +kubebuilder:validation:Optional
	ResourceType *string `json:"resourceType,omitempty" tf:"resource_type,omitempty"`
}

type TCPCheckInitParameters struct {

	// The port to the page to run the check against. Will be combined with host (specified within the MonitoredResource) to construct the full URL.
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`
}

type TCPCheckObservation struct {

	// The port to the page to run the check against. Will be combined with host (specified within the MonitoredResource) to construct the full URL.
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`
}

type TCPCheckParameters struct {

	// The port to the page to run the check against. Will be combined with host (specified within the MonitoredResource) to construct the full URL.
	// +kubebuilder:validation:Optional
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`
}

type UptimeCheckConfigInitParameters struct {

	// The checker type to use for the check. If the monitored resource type is servicedirectory_service, checkerType must be set to VPC_CHECKERS.
	// Possible values are: STATIC_IP_CHECKERS, VPC_CHECKERS.
	CheckerType *string `json:"checkerType,omitempty" tf:"checker_type,omitempty"`

	// The expected content on the page the check is run against. Currently, only the first entry in the list is supported, and other entries will be ignored. The server will look for an exact match of the string in the page response's content. This field is optional and should only be specified if a content match is required.
	// Structure is documented below.
	ContentMatchers []ContentMatchersInitParameters `json:"contentMatchers,omitempty" tf:"content_matchers,omitempty"`

	// A human-friendly name for the uptime check configuration. The display name should be unique within a Stackdriver Workspace in order to make it easier to identify; however, uniqueness is not enforced.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// Contains information needed to make an HTTP or HTTPS check.
	// Structure is documented below.
	HTTPCheck []HTTPCheckInitParameters `json:"httpCheck,omitempty" tf:"http_check,omitempty"`

	// The monitored resource (https://cloud.google.com/monitoring/api/resources) associated with the configuration. The following monitored resource types are supported for uptime checks:  uptime_url  gce_instance  gae_app  aws_ec2_instance aws_elb_load_balancer  k8s_service  servicedirectory_service
	// Structure is documented below.
	MonitoredResource []MonitoredResourceInitParameters `json:"monitoredResource,omitempty" tf:"monitored_resource,omitempty"`

	// How often, in seconds, the uptime check is performed. Currently, the only supported values are 60s (1 minute), 300s (5 minutes), 600s (10 minutes), and 900s (15 minutes). Optional, defaults to 300s.
	Period *string `json:"period,omitempty" tf:"period,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The group resource associated with the configuration.
	// Structure is documented below.
	ResourceGroup []ResourceGroupInitParameters `json:"resourceGroup,omitempty" tf:"resource_group,omitempty"`

	// The list of regions from which the check will be run. Some regions contain one location, and others contain more than one. If this field is specified, enough regions to include a minimum of 3 locations must be provided, or an error message is returned. Not specifying this field will result in uptime checks running from all regions.
	SelectedRegions []*string `json:"selectedRegions,omitempty" tf:"selected_regions,omitempty"`

	// Contains information needed to make a TCP check.
	// Structure is documented below.
	TCPCheck []TCPCheckInitParameters `json:"tcpCheck,omitempty" tf:"tcp_check,omitempty"`

	// The maximum amount of time to wait for the request to complete (must be between 1 and 60 seconds). Accepted formats https://developers.google.com/protocol-buffers/docs/reference/google.protobuf#google.protobuf.Duration
	Timeout *string `json:"timeout,omitempty" tf:"timeout,omitempty"`
}

type UptimeCheckConfigObservation struct {

	// The checker type to use for the check. If the monitored resource type is servicedirectory_service, checkerType must be set to VPC_CHECKERS.
	// Possible values are: STATIC_IP_CHECKERS, VPC_CHECKERS.
	CheckerType *string `json:"checkerType,omitempty" tf:"checker_type,omitempty"`

	// The expected content on the page the check is run against. Currently, only the first entry in the list is supported, and other entries will be ignored. The server will look for an exact match of the string in the page response's content. This field is optional and should only be specified if a content match is required.
	// Structure is documented below.
	ContentMatchers []ContentMatchersObservation `json:"contentMatchers,omitempty" tf:"content_matchers,omitempty"`

	// A human-friendly name for the uptime check configuration. The display name should be unique within a Stackdriver Workspace in order to make it easier to identify; however, uniqueness is not enforced.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// Contains information needed to make an HTTP or HTTPS check.
	// Structure is documented below.
	HTTPCheck []HTTPCheckObservation `json:"httpCheck,omitempty" tf:"http_check,omitempty"`

	// an identifier for the resource with format {{name}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The monitored resource (https://cloud.google.com/monitoring/api/resources) associated with the configuration. The following monitored resource types are supported for uptime checks:  uptime_url  gce_instance  gae_app  aws_ec2_instance aws_elb_load_balancer  k8s_service  servicedirectory_service
	// Structure is documented below.
	MonitoredResource []MonitoredResourceObservation `json:"monitoredResource,omitempty" tf:"monitored_resource,omitempty"`

	// A unique resource name for this UptimeCheckConfig. The format is projects/[PROJECT_ID]/uptimeCheckConfigs/[UPTIME_CHECK_ID].
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// How often, in seconds, the uptime check is performed. Currently, the only supported values are 60s (1 minute), 300s (5 minutes), 600s (10 minutes), and 900s (15 minutes). Optional, defaults to 300s.
	Period *string `json:"period,omitempty" tf:"period,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The group resource associated with the configuration.
	// Structure is documented below.
	ResourceGroup []ResourceGroupObservation `json:"resourceGroup,omitempty" tf:"resource_group,omitempty"`

	// The list of regions from which the check will be run. Some regions contain one location, and others contain more than one. If this field is specified, enough regions to include a minimum of 3 locations must be provided, or an error message is returned. Not specifying this field will result in uptime checks running from all regions.
	SelectedRegions []*string `json:"selectedRegions,omitempty" tf:"selected_regions,omitempty"`

	// Contains information needed to make a TCP check.
	// Structure is documented below.
	TCPCheck []TCPCheckObservation `json:"tcpCheck,omitempty" tf:"tcp_check,omitempty"`

	// The maximum amount of time to wait for the request to complete (must be between 1 and 60 seconds). Accepted formats https://developers.google.com/protocol-buffers/docs/reference/google.protobuf#google.protobuf.Duration
	Timeout *string `json:"timeout,omitempty" tf:"timeout,omitempty"`

	// The id of the uptime check
	UptimeCheckID *string `json:"uptimeCheckId,omitempty" tf:"uptime_check_id,omitempty"`
}

type UptimeCheckConfigParameters struct {

	// The checker type to use for the check. If the monitored resource type is servicedirectory_service, checkerType must be set to VPC_CHECKERS.
	// Possible values are: STATIC_IP_CHECKERS, VPC_CHECKERS.
	// +kubebuilder:validation:Optional
	CheckerType *string `json:"checkerType,omitempty" tf:"checker_type,omitempty"`

	// The expected content on the page the check is run against. Currently, only the first entry in the list is supported, and other entries will be ignored. The server will look for an exact match of the string in the page response's content. This field is optional and should only be specified if a content match is required.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	ContentMatchers []ContentMatchersParameters `json:"contentMatchers,omitempty" tf:"content_matchers,omitempty"`

	// A human-friendly name for the uptime check configuration. The display name should be unique within a Stackdriver Workspace in order to make it easier to identify; however, uniqueness is not enforced.
	// +kubebuilder:validation:Optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// Contains information needed to make an HTTP or HTTPS check.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	HTTPCheck []HTTPCheckParameters `json:"httpCheck,omitempty" tf:"http_check,omitempty"`

	// The monitored resource (https://cloud.google.com/monitoring/api/resources) associated with the configuration. The following monitored resource types are supported for uptime checks:  uptime_url  gce_instance  gae_app  aws_ec2_instance aws_elb_load_balancer  k8s_service  servicedirectory_service
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	MonitoredResource []MonitoredResourceParameters `json:"monitoredResource,omitempty" tf:"monitored_resource,omitempty"`

	// How often, in seconds, the uptime check is performed. Currently, the only supported values are 60s (1 minute), 300s (5 minutes), 600s (10 minutes), and 900s (15 minutes). Optional, defaults to 300s.
	// +kubebuilder:validation:Optional
	Period *string `json:"period,omitempty" tf:"period,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The group resource associated with the configuration.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	ResourceGroup []ResourceGroupParameters `json:"resourceGroup,omitempty" tf:"resource_group,omitempty"`

	// The list of regions from which the check will be run. Some regions contain one location, and others contain more than one. If this field is specified, enough regions to include a minimum of 3 locations must be provided, or an error message is returned. Not specifying this field will result in uptime checks running from all regions.
	// +kubebuilder:validation:Optional
	SelectedRegions []*string `json:"selectedRegions,omitempty" tf:"selected_regions,omitempty"`

	// Contains information needed to make a TCP check.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	TCPCheck []TCPCheckParameters `json:"tcpCheck,omitempty" tf:"tcp_check,omitempty"`

	// The maximum amount of time to wait for the request to complete (must be between 1 and 60 seconds). Accepted formats https://developers.google.com/protocol-buffers/docs/reference/google.protobuf#google.protobuf.Duration
	// +kubebuilder:validation:Optional
	Timeout *string `json:"timeout,omitempty" tf:"timeout,omitempty"`
}

// UptimeCheckConfigSpec defines the desired state of UptimeCheckConfig
type UptimeCheckConfigSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     UptimeCheckConfigParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	InitProvider UptimeCheckConfigInitParameters `json:"initProvider,omitempty"`
}

// UptimeCheckConfigStatus defines the observed state of UptimeCheckConfig.
type UptimeCheckConfigStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        UptimeCheckConfigObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// UptimeCheckConfig is the Schema for the UptimeCheckConfigs API. This message configures which resources and services to monitor for availability.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type UptimeCheckConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.displayName) || has(self.initProvider.displayName)",message="displayName is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.timeout) || has(self.initProvider.timeout)",message="timeout is a required parameter"
	Spec   UptimeCheckConfigSpec   `json:"spec"`
	Status UptimeCheckConfigStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// UptimeCheckConfigList contains a list of UptimeCheckConfigs
type UptimeCheckConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []UptimeCheckConfig `json:"items"`
}

// Repository type metadata.
var (
	UptimeCheckConfig_Kind             = "UptimeCheckConfig"
	UptimeCheckConfig_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: UptimeCheckConfig_Kind}.String()
	UptimeCheckConfig_KindAPIVersion   = UptimeCheckConfig_Kind + "." + CRDGroupVersion.String()
	UptimeCheckConfig_GroupVersionKind = CRDGroupVersion.WithKind(UptimeCheckConfig_Kind)
)

func init() {
	SchemeBuilder.Register(&UptimeCheckConfig{}, &UptimeCheckConfigList{})
}
