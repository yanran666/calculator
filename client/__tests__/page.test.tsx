// __tests__/page.test.tsx

// 覆盖 window.alert 以防测试环境中 alert 报错
window.alert = jest.fn();

import React from "react";
import { render, screen, fireEvent, waitFor } from "@testing-library/react";
import Home from "../src/app/page"; // 假设入口组件在 src/app/page.tsx
import "@testing-library/jest-dom";
import { CalculatorResponse } from "../src/gen/calculator_pb";

// 提供 TextEncoder/TextDecoder polyfill（如果未设置）
if (typeof global.TextEncoder === "undefined") {
  const { TextEncoder, TextDecoder } = require("util");
  global.TextEncoder = TextEncoder;
  global.TextDecoder = TextDecoder;
}

beforeAll(() => {
  global.fetch = jest.fn(() =>
    Promise.resolve({
      ok: true,
      status: 200,
      headers: new Headers({
        "content-type": "application/json",
        "connect-protocol-version": "1",
      }),
      json: () => Promise.resolve({ result: 10 }),
    } as unknown as Response)
  ) as jest.Mock;
});


afterAll(() => {
  jest.restoreAllMocks();
});

describe("Calculator UI", () => {
  test("renders page and computes result correctly", async () => {
    render(<Home />);

    const input1 = screen.getByRole("spinbutton", { name: /number1/i });
    const input2 = screen.getByRole("spinbutton", { name: /number2/i });
    const button = screen.getByRole("button", { name: "计算" });

    // 模拟用户输入
    fireEvent.change(input1, { target: { value: "5" } });
    fireEvent.change(input2, { target: { value: "5" } });
    fireEvent.click(button);

    await waitFor(() => {
      expect(screen.getByText("10")).toBeInTheDocument();
    });
  });
});
