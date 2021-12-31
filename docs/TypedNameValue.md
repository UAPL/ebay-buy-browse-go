# TypedNameValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The text representing the name of the aspect for the name/value pair, such as Color. | [optional] 
**Type** | Pointer to **string** | This indicates if the value being returned is a string or an array of values. &lt;br&gt;&lt;br&gt;&lt;b&gt; Valid Values: &lt;/b&gt; &lt;ul&gt;&lt;li&gt;&lt;b&gt; STRING&lt;/b&gt; - Indicates the value returned is a string.&lt;/li&gt;  &lt;li&gt;&lt;b&gt; STRING_ARRAY&lt;/b&gt; - Indicates the value returned is an array of strings.&lt;/li&gt;&lt;/ul&gt;  Code so that your app gracefully handles any future changes to this list. For implementation help, refer to &lt;a href&#x3D;&#39;https://developer.ebay.com/api-docs/buy/browse/types/gct:ValueTypeEnum&#39;&gt;eBay API documentation&lt;/a&gt; | [optional] 
**Value** | Pointer to **string** | The value of the aspect for the name/value pair, such as Red. | [optional] 

## Methods

### NewTypedNameValue

`func NewTypedNameValue() *TypedNameValue`

NewTypedNameValue instantiates a new TypedNameValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTypedNameValueWithDefaults

`func NewTypedNameValueWithDefaults() *TypedNameValue`

NewTypedNameValueWithDefaults instantiates a new TypedNameValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *TypedNameValue) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TypedNameValue) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TypedNameValue) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TypedNameValue) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *TypedNameValue) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TypedNameValue) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TypedNameValue) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *TypedNameValue) HasType() bool`

HasType returns a boolean if a field has been set.

### GetValue

`func (o *TypedNameValue) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *TypedNameValue) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *TypedNameValue) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *TypedNameValue) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


