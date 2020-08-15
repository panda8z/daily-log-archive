# golang 获取取文件的几种姿势



## 字节读取



```go
package main
import (
  "os"
  "file"
)

func main() {
  file, err := os.Open("mytext.txt")
  if err == nil {
    sum := 0
    buf := make([]byte, 1024)
    for {
      n, err := file.Read(buf)
      sum+=n
      if err == io.EOF {
        break
      }
    }
    fmt.Println("File size is:",sum)
  }
}
```

