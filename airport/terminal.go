

package airport

import (
    "time"
    
    "github.com/FreeFlightSim/theairlineproject-server/airline"
)

// Represents a Terminal Building
type Terminal struct {
    Name string
    DeliveryDate time.Time
    Airline airline.Airline
    Airport Airport
    Gates string
    IsBuilt bool
    IsBuyable bool  
}