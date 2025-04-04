import React from "react";
import { render } from "@testing-library/react";
import { screen, fireEvent, waitFor } from "@testing-library/dom";

import Home from "../src/app/page"; // 假设你的入口组件在 app/page.tsx
import "@testing-library/jest-dom";


// mock 全局 fetch 来模拟后端响应
beforeAll(() => {
  global.fetch = jest.fn(() =>
    Promise.resolve({
      ok: true,
      json: () => Promise.resolve({ result: 10 }),
    } as Response)
  ) as jest.Mock;
});

afterAll(() => {
  jest.restoreAllMocks();
});

describe("Calculator UI", () => {
  test("renders page and computes result correctly", async () => {
    render(<Home />);

    // 模拟用户输入：选择输入框和按钮
    const input1 = screen.getByRole("spinbutton", { name: /number1/i }) || screen.getByDisplayValue("0");
    const input2 = screen.getByRole("spinbutton", { name: /number2/i }) || screen.getByDisplayValue("0");
    const button = screen.getByRole("button", { name: "计算" });

    // 模拟输入值
    fireEvent.change(input1, { target: { value: "5" } });
    fireEvent.change(input2, { target: { value: "5" } });

    // 点击计算按钮
    fireEvent.click(button);

    // 等待异步结果更新，并断言显示正确结果
    await waitFor(() => {
      expect(screen.getByText("10")).toBeInTheDocument();
    });
  });
});
