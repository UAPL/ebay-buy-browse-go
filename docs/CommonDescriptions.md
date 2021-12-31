# CommonDescriptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | The item description that is used by more than one of the item variations. | [optional] 
**ItemIds** | Pointer to **[]string** | A list of item ids that have this description. | [optional] 

## Methods

### NewCommonDescriptions

`func NewCommonDescriptions() *CommonDescriptions`

NewCommonDescriptions instantiates a new CommonDescriptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommonDescriptionsWithDefaults

`func NewCommonDescriptionsWithDefaults() *CommonDescriptions`

NewCommonDescriptionsWithDefaults instantiates a new CommonDescriptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *CommonDescriptions) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CommonDescriptions) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CommonDescriptions) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CommonDescriptions) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetItemIds

`func (o *CommonDescriptions) GetItemIds() []string`

GetItemIds returns the ItemIds field if non-nil, zero value otherwise.

### GetItemIdsOk

`func (o *CommonDescriptions) GetItemIdsOk() (*[]string, bool)`

GetItemIdsOk returns a tuple with the ItemIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemIds

`func (o *CommonDescriptions) SetItemIds(v []string)`

SetItemIds sets ItemIds field to given value.

### HasItemIds

`func (o *CommonDescriptions) HasItemIds() bool`

HasItemIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


