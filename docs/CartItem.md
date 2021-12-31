# CartItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CartItemId** | Pointer to **string** | The identifier for the item being added to the cart. This is generated when the item is added to the cart. | [optional] 
**CartItemSubtotal** | Pointer to [**Amount**](Amount.md) |  | [optional] 
**Image** | Pointer to [**Image**](Image.md) |  | [optional] 
**ItemId** | Pointer to **string** | The RESTful identifier of the item. This identifier is generated when the item was listed. &lt;br /&gt;&lt;br /&gt; &lt;b&gt;RESTful Item ID Format: &lt;/b&gt;&lt;code&gt;v1&lt;/code&gt;|&lt;code&gt;&lt;i&gt;#&lt;/i&gt;&lt;/code&gt;|&lt;code&gt;&lt;i&gt;#&lt;/i&gt;&lt;/code&gt; &lt;br /&gt;&lt;b&gt; For example: &lt;/b&gt;&lt;br /&gt; &lt;code&gt;v1|2**********2|0&lt;/code&gt; &lt;br /&gt;&lt;code&gt;v1|1**********2|4**********2&lt;/code&gt; | [optional] 
**ItemWebUrl** | Pointer to **string** | The URL of the eBay view item page for the item. | [optional] 
**Price** | Pointer to [**Price**](Price.md) |  | [optional] 
**Quantity** | Pointer to **int32** | The number of this item the buyer wants to purchase. | [optional] 
**Title** | Pointer to **string** | The title of the item. This can be written by the seller or come from the eBay product catalog. | [optional] 

## Methods

### NewCartItem

`func NewCartItem() *CartItem`

NewCartItem instantiates a new CartItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCartItemWithDefaults

`func NewCartItemWithDefaults() *CartItem`

NewCartItemWithDefaults instantiates a new CartItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCartItemId

`func (o *CartItem) GetCartItemId() string`

GetCartItemId returns the CartItemId field if non-nil, zero value otherwise.

### GetCartItemIdOk

`func (o *CartItem) GetCartItemIdOk() (*string, bool)`

GetCartItemIdOk returns a tuple with the CartItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCartItemId

`func (o *CartItem) SetCartItemId(v string)`

SetCartItemId sets CartItemId field to given value.

### HasCartItemId

`func (o *CartItem) HasCartItemId() bool`

HasCartItemId returns a boolean if a field has been set.

### GetCartItemSubtotal

`func (o *CartItem) GetCartItemSubtotal() Amount`

GetCartItemSubtotal returns the CartItemSubtotal field if non-nil, zero value otherwise.

### GetCartItemSubtotalOk

`func (o *CartItem) GetCartItemSubtotalOk() (*Amount, bool)`

GetCartItemSubtotalOk returns a tuple with the CartItemSubtotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCartItemSubtotal

`func (o *CartItem) SetCartItemSubtotal(v Amount)`

SetCartItemSubtotal sets CartItemSubtotal field to given value.

### HasCartItemSubtotal

`func (o *CartItem) HasCartItemSubtotal() bool`

HasCartItemSubtotal returns a boolean if a field has been set.

### GetImage

`func (o *CartItem) GetImage() Image`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *CartItem) GetImageOk() (*Image, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *CartItem) SetImage(v Image)`

SetImage sets Image field to given value.

### HasImage

`func (o *CartItem) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetItemId

`func (o *CartItem) GetItemId() string`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *CartItem) GetItemIdOk() (*string, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *CartItem) SetItemId(v string)`

SetItemId sets ItemId field to given value.

### HasItemId

`func (o *CartItem) HasItemId() bool`

HasItemId returns a boolean if a field has been set.

### GetItemWebUrl

`func (o *CartItem) GetItemWebUrl() string`

GetItemWebUrl returns the ItemWebUrl field if non-nil, zero value otherwise.

### GetItemWebUrlOk

`func (o *CartItem) GetItemWebUrlOk() (*string, bool)`

GetItemWebUrlOk returns a tuple with the ItemWebUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemWebUrl

`func (o *CartItem) SetItemWebUrl(v string)`

SetItemWebUrl sets ItemWebUrl field to given value.

### HasItemWebUrl

`func (o *CartItem) HasItemWebUrl() bool`

HasItemWebUrl returns a boolean if a field has been set.

### GetPrice

`func (o *CartItem) GetPrice() Price`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *CartItem) GetPriceOk() (*Price, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *CartItem) SetPrice(v Price)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *CartItem) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetQuantity

`func (o *CartItem) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *CartItem) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *CartItem) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *CartItem) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetTitle

`func (o *CartItem) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CartItem) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CartItem) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *CartItem) HasTitle() bool`

HasTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


