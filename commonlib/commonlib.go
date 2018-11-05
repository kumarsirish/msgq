
//commonlib will contain list of functions which can used by producer/consumer across various message brokers
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
