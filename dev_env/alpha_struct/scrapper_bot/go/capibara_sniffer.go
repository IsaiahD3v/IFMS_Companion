// Package capibara_sniffer provides the sniffer implementation used by the scraper bot.

package capibara_sniffer;
import ("fmt" "github.com/gocolly/colly/v2 " "os" "encoding/json");

type FileRecordOrigin struct{
		EditalName string `json:"edital_name"`
		EditalLink string `json:"edital_link"`
	}


func main() {
	
	buffer_rotationary_dir := fullpath.Join("/workspaces/IFMS_Companion/dev_env/alpha_struct/buffer_rotationary_dir")

	fmt.Println("capibara siniffer is scavaging")
		//ss
	fmt.Println("debug: ___ ")



	surface_collector := colly.NewCollector(
		MaxDepth(1),
		colly.AllowedAbsoluteUrl("https://selecao.ifms.edu.br/destaques/todos"),

	) 

	depth_collector := colly.NewCollector(
		MaxDepth(2),
		colly.AllowedDomain("ifms.edu.br"),
	)

	

surface_colector.OnHtml("a[href]" func(html_element_link *colly.HTMLElement), {
	link := html_element_link.Attr("href")

	fmt.Printf("debug: %q [ %s\n ]", html_element_link.text, link) //could have used link only text is redundant

	if( string.Contains(link, "edital" ) ) {
		depth_collector.Visit(  ( html_element_link.Request.AbsoluteUrl(link) ) )
	}
		

		
		 depth_collector.OnHtml("a[href]", func(edital_html_element *colly.HTMLElement) {
			edital_link := edital_html_element.Attr("href") }
			if (string.Contains(edital_link, "edital")
			|| string.Contains(edital_link, "Edital") 
			|| string.Contains(edital_link, "EDITAL" ) )	){
				depth_collector.Visit(edital_link) }

				depth_collector.OnResponse(func(download *colly.Response) {

					
					
				  //if wrong invert the order
					edital_name := download.File()
					edital_link := download.Request.AbsoluteUrl.String() 

						

					log_file := json_filler{
						FileName: edital_name,
						FileLink: edital_link,


						download.Save(download.FileName(buffer_rotationary_dir))
					}
						

				}


					
					

					
					})
			} /// review from this point 5 lines up


//it must have a return for sucess and error