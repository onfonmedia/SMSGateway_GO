/*
 * onfonmediasmsgateway_lib
 *
 * This file was automatically generated by APIMATIC v2.0 ( https://apimatic.io ).
 */

package template_pkg

import "onfonmediasmsgateway_lib/configuration_pkg"

/*
 * Interface for the TEMPLATE_IMPL
 */
type TEMPLATE interface {
    GetTemplateList () (interface{}, error)

    CreateNewTemplate (string, string) (interface{}, error)

    UpdateTemplate (string, string, int64) (interface{}, error)

    DeleteTemplate (int64) (interface{}, error)
}

/*
 * Factory for the TEMPLATE interaface returning TEMPLATE_IMPL
 */
func NewTEMPLATE(config configuration_pkg.CONFIGURATION) *TEMPLATE_IMPL {
    client := new(TEMPLATE_IMPL)
    client.config = config
    return client
}