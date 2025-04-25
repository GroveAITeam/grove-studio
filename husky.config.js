module.exports = {
  hooks: {
    "pre-commit": "lint-staged"
  },
  config: {
    hooksDir: './frontend/.husky' // 指定自定义钩子路径
  }
};