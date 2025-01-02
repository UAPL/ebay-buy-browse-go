/*
Browse API

The Browse API has the following resources:<ul><li><b>item_summary:</b><br>Allows shoppers to search for specific items by keyword, GTIN, category, charity, product, image, or item aspects and refine the results by using filters, such as aspects, compatibility, and fields values, or UI parameters.</li><li><b>item:</b><br>Allows shoppers to retrieve the details of a specific item or all items in an item group, which is an item with variations such as color and size and check if a product is compatible with the specified item, such as if a specific car is compatible with a specific part.<br><br>This resource also provides a bridge between the eBay legacy APIs, such as the <a href=\"/api-docs/user-guides/static/finding-user-guide-landing.html\" target=\"_blank\">Finding</b>, and the RESTful APIs, which use different formats for the item IDs.</li></ul>The <b>item_summary</b>, <b>search_by_image</b>, and <b>item</b> resource calls require an <a href=\"/api-docs/static/oauth-client-credentials-grant.html\" target=\"_blank\">Application access token</a>.

API version: v1.19.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package buybrowse

import (
	"encoding/json"
)

// checks if the CompatibilityResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CompatibilityResponse{}

// CompatibilityResponse The type that defines the response fields for <b> checkCompatibility</b>.  
type CompatibilityResponse struct {
	// An enumeration value that tells you if the item is compatible with the product. <br><br>The values are: <ul>   <li>   <b> COMPATIBLE</b> - Indicates the item is compatible with the product specified in the request.</li>   <li>   <b> NOT_COMPATIBLE</b> - Indicates the item is not compatible with the product specified in the request. Be sure to check all the <b> value</b> fields to ensure they are correct as errors in the value can also cause this response.</li>   <li> <b> UNDETERMINED</b> - Indicates one or more attributes for the specified product are missing so compatibility cannot be determined.  The response returns the attributes that are missing.</li>  </ul>  Code so that your app gracefully handles any future changes to this list. For implementation help, refer to <a href='https://developer.ebay.com/api-docs/buy/browse/types/gct:CompatibilityStatus'>eBay API documentation</a>
	CompatibilityStatus *string `json:"compatibilityStatus,omitempty"`
	// An array of warning messages. These types of errors do not prevent the method from executing but should be checked.
	Warnings []Error `json:"warnings,omitempty"`
}

// NewCompatibilityResponse instantiates a new CompatibilityResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCompatibilityResponse() *CompatibilityResponse {
	this := CompatibilityResponse{}
	return &this
}

// NewCompatibilityResponseWithDefaults instantiates a new CompatibilityResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCompatibilityResponseWithDefaults() *CompatibilityResponse {
	this := CompatibilityResponse{}
	return &this
}

// GetCompatibilityStatus returns the CompatibilityStatus field value if set, zero value otherwise.
func (o *CompatibilityResponse) GetCompatibilityStatus() string {
	if o == nil || IsNil(o.CompatibilityStatus) {
		var ret string
		return ret
	}
	return *o.CompatibilityStatus
}

// GetCompatibilityStatusOk returns a tuple with the CompatibilityStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompatibilityResponse) GetCompatibilityStatusOk() (*string, bool) {
	if o == nil || IsNil(o.CompatibilityStatus) {
		return nil, false
	}
	return o.CompatibilityStatus, true
}

// HasCompatibilityStatus returns a boolean if a field has been set.
func (o *CompatibilityResponse) HasCompatibilityStatus() bool {
	if o != nil && !IsNil(o.CompatibilityStatus) {
		return true
	}

	return false
}

// SetCompatibilityStatus gets a reference to the given string and assigns it to the CompatibilityStatus field.
func (o *CompatibilityResponse) SetCompatibilityStatus(v string) {
	o.CompatibilityStatus = &v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *CompatibilityResponse) GetWarnings() []Error {
	if o == nil || IsNil(o.Warnings) {
		var ret []Error
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompatibilityResponse) GetWarningsOk() ([]Error, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *CompatibilityResponse) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []Error and assigns it to the Warnings field.
func (o *CompatibilityResponse) SetWarnings(v []Error) {
	o.Warnings = v
}

func (o CompatibilityResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CompatibilityResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CompatibilityStatus) {
		toSerialize["compatibilityStatus"] = o.CompatibilityStatus
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

type NullableCompatibilityResponse struct {
	value *CompatibilityResponse
	isSet bool
}

func (v NullableCompatibilityResponse) Get() *CompatibilityResponse {
	return v.value
}

func (v *NullableCompatibilityResponse) Set(val *CompatibilityResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCompatibilityResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCompatibilityResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCompatibilityResponse(val *CompatibilityResponse) *NullableCompatibilityResponse {
	return &NullableCompatibilityResponse{value: val, isSet: true}
}

func (v NullableCompatibilityResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCompatibilityResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


