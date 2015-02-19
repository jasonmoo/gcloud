// Package test provides access to the Google Cloud Test API.
//
// Usage example:
//
//   import "google.golang.org/api/test/v1"
//   ...
//   testService, err := test.New(oauthHttpClient)
package test

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

const apiId = "test:v1"
const apiName = "test"
const apiVersion = "v1"
const basePath = "https://test-devtools.googleapis.com/v1/"

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
	s.TestEnvironmentCatalog = NewTestEnvironmentCatalogService(s)
	return s, nil
}

type Service struct {
	client   *http.Client
	BasePath string // API endpoint base URL

	Projects *ProjectsService

	TestEnvironmentCatalog *TestEnvironmentCatalogService
}

func NewProjectsService(s *Service) *ProjectsService {
	rs := &ProjectsService{s: s}
	rs.Device = NewProjectsDeviceService(s)
	rs.Devices = NewProjectsDevicesService(s)
	rs.TestExecutions = NewProjectsTestExecutionsService(s)
	rs.TestMatrices = NewProjectsTestMatricesService(s)
	return rs
}

type ProjectsService struct {
	s *Service

	Device *ProjectsDeviceService

	Devices *ProjectsDevicesService

	TestExecutions *ProjectsTestExecutionsService

	TestMatrices *ProjectsTestMatricesService
}

func NewProjectsDeviceService(s *Service) *ProjectsDeviceService {
	rs := &ProjectsDeviceService{s: s}
	return rs
}

type ProjectsDeviceService struct {
	s *Service
}

func NewProjectsDevicesService(s *Service) *ProjectsDevicesService {
	rs := &ProjectsDevicesService{s: s}
	return rs
}

type ProjectsDevicesService struct {
	s *Service
}

func NewProjectsTestExecutionsService(s *Service) *ProjectsTestExecutionsService {
	rs := &ProjectsTestExecutionsService{s: s}
	return rs
}

type ProjectsTestExecutionsService struct {
	s *Service
}

func NewProjectsTestMatricesService(s *Service) *ProjectsTestMatricesService {
	rs := &ProjectsTestMatricesService{s: s}
	return rs
}

type ProjectsTestMatricesService struct {
	s *Service
}

func NewTestEnvironmentCatalogService(s *Service) *TestEnvironmentCatalogService {
	rs := &TestEnvironmentCatalogService{s: s}
	return rs
}

type TestEnvironmentCatalogService struct {
	s *Service
}

type AndroidDevice struct {
	AndroidModelId string `json:"androidModelId,omitempty"`

	AndroidVersionId string `json:"androidVersionId,omitempty"`

	Locale string `json:"locale,omitempty"`

	Orientation string `json:"orientation,omitempty"`
}

type AndroidDeviceCatalog struct {
	Models []*AndroidModel `json:"models,omitempty"`

	RuntimeConfiguration *AndroidRuntimeConfiguration `json:"runtimeConfiguration,omitempty"`

	Versions []*AndroidVersion `json:"versions,omitempty"`
}

type AndroidInstrumentationTest struct {
	AppApk *FileReference `json:"appApk,omitempty"`

	AppPackageId string `json:"appPackageId,omitempty"`

	TestApk *FileReference `json:"testApk,omitempty"`

	TestPackageId string `json:"testPackageId,omitempty"`

	TestRunnerClass string `json:"testRunnerClass,omitempty"`

	TestTargets []string `json:"testTargets,omitempty"`
}

type AndroidMatrix struct {
	AndroidModelIds []string `json:"androidModelIds,omitempty"`

	AndroidVersionIds []string `json:"androidVersionIds,omitempty"`

	Locales []string `json:"locales,omitempty"`

	Orientations []string `json:"orientations,omitempty"`
}

type AndroidModel struct {
	Form string `json:"form,omitempty"`

	Id string `json:"id,omitempty"`

	Manufacturer string `json:"manufacturer,omitempty"`

	Name string `json:"name,omitempty"`

	ScreenX int64 `json:"screenX,omitempty"`

	ScreenY int64 `json:"screenY,omitempty"`

	SupportedVersionIds []string `json:"supportedVersionIds,omitempty"`
}

type AndroidMonkeyTest struct {
	AppApk *FileReference `json:"appApk,omitempty"`

	AppPackageId string `json:"appPackageId,omitempty"`

	EventCount int64 `json:"eventCount,omitempty"`

	EventDelay string `json:"eventDelay,omitempty"`

	RandomSeed int64 `json:"randomSeed,omitempty"`
}

type AndroidRoboTest struct {
	AppApk *FileReference `json:"appApk,omitempty"`

	AppPackageId string `json:"appPackageId,omitempty"`
}

type AndroidRuntimeConfiguration struct {
	Locales []*Locale `json:"locales,omitempty"`

	Orientations []*Orientation `json:"orientations,omitempty"`
}

type AndroidVersion struct {
	ApiLevel int64 `json:"apiLevel,omitempty"`

	CodeName string `json:"codeName,omitempty"`

	Distribution *Distribution `json:"distribution,omitempty"`

	Id string `json:"id,omitempty"`

	ReleaseDate string `json:"releaseDate,omitempty"`

	VersionString string `json:"versionString,omitempty"`
}

type BlobstoreFile struct {
	BlobId string `json:"blobId,omitempty"`

	Md5Hash string `json:"md5Hash,omitempty"`
}

type CancelTestExecutionResponse struct {
	TestState string `json:"testState,omitempty"`
}

type CancelTestMatrixResponse struct {
	TestState string `json:"testState,omitempty"`
}

type ConnectionInfo struct {
	AdbPort int64 `json:"adbPort,omitempty"`

	IpAddress string `json:"ipAddress,omitempty"`

	SshPort int64 `json:"sshPort,omitempty"`

	VncPassword string `json:"vncPassword,omitempty"`

	VncPort int64 `json:"vncPort,omitempty"`
}

type Device struct {
	AndroidDevice *AndroidDevice `json:"androidDevice,omitempty"`

	CreationTime string `json:"creationTime,omitempty"`

	DeviceDetails *DeviceDetails `json:"deviceDetails,omitempty"`

	Id string `json:"id,omitempty"`

	State string `json:"state,omitempty"`

	StateDetails *DeviceStateDetails `json:"stateDetails,omitempty"`
}

type DeviceDetails struct {
	ConnectionInfo *ConnectionInfo `json:"connectionInfo,omitempty"`

	GceInstanceDetails *GceInstanceDetails `json:"gceInstanceDetails,omitempty"`
}

type DeviceStateDetails struct {
	ErrorDetails string `json:"errorDetails,omitempty"`

	ProgressDetails string `json:"progressDetails,omitempty"`
}

type Distribution struct {
	MarketShare float64 `json:"marketShare,omitempty"`

	MeasurementTime string `json:"measurementTime,omitempty"`
}

type Empty struct {
}

type Environment struct {
	AndroidDevice *AndroidDevice `json:"androidDevice,omitempty"`
}

type EnvironmentMatrix struct {
	AndroidMatrix *AndroidMatrix `json:"androidMatrix,omitempty"`
}

type FileReference struct {
	Blob *BlobstoreFile `json:"blob,omitempty"`

	GcsPath string `json:"gcsPath,omitempty"`
}

type GceInstanceDetails struct {
	Name string `json:"name,omitempty"`

	ProjectId string `json:"projectId,omitempty"`

	Zone string `json:"zone,omitempty"`
}

type GoogleCloudStorage struct {
	GcsPath string `json:"gcsPath,omitempty"`
}

type ListDevicesResponse struct {
	Devices []*Device `json:"devices,omitempty"`

	NextPageToken string `json:"nextPageToken,omitempty"`
}

type ListTestExecutionsResponse struct {
	TestExecutions []*TestExecution `json:"testExecutions,omitempty"`
}

type ListTestMatricesResponse struct {
	TestMatrices []*TestMatrix `json:"testMatrices,omitempty"`
}

type Locale struct {
	Id string `json:"id,omitempty"`

	Name string `json:"name,omitempty"`

	Region string `json:"region,omitempty"`
}

type Orientation struct {
	Id string `json:"id,omitempty"`

	Name string `json:"name,omitempty"`
}

type ResultStorage struct {
	GoogleCloudStorage *GoogleCloudStorage `json:"googleCloudStorage,omitempty"`

	ToolResultsExecutionId string `json:"toolResultsExecutionId,omitempty"`

	ToolResultsHistoryId string `json:"toolResultsHistoryId,omitempty"`

	ToolResultsStepId string `json:"toolResultsStepId,omitempty"`
}

type TestDetails struct {
	ErrorDetails string `json:"errorDetails,omitempty"`

	ProgressDetails string `json:"progressDetails,omitempty"`
}

type TestEnvironmentCatalog struct {
	AndroidDeviceCatalog *AndroidDeviceCatalog `json:"androidDeviceCatalog,omitempty"`
}

type TestExecution struct {
	Environment *Environment `json:"environment,omitempty"`

	Id string `json:"id,omitempty"`

	Labels map[string]string `json:"labels,omitempty"`

	ResultStorage *ResultStorage `json:"resultStorage,omitempty"`

	State string `json:"state,omitempty"`

	TestDetails *TestDetails `json:"testDetails,omitempty"`

	TestSpecification *TestSpecification `json:"testSpecification,omitempty"`

	Timestamp string `json:"timestamp,omitempty"`
}

type TestMatrix struct {
	EnvironmentMatrix *EnvironmentMatrix `json:"environmentMatrix,omitempty"`

	Labels map[string]string `json:"labels,omitempty"`

	ResultStorage *ResultStorage `json:"resultStorage,omitempty"`

	TestExecutions []*TestExecution `json:"testExecutions,omitempty"`

	TestMatrixId string `json:"testMatrixId,omitempty"`

	TestSpecification *TestSpecification `json:"testSpecification,omitempty"`
}

type TestSpecification struct {
	AndroidInstrumentationTest *AndroidInstrumentationTest `json:"androidInstrumentationTest,omitempty"`

	AndroidMonkeyTest *AndroidMonkeyTest `json:"androidMonkeyTest,omitempty"`

	AndroidRoboTest *AndroidRoboTest `json:"androidRoboTest,omitempty"`

	TestTimeout string `json:"testTimeout,omitempty"`
}

// method id "test.projects.device.delete":

type ProjectsDeviceDeleteCall struct {
	s         *Service
	projectId string
	deviceId  string
	opt_      map[string]interface{}
}

// Delete: Deletes a GCE Android device instance. May return any of the
// following canonical error codes: - PERMISSION_DENIED - if the user is
// not authorized to read project - INVALID_ARGUMENT - if the request is
// malformed - NOT_FOUND - if the project does not exist
func (r *ProjectsDeviceService) Delete(projectId string, deviceId string) *ProjectsDeviceDeleteCall {
	c := &ProjectsDeviceDeleteCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	c.deviceId = deviceId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsDeviceDeleteCall) Fields(s ...googleapi.Field) *ProjectsDeviceDeleteCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsDeviceDeleteCall) Do() (*Empty, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "projects/{projectId}/device/{deviceId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectId": c.projectId,
		"deviceId":  c.deviceId,
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
	//   "description": "Deletes a GCE Android device instance. May return any of the following canonical error codes: - PERMISSION_DENIED - if the user is not authorized to read project - INVALID_ARGUMENT - if the request is malformed - NOT_FOUND - if the project does not exist",
	//   "httpMethod": "DELETE",
	//   "id": "test.projects.device.delete",
	//   "parameterOrder": [
	//     "projectId",
	//     "deviceId"
	//   ],
	//   "parameters": {
	//     "deviceId": {
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projectId": {
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "projects/{projectId}/device/{deviceId}",
	//   "response": {
	//     "$ref": "Empty"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "test.projects.devices.create":

type ProjectsDevicesCreateCall struct {
	s         *Service
	projectId string
	device    *Device
	opt_      map[string]interface{}
}

// Create: Creates a new GCE Android device. May return any of the
// following canonical error codes: - PERMISSION_DENIED - if the user is
// not authorized to write to project - INVALID_ARGUMENT - if the
// request is malformed - NOT_FOUND - if the device type or project does
// not exist
func (r *ProjectsDevicesService) Create(projectId string, device *Device) *ProjectsDevicesCreateCall {
	c := &ProjectsDevicesCreateCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	c.device = device
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsDevicesCreateCall) Fields(s ...googleapi.Field) *ProjectsDevicesCreateCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsDevicesCreateCall) Do() (*Device, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.device)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "projects/{projectId}/devices")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectId": c.projectId,
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
	var ret *Device
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Creates a new GCE Android device. May return any of the following canonical error codes: - PERMISSION_DENIED - if the user is not authorized to write to project - INVALID_ARGUMENT - if the request is malformed - NOT_FOUND - if the device type or project does not exist",
	//   "httpMethod": "POST",
	//   "id": "test.projects.devices.create",
	//   "parameterOrder": [
	//     "projectId"
	//   ],
	//   "parameters": {
	//     "projectId": {
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "projects/{projectId}/devices",
	//   "request": {
	//     "$ref": "Device"
	//   },
	//   "response": {
	//     "$ref": "Device"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "test.projects.devices.get":

type ProjectsDevicesGetCall struct {
	s         *Service
	projectId string
	deviceId  string
	opt_      map[string]interface{}
}

// Get: Returns the GCE Android device. May return any of the following
// canonical error codes: - PERMISSION_DENIED - if the user is not
// authorized to read project - INVALID_ARGUMENT - if the request is
// malformed - NOT_FOUND - if the device type or project does not exist
func (r *ProjectsDevicesService) Get(projectId string, deviceId string) *ProjectsDevicesGetCall {
	c := &ProjectsDevicesGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	c.deviceId = deviceId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsDevicesGetCall) Fields(s ...googleapi.Field) *ProjectsDevicesGetCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsDevicesGetCall) Do() (*Device, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "projects/{projectId}/devices/{deviceId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectId": c.projectId,
		"deviceId":  c.deviceId,
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
	var ret *Device
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Returns the GCE Android device. May return any of the following canonical error codes: - PERMISSION_DENIED - if the user is not authorized to read project - INVALID_ARGUMENT - if the request is malformed - NOT_FOUND - if the device type or project does not exist",
	//   "httpMethod": "GET",
	//   "id": "test.projects.devices.get",
	//   "parameterOrder": [
	//     "projectId",
	//     "deviceId"
	//   ],
	//   "parameters": {
	//     "deviceId": {
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projectId": {
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "projects/{projectId}/devices/{deviceId}",
	//   "response": {
	//     "$ref": "Device"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "test.projects.devices.list":

type ProjectsDevicesListCall struct {
	s         *Service
	projectId string
	opt_      map[string]interface{}
}

// List: Lists all the current devices. May return any of the following
// canonical error codes: - PERMISSION_DENIED - if the user is not
// authorized to read project - INVALID_ARGUMENT - if the request is
// malformed - NOT_FOUND - if the project does not exist
func (r *ProjectsDevicesService) List(projectId string) *ProjectsDevicesListCall {
	c := &ProjectsDevicesListCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	return c
}

// PageSize sets the optional parameter "pageSize":
func (c *ProjectsDevicesListCall) PageSize(pageSize int64) *ProjectsDevicesListCall {
	c.opt_["pageSize"] = pageSize
	return c
}

// PageToken sets the optional parameter "pageToken":
func (c *ProjectsDevicesListCall) PageToken(pageToken string) *ProjectsDevicesListCall {
	c.opt_["pageToken"] = pageToken
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsDevicesListCall) Fields(s ...googleapi.Field) *ProjectsDevicesListCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsDevicesListCall) Do() (*ListDevicesResponse, error) {
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
	urls := googleapi.ResolveRelative(c.s.BasePath, "projects/{projectId}/devices")
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
	var ret *ListDevicesResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Lists all the current devices. May return any of the following canonical error codes: - PERMISSION_DENIED - if the user is not authorized to read project - INVALID_ARGUMENT - if the request is malformed - NOT_FOUND - if the project does not exist",
	//   "httpMethod": "GET",
	//   "id": "test.projects.devices.list",
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
	//   "path": "projects/{projectId}/devices",
	//   "response": {
	//     "$ref": "ListDevicesResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "test.projects.testExecutions.cancel":

type ProjectsTestExecutionsCancelCall struct {
	s               *Service
	projectId       string
	testExecutionId string
	opt_            map[string]interface{}
}

// Cancel: Cancel an individual test execution. If the specified test
// execution is running it will be aborted. If it's pending then it will
// simply be removed from the queue. The cancelled test execution will
// still be visible in the results of ListTestExecutions and
// GetTestExecution. May return any of the following canonical error
// codes: - PERMISSION_DENIED - if the user is not authorized to read
// project - INVALID_ARGUMENT - if the request is malformed - NOT_FOUND
// - if the Test Execution does not exist
func (r *ProjectsTestExecutionsService) Cancel(projectId string, testExecutionId string) *ProjectsTestExecutionsCancelCall {
	c := &ProjectsTestExecutionsCancelCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	c.testExecutionId = testExecutionId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsTestExecutionsCancelCall) Fields(s ...googleapi.Field) *ProjectsTestExecutionsCancelCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsTestExecutionsCancelCall) Do() (*CancelTestExecutionResponse, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "projects/{projectId}/testExecutions/{testExecutionId}:cancel")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectId":       c.projectId,
		"testExecutionId": c.testExecutionId,
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
	var ret *CancelTestExecutionResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Cancel an individual test execution. If the specified test execution is running it will be aborted. If it's pending then it will simply be removed from the queue. The cancelled test execution will still be visible in the results of ListTestExecutions and GetTestExecution. May return any of the following canonical error codes: - PERMISSION_DENIED - if the user is not authorized to read project - INVALID_ARGUMENT - if the request is malformed - NOT_FOUND - if the Test Execution does not exist",
	//   "httpMethod": "POST",
	//   "id": "test.projects.testExecutions.cancel",
	//   "parameterOrder": [
	//     "projectId",
	//     "testExecutionId"
	//   ],
	//   "parameters": {
	//     "projectId": {
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "testExecutionId": {
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "projects/{projectId}/testExecutions/{testExecutionId}:cancel",
	//   "response": {
	//     "$ref": "CancelTestExecutionResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "test.projects.testExecutions.create":

type ProjectsTestExecutionsCreateCall struct {
	s             *Service
	projectId     string
	testexecution *TestExecution
	opt_          map[string]interface{}
}

// Create: Request to execute a single test according to the given
// specification. May return any of the following canonical error codes:
// - PERMISSION_DENIED - if the user is not authorized to write to
// project - INVALID_ARGUMENT - if the request is malformed -
// UNSUPPORTED - if the given test environment is not supported.
func (r *ProjectsTestExecutionsService) Create(projectId string, testexecution *TestExecution) *ProjectsTestExecutionsCreateCall {
	c := &ProjectsTestExecutionsCreateCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	c.testexecution = testexecution
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsTestExecutionsCreateCall) Fields(s ...googleapi.Field) *ProjectsTestExecutionsCreateCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsTestExecutionsCreateCall) Do() (*TestExecution, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.testexecution)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "projects/{projectId}/testExecutions")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectId": c.projectId,
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
	var ret *TestExecution
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Request to execute a single test according to the given specification. May return any of the following canonical error codes: - PERMISSION_DENIED - if the user is not authorized to write to project - INVALID_ARGUMENT - if the request is malformed - UNSUPPORTED - if the given test environment is not supported.",
	//   "httpMethod": "POST",
	//   "id": "test.projects.testExecutions.create",
	//   "parameterOrder": [
	//     "projectId"
	//   ],
	//   "parameters": {
	//     "projectId": {
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "projects/{projectId}/testExecutions",
	//   "request": {
	//     "$ref": "TestExecution"
	//   },
	//   "response": {
	//     "$ref": "TestExecution"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "test.projects.testExecutions.delete":

type ProjectsTestExecutionsDeleteCall struct {
	s               *Service
	projectId       string
	testExecutionId string
	opt_            map[string]interface{}
}

// Delete: Delete all record of an individual test execution. The test
// execution will first be canceled if it is running or queued. It will
// no longer appear in the response from ListTestExecutions or
// GetTestExecution. May return any of the following canonical error
// codes: - PERMISSION_DENIED - if the user is not authorized to read
// project - INVALID_ARGUMENT - if the request is malformed - NOT_FOUND
// - if the Test Execution does not exist
func (r *ProjectsTestExecutionsService) Delete(projectId string, testExecutionId string) *ProjectsTestExecutionsDeleteCall {
	c := &ProjectsTestExecutionsDeleteCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	c.testExecutionId = testExecutionId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsTestExecutionsDeleteCall) Fields(s ...googleapi.Field) *ProjectsTestExecutionsDeleteCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsTestExecutionsDeleteCall) Do() (*Empty, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "projects/{projectId}/testExecutions/{testExecutionId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectId":       c.projectId,
		"testExecutionId": c.testExecutionId,
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
	//   "description": "Delete all record of an individual test execution. The test execution will first be canceled if it is running or queued. It will no longer appear in the response from ListTestExecutions or GetTestExecution. May return any of the following canonical error codes: - PERMISSION_DENIED - if the user is not authorized to read project - INVALID_ARGUMENT - if the request is malformed - NOT_FOUND - if the Test Execution does not exist",
	//   "httpMethod": "DELETE",
	//   "id": "test.projects.testExecutions.delete",
	//   "parameterOrder": [
	//     "projectId",
	//     "testExecutionId"
	//   ],
	//   "parameters": {
	//     "projectId": {
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "testExecutionId": {
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "projects/{projectId}/testExecutions/{testExecutionId}",
	//   "response": {
	//     "$ref": "Empty"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "test.projects.testExecutions.get":

type ProjectsTestExecutionsGetCall struct {
	s               *Service
	projectId       string
	testExecutionId string
	opt_            map[string]interface{}
}

// Get: Check the status of an individual test execution. May return any
// of the following canonical error codes: - PERMISSION_DENIED - if the
// user is not authorized to read project - INVALID_ARGUMENT - if the
// request is malformed - NOT_FOUND - if the Test Execution does not
// exist
func (r *ProjectsTestExecutionsService) Get(projectId string, testExecutionId string) *ProjectsTestExecutionsGetCall {
	c := &ProjectsTestExecutionsGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	c.testExecutionId = testExecutionId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsTestExecutionsGetCall) Fields(s ...googleapi.Field) *ProjectsTestExecutionsGetCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsTestExecutionsGetCall) Do() (*TestExecution, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "projects/{projectId}/testExecutions/{testExecutionId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectId":       c.projectId,
		"testExecutionId": c.testExecutionId,
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
	var ret *TestExecution
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Check the status of an individual test execution. May return any of the following canonical error codes: - PERMISSION_DENIED - if the user is not authorized to read project - INVALID_ARGUMENT - if the request is malformed - NOT_FOUND - if the Test Execution does not exist",
	//   "httpMethod": "GET",
	//   "id": "test.projects.testExecutions.get",
	//   "parameterOrder": [
	//     "projectId",
	//     "testExecutionId"
	//   ],
	//   "parameters": {
	//     "projectId": {
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "testExecutionId": {
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "projects/{projectId}/testExecutions/{testExecutionId}",
	//   "response": {
	//     "$ref": "TestExecution"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "test.projects.testExecutions.list":

type ProjectsTestExecutionsListCall struct {
	s         *Service
	projectId string
	opt_      map[string]interface{}
}

// List: List Test Executions The executions are sorted by creation_time
// in descending order. The test_execution_id key will be used to order
// the executions with the same creation_time. May return any of the
// following canonical error codes: - PERMISSION_DENIED - if the user is
// not authorized to read project - INVALID_ARGUMENT - if the request is
// malformed
func (r *ProjectsTestExecutionsService) List(projectId string) *ProjectsTestExecutionsListCall {
	c := &ProjectsTestExecutionsListCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	return c
}

// Query sets the optional parameter "query":
func (c *ProjectsTestExecutionsListCall) Query(query string) *ProjectsTestExecutionsListCall {
	c.opt_["query"] = query
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsTestExecutionsListCall) Fields(s ...googleapi.Field) *ProjectsTestExecutionsListCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsTestExecutionsListCall) Do() (*ListTestExecutionsResponse, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["query"]; ok {
		params.Set("query", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "projects/{projectId}/testExecutions")
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
	var ret *ListTestExecutionsResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "List Test Executions The executions are sorted by creation_time in descending order. The test_execution_id key will be used to order the executions with the same creation_time. May return any of the following canonical error codes: - PERMISSION_DENIED - if the user is not authorized to read project - INVALID_ARGUMENT - if the request is malformed",
	//   "httpMethod": "GET",
	//   "id": "test.projects.testExecutions.list",
	//   "parameterOrder": [
	//     "projectId"
	//   ],
	//   "parameters": {
	//     "projectId": {
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "query": {
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "projects/{projectId}/testExecutions",
	//   "response": {
	//     "$ref": "ListTestExecutionsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "test.projects.testMatrices.cancel":

type ProjectsTestMatricesCancelCall struct {
	s            *Service
	projectId    string
	testMatrixId string
	opt_         map[string]interface{}
}

// Cancel: Cancel a test matrix. If the test executions associated with
// the matrix are running they will be aborted. If they're pending then
// they will simply be removed from the queue. The cancelled tests may
// still be queried via calls to List* and Get*. This is equivalent to
// calling CancelTestExecution once for each test execution in the
// matrix. May return any of the following canonical error codes: -
// PERMISSION_DENIED - if the user is not authorized to read project -
// INVALID_ARGUMENT - if the request is malformed - NOT_FOUND - if the
// Test Matrix does not exist
func (r *ProjectsTestMatricesService) Cancel(projectId string, testMatrixId string) *ProjectsTestMatricesCancelCall {
	c := &ProjectsTestMatricesCancelCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	c.testMatrixId = testMatrixId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsTestMatricesCancelCall) Fields(s ...googleapi.Field) *ProjectsTestMatricesCancelCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsTestMatricesCancelCall) Do() (*CancelTestMatrixResponse, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "projects/{projectId}/testMatrices/{testMatrixId}:cancel")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectId":    c.projectId,
		"testMatrixId": c.testMatrixId,
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
	var ret *CancelTestMatrixResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Cancel a test matrix. If the test executions associated with the matrix are running they will be aborted. If they're pending then they will simply be removed from the queue. The cancelled tests may still be queried via calls to List* and Get*. This is equivalent to calling CancelTestExecution once for each test execution in the matrix. May return any of the following canonical error codes: - PERMISSION_DENIED - if the user is not authorized to read project - INVALID_ARGUMENT - if the request is malformed - NOT_FOUND - if the Test Matrix does not exist",
	//   "httpMethod": "POST",
	//   "id": "test.projects.testMatrices.cancel",
	//   "parameterOrder": [
	//     "projectId",
	//     "testMatrixId"
	//   ],
	//   "parameters": {
	//     "projectId": {
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "testMatrixId": {
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "projects/{projectId}/testMatrices/{testMatrixId}:cancel",
	//   "response": {
	//     "$ref": "CancelTestMatrixResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "test.projects.testMatrices.create":

type ProjectsTestMatricesCreateCall struct {
	s          *Service
	projectId  string
	testmatrix *TestMatrix
	opt_       map[string]interface{}
}

// Create: Request to run a matrix of tests according to the given
// specifications. Unsupported environments will be ignored. May return
// any of the following canonical error codes: - PERMISSION_DENIED - if
// the user is not authorized to write to project - INVALID_ARGUMENT -
// if the request is malformed
func (r *ProjectsTestMatricesService) Create(projectId string, testmatrix *TestMatrix) *ProjectsTestMatricesCreateCall {
	c := &ProjectsTestMatricesCreateCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	c.testmatrix = testmatrix
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsTestMatricesCreateCall) Fields(s ...googleapi.Field) *ProjectsTestMatricesCreateCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsTestMatricesCreateCall) Do() (*TestMatrix, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.testmatrix)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "projects/{projectId}/testMatrices")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectId": c.projectId,
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
	var ret *TestMatrix
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Request to run a matrix of tests according to the given specifications. Unsupported environments will be ignored. May return any of the following canonical error codes: - PERMISSION_DENIED - if the user is not authorized to write to project - INVALID_ARGUMENT - if the request is malformed",
	//   "httpMethod": "POST",
	//   "id": "test.projects.testMatrices.create",
	//   "parameterOrder": [
	//     "projectId"
	//   ],
	//   "parameters": {
	//     "projectId": {
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "projects/{projectId}/testMatrices",
	//   "request": {
	//     "$ref": "TestMatrix"
	//   },
	//   "response": {
	//     "$ref": "TestMatrix"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "test.projects.testMatrices.delete":

type ProjectsTestMatricesDeleteCall struct {
	s            *Service
	projectId    string
	testMatrixId string
	opt_         map[string]interface{}
}

// Delete: Delete all record of a test matrix plus any associated test
// executions. May return any of the following canonical error codes: -
// PERMISSION_DENIED - if the user is not authorized to read project -
// INVALID_ARGUMENT - if the request is malformed - NOT_FOUND - if the
// Test Matrix does not exist
func (r *ProjectsTestMatricesService) Delete(projectId string, testMatrixId string) *ProjectsTestMatricesDeleteCall {
	c := &ProjectsTestMatricesDeleteCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	c.testMatrixId = testMatrixId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsTestMatricesDeleteCall) Fields(s ...googleapi.Field) *ProjectsTestMatricesDeleteCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsTestMatricesDeleteCall) Do() (*Empty, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "projects/{projectId}/testMatrices/{testMatrixId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectId":    c.projectId,
		"testMatrixId": c.testMatrixId,
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
	//   "description": "Delete all record of a test matrix plus any associated test executions. May return any of the following canonical error codes: - PERMISSION_DENIED - if the user is not authorized to read project - INVALID_ARGUMENT - if the request is malformed - NOT_FOUND - if the Test Matrix does not exist",
	//   "httpMethod": "DELETE",
	//   "id": "test.projects.testMatrices.delete",
	//   "parameterOrder": [
	//     "projectId",
	//     "testMatrixId"
	//   ],
	//   "parameters": {
	//     "projectId": {
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "testMatrixId": {
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "projects/{projectId}/testMatrices/{testMatrixId}",
	//   "response": {
	//     "$ref": "Empty"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "test.projects.testMatrices.get":

type ProjectsTestMatricesGetCall struct {
	s            *Service
	projectId    string
	testMatrixId string
	opt_         map[string]interface{}
}

// Get: Check the status of a test matrix. May return any of the
// following canonical error codes: - PERMISSION_DENIED - if the user is
// not authorized to read project - INVALID_ARGUMENT - if the request is
// malformed - NOT_FOUND - if the Test Matrix does not exist
func (r *ProjectsTestMatricesService) Get(projectId string, testMatrixId string) *ProjectsTestMatricesGetCall {
	c := &ProjectsTestMatricesGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	c.testMatrixId = testMatrixId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsTestMatricesGetCall) Fields(s ...googleapi.Field) *ProjectsTestMatricesGetCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsTestMatricesGetCall) Do() (*TestMatrix, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "projects/{projectId}/testMatrices/{testMatrixId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectId":    c.projectId,
		"testMatrixId": c.testMatrixId,
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
	var ret *TestMatrix
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Check the status of a test matrix. May return any of the following canonical error codes: - PERMISSION_DENIED - if the user is not authorized to read project - INVALID_ARGUMENT - if the request is malformed - NOT_FOUND - if the Test Matrix does not exist",
	//   "httpMethod": "GET",
	//   "id": "test.projects.testMatrices.get",
	//   "parameterOrder": [
	//     "projectId",
	//     "testMatrixId"
	//   ],
	//   "parameters": {
	//     "projectId": {
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "testMatrixId": {
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "projects/{projectId}/testMatrices/{testMatrixId}",
	//   "response": {
	//     "$ref": "TestMatrix"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "test.projects.testMatrices.list":

type ProjectsTestMatricesListCall struct {
	s         *Service
	projectId string
	opt_      map[string]interface{}
}

// List: List test matrices. The returned matrices are currently
// unsorted. May return any of the following canonical error codes: -
// PERMISSION_DENIED - if the user is not authorized to read project -
// INVALID_ARGUMENT - if the request is malformed
func (r *ProjectsTestMatricesService) List(projectId string) *ProjectsTestMatricesListCall {
	c := &ProjectsTestMatricesListCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	return c
}

// Query sets the optional parameter "query":
func (c *ProjectsTestMatricesListCall) Query(query string) *ProjectsTestMatricesListCall {
	c.opt_["query"] = query
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsTestMatricesListCall) Fields(s ...googleapi.Field) *ProjectsTestMatricesListCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsTestMatricesListCall) Do() (*ListTestMatricesResponse, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["query"]; ok {
		params.Set("query", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "projects/{projectId}/testMatrices")
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
	var ret *ListTestMatricesResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "List test matrices. The returned matrices are currently unsorted. May return any of the following canonical error codes: - PERMISSION_DENIED - if the user is not authorized to read project - INVALID_ARGUMENT - if the request is malformed",
	//   "httpMethod": "GET",
	//   "id": "test.projects.testMatrices.list",
	//   "parameterOrder": [
	//     "projectId"
	//   ],
	//   "parameters": {
	//     "projectId": {
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "query": {
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "projects/{projectId}/testMatrices",
	//   "response": {
	//     "$ref": "ListTestMatricesResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "test.testEnvironmentCatalog.get":

type TestEnvironmentCatalogGetCall struct {
	s               *Service
	environmentType string
	opt_            map[string]interface{}
}

// Get: Get the catalog of supported test environments. May return any
// of the following canonical error codes: - INVALID_ARGUMENT - if the
// request is malformed - NOT_FOUND - if the environment type does not
// exist - INTERNAL - if an internal error occurred
func (r *TestEnvironmentCatalogService) Get(environmentType string) *TestEnvironmentCatalogGetCall {
	c := &TestEnvironmentCatalogGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.environmentType = environmentType
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *TestEnvironmentCatalogGetCall) Fields(s ...googleapi.Field) *TestEnvironmentCatalogGetCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *TestEnvironmentCatalogGetCall) Do() (*TestEnvironmentCatalog, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "testEnvironmentCatalog/{environmentType}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"environmentType": c.environmentType,
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
	var ret *TestEnvironmentCatalog
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Get the catalog of supported test environments. May return any of the following canonical error codes: - INVALID_ARGUMENT - if the request is malformed - NOT_FOUND - if the environment type does not exist - INTERNAL - if an internal error occurred",
	//   "httpMethod": "GET",
	//   "id": "test.testEnvironmentCatalog.get",
	//   "parameterOrder": [
	//     "environmentType"
	//   ],
	//   "parameters": {
	//     "environmentType": {
	//       "enum": [
	//         "ANDROID",
	//         "UNSPECIFIED"
	//       ],
	//       "enumDescriptions": [
	//         "",
	//         ""
	//       ],
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "testEnvironmentCatalog/{environmentType}",
	//   "response": {
	//     "$ref": "TestEnvironmentCatalog"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}
