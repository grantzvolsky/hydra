/*
 * Ory Oathkeeper API
 *
 * Documentation for all of Ory Oathkeeper's APIs.
 *
 * API version: 1.0.0
 * Contact: hi@ory.sh
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// PluginConfigInterface PluginConfigInterface PluginConfigInterface PluginConfigInterface PluginConfigInterface PluginConfigInterface PluginConfigInterface PluginConfigInterface PluginConfigInterface PluginConfigInterface PluginConfigInterface PluginConfigInterface PluginConfigInterface PluginConfigInterface PluginConfigInterface The interface between Docker and the plugin
type PluginConfigInterface struct {
	// Protocol to use for clients connecting to the plugin.
	ProtocolScheme *string `json:"ProtocolScheme,omitempty"`
	// socket
	Socket string `json:"Socket"`
	// types
	Types []PluginInterfaceType `json:"Types"`
}

// NewPluginConfigInterface instantiates a new PluginConfigInterface object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPluginConfigInterface(socket string, types []PluginInterfaceType) *PluginConfigInterface {
	this := PluginConfigInterface{}
	this.Socket = socket
	this.Types = types
	return &this
}

// NewPluginConfigInterfaceWithDefaults instantiates a new PluginConfigInterface object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPluginConfigInterfaceWithDefaults() *PluginConfigInterface {
	this := PluginConfigInterface{}
	return &this
}

// GetProtocolScheme returns the ProtocolScheme field value if set, zero value otherwise.
func (o *PluginConfigInterface) GetProtocolScheme() string {
	if o == nil || o.ProtocolScheme == nil {
		var ret string
		return ret
	}
	return *o.ProtocolScheme
}

// GetProtocolSchemeOk returns a tuple with the ProtocolScheme field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PluginConfigInterface) GetProtocolSchemeOk() (*string, bool) {
	if o == nil || o.ProtocolScheme == nil {
		return nil, false
	}
	return o.ProtocolScheme, true
}

// HasProtocolScheme returns a boolean if a field has been set.
func (o *PluginConfigInterface) HasProtocolScheme() bool {
	if o != nil && o.ProtocolScheme != nil {
		return true
	}

	return false
}

// SetProtocolScheme gets a reference to the given string and assigns it to the ProtocolScheme field.
func (o *PluginConfigInterface) SetProtocolScheme(v string) {
	o.ProtocolScheme = &v
}

// GetSocket returns the Socket field value
func (o *PluginConfigInterface) GetSocket() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Socket
}

// GetSocketOk returns a tuple with the Socket field value
// and a boolean to check if the value has been set.
func (o *PluginConfigInterface) GetSocketOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Socket, true
}

// SetSocket sets field value
func (o *PluginConfigInterface) SetSocket(v string) {
	o.Socket = v
}

// GetTypes returns the Types field value
func (o *PluginConfigInterface) GetTypes() []PluginInterfaceType {
	if o == nil {
		var ret []PluginInterfaceType
		return ret
	}

	return o.Types
}

// GetTypesOk returns a tuple with the Types field value
// and a boolean to check if the value has been set.
func (o *PluginConfigInterface) GetTypesOk() ([]PluginInterfaceType, bool) {
	if o == nil {
		return nil, false
	}
	return o.Types, true
}

// SetTypes sets field value
func (o *PluginConfigInterface) SetTypes(v []PluginInterfaceType) {
	o.Types = v
}

func (o PluginConfigInterface) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ProtocolScheme != nil {
		toSerialize["ProtocolScheme"] = o.ProtocolScheme
	}
	if true {
		toSerialize["Socket"] = o.Socket
	}
	if true {
		toSerialize["Types"] = o.Types
	}
	return json.Marshal(toSerialize)
}

type NullablePluginConfigInterface struct {
	value *PluginConfigInterface
	isSet bool
}

func (v NullablePluginConfigInterface) Get() *PluginConfigInterface {
	return v.value
}

func (v *NullablePluginConfigInterface) Set(val *PluginConfigInterface) {
	v.value = val
	v.isSet = true
}

func (v NullablePluginConfigInterface) IsSet() bool {
	return v.isSet
}

func (v *NullablePluginConfigInterface) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePluginConfigInterface(val *PluginConfigInterface) *NullablePluginConfigInterface {
	return &NullablePluginConfigInterface{value: val, isSet: true}
}

func (v NullablePluginConfigInterface) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePluginConfigInterface) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
