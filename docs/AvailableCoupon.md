# AvailableCoupon

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Constraint** | Pointer to [**CouponConstraint**](CouponConstraint.md) |  | [optional] 
**DiscountAmount** | Pointer to [**Amount**](Amount.md) |  | [optional] 
**DiscountType** | Pointer to **string** | The type of discount that the coupon applies. For implementation help, refer to &lt;a href&#x3D;&#39;https://developer.ebay.com/api-docs/buy/browse/types/gct:CouponDiscountType&#39;&gt;eBay API documentation&lt;/a&gt; | [optional] 
**Message** | Pointer to **string** | A description of the coupon.&lt;br /&gt;&lt;br /&gt;&lt;span class&#x3D;\&quot;tablenote\&quot;&gt;&lt;b&gt; Note: &lt;/b&gt; The value returned in the &lt;b&gt;termsWebUrl&lt;/b&gt; field should appear for all experiences when displaying coupons. The value in the &lt;b&gt;availableCoupons.message&lt;/b&gt; field must also be included, if returned in the API response.&lt;/span&gt; | [optional] 
**RedemptionCode** | Pointer to **string** | The coupon code. | [optional] 
**TermsWebUrl** | Pointer to **string** | The URL to the coupon terms of use.&lt;br /&gt;&lt;br /&gt;&lt;span class&#x3D;\&quot;tablenote\&quot;&gt;&lt;b&gt; Note: &lt;/b&gt; The value returned in the &lt;b&gt;termsWebUrl&lt;/b&gt; field should appear for all experiences when displaying coupons. The value in the &lt;b&gt;availableCoupons.message&lt;/b&gt; field must also be included, if returned in the API response.&lt;/span&gt; | [optional] 

## Methods

### NewAvailableCoupon

`func NewAvailableCoupon() *AvailableCoupon`

NewAvailableCoupon instantiates a new AvailableCoupon object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAvailableCouponWithDefaults

`func NewAvailableCouponWithDefaults() *AvailableCoupon`

NewAvailableCouponWithDefaults instantiates a new AvailableCoupon object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConstraint

`func (o *AvailableCoupon) GetConstraint() CouponConstraint`

GetConstraint returns the Constraint field if non-nil, zero value otherwise.

### GetConstraintOk

`func (o *AvailableCoupon) GetConstraintOk() (*CouponConstraint, bool)`

GetConstraintOk returns a tuple with the Constraint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstraint

`func (o *AvailableCoupon) SetConstraint(v CouponConstraint)`

SetConstraint sets Constraint field to given value.

### HasConstraint

`func (o *AvailableCoupon) HasConstraint() bool`

HasConstraint returns a boolean if a field has been set.

### GetDiscountAmount

`func (o *AvailableCoupon) GetDiscountAmount() Amount`

GetDiscountAmount returns the DiscountAmount field if non-nil, zero value otherwise.

### GetDiscountAmountOk

`func (o *AvailableCoupon) GetDiscountAmountOk() (*Amount, bool)`

GetDiscountAmountOk returns a tuple with the DiscountAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountAmount

`func (o *AvailableCoupon) SetDiscountAmount(v Amount)`

SetDiscountAmount sets DiscountAmount field to given value.

### HasDiscountAmount

`func (o *AvailableCoupon) HasDiscountAmount() bool`

HasDiscountAmount returns a boolean if a field has been set.

### GetDiscountType

`func (o *AvailableCoupon) GetDiscountType() string`

GetDiscountType returns the DiscountType field if non-nil, zero value otherwise.

### GetDiscountTypeOk

`func (o *AvailableCoupon) GetDiscountTypeOk() (*string, bool)`

GetDiscountTypeOk returns a tuple with the DiscountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountType

`func (o *AvailableCoupon) SetDiscountType(v string)`

SetDiscountType sets DiscountType field to given value.

### HasDiscountType

`func (o *AvailableCoupon) HasDiscountType() bool`

HasDiscountType returns a boolean if a field has been set.

### GetMessage

`func (o *AvailableCoupon) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *AvailableCoupon) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *AvailableCoupon) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *AvailableCoupon) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetRedemptionCode

`func (o *AvailableCoupon) GetRedemptionCode() string`

GetRedemptionCode returns the RedemptionCode field if non-nil, zero value otherwise.

### GetRedemptionCodeOk

`func (o *AvailableCoupon) GetRedemptionCodeOk() (*string, bool)`

GetRedemptionCodeOk returns a tuple with the RedemptionCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedemptionCode

`func (o *AvailableCoupon) SetRedemptionCode(v string)`

SetRedemptionCode sets RedemptionCode field to given value.

### HasRedemptionCode

`func (o *AvailableCoupon) HasRedemptionCode() bool`

HasRedemptionCode returns a boolean if a field has been set.

### GetTermsWebUrl

`func (o *AvailableCoupon) GetTermsWebUrl() string`

GetTermsWebUrl returns the TermsWebUrl field if non-nil, zero value otherwise.

### GetTermsWebUrlOk

`func (o *AvailableCoupon) GetTermsWebUrlOk() (*string, bool)`

GetTermsWebUrlOk returns a tuple with the TermsWebUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermsWebUrl

`func (o *AvailableCoupon) SetTermsWebUrl(v string)`

SetTermsWebUrl sets TermsWebUrl field to given value.

### HasTermsWebUrl

`func (o *AvailableCoupon) HasTermsWebUrl() bool`

HasTermsWebUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


