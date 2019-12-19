// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package s3

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type PutObjectLockConfigurationInput struct {
	_ struct{} `type:"structure" payload:"ObjectLockConfiguration"`

	// The bucket whose Object Lock configuration you want to create or replace.
	//
	// Bucket is a required field
	Bucket *string `location:"uri" locationName:"Bucket" type:"string" required:"true"`

	// The Object Lock configuration that you want to apply to the specified bucket.
	ObjectLockConfiguration *ObjectLockConfiguration `locationName:"ObjectLockConfiguration" type:"structure" xmlURI:"http://s3.amazonaws.com/doc/2006-03-01/"`

	// Confirms that the requester knows that she or he will be charged for the
	// request. Bucket owners need not specify this parameter in their requests.
	// Documentation on downloading objects from requester pays buckets can be found
	// at http://docs.aws.amazon.com/AmazonS3/latest/dev/ObjectsinRequesterPaysBuckets.html
	RequestPayer RequestPayer `location:"header" locationName:"x-amz-request-payer" type:"string" enum:"true"`

	// A token to allow Object Lock to be enabled for an existing bucket.
	Token *string `location:"header" locationName:"x-amz-bucket-object-lock-token" type:"string"`
}

// String returns the string representation
func (s PutObjectLockConfigurationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PutObjectLockConfigurationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PutObjectLockConfigurationInput"}

	if s.Bucket == nil {
		invalidParams.Add(aws.NewErrParamRequired("Bucket"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

func (s *PutObjectLockConfigurationInput) getBucket() (v string) {
	if s.Bucket == nil {
		return v
	}
	return *s.Bucket
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s PutObjectLockConfigurationInput) MarshalFields(e protocol.FieldEncoder) error {

	if len(s.RequestPayer) > 0 {
		v := s.RequestPayer

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-request-payer", v, metadata)
	}
	if s.Token != nil {
		v := *s.Token

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-bucket-object-lock-token", protocol.StringValue(v), metadata)
	}
	if s.Bucket != nil {
		v := *s.Bucket

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "Bucket", protocol.StringValue(v), metadata)
	}
	if s.ObjectLockConfiguration != nil {
		v := s.ObjectLockConfiguration

		metadata := protocol.Metadata{XMLNamespaceURI: "http://s3.amazonaws.com/doc/2006-03-01/"}
		e.SetFields(protocol.PayloadTarget, "ObjectLockConfiguration", v, metadata)
	}
	return nil
}

type PutObjectLockConfigurationOutput struct {
	_ struct{} `type:"structure"`

	// If present, indicates that the requester was successfully charged for the
	// request.
	RequestCharged RequestCharged `location:"header" locationName:"x-amz-request-charged" type:"string" enum:"true"`
}

// String returns the string representation
func (s PutObjectLockConfigurationOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s PutObjectLockConfigurationOutput) MarshalFields(e protocol.FieldEncoder) error {
	if len(s.RequestCharged) > 0 {
		v := s.RequestCharged

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-request-charged", v, metadata)
	}
	return nil
}

const opPutObjectLockConfiguration = "PutObjectLockConfiguration"

// PutObjectLockConfigurationRequest returns a request value for making API operation for
// Amazon Simple Storage Service.
//
// Places an Object Lock configuration on the specified bucket. The rule specified
// in the Object Lock configuration will be applied by default to every new
// object placed in the specified bucket.
//
// DefaultRetention requires either Days or Years. You can't specify both at
// the same time.
//
// Related Resources
//
//    * Locking Objects (https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lock.html)
//
//    // Example sending a request using PutObjectLockConfigurationRequest.
//    req := client.PutObjectLockConfigurationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/s3-2006-03-01/PutObjectLockConfiguration
func (c *Client) PutObjectLockConfigurationRequest(input *PutObjectLockConfigurationInput) PutObjectLockConfigurationRequest {
	op := &aws.Operation{
		Name:       opPutObjectLockConfiguration,
		HTTPMethod: "PUT",
		HTTPPath:   "/{Bucket}?object-lock",
	}

	if input == nil {
		input = &PutObjectLockConfigurationInput{}
	}

	req := c.newRequest(op, input, &PutObjectLockConfigurationOutput{})
	return PutObjectLockConfigurationRequest{Request: req, Input: input, Copy: c.PutObjectLockConfigurationRequest}
}

// PutObjectLockConfigurationRequest is the request type for the
// PutObjectLockConfiguration API operation.
type PutObjectLockConfigurationRequest struct {
	*aws.Request
	Input *PutObjectLockConfigurationInput
	Copy  func(*PutObjectLockConfigurationInput) PutObjectLockConfigurationRequest
}

// Send marshals and sends the PutObjectLockConfiguration API request.
func (r PutObjectLockConfigurationRequest) Send(ctx context.Context) (*PutObjectLockConfigurationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PutObjectLockConfigurationResponse{
		PutObjectLockConfigurationOutput: r.Request.Data.(*PutObjectLockConfigurationOutput),
		response:                         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PutObjectLockConfigurationResponse is the response type for the
// PutObjectLockConfiguration API operation.
type PutObjectLockConfigurationResponse struct {
	*PutObjectLockConfigurationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PutObjectLockConfiguration request.
func (r *PutObjectLockConfigurationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}