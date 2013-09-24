package airline

import (
    "time"
)


type AirlineLicense int
const (
    Domestic  AirlineLicense   = iota
    Regional
    Short_Haul
    Long_Haul
)

type AirlineMentality int
const (
    Safe  AirlineMentality   = iota
    Moderate
    Aggresive
)

type AirlineValue int
const (
    Very_low AirlineValue = iota
    Low
    Normal
    High
    Very_high
)

type AirlineFocus int
const (
    Global AirlineFocus = iota
    Regional
    Domestic
    Local
)
