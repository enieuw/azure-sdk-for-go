package web

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator 0.12.0.0
// Changes may cause incorrect behavior and will be lost if the code is
// regenerated.

import (
	"github.com/Azure/azure-sdk-for-go/Godeps/_workspace/src/github.com/Azure/go-autorest/autorest"
	"net/http"
	"net/url"
)

// ManagedHostingEnvironmentsClient is the use these APIs to manage Azure
// Websites resources through the Azure Resource Manager. All task operations
// conform to the HTTP/1.1 protocol specification and each operation returns
// an x-ms-request-id header that can be used to obtain information about the
// request. You must make sure that requests made to these resources are
// secure. For more information, see <a
// href="https://msdn.microsoft.com/en-us/library/azure/dn790557.aspx">Authenticating
// Azure Resource Manager requests.</a>
type ManagedHostingEnvironmentsClient struct {
	ManagementClient
}

// NewManagedHostingEnvironmentsClient creates an instance of the
// ManagedHostingEnvironmentsClient client.
func NewManagedHostingEnvironmentsClient(subscriptionID string) ManagedHostingEnvironmentsClient {
	return NewManagedHostingEnvironmentsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewManagedHostingEnvironmentsClientWithBaseURI creates an instance of the
// ManagedHostingEnvironmentsClient client.
func NewManagedHostingEnvironmentsClientWithBaseURI(baseURI string, subscriptionID string) ManagedHostingEnvironmentsClient {
	return ManagedHostingEnvironmentsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdateManagedHostingEnvironment sends the create or update managed
// hosting environment request.
//
// resourceGroupName is name of resource group name is name of managed hosting
// environment managedHostingEnvironmentEnvelope is properties of managed
// hosting environment
func (client ManagedHostingEnvironmentsClient) CreateOrUpdateManagedHostingEnvironment(resourceGroupName string, name string, managedHostingEnvironmentEnvelope HostingEnvironment) (result HostingEnvironment, ae error) {
	req, err := client.CreateOrUpdateManagedHostingEnvironmentPreparer(resourceGroupName, name, managedHostingEnvironmentEnvelope)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "web/ManagedHostingEnvironmentsClient", "CreateOrUpdateManagedHostingEnvironment", autorest.UndefinedStatusCode, "Failure preparing request")
	}

	resp, err := client.CreateOrUpdateManagedHostingEnvironmentSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "web/ManagedHostingEnvironmentsClient", "CreateOrUpdateManagedHostingEnvironment", resp.StatusCode, "Failure sending request")
	}

	result, err = client.CreateOrUpdateManagedHostingEnvironmentResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "web/ManagedHostingEnvironmentsClient", "CreateOrUpdateManagedHostingEnvironment", resp.StatusCode, "Failure responding to request")
	}

	return
}

// CreateOrUpdateManagedHostingEnvironmentPreparer prepares the CreateOrUpdateManagedHostingEnvironment request.
func (client ManagedHostingEnvironmentsClient) CreateOrUpdateManagedHostingEnvironmentPreparer(resourceGroupName string, name string, managedHostingEnvironmentEnvelope HostingEnvironment) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"name":              url.QueryEscape(name),
		"resourceGroupName": url.QueryEscape(resourceGroupName),
		"subscriptionId":    url.QueryEscape(client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/managedHostingEnvironments/{name}"),
		autorest.WithJSON(managedHostingEnvironmentEnvelope),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// CreateOrUpdateManagedHostingEnvironmentSender sends the CreateOrUpdateManagedHostingEnvironment request. The method will close the
// http.Response Body if it receives an error.
func (client ManagedHostingEnvironmentsClient) CreateOrUpdateManagedHostingEnvironmentSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, http.StatusOK, http.StatusAccepted, http.StatusBadRequest, http.StatusNotFound, http.StatusConflict)
}

// CreateOrUpdateManagedHostingEnvironmentResponder handles the response to the CreateOrUpdateManagedHostingEnvironment request. The method always
// closes the http.Response Body.
func (client ManagedHostingEnvironmentsClient) CreateOrUpdateManagedHostingEnvironmentResponder(resp *http.Response) (result HostingEnvironment, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		autorest.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusBadRequest, http.StatusNotFound, http.StatusConflict),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// DeleteManagedHostingEnvironment sends the delete managed hosting
// environment request.
//
// resourceGroupName is name of resource group name is name of managed hosting
// environment forceDelete is delete even if the managed hosting environment
// contains resources
func (client ManagedHostingEnvironmentsClient) DeleteManagedHostingEnvironment(resourceGroupName string, name string, forceDelete *bool) (result ObjectSet, ae error) {
	req, err := client.DeleteManagedHostingEnvironmentPreparer(resourceGroupName, name, forceDelete)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "web/ManagedHostingEnvironmentsClient", "DeleteManagedHostingEnvironment", autorest.UndefinedStatusCode, "Failure preparing request")
	}

	resp, err := client.DeleteManagedHostingEnvironmentSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "web/ManagedHostingEnvironmentsClient", "DeleteManagedHostingEnvironment", resp.StatusCode, "Failure sending request")
	}

	result, err = client.DeleteManagedHostingEnvironmentResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "web/ManagedHostingEnvironmentsClient", "DeleteManagedHostingEnvironment", resp.StatusCode, "Failure responding to request")
	}

	return
}

// DeleteManagedHostingEnvironmentPreparer prepares the DeleteManagedHostingEnvironment request.
func (client ManagedHostingEnvironmentsClient) DeleteManagedHostingEnvironmentPreparer(resourceGroupName string, name string, forceDelete *bool) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"name":              url.QueryEscape(name),
		"resourceGroupName": url.QueryEscape(resourceGroupName),
		"subscriptionId":    url.QueryEscape(client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if forceDelete != nil {
		queryParameters["forceDelete"] = forceDelete
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/managedHostingEnvironments/{name}"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// DeleteManagedHostingEnvironmentSender sends the DeleteManagedHostingEnvironment request. The method will close the
// http.Response Body if it receives an error.
func (client ManagedHostingEnvironmentsClient) DeleteManagedHostingEnvironmentSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, http.StatusOK, http.StatusAccepted, http.StatusBadRequest, http.StatusNotFound, http.StatusConflict)
}

// DeleteManagedHostingEnvironmentResponder handles the response to the DeleteManagedHostingEnvironment request. The method always
// closes the http.Response Body.
func (client ManagedHostingEnvironmentsClient) DeleteManagedHostingEnvironmentResponder(resp *http.Response) (result ObjectSet, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		autorest.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusBadRequest, http.StatusNotFound, http.StatusConflict),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetManagedHostingEnvironment sends the get managed hosting environment
// request.
//
// resourceGroupName is name of resource group name is name of managed hosting
// environment
func (client ManagedHostingEnvironmentsClient) GetManagedHostingEnvironment(resourceGroupName string, name string) (result ManagedHostingEnvironment, ae error) {
	req, err := client.GetManagedHostingEnvironmentPreparer(resourceGroupName, name)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "web/ManagedHostingEnvironmentsClient", "GetManagedHostingEnvironment", autorest.UndefinedStatusCode, "Failure preparing request")
	}

	resp, err := client.GetManagedHostingEnvironmentSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "web/ManagedHostingEnvironmentsClient", "GetManagedHostingEnvironment", resp.StatusCode, "Failure sending request")
	}

	result, err = client.GetManagedHostingEnvironmentResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "web/ManagedHostingEnvironmentsClient", "GetManagedHostingEnvironment", resp.StatusCode, "Failure responding to request")
	}

	return
}

// GetManagedHostingEnvironmentPreparer prepares the GetManagedHostingEnvironment request.
func (client ManagedHostingEnvironmentsClient) GetManagedHostingEnvironmentPreparer(resourceGroupName string, name string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"name":              url.QueryEscape(name),
		"resourceGroupName": url.QueryEscape(resourceGroupName),
		"subscriptionId":    url.QueryEscape(client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/managedHostingEnvironments/{name}"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// GetManagedHostingEnvironmentSender sends the GetManagedHostingEnvironment request. The method will close the
// http.Response Body if it receives an error.
func (client ManagedHostingEnvironmentsClient) GetManagedHostingEnvironmentSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, http.StatusOK)
}

// GetManagedHostingEnvironmentResponder handles the response to the GetManagedHostingEnvironment request. The method always
// closes the http.Response Body.
func (client ManagedHostingEnvironmentsClient) GetManagedHostingEnvironmentResponder(resp *http.Response) (result ManagedHostingEnvironment, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		autorest.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetManagedHostingEnvironmentOperation sends the get managed hosting
// environment operation request.
//
// resourceGroupName is name of resource group name is name of managed hosting
// environment operationID is operation identifier GUID
func (client ManagedHostingEnvironmentsClient) GetManagedHostingEnvironmentOperation(resourceGroupName string, name string, operationID string) (result ObjectSet, ae error) {
	req, err := client.GetManagedHostingEnvironmentOperationPreparer(resourceGroupName, name, operationID)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "web/ManagedHostingEnvironmentsClient", "GetManagedHostingEnvironmentOperation", autorest.UndefinedStatusCode, "Failure preparing request")
	}

	resp, err := client.GetManagedHostingEnvironmentOperationSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "web/ManagedHostingEnvironmentsClient", "GetManagedHostingEnvironmentOperation", resp.StatusCode, "Failure sending request")
	}

	result, err = client.GetManagedHostingEnvironmentOperationResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "web/ManagedHostingEnvironmentsClient", "GetManagedHostingEnvironmentOperation", resp.StatusCode, "Failure responding to request")
	}

	return
}

// GetManagedHostingEnvironmentOperationPreparer prepares the GetManagedHostingEnvironmentOperation request.
func (client ManagedHostingEnvironmentsClient) GetManagedHostingEnvironmentOperationPreparer(resourceGroupName string, name string, operationID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"name":              url.QueryEscape(name),
		"operationId":       url.QueryEscape(operationID),
		"resourceGroupName": url.QueryEscape(resourceGroupName),
		"subscriptionId":    url.QueryEscape(client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/managedHostingEnvironments/{name}/operations/{operationId}"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// GetManagedHostingEnvironmentOperationSender sends the GetManagedHostingEnvironmentOperation request. The method will close the
// http.Response Body if it receives an error.
func (client ManagedHostingEnvironmentsClient) GetManagedHostingEnvironmentOperationSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, http.StatusOK, http.StatusAccepted, http.StatusNotFound, http.StatusInternalServerError)
}

// GetManagedHostingEnvironmentOperationResponder handles the response to the GetManagedHostingEnvironmentOperation request. The method always
// closes the http.Response Body.
func (client ManagedHostingEnvironmentsClient) GetManagedHostingEnvironmentOperationResponder(resp *http.Response) (result ObjectSet, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		autorest.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNotFound, http.StatusInternalServerError),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetManagedHostingEnvironments sends the get managed hosting environments
// request.
//
// resourceGroupName is name of resource group
func (client ManagedHostingEnvironmentsClient) GetManagedHostingEnvironments(resourceGroupName string) (result HostingEnvironmentCollection, ae error) {
	req, err := client.GetManagedHostingEnvironmentsPreparer(resourceGroupName)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "web/ManagedHostingEnvironmentsClient", "GetManagedHostingEnvironments", autorest.UndefinedStatusCode, "Failure preparing request")
	}

	resp, err := client.GetManagedHostingEnvironmentsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "web/ManagedHostingEnvironmentsClient", "GetManagedHostingEnvironments", resp.StatusCode, "Failure sending request")
	}

	result, err = client.GetManagedHostingEnvironmentsResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "web/ManagedHostingEnvironmentsClient", "GetManagedHostingEnvironments", resp.StatusCode, "Failure responding to request")
	}

	return
}

// GetManagedHostingEnvironmentsPreparer prepares the GetManagedHostingEnvironments request.
func (client ManagedHostingEnvironmentsClient) GetManagedHostingEnvironmentsPreparer(resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": url.QueryEscape(resourceGroupName),
		"subscriptionId":    url.QueryEscape(client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/managedHostingEnvironments"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// GetManagedHostingEnvironmentsSender sends the GetManagedHostingEnvironments request. The method will close the
// http.Response Body if it receives an error.
func (client ManagedHostingEnvironmentsClient) GetManagedHostingEnvironmentsSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, http.StatusOK)
}

// GetManagedHostingEnvironmentsResponder handles the response to the GetManagedHostingEnvironments request. The method always
// closes the http.Response Body.
func (client ManagedHostingEnvironmentsClient) GetManagedHostingEnvironmentsResponder(resp *http.Response) (result HostingEnvironmentCollection, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		autorest.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetManagedHostingEnvironmentServerFarms sends the get managed hosting
// environment server farms request.
//
// resourceGroupName is name of resource group name is name of managed hosting
// environment
func (client ManagedHostingEnvironmentsClient) GetManagedHostingEnvironmentServerFarms(resourceGroupName string, name string) (result ServerFarmCollection, ae error) {
	req, err := client.GetManagedHostingEnvironmentServerFarmsPreparer(resourceGroupName, name)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "web/ManagedHostingEnvironmentsClient", "GetManagedHostingEnvironmentServerFarms", autorest.UndefinedStatusCode, "Failure preparing request")
	}

	resp, err := client.GetManagedHostingEnvironmentServerFarmsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "web/ManagedHostingEnvironmentsClient", "GetManagedHostingEnvironmentServerFarms", resp.StatusCode, "Failure sending request")
	}

	result, err = client.GetManagedHostingEnvironmentServerFarmsResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "web/ManagedHostingEnvironmentsClient", "GetManagedHostingEnvironmentServerFarms", resp.StatusCode, "Failure responding to request")
	}

	return
}

// GetManagedHostingEnvironmentServerFarmsPreparer prepares the GetManagedHostingEnvironmentServerFarms request.
func (client ManagedHostingEnvironmentsClient) GetManagedHostingEnvironmentServerFarmsPreparer(resourceGroupName string, name string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"name":              url.QueryEscape(name),
		"resourceGroupName": url.QueryEscape(resourceGroupName),
		"subscriptionId":    url.QueryEscape(client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/managedHostingEnvironments/{name}/serverfarms"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// GetManagedHostingEnvironmentServerFarmsSender sends the GetManagedHostingEnvironmentServerFarms request. The method will close the
// http.Response Body if it receives an error.
func (client ManagedHostingEnvironmentsClient) GetManagedHostingEnvironmentServerFarmsSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, http.StatusOK)
}

// GetManagedHostingEnvironmentServerFarmsResponder handles the response to the GetManagedHostingEnvironmentServerFarms request. The method always
// closes the http.Response Body.
func (client ManagedHostingEnvironmentsClient) GetManagedHostingEnvironmentServerFarmsResponder(resp *http.Response) (result ServerFarmCollection, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		autorest.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetManagedHostingEnvironmentSites sends the get managed hosting environment
// sites request.
//
// resourceGroupName is name of resource group name is name of managed hosting
// environment propertiesToInclude is comma separated list of site properties
// to include
func (client ManagedHostingEnvironmentsClient) GetManagedHostingEnvironmentSites(resourceGroupName string, name string, propertiesToInclude string) (result SiteCollection, ae error) {
	req, err := client.GetManagedHostingEnvironmentSitesPreparer(resourceGroupName, name, propertiesToInclude)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "web/ManagedHostingEnvironmentsClient", "GetManagedHostingEnvironmentSites", autorest.UndefinedStatusCode, "Failure preparing request")
	}

	resp, err := client.GetManagedHostingEnvironmentSitesSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "web/ManagedHostingEnvironmentsClient", "GetManagedHostingEnvironmentSites", resp.StatusCode, "Failure sending request")
	}

	result, err = client.GetManagedHostingEnvironmentSitesResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "web/ManagedHostingEnvironmentsClient", "GetManagedHostingEnvironmentSites", resp.StatusCode, "Failure responding to request")
	}

	return
}

// GetManagedHostingEnvironmentSitesPreparer prepares the GetManagedHostingEnvironmentSites request.
func (client ManagedHostingEnvironmentsClient) GetManagedHostingEnvironmentSitesPreparer(resourceGroupName string, name string, propertiesToInclude string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"name":              url.QueryEscape(name),
		"resourceGroupName": url.QueryEscape(resourceGroupName),
		"subscriptionId":    url.QueryEscape(client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(propertiesToInclude) > 0 {
		queryParameters["propertiesToInclude"] = propertiesToInclude
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/managedHostingEnvironments/{name}/sites"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// GetManagedHostingEnvironmentSitesSender sends the GetManagedHostingEnvironmentSites request. The method will close the
// http.Response Body if it receives an error.
func (client ManagedHostingEnvironmentsClient) GetManagedHostingEnvironmentSitesSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, http.StatusOK)
}

// GetManagedHostingEnvironmentSitesResponder handles the response to the GetManagedHostingEnvironmentSites request. The method always
// closes the http.Response Body.
func (client ManagedHostingEnvironmentsClient) GetManagedHostingEnvironmentSitesResponder(resp *http.Response) (result SiteCollection, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		autorest.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetManagedHostingEnvironmentVips sends the get managed hosting environment
// vips request.
//
// resourceGroupName is name of resource group name is name of managed hosting
// environment
func (client ManagedHostingEnvironmentsClient) GetManagedHostingEnvironmentVips(resourceGroupName string, name string) (result AddressResponse, ae error) {
	req, err := client.GetManagedHostingEnvironmentVipsPreparer(resourceGroupName, name)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "web/ManagedHostingEnvironmentsClient", "GetManagedHostingEnvironmentVips", autorest.UndefinedStatusCode, "Failure preparing request")
	}

	resp, err := client.GetManagedHostingEnvironmentVipsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "web/ManagedHostingEnvironmentsClient", "GetManagedHostingEnvironmentVips", resp.StatusCode, "Failure sending request")
	}

	result, err = client.GetManagedHostingEnvironmentVipsResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "web/ManagedHostingEnvironmentsClient", "GetManagedHostingEnvironmentVips", resp.StatusCode, "Failure responding to request")
	}

	return
}

// GetManagedHostingEnvironmentVipsPreparer prepares the GetManagedHostingEnvironmentVips request.
func (client ManagedHostingEnvironmentsClient) GetManagedHostingEnvironmentVipsPreparer(resourceGroupName string, name string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"name":              url.QueryEscape(name),
		"resourceGroupName": url.QueryEscape(resourceGroupName),
		"subscriptionId":    url.QueryEscape(client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/managedHostingEnvironments/{name}/capacities/virtualip"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// GetManagedHostingEnvironmentVipsSender sends the GetManagedHostingEnvironmentVips request. The method will close the
// http.Response Body if it receives an error.
func (client ManagedHostingEnvironmentsClient) GetManagedHostingEnvironmentVipsSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, http.StatusOK)
}

// GetManagedHostingEnvironmentVipsResponder handles the response to the GetManagedHostingEnvironmentVips request. The method always
// closes the http.Response Body.
func (client ManagedHostingEnvironmentsClient) GetManagedHostingEnvironmentVipsResponder(resp *http.Response) (result AddressResponse, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		autorest.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetManagedHostingEnvironmentWebHostingPlans sends the get managed hosting
// environment web hosting plans request.
//
// resourceGroupName is name of resource group name is name of managed hosting
// environment
func (client ManagedHostingEnvironmentsClient) GetManagedHostingEnvironmentWebHostingPlans(resourceGroupName string, name string) (result ServerFarmCollection, ae error) {
	req, err := client.GetManagedHostingEnvironmentWebHostingPlansPreparer(resourceGroupName, name)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "web/ManagedHostingEnvironmentsClient", "GetManagedHostingEnvironmentWebHostingPlans", autorest.UndefinedStatusCode, "Failure preparing request")
	}

	resp, err := client.GetManagedHostingEnvironmentWebHostingPlansSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "web/ManagedHostingEnvironmentsClient", "GetManagedHostingEnvironmentWebHostingPlans", resp.StatusCode, "Failure sending request")
	}

	result, err = client.GetManagedHostingEnvironmentWebHostingPlansResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "web/ManagedHostingEnvironmentsClient", "GetManagedHostingEnvironmentWebHostingPlans", resp.StatusCode, "Failure responding to request")
	}

	return
}

// GetManagedHostingEnvironmentWebHostingPlansPreparer prepares the GetManagedHostingEnvironmentWebHostingPlans request.
func (client ManagedHostingEnvironmentsClient) GetManagedHostingEnvironmentWebHostingPlansPreparer(resourceGroupName string, name string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"name":              url.QueryEscape(name),
		"resourceGroupName": url.QueryEscape(resourceGroupName),
		"subscriptionId":    url.QueryEscape(client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/managedHostingEnvironments/{name}/webhostingplans"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// GetManagedHostingEnvironmentWebHostingPlansSender sends the GetManagedHostingEnvironmentWebHostingPlans request. The method will close the
// http.Response Body if it receives an error.
func (client ManagedHostingEnvironmentsClient) GetManagedHostingEnvironmentWebHostingPlansSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, http.StatusOK)
}

// GetManagedHostingEnvironmentWebHostingPlansResponder handles the response to the GetManagedHostingEnvironmentWebHostingPlans request. The method always
// closes the http.Response Body.
func (client ManagedHostingEnvironmentsClient) GetManagedHostingEnvironmentWebHostingPlansResponder(resp *http.Response) (result ServerFarmCollection, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		autorest.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}