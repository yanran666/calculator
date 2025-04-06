// jest.setup.js
const { TextEncoder, TextDecoder } = require("util");
global.TextEncoder = TextEncoder;
global.TextDecoder = TextDecoder;

// 覆盖 window.alert，防止测试时报错
global.alert = jest.fn();
