/*
 * onfonmediasmsgateway_lib
 *
 * This file was automatically generated by APIMATIC v2.0 ( https://apimatic.io ).
 */

package OnfonMediaSMSGatewayClient

import(
	"onfonmediasmsgateway_lib/configuration_pkg"
	"onfonmediasmsgateway_lib/account_pkg"
	"onfonmediasmsgateway_lib/template_pkg"
	"onfonmediasmsgateway_lib/sms_pkg"
	"onfonmediasmsgateway_lib/group_pkg"
	"onfonmediasmsgateway_lib/campaign_pkg"
)
/*
 * Client structure as interface implementation
 */
type ONFONMEDIASMSGATEWAY_IMPL struct {
     account account_pkg.ACCOUNT
     template template_pkg.TEMPLATE
     sms sms_pkg.SMS
     group group_pkg.GROUP
     campaign campaign_pkg.CAMPAIGN
     config  configuration_pkg.CONFIGURATION
}

/**
     * Access to Configuration
     * @return Returns the Configuration instance
*/
func (me *ONFONMEDIASMSGATEWAY_IMPL) Configuration() configuration_pkg.CONFIGURATION {
    return me.config
}
/**
     * Access to Account controller
     * @return Returns the Account() instance
*/
func (me *ONFONMEDIASMSGATEWAY_IMPL) Account() account_pkg.ACCOUNT {
    if(me.account) == nil {
        me.account = account_pkg.NewACCOUNT(me.config)
    }
    return me.account
}
/**
     * Access to Template controller
     * @return Returns the Template() instance
*/
func (me *ONFONMEDIASMSGATEWAY_IMPL) Template() template_pkg.TEMPLATE {
    if(me.template) == nil {
        me.template = template_pkg.NewTEMPLATE(me.config)
    }
    return me.template
}
/**
     * Access to SMS controller
     * @return Returns the SMS() instance
*/
func (me *ONFONMEDIASMSGATEWAY_IMPL) SMS() sms_pkg.SMS {
    if(me.sms) == nil {
        me.sms = sms_pkg.NewSMS(me.config)
    }
    return me.sms
}
/**
     * Access to GROUP controller
     * @return Returns the GROUP() instance
*/
func (me *ONFONMEDIASMSGATEWAY_IMPL) GROUP() group_pkg.GROUP {
    if(me.group) == nil {
        me.group = group_pkg.NewGROUP(me.config)
    }
    return me.group
}
/**
     * Access to Campaign controller
     * @return Returns the Campaign() instance
*/
func (me *ONFONMEDIASMSGATEWAY_IMPL) Campaign() campaign_pkg.CAMPAIGN {
    if(me.campaign) == nil {
        me.campaign = campaign_pkg.NewCAMPAIGN(me.config)
    }
    return me.campaign
}

