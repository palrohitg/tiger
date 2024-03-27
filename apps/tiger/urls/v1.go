package urls

import (
	"fmt"
	"tigerhall/dtos"
)

func AddV1URLs(groups dtos.RouterGroups) {
	noAuthRg := groups.NoAuth
	fmt.Println(noAuthRg)
}
