package main

import (
	"fmt"
	"net/url"
)

const myUrl string = "https://instagram.fccu31-1.fna.fbcdn.net/v/t51.2885-15/465774232_1492772501529469_915749727888493764_n.jpg?stp=dst-jpg_e15&efg=eyJ2ZW5jb2RlX3RhZyI6ImltYWdlX3VybGdlbi42NDB4MTEzNi5zZHIuZjcxODc4LmRlZmF1bHRfY292ZXJfZnJhbWUifQ&_nc_ht=instagram.fccu31-1.fna.fbcdn.net&_nc_cat=101&_nc_ohc=DasL-p1KsmQQ7kNvgFUxU5a&_nc_gid=b42ba08a679040a9bd9de5fa1fd0b5b1&edm=APCB0h8BAAAA&ccb=7-5&ig_cache_key=MzQ5NjEzMDczNDc4MTkxMDI5MA%3D%3D.3-ccb7-5&oh=00_AYCFAyXxxYLvSvkVajtO_L_b17uABq6kwweCH_BEHaEK6w&oe=6732D15A"

func main() {
	result, err := url.Parse(myUrl)
	// var a *url.URL
	if err != nil {
		panic(err)
	}
	params := result.Query()
	fmt.Println(params)

	partsOfUrl := &url.URL{
		Scheme:   "https",
		Host:     "anindya.online",
		Path:     "/project",
		RawQuery: "userid=12eefnveuvbv1e22nmd347bve",
	}

	newUrl:= partsOfUrl.String()

	fmt.Println(newUrl)

}
