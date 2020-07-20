// Code generated by go-swagger; DO NOT EDIT.

package account_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/rhd-gitops-examples/gitops-backend/pkg/argocd/models"
)

// ListAccountsReader is a Reader for the ListAccounts structure.
type ListAccountsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListAccountsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListAccountsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListAccountsOK creates a ListAccountsOK with default headers values
func NewListAccountsOK() *ListAccountsOK {
	return &ListAccountsOK{}
}

/*ListAccountsOK handles this case with default header values.

(empty)
*/
type ListAccountsOK struct {
	Payload *models.AccountAccountsList
}

func (o *ListAccountsOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/account][%d] listAccountsOK  %+v", 200, o.Payload)
}

func (o *ListAccountsOK) GetPayload() *models.AccountAccountsList {
	return o.Payload
}

func (o *ListAccountsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AccountAccountsList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}