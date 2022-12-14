// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iam

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/iam-2010-05-08/GetSSHPublicKeyRequest
type GetSSHPublicKeyInput struct {
	_ struct{} `type:"structure"`

	// Specifies the public key encoding format to use in the response. To retrieve
	// the public key in ssh-rsa Some-Sha
	// PEM format, use PEM.
	//
	// Encoding is a required field
	Encoding EncodingType `type:"string" required:"true" enum:"true"`

	// The unique identifier for the SSH public key.
	//
	// This parameter allows (through its regex pattern (http://wikipedia.org/wiki/regex))
	// a string of characters that can consist of any upper or lowercased letter
	// or digit.
	//
	// SSHPublicKeyId is a required field
	SSHPublicKeyId *string `min:"20" type:"string" required:"true"`

	// The name of the IAM user associated with the SSH public key.
	//
	// This parameter allows (through its regex pattern (http://wikipedia.org/wiki/regex))
	// a string of characters consisting of upper and lowercase alphanumeric characters
	// with no spaces. You can also include any of the following characters: _+=,.@-
	//
	// UserName is a required field
	UserName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s GetSSHPublicKeyInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetSSHPublicKeyInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetSSHPublicKeyInput"}
	if len(s.Encoding) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("Encoding"))
	}

	if s.SSHPublicKeyId == nil {
		invalidParams.Add(aws.NewErrParamRequired("SSHPublicKeyId"))
	}
	if s.SSHPublicKeyId != nil && len(*s.SSHPublicKeyId) < 20 {
		invalidParams.Add(aws.NewErrParamMinLen("SSHPublicKeyId", 20))
	}

	if s.UserName == nil {
		invalidParams.Add(aws.NewErrParamRequired("UserName"))
	}
	if s.UserName != nil && len(*s.UserName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("UserName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Contains the response to a successful GetSSHPublicKey request.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iam-2010-05-08/GetSSHPublicKeyResponse
type GetSSHPublicKeyOutput struct {
	_ struct{} `type:"structure"`

	// A structure containing details about the SSH public key.
	SSHPublicKey *SSHPublicKey `type:"structure"`
}

// String returns the string representation
func (s GetSSHPublicKeyOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetSSHPublicKey = "GetSSHPublicKey"

// GetSSHPublicKeyRequest returns a request value for making API operation for
// AWS Identity and Access Management.
//
// Retrieves the specified SSH public key, including metadata about the key.
//
// The SSH public key retrieved by this operation is used only for authenticating
// the associated IAM user to an AWS CodeCommit repository. For more information
// about using SSH keys to authenticate to an AWS CodeCommit repository, see
// Set up AWS CodeCommit for SSH Connections (https://docs.aws.amazon.com/codecommit/latest/userguide/setting-up-credentials-ssh.html)
// in the AWS CodeCommit User Guide.
//
//    // Example sending a request using GetSSHPublicKeyRequest.
//    req := client.GetSSHPublicKeyRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iam-2010-05-08/GetSSHPublicKey
func (c *Client) GetSSHPublicKeyRequest(input *GetSSHPublicKeyInput) GetSSHPublicKeyRequest {
	op := &aws.Operation{
		Name:       opGetSSHPublicKey,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetSSHPublicKeyInput{}
	}

	req := c.newRequest(op, input, &GetSSHPublicKeyOutput{})
	return GetSSHPublicKeyRequest{Request: req, Input: input, Copy: c.GetSSHPublicKeyRequest}
}

// GetSSHPublicKeyRequest is the request type for the
// GetSSHPublicKey API operation.
type GetSSHPublicKeyRequest struct {
	*aws.Request
	Input *GetSSHPublicKeyInput
	Copy  func(*GetSSHPublicKeyInput) GetSSHPublicKeyRequest
}

// Send marshals and sends the GetSSHPublicKey API request.
func (r GetSSHPublicKeyRequest) Send(ctx context.Context) (*GetSSHPublicKeyResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetSSHPublicKeyResponse{
		GetSSHPublicKeyOutput: r.Request.Data.(*GetSSHPublicKeyOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetSSHPublicKeyResponse is the response type for the
// GetSSHPublicKey API operation.
type GetSSHPublicKeyResponse struct {
	*GetSSHPublicKeyOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetSSHPublicKey request.
func (r *GetSSHPublicKeyResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
