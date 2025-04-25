/** @type {import('tailwindcss').Config} */
export default {
  content: [
    "./index.html",
    "./src/**/*.{vue,js,ts,jsx,tsx}",
  ],
  theme: {
    extend: {},
  },
  plugins: [
    require('@tailwindcss/forms'),
    require('daisyui'),
  ],
  daisyui: {
    themes: [
      {
        grovechat: {
          "primary": "#0EA5E9",         // 浅蓝色
          "secondary": "#6B7280",       // 灰色
          "accent": "#10B981",          // 绿色，用于特殊强调
          "neutral": "#374151",         // 深灰色
          "base-100": "#FFFFFF",        // 白色背景
          "base-200": "#F3F4F6",        // 浅灰色背景
          "base-300": "#E5E7EB",        // 稍深灰色背景
          "info": "#3B82F6",            // 信息蓝
          "success": "#10B981",         // 成功绿
          "warning": "#F59E0B",         // 警告黄
          "error": "#EF4444",           // 错误红
        },
        grovedark: {
          "primary": "#0EA5E9",         // 浅蓝色
          "secondary": "#9CA3AF",       // 灰色
          "accent": "#10B981",          // 绿色
          "neutral": "#1F2937",         // 暗灰色
          "base-100": "#202123",        // 深灰背景
          "base-200": "#343541",        // 稍浅背景色
          "base-300": "#444654",        // 更浅背景色
          "info": "#3B82F6",            // 信息蓝
          "success": "#10B981",         // 成功绿
          "warning": "#F59E0B",         // 警告黄
          "error": "#EF4444",           // 错误红
        }
      },
      "light",
      "dark"
    ],
    darkTheme: "grovedark",
  },
}
