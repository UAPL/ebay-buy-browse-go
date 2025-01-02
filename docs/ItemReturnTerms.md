# ItemReturnTerms

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExtendedHolidayReturnsOffered** | Pointer to **bool** | This indicates if the seller has enabled the Extended Holiday Returns feature on the item. Extended Holiday Returns are only applicable during the US holiday season, and gives buyers extra time to return an item. This &#39;extra time&#39; will typically extend beyond what is set through the &lt;b&gt; returnPeriod&lt;/b&gt; value. | [optional] 
**RefundMethod** | Pointer to **string** | An enumeration value that indicates how a buyer is refunded when an item is returned. &lt;br&gt;&lt;br&gt;&lt;b&gt; Valid Values: &lt;/b&gt; MONEY_BACK or MERCHANDISE_CREDIT  &lt;br&gt;&lt;br&gt;Code so that your app gracefully handles any future changes to this list. For implementation help, refer to &lt;a href&#x3D;&#39;https://developer.ebay.com/api-docs/buy/browse/types/gct:RefundMethodEnum&#39;&gt;eBay API documentation&lt;/a&gt; | [optional] 
**RestockingFeePercentage** | Pointer to **string** | This string field indicates the restocking fee percentage that the seller has set on the item. Sellers have the option of setting no restocking fee for an item, or they can set the percentage to 10, 15, or 20 percent. So, if the cost of the item was $100, and the restocking percentage was 20 percent, the buyer would be charged $20 to return that item, so instead of receiving a $100 refund, they would receive $80 due to the restocking fee. | [optional] 
**ReturnInstructions** | Pointer to **string** | Text written by the seller describing what the buyer needs to do in order to return the item. | [optional] 
**ReturnMethod** | Pointer to **string** | An enumeration value that indicates the alternative methods for a full refund when an item is returned. This field is returned if the seller offers the buyer an item replacement or exchange instead of a monetary refund. &lt;br&gt;&lt;br&gt;&lt;b&gt; Valid Values: &lt;/b&gt;  &lt;ul&gt;&lt;li&gt;&lt;b&gt; REPLACEMENT&lt;/b&gt; -  Indicates that the buyer has the option of receiving money back for the returned item, or they can choose to have the seller replace the item with an identical item.&lt;/li&gt;  &lt;li&gt;&lt;b&gt; EXCHANGE&lt;/b&gt; - Indicates that the buyer has the option of receiving money back for the returned item, or they can exchange the item for another similar item.&lt;/li&gt;&lt;/ul&gt;  Code so that your app gracefully handles any future changes to this list. For implementation help, refer to &lt;a href&#x3D;&#39;https://developer.ebay.com/api-docs/buy/browse/types/gct:ReturnMethodEnum&#39;&gt;eBay API documentation&lt;/a&gt; | [optional] 
**ReturnPeriod** | Pointer to [**TimeDuration**](TimeDuration.md) |  | [optional] 
**ReturnsAccepted** | Pointer to **bool** | Indicates whether the seller accepts returns for the item. | [optional] 
**ReturnShippingCostPayer** | Pointer to **string** | This enumeration value indicates whether the buyer or seller is responsible for return shipping costs when an item is returned. &lt;br&gt;&lt;br&gt;&lt;b&gt; Valid Values: &lt;/b&gt; &lt;ul&gt;&lt;li&gt;&lt;b&gt; SELLER&lt;/b&gt; - Indicates the seller will pay for the shipping costs to return the item.&lt;/li&gt;  &lt;li&gt;&lt;b&gt; BUYER&lt;/b&gt; - Indicates the buyer will pay for the shipping costs to return the item.&lt;/li&gt;  &lt;/ul&gt;  Code so that your app gracefully handles any future changes to this list. For implementation help, refer to &lt;a href&#x3D;&#39;https://developer.ebay.com/api-docs/buy/browse/types/gct:ReturnShippingCostPayerEnum&#39;&gt;eBay API documentation&lt;/a&gt; | [optional] 

## Methods

### NewItemReturnTerms

`func NewItemReturnTerms() *ItemReturnTerms`

NewItemReturnTerms instantiates a new ItemReturnTerms object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItemReturnTermsWithDefaults

`func NewItemReturnTermsWithDefaults() *ItemReturnTerms`

NewItemReturnTermsWithDefaults instantiates a new ItemReturnTerms object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExtendedHolidayReturnsOffered

`func (o *ItemReturnTerms) GetExtendedHolidayReturnsOffered() bool`

GetExtendedHolidayReturnsOffered returns the ExtendedHolidayReturnsOffered field if non-nil, zero value otherwise.

### GetExtendedHolidayReturnsOfferedOk

`func (o *ItemReturnTerms) GetExtendedHolidayReturnsOfferedOk() (*bool, bool)`

GetExtendedHolidayReturnsOfferedOk returns a tuple with the ExtendedHolidayReturnsOffered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendedHolidayReturnsOffered

`func (o *ItemReturnTerms) SetExtendedHolidayReturnsOffered(v bool)`

SetExtendedHolidayReturnsOffered sets ExtendedHolidayReturnsOffered field to given value.

### HasExtendedHolidayReturnsOffered

`func (o *ItemReturnTerms) HasExtendedHolidayReturnsOffered() bool`

HasExtendedHolidayReturnsOffered returns a boolean if a field has been set.

### GetRefundMethod

`func (o *ItemReturnTerms) GetRefundMethod() string`

GetRefundMethod returns the RefundMethod field if non-nil, zero value otherwise.

### GetRefundMethodOk

`func (o *ItemReturnTerms) GetRefundMethodOk() (*string, bool)`

GetRefundMethodOk returns a tuple with the RefundMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundMethod

`func (o *ItemReturnTerms) SetRefundMethod(v string)`

SetRefundMethod sets RefundMethod field to given value.

### HasRefundMethod

`func (o *ItemReturnTerms) HasRefundMethod() bool`

HasRefundMethod returns a boolean if a field has been set.

### GetRestockingFeePercentage

`func (o *ItemReturnTerms) GetRestockingFeePercentage() string`

GetRestockingFeePercentage returns the RestockingFeePercentage field if non-nil, zero value otherwise.

### GetRestockingFeePercentageOk

`func (o *ItemReturnTerms) GetRestockingFeePercentageOk() (*string, bool)`

GetRestockingFeePercentageOk returns a tuple with the RestockingFeePercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestockingFeePercentage

`func (o *ItemReturnTerms) SetRestockingFeePercentage(v string)`

SetRestockingFeePercentage sets RestockingFeePercentage field to given value.

### HasRestockingFeePercentage

`func (o *ItemReturnTerms) HasRestockingFeePercentage() bool`

HasRestockingFeePercentage returns a boolean if a field has been set.

### GetReturnInstructions

`func (o *ItemReturnTerms) GetReturnInstructions() string`

GetReturnInstructions returns the ReturnInstructions field if non-nil, zero value otherwise.

### GetReturnInstructionsOk

`func (o *ItemReturnTerms) GetReturnInstructionsOk() (*string, bool)`

GetReturnInstructionsOk returns a tuple with the ReturnInstructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnInstructions

`func (o *ItemReturnTerms) SetReturnInstructions(v string)`

SetReturnInstructions sets ReturnInstructions field to given value.

### HasReturnInstructions

`func (o *ItemReturnTerms) HasReturnInstructions() bool`

HasReturnInstructions returns a boolean if a field has been set.

### GetReturnMethod

`func (o *ItemReturnTerms) GetReturnMethod() string`

GetReturnMethod returns the ReturnMethod field if non-nil, zero value otherwise.

### GetReturnMethodOk

`func (o *ItemReturnTerms) GetReturnMethodOk() (*string, bool)`

GetReturnMethodOk returns a tuple with the ReturnMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnMethod

`func (o *ItemReturnTerms) SetReturnMethod(v string)`

SetReturnMethod sets ReturnMethod field to given value.

### HasReturnMethod

`func (o *ItemReturnTerms) HasReturnMethod() bool`

HasReturnMethod returns a boolean if a field has been set.

### GetReturnPeriod

`func (o *ItemReturnTerms) GetReturnPeriod() TimeDuration`

GetReturnPeriod returns the ReturnPeriod field if non-nil, zero value otherwise.

### GetReturnPeriodOk

`func (o *ItemReturnTerms) GetReturnPeriodOk() (*TimeDuration, bool)`

GetReturnPeriodOk returns a tuple with the ReturnPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnPeriod

`func (o *ItemReturnTerms) SetReturnPeriod(v TimeDuration)`

SetReturnPeriod sets ReturnPeriod field to given value.

### HasReturnPeriod

`func (o *ItemReturnTerms) HasReturnPeriod() bool`

HasReturnPeriod returns a boolean if a field has been set.

### GetReturnsAccepted

`func (o *ItemReturnTerms) GetReturnsAccepted() bool`

GetReturnsAccepted returns the ReturnsAccepted field if non-nil, zero value otherwise.

### GetReturnsAcceptedOk

`func (o *ItemReturnTerms) GetReturnsAcceptedOk() (*bool, bool)`

GetReturnsAcceptedOk returns a tuple with the ReturnsAccepted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnsAccepted

`func (o *ItemReturnTerms) SetReturnsAccepted(v bool)`

SetReturnsAccepted sets ReturnsAccepted field to given value.

### HasReturnsAccepted

`func (o *ItemReturnTerms) HasReturnsAccepted() bool`

HasReturnsAccepted returns a boolean if a field has been set.

### GetReturnShippingCostPayer

`func (o *ItemReturnTerms) GetReturnShippingCostPayer() string`

GetReturnShippingCostPayer returns the ReturnShippingCostPayer field if non-nil, zero value otherwise.

### GetReturnShippingCostPayerOk

`func (o *ItemReturnTerms) GetReturnShippingCostPayerOk() (*string, bool)`

GetReturnShippingCostPayerOk returns a tuple with the ReturnShippingCostPayer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnShippingCostPayer

`func (o *ItemReturnTerms) SetReturnShippingCostPayer(v string)`

SetReturnShippingCostPayer sets ReturnShippingCostPayer field to given value.

### HasReturnShippingCostPayer

`func (o *ItemReturnTerms) HasReturnShippingCostPayer() bool`

HasReturnShippingCostPayer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


