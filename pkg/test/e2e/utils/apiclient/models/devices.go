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
)

// Devices devices
//
// swagger:model Devices
type Devices struct {

	// Whether to attach the default graphics device or not.
	// VNC will not be available if set to false. Defaults to true.
	AutoattachGraphicsDevice bool `json:"autoattachGraphicsDevice,omitempty"`

	// Whether to attach the Memory balloon device with default period.
	// Period can be adjusted in virt-config.
	// Defaults to true.
	// +optional
	AutoattachMemBalloon bool `json:"autoattachMemBalloon,omitempty"`

	// Whether to attach a pod network interface. Defaults to true.
	AutoattachPodInterface bool `json:"autoattachPodInterface,omitempty"`

	// Whether to attach the default serial console or not.
	// Serial console access will not be available if set to false. Defaults to true.
	AutoattachSerialConsole bool `json:"autoattachSerialConsole,omitempty"`

	// Whether or not to enable virtio multi-queue for block devices.
	// Defaults to false.
	// +optional
	BlockMultiQueue bool `json:"blockMultiQueue,omitempty"`

	// DisableHotplug disabled the ability to hotplug disks.
	DisableHotplug bool `json:"disableHotplug,omitempty"`

	// Disks describes disks, cdroms, floppy and luns which are connected to the vmi.
	Disks []*Disk `json:"disks"`

	// Filesystems describes filesystem which is connected to the vmi.
	// +optional
	// +listType=atomic
	Filesystems []*Filesystem `json:"filesystems"`

	// Whether to attach a GPU device to the vmi.
	// +optional
	// +listType=atomic
	GPUs []*GPU `json:"gpus"`

	// Whether to attach a host device to the vmi.
	// +optional
	// +listType=atomic
	HostDevices []*HostDevice `json:"hostDevices"`

	// Inputs describe input devices
	Inputs []*Input `json:"inputs"`

	// Interfaces describe network interfaces which are added to the vmi.
	Interfaces []*Interface `json:"interfaces"`

	// If specified, virtual network interfaces configured with a virtio bus will also enable the vhost multiqueue feature for network devices. The number of queues created depends on additional factors of the VirtualMachineInstance, like the number of guest CPUs.
	// +optional
	NetworkInterfaceMultiQueue bool `json:"networkInterfaceMultiqueue,omitempty"`

	// Fall back to legacy virtio 0.9 support if virtio bus is selected on devices.
	// This is helpful for old machines like CentOS6 or RHEL6 which
	// do not understand virtio_non_transitional (virtio 1.0).
	UseVirtioTransitional bool `json:"useVirtioTransitional,omitempty"`

	// client passthrough
	ClientPassthrough ClientPassthroughDevices `json:"clientPassthrough,omitempty"`

	// rng
	Rng Rng `json:"rng,omitempty"`

	// sound
	Sound *SoundDevice `json:"sound,omitempty"`

	// watchdog
	Watchdog *Watchdog `json:"watchdog,omitempty"`
}

// Validate validates this devices
func (m *Devices) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDisks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFilesystems(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGPUs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHostDevices(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInputs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInterfaces(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSound(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWatchdog(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Devices) validateDisks(formats strfmt.Registry) error {
	if swag.IsZero(m.Disks) { // not required
		return nil
	}

	for i := 0; i < len(m.Disks); i++ {
		if swag.IsZero(m.Disks[i]) { // not required
			continue
		}

		if m.Disks[i] != nil {
			if err := m.Disks[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("disks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Devices) validateFilesystems(formats strfmt.Registry) error {
	if swag.IsZero(m.Filesystems) { // not required
		return nil
	}

	for i := 0; i < len(m.Filesystems); i++ {
		if swag.IsZero(m.Filesystems[i]) { // not required
			continue
		}

		if m.Filesystems[i] != nil {
			if err := m.Filesystems[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("filesystems" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Devices) validateGPUs(formats strfmt.Registry) error {
	if swag.IsZero(m.GPUs) { // not required
		return nil
	}

	for i := 0; i < len(m.GPUs); i++ {
		if swag.IsZero(m.GPUs[i]) { // not required
			continue
		}

		if m.GPUs[i] != nil {
			if err := m.GPUs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("gpus" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Devices) validateHostDevices(formats strfmt.Registry) error {
	if swag.IsZero(m.HostDevices) { // not required
		return nil
	}

	for i := 0; i < len(m.HostDevices); i++ {
		if swag.IsZero(m.HostDevices[i]) { // not required
			continue
		}

		if m.HostDevices[i] != nil {
			if err := m.HostDevices[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("hostDevices" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Devices) validateInputs(formats strfmt.Registry) error {
	if swag.IsZero(m.Inputs) { // not required
		return nil
	}

	for i := 0; i < len(m.Inputs); i++ {
		if swag.IsZero(m.Inputs[i]) { // not required
			continue
		}

		if m.Inputs[i] != nil {
			if err := m.Inputs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("inputs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Devices) validateInterfaces(formats strfmt.Registry) error {
	if swag.IsZero(m.Interfaces) { // not required
		return nil
	}

	for i := 0; i < len(m.Interfaces); i++ {
		if swag.IsZero(m.Interfaces[i]) { // not required
			continue
		}

		if m.Interfaces[i] != nil {
			if err := m.Interfaces[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("interfaces" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Devices) validateSound(formats strfmt.Registry) error {
	if swag.IsZero(m.Sound) { // not required
		return nil
	}

	if m.Sound != nil {
		if err := m.Sound.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sound")
			}
			return err
		}
	}

	return nil
}

func (m *Devices) validateWatchdog(formats strfmt.Registry) error {
	if swag.IsZero(m.Watchdog) { // not required
		return nil
	}

	if m.Watchdog != nil {
		if err := m.Watchdog.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("watchdog")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this devices based on the context it is used
func (m *Devices) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDisks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFilesystems(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateGPUs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateHostDevices(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateInputs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateInterfaces(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSound(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateWatchdog(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Devices) contextValidateDisks(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Disks); i++ {

		if m.Disks[i] != nil {
			if err := m.Disks[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("disks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Devices) contextValidateFilesystems(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Filesystems); i++ {

		if m.Filesystems[i] != nil {
			if err := m.Filesystems[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("filesystems" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Devices) contextValidateGPUs(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.GPUs); i++ {

		if m.GPUs[i] != nil {
			if err := m.GPUs[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("gpus" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Devices) contextValidateHostDevices(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.HostDevices); i++ {

		if m.HostDevices[i] != nil {
			if err := m.HostDevices[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("hostDevices" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Devices) contextValidateInputs(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Inputs); i++ {

		if m.Inputs[i] != nil {
			if err := m.Inputs[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("inputs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Devices) contextValidateInterfaces(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Interfaces); i++ {

		if m.Interfaces[i] != nil {
			if err := m.Interfaces[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("interfaces" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Devices) contextValidateSound(ctx context.Context, formats strfmt.Registry) error {

	if m.Sound != nil {
		if err := m.Sound.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sound")
			}
			return err
		}
	}

	return nil
}

func (m *Devices) contextValidateWatchdog(ctx context.Context, formats strfmt.Registry) error {

	if m.Watchdog != nil {
		if err := m.Watchdog.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("watchdog")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Devices) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Devices) UnmarshalBinary(b []byte) error {
	var res Devices
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
