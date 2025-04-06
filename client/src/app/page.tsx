"use client";

import React, { useState } from "react";
import { createPromiseClient } from "@connectrpc/connect";
import { createConnectTransport } from "@connectrpc/connect-web";
import { CalculatorService } from "../gen/calculator_connect";
import type { CalculatorResponse } from "../gen/calculator_pb";

const Home = () => {
  const [number1, setNumber1] = useState("0");
  const [number2, setNumber2] = useState("0");
  const [operator, setOperator] = useState("+");
  const [result, setResult] = useState("");

  const transport = createConnectTransport({
    baseUrl: "http://localhost:8080",
  });

  const client = createPromiseClient(CalculatorService, transport);

  async function handleCalculate() {
    try {
      const response = await client.calculate({
        number1: parseFloat(number1),
        number2: parseFloat(number2),
        operator,
      });
      // 假设生成的 Stub 返回结果包装在 response.Msg 中
      setResult(response.result.toString());
    } catch (error: unknown) {
      console.error("Error during calculate:", error);
      let errMsg = "调用失败";
      if (error && typeof error === "object" && "message" in error) {
        errMsg += ": " + (error as { message: string }).message;
      }
      alert(errMsg);
    }
  }

  return (
    <main style={{ margin: "2rem" }}>
      <h1>Connect RPC 计算器</h1>
      <div style={{ display: "flex", gap: "1rem" }}>
        <input
          type="number"
          aria-label="number1"
          value={number1}
          onChange={(e) => setNumber1(e.target.value)}
        />
        <select value={operator} onChange={(e) => setOperator(e.target.value)}>
          <option value="+">+</option>
          <option value="-">-</option>
          <option value="*">*</option>
          <option value="/">/</option>
        </select>
        <input
          type="number"
          aria-label="number2"
          value={number2}
          onChange={(e) => setNumber2(e.target.value)}
        />
        <button onClick={handleCalculate}>计算</button>
      </div>
      <p>
        结果: <strong>{result}</strong>
      </p>
    </main>
  );
};

export default Home;
