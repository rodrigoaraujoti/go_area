package htmlutil

import (
	"io/ioutil"
	"net/http"
	"regexp"
)

//Titulo lista os titulos das paginas informadas
func Titulo(urls ...string) <-chan string {
	c := make(chan string)
	for _, url := range urls {
		go func(url string) { //funcao anonima
			resp, _ := http.Get(url)
			html, _ := ioutil.ReadAll(resp.Body)

			r, _ := regexp.Compile("<title>(.*?)<\\/title>")
			// fmt.Println(reflect.TypeOf(r)) ~ obs: o vscode nao autocompleta este tipo: *regexp.Regexp
			c <- r.FindStringSubmatch(string(html))[1] //--> [0] toda a regex, [1] soh o coringa: (.*?)

		}(url) // executando funcao anonima
	}
	return c
}

func main() {

}
