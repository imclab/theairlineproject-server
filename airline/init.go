
// Represents an Airline
package airline

import(
   // "fmt"
)

type AirlineModel struct{
    Facilities AirlineFacilities
    AirlineBudget
}


func LoadConfig(){
    
    LoadFacilities()
    
}
