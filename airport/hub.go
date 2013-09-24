
package airport


import (
    //"time"
    
    //"github.com/FreeFlightSim/theairlineproject-server/airline"
)

type TypeOfHub int
const (
    HT_Hub TypeOfHub = iota
    HT_Regional
    HT_FocusCity
    HT_FortessHub
)


type HubType struct {
    Name string
    Price float32
    Type TypeOfHub
    
}
    