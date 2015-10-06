## GD_Backup_GoLang 
Backup my cloud server data to specific google dirve folder 

## Precondition 
Before you start please finish below steps [Reference Link](https://developers.google.com/drive/web/quickstart/go)
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

## Step 3 : Get the sample 
```
git clone https://github.com/luyaotsung/GD_Backup_GoLang.git
```


