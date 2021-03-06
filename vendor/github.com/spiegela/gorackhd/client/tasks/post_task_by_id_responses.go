package tasks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/spiegela/gorackhd/models"
)

// PostTaskByIDReader is a Reader for the PostTaskByID structure.
type PostTaskByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostTaskByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewPostTaskByIDCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewPostTaskByIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewPostTaskByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostTaskByIDCreated creates a PostTaskByIDCreated with default headers values
func NewPostTaskByIDCreated() *PostTaskByIDCreated {
	return &PostTaskByIDCreated{}
}

/*PostTaskByIDCreated handles this case with default header values.

Successfully posted the specified task
*/
type PostTaskByIDCreated struct {
	Payload PostTaskByIDCreatedBody
}

func (o *PostTaskByIDCreated) Error() string {
	return fmt.Sprintf("[POST /tasks/{identifier}][%d] postTaskByIdCreated  %+v", 201, o.Payload)
}

func (o *PostTaskByIDCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTaskByIDNotFound creates a PostTaskByIDNotFound with default headers values
func NewPostTaskByIDNotFound() *PostTaskByIDNotFound {
	return &PostTaskByIDNotFound{}
}

/*PostTaskByIDNotFound handles this case with default header values.

The specified task was not found
*/
type PostTaskByIDNotFound struct {
	Payload *models.Error
}

func (o *PostTaskByIDNotFound) Error() string {
	return fmt.Sprintf("[POST /tasks/{identifier}][%d] postTaskByIdNotFound  %+v", 404, o.Payload)
}

func (o *PostTaskByIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTaskByIDDefault creates a PostTaskByIDDefault with default headers values
func NewPostTaskByIDDefault(code int) *PostTaskByIDDefault {
	return &PostTaskByIDDefault{
		_statusCode: code,
	}
}

/*PostTaskByIDDefault handles this case with default header values.

Unexpected error
*/
type PostTaskByIDDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the post task by Id default response
func (o *PostTaskByIDDefault) Code() int {
	return o._statusCode
}

func (o *PostTaskByIDDefault) Error() string {
	return fmt.Sprintf("[POST /tasks/{identifier}][%d] postTaskById default  %+v", o._statusCode, o.Payload)
}

func (o *PostTaskByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*PostTaskByIDCreatedBody post task by ID created body
swagger:model PostTaskByIDCreatedBody
*/
type PostTaskByIDCreatedBody interface{}
