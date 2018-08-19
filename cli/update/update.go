package update

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/google/go-github/github"
	"github.com/inconshreveable/go-update"
	"github.com/spf13/cobra"
	"github.com/yesnault/ghue/cli/internal"
)

// used by CI to inject architecture (linux-amd64, etc...) at build time
var architecture string
var urlGitubReleases = "https://github.com/yesnault/ghue/releases"

// Cmd update
var Cmd = &cobra.Command{
	Use:     "update",
	Short:   "Update ghue to the latest release version: ghue update",
	Long:    `ghue update`,
	Aliases: []string{"up"},
	Run: func(cmd *cobra.Command, args []string) {
		doUpdate("", architecture)
	},
}

func getURLArtifactFromGithub(architecture string) string {
	client := github.NewClient(nil)
	release, resp, err := client.Repositories.GetLatestRelease(context.Background(), "yesnault", "ghue")
	if err != nil {
		internal.Exit("Repositories.GetLatestRelease returned error: %v\n%v", err, resp.Body)
		os.Exit(1)
	}

	if len(release.Assets) > 0 {
		for _, asset := range release.Assets {
			if *asset.Name == "ghue-"+architecture {
				return *asset.BrowserDownloadURL
			}
		}
	}

	internal.Exit("Invalid Artifacts on latest release. Please try again in few minutes.\n")
	internal.Exit("If the problem persists, please open an issue on https://github.com/yesnault/ghue/issues\n")
	os.Exit(1)
	return ""
}

func getContentType(resp *http.Response) string {
	for k, v := range resp.Header {
		if k == "Content-Type" && len(v) >= 1 {
			return v[0]
		}
	}
	return ""
}

func doUpdate(baseurl, architecture string) {
	if architecture == "" {
		internal.Exit("You seem to have a custom build of ghue\n")
		internal.Exit("Please download latest release on %s\n", urlGitubReleases)
		os.Exit(1)
	}

	url := getURLArtifactFromGithub(architecture)
	if internal.Verbose {
		fmt.Printf("Url to update ghue: %s\n", url)
	}

	resp, err := http.Get(url)
	if err != nil {
		internal.Exit("Error when downloading ghue: %s\n", err.Error())
		fmt.Printf("Url: %s\n", url)
		os.Exit(1)
	}

	contentType := getContentType(resp)
	if contentType != "application/octet-stream" {
		internal.Exit("Invalid Binary (Content-Type: %s). Please try again or download it manually from %s\n", contentType, urlGitubReleases)
		fmt.Printf("Url: %s\n", url)
		os.Exit(1)
	}

	if resp.StatusCode != 200 {
		internal.Exit("Error http code: %d, url called: %s\n", resp.StatusCode, url)
		os.Exit(1)
	}

	fmt.Printf("Getting latest release from : %s ...\n", url)
	defer resp.Body.Close()
	err = update.Apply(resp.Body, update.Options{})
	if err != nil {
		internal.Exit("Error when updating ghue: %s\n", err.Error())
		internal.Exit("Url: %s\n", url)
		os.Exit(1)
	}
	fmt.Println("Update done.")
}
