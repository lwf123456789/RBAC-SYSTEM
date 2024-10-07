import { message } from 'antd';
import { signOut } from "next-auth/react";

const baseUrl = '/api';

export const clientRequest = async (url: string, options: RequestInit = {}) => {
  try {
    const response = await fetch(`${baseUrl}${url}`, options);
    const data = await response.json();

    if (!response.ok) {
      let errorMessage = '未知错误';
      let statusCode = response.status;

      if (typeof data.error === 'string') {
        try {
          const parsedError = JSON.parse(data.error);
          errorMessage = parsedError.message || errorMessage;
          statusCode = parsedError.status || statusCode;
        } catch (e) {
          // 如果解析失败，使用原始的 error 字符串
          errorMessage = data.error;
        }
      } else if (data.error && typeof data.error.message === 'string') {
        errorMessage = data.error.message;
        statusCode = data.error.status || statusCode;
      }

      // 如果 data 直接包含 status，优先使用它
      if (data.status) {
        statusCode = data.status;
      }

      switch (statusCode) {
        case 401:
          message.error('会话已过期,请重新登录');
          await signOut({ redirect: true, callbackUrl: "/" });
          break;
        case 403:
          message.error(`访问被拒绝: ${errorMessage}`);
          break;
        default:
          message.error(`请求失败: ${errorMessage}`);
      }
      throw new Error(errorMessage);
    }

    return data;
  } catch (error) {
    console.error('Request error:', error);
    throw error;
  }
};

export const $clientReq = {
  get: (url: string, options?: RequestInit) => clientRequest(url, { ...options, method: 'GET' }),
  post: (url: string, data: any, options?: RequestInit) => clientRequest(url, {
    ...options,
    method: 'POST',
    headers: { 'Content-Type': 'application/json', ...(options?.headers || {}) },
    body: JSON.stringify(data),
  }),
  put: (url: string, data: any, options?: RequestInit) => clientRequest(url, {
    ...options,
    method: 'PUT',
    headers: { 'Content-Type': 'application/json', ...(options?.headers || {}) },
    body: JSON.stringify(data),
  }),
  delete: (url: string, options?: RequestInit) => clientRequest(url, { ...options, method: 'DELETE' }),
  postFormData: (url: string, formData: FormData, options?: RequestInit) => clientRequest(url, {
    ...options,
    method: 'POST',
    body: formData,
  }),
  putFormData: (url: string, formData: FormData, options?: RequestInit) => clientRequest(url, {
    ...options,
    method: 'PUT',
    body: formData,
  }),
};