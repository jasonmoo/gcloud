// Package logging provides access to the Google Cloud Logging API.
//
// Usage example:
//
//   import "google.golang.org/api/logging/v1beta3"
//   ...
//   loggingService, err := logging.New(oauthHttpClient)
package logging

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"google.golang.org/api/googleapi"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// Always reference these packages, just in case the auto-generated code
// below doesn't.
var _ = bytes.NewBuffer
var _ = strconv.Itoa
var _ = fmt.Sprintf
var _ = json.NewDecoder
var _ = io.Copy
var _ = url.Parse
var _ = googleapi.Version
var _ = errors.New
var _ = strings.Replace

const apiId = "logging:v1beta3"
const apiName = "logging"
const apiVersion = "v1beta3"
const basePath = "https://logging.googleapis.com/v1beta3/"

// OAuth2 scopes used by this API.
const (
	// View and manage your data across Google Cloud Platform services
	CloudPlatformScope = "https://www.googleapis.com/auth/cloud-platform"
)

func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client, BasePath: basePath}
	s.Projects = NewProjectsService(s)
	return s, nil
}

type Service struct {
	client   *http.Client
	BasePath string // API endpoint base URL

	Projects *ProjectsService
}

func NewProjectsService(s *Service) *ProjectsService {
	rs := &ProjectsService{s: s}
	rs.LogEntries = NewProjectsLogEntriesService(s)
	rs.LogServices = NewProjectsLogServicesService(s)
	rs.Logs = NewProjectsLogsService(s)
	rs.Metrics = NewProjectsMetricsService(s)
	return rs
}

type ProjectsService struct {
	s *Service

	LogEntries *ProjectsLogEntriesService

	LogServices *ProjectsLogServicesService

	Logs *ProjectsLogsService

	Metrics *ProjectsMetricsService
}

func NewProjectsLogEntriesService(s *Service) *ProjectsLogEntriesService {
	rs := &ProjectsLogEntriesService{s: s}
	return rs
}

type ProjectsLogEntriesService struct {
	s *Service
}

func NewProjectsLogServicesService(s *Service) *ProjectsLogServicesService {
	rs := &ProjectsLogServicesService{s: s}
	rs.Indexes = NewProjectsLogServicesIndexesService(s)
	rs.Sinks = NewProjectsLogServicesSinksService(s)
	return rs
}

type ProjectsLogServicesService struct {
	s *Service

	Indexes *ProjectsLogServicesIndexesService

	Sinks *ProjectsLogServicesSinksService
}

func NewProjectsLogServicesIndexesService(s *Service) *ProjectsLogServicesIndexesService {
	rs := &ProjectsLogServicesIndexesService{s: s}
	return rs
}

type ProjectsLogServicesIndexesService struct {
	s *Service
}

func NewProjectsLogServicesSinksService(s *Service) *ProjectsLogServicesSinksService {
	rs := &ProjectsLogServicesSinksService{s: s}
	return rs
}

type ProjectsLogServicesSinksService struct {
	s *Service
}

func NewProjectsLogsService(s *Service) *ProjectsLogsService {
	rs := &ProjectsLogsService{s: s}
	rs.Entries = NewProjectsLogsEntriesService(s)
	rs.Sinks = NewProjectsLogsSinksService(s)
	return rs
}

type ProjectsLogsService struct {
	s *Service

	Entries *ProjectsLogsEntriesService

	Sinks *ProjectsLogsSinksService
}

func NewProjectsLogsEntriesService(s *Service) *ProjectsLogsEntriesService {
	rs := &ProjectsLogsEntriesService{s: s}
	return rs
}

type ProjectsLogsEntriesService struct {
	s *Service
}

func NewProjectsLogsSinksService(s *Service) *ProjectsLogsSinksService {
	rs := &ProjectsLogsSinksService{s: s}
	return rs
}

type ProjectsLogsSinksService struct {
	s *Service
}

func NewProjectsMetricsService(s *Service) *ProjectsMetricsService {
	rs := &ProjectsMetricsService{s: s}
	return rs
}

type ProjectsMetricsService struct {
	s *Service
}

type Empty struct {
}

type ListLogEntriesResponse struct {
	Entries []*LogEntry `json:"entries,omitempty"`

	NextPageToken string `json:"nextPageToken,omitempty"`
}

type ListLogMetricsResponse struct {
	Metrics []*LogMetric `json:"metrics,omitempty"`

	NextPageToken string `json:"nextPageToken,omitempty"`
}

type ListLogServiceIndexesResponse struct {
	NextPageToken string `json:"nextPageToken,omitempty"`

	ServiceIndexPrefixes []string `json:"serviceIndexPrefixes,omitempty"`
}

type ListLogServiceSinksResponse struct {
	Sinks []*LogSink `json:"sinks,omitempty"`
}

type ListLogServicesResponse struct {
	LogServices []*LogService `json:"logServices,omitempty"`

	NextPageToken string `json:"nextPageToken,omitempty"`
}

type ListLogSinksResponse struct {
	Sinks []*LogSink `json:"sinks,omitempty"`
}

type ListLogsResponse struct {
	Logs []*Log `json:"logs,omitempty"`

	NextPageToken string `json:"nextPageToken,omitempty"`
}

type Log struct {
	DisplayName string `json:"displayName,omitempty"`

	Name string `json:"name,omitempty"`

	PayloadType string `json:"payloadType,omitempty"`
}

type LogEntry struct {
	InsertId string `json:"insertId,omitempty"`

	Log string `json:"log,omitempty"`

	Metadata *LogEntryMetadata `json:"metadata,omitempty"`

	ProtoPayload *LogEntryProtoPayload `json:"protoPayload,omitempty"`

	TextPayload string `json:"textPayload,omitempty"`
}

type LogEntryProtoPayload struct {
}

type LogEntryMetadata struct {
	Labels map[string]string `json:"labels,omitempty"`

	ProjectId string `json:"projectId,omitempty"`

	ProjectNumber int64 `json:"projectNumber,omitempty,string"`

	Region string `json:"region,omitempty"`

	ServiceName string `json:"serviceName,omitempty"`

	Severity string `json:"severity,omitempty"`

	TimeNanos int64 `json:"timeNanos,omitempty,string"`

	Timestamp string `json:"timestamp,omitempty"`

	UserId string `json:"userId,omitempty"`

	Zone string `json:"zone,omitempty"`
}

type LogError struct {
	Resource string `json:"resource,omitempty"`

	Status *Status `json:"status,omitempty"`

	TimeNanos int64 `json:"timeNanos,omitempty,string"`
}

type LogMetric struct {
	Description string `json:"description,omitempty"`

	Filter string `json:"filter,omitempty"`

	Name string `json:"name,omitempty"`
}

type LogService struct {
	IndexKeys []string `json:"indexKeys,omitempty"`

	Name string `json:"name,omitempty"`
}

type LogSink struct {
	Destination string `json:"destination,omitempty"`

	Errors []*LogError `json:"errors,omitempty"`

	Name string `json:"name,omitempty"`
}

type Status struct {
	Code int64 `json:"code,omitempty"`

	Details []*StatusDetails `json:"details,omitempty"`

	Message string `json:"message,omitempty"`
}

type StatusDetails struct {
}

type WriteLogEntriesRequest struct {
	CommonLabels map[string]string `json:"commonLabels,omitempty"`

	Entries []*LogEntry `json:"entries,omitempty"`
}

type WriteLogEntriesResponse struct {
}

// method id "logging.projects.logEntries.list":

type ProjectsLogEntriesListCall struct {
	s         *Service
	projectId string
	opt_      map[string]interface{}
}

// List: Lists log entries in the specified project. Requires
// https://www.googleapis.com/auth/logging.read scope.
func (r *ProjectsLogEntriesService) List(projectId string) *ProjectsLogEntriesListCall {
	c := &ProjectsLogEntriesListCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	return c
}

// Filter sets the optional parameter "filter":
func (c *ProjectsLogEntriesListCall) Filter(filter string) *ProjectsLogEntriesListCall {
	c.opt_["filter"] = filter
	return c
}

// OrderBy sets the optional parameter "orderBy":
func (c *ProjectsLogEntriesListCall) OrderBy(orderBy string) *ProjectsLogEntriesListCall {
	c.opt_["orderBy"] = orderBy
	return c
}

// PageSize sets the optional parameter "pageSize":
func (c *ProjectsLogEntriesListCall) PageSize(pageSize int64) *ProjectsLogEntriesListCall {
	c.opt_["pageSize"] = pageSize
	return c
}

// PageToken sets the optional parameter "pageToken":
func (c *ProjectsLogEntriesListCall) PageToken(pageToken string) *ProjectsLogEntriesListCall {
	c.opt_["pageToken"] = pageToken
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsLogEntriesListCall) Fields(s ...googleapi.Field) *ProjectsLogEntriesListCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsLogEntriesListCall) Do() (*ListLogEntriesResponse, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["filter"]; ok {
		params.Set("filter", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["orderBy"]; ok {
		params.Set("orderBy", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["pageSize"]; ok {
		params.Set("pageSize", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["pageToken"]; ok {
		params.Set("pageToken", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "{+projectId}/logEntries")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectId": c.projectId,
	})
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *ListLogEntriesResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Lists log entries in the specified project. Requires https://www.googleapis.com/auth/logging.read scope.",
	//   "httpMethod": "GET",
	//   "id": "logging.projects.logEntries.list",
	//   "parameterOrder": [
	//     "projectId"
	//   ],
	//   "parameters": {
	//     "filter": {
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "orderBy": {
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "pageSize": {
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "projectId": {
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{+projectId}/logEntries",
	//   "response": {
	//     "$ref": "ListLogEntriesResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "logging.projects.logServices.list":

type ProjectsLogServicesListCall struct {
	s         *Service
	projectId string
	opt_      map[string]interface{}
}

// List: Lists log services associated with log entries ingested for a
// project. Requires https://www.googleapis.com/auth/logging.read scope.
func (r *ProjectsLogServicesService) List(projectId string) *ProjectsLogServicesListCall {
	c := &ProjectsLogServicesListCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	return c
}

// Log sets the optional parameter "log":
func (c *ProjectsLogServicesListCall) Log(log string) *ProjectsLogServicesListCall {
	c.opt_["log"] = log
	return c
}

// PageSize sets the optional parameter "pageSize":
func (c *ProjectsLogServicesListCall) PageSize(pageSize int64) *ProjectsLogServicesListCall {
	c.opt_["pageSize"] = pageSize
	return c
}

// PageToken sets the optional parameter "pageToken":
func (c *ProjectsLogServicesListCall) PageToken(pageToken string) *ProjectsLogServicesListCall {
	c.opt_["pageToken"] = pageToken
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsLogServicesListCall) Fields(s ...googleapi.Field) *ProjectsLogServicesListCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsLogServicesListCall) Do() (*ListLogServicesResponse, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["log"]; ok {
		params.Set("log", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["pageSize"]; ok {
		params.Set("pageSize", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["pageToken"]; ok {
		params.Set("pageToken", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "{+projectId}/logServices")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectId": c.projectId,
	})
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *ListLogServicesResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Lists log services associated with log entries ingested for a project. Requires https://www.googleapis.com/auth/logging.read scope.",
	//   "httpMethod": "GET",
	//   "id": "logging.projects.logServices.list",
	//   "parameterOrder": [
	//     "projectId"
	//   ],
	//   "parameters": {
	//     "log": {
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "pageSize": {
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "projectId": {
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{+projectId}/logServices",
	//   "response": {
	//     "$ref": "ListLogServicesResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "logging.projects.logServices.indexes.list":

type ProjectsLogServicesIndexesListCall struct {
	s       *Service
	service string
	opt_    map[string]interface{}
}

// List: Lists log service indexes associated with a log service.
// Requires https://www.googleapis.com/auth/logging.read scope.
func (r *ProjectsLogServicesIndexesService) List(service string) *ProjectsLogServicesIndexesListCall {
	c := &ProjectsLogServicesIndexesListCall{s: r.s, opt_: make(map[string]interface{})}
	c.service = service
	return c
}

// Depth sets the optional parameter "depth":
func (c *ProjectsLogServicesIndexesListCall) Depth(depth int64) *ProjectsLogServicesIndexesListCall {
	c.opt_["depth"] = depth
	return c
}

// IndexPrefix sets the optional parameter "indexPrefix":
func (c *ProjectsLogServicesIndexesListCall) IndexPrefix(indexPrefix string) *ProjectsLogServicesIndexesListCall {
	c.opt_["indexPrefix"] = indexPrefix
	return c
}

// Log sets the optional parameter "log":
func (c *ProjectsLogServicesIndexesListCall) Log(log string) *ProjectsLogServicesIndexesListCall {
	c.opt_["log"] = log
	return c
}

// PageSize sets the optional parameter "pageSize":
func (c *ProjectsLogServicesIndexesListCall) PageSize(pageSize int64) *ProjectsLogServicesIndexesListCall {
	c.opt_["pageSize"] = pageSize
	return c
}

// PageToken sets the optional parameter "pageToken":
func (c *ProjectsLogServicesIndexesListCall) PageToken(pageToken string) *ProjectsLogServicesIndexesListCall {
	c.opt_["pageToken"] = pageToken
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsLogServicesIndexesListCall) Fields(s ...googleapi.Field) *ProjectsLogServicesIndexesListCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsLogServicesIndexesListCall) Do() (*ListLogServiceIndexesResponse, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["depth"]; ok {
		params.Set("depth", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["indexPrefix"]; ok {
		params.Set("indexPrefix", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["log"]; ok {
		params.Set("log", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["pageSize"]; ok {
		params.Set("pageSize", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["pageToken"]; ok {
		params.Set("pageToken", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "{+service}/indexes")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"service": c.service,
	})
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *ListLogServiceIndexesResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Lists log service indexes associated with a log service. Requires https://www.googleapis.com/auth/logging.read scope.",
	//   "httpMethod": "GET",
	//   "id": "logging.projects.logServices.indexes.list",
	//   "parameterOrder": [
	//     "service"
	//   ],
	//   "parameters": {
	//     "depth": {
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "indexPrefix": {
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "log": {
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "pageSize": {
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "service": {
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{+service}/indexes",
	//   "response": {
	//     "$ref": "ListLogServiceIndexesResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "logging.projects.logServices.sinks.delete":

type ProjectsLogServicesSinksDeleteCall struct {
	s    *Service
	sink string
	opt_ map[string]interface{}
}

// Delete: Deletes the specified log service sink. Requires
// https://www.googleapis.com/auth/logging.admin scope.
func (r *ProjectsLogServicesSinksService) Delete(sink string) *ProjectsLogServicesSinksDeleteCall {
	c := &ProjectsLogServicesSinksDeleteCall{s: r.s, opt_: make(map[string]interface{})}
	c.sink = sink
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsLogServicesSinksDeleteCall) Fields(s ...googleapi.Field) *ProjectsLogServicesSinksDeleteCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsLogServicesSinksDeleteCall) Do() (*Empty, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "{+sink}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"sink": c.sink,
	})
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Empty
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Deletes the specified log service sink. Requires https://www.googleapis.com/auth/logging.admin scope.",
	//   "httpMethod": "DELETE",
	//   "id": "logging.projects.logServices.sinks.delete",
	//   "parameterOrder": [
	//     "sink"
	//   ],
	//   "parameters": {
	//     "sink": {
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{+sink}",
	//   "response": {
	//     "$ref": "Empty"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "logging.projects.logServices.sinks.get":

type ProjectsLogServicesSinksGetCall struct {
	s    *Service
	sink string
	opt_ map[string]interface{}
}

// Get: Get the specified log service sink resource. Requires
// https://www.googleapis.com/auth/logging.read scope.
func (r *ProjectsLogServicesSinksService) Get(sink string) *ProjectsLogServicesSinksGetCall {
	c := &ProjectsLogServicesSinksGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.sink = sink
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsLogServicesSinksGetCall) Fields(s ...googleapi.Field) *ProjectsLogServicesSinksGetCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsLogServicesSinksGetCall) Do() (*LogSink, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "{+sink}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"sink": c.sink,
	})
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *LogSink
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Get the specified log service sink resource. Requires https://www.googleapis.com/auth/logging.read scope.",
	//   "httpMethod": "GET",
	//   "id": "logging.projects.logServices.sinks.get",
	//   "parameterOrder": [
	//     "sink"
	//   ],
	//   "parameters": {
	//     "sink": {
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{+sink}",
	//   "response": {
	//     "$ref": "LogSink"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "logging.projects.logServices.sinks.list":

type ProjectsLogServicesSinksListCall struct {
	s       *Service
	service string
	opt_    map[string]interface{}
}

// List: List log service sinks associated with the specified service.
// Requires https://www.googleapis.com/auth/logging.read scope.
func (r *ProjectsLogServicesSinksService) List(service string) *ProjectsLogServicesSinksListCall {
	c := &ProjectsLogServicesSinksListCall{s: r.s, opt_: make(map[string]interface{})}
	c.service = service
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsLogServicesSinksListCall) Fields(s ...googleapi.Field) *ProjectsLogServicesSinksListCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsLogServicesSinksListCall) Do() (*ListLogServiceSinksResponse, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "{+service}/sinks")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"service": c.service,
	})
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *ListLogServiceSinksResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "List log service sinks associated with the specified service. Requires https://www.googleapis.com/auth/logging.read scope.",
	//   "httpMethod": "GET",
	//   "id": "logging.projects.logServices.sinks.list",
	//   "parameterOrder": [
	//     "service"
	//   ],
	//   "parameters": {
	//     "service": {
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{+service}/sinks",
	//   "response": {
	//     "$ref": "ListLogServiceSinksResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "logging.projects.logServices.sinks.patch":

type ProjectsLogServicesSinksPatchCall struct {
	s        *Service
	sinkName string
	sink     string
	logsink  *LogSink
	opt_     map[string]interface{}
}

// Patch: Create or update the specified log service sink resource.
// Requires https://www.googleapis.com/auth/logging.admin scope. This
// method supports patch semantics.
func (r *ProjectsLogServicesSinksService) Patch(sinkName string, sink string, logsink *LogSink) *ProjectsLogServicesSinksPatchCall {
	c := &ProjectsLogServicesSinksPatchCall{s: r.s, opt_: make(map[string]interface{})}
	c.sinkName = sinkName
	c.sink = sink
	c.logsink = logsink
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsLogServicesSinksPatchCall) Fields(s ...googleapi.Field) *ProjectsLogServicesSinksPatchCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsLogServicesSinksPatchCall) Do() (*LogSink, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.logsink)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	params.Set("sink", fmt.Sprintf("%v", c.sink))
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "{+sinkName}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("PATCH", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"sinkName": c.sinkName,
	})
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *LogSink
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Create or update the specified log service sink resource. Requires https://www.googleapis.com/auth/logging.admin scope. This method supports patch semantics.",
	//   "httpMethod": "PATCH",
	//   "id": "logging.projects.logServices.sinks.patch",
	//   "parameterOrder": [
	//     "sinkName",
	//     "sink"
	//   ],
	//   "parameters": {
	//     "sink": {
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "sinkName": {
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{+sinkName}",
	//   "request": {
	//     "$ref": "LogSink"
	//   },
	//   "response": {
	//     "$ref": "LogSink"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "logging.projects.logServices.sinks.update":

type ProjectsLogServicesSinksUpdateCall struct {
	s        *Service
	sinkName string
	logsink  *LogSink
	opt_     map[string]interface{}
}

// Update: Create or update the specified log service sink resource.
// Requires https://www.googleapis.com/auth/logging.admin scope.
func (r *ProjectsLogServicesSinksService) Update(sinkName string, logsink *LogSink) *ProjectsLogServicesSinksUpdateCall {
	c := &ProjectsLogServicesSinksUpdateCall{s: r.s, opt_: make(map[string]interface{})}
	c.sinkName = sinkName
	c.logsink = logsink
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsLogServicesSinksUpdateCall) Fields(s ...googleapi.Field) *ProjectsLogServicesSinksUpdateCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsLogServicesSinksUpdateCall) Do() (*LogSink, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.logsink)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "{+sinkName}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("PUT", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"sinkName": c.sinkName,
	})
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *LogSink
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Create or update the specified log service sink resource. Requires https://www.googleapis.com/auth/logging.admin scope.",
	//   "httpMethod": "PUT",
	//   "id": "logging.projects.logServices.sinks.update",
	//   "parameterOrder": [
	//     "sinkName"
	//   ],
	//   "parameters": {
	//     "sinkName": {
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{+sinkName}",
	//   "request": {
	//     "$ref": "LogSink"
	//   },
	//   "response": {
	//     "$ref": "LogSink"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "logging.projects.logs.delete":

type ProjectsLogsDeleteCall struct {
	s    *Service
	log  string
	opt_ map[string]interface{}
}

// Delete: Deletes the specified log resource and all log entries
// contained in it. Requires
// https://www.googleapis.com/auth/logging.admin scope.
func (r *ProjectsLogsService) Delete(log string) *ProjectsLogsDeleteCall {
	c := &ProjectsLogsDeleteCall{s: r.s, opt_: make(map[string]interface{})}
	c.log = log
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsLogsDeleteCall) Fields(s ...googleapi.Field) *ProjectsLogsDeleteCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsLogsDeleteCall) Do() (*Empty, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "{+log}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"log": c.log,
	})
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Empty
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Deletes the specified log resource and all log entries contained in it. Requires https://www.googleapis.com/auth/logging.admin scope.",
	//   "httpMethod": "DELETE",
	//   "id": "logging.projects.logs.delete",
	//   "parameterOrder": [
	//     "log"
	//   ],
	//   "parameters": {
	//     "log": {
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{+log}",
	//   "response": {
	//     "$ref": "Empty"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "logging.projects.logs.list":

type ProjectsLogsListCall struct {
	s         *Service
	projectId string
	opt_      map[string]interface{}
}

// List: Lists log resources belonging to the specified project.
// Requires https://www.googleapis.com/auth/logging.read scope.
func (r *ProjectsLogsService) List(projectId string) *ProjectsLogsListCall {
	c := &ProjectsLogsListCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	return c
}

// PageSize sets the optional parameter "pageSize":
func (c *ProjectsLogsListCall) PageSize(pageSize int64) *ProjectsLogsListCall {
	c.opt_["pageSize"] = pageSize
	return c
}

// PageToken sets the optional parameter "pageToken":
func (c *ProjectsLogsListCall) PageToken(pageToken string) *ProjectsLogsListCall {
	c.opt_["pageToken"] = pageToken
	return c
}

// ServiceIndexPrefix sets the optional parameter "serviceIndexPrefix":
func (c *ProjectsLogsListCall) ServiceIndexPrefix(serviceIndexPrefix string) *ProjectsLogsListCall {
	c.opt_["serviceIndexPrefix"] = serviceIndexPrefix
	return c
}

// ServiceName sets the optional parameter "serviceName":
func (c *ProjectsLogsListCall) ServiceName(serviceName string) *ProjectsLogsListCall {
	c.opt_["serviceName"] = serviceName
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsLogsListCall) Fields(s ...googleapi.Field) *ProjectsLogsListCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsLogsListCall) Do() (*ListLogsResponse, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["pageSize"]; ok {
		params.Set("pageSize", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["pageToken"]; ok {
		params.Set("pageToken", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["serviceIndexPrefix"]; ok {
		params.Set("serviceIndexPrefix", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["serviceName"]; ok {
		params.Set("serviceName", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "{+projectId}/logs")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectId": c.projectId,
	})
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *ListLogsResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Lists log resources belonging to the specified project. Requires https://www.googleapis.com/auth/logging.read scope.",
	//   "httpMethod": "GET",
	//   "id": "logging.projects.logs.list",
	//   "parameterOrder": [
	//     "projectId"
	//   ],
	//   "parameters": {
	//     "pageSize": {
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "projectId": {
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "serviceIndexPrefix": {
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "serviceName": {
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "{+projectId}/logs",
	//   "response": {
	//     "$ref": "ListLogsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "logging.projects.logs.entries.write":

type ProjectsLogsEntriesWriteCall struct {
	s                      *Service
	log                    string
	writelogentriesrequest *WriteLogEntriesRequest
	opt_                   map[string]interface{}
}

// Write: Creates several log entries in a log. Requires
// https://www.googleapis.com/auth/logging.write scope.
func (r *ProjectsLogsEntriesService) Write(log string, writelogentriesrequest *WriteLogEntriesRequest) *ProjectsLogsEntriesWriteCall {
	c := &ProjectsLogsEntriesWriteCall{s: r.s, opt_: make(map[string]interface{})}
	c.log = log
	c.writelogentriesrequest = writelogentriesrequest
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsLogsEntriesWriteCall) Fields(s ...googleapi.Field) *ProjectsLogsEntriesWriteCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsLogsEntriesWriteCall) Do() (*WriteLogEntriesResponse, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.writelogentriesrequest)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "{+log}/entries:write")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"log": c.log,
	})
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *WriteLogEntriesResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Creates several log entries in a log. Requires https://www.googleapis.com/auth/logging.write scope.",
	//   "httpMethod": "POST",
	//   "id": "logging.projects.logs.entries.write",
	//   "parameterOrder": [
	//     "log"
	//   ],
	//   "parameters": {
	//     "log": {
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{+log}/entries:write",
	//   "request": {
	//     "$ref": "WriteLogEntriesRequest"
	//   },
	//   "response": {
	//     "$ref": "WriteLogEntriesResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "logging.projects.logs.sinks.delete":

type ProjectsLogsSinksDeleteCall struct {
	s    *Service
	sink string
	opt_ map[string]interface{}
}

// Delete: Deletes the specified log sink. Requires
// https://www.googleapis.com/auth/logging.admin scope.
func (r *ProjectsLogsSinksService) Delete(sink string) *ProjectsLogsSinksDeleteCall {
	c := &ProjectsLogsSinksDeleteCall{s: r.s, opt_: make(map[string]interface{})}
	c.sink = sink
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsLogsSinksDeleteCall) Fields(s ...googleapi.Field) *ProjectsLogsSinksDeleteCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsLogsSinksDeleteCall) Do() (*Empty, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "{+sink}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"sink": c.sink,
	})
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Empty
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Deletes the specified log sink. Requires https://www.googleapis.com/auth/logging.admin scope.",
	//   "httpMethod": "DELETE",
	//   "id": "logging.projects.logs.sinks.delete",
	//   "parameterOrder": [
	//     "sink"
	//   ],
	//   "parameters": {
	//     "sink": {
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{+sink}",
	//   "response": {
	//     "$ref": "Empty"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "logging.projects.logs.sinks.get":

type ProjectsLogsSinksGetCall struct {
	s    *Service
	sink string
	opt_ map[string]interface{}
}

// Get: Get the specified log sink resource. Requires
// https://www.googleapis.com/auth/logging.read scope.
func (r *ProjectsLogsSinksService) Get(sink string) *ProjectsLogsSinksGetCall {
	c := &ProjectsLogsSinksGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.sink = sink
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsLogsSinksGetCall) Fields(s ...googleapi.Field) *ProjectsLogsSinksGetCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsLogsSinksGetCall) Do() (*LogSink, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "{+sink}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"sink": c.sink,
	})
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *LogSink
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Get the specified log sink resource. Requires https://www.googleapis.com/auth/logging.read scope.",
	//   "httpMethod": "GET",
	//   "id": "logging.projects.logs.sinks.get",
	//   "parameterOrder": [
	//     "sink"
	//   ],
	//   "parameters": {
	//     "sink": {
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{+sink}",
	//   "response": {
	//     "$ref": "LogSink"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "logging.projects.logs.sinks.list":

type ProjectsLogsSinksListCall struct {
	s    *Service
	log  string
	opt_ map[string]interface{}
}

// List: List log sinks associated with the specified log. Requires
// https://www.googleapis.com/auth/logging.read scope.
func (r *ProjectsLogsSinksService) List(log string) *ProjectsLogsSinksListCall {
	c := &ProjectsLogsSinksListCall{s: r.s, opt_: make(map[string]interface{})}
	c.log = log
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsLogsSinksListCall) Fields(s ...googleapi.Field) *ProjectsLogsSinksListCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsLogsSinksListCall) Do() (*ListLogSinksResponse, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "{+log}/sinks")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"log": c.log,
	})
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *ListLogSinksResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "List log sinks associated with the specified log. Requires https://www.googleapis.com/auth/logging.read scope.",
	//   "httpMethod": "GET",
	//   "id": "logging.projects.logs.sinks.list",
	//   "parameterOrder": [
	//     "log"
	//   ],
	//   "parameters": {
	//     "log": {
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{+log}/sinks",
	//   "response": {
	//     "$ref": "ListLogSinksResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "logging.projects.logs.sinks.patch":

type ProjectsLogsSinksPatchCall struct {
	s        *Service
	sinkName string
	sink     string
	logsink  *LogSink
	opt_     map[string]interface{}
}

// Patch: Create or update the specified log sink resource. Requires
// https://www.googleapis.com/auth/logging.admin scope. This method
// supports patch semantics.
func (r *ProjectsLogsSinksService) Patch(sinkName string, sink string, logsink *LogSink) *ProjectsLogsSinksPatchCall {
	c := &ProjectsLogsSinksPatchCall{s: r.s, opt_: make(map[string]interface{})}
	c.sinkName = sinkName
	c.sink = sink
	c.logsink = logsink
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsLogsSinksPatchCall) Fields(s ...googleapi.Field) *ProjectsLogsSinksPatchCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsLogsSinksPatchCall) Do() (*LogSink, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.logsink)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	params.Set("sink", fmt.Sprintf("%v", c.sink))
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "{+sinkName}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("PATCH", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"sinkName": c.sinkName,
	})
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *LogSink
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Create or update the specified log sink resource. Requires https://www.googleapis.com/auth/logging.admin scope. This method supports patch semantics.",
	//   "httpMethod": "PATCH",
	//   "id": "logging.projects.logs.sinks.patch",
	//   "parameterOrder": [
	//     "sinkName",
	//     "sink"
	//   ],
	//   "parameters": {
	//     "sink": {
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "sinkName": {
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{+sinkName}",
	//   "request": {
	//     "$ref": "LogSink"
	//   },
	//   "response": {
	//     "$ref": "LogSink"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "logging.projects.logs.sinks.update":

type ProjectsLogsSinksUpdateCall struct {
	s        *Service
	sinkName string
	logsink  *LogSink
	opt_     map[string]interface{}
}

// Update: Create or update the specified log sink resource. Requires
// https://www.googleapis.com/auth/logging.admin scope.
func (r *ProjectsLogsSinksService) Update(sinkName string, logsink *LogSink) *ProjectsLogsSinksUpdateCall {
	c := &ProjectsLogsSinksUpdateCall{s: r.s, opt_: make(map[string]interface{})}
	c.sinkName = sinkName
	c.logsink = logsink
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsLogsSinksUpdateCall) Fields(s ...googleapi.Field) *ProjectsLogsSinksUpdateCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsLogsSinksUpdateCall) Do() (*LogSink, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.logsink)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "{+sinkName}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("PUT", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"sinkName": c.sinkName,
	})
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *LogSink
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Create or update the specified log sink resource. Requires https://www.googleapis.com/auth/logging.admin scope.",
	//   "httpMethod": "PUT",
	//   "id": "logging.projects.logs.sinks.update",
	//   "parameterOrder": [
	//     "sinkName"
	//   ],
	//   "parameters": {
	//     "sinkName": {
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{+sinkName}",
	//   "request": {
	//     "$ref": "LogSink"
	//   },
	//   "response": {
	//     "$ref": "LogSink"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "logging.projects.metrics.delete":

type ProjectsMetricsDeleteCall struct {
	s      *Service
	metric string
	opt_   map[string]interface{}
}

// Delete: Deletes the specified log metric. Requires
// https://www.googleapis.com/auth/logging.admin scope.
func (r *ProjectsMetricsService) Delete(metric string) *ProjectsMetricsDeleteCall {
	c := &ProjectsMetricsDeleteCall{s: r.s, opt_: make(map[string]interface{})}
	c.metric = metric
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsMetricsDeleteCall) Fields(s ...googleapi.Field) *ProjectsMetricsDeleteCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsMetricsDeleteCall) Do() (*Empty, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "{+metric}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"metric": c.metric,
	})
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Empty
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Deletes the specified log metric. Requires https://www.googleapis.com/auth/logging.admin scope.",
	//   "httpMethod": "DELETE",
	//   "id": "logging.projects.metrics.delete",
	//   "parameterOrder": [
	//     "metric"
	//   ],
	//   "parameters": {
	//     "metric": {
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{+metric}",
	//   "response": {
	//     "$ref": "Empty"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "logging.projects.metrics.get":

type ProjectsMetricsGetCall struct {
	s      *Service
	metric string
	opt_   map[string]interface{}
}

// Get: Get the specified log metric resource. Requires
// https://www.googleapis.com/auth/logging.read scope.
func (r *ProjectsMetricsService) Get(metric string) *ProjectsMetricsGetCall {
	c := &ProjectsMetricsGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.metric = metric
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsMetricsGetCall) Fields(s ...googleapi.Field) *ProjectsMetricsGetCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsMetricsGetCall) Do() (*LogMetric, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "{+metric}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"metric": c.metric,
	})
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *LogMetric
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Get the specified log metric resource. Requires https://www.googleapis.com/auth/logging.read scope.",
	//   "httpMethod": "GET",
	//   "id": "logging.projects.metrics.get",
	//   "parameterOrder": [
	//     "metric"
	//   ],
	//   "parameters": {
	//     "metric": {
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{+metric}",
	//   "response": {
	//     "$ref": "LogMetric"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "logging.projects.metrics.list":

type ProjectsMetricsListCall struct {
	s         *Service
	projectId string
	opt_      map[string]interface{}
}

// List: List log metrics associated with the specified project.
// Requires https://www.googleapis.com/auth/logging.read scope.
func (r *ProjectsMetricsService) List(projectId string) *ProjectsMetricsListCall {
	c := &ProjectsMetricsListCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	return c
}

// PageSize sets the optional parameter "pageSize":
func (c *ProjectsMetricsListCall) PageSize(pageSize int64) *ProjectsMetricsListCall {
	c.opt_["pageSize"] = pageSize
	return c
}

// PageToken sets the optional parameter "pageToken":
func (c *ProjectsMetricsListCall) PageToken(pageToken string) *ProjectsMetricsListCall {
	c.opt_["pageToken"] = pageToken
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsMetricsListCall) Fields(s ...googleapi.Field) *ProjectsMetricsListCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsMetricsListCall) Do() (*ListLogMetricsResponse, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["pageSize"]; ok {
		params.Set("pageSize", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["pageToken"]; ok {
		params.Set("pageToken", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "{+projectId}/metrics")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectId": c.projectId,
	})
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *ListLogMetricsResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "List log metrics associated with the specified project. Requires https://www.googleapis.com/auth/logging.read scope.",
	//   "httpMethod": "GET",
	//   "id": "logging.projects.metrics.list",
	//   "parameterOrder": [
	//     "projectId"
	//   ],
	//   "parameters": {
	//     "pageSize": {
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "projectId": {
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{+projectId}/metrics",
	//   "response": {
	//     "$ref": "ListLogMetricsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "logging.projects.metrics.patch":

type ProjectsMetricsPatchCall struct {
	s          *Service
	metricName string
	metric     string
	logmetric  *LogMetric
	opt_       map[string]interface{}
}

// Patch: Create or update the specified log metric resource. Requires
// https://www.googleapis.com/auth/logging.admin scope. This method
// supports patch semantics.
func (r *ProjectsMetricsService) Patch(metricName string, metric string, logmetric *LogMetric) *ProjectsMetricsPatchCall {
	c := &ProjectsMetricsPatchCall{s: r.s, opt_: make(map[string]interface{})}
	c.metricName = metricName
	c.metric = metric
	c.logmetric = logmetric
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsMetricsPatchCall) Fields(s ...googleapi.Field) *ProjectsMetricsPatchCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsMetricsPatchCall) Do() (*LogMetric, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.logmetric)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	params.Set("metric", fmt.Sprintf("%v", c.metric))
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "{+metricName}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("PATCH", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"metricName": c.metricName,
	})
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *LogMetric
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Create or update the specified log metric resource. Requires https://www.googleapis.com/auth/logging.admin scope. This method supports patch semantics.",
	//   "httpMethod": "PATCH",
	//   "id": "logging.projects.metrics.patch",
	//   "parameterOrder": [
	//     "metricName",
	//     "metric"
	//   ],
	//   "parameters": {
	//     "metric": {
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "metricName": {
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{+metricName}",
	//   "request": {
	//     "$ref": "LogMetric"
	//   },
	//   "response": {
	//     "$ref": "LogMetric"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "logging.projects.metrics.update":

type ProjectsMetricsUpdateCall struct {
	s          *Service
	metricName string
	logmetric  *LogMetric
	opt_       map[string]interface{}
}

// Update: Create or update the specified log metric resource. Requires
// https://www.googleapis.com/auth/logging.admin scope.
func (r *ProjectsMetricsService) Update(metricName string, logmetric *LogMetric) *ProjectsMetricsUpdateCall {
	c := &ProjectsMetricsUpdateCall{s: r.s, opt_: make(map[string]interface{})}
	c.metricName = metricName
	c.logmetric = logmetric
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsMetricsUpdateCall) Fields(s ...googleapi.Field) *ProjectsMetricsUpdateCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsMetricsUpdateCall) Do() (*LogMetric, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.logmetric)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "{+metricName}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("PUT", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"metricName": c.metricName,
	})
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *LogMetric
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Create or update the specified log metric resource. Requires https://www.googleapis.com/auth/logging.admin scope.",
	//   "httpMethod": "PUT",
	//   "id": "logging.projects.metrics.update",
	//   "parameterOrder": [
	//     "metricName"
	//   ],
	//   "parameters": {
	//     "metricName": {
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{+metricName}",
	//   "request": {
	//     "$ref": "LogMetric"
	//   },
	//   "response": {
	//     "$ref": "LogMetric"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}
