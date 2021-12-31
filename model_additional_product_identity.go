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

// AdditionalProductIdentity The type that defines the array of product identifiers associated with the item. This container is returned if the seller has associated the eBay Product Identifier (ePID) with the item and in the request <b> fieldgroups</b> is set to <code>PRODUCT</code>.
type AdditionalProductIdentity struct {
	// An array of the product identifier/value pairs for the product associated with the item. This is returned if the seller has associated the eBay Product Identifier (ePID) with the item and the request has <b> fieldgroups</b> set to <code>PRODUCT</code>. <br /><br />The following table shows what is returned, based on the item information provided by the seller, when the <b> fieldgroups</b> set to <code>PRODUCT</code>.        <br /><br /><div style=\"overflow-x:auto;\"> <table border=1> <tr> <th> ePID Provided </th>  <th> Product&nbsp;ID(s) Provided</th> <th> Response</th> </tr> <tr> <td> No </td>  <td> No </td> <td> The <b> AdditionalProductIdentity</b> container is <i> not</i> returned.</td></tr>   <tr> <td> No </td>  <td> Yes </td>  <td> The <b> AdditionalProductIdentity</b> container is <i> not</i> returned but the product identifiers specified by the seller are returned in the <b> localizedAspects</b> container. </td>  </tr>   <tr> <td> Yes </td>  <td> No </td> <td>  The <b> AdditionalProductIdentity</b> container is returned listing the product identifiers of the product.</td></tr>   <tr> <td> Yes </td>  <td> Yes </td> <td> The <b> AdditionalProductIdentity</b> container is returned listing all the product identifiers of the product and the product identifiers specified by the seller are returned in the <b> localizedAspects</b> container.</td> </tr>   </table> </div>
	ProductIdentity *[]ProductIdentity `json:"productIdentity,omitempty"`
}

// NewAdditionalProductIdentity instantiates a new AdditionalProductIdentity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdditionalProductIdentity() *AdditionalProductIdentity {
	this := AdditionalProductIdentity{}
	return &this
}

// NewAdditionalProductIdentityWithDefaults instantiates a new AdditionalProductIdentity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdditionalProductIdentityWithDefaults() *AdditionalProductIdentity {
	this := AdditionalProductIdentity{}
	return &this
}

// GetProductIdentity returns the ProductIdentity field value if set, zero value otherwise.
func (o *AdditionalProductIdentity) GetProductIdentity() []ProductIdentity {
	if o == nil || o.ProductIdentity == nil {
		var ret []ProductIdentity
		return ret
	}
	return *o.ProductIdentity
}

// GetProductIdentityOk returns a tuple with the ProductIdentity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalProductIdentity) GetProductIdentityOk() (*[]ProductIdentity, bool) {
	if o == nil || o.ProductIdentity == nil {
		return nil, false
	}
	return o.ProductIdentity, true
}

// HasProductIdentity returns a boolean if a field has been set.
func (o *AdditionalProductIdentity) HasProductIdentity() bool {
	if o != nil && o.ProductIdentity != nil {
		return true
	}

	return false
}

// SetProductIdentity gets a reference to the given []ProductIdentity and assigns it to the ProductIdentity field.
func (o *AdditionalProductIdentity) SetProductIdentity(v []ProductIdentity) {
	o.ProductIdentity = &v
}

func (o AdditionalProductIdentity) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ProductIdentity != nil {
		toSerialize["productIdentity"] = o.ProductIdentity
	}
	return json.Marshal(toSerialize)
}

type NullableAdditionalProductIdentity struct {
	value *AdditionalProductIdentity
	isSet bool
}

func (v NullableAdditionalProductIdentity) Get() *AdditionalProductIdentity {
	return v.value
}

func (v *NullableAdditionalProductIdentity) Set(val *AdditionalProductIdentity) {
	v.value = val
	v.isSet = true
}

func (v NullableAdditionalProductIdentity) IsSet() bool {
	return v.isSet
}

func (v *NullableAdditionalProductIdentity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdditionalProductIdentity(val *AdditionalProductIdentity) *NullableAdditionalProductIdentity {
	return &NullableAdditionalProductIdentity{value: val, isSet: true}
}

func (v NullableAdditionalProductIdentity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdditionalProductIdentity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

