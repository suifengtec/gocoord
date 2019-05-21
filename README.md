# gocoord

[![wercker status](https://app.wercker.com/status/78da03fb6f650111e017cc31965a5ce7/s/master "wercker status")](https://app.wercker.com/project/byKey/78da03fb6f650111e017cc31965a5ce7)

在BD09,GCJ02,WGS84坐标系之间相互转换点的坐标。

* 高德地图、腾讯地图以及谷歌中国区地图使用的是GCJ-02坐标系;
* 百度地图使用的是BD-09坐标系;
* 底层接口(HTML5 Geolocation或iOS、安卓API)通过GPS设备获取的坐标使用的是WGS-84坐标系;

使用场景实例: 客户端或者远端的GPS设备上报GPS使用的WGS-84坐标系的点,从数据库中查找附近的点或者在百度地图或高德地图上显示某些点的位置。

## 使用

```bash

go get github.com/suifengtec/gocoord

```

源文件中导入并使用

```go

package main

import (
	"fmt"
	"github.com/suifengtec/gocoord"
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

