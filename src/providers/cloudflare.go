package providers

import (
	"context"
	"fmt"

	"github.com/go-diydns/src/config"
	"github.com/cloudflare/cloudflare-go"
	"github.com/go-diydns/utils"
)

func UpdateCloudflare(config config.Config, ip string) {
	api, err := cloudflare.NewWithAPIToken(config.Token)

	if err != nil {
		utils.PrintLog("error", "Unable to connect to Cloudflare, skipping update for " + config.Name)
	} else {
		zoneID, err := api.ZoneIDByName(utils.GetHostname(config.Domain))

		if err != nil {
			fmt.Println(err)
			utils.PrintLog("error", "Unable to get zone ID from Cloudflare, skipping update for " + config.Name)
		} else {
			recs, _, err := api.ListDNSRecords(context.Background(), cloudflare.ZoneIdentifier(zoneID), cloudflare.ListDNSRecordsParams{})

			if err != nil {
				utils.PrintLog("error", "Unable to get DNS records from Cloudflare, skipping update for " + config.Name)
			} else {
				for _, rec := range recs {
					if(rec.Name == config.Domain && rec.Type == "A") {						
						if(rec.Content != ip) {
							updated := cloudflare.UpdateDNSRecordParams{
								ID: rec.ID,
								Type: rec.Type,
								Name: rec.Name,
								Content: ip,
								Proxied: rec.Proxied,
								TTL: rec.TTL,
							}

							err := api.UpdateDNSRecord(context.Background(), cloudflare.ZoneIdentifier(zoneID), updated)

							if err != nil {
								utils.PrintLog("error", "Unable to update DNS record in Cloudflare, skipping update for " + config.Name)
							} else {
								utils.PrintLog("info", "Updated DNS record in Cloudflare for " + config.Name)
							}
						}
					}
				}
			}
		}
	}
}