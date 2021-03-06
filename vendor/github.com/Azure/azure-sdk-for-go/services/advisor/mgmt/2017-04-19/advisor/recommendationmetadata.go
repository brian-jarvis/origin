package advisor

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
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// RecommendationMetadataClient is the REST APIs for Azure Advisor
type RecommendationMetadataClient struct {
	BaseClient
}

// NewRecommendationMetadataClient creates an instance of the RecommendationMetadataClient client.
func NewRecommendationMetadataClient(subscriptionID string) RecommendationMetadataClient {
	return NewRecommendationMetadataClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewRecommendationMetadataClientWithBaseURI creates an instance of the RecommendationMetadataClient client.
func NewRecommendationMetadataClientWithBaseURI(baseURI string, subscriptionID string) RecommendationMetadataClient {
	return RecommendationMetadataClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Get sends the get request.
// Parameters:
// name - name of metadata entity.
func (client RecommendationMetadataClient) Get(ctx context.Context, name string) (result SetObject, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RecommendationMetadataClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, name)
	if err != nil {
		err = autorest.NewErrorWithError(err, "advisor.RecommendationMetadataClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "advisor.RecommendationMetadataClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "advisor.RecommendationMetadataClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client RecommendationMetadataClient) GetPreparer(ctx context.Context, name string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"name": autorest.Encode("path", name),
	}

	const APIVersion = "2017-04-19"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Advisor/metadata/{name}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client RecommendationMetadataClient) GetSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	return autorest.SendWithSender(client, req, sd...)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client RecommendationMetadataClient) GetResponder(resp *http.Response) (result SetObject, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNotFound),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List sends the list request.
func (client RecommendationMetadataClient) List(ctx context.Context) (result MetadataEntityListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RecommendationMetadataClient.List")
		defer func() {
			sc := -1
			if result.melr.Response.Response != nil {
				sc = result.melr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "advisor.RecommendationMetadataClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.melr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "advisor.RecommendationMetadataClient", "List", resp, "Failure sending request")
		return
	}

	result.melr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "advisor.RecommendationMetadataClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client RecommendationMetadataClient) ListPreparer(ctx context.Context) (*http.Request, error) {
	const APIVersion = "2017-04-19"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/providers/Microsoft.Advisor/metadata"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client RecommendationMetadataClient) ListSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	return autorest.SendWithSender(client, req, sd...)
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client RecommendationMetadataClient) ListResponder(resp *http.Response) (result MetadataEntityListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client RecommendationMetadataClient) listNextResults(ctx context.Context, lastResults MetadataEntityListResult) (result MetadataEntityListResult, err error) {
	req, err := lastResults.metadataEntityListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "advisor.RecommendationMetadataClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "advisor.RecommendationMetadataClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "advisor.RecommendationMetadataClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client RecommendationMetadataClient) ListComplete(ctx context.Context) (result MetadataEntityListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RecommendationMetadataClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx)
	return
}
