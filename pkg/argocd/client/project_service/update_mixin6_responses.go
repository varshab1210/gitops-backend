// Code generated by go-swagger; DO NOT EDIT.

package project_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/rhd-gitops-examples/gitops-backend/pkg/argocd/models"
)

// UpdateMixin6Reader is a Reader for the UpdateMixin6 structure.
type UpdateMixin6Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateMixin6Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateMixin6OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateMixin6OK creates a UpdateMixin6OK with default headers values
func NewUpdateMixin6OK() *UpdateMixin6OK {
	return &UpdateMixin6OK{}
}

/*UpdateMixin6OK handles this case with default header values.

(empty)
*/
type UpdateMixin6OK struct {
	Payload *models.V1alpha1AppProject
}

func (o *UpdateMixin6OK) Error() string {
	return fmt.Sprintf("[PUT /api/v1/projects/{project.metadata.name}][%d] updateMixin6OK  %+v", 200, o.Payload)
}

func (o *UpdateMixin6OK) GetPayload() *models.V1alpha1AppProject {
	return o.Payload
}

func (o *UpdateMixin6OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1alpha1AppProject)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}