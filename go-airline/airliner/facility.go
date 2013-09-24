

package airliner

import(
    "fmt"
    "log"
    "encoding/xml"
    "github.com/FreeFlightSim/theairlineproject-server/go-airline/loaders"
)



// airlinerfacilities.xml snippet
/*
    <airlinerfacilities> *
        <airlinerfacility uid="100" type="Audio" fromyear="1950">
            <level service="0"></level>
            <seats percent="0" price="0" uses="1"></seats>
            <translations>
                <en-US name="No Audio" />
                <zh-CN name="没有音乐服务" />
            </translations>
        </airlinerfacility>
*/
var FACILITIES_XML = "/airlinerfacilities.xml"

type AirlinerFacilities struct {
    Facility []AirlinerFacility `xml:"airlinerfacility"`
}

type AirlinerFacility struct{
    Section string 
    Uid uint `xml:"uid,attr"`    
    
    PricePerSeat float32 `xml:"seats>price,attr"`
    PercentOfSeats float32 `xml:"seats>percents,attr"`
    SeatUses float32 `xml:"seats>uses,attr"`
    
    Type string `xml:"monthlycost,attr"`
    ServiceLevel uint `xml:"level>service,attr"`
    
    FromYear uint `xml:"fromyear,attr"` 
    
    Name string 
}


func LoadFacilities(){
    
    //log.Println(file_path)
    buf, err := loaders.GetDataFileBytes(FACILITIES_XML)
    if err != nil {
        return
    }
    facils := new(AirlinerFacilities)
    errx := xml.Unmarshal(buf, facils)
    
    if errx != nil{
        log.Println("errun:", errx)
        return
    }
    //fmt.Fprintf("%v", facils)
    for _, cust := range facils.Facility {
        
        fmt.Printf("Airline-Facility: %v | %v | %v\n", cust.Uid, cust.FromYear, cust.Type )
    }
    fmt.Println("DONE")
}

