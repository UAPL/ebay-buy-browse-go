# AddCartItemInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ItemId** | Pointer to **string** | The eBay RESTful identifier of the item you want added to the cart. &lt;br /&gt;&lt;br /&gt; &lt;b&gt;RESTful Item ID Format: &lt;/b&gt;&lt;code&gt;v1&lt;/code&gt;|&lt;code&gt;&lt;i&gt;#&lt;/i&gt;&lt;/code&gt;|&lt;code&gt;&lt;i&gt;#&lt;/i&gt;&lt;/code&gt; &lt;br /&gt;&lt;b&gt; For example: &lt;/b&gt; &lt;br /&gt;&lt;code&gt;v1|2**********2|0&lt;/code&gt; &lt;br /&gt;&lt;code&gt;v1|1**********2|4**********2&lt;/code&gt; &lt;br /&gt;&lt;br /&gt;For more information about item ID for RESTful APIs, see the &lt;a href&#x3D;\&quot;/api-docs/buy/static/api-browse.html#Legacy\&quot;&gt;Legacy API compatibility&lt;/a&gt; section of the &lt;i&gt;Buy APIs Overview&lt;/i&gt;.&lt;br /&gt;&lt;br /&gt;&lt;b&gt; Maximum number of items in a cart: &lt;/b&gt; 100 | [optional] 
**Quantity** | Pointer to **int32** | The number of this item the buyer wants to purchase. If this value is greater than the number available, the service will change this value to the number available. If this happens, a warning is returned.&lt;br /&gt;&lt;br /&gt;&lt;b&gt; Maximum: &lt;/b&gt; &lt;i&gt;number available&lt;/i&gt; | [optional] 

## Methods

### NewAddCartItemInput

`func NewAddCartItemInput() *AddCartItemInput`

NewAddCartItemInput instantiates a new AddCartItemInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddCartItemInputWithDefaults

`func NewAddCartItemInputWithDefaults() *AddCartItemInput`

NewAddCartItemInputWithDefaults instantiates a new AddCartItemInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItemId

`func (o *AddCartItemInput) GetItemId() string`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *AddCartItemInput) GetItemIdOk() (*string, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *AddCartItemInput) SetItemId(v string)`

SetItemId sets ItemId field to given value.

### HasItemId

`func (o *AddCartItemInput) HasItemId() bool`

HasItemId returns a boolean if a field has been set.

### GetQuantity

`func (o *AddCartItemInput) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *AddCartItemInput) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *AddCartItemInput) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *AddCartItemInput) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


