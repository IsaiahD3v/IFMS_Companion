package main

// importing go libs
import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/gocolly/colly/v2"
)

// defining the structure for saving in the json
type file_log_origin struct {
	EditalName string `json:"edital_name"`
	EditalLink string `json:"edital_link"`
}

// starting the main func
func main() {
	// defining the path for editals download
	buffer_rotationary_dir := "/workspaces/IFMS_Companion/dev_env/alpha_struct/buffer_rotationary_dir"

	var file_log []file_log_origin

	fmt.Println("capibara siniffer is scavaging")
	fmt.Println("debug: ___ ")
	//creating the colectors (they are the actual web crawiling and the agents that do actions such as download)
	surface_collector := colly.NewCollector(
		colly.MaxDepth(1),
		colly.AllowedDomains("selecao.ifms.edu.br"),
	)

	depth_collector := colly.NewCollector(
		colly.MaxDepth(2),
		colly.AllowedDomains("selecao.ifms.edu.br"),
	)

	surface_collector.OnHTML("a[href]", func(html_element_link *colly.HTMLElement) {
		link := html_element_link.Attr("href")
		fmt.Printf("debug: %q [ %s\n ]", html_element_link.Text, link) //could have used link only text is redundant

		if strings.Contains(link, "edital") {
			depth_collector.Visit(
				html_element_link.Request.AbsoluteURL(link))
		}
	})
	/*commanding the colectors to go trough the page on the html tag (a link)
	  and store it to the edital_link var*/
	depth_collector.OnHTML("a[href]", func(edital_html_element *colly.HTMLElement) {
		edital_link := edital_html_element.Attr("href")
		edital_link = edital_html_element.Request.AbsoluteURL(edital_link)
		// send depth_colector to that download page
		if strings.Contains(edital_link, "edital") ||
			strings.Contains(edital_link, "Edital") ||
			strings.Contains(edital_link, "EDITAL") {
			depth_collector.Visit(edital_link)
		}
	})

	depth_collector.OnResponse(func(download *colly.Response) {
		edital_name := download.FileName()
		edital_link := download.Request.URL.String()
		// colector downloads the file
		content_type := download.Headers.Get("Content-Type")
		if !strings.Contains(content_type, "pdf") {
			return
		}

		default_download_dir := filepath.Join(buffer_rotationary_dir, edital_name)
		download.Save(default_download_dir)

		log_file := file_log_origin{
			EditalName: edital_name,
			EditalLink: edital_link,
		}

		file_log = append(file_log, log_file)

	})

	depth_collector.OnScraped(func(download *colly.Response) {
		json_data, download_failure_handler := json.MarshalIndent(file_log, "", "\t")
		if download_failure_handler == nil {
			json_path := filepath.Join(buffer_rotationary_dir, "edital_log.json")
			os.WriteFile(json_path, json_data, 0644)
			fmt.Println(" Debug ___ Json saved to:", json_path)
		} else {
			panic("depth_colector failed to save, download_save_handler ")
		}
	})
	// json handling

	surface_collector.Visit("https://selecao.ifms.edu.br/destaques/todos")

}

// missing error handler
// missing better comments
// good for prototype
// maybe better var names 