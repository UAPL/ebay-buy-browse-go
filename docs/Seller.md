# Seller

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FeedbackPercentage** | Pointer to **string** | The percentage of the total positive feedback. | [optional] 
**FeedbackScore** | Pointer to **int32** | The feedback score of the seller. This value is based on the ratings from eBay members that bought items from this seller. | [optional] 
**SellerAccountType** | Pointer to **string** | Indicates if the seller is a business or an individual. This is determined when the seller registers with eBay. If they register for a business account, this value will be BUSINESS. If they register for a private account, this value will be INDIVIDUAL. This designation is required by the tax laws in some countries.   &lt;br /&gt;&lt;br /&gt;This field is returned only on the following sites. &lt;br /&gt;&lt;br /&gt;EBAY_AT, EBAY_BE, EBAY_CH, EBAY_DE, EBAY_ES, EBAY_FR, EBAY_GB, EBAY_IE, EBAY_IT, EBAY_PL &lt;br /&gt;&lt;br /&gt;&lt;b&gt; Valid Values:&lt;/b&gt; BUSINESS or INDIVIDUAL &lt;br /&gt;&lt;br /&gt;Code so that your app gracefully handles any future changes to this list.  | [optional] 
**Username** | Pointer to **string** | The user name created by the seller for use on eBay. | [optional] 

## Methods

### NewSeller

`func NewSeller() *Seller`

NewSeller instantiates a new Seller object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSellerWithDefaults

`func NewSellerWithDefaults() *Seller`

NewSellerWithDefaults instantiates a new Seller object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFeedbackPercentage

`func (o *Seller) GetFeedbackPercentage() string`

GetFeedbackPercentage returns the FeedbackPercentage field if non-nil, zero value otherwise.

### GetFeedbackPercentageOk

`func (o *Seller) GetFeedbackPercentageOk() (*string, bool)`

GetFeedbackPercentageOk returns a tuple with the FeedbackPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeedbackPercentage

`func (o *Seller) SetFeedbackPercentage(v string)`

SetFeedbackPercentage sets FeedbackPercentage field to given value.

### HasFeedbackPercentage

`func (o *Seller) HasFeedbackPercentage() bool`

HasFeedbackPercentage returns a boolean if a field has been set.

### GetFeedbackScore

`func (o *Seller) GetFeedbackScore() int32`

GetFeedbackScore returns the FeedbackScore field if non-nil, zero value otherwise.

### GetFeedbackScoreOk

`func (o *Seller) GetFeedbackScoreOk() (*int32, bool)`

GetFeedbackScoreOk returns a tuple with the FeedbackScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeedbackScore

`func (o *Seller) SetFeedbackScore(v int32)`

SetFeedbackScore sets FeedbackScore field to given value.

### HasFeedbackScore

`func (o *Seller) HasFeedbackScore() bool`

HasFeedbackScore returns a boolean if a field has been set.

### GetSellerAccountType

`func (o *Seller) GetSellerAccountType() string`

GetSellerAccountType returns the SellerAccountType field if non-nil, zero value otherwise.

### GetSellerAccountTypeOk

`func (o *Seller) GetSellerAccountTypeOk() (*string, bool)`

GetSellerAccountTypeOk returns a tuple with the SellerAccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellerAccountType

`func (o *Seller) SetSellerAccountType(v string)`

SetSellerAccountType sets SellerAccountType field to given value.

### HasSellerAccountType

`func (o *Seller) HasSellerAccountType() bool`

HasSellerAccountType returns a boolean if a field has been set.

### GetUsername

`func (o *Seller) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *Seller) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *Seller) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *Seller) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


