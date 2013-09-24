
package airline

import (
    "time"
)

type AirlineBudget struct {
    
    TotalBudget int64

     BudgetActive time.Time

     BudgetExpires time.Time

    MarketingBudget int64

    MaintenanceBudget int64

    SecurityBudget int64

    CSBudget int64

    PrintBudget int64

    TelevisionBudget int64

    RadioBudget int64

    InternetBudget int64
    OverhaulBudget int64
    PartsBudget int64
    EnginesBudget int64
    RemoteBudget int64
    InFlightBudget int64
    AirportBudget int64
    EquipmentBudget int64
    ITBudget int64
    CompBudget int64
    PromoBudget int64
    ServCenterBudget int64
    PRBudget int64
    EndYearCash int64
    FleetSize int
    FleetValue int64
    Subsidiaries int
    TotalSubValue int64
    TotalEmployees int
    TotalPayroll int
    RemainingBudget int64
    Cash int64
}