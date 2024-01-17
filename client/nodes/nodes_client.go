//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2024 Weaviate B.V. All rights reserved.
//
//  CONTACT: hello@weaviate.io
//

// Code generated by go-swagger; DO NOT EDIT.

package nodes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new nodes API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for nodes API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	NodesGet(params *NodesGetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*NodesGetOK, error)

	NodesGetClass(params *NodesGetClassParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*NodesGetClassOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
NodesGet Returns status of Weaviate DB.
*/
func (a *Client) NodesGet(params *NodesGetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*NodesGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewNodesGetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "nodes.get",
		Method:             "GET",
		PathPattern:        "/nodes",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/yaml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &NodesGetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*NodesGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for nodes.get: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
NodesGetClass Returns status of Weaviate DB.
*/
func (a *Client) NodesGetClass(params *NodesGetClassParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*NodesGetClassOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewNodesGetClassParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "nodes.get.class",
		Method:             "GET",
		PathPattern:        "/nodes/{className}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/yaml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &NodesGetclassFinder{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*NodesGetClassOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for nodes.get.class: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
