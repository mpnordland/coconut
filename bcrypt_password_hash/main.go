package main

import "code.google.com/p/go.crypto/bcrypt"
import "fmt"
import "flag"
import "time"
import "os"

func main() {
    var st, dt time.Time
    workFactor := flag.Int("wf", 12, "bcrypt work factor, default is 12")
    doTime := flag.Bool("t", false, "print out how long it took to hash")
    flag.Usage = func(){ fmt.Fprintf(os.Stderr, "%s [flags] <password to hash>\n flags:\n", os.Args[0]); flag.PrintDefaults()}
    flag.Parse()
    tPass := flag.Arg(0)
    if tPass == "" {
        flag.Usage()
        return
    }
    if *doTime {
        st = time.Now()
    }
    pHash, err := bcrypt.GenerateFromPassword([]byte(tPass), *workFactor)
    if *doTime {
        dt = time.Now()
    }
    if err != nil {
        fmt.Println("Could not hash password because of this", err)
    } else {
        fmt.Println("password hashed in", dt.Sub(st))
        fmt.Println(string(pHash))
    }
}

