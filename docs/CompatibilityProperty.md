# CompatibilityProperty

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LocalizedName** | Pointer to **string** | The name of the product attribute that as been translated to the language of the site. | [optional] 
**Name** | Pointer to **string** | The name of the product attribute, such as Make, Model, Year, etc. | [optional] 
**Value** | Pointer to **string** | The value for the &lt;code&gt;name&lt;/code&gt; attribute, such as &lt;b&gt;BMW&lt;/b&gt;, &lt;b&gt;R1200GS&lt;/b&gt;, &lt;b&gt;2011&lt;/b&gt;, etc. | [optional] 

## Methods

### NewCompatibilityProperty

`func NewCompatibilityProperty() *CompatibilityProperty`

NewCompatibilityProperty instantiates a new CompatibilityProperty object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompatibilityPropertyWithDefaults

`func NewCompatibilityPropertyWithDefaults() *CompatibilityProperty`

NewCompatibilityPropertyWithDefaults instantiates a new CompatibilityProperty object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocalizedName

`func (o *CompatibilityProperty) GetLocalizedName() string`

GetLocalizedName returns the LocalizedName field if non-nil, zero value otherwise.

### GetLocalizedNameOk

`func (o *CompatibilityProperty) GetLocalizedNameOk() (*string, bool)`

GetLocalizedNameOk returns a tuple with the LocalizedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalizedName

`func (o *CompatibilityProperty) SetLocalizedName(v string)`

SetLocalizedName sets LocalizedName field to given value.

### HasLocalizedName

`func (o *CompatibilityProperty) HasLocalizedName() bool`

HasLocalizedName returns a boolean if a field has been set.

### GetName

`func (o *CompatibilityProperty) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CompatibilityProperty) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CompatibilityProperty) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CompatibilityProperty) HasName() bool`

HasName returns a boolean if a field has been set.

### GetValue

`func (o *CompatibilityProperty) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *CompatibilityProperty) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *CompatibilityProperty) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *CompatibilityProperty) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


