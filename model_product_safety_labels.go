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

// checks if the ProductSafetyLabels type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductSafetyLabels{}

// ProductSafetyLabels This type contains seller provided product safety pictograms and statements for the listing.
type ProductSafetyLabels struct {
	// An array of seller provided comma-separated string values that provides identifier, URL, and description for one or more pictograms associated with the listing.
	Pictograms []ProductSafetyLabelPictogram `json:"pictograms,omitempty"`
	// An array of seller provided comma-separated string values that provide identifier and description for one or more product safety statements associated with the listing.
	Statements []ProductSafetyLabelStatement `json:"statements,omitempty"`
}

// NewProductSafetyLabels instantiates a new ProductSafetyLabels object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductSafetyLabels() *ProductSafetyLabels {
	this := ProductSafetyLabels{}
	return &this
}

// NewProductSafetyLabelsWithDefaults instantiates a new ProductSafetyLabels object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductSafetyLabelsWithDefaults() *ProductSafetyLabels {
	this := ProductSafetyLabels{}
	return &this
}

// GetPictograms returns the Pictograms field value if set, zero value otherwise.
func (o *ProductSafetyLabels) GetPictograms() []ProductSafetyLabelPictogram {
	if o == nil || IsNil(o.Pictograms) {
		var ret []ProductSafetyLabelPictogram
		return ret
	}
	return o.Pictograms
}

// GetPictogramsOk returns a tuple with the Pictograms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductSafetyLabels) GetPictogramsOk() ([]ProductSafetyLabelPictogram, bool) {
	if o == nil || IsNil(o.Pictograms) {
		return nil, false
	}
	return o.Pictograms, true
}

// HasPictograms returns a boolean if a field has been set.
func (o *ProductSafetyLabels) HasPictograms() bool {
	if o != nil && !IsNil(o.Pictograms) {
		return true
	}

	return false
}

// SetPictograms gets a reference to the given []ProductSafetyLabelPictogram and assigns it to the Pictograms field.
func (o *ProductSafetyLabels) SetPictograms(v []ProductSafetyLabelPictogram) {
	o.Pictograms = v
}

// GetStatements returns the Statements field value if set, zero value otherwise.
func (o *ProductSafetyLabels) GetStatements() []ProductSafetyLabelStatement {
	if o == nil || IsNil(o.Statements) {
		var ret []ProductSafetyLabelStatement
		return ret
	}
	return o.Statements
}

// GetStatementsOk returns a tuple with the Statements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductSafetyLabels) GetStatementsOk() ([]ProductSafetyLabelStatement, bool) {
	if o == nil || IsNil(o.Statements) {
		return nil, false
	}
	return o.Statements, true
}

// HasStatements returns a boolean if a field has been set.
func (o *ProductSafetyLabels) HasStatements() bool {
	if o != nil && !IsNil(o.Statements) {
		return true
	}

	return false
}

// SetStatements gets a reference to the given []ProductSafetyLabelStatement and assigns it to the Statements field.
func (o *ProductSafetyLabels) SetStatements(v []ProductSafetyLabelStatement) {
	o.Statements = v
}

func (o ProductSafetyLabels) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProductSafetyLabels) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Pictograms) {
		toSerialize["pictograms"] = o.Pictograms
	}
	if !IsNil(o.Statements) {
		toSerialize["statements"] = o.Statements
	}
	return toSerialize, nil
}

type NullableProductSafetyLabels struct {
	value *ProductSafetyLabels
	isSet bool
}

func (v NullableProductSafetyLabels) Get() *ProductSafetyLabels {
	return v.value
}

func (v *NullableProductSafetyLabels) Set(val *ProductSafetyLabels) {
	v.value = val
	v.isSet = true
}

func (v NullableProductSafetyLabels) IsSet() bool {
	return v.isSet
}

func (v *NullableProductSafetyLabels) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductSafetyLabels(val *ProductSafetyLabels) *NullableProductSafetyLabels {
	return &NullableProductSafetyLabels{value: val, isSet: true}
}

func (v NullableProductSafetyLabels) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductSafetyLabels) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


