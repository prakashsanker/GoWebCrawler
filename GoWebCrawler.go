package main
import (
	"net/http"
	"encoding/json"
)


func main() {
	getDuckDuckGo("food")
    getGitHub("prakash")

}

func getDuckDuckGo(k string) (map[string]interface{}, error) {
    resp, err := http.Get("http://api.duckduckgo.com/?q=" + k + "&format=json&pretty=1")
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    r := make(map[string]interface{})
    d := json.NewDecoder(resp.Body)
    if err := d.Decode(&r); err != nil {
        return nil, err
    }
    return r, nil
}

func getGitHub(k string) (map[string]interface{}, error) {
    resp, err := http.Get("https://api.github.com/users/?q=" + k)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    r := make(map[string]interface{})
    d := json.NewDecoder(resp.Body)
    if err := d.Decode(&r); err != nil {
        return nil, err
    }
    return r, nil
}


