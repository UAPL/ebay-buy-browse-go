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

// UpdateCartItemInput The type that defines the fields for the <b>updateQuantity</b> request.
type UpdateCartItemInput struct {
	// The identifier of the item in the cart to be updated. This ID is generated when the item was added to the cart.
	CartItemId *string `json:"cartItemId,omitempty"`
	// The new quantity for the item that is being updated.
	Quantity *int32 `json:"quantity,omitempty"`
}

// NewUpdateCartItemInput instantiates a new UpdateCartItemInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateCartItemInput() *UpdateCartItemInput {
	this := UpdateCartItemInput{}
	return &this
}

// NewUpdateCartItemInputWithDefaults instantiates a new UpdateCartItemInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateCartItemInputWithDefaults() *UpdateCartItemInput {
	this := UpdateCartItemInput{}
	return &this
}

// GetCartItemId returns the CartItemId field value if set, zero value otherwise.
func (o *UpdateCartItemInput) GetCartItemId() string {
	if o == nil || o.CartItemId == nil {
		var ret string
		return ret
	}
	return *o.CartItemId
}

// GetCartItemIdOk returns a tuple with the CartItemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCartItemInput) GetCartItemIdOk() (*string, bool) {
	if o == nil || o.CartItemId == nil {
		return nil, false
	}
	return o.CartItemId, true
}

// HasCartItemId returns a boolean if a field has been set.
func (o *UpdateCartItemInput) HasCartItemId() bool {
	if o != nil && o.CartItemId != nil {
		return true
	}

	return false
}

// SetCartItemId gets a reference to the given string and assigns it to the CartItemId field.
func (o *UpdateCartItemInput) SetCartItemId(v string) {
	o.CartItemId = &v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *UpdateCartItemInput) GetQuantity() int32 {
	if o == nil || o.Quantity == nil {
		var ret int32
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCartItemInput) GetQuantityOk() (*int32, bool) {
	if o == nil || o.Quantity == nil {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *UpdateCartItemInput) HasQuantity() bool {
	if o != nil && o.Quantity != nil {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given int32 and assigns it to the Quantity field.
func (o *UpdateCartItemInput) SetQuantity(v int32) {
	o.Quantity = &v
}

func (o UpdateCartItemInput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CartItemId != nil {
		toSerialize["cartItemId"] = o.CartItemId
	}
	if o.Quantity != nil {
		toSerialize["quantity"] = o.Quantity
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateCartItemInput struct {
	value *UpdateCartItemInput
	isSet bool
}

func (v NullableUpdateCartItemInput) Get() *UpdateCartItemInput {
	return v.value
}

func (v *NullableUpdateCartItemInput) Set(val *UpdateCartItemInput) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateCartItemInput) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateCartItemInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateCartItemInput(val *UpdateCartItemInput) *NullableUpdateCartItemInput {
	return &NullableUpdateCartItemInput{value: val, isSet: true}
}

func (v NullableUpdateCartItemInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateCartItemInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


