package summary

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"urlshortener/shorten"
)

func GetSummary(w http.ResponseWriter,r *http.Request){
    urls:=shorten.Urls
	urls1:=shorten.Urls1

	for _,value:=range urls{

		urls1[value]++
	}

	keys:=make([]string,0,len(urls))

	for key:=range urls1{

		keys=append(keys, key)
	}

	sort.SliceStable(keys,func(i, j int) bool {

		return urls1[keys[i]] > urls1[keys[j]]
	})

	 
	count:=0

	for key,value:=range urls1{

		if count == 3 {

			break
		}
        
		url,err:=url.Parse(key)

		if err != nil{

			log.Fatal(err)
		}


		// fmt.Println(url.Hostname())

		parts:=strings.Split(url.Hostname() , ".")

		// fmt.Println("The value of parts is : ",parts)

		domain:=parts[len(parts)-2]

	    results:=fmt.Sprintf("%v : %d",domain,value)
		fmt.Fprintln(w,results)
		
		count ++;
}

}

