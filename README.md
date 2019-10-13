# Go Custom Log

Create nice custom logs with options

## Preview
![preview log](https://i.ibb.co/5MNQQMt/Screen-Shot-2019-07-26-at-11-17-24.png)

## Usage
```go
package  main

import  (
  log "github.com/igorpollo/go-custom-log
)

func  main()  {

  log.Info("test")
  log.Warning("test")
  log.Error("test")
  log.Success("test")
  
  //Show only if debug mode is true or env var DEBUG=true
  log.DebugMode(true)
  log.Debug("test")
   //Show time in messages
  log.ShowTime(true)

}
``` 

## Running Test

Run `go test -run ''` to see all logs outputs
