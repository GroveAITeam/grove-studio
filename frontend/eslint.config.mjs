import antfu from '@antfu/eslint-config'

export default antfu({
  // @stylistic/eslint-plugin-plus
  stylistic: true,
  // eslint-plugin-format
  formatters: true,
  // vue的eslint配置
  vue: true,
  // 保存删除未引入的代码
  // isInEditor: false,
  // 9x版本 忽略文件这种配置方式 废弃掉eslintignore
  ignores: [
    '*.sh',
    'node_modules',
    '*.md',
    '*.woff',
    '*.ttf',
    '.idea',
    '/public',
    '/docs',
    '.husky',
    '.local',
    '/bin',
    'Dockerfile',
  ],
  lessOpinionated: true,
  rules: {
    'eslint-comments/no-unlimited-disable': 'off',
    'ts/no-use-before-define': 'off',
    'style/no-mixed-operators': 'off',
    'no-console': 'warn',
    'ts/no-unused-expressions': 'off',
    'style/max-statements-per-line': 'off',
    'ts/prefer-namespace-keyword': 'off',
    'antfu/top-level-function': 'off',
    'node/prefer-global/process': 'off',
    'ts/consistent-type-definitions': 'off',
    'ts/ban-ts-comment': 'off',
    'vue/singleline-html-element-content-newline': 'off', // vue 标签强制换行
    // 关闭一些耗时的规则
    'import/no-cycle': 'off',
    'import/no-deprecated': 'off',
    'import/no-named-as-default': 'off',
    'prefer-promise-reject-errors': 'off',
  },
})
