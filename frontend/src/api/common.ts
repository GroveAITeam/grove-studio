import { get } from '@/api/request'

// 获取系统配置
export function getSystemConfig() {
  return get('/api/system/config')
}
