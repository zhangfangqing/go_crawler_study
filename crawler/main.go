package main
import (
	"net/http"
	"io/ioutil"
	"fmt"
	"golang.org/x/text/transform"
	"golang.org/x/text/encoding/simplifiedchinese"
)

func main() {
	resp, err := http.Get("http://www.zhenai.com/zhenghun")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		println("http get error.code=", http.StatusOK)
		return
	}

	uft8Reader := transform.NewReader(resp.Body, simplifiedchinese.GBK.NewDecoder())

	all, err := ioutil.ReadAll(uft8Reader)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", all)
}
