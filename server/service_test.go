package main

import (
    "context"
    "testing"

    // 你要测试的是 Connect RPC 形式
    "github.com/bufbuild/connect-go"
    "github.com/yanran666/calculator/server/calculatorv1"
)

func TestCalculatorServiceServer_Calculate(t *testing.T) {
    // 实例化你的服务（嵌入 Connect的 Unimplemented）
    s := &CalculatorServiceServer{}

    // 定义若干测试用例
    tests := []struct {
        name     string
        number1  float64
        number2  float64
        operator string
        want     float64
        wantErr  bool
    }{
        {"Add 1+2", 1, 2, "+", 3, false},
        {"Sub 5-3", 5, 3, "-", 2, false},
        {"Mul 2*3", 2, 3, "*", 6, false},
        {"Div 8/2", 8, 2, "/", 4, false},
        {"Div by zero", 10, 0, "/", 0, true},
        {"Invalid operator", 1, 1, "%", 0, true},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            // 构造 Connect 请求对象
            req := connect.NewRequest(&calculatorv1.CalculatorRequest{
                Number1:  tt.number1,
                Number2:  tt.number2,
                Operator: tt.operator,
            })

            // 调用我们实现的 Calculate
            resp, err := s.Calculate(context.Background(), req)
            if (err != nil) != tt.wantErr {
                t.Fatalf("error = %v, wantErr = %v", err, tt.wantErr)
            }

            // 如果预期不报错，则检查返回结果
            if !tt.wantErr {
                got := resp.Msg.GetResult()
                if got != tt.want {
                    t.Errorf("got = %v, want = %v", got, tt.want)
                }
            }
        })
    }
}
