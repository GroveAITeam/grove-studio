import type { NavigationFailure, RouteLocationRaw, Router } from 'vue-router'
import { router } from './index'

type NavigateOptions = {
  replace?: boolean
  query?: Record<string, string | number | boolean>
}

class Navigator {
  private router: Router

  constructor() {
    this.router = router
  }

  /**
   * 导航到指定路由
   * @param to 路由地址或对象
   * @param options 导航选项
   */
  async navigate(to: RouteLocationRaw, options: NavigateOptions = {}): Promise<void | NavigationFailure | undefined> {
    const { replace = false, query } = options

    // 如果是字符串路径且存在查询参数，构建完整的路由对象
    const finalTo: any = typeof to === 'string' && query
      ? { path: to, query }
      : to

    return replace
      ? this.router.replace(finalTo)
      : this.router.push(finalTo)
  }

  /**
   * 返回上一页
   * @param fallbackPath 如果没有上一页时的回退路径
   */
  async back(fallbackPath: string = '/'): Promise<void> {
    if (window.history.length > 1) {
      this.router.back()
    }
    else {
      await this.navigate(fallbackPath)
    }
  }

  /**
   * 刷新当前页面
   */
  async reload(): Promise<void> {
    const { fullPath } = this.router.currentRoute.value
    await this.router.replace(fullPath)
  }

  /**
   * 获取当前路由信息
   */
  getCurrentRoute() {
    return this.router.currentRoute.value
  }

  /**
   * 获取查询参数
   * @param key 参数名
   */
  getQuery(key?: string) {
    const query = this.getCurrentRoute().query
    return key ? query[key] : query
  }
}

// 导出单例实例
export const navigator = new Navigator()
