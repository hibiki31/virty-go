/*
VirtyAPI

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 4.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// checks if the NodePage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NodePage{}

// NodePage struct for NodePage
type NodePage struct {
	Name string `json:"name"`
	Description string `json:"description"`
	Domain string `json:"domain"`
	UserName string `json:"userName"`
	Port int32 `json:"port"`
	Core int32 `json:"core"`
	Memory int32 `json:"memory"`
	CpuGen string `json:"cpuGen"`
	OsLike string `json:"osLike"`
	OsName string `json:"osName"`
	OsVersion string `json:"osVersion"`
	Status int32 `json:"status"`
	QemuVersion *string `json:"qemuVersion,omitempty"`
	LibvirtVersion *string `json:"libvirtVersion,omitempty"`
	Roles []NodeRole `json:"roles"`
}

type _NodePage NodePage

// NewNodePage instantiates a new NodePage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNodePage(name string, description string, domain string, userName string, port int32, core int32, memory int32, cpuGen string, osLike string, osName string, osVersion string, status int32, roles []NodeRole) *NodePage {
	this := NodePage{}
	this.Name = name
	this.Description = description
	this.Domain = domain
	this.UserName = userName
	this.Port = port
	this.Core = core
	this.Memory = memory
	this.CpuGen = cpuGen
	this.OsLike = osLike
	this.OsName = osName
	this.OsVersion = osVersion
	this.Status = status
	this.Roles = roles
	return &this
}

// NewNodePageWithDefaults instantiates a new NodePage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNodePageWithDefaults() *NodePage {
	this := NodePage{}
	return &this
}

// GetName returns the Name field value
func (o *NodePage) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *NodePage) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *NodePage) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value
func (o *NodePage) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *NodePage) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *NodePage) SetDescription(v string) {
	o.Description = v
}

// GetDomain returns the Domain field value
func (o *NodePage) GetDomain() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Domain
}

// GetDomainOk returns a tuple with the Domain field value
// and a boolean to check if the value has been set.
func (o *NodePage) GetDomainOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Domain, true
}

// SetDomain sets field value
func (o *NodePage) SetDomain(v string) {
	o.Domain = v
}

// GetUserName returns the UserName field value
func (o *NodePage) GetUserName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserName
}

// GetUserNameOk returns a tuple with the UserName field value
// and a boolean to check if the value has been set.
func (o *NodePage) GetUserNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserName, true
}

// SetUserName sets field value
func (o *NodePage) SetUserName(v string) {
	o.UserName = v
}

// GetPort returns the Port field value
func (o *NodePage) GetPort() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Port
}

// GetPortOk returns a tuple with the Port field value
// and a boolean to check if the value has been set.
func (o *NodePage) GetPortOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Port, true
}

// SetPort sets field value
func (o *NodePage) SetPort(v int32) {
	o.Port = v
}

// GetCore returns the Core field value
func (o *NodePage) GetCore() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Core
}

// GetCoreOk returns a tuple with the Core field value
// and a boolean to check if the value has been set.
func (o *NodePage) GetCoreOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Core, true
}

// SetCore sets field value
func (o *NodePage) SetCore(v int32) {
	o.Core = v
}

// GetMemory returns the Memory field value
func (o *NodePage) GetMemory() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Memory
}

// GetMemoryOk returns a tuple with the Memory field value
// and a boolean to check if the value has been set.
func (o *NodePage) GetMemoryOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Memory, true
}

// SetMemory sets field value
func (o *NodePage) SetMemory(v int32) {
	o.Memory = v
}

// GetCpuGen returns the CpuGen field value
func (o *NodePage) GetCpuGen() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CpuGen
}

// GetCpuGenOk returns a tuple with the CpuGen field value
// and a boolean to check if the value has been set.
func (o *NodePage) GetCpuGenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CpuGen, true
}

// SetCpuGen sets field value
func (o *NodePage) SetCpuGen(v string) {
	o.CpuGen = v
}

// GetOsLike returns the OsLike field value
func (o *NodePage) GetOsLike() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OsLike
}

// GetOsLikeOk returns a tuple with the OsLike field value
// and a boolean to check if the value has been set.
func (o *NodePage) GetOsLikeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OsLike, true
}

// SetOsLike sets field value
func (o *NodePage) SetOsLike(v string) {
	o.OsLike = v
}

// GetOsName returns the OsName field value
func (o *NodePage) GetOsName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OsName
}

// GetOsNameOk returns a tuple with the OsName field value
// and a boolean to check if the value has been set.
func (o *NodePage) GetOsNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OsName, true
}

// SetOsName sets field value
func (o *NodePage) SetOsName(v string) {
	o.OsName = v
}

// GetOsVersion returns the OsVersion field value
func (o *NodePage) GetOsVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OsVersion
}

// GetOsVersionOk returns a tuple with the OsVersion field value
// and a boolean to check if the value has been set.
func (o *NodePage) GetOsVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OsVersion, true
}

// SetOsVersion sets field value
func (o *NodePage) SetOsVersion(v string) {
	o.OsVersion = v
}

// GetStatus returns the Status field value
func (o *NodePage) GetStatus() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *NodePage) GetStatusOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *NodePage) SetStatus(v int32) {
	o.Status = v
}

// GetQemuVersion returns the QemuVersion field value if set, zero value otherwise.
func (o *NodePage) GetQemuVersion() string {
	if o == nil || IsNil(o.QemuVersion) {
		var ret string
		return ret
	}
	return *o.QemuVersion
}

// GetQemuVersionOk returns a tuple with the QemuVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodePage) GetQemuVersionOk() (*string, bool) {
	if o == nil || IsNil(o.QemuVersion) {
		return nil, false
	}
	return o.QemuVersion, true
}

// HasQemuVersion returns a boolean if a field has been set.
func (o *NodePage) HasQemuVersion() bool {
	if o != nil && !IsNil(o.QemuVersion) {
		return true
	}

	return false
}

// SetQemuVersion gets a reference to the given string and assigns it to the QemuVersion field.
func (o *NodePage) SetQemuVersion(v string) {
	o.QemuVersion = &v
}

// GetLibvirtVersion returns the LibvirtVersion field value if set, zero value otherwise.
func (o *NodePage) GetLibvirtVersion() string {
	if o == nil || IsNil(o.LibvirtVersion) {
		var ret string
		return ret
	}
	return *o.LibvirtVersion
}

// GetLibvirtVersionOk returns a tuple with the LibvirtVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodePage) GetLibvirtVersionOk() (*string, bool) {
	if o == nil || IsNil(o.LibvirtVersion) {
		return nil, false
	}
	return o.LibvirtVersion, true
}

// HasLibvirtVersion returns a boolean if a field has been set.
func (o *NodePage) HasLibvirtVersion() bool {
	if o != nil && !IsNil(o.LibvirtVersion) {
		return true
	}

	return false
}

// SetLibvirtVersion gets a reference to the given string and assigns it to the LibvirtVersion field.
func (o *NodePage) SetLibvirtVersion(v string) {
	o.LibvirtVersion = &v
}

// GetRoles returns the Roles field value
func (o *NodePage) GetRoles() []NodeRole {
	if o == nil {
		var ret []NodeRole
		return ret
	}

	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value
// and a boolean to check if the value has been set.
func (o *NodePage) GetRolesOk() ([]NodeRole, bool) {
	if o == nil {
		return nil, false
	}
	return o.Roles, true
}

// SetRoles sets field value
func (o *NodePage) SetRoles(v []NodeRole) {
	o.Roles = v
}

func (o NodePage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NodePage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["description"] = o.Description
	toSerialize["domain"] = o.Domain
	toSerialize["userName"] = o.UserName
	toSerialize["port"] = o.Port
	toSerialize["core"] = o.Core
	toSerialize["memory"] = o.Memory
	toSerialize["cpuGen"] = o.CpuGen
	toSerialize["osLike"] = o.OsLike
	toSerialize["osName"] = o.OsName
	toSerialize["osVersion"] = o.OsVersion
	toSerialize["status"] = o.Status
	if !IsNil(o.QemuVersion) {
		toSerialize["qemuVersion"] = o.QemuVersion
	}
	if !IsNil(o.LibvirtVersion) {
		toSerialize["libvirtVersion"] = o.LibvirtVersion
	}
	toSerialize["roles"] = o.Roles
	return toSerialize, nil
}

func (o *NodePage) UnmarshalJSON(bytes []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"description",
		"domain",
		"userName",
		"port",
		"core",
		"memory",
		"cpuGen",
		"osLike",
		"osName",
		"osVersion",
		"status",
		"roles",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varNodePage := _NodePage{}

	err = json.Unmarshal(bytes, &varNodePage)

	if err != nil {
		return err
	}

	*o = NodePage(varNodePage)

	return err
}

type NullableNodePage struct {
	value *NodePage
	isSet bool
}

func (v NullableNodePage) Get() *NodePage {
	return v.value
}

func (v *NullableNodePage) Set(val *NodePage) {
	v.value = val
	v.isSet = true
}

func (v NullableNodePage) IsSet() bool {
	return v.isSet
}

func (v *NullableNodePage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNodePage(val *NodePage) *NullableNodePage {
	return &NullableNodePage{value: val, isSet: true}
}

func (v NullableNodePage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNodePage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


