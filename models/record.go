package models

type Record struct {
    Header Header
    Key    string //what's the diff b/w string and []byte why use one over the other 
    Value  []byte
}


