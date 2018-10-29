package commonlib
import (
  "log"
  "fmt"
  )

func FailOnError(err error, msgerr string, msgsuc string) {
        if err != nil {
                log.Fatalf("%s: %s",msgerr,err);
                
        } else {
                fmt.Printf("%s\n",msgsuc)
        }

}
