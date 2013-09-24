

package airline



type InsuranceType int
const (
    Public_Liability  InsuranceType   = iota
    Passenger_Liability     
    Combined_Single_Limit   
    Full_Coverage     
)

type InsuranceScope int
const (
    Airport  InsuranceScope   = iota
    Domestic     
    Hub   
    Global     
)