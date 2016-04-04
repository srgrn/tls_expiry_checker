// Copyright © 2016 Eran Zimbler <dev@zimbler.net>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"os"
	"time"
)

var days = flag.Int("days", 30, "How many days before certificate is invalid")
var url = flag.String("url", "google.com", "Specify url to check")
var silent = flag.Bool("s", false, "disables all output")

func main() {
	code := 0
	defer func() {
		os.Exit(code)
	}()
	flag.Parse()
	fullUrl := *url + ":443"
	//fmt.Println(fullUrl)
	conn, _ := tls.Dial("tcp", fullUrl, &tls.Config{})

	cert := conn.ConnectionState().PeerCertificates[0]

	end := cert.NotAfter
	diff := end.Sub(time.Now())
	if !*silent {
		fmt.Printf("Checked url: %v - days left: ", *url)
		fmt.Println((diff / (time.Hour * 24)).Nanoseconds())
	}

	if diff < (time.Duration(*days*24) * time.Hour) {
		code = 1
	}
}
