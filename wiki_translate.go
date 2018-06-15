package main
 
import (
	"net/http"
	"log"
	"io/ioutil"
	"fmt"
	"strings"
	"os"
)
var (
	a int = 1
	responseString2 string)

func main() {
	// Запрос данных с сайта
	response, err := http.Get("https://ru.wikipedia.org/wiki/"+os.Args[1])
	// Запись в переменную
    	responseData,err := ioutil.ReadAll(response.Body)
if err != nil {
    log.Fatal(err)
}
// преобразование в строку
responseString := string(responseData)
// поиск позиции первого вхождения тега
a = strings.Index(responseString, "interlanguage-link interwiki")
// перебор в цикле других вхождений
for (a<len(responseString)) {
	responseString2 = responseString[a:len(responseString)]
	//fmt.Println(responseString2[0:100])
	fmt.Println(responseString2[strings.Index(responseString2, "title")+7:strings.Index(responseString2, "lang=")-2]+" ("+responseString2[strings.Index(responseString2, "lang=")+6:strings.Index(responseString2, "hreflang")-2]+")")//+"		"+responseString2[strings.Index(responseString2, "interlanguage-link-target")+len("interlanguage-link-target"):strings.Index(responseString2, "interlanguage-link-target")+len("interlanguage-link-target")+30])
	responseString2 = strings.Replace(responseString2, "interlanguage-link interwiki", "", 1)
	//fmt.Println(responseString2[0:100])
	if strings.Index(responseString2[10:len(responseString2)], "interlanguage-link interwiki") == -1	{
		// если не найдено, выходим
		return
	}
	a = a+strings.Index(responseString2[10:len(responseString2)], "interlanguage-link interwiki")
	}
}
