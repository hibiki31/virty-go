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

// checks if the DomainDetail type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DomainDetail{}

// DomainDetail struct for DomainDetail
type DomainDetail struct {
	Uuid string `json:"uuid"`
	Name string `json:"name"`
	Core int32 `json:"core"`
	Memory int32 `json:"memory"`
	Status int32 `json:"status"`
	Description *string `json:"description,omitempty"`
	NodeName string `json:"nodeName"`
	OwnerUserId *string `json:"ownerUserId,omitempty"`
	OwnerProjectId *string `json:"ownerProjectId,omitempty"`
	OwnerProject *DomainProject `json:"ownerProject,omitempty"`
	VncPort *int32 `json:"vncPort,omitempty"`
	VncPassword *string `json:"vncPassword,omitempty"`
	Drives []DomainDrive `json:"drives,omitempty"`
	Interfaces []DomainInterface `json:"interfaces,omitempty"`
	Node Node `json:"node"`
}

type _DomainDetail DomainDetail

// NewDomainDetail instantiates a new DomainDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDomainDetail(uuid string, name string, core int32, memory int32, status int32, nodeName string, node Node) *DomainDetail {
	this := DomainDetail{}
	this.Uuid = uuid
	this.Name = name
	this.Core = core
	this.Memory = memory
	this.Status = status
	this.NodeName = nodeName
	this.Node = node
	return &this
}

// NewDomainDetailWithDefaults instantiates a new DomainDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDomainDetailWithDefaults() *DomainDetail {
	this := DomainDetail{}
	return &this
}

// GetUuid returns the Uuid field value
func (o *DomainDetail) GetUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value
// and a boolean to check if the value has been set.
func (o *DomainDetail) GetUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uuid, true
}

// SetUuid sets field value
func (o *DomainDetail) SetUuid(v string) {
	o.Uuid = v
}

// GetName returns the Name field value
func (o *DomainDetail) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DomainDetail) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DomainDetail) SetName(v string) {
	o.Name = v
}

// GetCore returns the Core field value
func (o *DomainDetail) GetCore() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Core
}

// GetCoreOk returns a tuple with the Core field value
// and a boolean to check if the value has been set.
func (o *DomainDetail) GetCoreOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Core, true
}

// SetCore sets field value
func (o *DomainDetail) SetCore(v int32) {
	o.Core = v
}

// GetMemory returns the Memory field value
func (o *DomainDetail) GetMemory() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Memory
}

// GetMemoryOk returns a tuple with the Memory field value
// and a boolean to check if the value has been set.
func (o *DomainDetail) GetMemoryOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Memory, true
}

// SetMemory sets field value
func (o *DomainDetail) SetMemory(v int32) {
	o.Memory = v
}

// GetStatus returns the Status field value
func (o *DomainDetail) GetStatus() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *DomainDetail) GetStatusOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *DomainDetail) SetStatus(v int32) {
	o.Status = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *DomainDetail) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainDetail) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *DomainDetail) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *DomainDetail) SetDescription(v string) {
	o.Description = &v
}

// GetNodeName returns the NodeName field value
func (o *DomainDetail) GetNodeName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NodeName
}

// GetNodeNameOk returns a tuple with the NodeName field value
// and a boolean to check if the value has been set.
func (o *DomainDetail) GetNodeNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NodeName, true
}

// SetNodeName sets field value
func (o *DomainDetail) SetNodeName(v string) {
	o.NodeName = v
}

// GetOwnerUserId returns the OwnerUserId field value if set, zero value otherwise.
func (o *DomainDetail) GetOwnerUserId() string {
	if o == nil || IsNil(o.OwnerUserId) {
		var ret string
		return ret
	}
	return *o.OwnerUserId
}

// GetOwnerUserIdOk returns a tuple with the OwnerUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainDetail) GetOwnerUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.OwnerUserId) {
		return nil, false
	}
	return o.OwnerUserId, true
}

// HasOwnerUserId returns a boolean if a field has been set.
func (o *DomainDetail) HasOwnerUserId() bool {
	if o != nil && !IsNil(o.OwnerUserId) {
		return true
	}

	return false
}

// SetOwnerUserId gets a reference to the given string and assigns it to the OwnerUserId field.
func (o *DomainDetail) SetOwnerUserId(v string) {
	o.OwnerUserId = &v
}

// GetOwnerProjectId returns the OwnerProjectId field value if set, zero value otherwise.
func (o *DomainDetail) GetOwnerProjectId() string {
	if o == nil || IsNil(o.OwnerProjectId) {
		var ret string
		return ret
	}
	return *o.OwnerProjectId
}

// GetOwnerProjectIdOk returns a tuple with the OwnerProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainDetail) GetOwnerProjectIdOk() (*string, bool) {
	if o == nil || IsNil(o.OwnerProjectId) {
		return nil, false
	}
	return o.OwnerProjectId, true
}

// HasOwnerProjectId returns a boolean if a field has been set.
func (o *DomainDetail) HasOwnerProjectId() bool {
	if o != nil && !IsNil(o.OwnerProjectId) {
		return true
	}

	return false
}

// SetOwnerProjectId gets a reference to the given string and assigns it to the OwnerProjectId field.
func (o *DomainDetail) SetOwnerProjectId(v string) {
	o.OwnerProjectId = &v
}

// GetOwnerProject returns the OwnerProject field value if set, zero value otherwise.
func (o *DomainDetail) GetOwnerProject() DomainProject {
	if o == nil || IsNil(o.OwnerProject) {
		var ret DomainProject
		return ret
	}
	return *o.OwnerProject
}

// GetOwnerProjectOk returns a tuple with the OwnerProject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainDetail) GetOwnerProjectOk() (*DomainProject, bool) {
	if o == nil || IsNil(o.OwnerProject) {
		return nil, false
	}
	return o.OwnerProject, true
}

// HasOwnerProject returns a boolean if a field has been set.
func (o *DomainDetail) HasOwnerProject() bool {
	if o != nil && !IsNil(o.OwnerProject) {
		return true
	}

	return false
}

// SetOwnerProject gets a reference to the given DomainProject and assigns it to the OwnerProject field.
func (o *DomainDetail) SetOwnerProject(v DomainProject) {
	o.OwnerProject = &v
}

// GetVncPort returns the VncPort field value if set, zero value otherwise.
func (o *DomainDetail) GetVncPort() int32 {
	if o == nil || IsNil(o.VncPort) {
		var ret int32
		return ret
	}
	return *o.VncPort
}

// GetVncPortOk returns a tuple with the VncPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainDetail) GetVncPortOk() (*int32, bool) {
	if o == nil || IsNil(o.VncPort) {
		return nil, false
	}
	return o.VncPort, true
}

// HasVncPort returns a boolean if a field has been set.
func (o *DomainDetail) HasVncPort() bool {
	if o != nil && !IsNil(o.VncPort) {
		return true
	}

	return false
}

// SetVncPort gets a reference to the given int32 and assigns it to the VncPort field.
func (o *DomainDetail) SetVncPort(v int32) {
	o.VncPort = &v
}

// GetVncPassword returns the VncPassword field value if set, zero value otherwise.
func (o *DomainDetail) GetVncPassword() string {
	if o == nil || IsNil(o.VncPassword) {
		var ret string
		return ret
	}
	return *o.VncPassword
}

// GetVncPasswordOk returns a tuple with the VncPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainDetail) GetVncPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.VncPassword) {
		return nil, false
	}
	return o.VncPassword, true
}

// HasVncPassword returns a boolean if a field has been set.
func (o *DomainDetail) HasVncPassword() bool {
	if o != nil && !IsNil(o.VncPassword) {
		return true
	}

	return false
}

// SetVncPassword gets a reference to the given string and assigns it to the VncPassword field.
func (o *DomainDetail) SetVncPassword(v string) {
	o.VncPassword = &v
}

// GetDrives returns the Drives field value if set, zero value otherwise.
func (o *DomainDetail) GetDrives() []DomainDrive {
	if o == nil || IsNil(o.Drives) {
		var ret []DomainDrive
		return ret
	}
	return o.Drives
}

// GetDrivesOk returns a tuple with the Drives field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainDetail) GetDrivesOk() ([]DomainDrive, bool) {
	if o == nil || IsNil(o.Drives) {
		return nil, false
	}
	return o.Drives, true
}

// HasDrives returns a boolean if a field has been set.
func (o *DomainDetail) HasDrives() bool {
	if o != nil && !IsNil(o.Drives) {
		return true
	}

	return false
}

// SetDrives gets a reference to the given []DomainDrive and assigns it to the Drives field.
func (o *DomainDetail) SetDrives(v []DomainDrive) {
	o.Drives = v
}

// GetInterfaces returns the Interfaces field value if set, zero value otherwise.
func (o *DomainDetail) GetInterfaces() []DomainInterface {
	if o == nil || IsNil(o.Interfaces) {
		var ret []DomainInterface
		return ret
	}
	return o.Interfaces
}

// GetInterfacesOk returns a tuple with the Interfaces field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainDetail) GetInterfacesOk() ([]DomainInterface, bool) {
	if o == nil || IsNil(o.Interfaces) {
		return nil, false
	}
	return o.Interfaces, true
}

// HasInterfaces returns a boolean if a field has been set.
func (o *DomainDetail) HasInterfaces() bool {
	if o != nil && !IsNil(o.Interfaces) {
		return true
	}

	return false
}

// SetInterfaces gets a reference to the given []DomainInterface and assigns it to the Interfaces field.
func (o *DomainDetail) SetInterfaces(v []DomainInterface) {
	o.Interfaces = v
}

// GetNode returns the Node field value
func (o *DomainDetail) GetNode() Node {
	if o == nil {
		var ret Node
		return ret
	}

	return o.Node
}

// GetNodeOk returns a tuple with the Node field value
// and a boolean to check if the value has been set.
func (o *DomainDetail) GetNodeOk() (*Node, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Node, true
}

// SetNode sets field value
func (o *DomainDetail) SetNode(v Node) {
	o.Node = v
}

func (o DomainDetail) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DomainDetail) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["uuid"] = o.Uuid
	toSerialize["name"] = o.Name
	toSerialize["core"] = o.Core
	toSerialize["memory"] = o.Memory
	toSerialize["status"] = o.Status
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["nodeName"] = o.NodeName
	if !IsNil(o.OwnerUserId) {
		toSerialize["ownerUserId"] = o.OwnerUserId
	}
	if !IsNil(o.OwnerProjectId) {
		toSerialize["ownerProjectId"] = o.OwnerProjectId
	}
	if !IsNil(o.OwnerProject) {
		toSerialize["ownerProject"] = o.OwnerProject
	}
	if !IsNil(o.VncPort) {
		toSerialize["vncPort"] = o.VncPort
	}
	if !IsNil(o.VncPassword) {
		toSerialize["vncPassword"] = o.VncPassword
	}
	if !IsNil(o.Drives) {
		toSerialize["drives"] = o.Drives
	}
	if !IsNil(o.Interfaces) {
		toSerialize["interfaces"] = o.Interfaces
	}
	toSerialize["node"] = o.Node
	return toSerialize, nil
}

func (o *DomainDetail) UnmarshalJSON(bytes []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"uuid",
		"name",
		"core",
		"memory",
		"status",
		"nodeName",
		"node",
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

	varDomainDetail := _DomainDetail{}

	err = json.Unmarshal(bytes, &varDomainDetail)

	if err != nil {
		return err
	}

	*o = DomainDetail(varDomainDetail)

	return err
}

type NullableDomainDetail struct {
	value *DomainDetail
	isSet bool
}

func (v NullableDomainDetail) Get() *DomainDetail {
	return v.value
}

func (v *NullableDomainDetail) Set(val *DomainDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableDomainDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableDomainDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDomainDetail(val *DomainDetail) *NullableDomainDetail {
	return &NullableDomainDetail{value: val, isSet: true}
}

func (v NullableDomainDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDomainDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


