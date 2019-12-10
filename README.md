# Getting started

Send SMS with Onfon Media SMS Platform.

## How to Build


* In order to successfully build and run your SDK files, you are required to have the following setup in your system:
    * **Go**  (Visit [https://golang.org/doc/install](https://golang.org/doc/install) for more details on how to install Go)
    * **Java VM** Version 8 or later
    * **Eclipse 4.6 (Neon)** or later ([http://www.eclipse.org/neon/](http://www.eclipse.org/neon/))
    * **GoClipse** setup within above installed Eclipse (Follow the instructions at [this link](https://github.com/GoClipse/goclipse/blob/latest/documentation/Installation.md#instructions) to setup GoClipse)
* Ensure that ```GOPATH``` environment variable is set in the system variables. If not, set it to your workspace directory where you will be adding your Go projects.
* The generated code uses unirest-go http library. Therefore, you will need internet access to resolve this dependency. If Go is properly installed and configured, run the following command to pull the dependency:

```Go
go get github.com/apimatic/unirest-go
```

This will install unirest-go in the ```GOPATH``` you specified in the system variables.

Now follow the steps mentioned below to build your SDK:

1. Open eclipse in the Go language perspective and click on the ```Import``` option in ```File``` menu.

![Importing SDK into Eclipse - Step 1](https://apidocs.io/illustration/go?step=import0)

2. Select ```General -> Existing Projects into Workspace``` option from the tree list.

![Importing SDK into Eclipse - Step 2](https://apidocs.io/illustration/go?step=import1)

3. In ```Select root directory```, provide path to the unzipped archive for the generated code. Once the path is set and the Project becomes visible under ```Projects``` click ```Finish```

![Importing SDK into Eclipse - Step 3](https://apidocs.io/illustration/go?step=import2&workspaceFolder=OnfonMedia%20SMS%20Gateway-GoLang&projectName=onfonmediasmsgateway_lib)

4. The Go library will be imported and its files will be visible in the Project Explorer

![Importing SDK into Eclipse - Step 4](https://apidocs.io/illustration/go?step=import3&projectName=onfonmediasmsgateway_lib)

## How to Use

The following section explains how to use the OnfonmediasmsgatewayLib library in a new project.

### 1. Add a new Test Project

Create a new project in Eclipse by ```File``` -> ```New``` -> ```Go Project```

![Add a new project in Eclipse](https://apidocs.io/illustration/go?step=createNewProject0)

Name the Project as ```Test``` and click ```Finish```

![Create a new Maven Project - Step 1](https://apidocs.io/illustration/go?step=createNewProject1)

Create a new directory in the ```src``` directory of this project

![Create a new Maven Project - Step 2](https://apidocs.io/illustration/go?step=createNewProject2&projectName=onfonmediasmsgateway_lib)

Name it ```test.com```

![Create a new Maven Project - Step 3](https://apidocs.io/illustration/go?step=createNewProject3&projectName=onfonmediasmsgateway_lib)

Now create a new file inside ```src/test.com```

![Create a new Maven Project - Step 4](https://apidocs.io/illustration/go?step=createNewProject4&projectName=onfonmediasmsgateway_lib)

Name it ```testsdk.go```

![Create a new Maven Project - Step 5](https://apidocs.io/illustration/go?step=createNewProject5&projectName=onfonmediasmsgateway_lib)

In this Go file, you can start adding code to initialize the client library. Sample code to initialize the client library and using its methods is given in the subsequent sections.

### 2. Configure the Test Project

You need to import your generated library in this project in order to make use of its functions. In order to import the library, you can add its path in the ```GOPATH``` for this project. Follow the below steps:

Right click on the project name and click on ```Properties```

![Adding dependency to the client library - Step 1](https://apidocs.io/illustration/go?step=testProject0&projectName=onfonmediasmsgateway_lib)

Choose ```Go Compiler``` from the side menu. Check ```Use project specific settings``` and uncheck ```Use same value as the GOPATH environment variable.```. By default, the GOPATH value from the environment variables will be visible in ```Eclipse GOPATH```. Do not remove this as this points to the unirest dependency.

![Adding dependency to the client library - Step 2](https://apidocs.io/illustration/go?step=testProject1)

Append the library path to this GOPATH

![Adding dependency to the client library - Step 3](https://apidocs.io/illustration/go?step=testProject2&workspaceFolder=OnfonMedia%20SMS%20Gateway-GoLang)

Once the path is appended, click on ```OK```

### 3. Build the Test Project

Right click on the project name and click on ```Build Project```

![Build Project](https://apidocs.io/illustration/go?step=buildProject&projectName=onfonmediasmsgateway_lib)

### 4. Run the Test Project

If the build is successful, right click on your Go file and click on ```Run As``` -> ```Go Application```

![Run Project](https://apidocs.io/illustration/go?step=runProject&projectName=onfonmediasmsgateway_lib)

## Initialization

### Authentication
In order to setup authentication of the API client, you need the following information.

| Parameter | Description |
|-----------|-------------|
| accessKey | Network Layer Access Key |
| apiKey | Used for authentication purpose and pass this parameter in URL encoded format. |
| clientId | Used for authentication purpose and pass this parameter in URL encoded format. |


To configure these for your generated code, open the file "Configuration.go" and edit it's contents.


# Class Reference

## <a name="list_of_controllers"></a>List of Controllers

* [account_pkg](#account_pkg)
* [template_pkg](#template_pkg)
* [sms_pkg](#sms_pkg)
* [group_pkg](#group_pkg)
* [campaign_pkg](#campaign_pkg)

## <a name="account_pkg"></a>![Class: ](https://apidocs.io/img/class.png ".account_pkg") account_pkg

### Get instance

Factory for the ``` ACCOUNT ``` interface can be accessed from the package account_pkg.

```go
account := account_pkg.NewACCOUNT()
```

### <a name="get_credit_balance"></a>![Method: ](https://apidocs.io/img/method.png ".account_pkg.GetCreditBalance") GetCreditBalance

> Get Credit Balance


```go
func (me *ACCOUNT_IMPL) GetCreditBalance()(interface{},error)
```

#### Example Usage

```go

var result interface{}
result,_ = account.GetCreditBalance()

```


[Back to List of Controllers](#list_of_controllers)

## <a name="template_pkg"></a>![Class: ](https://apidocs.io/img/class.png ".template_pkg") template_pkg

### Get instance

Factory for the ``` TEMPLATE ``` interface can be accessed from the package template_pkg.

```go
template := template_pkg.NewTEMPLATE()
```

### <a name="get_template_list"></a>![Method: ](https://apidocs.io/img/method.png ".template_pkg.GetTemplateList") GetTemplateList

> Get Template List


```go
func (me *TEMPLATE_IMPL) GetTemplateList()(interface{},error)
```

#### Example Usage

```go

var result interface{}
result,_ = template.GetTemplateList()

```


### <a name="create_new_template"></a>![Method: ](https://apidocs.io/img/method.png ".template_pkg.CreateNewTemplate") CreateNewTemplate

> Create New Template


```go
func (me *TEMPLATE_IMPL) CreateNewTemplate(
            messageTemplate string,
            templateName string)(interface{},error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| messageTemplate |  ``` Required ```  | Template text. |
| templateName |  ``` Required ```  | Name of template |


#### Example Usage

```go
messageTemplate := "MessageTemplate"
templateName := "TemplateName"

var result interface{}
result,_ = template.CreateNewTemplate(messageTemplate, templateName)

```


### <a name="update_template"></a>![Method: ](https://apidocs.io/img/method.png ".template_pkg.UpdateTemplate") UpdateTemplate

> Update Template


```go
func (me *TEMPLATE_IMPL) UpdateTemplate(
            messageTemplate string,
            templateName string,
            id int64)(interface{},error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| messageTemplate |  ``` Required ```  | Template text. |
| templateName |  ``` Required ```  | Name of template |
| id |  ``` Required ```  | id of template |


#### Example Usage

```go
messageTemplate := "MessageTemplate"
templateName := "TemplateName"
id,_ := strconv.ParseInt("57", 10, 8)

var result interface{}
result,_ = template.UpdateTemplate(messageTemplate, templateName, id)

```


### <a name="delete_template"></a>![Method: ](https://apidocs.io/img/method.png ".template_pkg.DeleteTemplate") DeleteTemplate

> Delete Template


```go
func (me *TEMPLATE_IMPL) DeleteTemplate(id int64)(interface{},error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| id |  ``` Required ```  | id of template |


#### Example Usage

```go
id,_ := strconv.ParseInt("57", 10, 8)

var result interface{}
result,_ = template.DeleteTemplate(id)

```


[Back to List of Controllers](#list_of_controllers)

## <a name="sms_pkg"></a>![Class: ](https://apidocs.io/img/class.png ".sms_pkg") sms_pkg

### Get instance

Factory for the ``` SMS ``` interface can be accessed from the package sms_pkg.

```go
sMS := sms_pkg.NewSMS()
```

### <a name="get_sent_message_list"></a>![Method: ](https://apidocs.io/img/method.png ".sms_pkg.GetSentMessageList") GetSentMessageList

> Get Sent Message List


```go
func (me *SMS_IMPL) GetSentMessageList(
            enddate *time.Time,
            fromdate *time.Time,
            length int64,
            start int64)(interface{},error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| enddate |  ``` Required ```  | Date format must be in yyyy-mm-dd |
| fromdate |  ``` Required ```  | Date format must be in yyyy-mm-dd |
| length |  ``` Required ```  | Ending index value to fetch the campaign detail. |
| start |  ``` Required ```  | Starting index value to fetch the campaign detail. |


#### Example Usage

```go
enddate := time.Now()
fromdate := time.Now()
length,_ := strconv.ParseInt("57", 10, 8)
start,_ := strconv.ParseInt("57", 10, 8)

var result interface{}
result,_ = sMS.GetSentMessageList(enddate, fromdate, length, start)

```


### <a name="get_sent_message_status"></a>![Method: ](https://apidocs.io/img/method.png ".sms_pkg.GetSentMessageStatus") GetSentMessageStatus

> Get Sent Message Status


```go
func (me *SMS_IMPL) GetSentMessageStatus(messageId int64)(interface{},error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| messageId |  ``` Required ```  | MessageId of message |


#### Example Usage

```go
messageId,_ := strconv.ParseInt("57", 10, 8)

var result interface{}
result,_ = sMS.GetSentMessageStatus(messageId)

```


### <a name="get_create_sms"></a>![Method: ](https://apidocs.io/img/method.png ".sms_pkg.GetCreateSMS") GetCreateSMS

> Create SMS


```go
func (me *SMS_IMPL) GetCreateSMS(
            message string,
            mobileNumber string,
            senderId string,
            coRelator *string,
            isFlash *bool,
            isUnicode *bool,
            linkId *string,
            groupId *string,
            scheduleTime *string,
            serviceId *string)(interface{},error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| message |  ``` Required ```  | text message to send |
| mobileNumber |  ``` Required ```  | Use mobile number as comma sepreated to send message on multiple mobile number e.g. 78461230,78945612 |
| senderId |  ``` Required ```  | Approved Sender Id |
| coRelator |  ``` Optional ```  | Parameter required for using SDP OnDemand Service |
| isFlash |  ``` Optional ```  | Is_Flash is true or false for flash message |
| isUnicode |  ``` Optional ```  | Is_Unicode is true or false for unicode message |
| linkId |  ``` Optional ```  | Parameter required for using SDP OnDemand Service |
| groupId |  ``` Optional ```  | Valid group-id of current user (only for group message otherwise leave empty string) |
| scheduleTime |  ``` Optional ```  | scheduleTime Date in yyyy-MM-dd HH:MM (only for schedule message) |
| serviceId |  ``` Optional ```  | Parameter required for using SDP OnSubscription Service |


#### Example Usage

```go
message := "Message"
mobileNumber := "MobileNumber"
senderId := "SenderId"
coRelator := "CoRelator"
isFlash := false
isUnicode := false
linkId := "LinkId"
groupId := "groupId"
scheduleTime := "scheduleTime"
serviceId := "serviceId"

var result interface{}
result,_ = sMS.GetCreateSMS(message, mobileNumber, senderId, coRelator, isFlash, isUnicode, linkId, groupId, scheduleTime, serviceId)

```


### <a name="create_sms"></a>![Method: ](https://apidocs.io/img/method.png ".sms_pkg.CreateSMS") CreateSMS

> Create SMS


```go
func (me *SMS_IMPL) CreateSMS(
            message string,
            mobileNumber string,
            senderId string,
            coRelator *string,
            isFlash *bool,
            isUnicode *bool,
            linkId *string,
            groupId *string,
            scheduleTime *string,
            serviceId *string)(interface{},error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| message |  ``` Required ```  | text message to send |
| mobileNumber |  ``` Required ```  | Use mobile number as comma sepreated to send message on multiple mobile number e.g. 78461230,78945612 |
| senderId |  ``` Required ```  | Approved Sender Id |
| coRelator |  ``` Optional ```  | Parameter required for using SDP OnDemand Service |
| isFlash |  ``` Optional ```  | Is_Flash is true or false for flash message |
| isUnicode |  ``` Optional ```  | Is_Unicode is true or false for unicode message |
| linkId |  ``` Optional ```  | Parameter required for using SDP OnDemand Service |
| groupId |  ``` Optional ```  | Valid group-id of current user (only for group message otherwise leave empty string) |
| scheduleTime |  ``` Optional ```  | scheduleTime Date in yyyy-MM-dd HH:MM (only for schedule message) |
| serviceId |  ``` Optional ```  | Parameter required for using SDP OnSubscription Service |


#### Example Usage

```go
message := "Message"
mobileNumber := "MobileNumber"
senderId := "SenderId"
coRelator := "CoRelator"
isFlash := false
isUnicode := false
linkId := "LinkId"
groupId := "groupId"
scheduleTime := "scheduleTime"
serviceId := "serviceId"

var result interface{}
result,_ = sMS.CreateSMS(message, mobileNumber, senderId, coRelator, isFlash, isUnicode, linkId, groupId, scheduleTime, serviceId)

```


### <a name="get_create_bulk_sms"></a>![Method: ](https://apidocs.io/img/method.png ".sms_pkg.GetCreateBulkSMS") GetCreateBulkSMS

> Create Bulk SMS


```go
func (me *SMS_IMPL) GetCreateBulkSMS(
            mobileNumberMessage string,
            senderId string,
            coRelator *string,
            isFlash *bool,
            isUnicode *bool,
            linkId *string,
            scheduleTime *time.Time,
            serviceId *string)(interface{},error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| mobileNumberMessage |  ``` Required ```  | Please ensure while submitting the request the message should be passed in encoded format. e.g. 78461230^test~78945612^hello |
| senderId |  ``` Required ```  | Approved Sender Id |
| coRelator |  ``` Optional ```  | Parameter required for using SDP OnDemand Service |
| isFlash |  ``` Optional ```  | Is_Flash is true or false for flash message |
| isUnicode |  ``` Optional ```  | Is_Unicode is true or false for unicode message |
| linkId |  ``` Optional ```  | Parameter required for using SDP OnDemand Service |
| scheduleTime |  ``` Optional ```  | scheduleTime Date in yyyy-MM-dd HH:MM (only for schedule message) |
| serviceId |  ``` Optional ```  | Parameter required for using SDP OnSubscription Service |


#### Example Usage

```go
mobileNumberMessage := "MobileNumber_Message"
senderId := "SenderId"
coRelator := "CoRelator"
isFlash := true
isUnicode := true
linkId := "LinkId"
scheduleTime := time.Now()
serviceId := "serviceId"

var result interface{}
result,_ = sMS.GetCreateBulkSMS(mobileNumberMessage, senderId, coRelator, isFlash, isUnicode, linkId, scheduleTime, serviceId)

```


### <a name="create_bulk_sms"></a>![Method: ](https://apidocs.io/img/method.png ".sms_pkg.CreateBulkSMS") CreateBulkSMS

> Create Bulk SMS


```go
func (me *SMS_IMPL) CreateBulkSMS(
            messageParameters []string,
            senderId string,
            isFlash *bool,
            isUnicode *bool,
            scheduleDateTime *time.Time)(interface{},error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| messageParameters |  ``` Required ```  ``` Collection ```  | TODO: Add a parameter description |
| senderId |  ``` Required ```  | Approved Sender Id |
| isFlash |  ``` Optional ```  | Is_Flash is true or false for flash message |
| isUnicode |  ``` Optional ```  | Is_Unicode is true or false for unicode message |
| scheduleDateTime |  ``` Optional ```  | scheduleTime Date in yyyy-MM-dd HH:MM (only for schedule message) |


#### Example Usage

```go
messageParameters := []string{"MessageParameters"}
senderId := "SenderId"
isFlash := true
isUnicode := true
scheduleDateTime := time.Now()

var result interface{}
result,_ = sMS.CreateBulkSMS(messageParameters, senderId, isFlash, isUnicode, scheduleDateTime)

```


[Back to List of Controllers](#list_of_controllers)

## <a name="group_pkg"></a>![Class: ](https://apidocs.io/img/class.png ".group_pkg") group_pkg

### Get instance

Factory for the ``` GROUP ``` interface can be accessed from the package group_pkg.

```go
gROUP := group_pkg.NewGROUP()
```

### <a name="get_group_list"></a>![Method: ](https://apidocs.io/img/method.png ".group_pkg.GetGroupList") GetGroupList

> Get Group List


```go
func (me *GROUP_IMPL) GetGroupList()(interface{},error)
```

#### Example Usage

```go

var result interface{}
result,_ = gROUP.GetGroupList()

```


### <a name="create_new_group"></a>![Method: ](https://apidocs.io/img/method.png ".group_pkg.CreateNewGroup") CreateNewGroup

> Create New Group


```go
func (me *GROUP_IMPL) CreateNewGroup(groupName string)(interface{},error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| groupName |  ``` Required ```  | Name for new group |


#### Example Usage

```go
groupName := "GroupName"

var result interface{}
result,_ = gROUP.CreateNewGroup(groupName)

```


### <a name="update_group"></a>![Method: ](https://apidocs.io/img/method.png ".group_pkg.UpdateGroup") UpdateGroup

> Update Group


```go
func (me *GROUP_IMPL) UpdateGroup(
            groupName string,
            id string)(interface{},error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| groupName |  ``` Required ```  | Name for new group |
| id |  ``` Required ```  | GroupID |


#### Example Usage

```go
groupName := "GroupName"
id := "id"

var result interface{}
result,_ = gROUP.UpdateGroup(groupName, id)

```


### <a name="create_sub_group_group"></a>![Method: ](https://apidocs.io/img/method.png ".group_pkg.CreateSubGroupGroup") CreateSubGroupGroup

> Create Sub-Group Group


```go
func (me *GROUP_IMPL) CreateSubGroupGroup(
            groupName string,
            id string)(interface{},error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| groupName |  ``` Required ```  | Name for new group |
| id |  ``` Required ```  | TODO: Add a parameter description |


#### Example Usage

```go
groupName := "GroupName"
id := "Id"

var result interface{}
result,_ = gROUP.CreateSubGroupGroup(groupName, id)

```


### <a name="delete_group"></a>![Method: ](https://apidocs.io/img/method.png ".group_pkg.DeleteGroup") DeleteGroup

> Delete Group


```go
func (me *GROUP_IMPL) DeleteGroup(id int64)(interface{},error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| id |  ``` Required ```  | TODO: Add a parameter description |


#### Example Usage

```go
id,_ := strconv.ParseInt("148", 10, 8)

var result interface{}
result,_ = gROUP.DeleteGroup(id)

```


[Back to List of Controllers](#list_of_controllers)

## <a name="campaign_pkg"></a>![Class: ](https://apidocs.io/img/class.png ".campaign_pkg") campaign_pkg

### Get instance

Factory for the ``` CAMPAIGN ``` interface can be accessed from the package campaign_pkg.

```go
campaign := campaign_pkg.NewCAMPAIGN()
```

### <a name="get_campaign_message_status"></a>![Method: ](https://apidocs.io/img/method.png ".campaign_pkg.GetCampaignMessageStatus") GetCampaignMessageStatus

> Get Campaign Message Status


```go
func (me *CAMPAIGN_IMPL) GetCampaignMessageStatus(campaignId int64)(interface{},error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| campaignId |  ``` Required ```  | First user have to call Get Campaigns api for CampaignId |


#### Example Usage

```go
campaignId,_ := strconv.ParseInt("148", 10, 8)

var result interface{}
result,_ = campaign.GetCampaignMessageStatus(campaignId)

```


### <a name="get_campaigns"></a>![Method: ](https://apidocs.io/img/method.png ".campaign_pkg.GetCampaigns") GetCampaigns

> Get Campaigns


```go
func (me *CAMPAIGN_IMPL) GetCampaigns(
            enddate *time.Time,
            fromdate *time.Time,
            length int64,
            start int64)(interface{},error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| enddate |  ``` Required ```  | Date format must be in yyyy-mm-dd |
| fromdate |  ``` Required ```  | Date format must be in yyyy-mm-dd |
| length |  ``` Required ```  | Ending index value to fetch the campaign detail. |
| start |  ``` Required ```  | Starting index value to fetch the campaign detail. |


#### Example Usage

```go
enddate := time.Now()
fromdate := time.Now()
length,_ := strconv.ParseInt("148", 10, 8)
start,_ := strconv.ParseInt("148", 10, 8)

var result interface{}
result,_ = campaign.GetCampaigns(enddate, fromdate, length, start)

```


[Back to List of Controllers](#list_of_controllers)



