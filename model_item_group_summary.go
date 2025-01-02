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

// checks if the ItemGroupSummary type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ItemGroupSummary{}

// ItemGroupSummary The type that defines the fields for the details of each item in an item group. An item group is  an item that has various aspect differences, such as color, size, storage capacity, etc. When an item group is created, one of the item variations, such as the red shirt size L, is chosen as the \"parent\". All the other items in the group are the children, such as the blue shirt size L, red shirt size M, etc. <br><br><span class=\"tablenote\"><b> Note: </b> This container is returned only if the <b> item_id</b> in the request is an item group (parent ID of an item with variations).</span>
type ItemGroupSummary struct {
	// An array of containers with the URLs for images that are in addition to the primary image of the item group.  The primary image is returned in the <b> itemGroupImage</b> field.
	ItemGroupAdditionalImages []Image `json:"itemGroupAdditionalImages,omitempty"`
	// The HATEOAS reference of the parent page of the item group. An item group is an item that has various aspect differences, such as color, size, storage capacity, etc. 
	ItemGroupHref *string `json:"itemGroupHref,omitempty"`
	// The unique identifier for the item group. An item group is an item that has various aspect differences, such as color, size, storage capacity, etc. 
	ItemGroupId *string `json:"itemGroupId,omitempty"`
	ItemGroupImage *Image `json:"itemGroupImage,omitempty"`
	// The title of the item that appears on the item group page. An item group is an item that has various aspect differences, such as color, size, storage capacity, etc. 
	ItemGroupTitle *string `json:"itemGroupTitle,omitempty"`
	// An enumeration value that indicates the type of the item group. An item group is an item that has various aspect differences, such as color, size, storage capacity, etc. For implementation help, refer to <a href='https://developer.ebay.com/api-docs/buy/browse/types/gct:ItemGroupTypeEnum'>eBay API documentation</a>
	ItemGroupType *string `json:"itemGroupType,omitempty"`
}

// NewItemGroupSummary instantiates a new ItemGroupSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewItemGroupSummary() *ItemGroupSummary {
	this := ItemGroupSummary{}
	return &this
}

// NewItemGroupSummaryWithDefaults instantiates a new ItemGroupSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewItemGroupSummaryWithDefaults() *ItemGroupSummary {
	this := ItemGroupSummary{}
	return &this
}

// GetItemGroupAdditionalImages returns the ItemGroupAdditionalImages field value if set, zero value otherwise.
func (o *ItemGroupSummary) GetItemGroupAdditionalImages() []Image {
	if o == nil || IsNil(o.ItemGroupAdditionalImages) {
		var ret []Image
		return ret
	}
	return o.ItemGroupAdditionalImages
}

// GetItemGroupAdditionalImagesOk returns a tuple with the ItemGroupAdditionalImages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemGroupSummary) GetItemGroupAdditionalImagesOk() ([]Image, bool) {
	if o == nil || IsNil(o.ItemGroupAdditionalImages) {
		return nil, false
	}
	return o.ItemGroupAdditionalImages, true
}

// HasItemGroupAdditionalImages returns a boolean if a field has been set.
func (o *ItemGroupSummary) HasItemGroupAdditionalImages() bool {
	if o != nil && !IsNil(o.ItemGroupAdditionalImages) {
		return true
	}

	return false
}

// SetItemGroupAdditionalImages gets a reference to the given []Image and assigns it to the ItemGroupAdditionalImages field.
func (o *ItemGroupSummary) SetItemGroupAdditionalImages(v []Image) {
	o.ItemGroupAdditionalImages = v
}

// GetItemGroupHref returns the ItemGroupHref field value if set, zero value otherwise.
func (o *ItemGroupSummary) GetItemGroupHref() string {
	if o == nil || IsNil(o.ItemGroupHref) {
		var ret string
		return ret
	}
	return *o.ItemGroupHref
}

// GetItemGroupHrefOk returns a tuple with the ItemGroupHref field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemGroupSummary) GetItemGroupHrefOk() (*string, bool) {
	if o == nil || IsNil(o.ItemGroupHref) {
		return nil, false
	}
	return o.ItemGroupHref, true
}

// HasItemGroupHref returns a boolean if a field has been set.
func (o *ItemGroupSummary) HasItemGroupHref() bool {
	if o != nil && !IsNil(o.ItemGroupHref) {
		return true
	}

	return false
}

// SetItemGroupHref gets a reference to the given string and assigns it to the ItemGroupHref field.
func (o *ItemGroupSummary) SetItemGroupHref(v string) {
	o.ItemGroupHref = &v
}

// GetItemGroupId returns the ItemGroupId field value if set, zero value otherwise.
func (o *ItemGroupSummary) GetItemGroupId() string {
	if o == nil || IsNil(o.ItemGroupId) {
		var ret string
		return ret
	}
	return *o.ItemGroupId
}

// GetItemGroupIdOk returns a tuple with the ItemGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemGroupSummary) GetItemGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.ItemGroupId) {
		return nil, false
	}
	return o.ItemGroupId, true
}

// HasItemGroupId returns a boolean if a field has been set.
func (o *ItemGroupSummary) HasItemGroupId() bool {
	if o != nil && !IsNil(o.ItemGroupId) {
		return true
	}

	return false
}

// SetItemGroupId gets a reference to the given string and assigns it to the ItemGroupId field.
func (o *ItemGroupSummary) SetItemGroupId(v string) {
	o.ItemGroupId = &v
}

// GetItemGroupImage returns the ItemGroupImage field value if set, zero value otherwise.
func (o *ItemGroupSummary) GetItemGroupImage() Image {
	if o == nil || IsNil(o.ItemGroupImage) {
		var ret Image
		return ret
	}
	return *o.ItemGroupImage
}

// GetItemGroupImageOk returns a tuple with the ItemGroupImage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemGroupSummary) GetItemGroupImageOk() (*Image, bool) {
	if o == nil || IsNil(o.ItemGroupImage) {
		return nil, false
	}
	return o.ItemGroupImage, true
}

// HasItemGroupImage returns a boolean if a field has been set.
func (o *ItemGroupSummary) HasItemGroupImage() bool {
	if o != nil && !IsNil(o.ItemGroupImage) {
		return true
	}

	return false
}

// SetItemGroupImage gets a reference to the given Image and assigns it to the ItemGroupImage field.
func (o *ItemGroupSummary) SetItemGroupImage(v Image) {
	o.ItemGroupImage = &v
}

// GetItemGroupTitle returns the ItemGroupTitle field value if set, zero value otherwise.
func (o *ItemGroupSummary) GetItemGroupTitle() string {
	if o == nil || IsNil(o.ItemGroupTitle) {
		var ret string
		return ret
	}
	return *o.ItemGroupTitle
}

// GetItemGroupTitleOk returns a tuple with the ItemGroupTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemGroupSummary) GetItemGroupTitleOk() (*string, bool) {
	if o == nil || IsNil(o.ItemGroupTitle) {
		return nil, false
	}
	return o.ItemGroupTitle, true
}

// HasItemGroupTitle returns a boolean if a field has been set.
func (o *ItemGroupSummary) HasItemGroupTitle() bool {
	if o != nil && !IsNil(o.ItemGroupTitle) {
		return true
	}

	return false
}

// SetItemGroupTitle gets a reference to the given string and assigns it to the ItemGroupTitle field.
func (o *ItemGroupSummary) SetItemGroupTitle(v string) {
	o.ItemGroupTitle = &v
}

// GetItemGroupType returns the ItemGroupType field value if set, zero value otherwise.
func (o *ItemGroupSummary) GetItemGroupType() string {
	if o == nil || IsNil(o.ItemGroupType) {
		var ret string
		return ret
	}
	return *o.ItemGroupType
}

// GetItemGroupTypeOk returns a tuple with the ItemGroupType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemGroupSummary) GetItemGroupTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ItemGroupType) {
		return nil, false
	}
	return o.ItemGroupType, true
}

// HasItemGroupType returns a boolean if a field has been set.
func (o *ItemGroupSummary) HasItemGroupType() bool {
	if o != nil && !IsNil(o.ItemGroupType) {
		return true
	}

	return false
}

// SetItemGroupType gets a reference to the given string and assigns it to the ItemGroupType field.
func (o *ItemGroupSummary) SetItemGroupType(v string) {
	o.ItemGroupType = &v
}

func (o ItemGroupSummary) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ItemGroupSummary) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ItemGroupAdditionalImages) {
		toSerialize["itemGroupAdditionalImages"] = o.ItemGroupAdditionalImages
	}
	if !IsNil(o.ItemGroupHref) {
		toSerialize["itemGroupHref"] = o.ItemGroupHref
	}
	if !IsNil(o.ItemGroupId) {
		toSerialize["itemGroupId"] = o.ItemGroupId
	}
	if !IsNil(o.ItemGroupImage) {
		toSerialize["itemGroupImage"] = o.ItemGroupImage
	}
	if !IsNil(o.ItemGroupTitle) {
		toSerialize["itemGroupTitle"] = o.ItemGroupTitle
	}
	if !IsNil(o.ItemGroupType) {
		toSerialize["itemGroupType"] = o.ItemGroupType
	}
	return toSerialize, nil
}

type NullableItemGroupSummary struct {
	value *ItemGroupSummary
	isSet bool
}

func (v NullableItemGroupSummary) Get() *ItemGroupSummary {
	return v.value
}

func (v *NullableItemGroupSummary) Set(val *ItemGroupSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableItemGroupSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableItemGroupSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableItemGroupSummary(val *ItemGroupSummary) *NullableItemGroupSummary {
	return &NullableItemGroupSummary{value: val, isSet: true}
}

func (v NullableItemGroupSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableItemGroupSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


