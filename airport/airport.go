

package airport


import (
    "time"
)


type Airport struct {
    
    Profile
    Statistics
    Weather []string
    Runways []string
    Teminals []string
    Terminals
    IsHub bool
    Income float32
    LastExpansionDate time.Time
    
}