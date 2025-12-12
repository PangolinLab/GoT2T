# Got2T - Go Type Conversion Toolkit

Got2T 是一个简洁高效的 Go 语言类型转换工具包，提供了丰富的类型转换接口和实现。

## 🌟 特性

- **全面的类型支持**：支持所有基本数据类型之间的相互转换
- **接口化设计**：清晰的接口定义，易于扩展和维护
- **错误处理**：完善的错误处理机制
- **零依赖**：纯 Go 标准库实现，无第三方依赖

## 📦 安装

```bash
go get github.com/pangolinlab/got2t
```


## 🔧 使用方法

### 基本用法

```go
import "github.com/pangolinlab/got2t"

// 整数转布尔值
converter := &got2t.FromIntImpl{}
result, err := converter.ToBool(1) // 返回 true, nil

// 字符串转整数
strConverter := &got2t.FromStringImpl{}
num, err := strConverter.ToInt("123") // 返回 123, nil

// 布尔值转字符串
boolConverter := &got2t.FromBoolImpl{}
str := boolConverter.ToString(true) // 返回 "true"
```


### 支持的转换类型

- `int`, `int8`, `int16`, `int32`, `int64`
- `uint`, `uint8`, `uint16`, `uint32`, `uint64`
- `float32`, `float64`
- `bool`
- `string`
- `[]byte`

## 🏗️ 架构设计

项目采用接口与实现分离的设计模式：

- [contracts.go](https://github.com/PangolinLab/GoT2T/blob/main/contracts.go) 定义了所有类型转换接口
- [imple.go](https://github.com/PangolinLab/GoT2T/blob/main/imple.go) 提供了具体的实现

## 📝 示例

```go
package main

import (
    "fmt"
    "github.com/pangolinlab/got2t"
)

func main() {
    // 整数转字节切片
    intConverter := &got2t.FromIntImpl{}
    bytes := intConverter.ToBytes(42)
    fmt.Printf("Integer 42 as bytes: %v\n", bytes)
    
    // 字符串转布尔值
    strConverter := &got2t.FromStringImpl{}
    boolVal, _ := strConverter.ToBool("true")
    fmt.Printf("String 'true' as boolean: %v\n", boolVal)
    
    // 布尔值转多种整数类型
    boolConverter := &got2t.FromBoolImpl{}
    fmt.Printf("Boolean true as int: %d\n", boolConverter.ToInt(true))
    fmt.Printf("Boolean true as int64: %d\n", boolConverter.ToInt64(true))
}
```


## ⚠️ 注意事项

- 布尔值与数值类型转换时，`false` 对应 `0`，`true` 对应 `1`
- 字符串转数值类型时，无效格式会返回错误
- 负数转换为无符号类型时会返回错误

## 🤝 贡献

欢迎提交 Issue 和 Pull Request 来改进这个项目！