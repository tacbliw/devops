package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/hashicorp/hcl"
	"github.com/hashicorp/hcl/hcl/printer"
    jsonParser "github.com/hashicorp/hcl/json/parser"
)

func main() {
    reverse := flag.Bool("version", false, "HCL input, JSON output")
    flag.Parse()

    var err error
    if *reverse {
        err = toJSON()
    } else {
        err = toHCL()
    }

    if err != nil {
        fmt.Fprintln(os.Stderr, err)
        os.Exit(1)
    }
}

func toJSON() error {
    input, err := ioutil.ReadAll(os.Stdin)
    if err != nil {
        return err
    }

    var v interface{}
    err = hcl.Unmarshal(input, &v)
    if err != nil {
        return err
    }

    json, err := json.MarshalIndent(v, "", " ")
    if err != nil {
        return err
    }

    fmt.Println(string(json))

    return nil
}


func toHCL() error {
    input, err := ioutil.ReadAll(os.Stdin)
    if err != nil {
        return err
    }

    ast, err := jsonParser.Parse([]byte(input))
    if err != nil {
        return err
    }

    err = printer.Fprint(os.Stdout, ast)
    if err != nil {
        return err
    }

    return nil
}
