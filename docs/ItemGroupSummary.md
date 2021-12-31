# ItemGroupSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ItemGroupAdditionalImages** | Pointer to [**[]Image**](Image.md) | An array of containers with the URLs for images that are in addition to the primary image of the item group.  The primary image is returned in the &lt;b&gt; itemGroupImage&lt;/b&gt; field. | [optional] 
**ItemGroupHref** | Pointer to **string** | The HATEOAS reference of the parent page of the item group. An item group is an item that has various aspect differences, such as color, size, storage capacity, etc.  | [optional] 
**ItemGroupId** | Pointer to **string** | The unique identifier for the item group. An item group is an item that has various aspect differences, such as color, size, storage capacity, etc.  | [optional] 
**ItemGroupImage** | Pointer to [**Image**](Image.md) |  | [optional] 
**ItemGroupTitle** | Pointer to **string** | The title of the item that appears on the item group page. An item group is an item that has various aspect differences, such as color, size, storage capacity, etc.  | [optional] 
**ItemGroupType** | Pointer to **string** | An enumeration value that indicates the type of the item group. An item group is an item that has various aspect differences, such as color, size, storage capacity, etc. For implementation help, refer to &lt;a href&#x3D;&#39;https://developer.ebay.com/api-docs/buy/browse/types/gct:ItemGroupTypeEnum&#39;&gt;eBay API documentation&lt;/a&gt; | [optional] 

## Methods

### NewItemGroupSummary

`func NewItemGroupSummary() *ItemGroupSummary`

NewItemGroupSummary instantiates a new ItemGroupSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItemGroupSummaryWithDefaults

`func NewItemGroupSummaryWithDefaults() *ItemGroupSummary`

NewItemGroupSummaryWithDefaults instantiates a new ItemGroupSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItemGroupAdditionalImages

`func (o *ItemGroupSummary) GetItemGroupAdditionalImages() []Image`

GetItemGroupAdditionalImages returns the ItemGroupAdditionalImages field if non-nil, zero value otherwise.

### GetItemGroupAdditionalImagesOk

`func (o *ItemGroupSummary) GetItemGroupAdditionalImagesOk() (*[]Image, bool)`

GetItemGroupAdditionalImagesOk returns a tuple with the ItemGroupAdditionalImages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemGroupAdditionalImages

`func (o *ItemGroupSummary) SetItemGroupAdditionalImages(v []Image)`

SetItemGroupAdditionalImages sets ItemGroupAdditionalImages field to given value.

### HasItemGroupAdditionalImages

`func (o *ItemGroupSummary) HasItemGroupAdditionalImages() bool`

HasItemGroupAdditionalImages returns a boolean if a field has been set.

### GetItemGroupHref

`func (o *ItemGroupSummary) GetItemGroupHref() string`

GetItemGroupHref returns the ItemGroupHref field if non-nil, zero value otherwise.

### GetItemGroupHrefOk

`func (o *ItemGroupSummary) GetItemGroupHrefOk() (*string, bool)`

GetItemGroupHrefOk returns a tuple with the ItemGroupHref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemGroupHref

`func (o *ItemGroupSummary) SetItemGroupHref(v string)`

SetItemGroupHref sets ItemGroupHref field to given value.

### HasItemGroupHref

`func (o *ItemGroupSummary) HasItemGroupHref() bool`

HasItemGroupHref returns a boolean if a field has been set.

### GetItemGroupId

`func (o *ItemGroupSummary) GetItemGroupId() string`

GetItemGroupId returns the ItemGroupId field if non-nil, zero value otherwise.

### GetItemGroupIdOk

`func (o *ItemGroupSummary) GetItemGroupIdOk() (*string, bool)`

GetItemGroupIdOk returns a tuple with the ItemGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemGroupId

`func (o *ItemGroupSummary) SetItemGroupId(v string)`

SetItemGroupId sets ItemGroupId field to given value.

### HasItemGroupId

`func (o *ItemGroupSummary) HasItemGroupId() bool`

HasItemGroupId returns a boolean if a field has been set.

### GetItemGroupImage

`func (o *ItemGroupSummary) GetItemGroupImage() Image`

GetItemGroupImage returns the ItemGroupImage field if non-nil, zero value otherwise.

### GetItemGroupImageOk

`func (o *ItemGroupSummary) GetItemGroupImageOk() (*Image, bool)`

GetItemGroupImageOk returns a tuple with the ItemGroupImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemGroupImage

`func (o *ItemGroupSummary) SetItemGroupImage(v Image)`

SetItemGroupImage sets ItemGroupImage field to given value.

### HasItemGroupImage

`func (o *ItemGroupSummary) HasItemGroupImage() bool`

HasItemGroupImage returns a boolean if a field has been set.

### GetItemGroupTitle

`func (o *ItemGroupSummary) GetItemGroupTitle() string`

GetItemGroupTitle returns the ItemGroupTitle field if non-nil, zero value otherwise.

### GetItemGroupTitleOk

`func (o *ItemGroupSummary) GetItemGroupTitleOk() (*string, bool)`

GetItemGroupTitleOk returns a tuple with the ItemGroupTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemGroupTitle

`func (o *ItemGroupSummary) SetItemGroupTitle(v string)`

SetItemGroupTitle sets ItemGroupTitle field to given value.

### HasItemGroupTitle

`func (o *ItemGroupSummary) HasItemGroupTitle() bool`

HasItemGroupTitle returns a boolean if a field has been set.

### GetItemGroupType

`func (o *ItemGroupSummary) GetItemGroupType() string`

GetItemGroupType returns the ItemGroupType field if non-nil, zero value otherwise.

### GetItemGroupTypeOk

`func (o *ItemGroupSummary) GetItemGroupTypeOk() (*string, bool)`

GetItemGroupTypeOk returns a tuple with the ItemGroupType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemGroupType

`func (o *ItemGroupSummary) SetItemGroupType(v string)`

SetItemGroupType sets ItemGroupType field to given value.

### HasItemGroupType

`func (o *ItemGroupSummary) HasItemGroupType() bool`

HasItemGroupType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


