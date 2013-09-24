

package stores

import(
    
    "log"
    
    "github.com/FreeFlightSim/theairlineproject-server/go-airline/config"
)


//= Fire up app here for now
func Initialise(){
    
    log.Println("Stores.Initialise()")
    config.Load()
    
}