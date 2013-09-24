

package airline

import (
    "time"
)


type AllianceType int
const (
    Codesharing  AllianceType   = iota
    Full
)


type Alliance struct {
    
    Type AllianceType
    Name string
    Members int // TODO
    FormationDate time.Time
    
    Logo string
    IsHumanAlliance bool // WTF is this ?
    
    
}


type AllianceMember {
    Airline Airline
    JoinedDate time.Time
    ExitedDate time.Time
}