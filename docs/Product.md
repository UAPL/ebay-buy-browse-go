# Product

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalImages** | Pointer to [**[]Image**](Image.md) | An array of containers with the URLs for the product images that are in addition to the primary image.  | [optional] 
**AdditionalProductIdentities** | Pointer to [**[]AdditionalProductIdentity**](AdditionalProductIdentity.md) | An array of product identifiers associated with the item. This container is returned if the seller has associated the eBay Product Identifier (ePID) with the item and in the request &lt;b&gt; fieldgroups&lt;/b&gt; is set to &lt;code&gt;PRODUCT&lt;/code&gt;. | [optional] 
**AspectGroups** | Pointer to [**[]AspectGroup**](AspectGroup.md) | An array of containers for the product aspects. Each group contains the aspect group name and the aspect name/value pairs. | [optional] 
**Brand** | Pointer to **string** | The brand associated with product. To identify the product, this is always used along with MPN (manufacturer part number). | [optional] 
**Description** | Pointer to **string** | The rich description of an eBay product, which might contain HTML. | [optional] 
**Gtins** | Pointer to **[]string** | An array of all the possible GTINs values associated with the product. A GTIN is a unique Global Trade Item number of the item as defined by &lt;a href&#x3D;\&quot;https://www.gtin.info\&quot; target&#x3D;\&quot;_blank\&quot;&gt;https://www.gtin.info&lt;/a&gt;. This can be a UPC (Universal Product Code), EAN (European Article Number), or an ISBN (International Standard Book Number) value. | [optional] 
**Image** | Pointer to [**Image**](Image.md) |  | [optional] 
**Mpns** | Pointer to **[]string** | An array of all possible MPN values associated with the product. A MPNs is manufacturer part number of the product. To identify the product, this is always used along with brand. | [optional] 
**Title** | Pointer to **string** | The title of the product. | [optional] 

## Methods

### NewProduct

`func NewProduct() *Product`

NewProduct instantiates a new Product object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductWithDefaults

`func NewProductWithDefaults() *Product`

NewProductWithDefaults instantiates a new Product object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdditionalImages

`func (o *Product) GetAdditionalImages() []Image`

GetAdditionalImages returns the AdditionalImages field if non-nil, zero value otherwise.

### GetAdditionalImagesOk

`func (o *Product) GetAdditionalImagesOk() (*[]Image, bool)`

GetAdditionalImagesOk returns a tuple with the AdditionalImages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalImages

`func (o *Product) SetAdditionalImages(v []Image)`

SetAdditionalImages sets AdditionalImages field to given value.

### HasAdditionalImages

`func (o *Product) HasAdditionalImages() bool`

HasAdditionalImages returns a boolean if a field has been set.

### GetAdditionalProductIdentities

`func (o *Product) GetAdditionalProductIdentities() []AdditionalProductIdentity`

GetAdditionalProductIdentities returns the AdditionalProductIdentities field if non-nil, zero value otherwise.

### GetAdditionalProductIdentitiesOk

`func (o *Product) GetAdditionalProductIdentitiesOk() (*[]AdditionalProductIdentity, bool)`

GetAdditionalProductIdentitiesOk returns a tuple with the AdditionalProductIdentities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalProductIdentities

`func (o *Product) SetAdditionalProductIdentities(v []AdditionalProductIdentity)`

SetAdditionalProductIdentities sets AdditionalProductIdentities field to given value.

### HasAdditionalProductIdentities

`func (o *Product) HasAdditionalProductIdentities() bool`

HasAdditionalProductIdentities returns a boolean if a field has been set.

### GetAspectGroups

`func (o *Product) GetAspectGroups() []AspectGroup`

GetAspectGroups returns the AspectGroups field if non-nil, zero value otherwise.

### GetAspectGroupsOk

`func (o *Product) GetAspectGroupsOk() (*[]AspectGroup, bool)`

GetAspectGroupsOk returns a tuple with the AspectGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAspectGroups

`func (o *Product) SetAspectGroups(v []AspectGroup)`

SetAspectGroups sets AspectGroups field to given value.

### HasAspectGroups

`func (o *Product) HasAspectGroups() bool`

HasAspectGroups returns a boolean if a field has been set.

### GetBrand

`func (o *Product) GetBrand() string`

GetBrand returns the Brand field if non-nil, zero value otherwise.

### GetBrandOk

`func (o *Product) GetBrandOk() (*string, bool)`

GetBrandOk returns a tuple with the Brand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrand

`func (o *Product) SetBrand(v string)`

SetBrand sets Brand field to given value.

### HasBrand

`func (o *Product) HasBrand() bool`

HasBrand returns a boolean if a field has been set.

### GetDescription

`func (o *Product) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Product) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Product) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Product) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetGtins

`func (o *Product) GetGtins() []string`

GetGtins returns the Gtins field if non-nil, zero value otherwise.

### GetGtinsOk

`func (o *Product) GetGtinsOk() (*[]string, bool)`

GetGtinsOk returns a tuple with the Gtins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGtins

`func (o *Product) SetGtins(v []string)`

SetGtins sets Gtins field to given value.

### HasGtins

`func (o *Product) HasGtins() bool`

HasGtins returns a boolean if a field has been set.

### GetImage

`func (o *Product) GetImage() Image`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *Product) GetImageOk() (*Image, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *Product) SetImage(v Image)`

SetImage sets Image field to given value.

### HasImage

`func (o *Product) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetMpns

`func (o *Product) GetMpns() []string`

GetMpns returns the Mpns field if non-nil, zero value otherwise.

### GetMpnsOk

`func (o *Product) GetMpnsOk() (*[]string, bool)`

GetMpnsOk returns a tuple with the Mpns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMpns

`func (o *Product) SetMpns(v []string)`

SetMpns sets Mpns field to given value.

### HasMpns

`func (o *Product) HasMpns() bool`

HasMpns returns a boolean if a field has been set.

### GetTitle

`func (o *Product) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Product) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Product) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *Product) HasTitle() bool`

HasTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


