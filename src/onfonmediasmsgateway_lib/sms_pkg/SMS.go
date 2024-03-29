/*
 * onfonmediasmsgateway_lib
 *
 * This file was automatically generated by APIMATIC v2.0 ( https://apimatic.io ).
 */

package sms_pkg

import "time"
import "onfonmediasmsgateway_lib/configuration_pkg"

/*
 * Interface for the SMS_IMPL
 */
type SMS interface {
    GetSentMessageList (*time.Time, *time.Time, int64, int64) (interface{}, error)

    GetSentMessageStatus (int64) (interface{}, error)

    GetCreateSMS (string, string, string, *string, *bool, *bool, *string, *string, *string, *string) (interface{}, error)

    CreateSMS (string, string, string, *string, *bool, *bool, *string, *string, *string, *string) (interface{}, error)

    GetCreateBulkSMS (string, string, *string, *bool, *bool, *string, *time.Time, *string) (interface{}, error)

    CreateBulkSMS ([]string, string, *bool, *bool, *time.Time) (interface{}, error)
}

/*
 * Factory for the SMS interaface returning SMS_IMPL
 */
func NewSMS(config configuration_pkg.CONFIGURATION) *SMS_IMPL {
    client := new(SMS_IMPL)
    client.config = config
    return client
}