

package config

import(
    //"fmt"

    "github.com/FreeFlightSim/theairlineproject-server/airline"
    "github.com/FreeFlightSim/theairlineproject-server/airliner"
    "github.com/FreeFlightSim/theairlineproject-server/finance"
)

func Load(){
    
    airline.LoadConfig()
    airliner.LoadConfig()
    
    finance.LoadConfig()
    
    
}





