// 代码生成时间: 2025-10-04 22:47:46
package main

import (
    "fmt"
    "math"
)

// NumericalIntegration 结构体，用于数值积分计算
type NumericalIntegration struct {
    // 定义变量
    Function    func(float64) float64 // 被积函数
    LowerBound  float64               // 积分下限
    UpperBound  float64               // 积分上限
}

// NewNumericalIntegration 创建一个新的数值积分计算器实例
func NewNumericalIntegration(function func(float64) float64, lowerBound, upperBound float64) *NumericalIntegration {
    return &NumericalIntegration{
        Function:    function,
        LowerBound:  lowerBound,
        UpperBound:  upperBound,
    }
}

// Integrate 使用梯形法则计算数值积分
func (ni *NumericalIntegration) Integrate() (float64, error) {
    // 定义梯形法则参数
    a := ni.LowerBound
    b := ni.UpperBound
    n := 1000 // 梯形分割的数量
    h := (b - a) / float64(n) // 每个梯形的宽度

    // 初始化积分值
    integral := 0.5 * (ni.Function(a) + ni.Function(b))

    // 计算中间梯形的面积并累加
    for i := 1; i < n; i++ {
        x := a + float64(i)*h
        integral += ni.Function(x)
    }

    // 计算最终积分结果
    integral *= h

    return integral, nil
}

func main() {
    // 定义被积函数，例如：x^2
    function := func(x float64) float64 {
        return x * x
    }

    // 创建数值积分计算器实例
    integrator := NewNumericalIntegration(function, 0, 1)

    // 计算积分
    result, err := integrator.Integrate()
    if err != nil {
        fmt.Println("Error: ", err)
    } else {
        fmt.Printf("The integral of x^2 from 0 to 1 is: %f
", result)
    }
}