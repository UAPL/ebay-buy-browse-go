# BuyingOptionDistribution

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BuyingOption** | Pointer to **string** | The container that returns the buying option type. This will be AUCTION, FIXED_PRICE, CLASSIFIED_AD, or a combination of these options. For details, see &lt;a href&#x3D;\&quot;/api-docs/buy/browse/resources/item_summary/methods/search#response.itemSummaries.buyingOptions\&quot;&gt;buyingOptions&lt;/a&gt;. | [optional] 
**MatchCount** | Pointer to **int32** | The number of items having this buying option. | [optional] 
**RefinementHref** | Pointer to **string** | The HATEOAS reference for this buying option. | [optional] 

## Methods

### NewBuyingOptionDistribution

`func NewBuyingOptionDistribution() *BuyingOptionDistribution`

NewBuyingOptionDistribution instantiates a new BuyingOptionDistribution object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBuyingOptionDistributionWithDefaults

`func NewBuyingOptionDistributionWithDefaults() *BuyingOptionDistribution`

NewBuyingOptionDistributionWithDefaults instantiates a new BuyingOptionDistribution object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuyingOption

`func (o *BuyingOptionDistribution) GetBuyingOption() string`

GetBuyingOption returns the BuyingOption field if non-nil, zero value otherwise.

### GetBuyingOptionOk

`func (o *BuyingOptionDistribution) GetBuyingOptionOk() (*string, bool)`

GetBuyingOptionOk returns a tuple with the BuyingOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyingOption

`func (o *BuyingOptionDistribution) SetBuyingOption(v string)`

SetBuyingOption sets BuyingOption field to given value.

### HasBuyingOption

`func (o *BuyingOptionDistribution) HasBuyingOption() bool`

HasBuyingOption returns a boolean if a field has been set.

### GetMatchCount

`func (o *BuyingOptionDistribution) GetMatchCount() int32`

GetMatchCount returns the MatchCount field if non-nil, zero value otherwise.

### GetMatchCountOk

`func (o *BuyingOptionDistribution) GetMatchCountOk() (*int32, bool)`

GetMatchCountOk returns a tuple with the MatchCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchCount

`func (o *BuyingOptionDistribution) SetMatchCount(v int32)`

SetMatchCount sets MatchCount field to given value.

### HasMatchCount

`func (o *BuyingOptionDistribution) HasMatchCount() bool`

HasMatchCount returns a boolean if a field has been set.

### GetRefinementHref

`func (o *BuyingOptionDistribution) GetRefinementHref() string`

GetRefinementHref returns the RefinementHref field if non-nil, zero value otherwise.

### GetRefinementHrefOk

`func (o *BuyingOptionDistribution) GetRefinementHrefOk() (*string, bool)`

GetRefinementHrefOk returns a tuple with the RefinementHref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefinementHref

`func (o *BuyingOptionDistribution) SetRefinementHref(v string)`

SetRefinementHref sets RefinementHref field to given value.

### HasRefinementHref

`func (o *BuyingOptionDistribution) HasRefinementHref() bool`

HasRefinementHref returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


