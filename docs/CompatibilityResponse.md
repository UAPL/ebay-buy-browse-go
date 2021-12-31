# CompatibilityResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompatibilityStatus** | Pointer to **string** | An enumeration value that tells you if the item is compatible with the product. &lt;br /&gt;&lt;br /&gt;The values are: &lt;ul&gt;   &lt;li&gt;   &lt;b&gt; COMPATIBLE&lt;/b&gt; - Indicates the item is compatible with the product specified in the request.&lt;/li&gt;   &lt;li&gt;   &lt;b&gt; NOT_COMPATIBLE&lt;/b&gt; - Indicates the item is not compatible with the product specified in the request. Be sure to check all the &lt;b&gt; value&lt;/b&gt; fields to ensure they are correct as errors in the value can also cause this response.&lt;/li&gt;   &lt;li&gt; &lt;b&gt; UNDETERMINED&lt;/b&gt; - Indicates one or more attributes for the specified product are missing so compatibility cannot be determined.  The response returns the attributes that are missing.&lt;/li&gt;  &lt;/ul&gt;  Code so that your app gracefully handles any future changes to this list. For implementation help, refer to &lt;a href&#x3D;&#39;https://developer.ebay.com/api-docs/buy/browse/types/gct:CompatibilityStatus&#39;&gt;eBay API documentation&lt;/a&gt; | [optional] 
**Warnings** | Pointer to [**[]Error**](Error.md) | An array of warning messages. These types of errors do not prevent the method from executing but should be checked. | [optional] 

## Methods

### NewCompatibilityResponse

`func NewCompatibilityResponse() *CompatibilityResponse`

NewCompatibilityResponse instantiates a new CompatibilityResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompatibilityResponseWithDefaults

`func NewCompatibilityResponseWithDefaults() *CompatibilityResponse`

NewCompatibilityResponseWithDefaults instantiates a new CompatibilityResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompatibilityStatus

`func (o *CompatibilityResponse) GetCompatibilityStatus() string`

GetCompatibilityStatus returns the CompatibilityStatus field if non-nil, zero value otherwise.

### GetCompatibilityStatusOk

`func (o *CompatibilityResponse) GetCompatibilityStatusOk() (*string, bool)`

GetCompatibilityStatusOk returns a tuple with the CompatibilityStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompatibilityStatus

`func (o *CompatibilityResponse) SetCompatibilityStatus(v string)`

SetCompatibilityStatus sets CompatibilityStatus field to given value.

### HasCompatibilityStatus

`func (o *CompatibilityResponse) HasCompatibilityStatus() bool`

HasCompatibilityStatus returns a boolean if a field has been set.

### GetWarnings

`func (o *CompatibilityResponse) GetWarnings() []Error`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *CompatibilityResponse) GetWarningsOk() (*[]Error, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *CompatibilityResponse) SetWarnings(v []Error)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *CompatibilityResponse) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


