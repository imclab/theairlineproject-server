

package airport


import (
    "time"
    
    "github.com/FreeFlightSim/theairlineproject-server/airline"
)


type Airport struct {
    
    Profile string
    Statistics string
    Weather []string
    Runways []string
    Teminals []string
    Terminals string
    IsHub bool
    Income float32
    LastExpansionDate time.Time
    
}



type AirportContract struct {
    
    Airline airline.Airline
    Airport Airport
    ContractDate time.Time
    Length int
    YearlyPayment float32
    NumberOfGates uint // from xplane
    IsExclusiveDeal bool
    MonthsLeft int
    Terminal string
    PayFull bool 
}