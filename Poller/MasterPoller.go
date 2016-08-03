package poller

import (
	//Standard Package
	"log"
	//Packages for this project
	//common
)

//PollMasters a structure that will poll all the masters
type PollMasters struct {
	DCs                //[]common.typ.DC
	Frequency int      //in seconds polling frequency
	ToPE      chan int //Channel to Policy Engine
}

//Returns a new PollMasters Variable
func NewPollMasters(Masters []Endpoint) *PollMasters {
}

//Poll Will Poll each master to obtain remaining resources and store that in its respective DC
func (P *PollMasters) Poll() {

	//Loop the DC array and poll each master for the remaining resources

}

func Run(Masters []Endpoint) {

	//populate DCs
	PM := NewPollMasters()
	//Should be

	for {
		select {
		case <-time.After(10):
			break
		case <-PM.ToPE:
			break
		}
	}
}
