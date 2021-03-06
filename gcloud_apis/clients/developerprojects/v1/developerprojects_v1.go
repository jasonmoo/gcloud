// Package developerprojects provides access to the Developer Projects API.
//
// Usage example:
//
//   import "google.golang.org/api/developerprojects/v1"
//   ...
//   developerprojectsService, err := developerprojects.New(oauthHttpClient)
package developerprojects

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

const apiId = "developerprojects:v1"
const apiName = "developerprojects"
const apiVersion = "v1"
const basePath = "https://www.googleapis.com/developerprojects/v1/"

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
	s.DeveloperProjects = NewDeveloperProjectsService(s)
	s.Projects = NewProjectsService(s)
	s.V1 = NewV1Service(s)
	return s, nil
}

type Service struct {
	client   *http.Client
	BasePath string // API endpoint base URL

	DeveloperProjects *DeveloperProjectsService

	Projects *ProjectsService

	V1 *V1Service
}

func NewDeveloperProjectsService(s *Service) *DeveloperProjectsService {
	rs := &DeveloperProjectsService{s: s}
	return rs
}

type DeveloperProjectsService struct {
	s *Service
}

func NewProjectsService(s *Service) *ProjectsService {
	rs := &ProjectsService{s: s}
	return rs
}

type ProjectsService struct {
	s *Service
}

func NewV1Service(s *Service) *V1Service {
	rs := &V1Service{s: s}
	return rs
}

type V1Service struct {
	s *Service
}

type Binding struct {
	// Members: Format of member entries: 1) * Matches any requesting
	// principal (users, service accounts or anonymous).
	//
	// 2) user:$emailid A
	// google user account using a fully qualified email address. For
	// example alice@gmail.com, joe@bestbuy.com
	//
	// 3)
	// serviceAccount:/projects/$project/serviceAccounts/$name An IAM
	// service account with a fully qualified REST name
	//
	// 4)
	// serviceAccount:$name An IAM service account with relative name.
	//
	// 5)
	// group:$emailid A google group with a fully qualified email address.
	// For example auth-ti-cloud@google.com
	//
	// 6)
	// authoritySelector:/projects/$project/authoritySelectors/$name An IAM
	// authority selector with a fully qualified REST name.
	//
	// 7)
	// authoritySelector:$name An IAM authority selector with a relative
	// name.
	//
	// 8) mdb:$name Ganpati group name. For example bigtable-prod,
	// zanzibar-dev
	//
	// Note: It is an error to use relative IAM names for
	// token policies. They can only be used for a resource (policy) which
	// has a project to infer from.
	// -------------------------------------------------------
	Members []string `json:"members,omitempty"`

	// Role: Required
	Role string `json:"role,omitempty"`
}

type Condition struct {
	// Iam: Trusted attributes supplied by the IAM system.
	Iam string `json:"iam,omitempty"`

	Op string `json:"op,omitempty"`

	// Svc: Trusted attributes discharged by the service.
	Svc string `json:"svc,omitempty"`

	// Sys: Trusted attributes supplied by any service that owns resources
	// and uses the IAM system for access control.
	Sys string `json:"sys,omitempty"`

	// Value: The object of the condition. Exactly one of these must be set.
	Value string `json:"value,omitempty"`

	Values []string `json:"values,omitempty"`
}

type CreateProjectRequest1 struct {
	// AppengineStorageLocation: The storage location for the AppEngine app.
	AppengineStorageLocation string `json:"appengineStorageLocation,omitempty"`

	// CreateAppengineProject: If true, an AppEngine project will be
	// created.
	CreateAppengineProject bool `json:"createAppengineProject,omitempty"`

	// Project: The project the user wishes to create. The user must supply
	// the 'project_id' field, and may optionally set the 'title' field. All
	// other fields must be unset.
	//
	// If the requested id is unavailable, the
	// request will fail.
	Project *Project `json:"project,omitempty"`
}

type ListProjectsResponse struct {
	// NextPageToken: Pagination token.
	//
	// If the result set is too large to
	// fit in a single response, this token will be filled in. It encodes
	// the position of the current result cursor. Feeding this value into a
	// new list request as 'page_token' parameter gives the next page of the
	// results.
	//
	// When next_page_token is not filled in, there is no next
	// page and the client is looking at the last page in the result
	// set.
	//
	// Pagination tokens have a limited lifetime.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// Projects: The list of projects that matched the list query, possibly
	// paginated.
	//
	// The resource is partially filled in, based on the
	// retrieval_options specified in the `retrieval_options` field of the
	// list request.
	Projects []*Project `json:"projects,omitempty"`
}

type LogConfig struct {
	Counter *LogConfigCounterOptions `json:"counter,omitempty"`
}

type LogConfigCounterOptions struct {
	Field string `json:"field,omitempty"`

	Metric string `json:"metric,omitempty"`
}

type Policy struct {
	// Bindings: It is an error to specify multiple bindings for the same
	// role. Bindings with no members may be silently dropped by the
	// underlying policy storage system.
	Bindings []*Binding `json:"bindings,omitempty"`

	// Etag: Can be used to perform a read-modify-write.
	Etag string `json:"etag,omitempty"`

	Rules []*Rule `json:"rules,omitempty"`

	// Version: The policy language version. The version of the policy
	// itself is represented by the etag.
	Version int64 `json:"version,omitempty"`
}

type Project struct {
	// AppengineName: Name identifying legacy projects. This field should
	// not be used for new projects. This field is read-only.
	AppengineName string `json:"appengineName,omitempty"`

	// CreatedMs: The time at which the project was created in milliseconds
	// since the epoch.
	//
	// This is a read-only field.
	CreatedMs int64 `json:"createdMs,omitempty,string"`

	Deprecated5 string `json:"deprecated5,omitempty"`

	// Labels: The labels associated with this project.
	Labels []*ProjectLabelsEntry `json:"labels,omitempty"`

	// LifecycleState: The project lifecycle state.
	//
	// This field is
	// read-only.
	LifecycleState string `json:"lifecycleState,omitempty"`

	// ProjectId: The unique, user-assigned id of the project. The id must
	// be 6?30 lowercase letters, digits, or hyphens. It must start with a
	// letter. Trailing hyphens are prohibited.
	//
	// Example: "tokyo-rain-123"
	// This field is read-only once set.
	ProjectId string `json:"projectId,omitempty"`

	// ProjectNumber: The number uniquely identifying the project.
	//
	// Example:
	// 415104041262 This field is read-only.
	ProjectNumber int64 `json:"projectNumber,omitempty,string"`

	// Title: The user-assigned title of the project. This field is optional
	// and may remain unset.
	//
	// Example: "My Project"
	//
	// This is a read-write
	// field.
	Title string `json:"title,omitempty"`
}

type ProjectLabelsEntry struct {
	Key string `json:"key,omitempty"`

	Value string `json:"value,omitempty"`
}

type Rule struct {
	// Action: Required
	Action string `json:"action,omitempty"`

	// Conditions: Additional restrictions that must be met
	Conditions []*Condition `json:"conditions,omitempty"`

	// Description: Human-readable description of the rule.
	Description string `json:"description,omitempty"`

	// In: The rule matches if the PRINCIPAL/AUTHORITY_SELECTOR is in this
	// set of entries.
	In []string `json:"in,omitempty"`

	// LogConfig: The config returned to callers of tech.iam.IAM.CheckPolicy
	// for any entries that match the LOG action.
	LogConfig []*LogConfig `json:"logConfig,omitempty"`

	// NotIn: The rule matches if the PRINCIPAL/AUTHORITY_SELECTOR is not in
	// this set of entries.
	NotIn []string `json:"notIn,omitempty"`

	// Permissions: A permission is a string of form '$type.$action' (e.g.,
	// 'storage.read').
	Permissions []string `json:"permissions,omitempty"`
}

type SetPolicyRequest struct {
	// ChangeDescription: Human readable summary of change, reason for
	// change, who requested the change etc. for auditing purposes. TODO: Do
	// we need more structure here for higher level tools?
	ChangeDescription string `json:"changeDescription,omitempty"`

	// Policy: REQUIRED: The complete policy to be applied to the
	// 'resource'. The size of the policy is limited to a few 10s of KB. An
	// empty policy is in general a valid policy but certain services (like
	// Projects) might reject them.
	Policy *Policy `json:"policy,omitempty"`

	// Wait: If true the SetPolicy call returns only after it is guaranteed
	// that a subsequent CheckPolicy call will observe the new policy. NOTE:
	// We may support this guarantee only for policy changes that broaden
	// access.
	Wait bool `json:"wait,omitempty"`
}

// method id "developerprojects.developerProjects.internalcreateproject":

type DeveloperProjectsInternalcreateprojectCall struct {
	s                     *Service
	createprojectrequest1 *CreateProjectRequest1
	opt_                  map[string]interface{}
}

// Internalcreateproject: Creates a project resource. This should
// produce the same project as CreateProject.
func (r *DeveloperProjectsService) Internalcreateproject(createprojectrequest1 *CreateProjectRequest1) *DeveloperProjectsInternalcreateprojectCall {
	c := &DeveloperProjectsInternalcreateprojectCall{s: r.s, opt_: make(map[string]interface{})}
	c.createprojectrequest1 = createprojectrequest1
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *DeveloperProjectsInternalcreateprojectCall) Fields(s ...googleapi.Field) *DeveloperProjectsInternalcreateprojectCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *DeveloperProjectsInternalcreateprojectCall) Do() (*Project, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.createprojectrequest1)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "developerProjects/internalcreateproject")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.SetOpaque(req.URL)
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
	var ret *Project
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Creates a project resource. This should produce the same project as CreateProject.",
	//   "httpMethod": "POST",
	//   "id": "developerprojects.developerProjects.internalcreateproject",
	//   "path": "developerProjects/internalcreateproject",
	//   "request": {
	//     "$ref": "CreateProjectRequest1"
	//   },
	//   "response": {
	//     "$ref": "Project"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "developerprojects.projects.create":

type ProjectsCreateCall struct {
	s       *Service
	project *Project
	opt_    map[string]interface{}
}

// Create: Creates a project resource.
//
// Initially, the project resource
// is owned by its creator exclusively. The creator may then grant
// permission to read or update the project to others.
func (r *ProjectsService) Create(project *Project) *ProjectsCreateCall {
	c := &ProjectsCreateCall{s: r.s, opt_: make(map[string]interface{})}
	c.project = project
	return c
}

// AppengineStorageLocation sets the optional parameter
// "appengineStorageLocation": The storage location for the AppEngine
// app.
func (c *ProjectsCreateCall) AppengineStorageLocation(appengineStorageLocation string) *ProjectsCreateCall {
	c.opt_["appengineStorageLocation"] = appengineStorageLocation
	return c
}

// CreateAppengineProject sets the optional parameter
// "createAppengineProject": If true, an AppEngine project will be
// created.
func (c *ProjectsCreateCall) CreateAppengineProject(createAppengineProject bool) *ProjectsCreateCall {
	c.opt_["createAppengineProject"] = createAppengineProject
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsCreateCall) Fields(s ...googleapi.Field) *ProjectsCreateCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsCreateCall) Do() (*Project, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.project)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["appengineStorageLocation"]; ok {
		params.Set("appengineStorageLocation", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["createAppengineProject"]; ok {
		params.Set("createAppengineProject", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "projects")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.SetOpaque(req.URL)
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
	var ret *Project
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Creates a project resource.\n\nInitially, the project resource is owned by its creator exclusively. The creator may then grant permission to read or update the project to others.",
	//   "httpMethod": "POST",
	//   "id": "developerprojects.projects.create",
	//   "parameters": {
	//     "appengineStorageLocation": {
	//       "description": "The storage location for the AppEngine app.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "createAppengineProject": {
	//       "default": "true",
	//       "description": "If true, an AppEngine project will be created.",
	//       "location": "query",
	//       "type": "boolean"
	//     }
	//   },
	//   "path": "projects",
	//   "request": {
	//     "$ref": "Project"
	//   },
	//   "response": {
	//     "$ref": "Project"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "developerprojects.projects.delete":

type ProjectsDeleteCall struct {
	s         *Service
	projectId string
	opt_      map[string]interface{}
}

// Delete: Initiates deletion for this project.
func (r *ProjectsService) Delete(projectId string) *ProjectsDeleteCall {
	c := &ProjectsDeleteCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsDeleteCall) Fields(s ...googleapi.Field) *ProjectsDeleteCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsDeleteCall) Do() error {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "projects/{projectId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectId": c.projectId,
	})
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return err
	}
	return nil
	// {
	//   "description": "Initiates deletion for this project.",
	//   "httpMethod": "DELETE",
	//   "id": "developerprojects.projects.delete",
	//   "parameterOrder": [
	//     "projectId"
	//   ],
	//   "parameters": {
	//     "projectId": {
	//       "description": "A reference that uniquely identifies the project.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "projects/{projectId}",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "developerprojects.projects.get":

type ProjectsGetCall struct {
	s         *Service
	projectId string
	opt_      map[string]interface{}
}

// Get: Retrieves a limited project metadata set, given any project
// identifier.
func (r *ProjectsService) Get(projectId string) *ProjectsGetCall {
	c := &ProjectsGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsGetCall) Fields(s ...googleapi.Field) *ProjectsGetCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsGetCall) Do() (*Project, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "projects/{projectId}")
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
	var ret *Project
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Retrieves a limited project metadata set, given any project identifier.",
	//   "httpMethod": "GET",
	//   "id": "developerprojects.projects.get",
	//   "parameterOrder": [
	//     "projectId"
	//   ],
	//   "parameters": {
	//     "projectId": {
	//       "description": "The unique identifier of a project.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "projects/{projectId}",
	//   "response": {
	//     "$ref": "Project"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "developerprojects.projects.list":

type ProjectsListCall struct {
	s    *Service
	opt_ map[string]interface{}
}

// List: Lists projects visible to the user.
func (r *ProjectsService) List() *ProjectsListCall {
	c := &ProjectsListCall{s: r.s, opt_: make(map[string]interface{})}
	return c
}

// MaxResults sets the optional parameter "maxResults": Maximum number
// of items to return on a single page.
func (c *ProjectsListCall) MaxResults(maxResults int64) *ProjectsListCall {
	c.opt_["maxResults"] = maxResults
	return c
}

// PageToken sets the optional parameter "pageToken": Pagination token.
func (c *ProjectsListCall) PageToken(pageToken string) *ProjectsListCall {
	c.opt_["pageToken"] = pageToken
	return c
}

// Query sets the optional parameter "query": A query expression for
// filtering the results of the request.
func (c *ProjectsListCall) Query(query string) *ProjectsListCall {
	c.opt_["query"] = query
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsListCall) Fields(s ...googleapi.Field) *ProjectsListCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsListCall) Do() (*ListProjectsResponse, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["maxResults"]; ok {
		params.Set("maxResults", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["pageToken"]; ok {
		params.Set("pageToken", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["query"]; ok {
		params.Set("query", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "projects")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.SetOpaque(req.URL)
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *ListProjectsResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Lists projects visible to the user.",
	//   "httpMethod": "GET",
	//   "id": "developerprojects.projects.list",
	//   "parameters": {
	//     "maxResults": {
	//       "description": "Maximum number of items to return on a single page.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Pagination token.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "query": {
	//       "description": "A query expression for filtering the results of the request.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "projects",
	//   "response": {
	//     "$ref": "ListProjectsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "developerprojects.projects.patch":

type ProjectsPatchCall struct {
	s         *Service
	projectId string
	project   *Project
	opt_      map[string]interface{}
}

// Patch: Updates the metadata associated with the project.
//
// . This
// method supports patch semantics.
func (r *ProjectsService) Patch(projectId string, project *Project) *ProjectsPatchCall {
	c := &ProjectsPatchCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	c.project = project
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsPatchCall) Fields(s ...googleapi.Field) *ProjectsPatchCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsPatchCall) Do() (*Project, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.project)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "projects/{projectId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("PATCH", urls, body)
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
	var ret *Project
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Updates the metadata associated with the project.\n\n. This method supports patch semantics.",
	//   "httpMethod": "PATCH",
	//   "id": "developerprojects.projects.patch",
	//   "parameterOrder": [
	//     "projectId"
	//   ],
	//   "parameters": {
	//     "projectId": {
	//       "description": "The unique identifier of the project to be updated.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "projects/{projectId}",
	//   "request": {
	//     "$ref": "Project"
	//   },
	//   "response": {
	//     "$ref": "Project"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "developerprojects.projects.undelete":

type ProjectsUndeleteCall struct {
	s         *Service
	projectId string
	opt_      map[string]interface{}
}

// Undelete: Request un-deletion for this project.
func (r *ProjectsService) Undelete(projectId string) *ProjectsUndeleteCall {
	c := &ProjectsUndeleteCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsUndeleteCall) Fields(s ...googleapi.Field) *ProjectsUndeleteCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsUndeleteCall) Do() error {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "projects/{projectId}:undelete")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectId": c.projectId,
	})
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return err
	}
	return nil
	// {
	//   "description": "Request un-deletion for this project.",
	//   "httpMethod": "POST",
	//   "id": "developerprojects.projects.undelete",
	//   "parameterOrder": [
	//     "projectId"
	//   ],
	//   "parameters": {
	//     "projectId": {
	//       "description": "A reference that uniquely identifies the project.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "projects/{projectId}:undelete",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "developerprojects.projects.update":

type ProjectsUpdateCall struct {
	s         *Service
	projectId string
	project   *Project
	opt_      map[string]interface{}
}

// Update: Updates the metadata associated with the project.
func (r *ProjectsService) Update(projectId string, project *Project) *ProjectsUpdateCall {
	c := &ProjectsUpdateCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	c.project = project
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsUpdateCall) Fields(s ...googleapi.Field) *ProjectsUpdateCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsUpdateCall) Do() (*Project, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.project)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "projects/{projectId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("PUT", urls, body)
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
	var ret *Project
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Updates the metadata associated with the project.",
	//   "httpMethod": "PUT",
	//   "id": "developerprojects.projects.update",
	//   "parameterOrder": [
	//     "projectId"
	//   ],
	//   "parameters": {
	//     "projectId": {
	//       "description": "The unique identifier of the project to be updated.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "projects/{projectId}",
	//   "request": {
	//     "$ref": "Project"
	//   },
	//   "response": {
	//     "$ref": "Project"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "developerprojects.v1.getIamPolicy":

type V1GetIamPolicyCall struct {
	s        *Service
	resource string
	opt_     map[string]interface{}
}

// GetIamPolicy: Returns the IAM access control policy for specified
// project. Resource should be in format: "/projectid" Uses POST instead
// of GET because Protiary v2 can't support repeated messages in GET
// requests.
func (r *V1Service) GetIamPolicy(resource string) *V1GetIamPolicyCall {
	c := &V1GetIamPolicyCall{s: r.s, opt_: make(map[string]interface{})}
	c.resource = resource
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *V1GetIamPolicyCall) Fields(s ...googleapi.Field) *V1GetIamPolicyCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *V1GetIamPolicyCall) Do() (*Policy, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/{+resource}:getIamPolicy")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"resource": c.resource,
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
	var ret *Policy
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Returns the IAM access control policy for specified project. Resource should be in format: \"/projectid\" Uses POST instead of GET because Protiary v2 can't support repeated messages in GET requests.",
	//   "httpMethod": "POST",
	//   "id": "developerprojects.v1.getIamPolicy",
	//   "parameterOrder": [
	//     "resource"
	//   ],
	//   "parameters": {
	//     "resource": {
	//       "description": "REQUIRED: The resource for which policy is being requested. Usually some path like projects/{$project}/zones/{$zone}/disks/{$disk}.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/{+resource}:getIamPolicy",
	//   "response": {
	//     "$ref": "Policy"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "developerprojects.v1.setIamPolicy":

type V1SetIamPolicyCall struct {
	s                *Service
	resource         string
	setpolicyrequest *SetPolicyRequest
	opt_             map[string]interface{}
}

// SetIamPolicy: Sets the IAM access control policy for the specified
// project. Resource should be in format: "/projectid"
func (r *V1Service) SetIamPolicy(resource string, setpolicyrequest *SetPolicyRequest) *V1SetIamPolicyCall {
	c := &V1SetIamPolicyCall{s: r.s, opt_: make(map[string]interface{})}
	c.resource = resource
	c.setpolicyrequest = setpolicyrequest
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *V1SetIamPolicyCall) Fields(s ...googleapi.Field) *V1SetIamPolicyCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *V1SetIamPolicyCall) Do() (*Policy, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.setpolicyrequest)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/{+resource}:setIamPolicy")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"resource": c.resource,
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
	var ret *Policy
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Sets the IAM access control policy for the specified project. Resource should be in format: \"/projectid\"",
	//   "httpMethod": "POST",
	//   "id": "developerprojects.v1.setIamPolicy",
	//   "parameterOrder": [
	//     "resource"
	//   ],
	//   "parameters": {
	//     "resource": {
	//       "description": "REQUIRED: The resource for which policy is being specified. Usually some path like projects/{$project}/zones/{$zone}/disks/{$disk}.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/{+resource}:setIamPolicy",
	//   "request": {
	//     "$ref": "SetPolicyRequest"
	//   },
	//   "response": {
	//     "$ref": "Policy"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}
