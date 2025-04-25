#!/usr/bin/env node

import { execSync } from 'node:child_process'

try {
  // 使用--no-warnings标志运行eslint命令以抑制实验性警告
  execSync('node --no-warnings node_modules/.bin/eslint . --fix', {
    stdio: 'inherit',
  })

  // 如果命令成功执行，输出成功消息
  process.stdout.write('\n✅ Lint修复成功完成！\n')
  process.exit(0)
}
catch (error) {
  // 如果命令执行失败，输出错误信息
  process.stderr.write('\n❌ Lint修复过程中出现错误\n', error)
  process.exit(1)
}
