// src/types/calculator.ts

export class CalculatorRequest {
  number1: number;
  number2: number;
  operator: string;

  constructor(data?: Partial<CalculatorRequest>) {
    this.number1 = data?.number1 ?? 0;
    this.number2 = data?.number2 ?? 0;
    this.operator = data?.operator ?? "";
  }

  toJsonString(): string {
    // 返回 JSON 字符串形式的消息
    return JSON.stringify({
      number1: this.number1,
      number2: this.number2,
      operator: this.operator,
    });
  }
}

export class CalculatorResponse {
  result: number;

  constructor(data?: Partial<CalculatorResponse>) {
    this.result = data?.result ?? 0;
  }

  toJsonString(): string {
    return JSON.stringify({ result: this.result });
  }
}
