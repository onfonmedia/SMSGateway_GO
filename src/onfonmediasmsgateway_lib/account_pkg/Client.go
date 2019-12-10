/*
 * onfonmediasmsgateway_lib
 *
 * This file was automatically generated by APIMATIC v2.0 ( https://apimatic.io ).
 */

package account_pkg


import(
	"encoding/json"
	"github.com/apimatic/unirest-go"
	"onfonmediasmsgateway_lib/apihelper_pkg"
	"onfonmediasmsgateway_lib/configuration_pkg"
)
/*
 * Client structure as interface implementation
 */
type ACCOUNT_IMPL struct {
     config configuration_pkg.CONFIGURATION
}

/**
 * Get Credit Balance
 * @return	Returns the interface{} response from the API call
 */
func (me *ACCOUNT_IMPL) GetCreditBalance () (interface{}, error) {
    //the endpoint path uri
    _pathUrl := "/Balance"

    //variable to hold errors
    var err error = nil
    //the base uri for api requests
    _queryBuilder := configuration_pkg.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + _pathUrl

    //process optional query parameters
    _queryBuilder, err = apihelper_pkg.AppendUrlWithQueryParameters(_queryBuilder, map[string]interface{} {
        "ApiKey" : onfonmediasmsgateway_lib.config.ApiKey,
        "ClientId" : onfonmediasmsgateway_lib.config.ClientId,
    })
    if err != nil {
        //error in query param handling
        return nil, err
    }

    //validate and preprocess url
    _queryBuilder, err = apihelper_pkg.CleanUrl(_queryBuilder)
    if err != nil {
        //error in url validation or cleaning
        return nil, err
    }
    //prepare headers for the outgoing request
    headers := map[string]interface{} {
        "user-agent" : "APIMATIC 2.0",
        "accept" : "application/json",
    }

    //prepare API request
    _request := unirest.Get(_queryBuilder, headers)
    //and invoke the API call request to fetch the response
    _response, err := unirest.AsString(_request,false);
    if err != nil {
        //error in API invocation
        return nil, err
    }

    //error handling using HTTP status codes
    if (_response.Code < 200) || (_response.Code > 206) { //[200,206] = HTTP OK
        err = apihelper_pkg.NewAPIError("HTTP Response Not OK" , _response.Code, _response.RawBody)
    }
    if(err != nil) {
        //error detected in status code validation
        return nil, err
    }

    //returning the response
    var retVal interface{}
    err = json.Unmarshal(_response.RawBody, &retVal)

    if err != nil {
        //error in parsing
        return nil, err
    }
    return retVal, nil

}

