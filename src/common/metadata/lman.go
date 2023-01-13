package metadata

/**
 * @Author: Tao Jun
 * @Since: 2023/1/11
 * @Desc: lman.go
**/
type SearchInstResultLman struct {
	Field string      `form:"field" json:"field" example:"模型属性名"`
	Value interface{} `form:"value" json:"value" swaggertype:"string" example:"对应数据"`
}
