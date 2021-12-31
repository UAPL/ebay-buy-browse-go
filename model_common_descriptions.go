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

// CommonDescriptions The type that defines the fields for the item ids that all use a common description.  Often the item variations within an item group all have the same description. Instead of repeating this description in the item details of each item, a description that is shared by at least one other item is returned in this container. If the description is unique, it is returned in the <b> items.description</b> field.
type CommonDescriptions struct {
	// The item description that is used by more than one of the item variations.
	Description *string `json:"description,omitempty"`
	// A list of item ids that have this description.
	ItemIds *[]string `json:"itemIds,omitempty"`
}

// NewCommonDescriptions instantiates a new CommonDescriptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommonDescriptions() *CommonDescriptions {
	this := CommonDescriptions{}
	return &this
}

// NewCommonDescriptionsWithDefaults instantiates a new CommonDescriptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommonDescriptionsWithDefaults() *CommonDescriptions {
	this := CommonDescriptions{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CommonDescriptions) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonDescriptions) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CommonDescriptions) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CommonDescriptions) SetDescription(v string) {
	o.Description = &v
}

// GetItemIds returns the ItemIds field value if set, zero value otherwise.
func (o *CommonDescriptions) GetItemIds() []string {
	if o == nil || o.ItemIds == nil {
		var ret []string
		return ret
	}
	return *o.ItemIds
}

// GetItemIdsOk returns a tuple with the ItemIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonDescriptions) GetItemIdsOk() (*[]string, bool) {
	if o == nil || o.ItemIds == nil {
		return nil, false
	}
	return o.ItemIds, true
}

// HasItemIds returns a boolean if a field has been set.
func (o *CommonDescriptions) HasItemIds() bool {
	if o != nil && o.ItemIds != nil {
		return true
	}

	return false
}

// SetItemIds gets a reference to the given []string and assigns it to the ItemIds field.
func (o *CommonDescriptions) SetItemIds(v []string) {
	o.ItemIds = &v
}

func (o CommonDescriptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.ItemIds != nil {
		toSerialize["itemIds"] = o.ItemIds
	}
	return json.Marshal(toSerialize)
}

type NullableCommonDescriptions struct {
	value *CommonDescriptions
	isSet bool
}

func (v NullableCommonDescriptions) Get() *CommonDescriptions {
	return v.value
}

func (v *NullableCommonDescriptions) Set(val *CommonDescriptions) {
	v.value = val
	v.isSet = true
}

func (v NullableCommonDescriptions) IsSet() bool {
	return v.isSet
}

func (v *NullableCommonDescriptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommonDescriptions(val *CommonDescriptions) *NullableCommonDescriptions {
	return &NullableCommonDescriptions{value: val, isSet: true}
}

func (v NullableCommonDescriptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommonDescriptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


