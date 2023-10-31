package models

type LogRow struct {
	Lat  float64                `json:"latitude"`
	Lon  float64                `json:"longitude"`
	Id   string                 `json:"id"`
	Tags map[string]interface{} `json:"tags"`
}
