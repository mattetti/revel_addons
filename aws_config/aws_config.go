/* 
*  This package is a Revel module that extracts config values fron app.conf
* when the app boots. While this package exports the extracted values, it also
* makes available variables of types defined by the popular launchpad.net/goamz/aws package.
* 
* Using these variables, you can start using S3, EC2 etc.. packages without having to 
* to worry about extracting the AWS config settings.
*
*/
package awsConfig

import (
  "github.com/robfig/revel"
  "launchpad.net/goamz/aws"
  "fmt"
)

var (
	AccessKey  string
  SecretKey  string
  Region aws.Region
  Auth aws.Auth
)

func init() {
  revel.OnAppStart(func() {

    configRequired := func(key string) string {
      value, found := revel.Config.String(key); 
      if !found {
        revel.ERROR.Fatal(fmt.Sprintf("Configuration for %s missing in app.conf.", key))
      }
      return value
    }

    AccessKey = configRequired("aws.access_key")
    SecretKey = configRequired("aws.secret_key")
    Region = aws.Regions[configRequired("aws.region")]
    Auth = aws.Auth{AccessKey: AccessKey, SecretKey: SecretKey}
    revel.TRACE.Println("AwsConfig setup")
  })
}
