package src

import (
	"time"

	"github.com/go-diydns/src/config"
	"github.com/go-diydns/src/dns"
	"github.com/go-diydns/src/providers"
	"github.com/go-diydns/utils"
)

type LastUpdated struct {
	Name        string
	LastUpdated int64
}

type AlreadyLogged struct {
	Name string
}

func RunUpdate(configPath string) {
	configs, err := config.ReadConfig(configPath)
	var lastUpdated []LastUpdated

	if err != nil {
		utils.PrintLog("error", err.Error())
	} else {
		for {
			for _, config := range configs {
				var found bool = false
				for _, last := range lastUpdated {
					if last.Name == config.Name {
						found = true
					}
				}

				if !found {
					lastUpdated = append(lastUpdated, LastUpdated{Name: config.Name, LastUpdated: time.Now().Unix()})
					updateOne(config)
				}

				for _, last := range lastUpdated {
					if last.Name == config.Name {
						if time.Now().Unix()-last.LastUpdated > int64(config.Interval) {
							last.LastUpdated = time.Now().Unix()
							updateOne(config)
						}
					}
				}
			}

			time.Sleep(time.Second * time.Duration(3))
		}
	}
}

var alreadyLogged []AlreadyLogged

func updateOne(config config.Config) {
	ip := dns.GetPublicIP()
	if ip == "" {
		utils.PrintLog("error", "Unable to get public IP, skipping update for "+config.Name)
	} else {

		res, err := dns.ResolveDomain(config.Domain)

		if err != nil {
			utils.PrintLog("info", "Domain for '"+config.Name+"' has no records, please create one manually")
		} else {
			if len(res) > 1 {
				utils.PrintLog("error", "Domain for '"+config.Name+"' has more than one record, please remove all but one")
			} else {
				if res[0] != ip {
					for i, logged := range alreadyLogged {
						if logged.Name == config.Name {
							alreadyLogged = append(alreadyLogged[:i], alreadyLogged[i+1:]...)
						}
					}

					utils.PrintLog("info", "Updating DNS record for "+config.Name+" to "+ip+" from "+res[0])

					switch config.Provider {
					case "cloudflare":
						providers.UpdateCloudflare(config, ip)
					default:
						utils.PrintLog("error", "Provider for '"+config.Name+"' is not supported")
					}
				} else {
					var found bool = false

					for _, logged := range alreadyLogged {
						if logged.Name == config.Name {
							found = true
						}
					}

					if !found {
						alreadyLogged = append(alreadyLogged, AlreadyLogged{Name: config.Name})
						utils.PrintLog("info", "DNS record for "+config.Name+" is up to date")
					}
				}
			}
		}
	}
}
