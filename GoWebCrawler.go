package main
import (
	"net/http"
	"encoding/json"
    "fmt"
    "io/ioutil"
)


func main() {
	getDuckDuckGo("food")
    getGitHub("defunkt")

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

func getDuckDuckGo(k string) (map[string]interface{}, error) {
    resp, err := http.Get("http://api.duckduckgo.com/?q=" + k + "&format=json&pretty=1")
    if err != nil {
        return nil, err
    }
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

    // r := make(map[string]interface{})
    // d := json.NewDecoder(resp.Body)
    // if err := d.Decode(&r); err != nil {
    //     return nil, err
    // }
    return nil, nil
}

func getGitHub(k string) (map[string]interface{}, error) {
    resp, err := http.Get("https://api.github.com/users/?q=" + k)
    if err != nil {
        return nil, err
    }

    var githubParsed GitHubResponse
    jsonDataFromHttp, jsonErr := ioutil.ReadAll(resp.Body)

    defer resp.Body.Close()

    if err:= json.Unmarshal(jsonDataFromHttp, &githubParsed); err != nil {
        panic(err)
    }


    return githubParsed, nil
}


