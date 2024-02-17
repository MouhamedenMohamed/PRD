// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cells-mailer.proto

package mailer

import (
	fmt "fmt"
	math "math"
	proto "google.golang.org/protobuf/proto"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *User) Validate() error {
	return nil
}
func (this *Mail) Validate() error {
	if this.From != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.From); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("From", err)
		}
	}
	for _, item := range this.To {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("To", err)
			}
		}
	}
	for _, item := range this.Cc {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Cc", err)
			}
		}
	}
	// Validation of proto3 map<> fields is unsupported.
	if this.Sender != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Sender); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Sender", err)
		}
	}
	return nil
}
func (this *SendMailRequest) Validate() error {
	if this.Mail != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Mail); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Mail", err)
		}
	}
	return nil
}
func (this *SendMailResponse) Validate() error {
	return nil
}
func (this *ConsumeQueueRequest) Validate() error {
	return nil
}
func (this *ConsumeQueueResponse) Validate() error {
	return nil
}
