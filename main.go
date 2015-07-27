
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/user"
	"path/filepath"
	"time"

	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/drive/v2"
)

// getClient uses a Context and Config to retrieve a Token
// then generate a Client. It returns the generated Client.
func getClient(ctx context.Context, config *oauth2.Config) *http.Client {
	cacheFile, err := tokenCacheFile()
	if err != nil {
		log.Fatalf("Unable to get path to cached credential file. %v", err)
	}
	tok, err := tokenFromFile(cacheFile)
	if err != nil {
		tok = getTokenFromWeb(config)
		saveToken(cacheFile, tok)
	}

	return config.Client(ctx, tok)
}

// getTokenFromWeb uses Config to request a Token.
// It returns the retrieved Token.
func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)

	var code string
	if _, err := fmt.Scan(&code); err != nil {
		log.Fatalf("Unable to read authorization code %v", err)
	}

	tok, err := config.Exchange(oauth2.NoContext, code)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web %v", err)
	}
	return tok
}

// tokenCacheFile generates credential file path/filename.
// It returns the generated credential path/filename.
func tokenCacheFile() (string, error) {
	usr, err := user.Current()
	if err != nil {
		return "", err
	}
	tokenCacheDir := filepath.Join(usr.HomeDir, ".credentials")
	os.MkdirAll(tokenCacheDir, 0700)
	return filepath.Join(tokenCacheDir,
		url.QueryEscape("drive-api-quickstart.json")), err
}

// tokenFromFile retrieves a Token from a given file path.
// It returns the retrieved Token and any read error encountered.
func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	t := &oauth2.Token{}

	err = json.NewDecoder(f).Decode(t)
	defer f.Close()
	return t, err
}

// saveToken uses a file path to create a file and store the
// token in it.
func saveToken(file string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", file)
	f, err := os.Create(file)
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
}

type My_ID struct {
	Folder string
	File   string
}

// My Folder ID
// Link ==> https://drive.google.com/drive/folders/0Bz1AlpFYTBpyfm1QUFI4SW8xaDNpVkJiWm05QTZ0S08tSDVPdzF1MzR0S3JGLVZZTXluN2c
const luyaotsung_Lindox_Backup_id string = "0Bz1AlpFYTBpyfm1QUFI4SW8xaDNpVkJiWm05QTZ0S08tSDVPdzF1MzR0S3JGLVZZTXluN2c"

func main() {
	ctx := context.Background()

	b, err := ioutil.ReadFile("/home/Backup/Source/client_secret_37504509164-suvnccq460nhbjsah887lkat9ddud873.apps.googleusercontent.com.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	config, err := google.ConfigFromJSON(b, drive.DriveScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}

	client := getClient(ctx, config)

	srv, err := drive.New(client)
	if err != nil {
		log.Fatalf("Unable to create Drive service: %v", err)
	}

	//r, err := srv.Files.List().MaxResults(10).Do()
	//if err != nil {
	//	log.Fatalf("Unable to retrieve files.", err)
	//}
	//fmt.Println("Files:")
	//if len(r.Items) > 0 {
	//	for _, i := range r.Items {
	//		fmt.Printf("%s (%s)\n", i.Title, i.Id)
	//	}
	//} else {
	//	fmt.Print("No files found.")
	//}

	//filename := os.Args[1]
	filename := "/home/Backup/Backup.tar.bz2"
	fmt.Println("File Name -> ", filename)
	goFile, err := os.Open(filename)
	uploadfiles := new(drive.File)
	p := &drive.ParentReference{Id: luyaotsung_Lindox_Backup_id}
	uploadfiles.Parents = []*drive.ParentReference{p}

	currentTime := time.Now()
	abc := fmt.Sprintf("%04d%02d%02d-%02d%02d.tar.bz2", currentTime.Year(), currentTime.Month(), currentTime.Day(), currentTime.Hour(), currentTime.Minute())

	uploadfiles.Title = abc

	r, err := srv.Files.Insert(uploadfiles).Media(goFile).Do()

	//srv.Files.Delete("0Bz1AlpFYTBpya2hhVTJldlEtOFk").Do()

	if err != nil {
		log.Fatalf("Unable to retrieve files.", err)
	}

	fmt.Println("Download URL = ", r.DownloadUrl, " \n ID = ", r.OriginalFilename, r.Id)

	//googleDrive_ID := new(My_ID)
	//googleDrive_ID.Folder = luyaotsung_Lindox_Backup_id
	//googleDrive_ID.File = r.Id

	//usr, _ := user.Current()

	//CacheDir := filepath.Join(usr.HomeDir, ".credentials")
	//os.MkdirAll(CacheDir, 0700)
	//saveID(filepath.Join(CacheDir, url.QueryEscape("id.json")), googleDrive_ID)

}

