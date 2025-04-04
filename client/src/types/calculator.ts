// src/types/calculator.ts
export interface CalculatorRequest {
  number1: number;
  number2: number;
  operator: string;
}

export interface CalculatorResponse {
  result: number;
}
