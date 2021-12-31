# CompatibilityPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompatibilityProperties** | Pointer to [**[]AttributeNameValue**](AttributeNameValue.md) | An array of attribute name/value pairs used to define a specific product. For example: If you wanted to specify a specific car, one of the name/value pairs would be &lt;br /&gt;&lt;code&gt;\&quot;name\&quot; : \&quot;Year\&quot;, &lt;br /&gt;\&quot;value\&quot; : \&quot;2019\&quot;&lt;/code&gt;  &lt;p&gt; For a list of the attributes required for cars and trucks and motorcycles see &lt;a href&#x3D;\&quot;/api-docs/buy/static/api-browse.html#Check\&quot;&gt;Check compatibility&lt;/a&gt; in the Buy Integration Guide.&lt;/p&gt; | [optional] 

## Methods

### NewCompatibilityPayload

`func NewCompatibilityPayload() *CompatibilityPayload`

NewCompatibilityPayload instantiates a new CompatibilityPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompatibilityPayloadWithDefaults

`func NewCompatibilityPayloadWithDefaults() *CompatibilityPayload`

NewCompatibilityPayloadWithDefaults instantiates a new CompatibilityPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompatibilityProperties

`func (o *CompatibilityPayload) GetCompatibilityProperties() []AttributeNameValue`

GetCompatibilityProperties returns the CompatibilityProperties field if non-nil, zero value otherwise.

### GetCompatibilityPropertiesOk

`func (o *CompatibilityPayload) GetCompatibilityPropertiesOk() (*[]AttributeNameValue, bool)`

GetCompatibilityPropertiesOk returns a tuple with the CompatibilityProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompatibilityProperties

`func (o *CompatibilityPayload) SetCompatibilityProperties(v []AttributeNameValue)`

SetCompatibilityProperties sets CompatibilityProperties field to given value.

### HasCompatibilityProperties

`func (o *CompatibilityPayload) HasCompatibilityProperties() bool`

HasCompatibilityProperties returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


