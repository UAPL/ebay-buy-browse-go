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

// checks if the Product type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Product{}

// Product The type that defines the fields for the product information of the item.
type Product struct {
	// An array of containers with the URLs for the product images that are in addition to the primary image. 
	AdditionalImages []Image `json:"additionalImages,omitempty"`
	// An array of product identifiers associated with the item. This container is returned if the seller has associated the eBay Product Identifier (ePID) with the item and in the request <b> fieldgroups</b> is set to <code>PRODUCT</code>.
	AdditionalProductIdentities []AdditionalProductIdentity `json:"additionalProductIdentities,omitempty"`
	// An array of containers for the product aspects. Each group contains the aspect group name and the aspect name/value pairs.
	AspectGroups []AspectGroup `json:"aspectGroups,omitempty"`
	// The brand associated with product. To identify the product, this is always used along with MPN (manufacturer part number).
	Brand *string `json:"brand,omitempty"`
	// The rich description of an eBay product, which might contain HTML.
	Description *string `json:"description,omitempty"`
	// An array of all the possible GTINs values associated with the product. A GTIN is a unique Global Trade Item number of the item as defined by <a href=\"https://www.gtin.info \" target=\"_blank\">https://www.gtin.info</a>. This can be a UPC (Universal Product Code), EAN (European Article Number), or an ISBN (International Standard Book Number) value.
	Gtins []string `json:"gtins,omitempty"`
	Image *Image `json:"image,omitempty"`
	// An array of all possible MPN values associated with the product. A MPNs is manufacturer part number of the product. To identify the product, this is always used along with brand.
	Mpns []string `json:"mpns,omitempty"`
	// The title of the product.
	Title *string `json:"title,omitempty"`
}

// NewProduct instantiates a new Product object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProduct() *Product {
	this := Product{}
	return &this
}

// NewProductWithDefaults instantiates a new Product object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductWithDefaults() *Product {
	this := Product{}
	return &this
}

// GetAdditionalImages returns the AdditionalImages field value if set, zero value otherwise.
func (o *Product) GetAdditionalImages() []Image {
	if o == nil || IsNil(o.AdditionalImages) {
		var ret []Image
		return ret
	}
	return o.AdditionalImages
}

// GetAdditionalImagesOk returns a tuple with the AdditionalImages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product) GetAdditionalImagesOk() ([]Image, bool) {
	if o == nil || IsNil(o.AdditionalImages) {
		return nil, false
	}
	return o.AdditionalImages, true
}

// HasAdditionalImages returns a boolean if a field has been set.
func (o *Product) HasAdditionalImages() bool {
	if o != nil && !IsNil(o.AdditionalImages) {
		return true
	}

	return false
}

// SetAdditionalImages gets a reference to the given []Image and assigns it to the AdditionalImages field.
func (o *Product) SetAdditionalImages(v []Image) {
	o.AdditionalImages = v
}

// GetAdditionalProductIdentities returns the AdditionalProductIdentities field value if set, zero value otherwise.
func (o *Product) GetAdditionalProductIdentities() []AdditionalProductIdentity {
	if o == nil || IsNil(o.AdditionalProductIdentities) {
		var ret []AdditionalProductIdentity
		return ret
	}
	return o.AdditionalProductIdentities
}

// GetAdditionalProductIdentitiesOk returns a tuple with the AdditionalProductIdentities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product) GetAdditionalProductIdentitiesOk() ([]AdditionalProductIdentity, bool) {
	if o == nil || IsNil(o.AdditionalProductIdentities) {
		return nil, false
	}
	return o.AdditionalProductIdentities, true
}

// HasAdditionalProductIdentities returns a boolean if a field has been set.
func (o *Product) HasAdditionalProductIdentities() bool {
	if o != nil && !IsNil(o.AdditionalProductIdentities) {
		return true
	}

	return false
}

// SetAdditionalProductIdentities gets a reference to the given []AdditionalProductIdentity and assigns it to the AdditionalProductIdentities field.
func (o *Product) SetAdditionalProductIdentities(v []AdditionalProductIdentity) {
	o.AdditionalProductIdentities = v
}

// GetAspectGroups returns the AspectGroups field value if set, zero value otherwise.
func (o *Product) GetAspectGroups() []AspectGroup {
	if o == nil || IsNil(o.AspectGroups) {
		var ret []AspectGroup
		return ret
	}
	return o.AspectGroups
}

// GetAspectGroupsOk returns a tuple with the AspectGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product) GetAspectGroupsOk() ([]AspectGroup, bool) {
	if o == nil || IsNil(o.AspectGroups) {
		return nil, false
	}
	return o.AspectGroups, true
}

// HasAspectGroups returns a boolean if a field has been set.
func (o *Product) HasAspectGroups() bool {
	if o != nil && !IsNil(o.AspectGroups) {
		return true
	}

	return false
}

// SetAspectGroups gets a reference to the given []AspectGroup and assigns it to the AspectGroups field.
func (o *Product) SetAspectGroups(v []AspectGroup) {
	o.AspectGroups = v
}

// GetBrand returns the Brand field value if set, zero value otherwise.
func (o *Product) GetBrand() string {
	if o == nil || IsNil(o.Brand) {
		var ret string
		return ret
	}
	return *o.Brand
}

// GetBrandOk returns a tuple with the Brand field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product) GetBrandOk() (*string, bool) {
	if o == nil || IsNil(o.Brand) {
		return nil, false
	}
	return o.Brand, true
}

// HasBrand returns a boolean if a field has been set.
func (o *Product) HasBrand() bool {
	if o != nil && !IsNil(o.Brand) {
		return true
	}

	return false
}

// SetBrand gets a reference to the given string and assigns it to the Brand field.
func (o *Product) SetBrand(v string) {
	o.Brand = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Product) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Product) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Product) SetDescription(v string) {
	o.Description = &v
}

// GetGtins returns the Gtins field value if set, zero value otherwise.
func (o *Product) GetGtins() []string {
	if o == nil || IsNil(o.Gtins) {
		var ret []string
		return ret
	}
	return o.Gtins
}

// GetGtinsOk returns a tuple with the Gtins field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product) GetGtinsOk() ([]string, bool) {
	if o == nil || IsNil(o.Gtins) {
		return nil, false
	}
	return o.Gtins, true
}

// HasGtins returns a boolean if a field has been set.
func (o *Product) HasGtins() bool {
	if o != nil && !IsNil(o.Gtins) {
		return true
	}

	return false
}

// SetGtins gets a reference to the given []string and assigns it to the Gtins field.
func (o *Product) SetGtins(v []string) {
	o.Gtins = v
}

// GetImage returns the Image field value if set, zero value otherwise.
func (o *Product) GetImage() Image {
	if o == nil || IsNil(o.Image) {
		var ret Image
		return ret
	}
	return *o.Image
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product) GetImageOk() (*Image, bool) {
	if o == nil || IsNil(o.Image) {
		return nil, false
	}
	return o.Image, true
}

// HasImage returns a boolean if a field has been set.
func (o *Product) HasImage() bool {
	if o != nil && !IsNil(o.Image) {
		return true
	}

	return false
}

// SetImage gets a reference to the given Image and assigns it to the Image field.
func (o *Product) SetImage(v Image) {
	o.Image = &v
}

// GetMpns returns the Mpns field value if set, zero value otherwise.
func (o *Product) GetMpns() []string {
	if o == nil || IsNil(o.Mpns) {
		var ret []string
		return ret
	}
	return o.Mpns
}

// GetMpnsOk returns a tuple with the Mpns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product) GetMpnsOk() ([]string, bool) {
	if o == nil || IsNil(o.Mpns) {
		return nil, false
	}
	return o.Mpns, true
}

// HasMpns returns a boolean if a field has been set.
func (o *Product) HasMpns() bool {
	if o != nil && !IsNil(o.Mpns) {
		return true
	}

	return false
}

// SetMpns gets a reference to the given []string and assigns it to the Mpns field.
func (o *Product) SetMpns(v []string) {
	o.Mpns = v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *Product) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *Product) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *Product) SetTitle(v string) {
	o.Title = &v
}

func (o Product) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Product) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AdditionalImages) {
		toSerialize["additionalImages"] = o.AdditionalImages
	}
	if !IsNil(o.AdditionalProductIdentities) {
		toSerialize["additionalProductIdentities"] = o.AdditionalProductIdentities
	}
	if !IsNil(o.AspectGroups) {
		toSerialize["aspectGroups"] = o.AspectGroups
	}
	if !IsNil(o.Brand) {
		toSerialize["brand"] = o.Brand
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Gtins) {
		toSerialize["gtins"] = o.Gtins
	}
	if !IsNil(o.Image) {
		toSerialize["image"] = o.Image
	}
	if !IsNil(o.Mpns) {
		toSerialize["mpns"] = o.Mpns
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	return toSerialize, nil
}

type NullableProduct struct {
	value *Product
	isSet bool
}

func (v NullableProduct) Get() *Product {
	return v.value
}

func (v *NullableProduct) Set(val *Product) {
	v.value = val
	v.isSet = true
}

func (v NullableProduct) IsSet() bool {
	return v.isSet
}

func (v *NullableProduct) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProduct(val *Product) *NullableProduct {
	return &NullableProduct{value: val, isSet: true}
}

func (v NullableProduct) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProduct) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


