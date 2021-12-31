# ItemGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CommonDescriptions** | Pointer to [**[]CommonDescriptions**](CommonDescriptions.md) | An array of containers for a description and the item IDs of all the items that have this exact description. Often the item variations within an item group all have the same description. Instead of repeating this description in the item details of each item, a description that is shared by at least one other item is returned in this container. If the description is unique, it is returned in the &lt;b&gt; items.description&lt;/b&gt; field. | [optional] 
**Items** | Pointer to [**[]Item**](Item.md) | An array of containers for all the item variation details, excluding the description. | [optional] 
**Warnings** | Pointer to [**[]Error**](Error.md) | An array of warning messages. These types of errors do not prevent the method from executing but should be checked. | [optional] 

## Methods

### NewItemGroup

`func NewItemGroup() *ItemGroup`

NewItemGroup instantiates a new ItemGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItemGroupWithDefaults

`func NewItemGroupWithDefaults() *ItemGroup`

NewItemGroupWithDefaults instantiates a new ItemGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommonDescriptions

`func (o *ItemGroup) GetCommonDescriptions() []CommonDescriptions`

GetCommonDescriptions returns the CommonDescriptions field if non-nil, zero value otherwise.

### GetCommonDescriptionsOk

`func (o *ItemGroup) GetCommonDescriptionsOk() (*[]CommonDescriptions, bool)`

GetCommonDescriptionsOk returns a tuple with the CommonDescriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommonDescriptions

`func (o *ItemGroup) SetCommonDescriptions(v []CommonDescriptions)`

SetCommonDescriptions sets CommonDescriptions field to given value.

### HasCommonDescriptions

`func (o *ItemGroup) HasCommonDescriptions() bool`

HasCommonDescriptions returns a boolean if a field has been set.

### GetItems

`func (o *ItemGroup) GetItems() []Item`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ItemGroup) GetItemsOk() (*[]Item, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ItemGroup) SetItems(v []Item)`

SetItems sets Items field to given value.

### HasItems

`func (o *ItemGroup) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetWarnings

`func (o *ItemGroup) GetWarnings() []Error`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *ItemGroup) GetWarningsOk() (*[]Error, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *ItemGroup) SetWarnings(v []Error)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *ItemGroup) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


