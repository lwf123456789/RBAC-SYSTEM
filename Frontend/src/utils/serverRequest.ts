import { getServerSession } from "next-auth/next";
import { authOptions } from "@/app/api/auth/[...nextauth]/route";

const baseUrl = process.env.API_BASE_URL;

export const serverRequest = async (url: string, options: RequestInit = {}) => {
  const session = await getServerSession(authOptions);
  const headers = new Headers(options.headers || {});

  if (session?.accessToken) {
    headers.set('Authorization', `Bearer ${session.accessToken}`);
  }

  const response = await fetch(`${baseUrl}${url}`, {
    ...options,
    headers,
  });

  const data = await response.json();

  if (!response.ok) {
    if (response.status === 401) {
      throw new Error(JSON.stringify({
        status: 401,
        message: "无效或者过期的token"
      }));
    }
    throw new Error(JSON.stringify({
      status: response.status,
      message: data.error || '未知错误'
    }));
  }

  return data;
};

export const $serverReq = {
  get: (url: string, options?: RequestInit) => serverRequest(url, { ...options, method: 'GET' }),
  post: (url: string, data: any, options?: RequestInit) => serverRequest(url, {
    ...options,
    method: 'POST',
    headers: { 'Content-Type': 'application/json', ...(options?.headers || {}) },
    body: JSON.stringify(data),
  }),
  put: (url: string, data: any, options?: RequestInit) => serverRequest(url, {
    ...options,
    method: 'PUT',
    headers: { 'Content-Type': 'application/json', ...(options?.headers || {}) },
    body: JSON.stringify(data),
  }),
  delete: (url: string, options?: RequestInit) => serverRequest(url, { ...options, method: 'DELETE' }),
  postFormData: (url: string, formData: FormData, options?: RequestInit) => serverRequest(url, {
    ...options,
    method: 'POST',
    // headers: { 'Content-Type': 'multipart/form-data' },
    body: formData,
  }),
  putFormData: (url: string, formData: FormData, options?: RequestInit) => serverRequest(url, {
    ...options,
    method: 'PUT',
    // headers: { 'Content-Type': 'multipart/form-data' },
    body: formData,
  }),
};