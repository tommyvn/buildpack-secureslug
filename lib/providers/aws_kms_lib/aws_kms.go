package main

import (
    "fmt"
    "os"
    "encoding/json"
    "io/ioutil"

    "github.com/codegangsta/cli"
    "github.com/aws/aws-sdk-go/service/kms"
    "github.com/aws/aws-sdk-go/aws"
)

func main() {
  app := cli.NewApp()

  svc := kms.New(nil)

  app.Commands = []cli.Command{
    {
      Name: "decrypt",
      Action: func(c *cli.Context) {
        dat, err := ioutil.ReadFile(c.Args().First())
        if err != nil {
            panic(err)
        }
        params := &kms.DecryptInput{CiphertextBlob: dat}
        resp, err := svc.Decrypt(params)
        if err != nil {
            panic(err)
        }
        fmt.Printf("%s", resp.Plaintext)
      },
    },
    {
      Name: "generate-data-key",
      Flags: []cli.Flag {cli.StringFlag{Name: "key-id"},},
      Action: func(c *cli.Context) {
        key_id := c.String("key-id")
        resp, err := svc.GenerateDataKey(&kms.GenerateDataKeyInput{KeyID: aws.String(key_id), NumberOfBytes: aws.Int64(128)})
        if err != nil {
            panic(err)
        }
        j, err := json.Marshal(resp)
        if err != nil {
            panic(err)
        }
        fmt.Printf("%s", j)
      },
    },
  }

  app.Run(os.Args)
}
