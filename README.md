# ipdb-go
IPIP.net officially supported IP database ipdb format parsing library

# Installing
<code>
    go get github.com/ipipdotnet/ipdb-go/ipdb
</code>

# Example
<pre>
<code>
package main

import (
	"github.com/ipipdotnet/ipdb-go/ipdb"
	"fmt"
	"log"
)

func main() {
	db, err := ipdb.New("c:/work/tiantexin/bb/v6/mydata6vipday4.ipdb")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("IPv6 Support:", db.IsIPv6Support()) // Whether IPv6 is supported
	fmt.Println("IPv4 Support:", db.IsIPv4Support()) // Whether IPv4 is supported
	fmt.Println("Languages:", db.Languages())		// Supported language items
	fmt.Println("UTC Time:", db.Build())			// time.Time UTC

	fmt.Println(db.Find("2001:19F0:7000::", "CN"))
    fmt.Println(db.FindMap("2001:19F0:7000::", "CN"))
}
</code>
</pre>