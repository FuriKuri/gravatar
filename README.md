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
	g, _ := gravatar.New("<your-mail>", "password")
	images := g.Images()
	for _, image := range images {
		fmt.Printf("Address: %s, Image ID: %s, Image URL: %s \n", 
			image.ID, image.Image.ID, image.Image.URL)
	}
}
```

### Upload new image over URL
```
package main

import "github.com/furikuri/gravatar"

func main() {
	g, _ := gravatar.New("<your-mail>", "password")
	imageId := g.SaveURL("https://<image-url>", 0)
	fmt.Printf("Image ID: %s\n", imageId)
}
```

### Upload new image as data
```
data, _ := ioutil.ReadFile("<image-path>)
imageID := g.SaveData(data, 0)
fmt.Printf("Image ID: %s", imageID)
```

### Use image
```
g.UseImage("<image-id>", []string {"<first-mail>", "<second-mail>"})
```

### Delete image
```
g.DeleteImage("<image-id>")
```

### Address informations
```
result := g.Addresses()
for _, value := range result {
	fmt.Printf("Address: %s, Image ID: %s, Image URL: %s \n",
		value.ID, value.Image.ID, value.Image.URL)
}
```

TBD: more examples

## License
Code released under [the MIT license](https://github.com/FuriKuri/gravatar/blob/master/LICENSE).