# ProductSafetyLabels

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pictograms** | Pointer to [**[]ProductSafetyLabelPictogram**](ProductSafetyLabelPictogram.md) | An array of seller provided comma-separated string values that provides identifier, URL, and description for one or more pictograms associated with the listing. | [optional] 
**Statements** | Pointer to [**[]ProductSafetyLabelStatement**](ProductSafetyLabelStatement.md) | An array of seller provided comma-separated string values that provide identifier and description for one or more product safety statements associated with the listing. | [optional] 

## Methods

### NewProductSafetyLabels

`func NewProductSafetyLabels() *ProductSafetyLabels`

NewProductSafetyLabels instantiates a new ProductSafetyLabels object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductSafetyLabelsWithDefaults

`func NewProductSafetyLabelsWithDefaults() *ProductSafetyLabels`

NewProductSafetyLabelsWithDefaults instantiates a new ProductSafetyLabels object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPictograms

`func (o *ProductSafetyLabels) GetPictograms() []ProductSafetyLabelPictogram`

GetPictograms returns the Pictograms field if non-nil, zero value otherwise.

### GetPictogramsOk

`func (o *ProductSafetyLabels) GetPictogramsOk() (*[]ProductSafetyLabelPictogram, bool)`

GetPictogramsOk returns a tuple with the Pictograms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPictograms

`func (o *ProductSafetyLabels) SetPictograms(v []ProductSafetyLabelPictogram)`

SetPictograms sets Pictograms field to given value.

### HasPictograms

`func (o *ProductSafetyLabels) HasPictograms() bool`

HasPictograms returns a boolean if a field has been set.

### GetStatements

`func (o *ProductSafetyLabels) GetStatements() []ProductSafetyLabelStatement`

GetStatements returns the Statements field if non-nil, zero value otherwise.

### GetStatementsOk

`func (o *ProductSafetyLabels) GetStatementsOk() (*[]ProductSafetyLabelStatement, bool)`

GetStatementsOk returns a tuple with the Statements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatements

`func (o *ProductSafetyLabels) SetStatements(v []ProductSafetyLabelStatement)`

SetStatements sets Statements field to given value.

### HasStatements

`func (o *ProductSafetyLabels) HasStatements() bool`

HasStatements returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


