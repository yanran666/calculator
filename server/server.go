package main

import (
    "context"
    "fmt"

    // 必须导入 connect-go，用于 Request/Response
    connect_go "github.com/bufbuild/connect-go"

    // 你的 Protobuf 消息 (CalculatorRequest, CalculatorResponse)
    "github.com/yanran666/calculator/server/calculatorv1"

    // 你的 Connect 服务接口/Handler (UnimplementedCalculatorServiceHandler)
    "github.com/yanran666/calculator/server/calculatorv1connect"
)

type CalculatorServiceServer struct {
    // 注意：嵌入的是 calculatorv1connect 包里的 UnimplementedCalculatorServiceHandler
    calculatorv1connect.UnimplementedCalculatorServiceHandler
}

// 这是 Connect 规定的接口签名：Calculate(ctx, *connect_go.Request[CalculatorRequest]) 
// 返回 (*connect_go.Response[CalculatorResponse], error)
func (s *CalculatorServiceServer) Calculate(
    ctx context.Context,
    req *connect_go.Request[calculatorv1.CalculatorRequest],
) (*connect_go.Response[calculatorv1.CalculatorResponse], error) {

    // Connect 的 Request 里，真正的消息体在 req.Msg
    n1 := req.Msg.GetNumber1()
    n2 := req.Msg.GetNumber2()
    op := req.Msg.GetOperator()

    var result float64
    switch op {
    case "+":
        result = n1 + n2
    case "-":
        result = n1 - n2
    case "*":
        result = n1 * n2
    case "/":
        if n2 == 0 {
            return nil, fmt.Errorf("division by zero")
        }
        result = n1 / n2
    default:
        return nil, fmt.Errorf("unknown operator: %s", op)
    }

    // Connect 要求返回一个 *connect_go.Response[T]，而不是直接 *CalculatorResponse
    return connect_go.NewResponse(&calculatorv1.CalculatorResponse{Result: result}), nil
}
