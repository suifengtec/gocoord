# goCoord
在BD09,GCJ02,WGS84坐标系之间相互转换点的坐标。

## 使用

```bash

go get https://github.com/suifengtec/goCoord

```

源文件中导入并使用

```go

import(
    "github.com/suifengtec/goCoord"
    "fmt"
)

func main() {

	p := Position{113.739873, 34.356696}
	//BD09ToGCJ02
	// 谷歌地图显示为
	// 113.733355， 34.350604
	fmt.Println("BD09转GCJ02")
	p2 := BD09ToGCJ02(p)
	fmt.Printf("%#v\n", p2)

	//BD09ToWGS84
	fmt.Println("BD09转WGS84")
	p2 = BD09ToWGS84(p)
	fmt.Printf("%#v\n", p2)



}



```
