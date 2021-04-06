/**
Library: bitbucket.org/rj/xmldom-go
A partial solution based on an incomplete library. The library is missing functions needed to create a DOM piece by piece like other other solutions here. It can however create a DOM by parsing XML. Also, it lacks a function to access the processing instruction, so not surprisingly this is missing from the serialized output.
*/

package main

import (
	"fmt"
	dom "gitlab.com/stone.code/xmldom-go.git"
)

func main() {
	d, err := dom.ParseStringXml(`
<?mxl version="1.0" ?>
<root>
	<element>
		Some text here
	</element>
</root>`)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(d.ToXml()))
}
