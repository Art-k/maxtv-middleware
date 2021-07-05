package jobs

import (
	"fmt"
	"maxtv_middleware/pkg/common"
	"maxtv_middleware/pkg/db_interface"
	"net/http"
)

func ProposalMakerCheckAvailability() {

	jb, err := Sch.Every(1).Hour().Do(checkProposal)
	if err != nil {
		fmt.Println("Error ", err)
	} else {
		fmt.Println("Task, 'Check Proposal' next run :", jb.NextRun())
	}

}

func checkProposal() {

	if common.DB != nil {
		var users []db_interface.MaxtvUser
		err := common.DB.Where("access_level = ?", 200).Find(&users).Error
		if err != nil {
			fmt.Println(err.Error())
		}

		for _, user := range users {
			if user.ApiToken != "" {
				client := http.Client{}
				resp, err := client.Get("https://proposal.maxtvmedia.com/?token=" + user.ApiToken)
				if err != nil {
					fmt.Println("GET proposal, error received ", err.Error())
					common.PostTelegrammMessage("GET proposal, error received " + err.Error())
					break
				}
				if resp == nil {
					err := fmt.Errorf("GET proposal, empty response received")
					common.PostTelegrammMessage(err.Error())
					break
				}

				if resp.StatusCode != 200 {
					err := fmt.Errorf("GET proposal, status code is not 200, received %d", resp.StatusCode)
					fmt.Println(err.Error())
					common.PostTelegrammMessage(err.Error())
					break
				}

				if resp.StatusCode == 200 {
					fmt.Println("GET proposal, all good...")
					break
				}

			}
		}
	}

}
