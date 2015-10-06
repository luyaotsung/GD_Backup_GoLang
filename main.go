package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/user"
	"path/filepath"
	//"strconv"
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
//const luyaotsung_Lindox_Backup_id string = "0Bz1AlpFYTBpyfm1QUFI4SW8xaDNpVkJiWm05QTZ0S08tSDVPdzF1MzR0S3JGLVZZTXluN2c"

func main() {

	client_secret_file := flag.String("client_secret_file", "", "Full Path of Client Secret File")
	gDrive_folder_id := flag.String("gDrive_folder_id", "", "Folder ID of Google Drive")
	backup_package_file := flag.String("backup_file", "", "Full Paht of backup file")

	flag.Parse()

	ctx := context.Background()

	b, err := ioutil.ReadFile(*client_secret_file)
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

	fmt.Println("File Name -> ", *backup_package_file)
	goFile, err := os.Open(*backup_package_file)
	uploadfiles := new(drive.File)
	p := &drive.ParentReference{Id: *gDrive_folder_id}
	uploadfiles.Parents = []*drive.ParentReference{p}

	currentTime := time.Now()
	uploadfiles.Title = fmt.Sprintf("%04d%02d%02d-%02d%02d.tar.bz2", currentTime.Year(), currentTime.Month(), currentTime.Day(), currentTime.Hour(), currentTime.Minute())

	r, err := srv.Files.Insert(uploadfiles).Media(goFile).Do()

	if err != nil {
		log.Fatalf("Unable to retrieve files.", err)
	}

	fmt.Println("Download URL = ", r.DownloadUrl, " \n ID = ", r.OriginalFilename, r.Id)

}
