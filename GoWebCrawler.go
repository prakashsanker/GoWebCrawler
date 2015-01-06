package main
import (
	"net/http"
	"encoding/json"
    "fmt"
    "io/ioutil"
)


func main() {
	// getDuckDuckGo("food")
 //    getGitHub("defunkt")
    fanIn(getDuckDuckGo("food"), getGitHub("defunkt"))

}

type DuckDuckGoResponse struct {
    RelatedTopics []struct {
        Result string `json:"Result"`
        FirstUrl string `json:"FirstURL"`
        Text string `json:"Text"`
    } `json:"RelatedTopics"`
}

type GitHubResponse struct {
    Login string `json:"login"`
    Email string `json:"email"`
    Name string `json:"name"`
}

type Response struct {
    GitHub GitHubResponse
    DuckDuckGo DuckDuckGoResponse
}

func fanIn(input1 <- DuckDuckGoResponse, input2 <- GitHubResponse) <-chan string {
    c := make(chan string)

    go func() {
        select {
            case s := <- input1: c <- s
            case s := <- input2: c <- s
        }
    }()
    return c
}


func getDuckDuckGo(k string) <-chan Response {
    resp, err := http.Get("http://api.duckduckgo.com/?q=" + k + "&format=json&pretty=1")
    if err != nil {
        return nil, err
    }
    c := make(chan DuckDuckGoResponse)
    var duckDuckParsed DuckDuckGoResponse
    jsonDataFromHttp, jsonErr := ioutil.ReadAll(resp.Body)

    if jsonErr != nil {
        fmt.Println("Json error!")
    }
    defer resp.Body.Close()


    if err:= json.Unmarshal(jsonDataFromHttp, &duckDuckParsed); err != nil {
        panic(err)
    }

    fmt.Println(duckDuckParsed)
    return c
}

func getGitHub(k string) <-chan Response {
    resp, err := http.Get("https://api.github.com/users/?q=" + k)
    if err != nil {
        return nil, err
    }
    c := make(chan GitHubResponse)

    var githubParsed GitHubResponse
    jsonDataFromHttp, jsonErr := ioutil.ReadAll(resp.Body)

    defer resp.Body.Close()

    if err:= json.Unmarshal(jsonDataFromHttp, &githubParsed); err != nil {
        panic(err)
    }
    return c
}


