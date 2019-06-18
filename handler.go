package splunkpersist

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

const opCodeInit = "1"
const opCodeBlock = "2"
const opCodeEnd = "4"
const opCodeAllowStream = "8"

// Request packet coming form splunk
type Request struct {
	OutputMode         string `json:"output_mode"`
	OutputModeExplicit bool   `json:"output_mode_explicit"`
	Server             struct {
		RestURI    string `json:"rest_uri"`
		Hostname   string
		Servername string
		GUID       string `json:"guid"`
	}
	Restmap struct {
		Name string
		Conf struct {
			Driver                string
			Match                 string
			OutputModes           string `json:"output_modes"`
			PassHTTPCookies       string `json:"passHttpCookies"`
			PassHTTPHeaders       string `json:"passHttpHeaders"`
			PassPayload           string
			RequireAuthentication string
			Script                string
			Scripttype            string
		}
	}
	Query      map[string]string
	Connection struct {
		SrcIP         string `json:"src_ip"`
		Ssl           bool
		ListeningPort int `json:"listening_port"`
	}
	Session struct {
		User      string
		Authtoken string
	}
	RestPath string `json:"rest_path"`
	Method   string
}

//Run read packets coming from splunk
func Run(r *Router) {
	reader := bufio.NewReader(os.Stdin)

	for {
		// Get Op Code
		var opbyte byte
		for {
			newbyte, _ := reader.ReadByte()
			if string(newbyte) == "\n" {
				break
			}
			opbyte = newbyte
		}
		//Get Command
		if string(opbyte) == opCodeInit {
			line, _, _ := reader.ReadLine()
			i, _ := strconv.Atoi(string(line))
			command := make([]byte, i)
			_, _ = reader.Read(command)
		}

		//Get Command Args
		_, _, _ = reader.ReadLine()
		line, _, _ := reader.ReadLine()
		i, _ := strconv.Atoi(string(line))
		restArgs := make([]byte, i)
		_, _ = reader.Read(restArgs)

		//Parse the JSON request
		var query Request
		err := json.Unmarshal(restArgs, &query)
		var response Response
		if err != nil {
			response.AddEntry(err.Error())
		} else {
			response = r.routes[query.RestPath](query)
		}
		outString, _ := json.Marshal(response)
		fmt.Println("0")
		fmt.Printf("%v\n", len(string(outString)))
		fmt.Printf("%v", string(outString))
	}

}

func (r *Request) UnmarshalJSON(data []byte) error {
	type Alias Request
	aux := &struct {
		Query [][]string `json:"query"`
		Form  [][]string `json:"form"`
		*Alias
	}{
		Alias: (*Alias)(r),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	r.Query = make(map[string]string)
	for _, q := range aux.Query {
		r.Query[q[0]] = q[1]
	}
	for _, q := range aux.Form {
		r.Query[q[0]] = q[1]
	}
	return nil
}
