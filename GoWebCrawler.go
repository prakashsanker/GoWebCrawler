package main
import (
	"net/http"
	"encoding/json"
    "fmt"
    "io/ioutil"
)


func main() {
    d := make(chan Response)
    c := fanIn(getDuckDuckGo("food"), d) 
    fmt.Println(<- c)
    fmt.Println(<- c)
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
    DuckDuckGoResponse
    GitHubResponse
}

func fanIn(input1 <-chan Response, input2 <-chan Response) <-chan string {
    c := make(chan string)
    go func() {
        for {
            select {
            case s := <-input1:
                c <- s.RelatedTopics[0].Result
            case s := <-input2:
                c <- s.Name
            }
        }
    }()
    return c
}


func getDuckDuckGo(k string) <-chan Response {
    c := make(chan Response)

    go func() {
        resp, err := http.Get("http://api.duckduckgo.com/?q=" + k + "&format=json&pretty=1")
        if err != nil {
            fmt.Println(err)
        }
        var wrapper Response
        var duckDuckParsed DuckDuckGoResponse
        jsonDataFromHttp, jsonErr := ioutil.ReadAll(resp.Body)

        if jsonErr != nil {
            fmt.Println(jsonErr)
        }
        defer resp.Body.Close()


        if err:= json.Unmarshal(jsonDataFromHttp, &duckDuckParsed); err != nil {
            panic(err)
        }

        wrapper.DuckDuckGoResponse = duckDuckParsed
        c <- wrapper
    } ()

    return c
}

func getGitHub(k string) <-chan Response {
    c := make(chan Response)
    go func() {
        resp, err := http.Get("https://api.github.com/users/?q=" + k)
        if err != nil {
            fmt.Println(err)
        }

        var githubParsed GitHubResponse
        var wrapper Response
        jsonDataFromHttp, jsonErr := ioutil.ReadAll(resp.Body)

        if jsonErr != nil {
            fmt.Println(jsonErr)
        }

        defer resp.Body.Close()

        if err:= json.Unmarshal(jsonDataFromHttp, &githubParsed); err != nil {
            fmt.Println(err)
        }
        wrapper.GitHubResponse = githubParsed
        c <- wrapper
    }()
 
    return c
}


