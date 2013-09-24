

package airline


/*
 < airlinefacilities> *
 <airlinefacility uid="100" price="100000" fromyear="1970" monthlycost="10000">
 <level service="0" luxury="10">
 </level>
 <translations>
 <en-US name="Frequent Flyer Program" shortname="Flyer" />
 <de-DE name="Werbemaßnahme wechselnde Broschüre" shortname="Broschüre" />
 <en-GB name="Frequent Flyer Program" shortname="Flyer" />
 <nl-NL name="Frequent flyer-programma" shortname="Passagier" />
 <zh-CN name="飞行常客奖励计划" shortname="常客" />
 <da-DK name="Loyalitetsprogram" shortname="Flyer" />
 <es-ES name="Programa de Viajero Frecuente" shortname="Folleto" />
 </translations>
 </airlinefacility>
 */

//
type AirlineFacility struct{
	
    Uid uint `xml:"uid"` //`json:"uid"` 
	
	// 
	Section string //`xml:"section"`
	
	// 
	Price uint `xml:"price"`
	
	// 
	MonthlyCost string `xml:"monthlycost"`
	
	LuxuryLevel uint
	ServiceLevel uint
	FromEar int
	
	Name string
	SortName string
	
}



type Facilities struct{
	List *Facility[]
}
