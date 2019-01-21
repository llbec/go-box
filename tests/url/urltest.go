package main

import (
	"fmt"
	"net/url"
	"os"
)

func main() {
	//s := "https://golang.org/pkg/net/url/"
	if len(os.Args) != 2 {
		fmt.Println("Please specify the url!")
		return
	}
	s := os.Args[1]
	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}
	fmt.Println("\tParse: ", u)
	{
		fmt.Println("\t\tScheme: ", u.Scheme)
		fmt.Println("\t\tOpaque: ", u.Opaque)
		fmt.Println("\t\tUser: ", u.User)
		fmt.Println("\t\tHost: ", u.Host)
		fmt.Println("\t\tPath: ", u.Path)
		fmt.Println("\t\tRawOath: ", u.RawPath)
		fmt.Println("\t\tForceQuery: ", u.ForceQuery)
		fmt.Println("\t\tRawQuery: ", u.RawQuery)
		fmt.Println("\t\tFragment: ", u.Fragment)
		fmt.Println("\t\tEscapedPath: ", u.EscapedPath())
		fmt.Println("\t\tHostname: ", u.Hostname())
		fmt.Println("\t\tIsAbs: ", u.IsAbs())
		b, err := u.MarshalBinary()
		if err != nil {
			panic(err)
		}
		fmt.Println("\t\tMarshalBinary: ", b)
		fmt.Println("\t\tPort: ", u.Port())
		fmt.Println("\t\tQuery: ", u.Query())
		fmt.Println("\t\tRequestURI: ", u.RequestURI())
		fmt.Println("\t\tResolveReference: ", u.ResolveReference(u))
		u.Opaque = "test"
		fmt.Println("\t\tString: ", u.String(), "\t", u)
		err = u.UnmarshalBinary(b)
		if err != nil {
			panic(err)
		}
		fmt.Println("\t\tUnmarshalBinary: ", u, "\t", u.String())
	}

	fmt.Println("\tPathEscape: ", url.PathEscape(s))
	r, err := url.PathUnescape(s)
	if err != nil {
		panic(err)
	}
	fmt.Println("\tPathUnescape: ", r)
	fmt.Println("\tQueryEscape: ", url.QueryEscape(s))
	r, err = url.QueryUnescape(s)
	if err != nil {
		panic(err)
	}
	fmt.Println("\tQueryUnescape: ", r)
}
