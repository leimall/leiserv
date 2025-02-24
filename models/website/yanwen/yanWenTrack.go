package model

import "time"

// type YanWenTrackResponse struct {
// 	Success bool   `json:"success"`
// 	Code    string `json:"code"`
// 	Message string `json:"message"`
// }

type YanWenTrackResponse struct {
	Code                string              `json:"code"`
	Message             string              `json:"message"`
	Result              []Result            `json:"result"`
	RequestTime         time.Time           `json:"requestTime"`
	ElapsedMilliseconds ElapsedMilliseconds `json:"elapsedMilliseconds"`
}
type ExtraProperties struct {
	City     string `json:"City"`
	Country  string `json:"Country"`
	State    string `json:"State"`
	Attached string `json:"Attached"`
	Image    string `json:"Image"`
	Images   string `json:"Images"`
	MqpID    int    `json:"MqpId"`
}
type Checkpoints struct {
	TimeStamp            string          `json:"time_stamp"`
	TimeZone             string          `json:"time_zone"`
	TrackingStatus       string          `json:"tracking_status"`
	Message              string          `json:"message"`
	Location             string          `json:"location"`
	IsLastMileCheckpoint int             `json:"is_last_mile_checkpoint"`
	ExtraProperties      ExtraProperties `json:"extraProperties,omitempty"`
	TimeCreate           string          `json:"time_create"`
}
type TrackingStatusWaybill struct {
	Level1 string `json:"level1"`
	Level2 string `json:"level2"`
	Level3 string `json:"level3"`
	Level4 string `json:"level4"`
}
type Result struct {
	TrackingNumber               string                `json:"tracking_number"`
	WaybillNumber                string                `json:"waybill_number"`
	ExchangeNumber               string                `json:"exchange_number"`
	YanwenNumber                 string                `json:"yanwen_number"`
	OrderNumber                  string                `json:"order_number"`
	Checkpoints                  []Checkpoints         `json:"checkpoints"`
	TrackingStatus               string                `json:"tracking_status"`
	TrackingStatusWaybill        TrackingStatusWaybill `json:"tracking_status_waybill"`
	LastMileTrackingExpected     bool                  `json:"last_mile_tracking_expected"`
	OriginCountry                interface{}           `json:"origin_country"`
	DestinationCountry           string                `json:"destination_country"`
	LastMileCarrier              interface{}           `json:"last_mile_carrier"`
	LastMileCarrierWebsite       interface{}           `json:"last_mile_carrier_website"`
	LastMileCarrierContactNumber interface{}           `json:"last_mile_carrier_contact_number"`
}
type ElapsedMilliseconds struct {
	Total int `json:"total"`
}
