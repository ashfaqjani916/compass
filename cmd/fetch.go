package cmd

import "web-scraper/modules"

func FetchHackathons()  {
	modules.GetHackathonData()
}

func FetchJobs() {
	modules.GetJobData()
}
