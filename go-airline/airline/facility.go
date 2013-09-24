

package airline

import(
    "log"
    "encoding/xml"
    
    "github.com/FreeFlightSim/theairlineproject-server/go-airline/loaders"
)

// airlinefacilities.xml Snippet
/*
    <airlinefacilities>
        <airlinefacility uid="100" price="100000" fromyear="1970" monthlycost="10000">
            <level service="0" luxury="10"></level>
            <translations>
                <en-US name="Frequent Flyer Program" shortname="Flyer" />
                <es-ES name="Programa de Viajero Frecuente" shortname="Folleto" />
            </translations>
        </airlinefacility>
 */
var FACILITIES_XML = "/airlinefacilities.xml"



type AirlineFacilities struct {
    Facility []AirlineFacility `xml:"airlinefacility"`
}

type AirlineFacility struct{
	
    
    Uid uint `xml:"uid"` 

	Section string 

	Price string `xml:"price,attr"`
	
	MonthlyCost string `xml:"monthlycost,attr"`
	
	LuxuryLevel string `xml:"level>luxury,attr"` //NOT WORKING
	ServiceLevel uint
	FromYear int `xml:"fromyear,attr"`
	
	Name string
	SortName string
	
}



func LoadFacilities(){
    
    log.Println(FACILITIES_XML)
    buf, err := loaders.GetDataFileBytes(FACILITIES_XML)
    if err != nil {
        return
    }
    facils := new(AirlineFacilities)
    errx := xml.Unmarshal(buf, facils)
    
    if errx != nil{
        log.Fatalln("errun:", errx)
        return
    }
    //fmt.Fprintf("%v", facils)
    //for _, cust := range facils.Facility {
     //   
       // fmt.Printf("Airline-Facility: %v | %v | %v\n", cust.FromYear, cust.LuxuryLevel, cust.Price )
    //}
    log.Println("> ok")
}


