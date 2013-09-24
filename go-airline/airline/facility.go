

package airline



//
type Facility struct{
	
	Uid uint `json:"uid"` 
	
	// 
	Section string `json:"section"`
	
	// 
	Price uint `json:"price"`
	
	// 
	MonthlyCost string `json:"monthly_cost"`
	
	LuxuryLevel uint
	ServiceLevel uint
	FromEar int
	
	Name string
	SortName string
	
}



type Facilities struct{
	List *Facility[]
}
