package airline

import (
    "time"
)


type AirlineLicense int //
const (
    LIC_Domestic   AirlineLicense   = iota
    LIC_Regional
    LIC_Short_Haul
    LIC_Long_Haul
)

type AirlineMentality int
const (
    Safe AirlineMentality   = iota
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




type Airline struct {
    
    Foo time.Time
    
    
    
    
}
