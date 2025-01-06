/*
Logto API references

API references for Logto services.  Note: The documentation is for Logto Cloud. If you are using Logto OSS, please refer to the response of `/api/swagger.json` endpoint on your Logto instance.

API version: Cloud
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package logto

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the CreateHook201Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateHook201Response{}

// CreateHook201Response struct for CreateHook201Response
type CreateHook201Response struct {
	TenantId string `json:"tenantId"`
	Id string `json:"id"`
	Name string `json:"name"`
	Event NullableString `json:"event"`
	Events []string `json:"events"`
	Config ListHooks200ResponseInnerConfig `json:"config"`
	SigningKey string `json:"signingKey"`
	Enabled bool `json:"enabled"`
	CreatedAt float32 `json:"createdAt"`
}

type _CreateHook201Response CreateHook201Response

// NewCreateHook201Response instantiates a new CreateHook201Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateHook201Response(tenantId string, id string, name string, event NullableString, events []string, config ListHooks200ResponseInnerConfig, signingKey string, enabled bool, createdAt float32) *CreateHook201Response {
	this := CreateHook201Response{}
	this.TenantId = tenantId
	this.Id = id
	this.Name = name
	this.Event = event
	this.Events = events
	this.Config = config
	this.SigningKey = signingKey
	this.Enabled = enabled
	this.CreatedAt = createdAt
	return &this
}

// NewCreateHook201ResponseWithDefaults instantiates a new CreateHook201Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateHook201ResponseWithDefaults() *CreateHook201Response {
	this := CreateHook201Response{}
	return &this
}

// GetTenantId returns the TenantId field value
func (o *CreateHook201Response) GetTenantId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value
// and a boolean to check if the value has been set.
func (o *CreateHook201Response) GetTenantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TenantId, true
}

// SetTenantId sets field value
func (o *CreateHook201Response) SetTenantId(v string) {
	o.TenantId = v
}

// GetId returns the Id field value
func (o *CreateHook201Response) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CreateHook201Response) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CreateHook201Response) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *CreateHook201Response) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateHook201Response) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateHook201Response) SetName(v string) {
	o.Name = v
}

// GetEvent returns the Event field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CreateHook201Response) GetEvent() string {
	if o == nil || o.Event.Get() == nil {
		var ret string
		return ret
	}

	return *o.Event.Get()
}

// GetEventOk returns a tuple with the Event field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateHook201Response) GetEventOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Event.Get(), o.Event.IsSet()
}

// SetEvent sets field value
func (o *CreateHook201Response) SetEvent(v string) {
	o.Event.Set(&v)
}

// GetEvents returns the Events field value
func (o *CreateHook201Response) GetEvents() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Events
}

// GetEventsOk returns a tuple with the Events field value
// and a boolean to check if the value has been set.
func (o *CreateHook201Response) GetEventsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Events, true
}

// SetEvents sets field value
func (o *CreateHook201Response) SetEvents(v []string) {
	o.Events = v
}

// GetConfig returns the Config field value
func (o *CreateHook201Response) GetConfig() ListHooks200ResponseInnerConfig {
	if o == nil {
		var ret ListHooks200ResponseInnerConfig
		return ret
	}

	return o.Config
}

// GetConfigOk returns a tuple with the Config field value
// and a boolean to check if the value has been set.
func (o *CreateHook201Response) GetConfigOk() (*ListHooks200ResponseInnerConfig, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Config, true
}

// SetConfig sets field value
func (o *CreateHook201Response) SetConfig(v ListHooks200ResponseInnerConfig) {
	o.Config = v
}

// GetSigningKey returns the SigningKey field value
func (o *CreateHook201Response) GetSigningKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SigningKey
}

// GetSigningKeyOk returns a tuple with the SigningKey field value
// and a boolean to check if the value has been set.
func (o *CreateHook201Response) GetSigningKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SigningKey, true
}

// SetSigningKey sets field value
func (o *CreateHook201Response) SetSigningKey(v string) {
	o.SigningKey = v
}

// GetEnabled returns the Enabled field value
func (o *CreateHook201Response) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *CreateHook201Response) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *CreateHook201Response) SetEnabled(v bool) {
	o.Enabled = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *CreateHook201Response) GetCreatedAt() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *CreateHook201Response) GetCreatedAtOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *CreateHook201Response) SetCreatedAt(v float32) {
	o.CreatedAt = v
}

func (o CreateHook201Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateHook201Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["tenantId"] = o.TenantId
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["event"] = o.Event.Get()
	toSerialize["events"] = o.Events
	toSerialize["config"] = o.Config
	toSerialize["signingKey"] = o.SigningKey
	toSerialize["enabled"] = o.Enabled
	toSerialize["createdAt"] = o.CreatedAt
	return toSerialize, nil
}

func (o *CreateHook201Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"tenantId",
		"id",
		"name",
		"event",
		"events",
		"config",
		"signingKey",
		"enabled",
		"createdAt",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varCreateHook201Response := _CreateHook201Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateHook201Response)

	if err != nil {
		return err
	}

	*o = CreateHook201Response(varCreateHook201Response)

	return err
}

type NullableCreateHook201Response struct {
	value *CreateHook201Response
	isSet bool
}

func (v NullableCreateHook201Response) Get() *CreateHook201Response {
	return v.value
}

func (v *NullableCreateHook201Response) Set(val *CreateHook201Response) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateHook201Response) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateHook201Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateHook201Response(val *CreateHook201Response) *NullableCreateHook201Response {
	return &NullableCreateHook201Response{value: val, isSet: true}
}

func (v NullableCreateHook201Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateHook201Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


