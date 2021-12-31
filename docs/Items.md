# Items

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]CoreItem**](CoreItem.md) | An arraylist of all the items. | [optional] 
**Total** | Pointer to **int32** | The total number of items retrieved. | [optional] 
**Warnings** | Pointer to [**[]Error**](Error.md) | An array of warning messages. These types of errors do not prevent the method from executing but should be checked. | [optional] 

## Methods

### NewItems

`func NewItems() *Items`

NewItems instantiates a new Items object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItemsWithDefaults

`func NewItemsWithDefaults() *Items`

NewItemsWithDefaults instantiates a new Items object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *Items) GetItems() []CoreItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *Items) GetItemsOk() (*[]CoreItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *Items) SetItems(v []CoreItem)`

SetItems sets Items field to given value.

### HasItems

`func (o *Items) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetTotal

`func (o *Items) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *Items) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *Items) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *Items) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetWarnings

`func (o *Items) GetWarnings() []Error`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *Items) GetWarningsOk() (*[]Error, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *Items) SetWarnings(v []Error)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *Items) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


