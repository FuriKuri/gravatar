# gravatar
Simlpe library to use the Gravatar API

## Install
```
go get github.com/furikuri/gravatar
```

## Usage
### Simple example
```
package main

import (
	"fmt"
	"github.com/furikuri/gravatar"
)

func main() {
	g, _ := gravatar.New("your@mail.com", "password")
	images := g.Images()
	for _, image := range images {
		fmt.Printf("Address: %s, Image ID: %s, Image URL: %s \n", 
			image.ID, image.Image.ID, image.Image.URL)
	}
}
```

### Upload new image
```
package main

import "github.com/furikuri/gravatar"

func main() {
	g, _ := gravatar.New("your@mail.com", "password")
	imageId := g.SaveURL("https://image.url", 0)
	fmt.Printf("Image ID: %s\n", imageId)
}
```

TBD: more examples

## License
Code released under [the MIT license](https://github.com/FuriKuri/gravatar/blob/master/LICENSE).