# ConditionDescriptor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of a condition descriptor. The value(s) for this condition descriptor is returned in the associated &lt;b&gt;values&lt;/b&gt; array. | [optional] 
**Values** | Pointer to [**[]ConditionDescriptorValue**](ConditionDescriptorValue.md) | This array displays the value(s) for a condition descriptor (denoted by the associated &lt;b&gt;name&lt;/b&gt; field), as well as any other additional information about the condition of the item. | [optional] 

## Methods

### NewConditionDescriptor

`func NewConditionDescriptor() *ConditionDescriptor`

NewConditionDescriptor instantiates a new ConditionDescriptor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConditionDescriptorWithDefaults

`func NewConditionDescriptorWithDefaults() *ConditionDescriptor`

NewConditionDescriptorWithDefaults instantiates a new ConditionDescriptor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ConditionDescriptor) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConditionDescriptor) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConditionDescriptor) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ConditionDescriptor) HasName() bool`

HasName returns a boolean if a field has been set.

### GetValues

`func (o *ConditionDescriptor) GetValues() []ConditionDescriptorValue`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *ConditionDescriptor) GetValuesOk() (*[]ConditionDescriptorValue, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *ConditionDescriptor) SetValues(v []ConditionDescriptorValue)`

SetValues sets Values field to given value.

### HasValues

`func (o *ConditionDescriptor) HasValues() bool`

HasValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


