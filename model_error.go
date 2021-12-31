/*
Browse API

<p>The Browse API has the following resources:</p>   <ul> <li><b> item_summary: </b> Lets shoppers search for specific items by keyword, GTIN, category, charity, product, or item aspects and refine the results by using filters, such as aspects, compatibility, and fields values.</li>  <li><b> search_by_image: </b><a href=\"https://developer.ebay.com/api-docs/static/versioning.html#experimental\" target=\"_blank\"><img src=\"/cms/img/docs/experimental-icon.svg\" class=\"legend-icon experimental-icon\" alt=\"Experimental Release\" title=\"Experimental Release\" />&nbsp;(Experimental)</a> Lets shoppers search for specific items by image. You can refine the results by using URI parameters and filters.</li>   <li><b> item: </b> <ul><li>Lets you retrieve the details of a specific item or all the items in an item group, which is an item with variations such as color and size and check if a product is compatible with the specified item, such as if a specific car is compatible with a specific part.</li> <li>Provides a bridge between the eBay legacy APIs, such as <b> Finding</b>, and the RESTful APIs, which use different formats for the item IDs.</li>  </ul> </li>  <li> <b> shopping_cart: </b> <a href=\"https://developer.ebay.com/api-docs/static/versioning.html#experimental\" target=\"_blank\"><img src=\"/cms/img/docs/experimental-icon.svg\" class=\"legend-icon experimental-icon\" alt=\"Experimental Release\" title=\"Experimental Release\" />&nbsp;(Experimental)</a> <a href=\"https://developer.ebay.com/api-docs/static/versioning.html#limited\" target=\"_blank\"> <img src=\"/cms/img/docs/partners-api.svg\" class=\"legend-icon partners-icon\" title=\"Limited Release\"  alt=\"Limited Release\" />(Limited Release)</a> Provides the ability for eBay members to see the contents of their eBay cart, and add, remove, and change the quantity of items in their eBay cart.&nbsp;&nbsp;<b> Note: </b> This resource is not available in the eBay API Explorer.</li></ul>       <p>The <b> item_summary</b>, <b> search_by_image</b>, and <b> item</b> resource calls require an <a href=\"/api-docs/static/oauth-client-credentials-grant.html\">Application access token</a>. The <b> shopping_cart</b> resource calls require a <a href=\"/api-docs/static/oauth-authorization-code-grant.html\">User access token</a>.</p>

API version: v1.11.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package buybrowse

import (
	"encoding/json"
)

// Error The type that defines the fields that can be returned in an error.
type Error struct {
	// This string value indicates the error category. There are three categories of errors: request errors, application errors, and system errors. 
	Category *string `json:"category,omitempty"`
	// The name of the primary system where the error occurred. This is relevant for application errors.
	Domain *string `json:"domain,omitempty"`
	// A unique code that identifies the particular error or warning that occurred. Your application can use error codes as identifiers in your customized error-handling algorithms.
	ErrorId *int32 `json:"errorId,omitempty"`
	// An array of reference IDs that identify the specific request elements most closely associated to the error or warning, if any.
	InputRefIds *[]string `json:"inputRefIds,omitempty"`
	// A detailed description of the condition that caused the error or warning, and information on what to do to correct the problem.
	LongMessage *string `json:"longMessage,omitempty"`
	// A description of the condition that caused the error or warning.
	Message *string `json:"message,omitempty"`
	// An array of reference IDs that identify the specific response elements most closely associated to the error or warning, if any.
	OutputRefIds *[]string `json:"outputRefIds,omitempty"`
	// An array of warning and error messages that return one or more variables contextual information about the error or warning. This is often the field or value that triggered the error or warning.
	Parameters *[]ErrorParameter `json:"parameters,omitempty"`
	// The name of the subdomain in which the error or warning occurred.
	Subdomain *string `json:"subdomain,omitempty"`
}

// NewError instantiates a new Error object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewError() *Error {
	this := Error{}
	return &this
}

// NewErrorWithDefaults instantiates a new Error object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewErrorWithDefaults() *Error {
	this := Error{}
	return &this
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *Error) GetCategory() string {
	if o == nil || o.Category == nil {
		var ret string
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Error) GetCategoryOk() (*string, bool) {
	if o == nil || o.Category == nil {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *Error) HasCategory() bool {
	if o != nil && o.Category != nil {
		return true
	}

	return false
}

// SetCategory gets a reference to the given string and assigns it to the Category field.
func (o *Error) SetCategory(v string) {
	o.Category = &v
}

// GetDomain returns the Domain field value if set, zero value otherwise.
func (o *Error) GetDomain() string {
	if o == nil || o.Domain == nil {
		var ret string
		return ret
	}
	return *o.Domain
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Error) GetDomainOk() (*string, bool) {
	if o == nil || o.Domain == nil {
		return nil, false
	}
	return o.Domain, true
}

// HasDomain returns a boolean if a field has been set.
func (o *Error) HasDomain() bool {
	if o != nil && o.Domain != nil {
		return true
	}

	return false
}

// SetDomain gets a reference to the given string and assigns it to the Domain field.
func (o *Error) SetDomain(v string) {
	o.Domain = &v
}

// GetErrorId returns the ErrorId field value if set, zero value otherwise.
func (o *Error) GetErrorId() int32 {
	if o == nil || o.ErrorId == nil {
		var ret int32
		return ret
	}
	return *o.ErrorId
}

// GetErrorIdOk returns a tuple with the ErrorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Error) GetErrorIdOk() (*int32, bool) {
	if o == nil || o.ErrorId == nil {
		return nil, false
	}
	return o.ErrorId, true
}

// HasErrorId returns a boolean if a field has been set.
func (o *Error) HasErrorId() bool {
	if o != nil && o.ErrorId != nil {
		return true
	}

	return false
}

// SetErrorId gets a reference to the given int32 and assigns it to the ErrorId field.
func (o *Error) SetErrorId(v int32) {
	o.ErrorId = &v
}

// GetInputRefIds returns the InputRefIds field value if set, zero value otherwise.
func (o *Error) GetInputRefIds() []string {
	if o == nil || o.InputRefIds == nil {
		var ret []string
		return ret
	}
	return *o.InputRefIds
}

// GetInputRefIdsOk returns a tuple with the InputRefIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Error) GetInputRefIdsOk() (*[]string, bool) {
	if o == nil || o.InputRefIds == nil {
		return nil, false
	}
	return o.InputRefIds, true
}

// HasInputRefIds returns a boolean if a field has been set.
func (o *Error) HasInputRefIds() bool {
	if o != nil && o.InputRefIds != nil {
		return true
	}

	return false
}

// SetInputRefIds gets a reference to the given []string and assigns it to the InputRefIds field.
func (o *Error) SetInputRefIds(v []string) {
	o.InputRefIds = &v
}

// GetLongMessage returns the LongMessage field value if set, zero value otherwise.
func (o *Error) GetLongMessage() string {
	if o == nil || o.LongMessage == nil {
		var ret string
		return ret
	}
	return *o.LongMessage
}

// GetLongMessageOk returns a tuple with the LongMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Error) GetLongMessageOk() (*string, bool) {
	if o == nil || o.LongMessage == nil {
		return nil, false
	}
	return o.LongMessage, true
}

// HasLongMessage returns a boolean if a field has been set.
func (o *Error) HasLongMessage() bool {
	if o != nil && o.LongMessage != nil {
		return true
	}

	return false
}

// SetLongMessage gets a reference to the given string and assigns it to the LongMessage field.
func (o *Error) SetLongMessage(v string) {
	o.LongMessage = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *Error) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Error) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *Error) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *Error) SetMessage(v string) {
	o.Message = &v
}

// GetOutputRefIds returns the OutputRefIds field value if set, zero value otherwise.
func (o *Error) GetOutputRefIds() []string {
	if o == nil || o.OutputRefIds == nil {
		var ret []string
		return ret
	}
	return *o.OutputRefIds
}

// GetOutputRefIdsOk returns a tuple with the OutputRefIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Error) GetOutputRefIdsOk() (*[]string, bool) {
	if o == nil || o.OutputRefIds == nil {
		return nil, false
	}
	return o.OutputRefIds, true
}

// HasOutputRefIds returns a boolean if a field has been set.
func (o *Error) HasOutputRefIds() bool {
	if o != nil && o.OutputRefIds != nil {
		return true
	}

	return false
}

// SetOutputRefIds gets a reference to the given []string and assigns it to the OutputRefIds field.
func (o *Error) SetOutputRefIds(v []string) {
	o.OutputRefIds = &v
}

// GetParameters returns the Parameters field value if set, zero value otherwise.
func (o *Error) GetParameters() []ErrorParameter {
	if o == nil || o.Parameters == nil {
		var ret []ErrorParameter
		return ret
	}
	return *o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Error) GetParametersOk() (*[]ErrorParameter, bool) {
	if o == nil || o.Parameters == nil {
		return nil, false
	}
	return o.Parameters, true
}

// HasParameters returns a boolean if a field has been set.
func (o *Error) HasParameters() bool {
	if o != nil && o.Parameters != nil {
		return true
	}

	return false
}

// SetParameters gets a reference to the given []ErrorParameter and assigns it to the Parameters field.
func (o *Error) SetParameters(v []ErrorParameter) {
	o.Parameters = &v
}

// GetSubdomain returns the Subdomain field value if set, zero value otherwise.
func (o *Error) GetSubdomain() string {
	if o == nil || o.Subdomain == nil {
		var ret string
		return ret
	}
	return *o.Subdomain
}

// GetSubdomainOk returns a tuple with the Subdomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Error) GetSubdomainOk() (*string, bool) {
	if o == nil || o.Subdomain == nil {
		return nil, false
	}
	return o.Subdomain, true
}

// HasSubdomain returns a boolean if a field has been set.
func (o *Error) HasSubdomain() bool {
	if o != nil && o.Subdomain != nil {
		return true
	}

	return false
}

// SetSubdomain gets a reference to the given string and assigns it to the Subdomain field.
func (o *Error) SetSubdomain(v string) {
	o.Subdomain = &v
}

func (o Error) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Category != nil {
		toSerialize["category"] = o.Category
	}
	if o.Domain != nil {
		toSerialize["domain"] = o.Domain
	}
	if o.ErrorId != nil {
		toSerialize["errorId"] = o.ErrorId
	}
	if o.InputRefIds != nil {
		toSerialize["inputRefIds"] = o.InputRefIds
	}
	if o.LongMessage != nil {
		toSerialize["longMessage"] = o.LongMessage
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if o.OutputRefIds != nil {
		toSerialize["outputRefIds"] = o.OutputRefIds
	}
	if o.Parameters != nil {
		toSerialize["parameters"] = o.Parameters
	}
	if o.Subdomain != nil {
		toSerialize["subdomain"] = o.Subdomain
	}
	return json.Marshal(toSerialize)
}

type NullableError struct {
	value *Error
	isSet bool
}

func (v NullableError) Get() *Error {
	return v.value
}

func (v *NullableError) Set(val *Error) {
	v.value = val
	v.isSet = true
}

func (v NullableError) IsSet() bool {
	return v.isSet
}

func (v *NullableError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableError(val *Error) *NullableError {
	return &NullableError{value: val, isSet: true}
}

func (v NullableError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


