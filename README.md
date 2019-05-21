# goCoord
在BD09,GCJ02,WGS84坐标系之间相互转换点的坐标。

## 使用

```bash

go get https://github.com/suifengtec/goCoord

```

源文件中导入并使用

```go

package main

import (
	"fmt"
	gocoord "github.com/suifengtec/goCoord"
)

func main() {
	p := gocoord.Position{Lon: 113.739873, Lat: 34.356696}
	//BD09ToGCJ02
	// 谷歌地图显示为
	// 113.733355， 34.350604
	// 实际输出为gocoord.Position{Lon:113.73337197243862, Lat:34.350630274732744}
	fmt.Println("BD09转GCJ02")
	p2 := gocoord.BD09ToGCJ02(p)
	fmt.Printf("%#v\n", p2)

	//BD09ToWGS84
	// 实际输出为 gocoord.Position{Lon:113.7272281721665, Lat:34.351951705458674}
	fmt.Println("BD09转WGS84")
	p2 = gocoord.BD09ToWGS84(p)
	fmt.Printf("%#v\n", p2)

}


```
