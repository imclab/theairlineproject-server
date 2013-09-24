


package airport

import (
    "time"
)


type Gate struct {
    DeliveryDate time.Time
}

type Gates struct {
    DeliveryDate time.Time
    Gates []Gate
}

func (me Gates) AddGate(g Gate) {
    //todo
}