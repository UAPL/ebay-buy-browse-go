# AttributeNameValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the product attribute, such as &lt;i&gt;Make&lt;/i&gt;, &lt;i&gt;Model&lt;/i&gt;, &lt;i&gt;Year&lt;/i&gt;, etc. | [optional] 
**Value** | Pointer to **string** | The value for the &lt;code&gt;name&lt;/code&gt; attribute, such as &lt;i&gt;BMW&lt;/i&gt;, &lt;i&gt;R1200GS&lt;/i&gt;, &lt;i&gt;2011&lt;/i&gt;, etc. | [optional] 

## Methods

### NewAttributeNameValue

`func NewAttributeNameValue() *AttributeNameValue`

NewAttributeNameValue instantiates a new AttributeNameValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttributeNameValueWithDefaults

`func NewAttributeNameValueWithDefaults() *AttributeNameValue`

NewAttributeNameValueWithDefaults instantiates a new AttributeNameValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AttributeNameValue) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AttributeNameValue) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AttributeNameValue) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AttributeNameValue) HasName() bool`

HasName returns a boolean if a field has been set.

### GetValue

`func (o *AttributeNameValue) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *AttributeNameValue) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *AttributeNameValue) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *AttributeNameValue) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


