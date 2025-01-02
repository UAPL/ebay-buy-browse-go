# SellerDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FeedbackPercentage** | Pointer to **string** | The percentage of the total positive feedback. | [optional] 
**FeedbackScore** | Pointer to **int32** | The feedback score of the seller. This value is based on the ratings from eBay members that bought items from this seller. | [optional] 
**SellerAccountType** | Pointer to **string** | This indicates if the seller is a business or an individual. This is determined when the seller registers with eBay. If they register for a business account, this value will be BUSINESS. If they register for a private account, this value will be INDIVIDUAL. This designation is required by the tax laws in the following countries:  &lt;br&gt;&lt;br&gt; This field is returned only on the following sites. &lt;br&gt;&lt;br&gt;EBAY_AT, EBAY_BE, EBAY_CH, EBAY_DE, EBAY_ES, EBAY_FR, EBAY_GB, EBAY_IE, EBAY_IT, EBAY_PL &lt;br&gt;&lt;br&gt;&lt;b&gt; Valid Values:&lt;/b&gt; BUSINESS or INDIVIDUAL &lt;br&gt;&lt;br&gt;Code so that your app gracefully handles any future changes to this list.  | [optional] 
**SellerLegalInfo** | Pointer to [**SellerLegalInfo**](SellerLegalInfo.md) |  | [optional] 
**UserId** | Pointer to **string** | The unique identifier of an eBay user across all eBay sites. This value does not change, even when a user changes their username. | [optional] 
**Username** | Pointer to **string** | The user name created by the seller for use on eBay. | [optional] 

## Methods

### NewSellerDetail

`func NewSellerDetail() *SellerDetail`

NewSellerDetail instantiates a new SellerDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSellerDetailWithDefaults

`func NewSellerDetailWithDefaults() *SellerDetail`

NewSellerDetailWithDefaults instantiates a new SellerDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFeedbackPercentage

`func (o *SellerDetail) GetFeedbackPercentage() string`

GetFeedbackPercentage returns the FeedbackPercentage field if non-nil, zero value otherwise.

### GetFeedbackPercentageOk

`func (o *SellerDetail) GetFeedbackPercentageOk() (*string, bool)`

GetFeedbackPercentageOk returns a tuple with the FeedbackPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeedbackPercentage

`func (o *SellerDetail) SetFeedbackPercentage(v string)`

SetFeedbackPercentage sets FeedbackPercentage field to given value.

### HasFeedbackPercentage

`func (o *SellerDetail) HasFeedbackPercentage() bool`

HasFeedbackPercentage returns a boolean if a field has been set.

### GetFeedbackScore

`func (o *SellerDetail) GetFeedbackScore() int32`

GetFeedbackScore returns the FeedbackScore field if non-nil, zero value otherwise.

### GetFeedbackScoreOk

`func (o *SellerDetail) GetFeedbackScoreOk() (*int32, bool)`

GetFeedbackScoreOk returns a tuple with the FeedbackScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeedbackScore

`func (o *SellerDetail) SetFeedbackScore(v int32)`

SetFeedbackScore sets FeedbackScore field to given value.

### HasFeedbackScore

`func (o *SellerDetail) HasFeedbackScore() bool`

HasFeedbackScore returns a boolean if a field has been set.

### GetSellerAccountType

`func (o *SellerDetail) GetSellerAccountType() string`

GetSellerAccountType returns the SellerAccountType field if non-nil, zero value otherwise.

### GetSellerAccountTypeOk

`func (o *SellerDetail) GetSellerAccountTypeOk() (*string, bool)`

GetSellerAccountTypeOk returns a tuple with the SellerAccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellerAccountType

`func (o *SellerDetail) SetSellerAccountType(v string)`

SetSellerAccountType sets SellerAccountType field to given value.

### HasSellerAccountType

`func (o *SellerDetail) HasSellerAccountType() bool`

HasSellerAccountType returns a boolean if a field has been set.

### GetSellerLegalInfo

`func (o *SellerDetail) GetSellerLegalInfo() SellerLegalInfo`

GetSellerLegalInfo returns the SellerLegalInfo field if non-nil, zero value otherwise.

### GetSellerLegalInfoOk

`func (o *SellerDetail) GetSellerLegalInfoOk() (*SellerLegalInfo, bool)`

GetSellerLegalInfoOk returns a tuple with the SellerLegalInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellerLegalInfo

`func (o *SellerDetail) SetSellerLegalInfo(v SellerLegalInfo)`

SetSellerLegalInfo sets SellerLegalInfo field to given value.

### HasSellerLegalInfo

`func (o *SellerDetail) HasSellerLegalInfo() bool`

HasSellerLegalInfo returns a boolean if a field has been set.

### GetUserId

`func (o *SellerDetail) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *SellerDetail) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *SellerDetail) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *SellerDetail) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetUsername

`func (o *SellerDetail) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *SellerDetail) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *SellerDetail) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *SellerDetail) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


