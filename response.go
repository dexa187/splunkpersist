package splunkpersist

//Perms item
type Perms struct {
	Read  []string `json:"read,omitempty"`
	Write []string `json:"write,omitempty"`
}

//ACL item
type ACL struct {
	Perms `json:"perms,omitempty"`
}

//Fields item
type Fields struct {
	Required []string `json:"required,omitempty"`
	Optional []string `json:"optional,omitempty"`
	Wildcard []string `json:"wildcard,omitempty"`
}

//Entry item
type Entry struct {
	ACL     `json:"acl,omitempty"`
	Fields  `json:"fields,omitempty"`
	Content interface{} `json:"content,omitempty"`
}

//Payload of Response object
type Payload struct {
	Entry    []Entry  `json:"entry,omitempty"`
	Messages []string `json:"messages,omitempty"`
	Status   int      `json:"status,omitempty"`
}

//Response Object
//response := Responce{Payload{ Entry: []Entry{Entry{Content: "test"}}}}
type Response struct {
	Payload `json:"payload"`
}

func (r *Response) AddEntry(s interface{}) {
	r.Entry = append(r.Entry, Entry{Content: s})
}
