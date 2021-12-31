/*
Browse API

<p>The Browse API has the following resources:</p>   <ul> <li><b> item_summary: </b> Lets shoppers search for specific items by keyword, GTIN, category, charity, product, or item aspects and refine the results by using filters, such as aspects, compatibility, and fields values.</li>  <li><b> search_by_image: </b><a href=\"https://developer.ebay.com/api-docs/static/versioning.html#experimental\" target=\"_blank\"><img src=\"/cms/img/docs/experimental-icon.svg\" class=\"legend-icon experimental-icon\" alt=\"Experimental Release\" title=\"Experimental Release\" />&nbsp;(Experimental)</a> Lets shoppers search for specific items by image. You can refine the results by using URI parameters and filters.</li>   <li><b> item: </b> <ul><li>Lets you retrieve the details of a specific item or all the items in an item group, which is an item with variations such as color and size and check if a product is compatible with the specified item, such as if a specific car is compatible with a specific part.</li> <li>Provides a bridge between the eBay legacy APIs, such as <b> Finding</b>, and the RESTful APIs, which use different formats for the item IDs.</li>  </ul> </li>  <li> <b> shopping_cart: </b> <a href=\"https://developer.ebay.com/api-docs/static/versioning.html#experimental\" target=\"_blank\"><img src=\"/cms/img/docs/experimental-icon.svg\" class=\"legend-icon experimental-icon\" alt=\"Experimental Release\" title=\"Experimental Release\" />&nbsp;(Experimental)</a> <a href=\"https://developer.ebay.com/api-docs/static/versioning.html#limited\" target=\"_blank\"> <img src=\"/cms/img/docs/partners-api.svg\" class=\"legend-icon partners-icon\" title=\"Limited Release\"  alt=\"Limited Release\" />(Limited Release)</a> Provides the ability for eBay members to see the contents of their eBay cart, and add, remove, and change the quantity of items in their eBay cart.&nbsp;&nbsp;<b> Note: </b> This resource is not available in the eBay API Explorer.</li></ul>       <p>The <b> item_summary</b>, <b> search_by_image</b>, and <b> item</b> resource calls require an <a href=\"/api-docs/static/oauth-client-credentials-grant.html\">Application access token</a>. The <b> shopping_cart</b> resource calls require a <a href=\"/api-docs/static/oauth-authorization-code-grant.html\">User access token</a>.</p>

API version: v1.11.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package buybrowse

import (
	"bytes"
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
)

// Linger please
var (
	_ _context.Context
)

// ItemSummaryApiService ItemSummaryApi service
type ItemSummaryApiService service

type ApiSearchRequest struct {
	ctx _context.Context
	ApiService *ItemSummaryApiService
	aspectFilter *string
	autoCorrect *string
	categoryIds *string
	charityIds *string
	compatibilityFilter *string
	epid *string
	fieldgroups *string
	filter *string
	gtin *string
	limit *string
	offset *string
	q *string
	sort *string
}

// This field lets you filter by item aspects. The aspect name/value pairs and category, which is required, is used to limit the results to specific aspects of the item. For example, in a clothing category one aspect pair would be Color/Red. &lt;br /&gt;&lt;br /&gt;For example, the method below uses the category ID for Women&#39;s Clothing. This will return only items for a woman&#39;s red shirt.&lt;br /&gt;&lt;br /&gt;&lt;code&gt;/buy/browse/v1/item_summary/search?q&#x3D;shirt&amp;category_ids&#x3D;15724&amp;aspect_filter&#x3D;categoryId:15724,Color:{Red}&lt;/code&gt; &lt;br /&gt;&lt;br /&gt;To get a list of the aspects pairs and the category, which is returned in the &lt;b&gt; dominantCategoryId&lt;/b&gt; field, set &lt;b&gt; fieldgroups&lt;/b&gt; to &lt;code&gt;ASPECT_REFINEMENTS&lt;/code&gt;.   &lt;br /&gt;&lt;br /&gt; &lt;code&gt;/buy/browse/v1/item_summary/search?q&#x3D;shirt&amp;amp;fieldgroups&#x3D;ASPECT_REFINEMENTS&lt;/code&gt; &lt;br /&gt;&lt;br /&gt;&lt;b&gt;Required: &lt;/b&gt; The category ID is required &lt;i&gt;twice&lt;/i&gt;; once as a URI parameter and as part of the &lt;b&gt; aspect_filter&lt;/b&gt;. For implementation help, refer to eBay API documentation at https://developer.ebay.com/api-docs/buy/browse/types/gct:AspectFilter
func (r ApiSearchRequest) AspectFilter(aspectFilter string) ApiSearchRequest {
	r.aspectFilter = &aspectFilter
	return r
}
// A query parameter that enables auto correction.&lt;br /&gt;&lt;br /&gt;&lt;b&gt;Valid Values:&lt;/b&gt; &lt;code&gt;KEYWORD&lt;/code&gt;
func (r ApiSearchRequest) AutoCorrect(autoCorrect string) ApiSearchRequest {
	r.autoCorrect = &autoCorrect
	return r
}
// &lt;a name&#x3D;\&quot;category_ids\&quot;&gt;&lt;/a&gt;The category ID is used to limit the results. This field can have one category ID or a comma separated list of IDs.&lt;br /&gt;&lt;br /&gt;&lt;b&gt; For example: &lt;/b&gt;&lt;br/&gt;&lt;code&gt;/buy/browse/v1/item_summary/search?category_ids&#x3D;29792&lt;/code&gt; &lt;br /&gt;&lt;br /&gt;&lt;span class&#x3D;\&quot;tablenote\&quot;&gt;&lt;b&gt; Note: &lt;/b&gt;Currently, you can pass in only one category ID per request.&lt;/span&gt; &lt;br /&gt; &lt;br /&gt;You can also use any combination of the &lt;b&gt; category_Ids&lt;/b&gt;, &lt;b&gt; epid&lt;/b&gt;, and &lt;b&gt; q&lt;/b&gt; fields. This gives you additional control over the result set. &lt;br /&gt;&lt;br /&gt;For example, let&#39;s say you are looking of a toy phone. If you search for \&quot;phone\&quot;, the result set will be mobile phones because this is the \&quot;Best Match\&quot; for this search. But if you also include the toy category ID, the results will be what you wanted. &lt;br /&gt;&lt;br /&gt;&lt;b&gt; For example: &lt;/b&gt;&lt;br /&gt;&lt;code&gt;/buy/browse/v1/item_summary/search?q&#x3D;phone&amp;category_ids&#x3D;220&lt;/code&gt;&lt;br /&gt; &lt;br /&gt;The list of eBay category IDs is not published and category IDs are not the same across all the eBay marketplaces. You can use the following techniques to find a category by site: &lt;ul&gt; &lt;li&gt;Use the &lt;a href&#x3D;\&quot;https://pages.ebay.com/sellerinformation/news/categorychanges.html\&quot; target&#x3D;\&quot;_blank\&quot;&gt;Category Changes page&lt;/a&gt;.&lt;/li&gt; &lt;li&gt;Use the Taxonomy API. For details see &lt;a href&#x3D;\&quot;/api-docs/buy/buy-categories.html\&quot;&gt;Get Categories for Buy APIs&lt;/a&gt;. &lt;/li&gt;  &lt;li&gt;Submit the following method to get the &lt;b&gt; dominantCategoryId&lt;/b&gt; for an item. &lt;br /&gt;&lt;br /&gt;&lt;code&gt;/buy/browse/v1/item_summary/search?q&#x3D;&lt;em&gt; keyword&lt;/em&gt;&amp;fieldgroups&#x3D;ASPECT_REFINEMENTS  &lt;/code&gt;&lt;/li&gt;&lt;/ul&gt;  &lt;span class&#x3D;\&quot;tablenote\&quot;&gt; &lt;b&gt; Note:&lt;/b&gt; If a top-level (L1) category is specified, you &lt;b&gt; must&lt;/b&gt; also include the &lt;b&gt; q&lt;/b&gt; query parameter.&lt;/span&gt; &lt;br /&gt;&lt;br /&gt;&lt;b&gt; Required: &lt;/b&gt; The method must have &lt;b&gt; category_ids&lt;/b&gt;, &lt;b&gt; epid&lt;/b&gt;, &lt;b&gt; gtin&lt;/b&gt;, or &lt;b&gt; q&lt;/b&gt;  (or any combination of these)
func (r ApiSearchRequest) CategoryIds(categoryIds string) ApiSearchRequest {
	r.categoryIds = &categoryIds
	return r
}
// The charity ID is used to limit the results to only items associated with the specified charity. This field can have one charity ID or a comma separated list of IDs. The method will return all the items associated with the specified charities.&lt;br /&gt;&lt;br /&gt; &lt;b&gt;For example:&lt;/b&gt;&lt;br /&gt;&lt;code&gt;/buy/browse/v1/item_summary/search?charity_ids&#x3D;13-1788491,300108469&lt;/code&gt;&lt;br /&gt;&lt;br /&gt;The charity ID is the charity&#39;s registration ID, also known as the Employer Identification Number (EIN). In GB, it is the Charity Registration Number (CRN), commonly called \&quot;Charity Number\&quot;.   &lt;ul&gt;&lt;li&gt;To find the charities eBay supports, you can search for a charity at &lt;a href&#x3D;\&quot;https://charity.ebay.com/search\&quot; target&#x3D;\&quot;_blank\&quot;&gt;Charity Search &lt;/a&gt; or go to &lt;a href&#x3D;\&quot;https://www.ebay.com/b/Charity/bn_7114598164\&quot; target&#x3D;\&quot;_blank\&quot;&gt;Charity Shop&lt;/a&gt;.&lt;/li&gt;   &lt;li&gt;To find the charity ID of a specific charity, click on a charity and use the EIN number. For example, the charity ID for  &lt;a href&#x3D;\&quot;https://charity.ebay.com/charity/American-Red-Cross/3843\&quot; target&#x3D;\&quot;_blank\&quot;&gt;American Red Cross&lt;/a&gt;, is &lt;code&gt;530196605&lt;/code&gt;.&lt;/li&gt;&lt;/ul&gt; You  can also use any combination of the &lt;code&gt;category_Ids&lt;/code&gt; and &lt;code&gt;q&lt;/code&gt; fields with a &lt;code&gt;charity_Ids&lt;/code&gt; to filter the result set. This gives you additional control over the result set. &lt;br /&gt;&lt;br /&gt;&lt;b&gt;Restriction: &lt;/b&gt; This is supported only on the US and GB marketplaces.&lt;br /&gt;&lt;br /&gt;&lt;b&gt;Maximum: &lt;/b&gt; 20 IDs &lt;br /&gt;&lt;br /&gt;&lt;b&gt;Required:&lt;/b&gt; One ID
func (r ApiSearchRequest) CharityIds(charityIds string) ApiSearchRequest {
	r.charityIds = &charityIds
	return r
}
// This field specifies the attributes used to define a specific product. The service searches for items matching the keyword or matching the keyword and a product attribute value in the title of the item.  &lt;br /&gt;&lt;br /&gt;For example, if the keyword is &lt;code&gt;brakes&lt;/code&gt; and &lt;code&gt;compatibility-filter&#x3D;Year:2018;Make:Honda&lt;/code&gt;, the items returned are items with brakes, 2018, or Honda in the title.  &lt;br /&gt;&lt;br /&gt;The service uses the product attributes to determine if the item is compatible.    The service returns the attributes that are compatible and the &lt;a href&#x3D;\&quot;/api-docs/buy/browse/resources/item_summary/methods/search#response.itemSummaries.compatibilityMatch\&quot;&gt; CompatibilityMatchEnum&lt;/a&gt; value that indicates how well the item matches the attributes.  &lt;br /&gt;&lt;br /&gt;For the best compatibility results, submit all the attributes used to define the product. &lt;br /&gt;&lt;br /&gt;&lt;b&gt;Best Practice: &lt;/b&gt; Submit all the &lt;a href&#x3D;\&quot;/api-docs/buy/static/api-browse.html#product-attributes\&quot;&gt;product attributes&lt;/a&gt; for the specific product.   &lt;br /&gt;&lt;br /&gt;For more details, see &lt;a href&#x3D;\&quot;/api-docs/buy/static/api-browse.html#Check\&quot;&gt;Check compatibility&lt;/a&gt; in the Buy Integration Guide.  &lt;br /&gt;&lt;br /&gt; &lt;span class&#x3D;\&quot;tablenote\&quot;&gt;&lt;b&gt;Note: &lt;/b&gt; The only products supported are cars, trucks, and motorcycles. &lt;/span&gt;  &lt;br /&gt;&lt;br /&gt; &lt;p&gt;For an example, see the &lt;a href&#x3D;\&quot;/api-docs/buy/browse/resources/item_summary/methods/search#w4-w1-w5-ReturnItemsthatareCompatiblewiththeKeywordandVehicle-9\&quot;&gt;Samples&lt;/a&gt; section. &lt;br /&gt;&lt;br /&gt; &lt;span class&#x3D;\&quot;tablenote\&quot;&gt;&lt;b&gt;Note: &lt;/b&gt; Testing in Sandbox is only supported using mock data. See &lt;a href&#x3D;\&quot;/api-docs/buy/static/api-browse.html#sbox-test\&quot;&gt;Testing search in the Sandbox&lt;/a&gt; for details. &lt;/span&gt;   &lt;br /&gt;&lt;br /&gt;&lt;b&gt;Required: &lt;/b&gt;&lt;ul&gt; &lt;li&gt;&lt;b&gt; q&lt;/b&gt; (keyword)&lt;/li&gt;  &lt;li&gt; one fitment supported category ID (such as &lt;code&gt;33559&lt;/code&gt; Brakes)&lt;/li&gt;  &lt;li&gt; a least one &lt;a href&#x3D;\&quot;/api-docs/buy/static/api-browse.html#product-attributes\&quot;&gt;product attribute name/value pair&lt;/a&gt;&lt;/li&gt;&lt;/ul&gt; For implementation help, refer to eBay API documentation at https://developer.ebay.com/api-docs/buy/browse/types/gct:CompatibilityFilter
func (r ApiSearchRequest) CompatibilityFilter(compatibilityFilter string) ApiSearchRequest {
	r.compatibilityFilter = &compatibilityFilter
	return r
}
// The ePID is the eBay product identifier of a product from the eBay product catalog. This field limits the results to only items in the specified ePID. &lt;br /&gt;&lt;br /&gt;The &lt;b&gt; Marketing&lt;/b&gt; API &lt;b&gt;getMerchandisedProducts&lt;/b&gt; method and the &lt;b&gt;Browse&lt;/b&gt; API &lt;b&gt; getItem&lt;/b&gt;, &lt;b&gt; getItemByLegacyId&lt;/b&gt;, and &lt;b&gt; getItemsByItemGroup&lt;/b&gt; calls return the ePID of the product.  You can also use the &lt;a href&#x3D;\&quot;/api-docs/commerce/catalog/resources/product_summary/methods/search\&quot;&gt;product_summary/search&lt;/a&gt; method in the &lt;b&gt;Catalog&lt;/b&gt; API to search for the ePID of the product. &lt;br /&gt;&lt;br /&gt;&lt;b&gt; For example: &lt;/b&gt;&lt;br/&gt;&lt;code&gt;/buy/browse/v1/item_summary/search?epid&#x3D;15032&lt;/code&gt; &lt;br /&gt;&lt;br /&gt;&lt;b&gt; Maximum: &lt;/b&gt; 1    &lt;br /&gt;&lt;br /&gt;&lt;b&gt; Required: &lt;/b&gt; The method must have  &lt;b&gt; category_ids&lt;/b&gt;, &lt;b&gt; epid&lt;/b&gt;,  &lt;b&gt; gtin&lt;/b&gt;,  or &lt;b&gt; q&lt;/b&gt;  (or any combination of these)
func (r ApiSearchRequest) Epid(epid string) ApiSearchRequest {
	r.epid = &epid
	return r
}
// This field is a comma separated list of values that lets you control what is returned in the response. The default is &lt;b&gt; MATCHING_ITEMS&lt;/b&gt;, which returns the items that match the keyword or category specified. The other values return data that can be used to create histograms or provide additional information.  &lt;br /&gt;&lt;br /&gt;&lt;b&gt; Valid Values: &lt;/b&gt; &lt;ul&gt;    &lt;li&gt;&lt;b&gt; ASPECT_REFINEMENTS&lt;/b&gt; - This returns the &lt;a href&#x3D;\&quot;#response.refinement.aspectDistributions\&quot;&gt;aspectDistributions&lt;/a&gt; container, which has the &lt;b&gt; dominantCategoryId&lt;/b&gt;, &lt;b&gt; matchCount&lt;/b&gt;, and &lt;b&gt; refinementHref&lt;/b&gt; for the various aspects of the items found. For example, if you searched for &#39;Mustang&#39;, some of the aspect would be &lt;b&gt; Model Year&lt;/b&gt;,  &lt;b&gt; Exterior Color&lt;/b&gt;, &lt;b&gt; Vehicle Mileage&lt;/b&gt;, etc. &lt;br /&gt; &lt;br /&gt;&lt;span class&#x3D;\&quot;tablenote\&quot;&gt; &lt;b&gt;Note: &lt;/b&gt; ASPECT_REFINEMENTS are category specific.&lt;/span&gt; &lt;br /&gt;&lt;br /&gt;&lt;/li&gt;   &lt;li&gt;&lt;b&gt; BUYING_OPTION_REFINEMENTS&lt;/b&gt; - This returns the &lt;a href&#x3D;\&quot;#response.refinement.buyingOptionDistributions\&quot;&gt;buyingOptionDistributions&lt;/a&gt;  container, which has the &lt;b&gt; matchCount&lt;/b&gt; and &lt;b&gt; refinementHref&lt;/b&gt; for &lt;b&gt; AUCTION&lt;/b&gt; and &lt;b&gt; FIXED_PRICE&lt;/b&gt; (Buy It Now) items. &lt;br /&gt;&lt;br /&gt;&lt;span class&#x3D;\&quot;tablenote\&quot;&gt; &lt;b&gt;Note: &lt;/b&gt;Classified items are not supported and only \&quot;Buy It Now\&quot; (non-auction) items are returned.&lt;/span&gt; &lt;br /&gt;&lt;br /&gt; &lt;/li&gt;   &lt;li&gt;&lt;b&gt; CATEGORY_REFINEMENTS&lt;/b&gt; - This returns the &lt;a href&#x3D;\&quot;#response.refinement.categoryDistributions\&quot;&gt;categoryDistributions&lt;/a&gt; container, which has the categories that the item is in.   &lt;/li&gt;   &lt;li&gt;&lt;b&gt; CONDITION_REFINEMENTS&lt;/b&gt; - This returns the &lt;a href&#x3D;\&quot;#response.refinement.conditionDistributions\&quot;&gt;conditionDistributions&lt;/a&gt;  container, such as &lt;b&gt; NEW&lt;/b&gt;, &lt;b&gt; USED&lt;/b&gt;, etc. Within these groups are multiple states of the condition. For example, &lt;b&gt; New &lt;/b&gt; can be New without tag, New in box, New without box, etc. &lt;/li&gt;   &lt;li&gt;&lt;b&gt; EXTENDED&lt;/b&gt; - This returns the &lt;a href&#x3D;\&quot;/api-docs/buy/browse/resources/item_summary/methods/search#response.itemSummaries.shortDescription\&quot;&gt;shortDescription&lt;/a&gt; field, which provides condition and item aspect information and the &lt;a href&#x3D;\&quot;/api-docs/buy/browse/resources/item_summary/methods/search#response.itemSummaries.itemLocation.city\&quot;&gt;itemLocation.city&lt;/a&gt; field.   &lt;/li&gt;  &lt;li&gt;&lt;b&gt; MATCHING_ITEMS&lt;/b&gt; - This is meant to be used with one or more of the refinement values above. You use this to return the specified refinements and all the matching items. &lt;/li&gt; &lt;li&gt;&lt;b&gt; FULL &lt;/b&gt; - This returns all the refinement containers and all the matching items.&lt;/li&gt;   &lt;/ul&gt; Code so that your app gracefully handles any future changes to this list.  &lt;br /&gt;&lt;br /&gt;&lt;b&gt;Default: &lt;/b&gt; MATCHING_ITEMS
func (r ApiSearchRequest) Fieldgroups(fieldgroups string) ApiSearchRequest {
	r.fieldgroups = &fieldgroups
	return r
}
// This field supports multiple field filters that can be used to limit/customize the result set. &lt;br /&gt;&lt;br /&gt;&lt;b&gt; For example: &lt;/b&gt;&lt;br /&gt;&lt;code&gt;/buy/browse/v1/item_summary/search?q&#x3D;shirt&amp;filter&#x3D;price:[10..50]&lt;/code&gt;&lt;br /&gt;&lt;br /&gt;You can also combine filters. &lt;br /&gt;&lt;code&gt;/buy/browse/v1/item_summary/search?q&#x3D;shirt&amp;filter&#x3D;price:[10..50],sellers:{rpseller|bigSal} &lt;/code&gt;   &lt;br /&gt;&lt;br /&gt;The following are the supported filters. For details and examples for all the filters, see &lt;a href&#x3D;\&quot;/api-docs/buy/static/ref-buy-browse-filters.html\&quot;&gt;Buy API Field Filters&lt;/a&gt;. &lt;div style&#x3D;\&quot;overflow-x:auto;\&quot;&gt; &lt;table&gt;  &lt;tr&gt;  &lt;td&gt;  &lt;ul&gt;  &lt;li&gt;&lt;a href&#x3D;\&quot;/api-docs/buy/static/ref-buy-browse-filters.html#bidCount\&quot;&gt;bidCount&lt;/a&gt; &lt;/li&gt;  &lt;li&gt;&lt;a href&#x3D;\&quot;/api-docs/buy/static/ref-buy-browse-filters.html#buyingOptions\&quot;&gt;buyingOptions&lt;/a&gt; &lt;/li&gt;  &lt;li&gt;&lt;a href&#x3D;\&quot;/api-docs/buy/static/ref-buy-browse-filters.html#charityOnly\&quot;&gt;charityOnly&lt;/a&gt; &lt;/li&gt;  &lt;li&gt;&lt;a href&#x3D;\&quot;/api-docs/buy/static/ref-buy-browse-filters.html#conditionIds\&quot;&gt;conditionIds&lt;/a&gt; &lt;/li&gt;    &lt;li&gt;&lt;a href&#x3D;\&quot;/api-docs/buy/static/ref-buy-browse-filters.html#conditions\&quot;&gt;conditions&lt;/a&gt; &lt;/li&gt;    &lt;li&gt;&lt;a href&#x3D;\&quot;/api-docs/buy/static/ref-buy-browse-filters.html#deliveryCountry\&quot;&gt;deliveryCountry&lt;/a&gt; &lt;/li&gt;    &lt;li&gt;&lt;a href&#x3D;\&quot;/api-docs/buy/static/ref-buy-browse-filters.html#deliveryOptions\&quot;&gt;deliveryOptions&lt;/a&gt; &lt;/li&gt;    &lt;li&gt;&lt;a href&#x3D;\&quot;/api-docs/buy/static/ref-buy-browse-filters.html#deliveryPostalCode\&quot;&gt;deliveryPostalCode&lt;/a&gt; &lt;/li&gt;   &lt;li&gt;&lt;a href&#x3D;\&quot;/api-docs/buy/static/ref-buy-browse-filters.html#excludeCategoryIds\&quot;&gt;excludeCategoryIds&lt;/a&gt; &lt;/li&gt;   &lt;li&gt;&lt;a href&#x3D;\&quot;/api-docs/buy/static/ref-buy-browse-filters.html#excludeSellers\&quot;&gt;excludeSellers&lt;/a&gt; &lt;/li&gt;  &lt;/ul&gt;&lt;/td&gt; &lt;td&gt;  &lt;ul&gt; &lt;li&gt;&lt;a href&#x3D;\&quot;/api-docs/buy/static/ref-buy-browse-filters.html#guaranteedDeliveryInDays\&quot;&gt;guaranteedDeliveryInDays&lt;/a&gt; &lt;/li&gt;     &lt;li&gt;&lt;a href&#x3D;\&quot;/api-docs/buy/static/ref-buy-browse-filters.html#itemEndDate\&quot;&gt;itemEndDate&lt;/a&gt; &lt;/li&gt;    &lt;li&gt;&lt;a href&#x3D;\&quot;/api-docs/buy/static/ref-buy-browse-filters.html#itemLocationCountry\&quot;&gt;itemLocationCountry&lt;/a&gt; &lt;/li&gt;    &lt;li&gt;&lt;a href&#x3D;\&quot;/api-docs/buy/static/ref-buy-browse-filters.html#itemStartDate\&quot;&gt;itemStartDate&lt;/a&gt; &lt;/li&gt;  &lt;li&gt;&lt;a href&#x3D;\&quot;/api-docs/buy/static/ref-buy-browse-filters.html#maxDeliveryCost\&quot;&gt;maxDeliveryCost&lt;/a&gt;&lt;/li&gt;  &lt;li&gt;&lt;a href&#x3D;\&quot;/api-docs/buy/static/ref-buy-browse-filters.html#paymentMethods\&quot;&gt;paymentMethods&lt;/a&gt; &lt;/li&gt;    &lt;li&gt;&lt;a href&#x3D;\&quot;/api-docs/buy/static/ref-buy-browse-filters.html#pickupCountry\&quot;&gt;pickupCountry&lt;/a&gt; &lt;/li&gt;    &lt;li&gt;&lt;a href&#x3D;\&quot;/api-docs/buy/static/ref-buy-browse-filters.html#pickupPostalCode\&quot;&gt;pickupPostalCode&lt;/a&gt; &lt;/li&gt;   &lt;li&gt;&lt;a href&#x3D;\&quot;/api-docs/buy/static/ref-buy-browse-filters.html#pickupRadius\&quot;&gt;pickupRadius&lt;/a&gt; &lt;/li&gt;  &lt;li&gt;&lt;a href&#x3D;\&quot;/api-docs/buy/static/ref-buy-browse-filters.html#pickupRadiusUnit\&quot;&gt;pickupRadiusUnit&lt;/a&gt; &lt;/li&gt;  &lt;/ul&gt;&lt;/td&gt; &lt;td&gt;  &lt;ul&gt;  &lt;li&gt;&lt;a href&#x3D;\&quot;/api-docs/buy/static/ref-buy-browse-filters.html#price\&quot;&gt;price&lt;/a&gt; &lt;/li&gt;    &lt;li&gt;&lt;a href&#x3D;\&quot;/api-docs/buy/static/ref-buy-browse-filters.html#priceCurrency\&quot;&gt;priceCurrency&lt;/a&gt; &lt;/li&gt;  &lt;li&gt;&lt;a href&#x3D;\&quot;/api-docs/buy/static/ref-buy-browse-filters.html#priorityListing\&quot;&gt;priorityListing&lt;/a&gt; &lt;/li&gt;  &lt;li&gt;&lt;a href&#x3D;\&quot;/api-docs/buy/static/ref-buy-browse-filters.html#qualifiedPrograms\&quot;&gt;qualifiedPrograms&lt;/a&gt; &lt;/li&gt;          &lt;li&gt;&lt;a href&#x3D;\&quot;/api-docs/buy/static/ref-buy-browse-filters.html#returnsAccepted\&quot;&gt;returnsAccepted&lt;/a&gt; &lt;/li&gt;    &lt;li&gt;&lt;a href&#x3D;\&quot;/api-docs/buy/static/ref-buy-browse-filters.html#searchInDescription\&quot;&gt;searchInDescription&lt;/a&gt; &lt;/li&gt;    &lt;li&gt;&lt;a href&#x3D;\&quot;/api-docs/buy/static/ref-buy-browse-filters.html#sellerAccountTypes\&quot;&gt;sellerAccountTypes&lt;/a&gt; &lt;/li&gt;    &lt;li&gt;&lt;a href&#x3D;\&quot;/api-docs/buy/static/ref-buy-browse-filters.html#sellers\&quot;&gt;sellers&lt;/a&gt; &lt;/li&gt;  &lt;/ul&gt;&lt;/td&gt;  &lt;/tr&gt;  &lt;/table&gt;  &lt;/div&gt; For implementation help, refer to eBay API documentation at https://developer.ebay.com/api-docs/buy/browse/types/cos:FilterField
func (r ApiSearchRequest) Filter(filter string) ApiSearchRequest {
	r.filter = &filter
	return r
}
// This field lets you search by the Global Trade Item Number of the item as defined by &lt;a href&#x3D;\&quot;https://www.gtin.info\&quot; target&#x3D;\&quot;_blank\&quot;&gt;https://www.gtin.info&lt;/a&gt;. You can search only by UPC (Universal Product Code). If you have other formats of GTIN, you need to search by keyword.  &lt;br /&gt;&lt;br /&gt;&lt;b&gt; For example: &lt;/b&gt;&lt;br/&gt;&lt;code&gt; /buy/browse/v1/item_summary/search?gtin&#x3D;099482432621&lt;/code&gt; &lt;br /&gt;&lt;br /&gt; &lt;b&gt; Maximum: &lt;/b&gt; 1     &lt;br /&gt;&lt;br /&gt;&lt;b&gt; Required: &lt;/b&gt; The method must have &lt;b&gt; category_ids&lt;/b&gt;, &lt;b&gt; epid&lt;/b&gt;, &lt;b&gt; gtin&lt;/b&gt;, or &lt;b&gt; q&lt;/b&gt; (or any combination of these)
func (r ApiSearchRequest) Gtin(gtin string) ApiSearchRequest {
	r.gtin = &gtin
	return r
}
// The number of items, from the result set, returned in a single page.  &lt;br /&gt;&lt;br /&gt;&lt;b&gt; Default:&lt;/b&gt; 50    &lt;br /&gt;&lt;br /&gt;&lt;b&gt; Maximum number of items per page (limit): &lt;/b&gt;200     &lt;br /&gt;&lt;br /&gt; &lt;b&gt; Maximum number of items in a result set: &lt;/b&gt; 10,000
func (r ApiSearchRequest) Limit(limit string) ApiSearchRequest {
	r.limit = &limit
	return r
}
// Specifies the number of items to skip in the result set. This is used with the &lt;b&gt; limit&lt;/b&gt; field to control the pagination of the output.  &lt;br /&gt;&lt;br /&gt;If &lt;b&gt; offset&lt;/b&gt; is 0 and &lt;b&gt; limit&lt;/b&gt; is 10, the method will retrieve items 1-10 from the list of items returned, if &lt;b&gt; offset&lt;/b&gt; is 10 and &lt;b&gt; limit&lt;/b&gt; is 10, the method will retrieve items 11 thru 20 from the list of items returned.  &lt;br /&gt;&lt;br /&gt;&lt;b&gt; Valid Values&lt;/b&gt;: 0-10,000 (inclusive)   &lt;br /&gt; &lt;br /&gt; &lt;b&gt; Default:&lt;/b&gt; 0    &lt;br /&gt; &lt;br /&gt; &lt;b&gt; Maximum number of items returned: &lt;/b&gt; 10,000  
func (r ApiSearchRequest) Offset(offset string) ApiSearchRequest {
	r.offset = &offset
	return r
}
// A string consisting of one or more keywords that are used to search for items on eBay. The keywords are handled as follows:&lt;ul&gt;&lt;li&gt;If the keywords are separated by a space, it is treated as an AND. In the following example, the query returns items that have iphone &lt;b&gt;AND&lt;/b&gt; ipad.&lt;br /&gt;&lt;br /&gt;&lt;code&gt;/buy/browse/v1/item_summary/search?q&#x3D;iphone ipad&lt;/code&gt;&lt;br/&gt;&lt;br /&gt;&lt;/li&gt;&lt;li&gt;If the keywords are input using parentheses and separated by a comma, or if they are URL-encoded, it is treated as an OR. In the following examples, the query returns items that have iphone &lt;b&gt;OR&lt;/b&gt; ipad.&lt;br /&gt;&lt;br /&gt;&lt;code&gt;/buy/browse/v1/item_summary/search?q&#x3D;(iphone, ipad)&lt;/code&gt;&lt;br /&gt;&lt;br /&gt;&lt;code&gt;/buy/browse/v1/item_summary/search?q&#x3D;%28iphone%2c%20ipad%29&lt;/code&gt;&lt;br /&gt;&lt;br /&gt;&lt;/li&gt;&lt;/ul&gt;&lt;b&gt;Restriction:&lt;/b&gt; The &lt;code&gt;*&lt;/code&gt; wildcard character is &lt;b&gt;not&lt;/b&gt; allowed in this field.&lt;br /&gt;&lt;br /&gt;&lt;b&gt;Required:&lt;/b&gt; The method must have &lt;b&gt;category_ids&lt;/b&gt;, &lt;b&gt;epid&lt;/b&gt;, &lt;b&gt;gtin&lt;/b&gt;, or &lt;b&gt;q&lt;/b&gt; (or any combination of these). 
func (r ApiSearchRequest) Q(q string) ApiSearchRequest {
	r.q = &q
	return r
}
// Specifies the order and the field name to use to sort the items. &lt;br /&gt;&lt;br /&gt;You can sort items by price (in ascending or descending order) or by distance (only applicable if the &lt;a href&#x3D;\&quot;/api-docs/buy/static/ref-buy-browse-filters.html#pickupCountry\&quot;&gt;\&quot;pickup\&quot; filters&lt;/a&gt; are used, and only ascending order is supported). You can also sort items by listing date, with the most recently listed (newest) items appearing first.&lt;br /&gt;&lt;br /&gt;&lt;span class&#x3D;\&quot;tablenote\&quot;&gt;&lt;b&gt;Note: &lt;/b&gt; To sort in descending order, insert a hyphen (&lt;code&gt;-&lt;/code&gt;) before the field name. If no &lt;b&gt;sort&lt;/b&gt; parameter is submitted, the result set is sorted by &amp;quot;&lt;a href&#x3D;\&quot;https://pages.ebay.com/help/sell/searchstanding.html\&quot; target&#x3D;\&quot;_blank\&quot;&gt;Best Match&lt;/a&gt;&amp;quot;.&lt;/span&gt;&lt;br /&gt;&lt;br /&gt;The following are examples of using the &lt;b&gt; sort&lt;/b&gt; query parameter.&lt;br /&gt;&lt;br /&gt;&lt;table&gt;&lt;tr&gt;&lt;th&gt;Sort&lt;/th&gt;&lt;th&gt;Result&lt;/th&gt;  &lt;/tr&gt; &lt;tr&gt; &lt;td&gt;&lt;code&gt;sort&#x3D;price&lt;/code&gt;&lt;/td&gt;  &lt;td&gt; Sorts by &lt;b&gt; price&lt;/b&gt; in ascending order (lowest price first)&lt;/td&gt; &lt;/tr&gt;   &lt;tr&gt;  &lt;td&gt;&lt;code&gt;sort&#x3D;-price&lt;/code&gt;&lt;/td&gt;  &lt;td&gt; Sorts by &lt;b&gt; price&lt;/b&gt; in descending order (highest price first)&lt;/td&gt; &lt;/tr&gt;   &lt;tr&gt;  &lt;td&gt;&lt;code&gt;sort&#x3D;distance&lt;/code&gt;&lt;/td&gt;  &lt;td&gt; Sorts by &lt;b&gt; distance&lt;/b&gt; in ascending order (shortest distance first)&lt;/td&gt; &lt;/tr&gt; &lt;tr&gt; &lt;td&gt;&lt;code&gt;sort&#x3D;newlyListed&lt;/code&gt;&lt;/td&gt;  &lt;td&gt;Sorts by &lt;b&gt;listing date&lt;/b&gt; (most recently listed/newest items first)&lt;/td&gt; &lt;/tr&gt; &lt;tr&gt; &lt;td&gt;&lt;code&gt;sort&#x3D;endingSoonest&lt;/code&gt;&lt;/td&gt;  &lt;td&gt;Sorts by &lt;b&gt;date/time&lt;/b&gt; the listing ends (listings nearest to end date/time first)&lt;/td&gt; &lt;/tr&gt;&lt;/table&gt;  &lt;br /&gt;&lt;b&gt; Default: &lt;/b&gt; Ascending For implementation help, refer to eBay API documentation at https://developer.ebay.com/api-docs/buy/browse/types/cos:SortField
func (r ApiSearchRequest) Sort(sort string) ApiSearchRequest {
	r.sort = &sort
	return r
}

func (r ApiSearchRequest) Execute() (SearchPagedCollection, *_nethttp.Response, error) {
	return r.ApiService.SearchExecute(r)
}

/*
Search Method for Search

<p>This method searches for eBay items by various query parameters and retrieves summaries of the items. You can search by keyword, category, eBay product ID (ePID), or GTIN, charity ID, or a combination of these. </p><p><span class="tablenote"><b> Note: </b>Only FIXED_PRICE (Buy It Now) items are returned by default. However, this method does return items where both FIXED_PRICE and AUCTION are available as a buying option. After a bid has been placed, items become active auction items and are no longer returned by default, but they remain accessible by filtering for the AUCTION buying option.</span></p><p>This method also supports the following:  <ul> <li>Filtering by the value of one or multiple fields, such as listing format, item condition, price range, location, and more.  For the fields supported by this method, see the <a href="#uri.filter">filter</a> parameter.</li> <li>Retrieving the refinements (metadata) of an item , such as item aspects (color, brand), condition, category, etc. using the <a href="#uri.fieldgroups">fieldgroups</a> parameter. </li>  <li>Filtering by item aspects and other refinements using the <a href="#uri.aspect_filter">aspect_filter</a> parameter. </li> <li>Filtering for items that are compatible with a specific product, using the <a href="#uri.compatibility_filter">compatibility_filter</a> parameter. </li><li>Creating aspects histograms, which enables shoppers to drill down in each refinement narrowing the search results.  </li>  </ul></p>  <p>For details and examples of these capabilities, see <a href="/api-docs/buy/static/api-browse.html">Browse API</a> in the Buying Integration Guide.</p>     <h3><b> Pagination and sort controls</b></h3>  <p>There are pagination controls (<b>limit</b> and <b>offset</b> fields) and <b> sort</b> query parameters that control/sort the data that is returned. By default, the results are sorted by &quot;Best Match&quot;. For more information about  Best Match, see the eBay help page <a href="https://pages.ebay.com/help/sell/searchstanding.html" target="_blank">Best Match</a>.  </p>    <h3><b>URLs for this method</b></h3>           <p><ul>            <li><b> Production URL: </b> <code>https://api.ebay.com/buy/browse/v1/item_summary/search?</code></li>            <li><b> Sandbox URL:  </b><code>https://api.sandbox.ebay.com/buy/browse/v1/item_summary/search?</code></li>           </ul>    </p>             <h3><b> Request headers</b></h3> This method uses the  <b>X-EBAY-C-ENDUSERCTX</b> request header to support revenue sharing for eBay Partner Networks and to improve the accuracy of shipping and delivery time estimations. For details see, <a href="/api-docs/buy/static/api-browse.html#Headers">Request headers</a> in the Buying Integration Guide.      <h3><b>Restrictions </b></h3> <p>This method can return a maximum of 10,000 items. For a list of supported sites and other restrictions, see <a href="/api-docs/buy/browse/overview.html#API">API Restrictions</a>.</p>    <span class="tablenote"><b>eBay Partner Network: </b> In order to receive a commission for your sales, you must use the URL returned in the <code>itemAffiliateWebUrl</code> field to forward your buyer to the ebay.com site. </span>

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSearchRequest
*/
func (a *ItemSummaryApiService) Search(ctx _context.Context) ApiSearchRequest {
	return ApiSearchRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return SearchPagedCollection
func (a *ItemSummaryApiService) SearchExecute(r ApiSearchRequest) (SearchPagedCollection, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  SearchPagedCollection
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ItemSummaryApiService.Search")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/item_summary/search"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.aspectFilter != nil {
		localVarQueryParams.Add("aspect_filter", parameterToString(*r.aspectFilter, ""))
	}
	if r.autoCorrect != nil {
		localVarQueryParams.Add("auto_correct", parameterToString(*r.autoCorrect, ""))
	}
	if r.categoryIds != nil {
		localVarQueryParams.Add("category_ids", parameterToString(*r.categoryIds, ""))
	}
	if r.charityIds != nil {
		localVarQueryParams.Add("charity_ids", parameterToString(*r.charityIds, ""))
	}
	if r.compatibilityFilter != nil {
		localVarQueryParams.Add("compatibility_filter", parameterToString(*r.compatibilityFilter, ""))
	}
	if r.epid != nil {
		localVarQueryParams.Add("epid", parameterToString(*r.epid, ""))
	}
	if r.fieldgroups != nil {
		localVarQueryParams.Add("fieldgroups", parameterToString(*r.fieldgroups, ""))
	}
	if r.filter != nil {
		localVarQueryParams.Add("filter", parameterToString(*r.filter, ""))
	}
	if r.gtin != nil {
		localVarQueryParams.Add("gtin", parameterToString(*r.gtin, ""))
	}
	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
	}
	if r.offset != nil {
		localVarQueryParams.Add("offset", parameterToString(*r.offset, ""))
	}
	if r.q != nil {
		localVarQueryParams.Add("q", parameterToString(*r.q, ""))
	}
	if r.sort != nil {
		localVarQueryParams.Add("sort", parameterToString(*r.sort, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
