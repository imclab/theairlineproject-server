

package airline

import (
    "time"
)


type InsuranceType int
const (
    Public_Liability  InsuranceType   = iota
    Passenger_Liability     
    Combined_Single_Limit   
    Full_Coverage     
)

type InsuranceScope int
const (
    IN_Airport  InsuranceScope   = iota
    IN_Domestic     
    IN_Hub   
    IN_Global     
)

type InsuranceTerms int
const (
    Annual  InsuranceTerms   = iota
    Biannual     
    Quarterly   
    Monthly     
)


type AirlineInsurance struct {
    InsType InsuranceType
    InsScope InsuranceScope
    InsTerms InsuranceTerms

    InsuredAmount int

    Deductible float32

    TermLength int
    PaymentAmount float32

    CancellationFee int

    PolicyIndex string 

    AllFleetAirliners bool 

    InsuranceEffective time.Time
    InsuranceExpires time.Time
    NextPaymentDue time.Time
    RemainingPayments int
    
}