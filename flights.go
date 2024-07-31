package main

import "time"

type FlightRequest struct {
	Date         time.Time `json:"date"`
	SrcCity      string    `json:"src_city"`
	DstCity      string    `json:"dst_city"`
	Adults       int       `json:"adults"`
	Children     int       `json:"children"`
	InfantInSeat int       `json:"infant_in_seat"`
	Stops        int       `json:"stops"`
	Class        int       `json:"class"`
}
