import type { AxiosError, AxiosRequestConfig, AxiosResponse, InternalAxiosRequestConfig } from 'axios'
import axios from 'axios'

// 创建axios实例
const service = axios.create({
  baseURL: import.meta.env.VITE_API_BASE_URL || '',
  timeout: 15000, // 请求超时时间
  headers: {
    'Content-Type': 'application/json;charset=utf-8',
  },
})

// 请求拦截器
service.interceptors.request.use(
  (config: InternalAxiosRequestConfig<any>) => {
    // 在请求发送前做一些处理，例如添加token
    const token = localStorage.getItem('token')
    if (token) {
      config.headers.set('Authorization', `Bearer ${token}`)
    }
    return config
  },
  (error: AxiosError) => {
    // 请求错误处理
    console.error('Request error:', error)
    return Promise.reject(error)
  },
)

// 响应拦截器
service.interceptors.response.use(
  (response: AxiosResponse) => {
    const { data } = response

    // 根据自定义错误码处理错误
    if (data.code && data.code !== 200) {
      // 处理业务错误
      console.error('Business error:', data.message || 'Error')
      return Promise.reject(new Error(data.message || 'Error'))
    }

    return data
  },
  (error: AxiosError) => {
    // 处理HTTP错误状态码
    if (error.response) {
      const { status } = error.response

      // 处理常见HTTP错误
      switch (status) {
        case 401:
          // 未授权，可能需要重新登录
          console.error('Unauthorized, please login again')
          // 可以在这里执行登出操作或跳转到登录页
          break
        case 403:
          console.error('Forbidden, no permission')
          break
        case 404:
          console.error('API not found')
          break
        case 500:
          console.error('Server error')
          break
        default:
          console.error(`Request failed with status code ${status}`)
      }
    }
    else if (error.request) {
      // 请求已发出但没有收到响应
      console.error('No response received:', error.request)
    }
    else {
      // 请求配置出错
      console.error('Request config error:', error.message)
    }

    return Promise.reject(error)
  },
)

// 封装GET请求
export function get<T>(url: string, params?: any, config?: AxiosRequestConfig): Promise<T> {
  return service.get(url, { params, ...config })
}

// 封装POST请求
export function post<T>(url: string, data?: any, config?: AxiosRequestConfig): Promise<T> {
  return service.post(url, data, config)
}

// 封装PUT请求
export function put<T>(url: string, data?: any, config?: AxiosRequestConfig): Promise<T> {
  return service.put(url, data, config)
}

// 封装DELETE请求
export function del<T>(url: string, config?: AxiosRequestConfig): Promise<T> {
  return service.delete(url, config)
}

// 导出axios实例
export default service
