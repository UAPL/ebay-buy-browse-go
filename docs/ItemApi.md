# \ItemAPI

All URIs are relative to *https://api.ebay.com/buy/browse/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CheckCompatibility**](ItemAPI.md#CheckCompatibility) | **Post** /item/{item_id}/check_compatibility | 
[**GetItem**](ItemAPI.md#GetItem) | **Get** /item/{item_id} | 
[**GetItemByLegacyId**](ItemAPI.md#GetItemByLegacyId) | **Get** /item/get_item_by_legacy_id | 
[**GetItems**](ItemAPI.md#GetItems) | **Get** /item/ | 
[**GetItemsByItemGroup**](ItemAPI.md#GetItemsByItemGroup) | **Get** /item/get_items_by_item_group | 



## CheckCompatibility

> CompatibilityResponse CheckCompatibility(ctx, itemId).ContentType(contentType).XEBAYCMARKETPLACEID(xEBAYCMARKETPLACEID).AcceptLanguage(acceptLanguage).CompatibilityPayload(compatibilityPayload).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	itemId := "itemId_example" // string | This path parameter specifies the unique RESTful identifier of an item (such as the park you want to check).<br><br><b>RESTful Item ID Format: </b><code>v1</code>|<code><i>#</i></code>|<code><i>#</i></code><br><br>For a single SKU listing, pass in the item ID: <pre>v1|2**********2|0</pre>For a multi-SKU listing, pass in the identifier of the variation:<pre>v1|1**********2|4**********2</pre><br>For more information about item IDs for RESTful APIs, refer to <a href=\"/api-docs/buy/static/api-browse.html#Legacy\" target=\"_blank\">Item ID legacy API compatibility overview</a> in the <a href=\"/api-docs/buy/static/buying-ig-landing.html\" target=\"_blank\">Buying Integration Guide</a>.
	contentType := "contentType_example" // string | This header indicates the format of the request body provided by the client.<br><br>Its value should be set to <code>application/json</code>.<br><br>For more information, refer to <a href=\"/api-docs/static/rest-request-components.html#HTTP\" target=\"_blank \">HTTP request headers</a> in the <a href=\"/api-docs/static/ebay-rest-landing.html\" target=\"_blank\">Using eBay RESTful APIs</a> guide.
	xEBAYCMARKETPLACEID := "xEBAYCMARKETPLACEID_example" // string | This header identifies the seller's eBay marketplace. It is required for all marketplaces outside of the US.<br><br><span class=\"tablenote\"><b>Note:</b> If a marketplace ID value is not provided, the default value of <code>EBAY_US</code> is used.</span><br>See <a href=\"/api-docs/buy/browse/types/ba:MarketplaceIdEnum\" target=\"_blank\">MarketplaceIdEnum</a> for supported values. (optional)
	acceptLanguage := "acceptLanguage_example" // string | This header is used to indicate the natural language and locale preferred by the user for the response.<br><br>This header is required when targeting a specific locale of a marketplace that supports multiple locales. For example:<ul><li>When targeting the French locale of the Belgium marketplace, it is required to pass in <code>fr-BE</code> to specify this. If this locale is not specified, the language will default to Dutch.</li><li>When targeting the French locale of the Canadian marketplace, it is required to pass in <code>fr-CA</code> to specify this. If this locale is not specified, the language will default to English.</li></ul> (optional)
	compatibilityPayload := *openapiclient.NewCompatibilityPayload() // CompatibilityPayload |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ItemAPI.CheckCompatibility(context.Background(), itemId).ContentType(contentType).XEBAYCMARKETPLACEID(xEBAYCMARKETPLACEID).AcceptLanguage(acceptLanguage).CompatibilityPayload(compatibilityPayload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ItemAPI.CheckCompatibility``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CheckCompatibility`: CompatibilityResponse
	fmt.Fprintf(os.Stdout, "Response from `ItemAPI.CheckCompatibility`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | **string** | This path parameter specifies the unique RESTful identifier of an item (such as the park you want to check).&lt;br&gt;&lt;br&gt;&lt;b&gt;RESTful Item ID Format: &lt;/b&gt;&lt;code&gt;v1&lt;/code&gt;|&lt;code&gt;&lt;i&gt;#&lt;/i&gt;&lt;/code&gt;|&lt;code&gt;&lt;i&gt;#&lt;/i&gt;&lt;/code&gt;&lt;br&gt;&lt;br&gt;For a single SKU listing, pass in the item ID: &lt;pre&gt;v1|2**********2|0&lt;/pre&gt;For a multi-SKU listing, pass in the identifier of the variation:&lt;pre&gt;v1|1**********2|4**********2&lt;/pre&gt;&lt;br&gt;For more information about item IDs for RESTful APIs, refer to &lt;a href&#x3D;\&quot;/api-docs/buy/static/api-browse.html#Legacy\&quot; target&#x3D;\&quot;_blank\&quot;&gt;Item ID legacy API compatibility overview&lt;/a&gt; in the &lt;a href&#x3D;\&quot;/api-docs/buy/static/buying-ig-landing.html\&quot; target&#x3D;\&quot;_blank\&quot;&gt;Buying Integration Guide&lt;/a&gt;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCheckCompatibilityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **contentType** | **string** | This header indicates the format of the request body provided by the client.&lt;br&gt;&lt;br&gt;Its value should be set to &lt;code&gt;application/json&lt;/code&gt;.&lt;br&gt;&lt;br&gt;For more information, refer to &lt;a href&#x3D;\&quot;/api-docs/static/rest-request-components.html#HTTP\&quot; target&#x3D;\&quot;_blank \&quot;&gt;HTTP request headers&lt;/a&gt; in the &lt;a href&#x3D;\&quot;/api-docs/static/ebay-rest-landing.html\&quot; target&#x3D;\&quot;_blank\&quot;&gt;Using eBay RESTful APIs&lt;/a&gt; guide. | 
 **xEBAYCMARKETPLACEID** | **string** | This header identifies the seller&#39;s eBay marketplace. It is required for all marketplaces outside of the US.&lt;br&gt;&lt;br&gt;&lt;span class&#x3D;\&quot;tablenote\&quot;&gt;&lt;b&gt;Note:&lt;/b&gt; If a marketplace ID value is not provided, the default value of &lt;code&gt;EBAY_US&lt;/code&gt; is used.&lt;/span&gt;&lt;br&gt;See &lt;a href&#x3D;\&quot;/api-docs/buy/browse/types/ba:MarketplaceIdEnum\&quot; target&#x3D;\&quot;_blank\&quot;&gt;MarketplaceIdEnum&lt;/a&gt; for supported values. | 
 **acceptLanguage** | **string** | This header is used to indicate the natural language and locale preferred by the user for the response.&lt;br&gt;&lt;br&gt;This header is required when targeting a specific locale of a marketplace that supports multiple locales. For example:&lt;ul&gt;&lt;li&gt;When targeting the French locale of the Belgium marketplace, it is required to pass in &lt;code&gt;fr-BE&lt;/code&gt; to specify this. If this locale is not specified, the language will default to Dutch.&lt;/li&gt;&lt;li&gt;When targeting the French locale of the Canadian marketplace, it is required to pass in &lt;code&gt;fr-CA&lt;/code&gt; to specify this. If this locale is not specified, the language will default to English.&lt;/li&gt;&lt;/ul&gt; | 
 **compatibilityPayload** | [**CompatibilityPayload**](CompatibilityPayload.md) |  | 

### Return type

[**CompatibilityResponse**](CompatibilityResponse.md)

### Authorization

[api_auth](../README.md#api_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetItem

> Item GetItem(ctx, itemId).Fieldgroups(fieldgroups).XEBAYCENDUSERCTX(xEBAYCENDUSERCTX).XEBAYCMARKETPLACEID(xEBAYCMARKETPLACEID).AcceptLanguage(acceptLanguage).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	itemId := "itemId_example" // string | This path parameter specifies the unique RESTful identifier of the item being retrieved. <br><br><b>RESTful Item ID Format: </b><code>v1</code>|<code><i>#</i></code>|<code><i>#</i></code><br><br>For a single SKU listing, pass in the item ID: <pre>v1|2**********2|0</pre>For a multi-SKU listing, pass in the identifier of the variation:<pre>v1|1**********2|4**********2</pre><br>For more information about item IDs for RESTful APIs, refer to <a href=\"/api-docs/buy/static/api-browse.html#Legacy\" target=\"_blank\">Item ID legacy API compatibility overview</a> in the <a href=\"/api-docs/buy/static/buying-ig-landing.html\" target=\"_blank\">Buying Integration Guide</a>.
	fieldgroups := "fieldgroups_example" // string | This parameter controls what is returned in the response. If this field is not set, the method returns all the details of the item.<br><br><span class=\"tablenote\"><b>Note:</b> Multiple <b>fieldgroups</b> can be set and applied simultaneously. However, <b>COMPACT</b> <b>must</b> be used alone. Otherwise, an error will occur.</span><br><b>Valid Values:</b><ul><li><b>PRODUCT</b><br>This field group adds the <a href=\"/api-docs/buy/browse/resources/item/methods/getItem#response.product\">product</a> container to the response.</li><li><b>COMPACT</b></br>This field group returns only the following fields which let you quickly check if the availability or price of the item has changed, if the item has been revised by the seller, or if an item's top-rated plus status has changed for items you have stored.<ul><li><b>itemId</b></li><li><b>bidCount</b></li><li><b>currentBidPrice</b></li><li><b>eligibleForInlineCheckout</b></li><li><b>estimatedAvailabilities</b></li><li><b>gtin</b></li><li><b>itemAffiliateWebURL</b></li><li><b>itemCreationDate</b></li><li><b>ItemWebURL</b></li><li><b>legacyItemId</b></li><li><b>minimumPriceToBid</b></li><li><b>price</b></li><li><b>priorityListing</b></li><li><b>reservePriceMet</b></li><li><b>sellerItemRevision</b></li><li><b>shippingOptions</b></li><li><b>taxes</b></li><li><b>topRatedBuyingExperience</b></li><li><b>uniqueBidderCount</b></li></ul>For Example:<br>To determine if a stored item's information is current, perform the following:<ol><li>Pass in the item ID and set <code>fieldgroups</code> to <code>COMPACT</code>.<pre>item/v1|4**********8|0?fieldgroups=COMPACT</pre></li><li>Do one of the following:<ul><li>If <b>sellerItemRevision</b> is returned and a revision number for this item <i>has not</i> been stored, record the number and pass in the item ID in the <b>getItem</b> method to retrieve the most recent information.</li><li>If the revision number is different from the value you have stored, update the value and pass in the item ID in the <b>getItem</b> method to retrieve the most recent information.</li><li>If <code>sellerItemRevision</code> is <i>not</i> returned or has not changed, where needed, update the item information with the information returned in the response.</li></ul></li></ol></li><li><b>ADDITIONAL_SELLER_DETAILS</b><br>This adds the <a href=\"/api-docs/buy/browse/resources/item/methods/getItem#response.seller.userId\">userId</a> field to the response.</li></ul> (optional)
	xEBAYCENDUSERCTX := "xEBAYCENDUSERCTX_example" // string | This header is required to support revenue sharing for eBay Partner Network and to improve the accuracy of shipping and delivery time estimations.<br><br>For additional information, refer to <a href=\"/api-docs/buy/static/api-browse.html#Headers\" target=\"_blank \">Use request headers</a> in the <a href=\"/api-docs/buy/static/buying-ig-landing.html\" target=\"_blank\">Buying Integration Guide</a>. (optional)
	xEBAYCMARKETPLACEID := "xEBAYCMARKETPLACEID_example" // string | This header identifies the seller's eBay marketplace. It is required for all marketplaces outside of the US.<br><br><span class=\"tablenote\"><b>Note:</b> If a marketplace ID value is not provided, the default value of <code>EBAY_US</code> is used.</span><br>See <a href=\"/api-docs/buy/browse/types/ba:MarketplaceIdEnum\" target=\"_blank\">MarketplaceIdEnum</a> for a list of supported marketplaces. (optional)
	acceptLanguage := "acceptLanguage_example" // string | This header is used to indicate the natural language and locale preferred by the user for the response.<br><br>This header is required when targeting a specific locale of a marketplace that supports multiple locales. For example:<ul><li>When targeting the French locale of the Belgium marketplace, it is required to pass in <code>fr-BE</code> to specify this. If this locale is not specified, the language will default to Dutch.</li><li>When targeting the French locale of the Canadian marketplace, it is required to pass in <code>fr-CA</code> to specify this. If this locale is not specified, the language will default to English.</li></ul> (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ItemAPI.GetItem(context.Background(), itemId).Fieldgroups(fieldgroups).XEBAYCENDUSERCTX(xEBAYCENDUSERCTX).XEBAYCMARKETPLACEID(xEBAYCMARKETPLACEID).AcceptLanguage(acceptLanguage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ItemAPI.GetItem``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetItem`: Item
	fmt.Fprintf(os.Stdout, "Response from `ItemAPI.GetItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | **string** | This path parameter specifies the unique RESTful identifier of the item being retrieved. &lt;br&gt;&lt;br&gt;&lt;b&gt;RESTful Item ID Format: &lt;/b&gt;&lt;code&gt;v1&lt;/code&gt;|&lt;code&gt;&lt;i&gt;#&lt;/i&gt;&lt;/code&gt;|&lt;code&gt;&lt;i&gt;#&lt;/i&gt;&lt;/code&gt;&lt;br&gt;&lt;br&gt;For a single SKU listing, pass in the item ID: &lt;pre&gt;v1|2**********2|0&lt;/pre&gt;For a multi-SKU listing, pass in the identifier of the variation:&lt;pre&gt;v1|1**********2|4**********2&lt;/pre&gt;&lt;br&gt;For more information about item IDs for RESTful APIs, refer to &lt;a href&#x3D;\&quot;/api-docs/buy/static/api-browse.html#Legacy\&quot; target&#x3D;\&quot;_blank\&quot;&gt;Item ID legacy API compatibility overview&lt;/a&gt; in the &lt;a href&#x3D;\&quot;/api-docs/buy/static/buying-ig-landing.html\&quot; target&#x3D;\&quot;_blank\&quot;&gt;Buying Integration Guide&lt;/a&gt;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldgroups** | **string** | This parameter controls what is returned in the response. If this field is not set, the method returns all the details of the item.&lt;br&gt;&lt;br&gt;&lt;span class&#x3D;\&quot;tablenote\&quot;&gt;&lt;b&gt;Note:&lt;/b&gt; Multiple &lt;b&gt;fieldgroups&lt;/b&gt; can be set and applied simultaneously. However, &lt;b&gt;COMPACT&lt;/b&gt; &lt;b&gt;must&lt;/b&gt; be used alone. Otherwise, an error will occur.&lt;/span&gt;&lt;br&gt;&lt;b&gt;Valid Values:&lt;/b&gt;&lt;ul&gt;&lt;li&gt;&lt;b&gt;PRODUCT&lt;/b&gt;&lt;br&gt;This field group adds the &lt;a href&#x3D;\&quot;/api-docs/buy/browse/resources/item/methods/getItem#response.product\&quot;&gt;product&lt;/a&gt; container to the response.&lt;/li&gt;&lt;li&gt;&lt;b&gt;COMPACT&lt;/b&gt;&lt;/br&gt;This field group returns only the following fields which let you quickly check if the availability or price of the item has changed, if the item has been revised by the seller, or if an item&#39;s top-rated plus status has changed for items you have stored.&lt;ul&gt;&lt;li&gt;&lt;b&gt;itemId&lt;/b&gt;&lt;/li&gt;&lt;li&gt;&lt;b&gt;bidCount&lt;/b&gt;&lt;/li&gt;&lt;li&gt;&lt;b&gt;currentBidPrice&lt;/b&gt;&lt;/li&gt;&lt;li&gt;&lt;b&gt;eligibleForInlineCheckout&lt;/b&gt;&lt;/li&gt;&lt;li&gt;&lt;b&gt;estimatedAvailabilities&lt;/b&gt;&lt;/li&gt;&lt;li&gt;&lt;b&gt;gtin&lt;/b&gt;&lt;/li&gt;&lt;li&gt;&lt;b&gt;itemAffiliateWebURL&lt;/b&gt;&lt;/li&gt;&lt;li&gt;&lt;b&gt;itemCreationDate&lt;/b&gt;&lt;/li&gt;&lt;li&gt;&lt;b&gt;ItemWebURL&lt;/b&gt;&lt;/li&gt;&lt;li&gt;&lt;b&gt;legacyItemId&lt;/b&gt;&lt;/li&gt;&lt;li&gt;&lt;b&gt;minimumPriceToBid&lt;/b&gt;&lt;/li&gt;&lt;li&gt;&lt;b&gt;price&lt;/b&gt;&lt;/li&gt;&lt;li&gt;&lt;b&gt;priorityListing&lt;/b&gt;&lt;/li&gt;&lt;li&gt;&lt;b&gt;reservePriceMet&lt;/b&gt;&lt;/li&gt;&lt;li&gt;&lt;b&gt;sellerItemRevision&lt;/b&gt;&lt;/li&gt;&lt;li&gt;&lt;b&gt;shippingOptions&lt;/b&gt;&lt;/li&gt;&lt;li&gt;&lt;b&gt;taxes&lt;/b&gt;&lt;/li&gt;&lt;li&gt;&lt;b&gt;topRatedBuyingExperience&lt;/b&gt;&lt;/li&gt;&lt;li&gt;&lt;b&gt;uniqueBidderCount&lt;/b&gt;&lt;/li&gt;&lt;/ul&gt;For Example:&lt;br&gt;To determine if a stored item&#39;s information is current, perform the following:&lt;ol&gt;&lt;li&gt;Pass in the item ID and set &lt;code&gt;fieldgroups&lt;/code&gt; to &lt;code&gt;COMPACT&lt;/code&gt;.&lt;pre&gt;item/v1|4**********8|0?fieldgroups&#x3D;COMPACT&lt;/pre&gt;&lt;/li&gt;&lt;li&gt;Do one of the following:&lt;ul&gt;&lt;li&gt;If &lt;b&gt;sellerItemRevision&lt;/b&gt; is returned and a revision number for this item &lt;i&gt;has not&lt;/i&gt; been stored, record the number and pass in the item ID in the &lt;b&gt;getItem&lt;/b&gt; method to retrieve the most recent information.&lt;/li&gt;&lt;li&gt;If the revision number is different from the value you have stored, update the value and pass in the item ID in the &lt;b&gt;getItem&lt;/b&gt; method to retrieve the most recent information.&lt;/li&gt;&lt;li&gt;If &lt;code&gt;sellerItemRevision&lt;/code&gt; is &lt;i&gt;not&lt;/i&gt; returned or has not changed, where needed, update the item information with the information returned in the response.&lt;/li&gt;&lt;/ul&gt;&lt;/li&gt;&lt;/ol&gt;&lt;/li&gt;&lt;li&gt;&lt;b&gt;ADDITIONAL_SELLER_DETAILS&lt;/b&gt;&lt;br&gt;This adds the &lt;a href&#x3D;\&quot;/api-docs/buy/browse/resources/item/methods/getItem#response.seller.userId\&quot;&gt;userId&lt;/a&gt; field to the response.&lt;/li&gt;&lt;/ul&gt; | 
 **xEBAYCENDUSERCTX** | **string** | This header is required to support revenue sharing for eBay Partner Network and to improve the accuracy of shipping and delivery time estimations.&lt;br&gt;&lt;br&gt;For additional information, refer to &lt;a href&#x3D;\&quot;/api-docs/buy/static/api-browse.html#Headers\&quot; target&#x3D;\&quot;_blank \&quot;&gt;Use request headers&lt;/a&gt; in the &lt;a href&#x3D;\&quot;/api-docs/buy/static/buying-ig-landing.html\&quot; target&#x3D;\&quot;_blank\&quot;&gt;Buying Integration Guide&lt;/a&gt;. | 
 **xEBAYCMARKETPLACEID** | **string** | This header identifies the seller&#39;s eBay marketplace. It is required for all marketplaces outside of the US.&lt;br&gt;&lt;br&gt;&lt;span class&#x3D;\&quot;tablenote\&quot;&gt;&lt;b&gt;Note:&lt;/b&gt; If a marketplace ID value is not provided, the default value of &lt;code&gt;EBAY_US&lt;/code&gt; is used.&lt;/span&gt;&lt;br&gt;See &lt;a href&#x3D;\&quot;/api-docs/buy/browse/types/ba:MarketplaceIdEnum\&quot; target&#x3D;\&quot;_blank\&quot;&gt;MarketplaceIdEnum&lt;/a&gt; for a list of supported marketplaces. | 
 **acceptLanguage** | **string** | This header is used to indicate the natural language and locale preferred by the user for the response.&lt;br&gt;&lt;br&gt;This header is required when targeting a specific locale of a marketplace that supports multiple locales. For example:&lt;ul&gt;&lt;li&gt;When targeting the French locale of the Belgium marketplace, it is required to pass in &lt;code&gt;fr-BE&lt;/code&gt; to specify this. If this locale is not specified, the language will default to Dutch.&lt;/li&gt;&lt;li&gt;When targeting the French locale of the Canadian marketplace, it is required to pass in &lt;code&gt;fr-CA&lt;/code&gt; to specify this. If this locale is not specified, the language will default to English.&lt;/li&gt;&lt;/ul&gt; | 

### Return type

[**Item**](Item.md)

### Authorization

[api_auth](../README.md#api_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetItemByLegacyId

> Item GetItemByLegacyId(ctx).LegacyItemId(legacyItemId).Fieldgroups(fieldgroups).LegacyVariationId(legacyVariationId).LegacyVariationSku(legacyVariationSku).XEBAYCENDUSERCTX(xEBAYCENDUSERCTX).XEBAYCMARKETPLACEID(xEBAYCMARKETPLACEID).AcceptLanguage(acceptLanguage).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	legacyItemId := "legacyItemId_example" // string | This query parameter is the unique identifier that specifies the item being retrieved.<br><br><span class=\"tablenote\"><b> Note:</b> When passing in the ID for a multi-variation listing, you must also use the <code>legacy_variation_id</code> field and pass in the ID of the specific item variation. If not, all variation within the multi-SKU listing will be retrieved.</span></li></ul><br>The following is an example of using the value of the <b>ItemID</b> field for a specific item to get the RESTful <code>itemId</code> value.<pre><code>browse/v1/item/get_item_by_legacy_id?legacy_item_id=1**********9</code></pre>
	fieldgroups := "fieldgroups_example" // string | This field controls what is returned in the response. If this field is not set, the method returns all details about the item. Multiple <code>fieldgroups</code> can be set.<br><br><b>Valid Values:</b><ul><li> <b>PRODUCT</b><br>This field group adds the <a href=\"/api-docs/buy/browse/resources/item/methods/getItemByLegacyId#response.product\">product</a> container to the response. </li><li><b>ADDITIONAL_SELLER_DETAILS</b><br>This field group adds the <a href=\"/api-docs/buy/browse/resources/item/methods/getItemByLegacyId#response.seller.userId\">userId</a> field to the response.</li></ul> (optional)
	legacyVariationId := "legacyVariationId_example" // string | This query parameter specifies the legacy item ID of a specific item in a multi-variation listing, such as that for the <i>red shirt size L</i> item.<br><br><div class=\"msgbox_important\"><p class=\"msgbox_importantInDiv\" data-mc-autonum=\"&lt;b&gt;&lt;span style=&quot;color: #dd1e31;&quot; class=&quot;mcFormatColor&quot;&gt;Important! &lt;/span&gt;&lt;/b&gt;\"><span class=\"autonumber\"><span><b><span style=\"color: #dd1e31;\" class=\"mcFormatColor\">Important!</span></b></span></span> A <code>legacy_item_id</code> value must always be passed in when specifying a <code>legacy_variation_id</code> value.</p></div> (optional)
	legacyVariationSku := "legacyVariationSku_example" // string | This query parameter specifies the legacy SKU of an item. SKUs are the unique identifiers of an item created by the seller.<br><br>The following is an example of using the value of the <code>ItemID</code> and <code>SKU</code> fields to get the RESTful <code>itemId</code> value.<pre><code>browse/v1/item/get_item_by_legacy_id?legacy_item_id=1**********9&amp;legacy_variation_sku=V**********M</code></pre><div class=\"msgbox_important\"><p class=\"msgbox_importantInDiv\" data-mc-autonum=\"&lt;b&gt;&lt;span style=&quot;color: #dd1e31;&quot; class=&quot;mcFormatColor&quot;&gt;Important! &lt;/span&gt;&lt;/b&gt;\"><span class=\"autonumber\"><span><b><span style=\"color: #dd1e31;\" class=\"mcFormatColor\">Important!</span></b></span></span> A <code>legacy_item_id</code> value must always be passed in when specifying a <code>legacy_variation_sku</code> value.</p></div> (optional)
	xEBAYCENDUSERCTX := "xEBAYCENDUSERCTX_example" // string | This header is required to support revenue sharing for eBay Partner Network and to improve the accuracy of shipping and delivery time estimations.<br><br>For additional information, refer to <a href=\"/api-docs/buy/static/api-browse.html#Headers\" target=\"_blank \">Use request headers</a> in the <a href=\"/api-docs/buy/static/buying-ig-landing.html\" target=\"_blank\">Buying Integration Guide</a>. (optional)
	xEBAYCMARKETPLACEID := "xEBAYCMARKETPLACEID_example" // string | This header identifies the seller's eBay marketplace. It is required for all marketplaces outside of the US.<br><br><span class=\"tablenote\"><b>Note:</b> If a marketplace ID value is not provided, the default value of <code>EBAY_US</code> is used.</span><br>See <a href=\"/api-docs/buy/browse/types/ba:MarketplaceIdEnum\" target=\"_blank\">MarketplaceIdEnum</a> for a list of supported marketplaces. (optional)
	acceptLanguage := "acceptLanguage_example" // string | This header is used to indicate the natural language and locale preferred by the user for the response.<br><br>This header is required when targeting a specific locale of a marketplace that supports multiple locales. For example:<ul><li>When targeting the French locale of the Belgium marketplace, it is required to pass in <code>fr-BE</code> to specify this. If this locale is not specified, the language will default to Dutch.</li><li>When targeting the French locale of the Canadian marketplace, it is required to pass in <code>fr-CA</code> to specify this. If this locale is not specified, the language will default to English.</li></ul> (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ItemAPI.GetItemByLegacyId(context.Background()).LegacyItemId(legacyItemId).Fieldgroups(fieldgroups).LegacyVariationId(legacyVariationId).LegacyVariationSku(legacyVariationSku).XEBAYCENDUSERCTX(xEBAYCENDUSERCTX).XEBAYCMARKETPLACEID(xEBAYCMARKETPLACEID).AcceptLanguage(acceptLanguage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ItemAPI.GetItemByLegacyId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetItemByLegacyId`: Item
	fmt.Fprintf(os.Stdout, "Response from `ItemAPI.GetItemByLegacyId`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetItemByLegacyIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **legacyItemId** | **string** | This query parameter is the unique identifier that specifies the item being retrieved.&lt;br&gt;&lt;br&gt;&lt;span class&#x3D;\&quot;tablenote\&quot;&gt;&lt;b&gt; Note:&lt;/b&gt; When passing in the ID for a multi-variation listing, you must also use the &lt;code&gt;legacy_variation_id&lt;/code&gt; field and pass in the ID of the specific item variation. If not, all variation within the multi-SKU listing will be retrieved.&lt;/span&gt;&lt;/li&gt;&lt;/ul&gt;&lt;br&gt;The following is an example of using the value of the &lt;b&gt;ItemID&lt;/b&gt; field for a specific item to get the RESTful &lt;code&gt;itemId&lt;/code&gt; value.&lt;pre&gt;&lt;code&gt;browse/v1/item/get_item_by_legacy_id?legacy_item_id&#x3D;1**********9&lt;/code&gt;&lt;/pre&gt; | 
 **fieldgroups** | **string** | This field controls what is returned in the response. If this field is not set, the method returns all details about the item. Multiple &lt;code&gt;fieldgroups&lt;/code&gt; can be set.&lt;br&gt;&lt;br&gt;&lt;b&gt;Valid Values:&lt;/b&gt;&lt;ul&gt;&lt;li&gt; &lt;b&gt;PRODUCT&lt;/b&gt;&lt;br&gt;This field group adds the &lt;a href&#x3D;\&quot;/api-docs/buy/browse/resources/item/methods/getItemByLegacyId#response.product\&quot;&gt;product&lt;/a&gt; container to the response. &lt;/li&gt;&lt;li&gt;&lt;b&gt;ADDITIONAL_SELLER_DETAILS&lt;/b&gt;&lt;br&gt;This field group adds the &lt;a href&#x3D;\&quot;/api-docs/buy/browse/resources/item/methods/getItemByLegacyId#response.seller.userId\&quot;&gt;userId&lt;/a&gt; field to the response.&lt;/li&gt;&lt;/ul&gt; | 
 **legacyVariationId** | **string** | This query parameter specifies the legacy item ID of a specific item in a multi-variation listing, such as that for the &lt;i&gt;red shirt size L&lt;/i&gt; item.&lt;br&gt;&lt;br&gt;&lt;div class&#x3D;\&quot;msgbox_important\&quot;&gt;&lt;p class&#x3D;\&quot;msgbox_importantInDiv\&quot; data-mc-autonum&#x3D;\&quot;&amp;lt;b&amp;gt;&amp;lt;span style&#x3D;&amp;quot;color: #dd1e31;&amp;quot; class&#x3D;&amp;quot;mcFormatColor&amp;quot;&amp;gt;Important! &amp;lt;/span&amp;gt;&amp;lt;/b&amp;gt;\&quot;&gt;&lt;span class&#x3D;\&quot;autonumber\&quot;&gt;&lt;span&gt;&lt;b&gt;&lt;span style&#x3D;\&quot;color: #dd1e31;\&quot; class&#x3D;\&quot;mcFormatColor\&quot;&gt;Important!&lt;/span&gt;&lt;/b&gt;&lt;/span&gt;&lt;/span&gt; A &lt;code&gt;legacy_item_id&lt;/code&gt; value must always be passed in when specifying a &lt;code&gt;legacy_variation_id&lt;/code&gt; value.&lt;/p&gt;&lt;/div&gt; | 
 **legacyVariationSku** | **string** | This query parameter specifies the legacy SKU of an item. SKUs are the unique identifiers of an item created by the seller.&lt;br&gt;&lt;br&gt;The following is an example of using the value of the &lt;code&gt;ItemID&lt;/code&gt; and &lt;code&gt;SKU&lt;/code&gt; fields to get the RESTful &lt;code&gt;itemId&lt;/code&gt; value.&lt;pre&gt;&lt;code&gt;browse/v1/item/get_item_by_legacy_id?legacy_item_id&#x3D;1**********9&amp;amp;legacy_variation_sku&#x3D;V**********M&lt;/code&gt;&lt;/pre&gt;&lt;div class&#x3D;\&quot;msgbox_important\&quot;&gt;&lt;p class&#x3D;\&quot;msgbox_importantInDiv\&quot; data-mc-autonum&#x3D;\&quot;&amp;lt;b&amp;gt;&amp;lt;span style&#x3D;&amp;quot;color: #dd1e31;&amp;quot; class&#x3D;&amp;quot;mcFormatColor&amp;quot;&amp;gt;Important! &amp;lt;/span&amp;gt;&amp;lt;/b&amp;gt;\&quot;&gt;&lt;span class&#x3D;\&quot;autonumber\&quot;&gt;&lt;span&gt;&lt;b&gt;&lt;span style&#x3D;\&quot;color: #dd1e31;\&quot; class&#x3D;\&quot;mcFormatColor\&quot;&gt;Important!&lt;/span&gt;&lt;/b&gt;&lt;/span&gt;&lt;/span&gt; A &lt;code&gt;legacy_item_id&lt;/code&gt; value must always be passed in when specifying a &lt;code&gt;legacy_variation_sku&lt;/code&gt; value.&lt;/p&gt;&lt;/div&gt; | 
 **xEBAYCENDUSERCTX** | **string** | This header is required to support revenue sharing for eBay Partner Network and to improve the accuracy of shipping and delivery time estimations.&lt;br&gt;&lt;br&gt;For additional information, refer to &lt;a href&#x3D;\&quot;/api-docs/buy/static/api-browse.html#Headers\&quot; target&#x3D;\&quot;_blank \&quot;&gt;Use request headers&lt;/a&gt; in the &lt;a href&#x3D;\&quot;/api-docs/buy/static/buying-ig-landing.html\&quot; target&#x3D;\&quot;_blank\&quot;&gt;Buying Integration Guide&lt;/a&gt;. | 
 **xEBAYCMARKETPLACEID** | **string** | This header identifies the seller&#39;s eBay marketplace. It is required for all marketplaces outside of the US.&lt;br&gt;&lt;br&gt;&lt;span class&#x3D;\&quot;tablenote\&quot;&gt;&lt;b&gt;Note:&lt;/b&gt; If a marketplace ID value is not provided, the default value of &lt;code&gt;EBAY_US&lt;/code&gt; is used.&lt;/span&gt;&lt;br&gt;See &lt;a href&#x3D;\&quot;/api-docs/buy/browse/types/ba:MarketplaceIdEnum\&quot; target&#x3D;\&quot;_blank\&quot;&gt;MarketplaceIdEnum&lt;/a&gt; for a list of supported marketplaces. | 
 **acceptLanguage** | **string** | This header is used to indicate the natural language and locale preferred by the user for the response.&lt;br&gt;&lt;br&gt;This header is required when targeting a specific locale of a marketplace that supports multiple locales. For example:&lt;ul&gt;&lt;li&gt;When targeting the French locale of the Belgium marketplace, it is required to pass in &lt;code&gt;fr-BE&lt;/code&gt; to specify this. If this locale is not specified, the language will default to Dutch.&lt;/li&gt;&lt;li&gt;When targeting the French locale of the Canadian marketplace, it is required to pass in &lt;code&gt;fr-CA&lt;/code&gt; to specify this. If this locale is not specified, the language will default to English.&lt;/li&gt;&lt;/ul&gt; | 

### Return type

[**Item**](Item.md)

### Authorization

[api_auth](../README.md#api_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetItems

> Items GetItems(ctx).ItemIds(itemIds).ItemGroupIds(itemGroupIds).XEBAYCENDUSERCTX(xEBAYCENDUSERCTX).XEBAYCMARKETPLACEID(xEBAYCMARKETPLACEID).AcceptLanguage(acceptLanguage).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	itemIds := "itemIds_example" // string | A comma separated list of the unique identifiers of the items to retrieve (maximum 20).<br><br><span class=\"tablenote\"><b>Note:</b> In any given request, either <code>item_ids</code> <b>or</b> <code>item_group_ids</code> can be retrieved. Attempting to retrieve both will result in an error.</span><br><b>RESTful Item ID Format: </b><code>v1</code>|<code><i>#</i></code>|<code><i>#</i></code><br><br>For a single SKU listing, pass in the item ID: <pre>v1|2**********2|0</pre>For a multi-SKU listing, pass in the identifier of the variation:<pre>v1|1**********2|4**********2</pre><br>For more information about item IDs for RESTful APIs, refer to <a href=\"/api-docs/buy/static/api-browse.html#Legacy\" target=\"_blank\">Item ID legacy API compatibility overview</a> in the <a href=\"/api-docs/buy/static/buying-ig-landing.html\" target=\"_blank\">Buying Integration Guide</a>. (optional)
	itemGroupIds := "itemGroupIds_example" // string | A comma separated list of the unique identifiers of the item groups being retrieved (maximum 10).<br><br><span class=\"tablenote\"><b>Note:</b> In any given request, either <code>item_ids</code> <b>or</b> <code>item_group_ids</code> can be retrieved. Attempting to retrieve both will result in an error.</span><br><b>RESTful Group Item ID Format:</b> <code>############</code><br><br>For example:<pre>3**********9</pre> (optional)
	xEBAYCENDUSERCTX := "xEBAYCENDUSERCTX_example" // string | This header is required to support revenue sharing for eBay Partner Network and to improve the accuracy of shipping and delivery time estimations.<br><br>For additional information, refer to <a href=\"/api-docs/buy/static/api-browse.html#Headers\" target=\"_blank \">Use request headers</a> in the <a href=\"/api-docs/buy/static/buying-ig-landing.html\" target=\"_blank\">Buying Integration Guide</a>. (optional)
	xEBAYCMARKETPLACEID := "xEBAYCMARKETPLACEID_example" // string | This header identifies the seller's eBay marketplace. It is required for all marketplaces outside of the US.<br><br><span class=\"tablenote\"><b>Note:</b> If a marketplace ID value is not provided, the default value of <code>EBAY_US</code> is used.</span><br>See <a href=\"/api-docs/buy/browse/types/ba:MarketplaceIdEnum\" target=\"_blank\">MarketplaceIdEnum</a> for a list of supported marketplaces. (optional)
	acceptLanguage := "acceptLanguage_example" // string | This header is used to indicate the natural language and locale preferred by the user for the response.<br><br>This header is required when targeting a specific locale of a marketplace that supports multiple locales. For example:<ul><li>When targeting the French locale of the Belgium marketplace, it is required to pass in <code>fr-BE</code> to specify this. If this locale is not specified, the language will default to Dutch.</li><li>When targeting the French locale of the Canadian marketplace, it is required to pass in <code>fr-CA</code> to specify this. If this locale is not specified, the language will default to English.</li></ul> (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ItemAPI.GetItems(context.Background()).ItemIds(itemIds).ItemGroupIds(itemGroupIds).XEBAYCENDUSERCTX(xEBAYCENDUSERCTX).XEBAYCMARKETPLACEID(xEBAYCMARKETPLACEID).AcceptLanguage(acceptLanguage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ItemAPI.GetItems``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetItems`: Items
	fmt.Fprintf(os.Stdout, "Response from `ItemAPI.GetItems`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **itemIds** | **string** | A comma separated list of the unique identifiers of the items to retrieve (maximum 20).&lt;br&gt;&lt;br&gt;&lt;span class&#x3D;\&quot;tablenote\&quot;&gt;&lt;b&gt;Note:&lt;/b&gt; In any given request, either &lt;code&gt;item_ids&lt;/code&gt; &lt;b&gt;or&lt;/b&gt; &lt;code&gt;item_group_ids&lt;/code&gt; can be retrieved. Attempting to retrieve both will result in an error.&lt;/span&gt;&lt;br&gt;&lt;b&gt;RESTful Item ID Format: &lt;/b&gt;&lt;code&gt;v1&lt;/code&gt;|&lt;code&gt;&lt;i&gt;#&lt;/i&gt;&lt;/code&gt;|&lt;code&gt;&lt;i&gt;#&lt;/i&gt;&lt;/code&gt;&lt;br&gt;&lt;br&gt;For a single SKU listing, pass in the item ID: &lt;pre&gt;v1|2**********2|0&lt;/pre&gt;For a multi-SKU listing, pass in the identifier of the variation:&lt;pre&gt;v1|1**********2|4**********2&lt;/pre&gt;&lt;br&gt;For more information about item IDs for RESTful APIs, refer to &lt;a href&#x3D;\&quot;/api-docs/buy/static/api-browse.html#Legacy\&quot; target&#x3D;\&quot;_blank\&quot;&gt;Item ID legacy API compatibility overview&lt;/a&gt; in the &lt;a href&#x3D;\&quot;/api-docs/buy/static/buying-ig-landing.html\&quot; target&#x3D;\&quot;_blank\&quot;&gt;Buying Integration Guide&lt;/a&gt;. | 
 **itemGroupIds** | **string** | A comma separated list of the unique identifiers of the item groups being retrieved (maximum 10).&lt;br&gt;&lt;br&gt;&lt;span class&#x3D;\&quot;tablenote\&quot;&gt;&lt;b&gt;Note:&lt;/b&gt; In any given request, either &lt;code&gt;item_ids&lt;/code&gt; &lt;b&gt;or&lt;/b&gt; &lt;code&gt;item_group_ids&lt;/code&gt; can be retrieved. Attempting to retrieve both will result in an error.&lt;/span&gt;&lt;br&gt;&lt;b&gt;RESTful Group Item ID Format:&lt;/b&gt; &lt;code&gt;############&lt;/code&gt;&lt;br&gt;&lt;br&gt;For example:&lt;pre&gt;3**********9&lt;/pre&gt; | 
 **xEBAYCENDUSERCTX** | **string** | This header is required to support revenue sharing for eBay Partner Network and to improve the accuracy of shipping and delivery time estimations.&lt;br&gt;&lt;br&gt;For additional information, refer to &lt;a href&#x3D;\&quot;/api-docs/buy/static/api-browse.html#Headers\&quot; target&#x3D;\&quot;_blank \&quot;&gt;Use request headers&lt;/a&gt; in the &lt;a href&#x3D;\&quot;/api-docs/buy/static/buying-ig-landing.html\&quot; target&#x3D;\&quot;_blank\&quot;&gt;Buying Integration Guide&lt;/a&gt;. | 
 **xEBAYCMARKETPLACEID** | **string** | This header identifies the seller&#39;s eBay marketplace. It is required for all marketplaces outside of the US.&lt;br&gt;&lt;br&gt;&lt;span class&#x3D;\&quot;tablenote\&quot;&gt;&lt;b&gt;Note:&lt;/b&gt; If a marketplace ID value is not provided, the default value of &lt;code&gt;EBAY_US&lt;/code&gt; is used.&lt;/span&gt;&lt;br&gt;See &lt;a href&#x3D;\&quot;/api-docs/buy/browse/types/ba:MarketplaceIdEnum\&quot; target&#x3D;\&quot;_blank\&quot;&gt;MarketplaceIdEnum&lt;/a&gt; for a list of supported marketplaces. | 
 **acceptLanguage** | **string** | This header is used to indicate the natural language and locale preferred by the user for the response.&lt;br&gt;&lt;br&gt;This header is required when targeting a specific locale of a marketplace that supports multiple locales. For example:&lt;ul&gt;&lt;li&gt;When targeting the French locale of the Belgium marketplace, it is required to pass in &lt;code&gt;fr-BE&lt;/code&gt; to specify this. If this locale is not specified, the language will default to Dutch.&lt;/li&gt;&lt;li&gt;When targeting the French locale of the Canadian marketplace, it is required to pass in &lt;code&gt;fr-CA&lt;/code&gt; to specify this. If this locale is not specified, the language will default to English.&lt;/li&gt;&lt;/ul&gt; | 

### Return type

[**Items**](Items.md)

### Authorization

[api_auth](../README.md#api_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetItemsByItemGroup

> ItemGroup GetItemsByItemGroup(ctx).ItemGroupId(itemGroupId).Fieldgroups(fieldgroups).XEBAYCENDUSERCTX(xEBAYCENDUSERCTX).XEBAYCMARKETPLACEID(xEBAYCMARKETPLACEID).AcceptLanguage(acceptLanguage).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	itemGroupId := "itemGroupId_example" // string | This query parameter specifies the unique identifier of an item group for which information is to be returned. An item group is an item that has various aspect differences, such as color, size, storage capacity, etc.<br><br>This ID is returned in the <b>itemGroupHref</b> field of the <a href=\"/api-docs/buy/browse/resources/item_summary/methods/search\">search</a> and <a href=\"/api-docs/buy/browse/resources/item/methods/getItem\">getItem</a> methods.<br><br><b>For Example:</b><pre>https://api.ebay.com/buy/browse/v1/item/get_items_by_item_group?item_group_id=3**********6</pre>
	fieldgroups := "fieldgroups_example" // string | This field controls what is returned in the response. If this field is not set, the method returns all details about the item.<br><br><b>Valid Values:</b><br><br><b>ADDITIONAL_SELLER_DETAILS</b> - This field group adds the <a href=\"/api-docs/buy/browse/resources/item/methods/getItemsByItemGroup#response.items.seller.userId\">userId</a> field to the response. (optional)
	xEBAYCENDUSERCTX := "xEBAYCENDUSERCTX_example" // string | This header is required to support revenue sharing for eBay Partner Network and to improve the accuracy of shipping and delivery time estimations.<br><br>For additional information, refer to <a href=\"/api-docs/buy/static/api-browse.html#Headers\" target=\"_blank \">Use request headers</a> in the <a href=\"/api-docs/buy/static/buying-ig-landing.html\" target=\"_blank\">Buying Integration Guide</a>. (optional)
	xEBAYCMARKETPLACEID := "xEBAYCMARKETPLACEID_example" // string | This header identifies the seller's eBay marketplace. It is required for all marketplaces outside of the US.<br><br><span class=\"tablenote\"><b>Note:</b> If a marketplace ID value is not provided, the default value of <code>EBAY_US</code> is used.</span><br>See <a href=\"/api-docs/buy/browse/types/ba:MarketplaceIdEnum\" target=\"_blank\">MarketplaceIdEnum</a> for a list of supported marketplaces. (optional)
	acceptLanguage := "acceptLanguage_example" // string | This header is used to indicate the natural language and locale preferred by the user for the response.<br><br>This header is required when targeting a specific locale of a marketplace that supports multiple locales. For example:<ul><li>When targeting the French locale of the Belgium marketplace, it is required to pass in <code>fr-BE</code> to specify this. If this locale is not specified, the language will default to Dutch.</li><li>When targeting the French locale of the Canadian marketplace, it is required to pass in <code>fr-CA</code> to specify this. If this locale is not specified, the language will default to English.</li></ul> (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ItemAPI.GetItemsByItemGroup(context.Background()).ItemGroupId(itemGroupId).Fieldgroups(fieldgroups).XEBAYCENDUSERCTX(xEBAYCENDUSERCTX).XEBAYCMARKETPLACEID(xEBAYCMARKETPLACEID).AcceptLanguage(acceptLanguage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ItemAPI.GetItemsByItemGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetItemsByItemGroup`: ItemGroup
	fmt.Fprintf(os.Stdout, "Response from `ItemAPI.GetItemsByItemGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetItemsByItemGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **itemGroupId** | **string** | This query parameter specifies the unique identifier of an item group for which information is to be returned. An item group is an item that has various aspect differences, such as color, size, storage capacity, etc.&lt;br&gt;&lt;br&gt;This ID is returned in the &lt;b&gt;itemGroupHref&lt;/b&gt; field of the &lt;a href&#x3D;\&quot;/api-docs/buy/browse/resources/item_summary/methods/search\&quot;&gt;search&lt;/a&gt; and &lt;a href&#x3D;\&quot;/api-docs/buy/browse/resources/item/methods/getItem\&quot;&gt;getItem&lt;/a&gt; methods.&lt;br&gt;&lt;br&gt;&lt;b&gt;For Example:&lt;/b&gt;&lt;pre&gt;https://api.ebay.com/buy/browse/v1/item/get_items_by_item_group?item_group_id&#x3D;3**********6&lt;/pre&gt; | 
 **fieldgroups** | **string** | This field controls what is returned in the response. If this field is not set, the method returns all details about the item.&lt;br&gt;&lt;br&gt;&lt;b&gt;Valid Values:&lt;/b&gt;&lt;br&gt;&lt;br&gt;&lt;b&gt;ADDITIONAL_SELLER_DETAILS&lt;/b&gt; - This field group adds the &lt;a href&#x3D;\&quot;/api-docs/buy/browse/resources/item/methods/getItemsByItemGroup#response.items.seller.userId\&quot;&gt;userId&lt;/a&gt; field to the response. | 
 **xEBAYCENDUSERCTX** | **string** | This header is required to support revenue sharing for eBay Partner Network and to improve the accuracy of shipping and delivery time estimations.&lt;br&gt;&lt;br&gt;For additional information, refer to &lt;a href&#x3D;\&quot;/api-docs/buy/static/api-browse.html#Headers\&quot; target&#x3D;\&quot;_blank \&quot;&gt;Use request headers&lt;/a&gt; in the &lt;a href&#x3D;\&quot;/api-docs/buy/static/buying-ig-landing.html\&quot; target&#x3D;\&quot;_blank\&quot;&gt;Buying Integration Guide&lt;/a&gt;. | 
 **xEBAYCMARKETPLACEID** | **string** | This header identifies the seller&#39;s eBay marketplace. It is required for all marketplaces outside of the US.&lt;br&gt;&lt;br&gt;&lt;span class&#x3D;\&quot;tablenote\&quot;&gt;&lt;b&gt;Note:&lt;/b&gt; If a marketplace ID value is not provided, the default value of &lt;code&gt;EBAY_US&lt;/code&gt; is used.&lt;/span&gt;&lt;br&gt;See &lt;a href&#x3D;\&quot;/api-docs/buy/browse/types/ba:MarketplaceIdEnum\&quot; target&#x3D;\&quot;_blank\&quot;&gt;MarketplaceIdEnum&lt;/a&gt; for a list of supported marketplaces. | 
 **acceptLanguage** | **string** | This header is used to indicate the natural language and locale preferred by the user for the response.&lt;br&gt;&lt;br&gt;This header is required when targeting a specific locale of a marketplace that supports multiple locales. For example:&lt;ul&gt;&lt;li&gt;When targeting the French locale of the Belgium marketplace, it is required to pass in &lt;code&gt;fr-BE&lt;/code&gt; to specify this. If this locale is not specified, the language will default to Dutch.&lt;/li&gt;&lt;li&gt;When targeting the French locale of the Canadian marketplace, it is required to pass in &lt;code&gt;fr-CA&lt;/code&gt; to specify this. If this locale is not specified, the language will default to English.&lt;/li&gt;&lt;/ul&gt; | 

### Return type

[**ItemGroup**](ItemGroup.md)

### Authorization

[api_auth](../README.md#api_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

