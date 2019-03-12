// Code generated by go-swagger; DO NOT EDIT.

// Copyright Prometheus Team
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package silence

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new silence API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for silence API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
DeleteSilence Delete a silence by its ID
*/
func (a *Client) DeleteSilence(params *DeleteSilenceParams) (*DeleteSilenceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteSilenceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteSilence",
		Method:             "DELETE",
		PathPattern:        "/silence/{silenceID}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteSilenceReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteSilenceOK), nil

}

/*
GetSilence Get a silence by its ID
*/
func (a *Client) GetSilence(params *GetSilenceParams) (*GetSilenceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSilenceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getSilence",
		Method:             "GET",
		PathPattern:        "/silence/{silenceID}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetSilenceReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetSilenceOK), nil

}

/*
GetSilences Get a list of silences
*/
func (a *Client) GetSilences(params *GetSilencesParams) (*GetSilencesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSilencesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getSilences",
		Method:             "GET",
		PathPattern:        "/silences",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetSilencesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetSilencesOK), nil

}

/*
PostSilences Post a new silence or update an existing one
*/
func (a *Client) PostSilences(params *PostSilencesParams) (*PostSilencesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostSilencesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "postSilences",
		Method:             "POST",
		PathPattern:        "/silences",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PostSilencesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostSilencesOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}