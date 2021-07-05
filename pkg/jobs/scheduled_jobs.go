package jobs

import (
	"github.com/go-co-op/gocron"
)

var Sch *gocron.Scheduler

func JobInit() {

	ProposalMakerCheckAvailability()

}
