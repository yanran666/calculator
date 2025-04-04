"use client";

import React, { useState } from "react";

const Home = () => {
  const [number1, setNumber1] = useState("0");
  const [number2, setNumber2] = useState("0");
  const [operator, setOperator] = useState("+");
  const [result, setResult] = useState("");

  async function handleCalculate() {
    // 客户端提前检查：如果运算符是除法且第二个数字为 0，则直接提示用户
    if (operator === "/" && parseFloat(number2) === 0) {
      alert("除数不能为零");
      return;
    }
    try {
      const response = await fetch("http://localhost:8080/calculator.v1.CalculatorService/Calculate", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          "Accept": "application/json",
          "Connect-Protocol-Version": "1"
        },
        body: JSON.stringify({
          number1: parseFloat(number1),
          number2: parseFloat(number2),
          operator: operator,
        }),
      });
      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`);
      }
      const data = await response.json();
      console.log("Response data:", data);
      if (data && data.result !== undefined) {
        setResult(data.result.toString());
      } else {
        throw new Error("响应数据中没有 result 字段: " + JSON.stringify(data));
      }
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
      <h1>简单计算器</h1>
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
