# EstimatedAvailability

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvailabilityThreshold** | Pointer to **int32** | This field is return only when the seller sets their &#39;&lt;a href&#x3D;\&quot;#display-item-quantity\&quot;&gt;display item quantity&lt;/a&gt;&#39; preference to &lt;b&gt; Display \&quot;More than 10 available\&quot; in your listing (if applicable)&lt;/b&gt;. The value of this field will be \&quot;10\&quot;, which is the threshold value. &lt;br&gt;&lt;br&gt;Code so that your app gracefully handles any future changes to this value. | [optional] 
**AvailabilityThresholdType** | Pointer to **string** | &lt;a name&#x3D;\&quot;display-item-quantity\&quot;&gt;&lt;/a&gt; This field is return only when the seller sets their &lt;b&gt; Display Item Quantity&lt;/b&gt; preference to &lt;b&gt; Display \&quot;More than 10 available\&quot; in your listing (if applicable)&lt;/b&gt;. The value of this field will be &lt;code&gt; MORE_THAN&lt;/code&gt;. This indicates that the seller has more than the &#39;quantity display preference&#39;, which is 10, in stock for this item.    &lt;br&gt;&lt;br&gt; The following are the display item quantity preferences the seller can set. &lt;br&gt;&lt;ul&gt;&lt;li&gt; &lt;b&gt; Display \&quot;More than 10 available\&quot; in your listing (if applicable) &lt;/b&gt;&lt;ul&gt; &lt;li&gt;If the seller enables this preference, this field is returned as long as there are more than 10 of this item in inventory.&lt;/li&gt;  &lt;li&gt; If the quantity is equal to 10 or drops below 10, this field is not returned and the estimated quantity of the item is returned in the &lt;b&gt; estimatedAvailableQuantity&lt;/b&gt; field.&lt;/li&gt;&lt;/ul&gt; &lt;/li&gt; &lt;li&gt; &lt;b&gt; Display the exact quantity in your items&lt;/b&gt; &lt;br&gt;If the seller enables this preference, the &lt;b&gt; availabilityThresholdType&lt;/b&gt; and &lt;b&gt; availabilityThreshold&lt;/b&gt; fields are not returned and the estimated quantity of the item is returned in the &lt;b&gt; estimatedAvailableQuantity&lt;/b&gt; field.&lt;br&gt;&lt;br&gt;&lt;b&gt; Note: &lt;/b&gt; Because the quantity of an item can change several times within a second, it is impossible to return the exact quantity. &lt;/li&gt;&lt;/ul&gt;   &lt;br&gt;Code so that your app gracefully handles any future changes to these preferences. For implementation help, refer to &lt;a href&#x3D;&#39;https://developer.ebay.com/api-docs/buy/browse/types/gct:AvailabilityThresholdEnum&#39;&gt;eBay API documentation&lt;/a&gt; | [optional] 
**DeliveryOptions** | Pointer to **[]string** | An array of available delivery options. &lt;br&gt;&lt;br&gt;&lt;b&gt; Valid Values: &lt;/b&gt; SHIP_TO_HOME, SELLER_ARRANGED_LOCAL_PICKUP, IN_STORE_PICKUP, PICKUP_DROP_OFF, or DIGITAL_DELIVERY &lt;br&gt;&lt;br&gt;Code so that your app gracefully handles any future changes to this list.  | [optional] 
**EstimatedAvailabilityStatus** | Pointer to **string** | An enumeration value representing the inventory status of this item.&lt;br&gt;&lt;br&gt;&lt;span class&#x3D;\&quot;tablenote\&quot;&gt;&lt;b&gt; Note: &lt;/b&gt;Be sure to review the &lt;b&gt;itemEndDate&lt;/b&gt; field to determine whether the item is available for purchase.&lt;/span&gt;&lt;br&gt;&lt;br&gt;&lt;b&gt; Valid Values: &lt;/b&gt; IN_STOCK, LIMITED_STOCK, or OUT_OF_STOCK &lt;br&gt;&lt;br&gt;Code so that your app gracefully handles any future changes to this list. For implementation help, refer to &lt;a href&#x3D;&#39;https://developer.ebay.com/api-docs/buy/browse/types/gct:AvailabilityStatusEnum&#39;&gt;eBay API documentation&lt;/a&gt; | [optional] 
**EstimatedAvailableQuantity** | Pointer to **int32** | The estimated number of this item that are available for purchase. Because the quantity of an item can change several times within a second, it is impossible to return the exact quantity. So instead of returning quantity, the estimated availability of the item is returned.&lt;br&gt;&lt;br&gt;&lt;span class&#x3D;\&quot;tablenote\&quot;&gt;&lt;b&gt;Note:&lt;/b&gt; To see if a listing is available for purchase, review the &lt;a href&#x3D;\&quot;/api-docs/buy/browse/resources/item/methods/getItem#response.itemEndDate\&quot;&gt;itemEndDate&lt;/a&gt; and &lt;a href&#x3D;\&quot;/api-docs/buy/browse/resources/item/methods/getItem#response.estimatedAvailabilities.estimatedAvailabilityStatus\&quot;&gt;estimatedAvailablityStatus&lt;/a&gt; fields. If the item has an &lt;b&gt;EndDate&lt;/b&gt; in the past, or the &lt;b&gt;estimatedAvailabilityStatus&lt;/b&gt; is &lt;code&gt;OUT_OF_STOCK&lt;/code&gt;, the item is unavailable for purchase.&lt;/span&gt; | [optional] 
**EstimatedRemainingQuantity** | Pointer to **int32** | The estimated number of this item that are available for purchase. Because the quantity of an item can change several times within a second, it is impossible to return the exact quantity. So instead of returning quantity, the estimated availability of the item is returned.&lt;br&gt;&lt;br&gt;&lt;span class&#x3D;\&quot;tablenote\&quot;&gt;&lt;b&gt;Note:&lt;/b&gt; To see if a listing is available for purchase, review the &lt;a href&#x3D;\&quot;/api-docs/buy/browse/resources/item/methods/getItem#response.itemEndDate\&quot;&gt;itemEndDate&lt;/a&gt; and &lt;a href&#x3D;\&quot;/api-docs/buy/browse/resources/item/methods/getItem#response.estimatedAvailabilities.estimatedAvailabilityStatus\&quot;&gt;estimatedAvailablityStatus&lt;/a&gt; fields. If the item has an &lt;b&gt;EndDate&lt;/b&gt; in the past, or the &lt;b&gt;estimatedAvailabilityStatus&lt;/b&gt; is &lt;code&gt;OUT_OF_STOCK&lt;/code&gt;, the item is unavailable for purchase.&lt;/span&gt; | [optional] 
**EstimatedSoldQuantity** | Pointer to **int32** | The estimated number of this item that have been sold. | [optional] 

## Methods

### NewEstimatedAvailability

`func NewEstimatedAvailability() *EstimatedAvailability`

NewEstimatedAvailability instantiates a new EstimatedAvailability object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEstimatedAvailabilityWithDefaults

`func NewEstimatedAvailabilityWithDefaults() *EstimatedAvailability`

NewEstimatedAvailabilityWithDefaults instantiates a new EstimatedAvailability object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailabilityThreshold

`func (o *EstimatedAvailability) GetAvailabilityThreshold() int32`

GetAvailabilityThreshold returns the AvailabilityThreshold field if non-nil, zero value otherwise.

### GetAvailabilityThresholdOk

`func (o *EstimatedAvailability) GetAvailabilityThresholdOk() (*int32, bool)`

GetAvailabilityThresholdOk returns a tuple with the AvailabilityThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityThreshold

`func (o *EstimatedAvailability) SetAvailabilityThreshold(v int32)`

SetAvailabilityThreshold sets AvailabilityThreshold field to given value.

### HasAvailabilityThreshold

`func (o *EstimatedAvailability) HasAvailabilityThreshold() bool`

HasAvailabilityThreshold returns a boolean if a field has been set.

### GetAvailabilityThresholdType

`func (o *EstimatedAvailability) GetAvailabilityThresholdType() string`

GetAvailabilityThresholdType returns the AvailabilityThresholdType field if non-nil, zero value otherwise.

### GetAvailabilityThresholdTypeOk

`func (o *EstimatedAvailability) GetAvailabilityThresholdTypeOk() (*string, bool)`

GetAvailabilityThresholdTypeOk returns a tuple with the AvailabilityThresholdType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityThresholdType

`func (o *EstimatedAvailability) SetAvailabilityThresholdType(v string)`

SetAvailabilityThresholdType sets AvailabilityThresholdType field to given value.

### HasAvailabilityThresholdType

`func (o *EstimatedAvailability) HasAvailabilityThresholdType() bool`

HasAvailabilityThresholdType returns a boolean if a field has been set.

### GetDeliveryOptions

`func (o *EstimatedAvailability) GetDeliveryOptions() []string`

GetDeliveryOptions returns the DeliveryOptions field if non-nil, zero value otherwise.

### GetDeliveryOptionsOk

`func (o *EstimatedAvailability) GetDeliveryOptionsOk() (*[]string, bool)`

GetDeliveryOptionsOk returns a tuple with the DeliveryOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryOptions

`func (o *EstimatedAvailability) SetDeliveryOptions(v []string)`

SetDeliveryOptions sets DeliveryOptions field to given value.

### HasDeliveryOptions

`func (o *EstimatedAvailability) HasDeliveryOptions() bool`

HasDeliveryOptions returns a boolean if a field has been set.

### GetEstimatedAvailabilityStatus

`func (o *EstimatedAvailability) GetEstimatedAvailabilityStatus() string`

GetEstimatedAvailabilityStatus returns the EstimatedAvailabilityStatus field if non-nil, zero value otherwise.

### GetEstimatedAvailabilityStatusOk

`func (o *EstimatedAvailability) GetEstimatedAvailabilityStatusOk() (*string, bool)`

GetEstimatedAvailabilityStatusOk returns a tuple with the EstimatedAvailabilityStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedAvailabilityStatus

`func (o *EstimatedAvailability) SetEstimatedAvailabilityStatus(v string)`

SetEstimatedAvailabilityStatus sets EstimatedAvailabilityStatus field to given value.

### HasEstimatedAvailabilityStatus

`func (o *EstimatedAvailability) HasEstimatedAvailabilityStatus() bool`

HasEstimatedAvailabilityStatus returns a boolean if a field has been set.

### GetEstimatedAvailableQuantity

`func (o *EstimatedAvailability) GetEstimatedAvailableQuantity() int32`

GetEstimatedAvailableQuantity returns the EstimatedAvailableQuantity field if non-nil, zero value otherwise.

### GetEstimatedAvailableQuantityOk

`func (o *EstimatedAvailability) GetEstimatedAvailableQuantityOk() (*int32, bool)`

GetEstimatedAvailableQuantityOk returns a tuple with the EstimatedAvailableQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedAvailableQuantity

`func (o *EstimatedAvailability) SetEstimatedAvailableQuantity(v int32)`

SetEstimatedAvailableQuantity sets EstimatedAvailableQuantity field to given value.

### HasEstimatedAvailableQuantity

`func (o *EstimatedAvailability) HasEstimatedAvailableQuantity() bool`

HasEstimatedAvailableQuantity returns a boolean if a field has been set.

### GetEstimatedRemainingQuantity

`func (o *EstimatedAvailability) GetEstimatedRemainingQuantity() int32`

GetEstimatedRemainingQuantity returns the EstimatedRemainingQuantity field if non-nil, zero value otherwise.

### GetEstimatedRemainingQuantityOk

`func (o *EstimatedAvailability) GetEstimatedRemainingQuantityOk() (*int32, bool)`

GetEstimatedRemainingQuantityOk returns a tuple with the EstimatedRemainingQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedRemainingQuantity

`func (o *EstimatedAvailability) SetEstimatedRemainingQuantity(v int32)`

SetEstimatedRemainingQuantity sets EstimatedRemainingQuantity field to given value.

### HasEstimatedRemainingQuantity

`func (o *EstimatedAvailability) HasEstimatedRemainingQuantity() bool`

HasEstimatedRemainingQuantity returns a boolean if a field has been set.

### GetEstimatedSoldQuantity

`func (o *EstimatedAvailability) GetEstimatedSoldQuantity() int32`

GetEstimatedSoldQuantity returns the EstimatedSoldQuantity field if non-nil, zero value otherwise.

### GetEstimatedSoldQuantityOk

`func (o *EstimatedAvailability) GetEstimatedSoldQuantityOk() (*int32, bool)`

GetEstimatedSoldQuantityOk returns a tuple with the EstimatedSoldQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedSoldQuantity

`func (o *EstimatedAvailability) SetEstimatedSoldQuantity(v int32)`

SetEstimatedSoldQuantity sets EstimatedSoldQuantity field to given value.

### HasEstimatedSoldQuantity

`func (o *EstimatedAvailability) HasEstimatedSoldQuantity() bool`

HasEstimatedSoldQuantity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


