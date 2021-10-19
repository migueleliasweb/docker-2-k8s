// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// SystemVersion Response of Engine API: GET "/version"
//
//
// swagger:model SystemVersion
type SystemVersion struct {

	// The default (and highest) API version that is supported by the daemon
	//
	// Example: 1.40
	APIVersion string `json:"ApiVersion,omitempty"`

	// The architecture that the daemon is running on
	//
	// Example: amd64
	Arch string `json:"Arch,omitempty"`

	// The date and time that the daemon was compiled.
	//
	// Example: 2020-06-22T15:49:27.000000000+00:00
	BuildTime string `json:"BuildTime,omitempty"`

	// Information about system components
	//
	Components []*ComponentVersion `json:"Components"`

	// Indicates if the daemon is started with experimental features enabled.
	//
	// This field is omitted when empty / false.
	//
	// Example: true
	Experimental bool `json:"Experimental,omitempty"`

	// The Git commit of the source code that was used to build the daemon
	//
	// Example: 48a66213fe
	GitCommit string `json:"GitCommit,omitempty"`

	// The version Go used to compile the daemon, and the version of the Go
	// runtime in use.
	//
	// Example: go1.13.14
	GoVersion string `json:"GoVersion,omitempty"`

	// The kernel version (`uname -r`) that the daemon is running on.
	//
	// This field is omitted when empty.
	//
	// Example: 4.19.76-linuxkit
	KernelVersion string `json:"KernelVersion,omitempty"`

	// The minimum API version that is supported by the daemon
	//
	// Example: 1.12
	MinAPIVersion string `json:"MinAPIVersion,omitempty"`

	// The operating system that the daemon is running on ("linux" or "windows")
	//
	// Example: linux
	Os string `json:"Os,omitempty"`

	// platform
	Platform *SystemVersionPlatform `json:"Platform,omitempty"`

	// The version of the daemon
	// Example: 19.03.12
	Version string `json:"Version,omitempty"`
}

// Validate validates this system version
func (m *SystemVersion) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateComponents(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePlatform(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SystemVersion) validateComponents(formats strfmt.Registry) error {
	if swag.IsZero(m.Components) { // not required
		return nil
	}

	for i := 0; i < len(m.Components); i++ {
		if swag.IsZero(m.Components[i]) { // not required
			continue
		}

		if m.Components[i] != nil {
			if err := m.Components[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Components" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *SystemVersion) validatePlatform(formats strfmt.Registry) error {
	if swag.IsZero(m.Platform) { // not required
		return nil
	}

	if m.Platform != nil {
		if err := m.Platform.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Platform")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this system version based on the context it is used
func (m *SystemVersion) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateComponents(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePlatform(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SystemVersion) contextValidateComponents(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Components); i++ {

		if m.Components[i] != nil {
			if err := m.Components[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Components" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *SystemVersion) contextValidatePlatform(ctx context.Context, formats strfmt.Registry) error {

	if m.Platform != nil {
		if err := m.Platform.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Platform")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SystemVersion) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SystemVersion) UnmarshalBinary(b []byte) error {
	var res SystemVersion
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ComponentVersion component version
//
// swagger:model ComponentVersion
type ComponentVersion struct {

	// Key/value pairs of strings with additional information about the
	// component. These values are intended for informational purposes
	// only, and their content is not defined, and not part of the API
	// specification.
	//
	// These messages can be printed by the client as information to the user.
	//
	Details interface{} `json:"Details,omitempty"`

	// Name of the component
	//
	// Example: Engine
	// Required: true
	Name *string `json:"Name"`

	// Version of the component
	//
	// Example: 19.03.12
	// Required: true
	Version string `json:"Version"`
}

// Validate validates this component version
func (m *ComponentVersion) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVersion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ComponentVersion) validateName(formats strfmt.Registry) error {

	if err := validate.Required("Name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *ComponentVersion) validateVersion(formats strfmt.Registry) error {

	if err := validate.RequiredString("Version", "body", m.Version); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this component version based on context it is used
func (m *ComponentVersion) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ComponentVersion) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ComponentVersion) UnmarshalBinary(b []byte) error {
	var res ComponentVersion
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SystemVersionPlatform system version platform
//
// swagger:model SystemVersionPlatform
type SystemVersionPlatform struct {

	// name
	// Required: true
	Name *string `json:"Name"`
}

// Validate validates this system version platform
func (m *SystemVersionPlatform) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SystemVersionPlatform) validateName(formats strfmt.Registry) error {

	if err := validate.Required("Platform"+"."+"Name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this system version platform based on context it is used
func (m *SystemVersionPlatform) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SystemVersionPlatform) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SystemVersionPlatform) UnmarshalBinary(b []byte) error {
	var res SystemVersionPlatform
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
