# ConditionDescriptorValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalInfo** | Pointer to **[]string** | Additional information about the condition of an item as it relates to a condition descriptor. This array elaborates on the value specified in the &lt;b&gt;content&lt;/b&gt; field and provides additional details about the condition of an item. | [optional] 
**Content** | Pointer to **string** | The value for the condition descriptor indicated in the associated &lt;b&gt;name&lt;/b&gt; field. | [optional] 

## Methods

### NewConditionDescriptorValue

`func NewConditionDescriptorValue() *ConditionDescriptorValue`

NewConditionDescriptorValue instantiates a new ConditionDescriptorValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConditionDescriptorValueWithDefaults

`func NewConditionDescriptorValueWithDefaults() *ConditionDescriptorValue`

NewConditionDescriptorValueWithDefaults instantiates a new ConditionDescriptorValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdditionalInfo

`func (o *ConditionDescriptorValue) GetAdditionalInfo() []string`

GetAdditionalInfo returns the AdditionalInfo field if non-nil, zero value otherwise.

### GetAdditionalInfoOk

`func (o *ConditionDescriptorValue) GetAdditionalInfoOk() (*[]string, bool)`

GetAdditionalInfoOk returns a tuple with the AdditionalInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalInfo

`func (o *ConditionDescriptorValue) SetAdditionalInfo(v []string)`

SetAdditionalInfo sets AdditionalInfo field to given value.

### HasAdditionalInfo

`func (o *ConditionDescriptorValue) HasAdditionalInfo() bool`

HasAdditionalInfo returns a boolean if a field has been set.

### GetContent

`func (o *ConditionDescriptorValue) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *ConditionDescriptorValue) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *ConditionDescriptorValue) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *ConditionDescriptorValue) HasContent() bool`

HasContent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


