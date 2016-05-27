package main

import (
	"encoding/base32"
	"flag"
	"fmt"
	"strings"

	"github.com/hgfischer/go-otp"
)

var (
	secret   = flag.String("secret", "", "Secret key")
	isBase32 = flag.Bool("base32", true, "If true, the secret is interpreted as a Base32 string")
	length   = flag.Uint("length", otp.DefaultLength, "OTP length")
	period   = flag.Uint("period", otp.DefaultPeriod, "Period in seconds")
	counter  = flag.Uint64("counter", 0, "Counter")
)

func main() {
	flag.Parse()

	if secret == "" {
		log.Fatal("Must provide a secret.")
	}

	key := *secret
	if !*isBase32 {
		key = base32.StdEncoding.EncodeToString([]byte(*secret))
	}

	key = strings.ToUpper(key)

	totp := &otp.TOTP{
		Secret:         key,
		Length:         uint8(*length),
		Period:         uint8(*period),
		IsBase32Secret: true,
	}
	fmt.Println(totp.Get())
}
