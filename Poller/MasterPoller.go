package poller

import (
	//Standard Package
	"log"
	"time"
	//Packages for this project
	"common/quotalib"
	"common/types"
	"../Config"
)

//PollMasters a structure that will poll all the masters
type PollMasters struct {
	DCs     typ.DC          //[]common.typ.DC
	Frequency int      //in seconds polling frequency
	ToPE      chan int //Channel to Policy Engine
}

//Returns a new PollMasters Variable
func NewPollMasters(Masters []config.DCList) *PollMasters {
		return &PollMasters{}
}


//Poll Will Poll each master to obtain remaining resources and store that in its respective DC
func (P *PollMasters) Poll(Masters []config.DCList) {
	var role string

	//Loop the DC array and poll each master for the remaining resources
	for _,val := range Masters {
		P.DCs.Endpoint = val.Master[0].Ip
		role = val.Master[0].Role
		cpu,mem,disk,err := quotalib.RemainingResource(P.DCs,role)
		if err != nil {
			log.Println("RemainingResource:%v",err)
		}
		P.DCs.Ucpu = cpu
		P.DCs.Umem = mem
		P.DCs.Udisk = disk
		log.Println("values are\n",cpu,mem,disk)	
	}

}

func Run(Masters []config.DCList) {

	//populate DCs
	PM := NewPollMasters(Masters)
	PM.Poll(Masters)
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
