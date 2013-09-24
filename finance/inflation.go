

package finance

import(
    "log"
    "encoding/xml"
    "github.com/FreeFlightSim/theairlineproject-server/loaders"
)

// inflations.xml snippet
/*  
    <inflations>
        <inflation year="1960" fuelprice="0.119" inflation="0.0146" pricemodifier="1.000"/>
        <inflation year="1961" fuelprice="0.082" inflation="0.0107" pricemodifier="1.010"/>
*/
var INFLATION_XML = "/inflations.xml"


type Inflations struct {
    Inflations []Inflation `xml:"inflation"`
}

type Inflation struct{
    FuelPrice float32 `xml:"fuelprice,attr"` 
    InflationPercent float32 `xml:"inflation,attr"` 
    Modifier float32 `xml:"pricemodifier,attr"`
    Year uint `xml:"year,attr"`
}


func LoadInflations(){
    
    log.Println(INFLATION_XML)
    
    buf, err := loaders.GetDataFileBytes(INFLATION_XML)
    if err != nil {
        return
    }
    
    vals := new(Inflations)
    errx := xml.Unmarshal(buf, vals)
    if errx != nil{
        log.Println("Panic", errx)
        return
    }
    //fmt.Fprintf("%v", facils)
    //for _, inf := range vals.Inflations {
    //    
    //    fmt.Printf("Inflation: %v | %v | %v\n", inf.Year, inf.InflationPercent, inf.FuelPrice )
    //}
    log.Println("> ok")
}

