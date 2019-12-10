/*
 * onfonmediasmsgateway_lib
 *
 * This file was automatically generated by APIMATIC v2.0 ( https://apimatic.io ).
 */

package configuration_pkg


/** The base Uri for API calls */
const BASEURI string = "https://api.onfonmedia.co.ke/v1/sms"



type CONFIGURATION_IMPL struct {
    /** Network Layer Access Key */
    /** Replace the value of AccessKey with an appropriate value */
    accesskey string
    /** Used for authentication purpose and pass this parameter in URL encoded format. */
    /** Replace the value of ApiKey with an appropriate value */
    apikey string
    /** Used for authentication purpose and pass this parameter in URL encoded format. */
    /** Replace the value of ClientId with an appropriate value */
    clientid string
}

/*
 * Getter function returning accesskey
 */
func (me *CONFIGURATION_IMPL) AccessKey() string{
    return me.accesskey
}

/*
 * Setter function setting up accesskey
 */
func (me *CONFIGURATION_IMPL) SetAccessKey(accessKey string) {
    me.accesskey = accessKey
}

/*
 * Getter function returning apikey
 */
func (me *CONFIGURATION_IMPL) ApiKey() string{
    return me.apikey
}

/*
 * Setter function setting up apikey
 */
func (me *CONFIGURATION_IMPL) SetApiKey(apiKey string) {
    me.apikey = apiKey
}

/*
 * Getter function returning clientid
 */
func (me *CONFIGURATION_IMPL) ClientId() string{
    return me.clientid
}

/*
 * Setter function setting up clientid
 */
func (me *CONFIGURATION_IMPL) SetClientId(clientId string) {
    me.clientid = clientId
}
