// Code generated by go-swagger; DO NOT EDIT.

package account_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new account service API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for account service API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	CanI(params *CanIParams) (*CanIOK, error)

	CreateTokenMixin9(params *CreateTokenMixin9Params) (*CreateTokenMixin9OK, error)

	DeleteTokenMixin9(params *DeleteTokenMixin9Params) (*DeleteTokenMixin9OK, error)

	GetAccount(params *GetAccountParams) (*GetAccountOK, error)

	ListAccounts(params *ListAccountsParams) (*ListAccountsOK, error)

	UpdatePassword(params *UpdatePasswordParams) (*UpdatePasswordOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CanI can i API
*/
func (a *Client) CanI(params *CanIParams) (*CanIOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCanIParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CanI",
		Method:             "GET",
		PathPattern:        "/api/v1/account/can-i/{resource}/{action}/{subresource}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CanIReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CanIOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CanI: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CreateTokenMixin9 create token mixin9 API
*/
func (a *Client) CreateTokenMixin9(params *CreateTokenMixin9Params) (*CreateTokenMixin9OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateTokenMixin9Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CreateTokenMixin9",
		Method:             "POST",
		PathPattern:        "/api/v1/account/{name}/token",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CreateTokenMixin9Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateTokenMixin9OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CreateTokenMixin9: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteTokenMixin9 delete token mixin9 API
*/
func (a *Client) DeleteTokenMixin9(params *DeleteTokenMixin9Params) (*DeleteTokenMixin9OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteTokenMixin9Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteTokenMixin9",
		Method:             "DELETE",
		PathPattern:        "/api/v1/account/{name}/token/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteTokenMixin9Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteTokenMixin9OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DeleteTokenMixin9: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetAccount get account API
*/
func (a *Client) GetAccount(params *GetAccountParams) (*GetAccountOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAccountParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetAccount",
		Method:             "GET",
		PathPattern:        "/api/v1/account/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetAccountReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetAccountOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetAccount: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListAccounts list accounts API
*/
func (a *Client) ListAccounts(params *ListAccountsParams) (*ListAccountsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListAccountsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ListAccounts",
		Method:             "GET",
		PathPattern:        "/api/v1/account",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListAccountsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListAccountsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ListAccounts: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdatePassword updates password updates an account s password to a new value
*/
func (a *Client) UpdatePassword(params *UpdatePasswordParams) (*UpdatePasswordOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdatePasswordParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "UpdatePassword",
		Method:             "PUT",
		PathPattern:        "/api/v1/account/password",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UpdatePasswordReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdatePasswordOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for UpdatePassword: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}