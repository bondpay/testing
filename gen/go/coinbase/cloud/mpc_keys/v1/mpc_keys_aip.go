// Code generated by protoc-gen-go-aip. DO NOT EDIT.
//
// versions:
// 	protoc-gen-go-aip development
// 	protoc (unknown)
// source: coinbase/cloud/mpc_keys/v1/mpc_keys.proto

package v1

import (
	fmt "fmt"
	resourcename "go.einride.tech/aip/resourcename"
	strings "strings"
)

type DeviceResourceName struct {
	DeviceId string
}

func (n DeviceResourceName) Validate() error {
	if n.DeviceId == "" {
		return fmt.Errorf("device_id: empty")
	}
	if strings.IndexByte(n.DeviceId, '/') != -1 {
		return fmt.Errorf("device_id: contains illegal character '/'")
	}
	return nil
}

func (n DeviceResourceName) ContainsWildcard() bool {
	return false || n.DeviceId == "-"
}

func (n DeviceResourceName) String() string {
	return resourcename.Sprint(
		"devices/{device_id}",
		n.DeviceId,
	)
}

func (n DeviceResourceName) MarshalString() (string, error) {
	if err := n.Validate(); err != nil {
		return "", err
	}
	return n.String(), nil
}

func (n *DeviceResourceName) UnmarshalString(name string) error {
	return resourcename.Sscan(
		name,
		"devices/{device_id}",
		&n.DeviceId,
	)
}

type DeviceGroupResourceName struct {
	PoolId        string
	DeviceGroupId string
}

func (n DeviceGroupResourceName) Validate() error {
	if n.PoolId == "" {
		return fmt.Errorf("pool_id: empty")
	}
	if strings.IndexByte(n.PoolId, '/') != -1 {
		return fmt.Errorf("pool_id: contains illegal character '/'")
	}
	if n.DeviceGroupId == "" {
		return fmt.Errorf("device_group_id: empty")
	}
	if strings.IndexByte(n.DeviceGroupId, '/') != -1 {
		return fmt.Errorf("device_group_id: contains illegal character '/'")
	}
	return nil
}

func (n DeviceGroupResourceName) ContainsWildcard() bool {
	return false || n.PoolId == "-" || n.DeviceGroupId == "-"
}

func (n DeviceGroupResourceName) String() string {
	return resourcename.Sprint(
		"pools/{pool_id}/deviceGroups/{device_group_id}",
		n.PoolId,
		n.DeviceGroupId,
	)
}

func (n DeviceGroupResourceName) MarshalString() (string, error) {
	if err := n.Validate(); err != nil {
		return "", err
	}
	return n.String(), nil
}

func (n *DeviceGroupResourceName) UnmarshalString(name string) error {
	return resourcename.Sscan(
		name,
		"pools/{pool_id}/deviceGroups/{device_group_id}",
		&n.PoolId,
		&n.DeviceGroupId,
	)
}

type MPCOperationResourceName struct {
	PoolId         string
	DeviceGroupId  string
	MpcOperationId string
}

func (n DeviceGroupResourceName) MPCOperationResourceName(
	mpcOperationId string,
) MPCOperationResourceName {
	return MPCOperationResourceName{
		PoolId:         n.PoolId,
		DeviceGroupId:  n.DeviceGroupId,
		MpcOperationId: mpcOperationId,
	}
}

func (n MPCOperationResourceName) Validate() error {
	if n.PoolId == "" {
		return fmt.Errorf("pool_id: empty")
	}
	if strings.IndexByte(n.PoolId, '/') != -1 {
		return fmt.Errorf("pool_id: contains illegal character '/'")
	}
	if n.DeviceGroupId == "" {
		return fmt.Errorf("device_group_id: empty")
	}
	if strings.IndexByte(n.DeviceGroupId, '/') != -1 {
		return fmt.Errorf("device_group_id: contains illegal character '/'")
	}
	if n.MpcOperationId == "" {
		return fmt.Errorf("mpc_operation_id: empty")
	}
	if strings.IndexByte(n.MpcOperationId, '/') != -1 {
		return fmt.Errorf("mpc_operation_id: contains illegal character '/'")
	}
	return nil
}

func (n MPCOperationResourceName) ContainsWildcard() bool {
	return false || n.PoolId == "-" || n.DeviceGroupId == "-" || n.MpcOperationId == "-"
}

func (n MPCOperationResourceName) String() string {
	return resourcename.Sprint(
		"pools/{pool_id}/deviceGroups/{device_group_id}/mpcOperations/{mpc_operation_id}",
		n.PoolId,
		n.DeviceGroupId,
		n.MpcOperationId,
	)
}

func (n MPCOperationResourceName) MarshalString() (string, error) {
	if err := n.Validate(); err != nil {
		return "", err
	}
	return n.String(), nil
}

func (n *MPCOperationResourceName) UnmarshalString(name string) error {
	return resourcename.Sscan(
		name,
		"pools/{pool_id}/deviceGroups/{device_group_id}/mpcOperations/{mpc_operation_id}",
		&n.PoolId,
		&n.DeviceGroupId,
		&n.MpcOperationId,
	)
}

func (n MPCOperationResourceName) DeviceGroupResourceName() DeviceGroupResourceName {
	return DeviceGroupResourceName{
		PoolId:        n.PoolId,
		DeviceGroupId: n.DeviceGroupId,
	}
}

type MPCKeyResourceName struct {
	PoolId        string
	DeviceGroupId string
	MpcKeyId      string
}

func (n DeviceGroupResourceName) MPCKeyResourceName(
	mpcKeyId string,
) MPCKeyResourceName {
	return MPCKeyResourceName{
		PoolId:        n.PoolId,
		DeviceGroupId: n.DeviceGroupId,
		MpcKeyId:      mpcKeyId,
	}
}

func (n MPCKeyResourceName) Validate() error {
	if n.PoolId == "" {
		return fmt.Errorf("pool_id: empty")
	}
	if strings.IndexByte(n.PoolId, '/') != -1 {
		return fmt.Errorf("pool_id: contains illegal character '/'")
	}
	if n.DeviceGroupId == "" {
		return fmt.Errorf("device_group_id: empty")
	}
	if strings.IndexByte(n.DeviceGroupId, '/') != -1 {
		return fmt.Errorf("device_group_id: contains illegal character '/'")
	}
	if n.MpcKeyId == "" {
		return fmt.Errorf("mpc_key_id: empty")
	}
	if strings.IndexByte(n.MpcKeyId, '/') != -1 {
		return fmt.Errorf("mpc_key_id: contains illegal character '/'")
	}
	return nil
}

func (n MPCKeyResourceName) ContainsWildcard() bool {
	return false || n.PoolId == "-" || n.DeviceGroupId == "-" || n.MpcKeyId == "-"
}

func (n MPCKeyResourceName) String() string {
	return resourcename.Sprint(
		"pools/{pool_id}/deviceGroups/{device_group_id}/mpcKeys/{mpc_key_id}",
		n.PoolId,
		n.DeviceGroupId,
		n.MpcKeyId,
	)
}

func (n MPCKeyResourceName) MarshalString() (string, error) {
	if err := n.Validate(); err != nil {
		return "", err
	}
	return n.String(), nil
}

func (n *MPCKeyResourceName) UnmarshalString(name string) error {
	return resourcename.Sscan(
		name,
		"pools/{pool_id}/deviceGroups/{device_group_id}/mpcKeys/{mpc_key_id}",
		&n.PoolId,
		&n.DeviceGroupId,
		&n.MpcKeyId,
	)
}

func (n MPCKeyResourceName) DeviceGroupResourceName() DeviceGroupResourceName {
	return DeviceGroupResourceName{
		PoolId:        n.PoolId,
		DeviceGroupId: n.DeviceGroupId,
	}
}

type SignatureResourceName struct {
	PoolId        string
	DeviceGroupId string
	MpcKeyId      string
	SignatureId   string
}

func (n DeviceGroupResourceName) SignatureResourceName(
	mpcKeyId string,
	signatureId string,
) SignatureResourceName {
	return SignatureResourceName{
		PoolId:        n.PoolId,
		DeviceGroupId: n.DeviceGroupId,
		MpcKeyId:      mpcKeyId,
		SignatureId:   signatureId,
	}
}

func (n MPCKeyResourceName) SignatureResourceName(
	signatureId string,
) SignatureResourceName {
	return SignatureResourceName{
		PoolId:        n.PoolId,
		DeviceGroupId: n.DeviceGroupId,
		MpcKeyId:      n.MpcKeyId,
		SignatureId:   signatureId,
	}
}

func (n SignatureResourceName) Validate() error {
	if n.PoolId == "" {
		return fmt.Errorf("pool_id: empty")
	}
	if strings.IndexByte(n.PoolId, '/') != -1 {
		return fmt.Errorf("pool_id: contains illegal character '/'")
	}
	if n.DeviceGroupId == "" {
		return fmt.Errorf("device_group_id: empty")
	}
	if strings.IndexByte(n.DeviceGroupId, '/') != -1 {
		return fmt.Errorf("device_group_id: contains illegal character '/'")
	}
	if n.MpcKeyId == "" {
		return fmt.Errorf("mpc_key_id: empty")
	}
	if strings.IndexByte(n.MpcKeyId, '/') != -1 {
		return fmt.Errorf("mpc_key_id: contains illegal character '/'")
	}
	if n.SignatureId == "" {
		return fmt.Errorf("signature_id: empty")
	}
	if strings.IndexByte(n.SignatureId, '/') != -1 {
		return fmt.Errorf("signature_id: contains illegal character '/'")
	}
	return nil
}

func (n SignatureResourceName) ContainsWildcard() bool {
	return false || n.PoolId == "-" || n.DeviceGroupId == "-" || n.MpcKeyId == "-" || n.SignatureId == "-"
}

func (n SignatureResourceName) String() string {
	return resourcename.Sprint(
		"pools/{pool_id}/deviceGroups/{device_group_id}/mpcKeys/{mpc_key_id}/signatures/{signature_id}",
		n.PoolId,
		n.DeviceGroupId,
		n.MpcKeyId,
		n.SignatureId,
	)
}

func (n SignatureResourceName) MarshalString() (string, error) {
	if err := n.Validate(); err != nil {
		return "", err
	}
	return n.String(), nil
}

func (n *SignatureResourceName) UnmarshalString(name string) error {
	return resourcename.Sscan(
		name,
		"pools/{pool_id}/deviceGroups/{device_group_id}/mpcKeys/{mpc_key_id}/signatures/{signature_id}",
		&n.PoolId,
		&n.DeviceGroupId,
		&n.MpcKeyId,
		&n.SignatureId,
	)
}

func (n SignatureResourceName) DeviceGroupResourceName() DeviceGroupResourceName {
	return DeviceGroupResourceName{
		PoolId:        n.PoolId,
		DeviceGroupId: n.DeviceGroupId,
	}
}

func (n SignatureResourceName) MPCKeyResourceName() MPCKeyResourceName {
	return MPCKeyResourceName{
		PoolId:        n.PoolId,
		DeviceGroupId: n.DeviceGroupId,
		MpcKeyId:      n.MpcKeyId,
	}
}
