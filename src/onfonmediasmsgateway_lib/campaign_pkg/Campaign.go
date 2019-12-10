/*
 * onfonmediasmsgateway_lib
 *
 * This file was automatically generated by APIMATIC v2.0 ( https://apimatic.io ).
 */

package campaign_pkg

import "time"
import "onfonmediasmsgateway_lib/configuration_pkg"

/*
 * Interface for the CAMPAIGN_IMPL
 */
type CAMPAIGN interface {
    GetCampaignMessageStatus (int64) (interface{}, error)

    GetCampaigns (*time.Time, *time.Time, int64, int64) (interface{}, error)
}

/*
 * Factory for the CAMPAIGN interaface returning CAMPAIGN_IMPL
 */
func NewCAMPAIGN(config configuration_pkg.CONFIGURATION) *CAMPAIGN_IMPL {
    client := new(CAMPAIGN_IMPL)
    client.config = config
    return client
}