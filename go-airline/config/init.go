

package config

import(
    //"fmt"
   // "log"
   // "encoding/xml"
    //"github.com/FreeFlightSim/theairlineproject-server/go-airline/loaders"
    "github.com/FreeFlightSim/theairlineproject-server/go-airline/airline"
    "github.com/FreeFlightSim/theairlineproject-server/go-airline/airliner"
    "github.com/FreeFlightSim/theairlineproject-server/go-airline/finance"
)

func Load(){
    
    airline.LoadConfig()
    airliner.LoadConfig()
    
    finance.LoadConfig()
    
    
}





