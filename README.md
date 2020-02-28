# nicrebooter
how to spell: nic - rebooter inspired from panic + rebooter


## How does it work
```golang
package main

import "github.com/toeydevelopment/nicrebooter"

func worker() {
    for i := 1 i < 10 ; i++ {
        if i == 9 {
            panic("I Need to restart !!!")
        }
    }
}

func main(){
    nicrebooter.Run(worker,2 // repeat every 2 seconds)
}
```