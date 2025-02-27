// Code generated by go-swagger; DO NOT EDIT.

package digitalocean

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"k8c.io/kubermatic/v2/pkg/test/e2e/utils/apiclient/models"
)

// ListDigitaloceanSizesNoCredentialsReader is a Reader for the ListDigitaloceanSizesNoCredentials structure.
type ListDigitaloceanSizesNoCredentialsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListDigitaloceanSizesNoCredentialsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListDigitaloceanSizesNoCredentialsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListDigitaloceanSizesNoCredentialsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListDigitaloceanSizesNoCredentialsOK creates a ListDigitaloceanSizesNoCredentialsOK with default headers values
func NewListDigitaloceanSizesNoCredentialsOK() *ListDigitaloceanSizesNoCredentialsOK {
	return &ListDigitaloceanSizesNoCredentialsOK{}
}

/* ListDigitaloceanSizesNoCredentialsOK describes a response with status code 200, with default header values.

DigitaloceanSizeList
*/
type ListDigitaloceanSizesNoCredentialsOK struct {
	Payload *models.DigitaloceanSizeList
}

func (o *ListDigitaloceanSizesNoCredentialsOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/providers/digitalocean/sizes][%d] listDigitaloceanSizesNoCredentialsOK  %+v", 200, o.Payload)
}
func (o *ListDigitaloceanSizesNoCredentialsOK) GetPayload() *models.DigitaloceanSizeList {
	return o.Payload
}

func (o *ListDigitaloceanSizesNoCredentialsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DigitaloceanSizeList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListDigitaloceanSizesNoCredentialsDefault creates a ListDigitaloceanSizesNoCredentialsDefault with default headers values
func NewListDigitaloceanSizesNoCredentialsDefault(code int) *ListDigitaloceanSizesNoCredentialsDefault {
	return &ListDigitaloceanSizesNoCredentialsDefault{
		_statusCode: code,
	}
}

/* ListDigitaloceanSizesNoCredentialsDefault describes a response with status code -1, with default header values.

errorResponse
*/
type ListDigitaloceanSizesNoCredentialsDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the list digitalocean sizes no credentials default response
func (o *ListDigitaloceanSizesNoCredentialsDefault) Code() int {
	return o._statusCode
}

func (o *ListDigitaloceanSizesNoCredentialsDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/providers/digitalocean/sizes][%d] listDigitaloceanSizesNoCredentials default  %+v", o._statusCode, o.Payload)
}
func (o *ListDigitaloceanSizesNoCredentialsDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListDigitaloceanSizesNoCredentialsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
