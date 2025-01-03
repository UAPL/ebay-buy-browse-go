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

// checks if the ShipToLocations type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ShipToLocations{}

// ShipToLocations The type that defines the fields that include and exclude geographic regions affecting where the item can be shipped. The seller defines these regions when listing the item.
type ShipToLocations struct {
	// An array of containers that express the large geographical regions, countries, state/provinces, or special locations within a country where the seller is not willing to ship to.
	RegionExcluded []ShipToRegion `json:"regionExcluded,omitempty"`
	// An array of containers that express the large geographical regions, countries, or state/provinces within a country where the seller is willing to ship to. Prospective buyers must look at the shipping regions under this container, as well as the shipping regions that are under the <b>regionExcluded</b> to see where the seller is willing to ship items. Sellers can specify that they ship 'Worldwide', but then add several large geographical regions (e.g. Asia, Oceania, Middle East) to the exclusion list, or sellers can specify that they ship to Europe and Africa, but then add several individual countries to the exclusion list.
	RegionIncluded []ShipToRegion `json:"regionIncluded,omitempty"`
}

// NewShipToLocations instantiates a new ShipToLocations object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShipToLocations() *ShipToLocations {
	this := ShipToLocations{}
	return &this
}

// NewShipToLocationsWithDefaults instantiates a new ShipToLocations object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShipToLocationsWithDefaults() *ShipToLocations {
	this := ShipToLocations{}
	return &this
}

// GetRegionExcluded returns the RegionExcluded field value if set, zero value otherwise.
func (o *ShipToLocations) GetRegionExcluded() []ShipToRegion {
	if o == nil || IsNil(o.RegionExcluded) {
		var ret []ShipToRegion
		return ret
	}
	return o.RegionExcluded
}

// GetRegionExcludedOk returns a tuple with the RegionExcluded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipToLocations) GetRegionExcludedOk() ([]ShipToRegion, bool) {
	if o == nil || IsNil(o.RegionExcluded) {
		return nil, false
	}
	return o.RegionExcluded, true
}

// HasRegionExcluded returns a boolean if a field has been set.
func (o *ShipToLocations) HasRegionExcluded() bool {
	if o != nil && !IsNil(o.RegionExcluded) {
		return true
	}

	return false
}

// SetRegionExcluded gets a reference to the given []ShipToRegion and assigns it to the RegionExcluded field.
func (o *ShipToLocations) SetRegionExcluded(v []ShipToRegion) {
	o.RegionExcluded = v
}

// GetRegionIncluded returns the RegionIncluded field value if set, zero value otherwise.
func (o *ShipToLocations) GetRegionIncluded() []ShipToRegion {
	if o == nil || IsNil(o.RegionIncluded) {
		var ret []ShipToRegion
		return ret
	}
	return o.RegionIncluded
}

// GetRegionIncludedOk returns a tuple with the RegionIncluded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipToLocations) GetRegionIncludedOk() ([]ShipToRegion, bool) {
	if o == nil || IsNil(o.RegionIncluded) {
		return nil, false
	}
	return o.RegionIncluded, true
}

// HasRegionIncluded returns a boolean if a field has been set.
func (o *ShipToLocations) HasRegionIncluded() bool {
	if o != nil && !IsNil(o.RegionIncluded) {
		return true
	}

	return false
}

// SetRegionIncluded gets a reference to the given []ShipToRegion and assigns it to the RegionIncluded field.
func (o *ShipToLocations) SetRegionIncluded(v []ShipToRegion) {
	o.RegionIncluded = v
}

func (o ShipToLocations) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ShipToLocations) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RegionExcluded) {
		toSerialize["regionExcluded"] = o.RegionExcluded
	}
	if !IsNil(o.RegionIncluded) {
		toSerialize["regionIncluded"] = o.RegionIncluded
	}
	return toSerialize, nil
}

type NullableShipToLocations struct {
	value *ShipToLocations
	isSet bool
}

func (v NullableShipToLocations) Get() *ShipToLocations {
	return v.value
}

func (v *NullableShipToLocations) Set(val *ShipToLocations) {
	v.value = val
	v.isSet = true
}

func (v NullableShipToLocations) IsSet() bool {
	return v.isSet
}

func (v *NullableShipToLocations) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShipToLocations(val *ShipToLocations) *NullableShipToLocations {
	return &NullableShipToLocations{value: val, isSet: true}
}

func (v NullableShipToLocations) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShipToLocations) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


