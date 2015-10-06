## GD_Backup_GoLang 
Backup my cloud server data to specific google dirve folder 

## Quick Start
Before you start please finish below steps 
### Step 1 : Enable the Drive API 
1. Use [this wizard](https://console.developers.google.com/start/api?id=drive) to create or select a project in the Google Developers Console and automatically enable the API. Click the **Go to credentials** button to continue.
2. At the top of the page, select the **OAuth consent screen** tab. Select an **Email address**, enter a **Product name** if not already set, and click the **Save** button.
3. Back on the **Credentials** tab, click the **Add credentials** button and select **OAuth 2.0 client ID**.
4. Select the application type **Other** and click the **Create** button.
5. Click OK to dismiss the resulting dialog.
6. Click the  (Download JSON) button to the right of the client ID. Move this file to your working directory and rename it **client_secret.json**.

## Step 2 : Prepare the workspace 
1) Set the **GOPATH** environment variable to your working directory. 
2) Get the Drive API Go client library and OAuth2 package using the following commands:
```
$ go get -u google.golang.org/api/drive/v2
$ go get -u golang.org/x/oauth2/...
```

## Step 3 : Get the Sample 
```
$ git clone https://github.com/luyaotsung/GD_Backup_GoLang.git
```

## Step 4 : Run the Sample 
Build and run the sample using the following command from your working directory:
```
$ go run quickstart.go
```
The first time you run the sample, it will prompt you to authorize access:

Browse to the provided URL in your web browser.

1. If you are not already logged into your Google account, you will be prompted to log in. If you are logged into multiple Google accounts, you will be asked to select one account to use for the authorization.
2. Click the **Accept** button.
3. Copy the code you're given, paste it into the command-line prompt, and press **Enter**.

Authorization information is stored on the file system, so subsequent executions will not prompt for authorization.


## Reference Link 
[https://developers.google.com/drive/web/quickstart/go](https://developers.google.com/drive/web/quickstart/go)
