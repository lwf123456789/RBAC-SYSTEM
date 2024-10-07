'use client'
import React, { useEffect, useState } from "react";
import { Form, Input, Button, message } from 'antd';
import { useRouter } from 'next/navigation';
import { UserOutlined, LockOutlined, EyeInvisibleOutlined, EyeOutlined, LoadingOutlined } from '@ant-design/icons';
import Link from "next/link";
import Image from "next/image";
import loginBg from '@/public/bg/login-bg.jpg';
import logoIcon from '@/public/icons/logo.png'
import Notification from '@/components/Notification';
import { useLayout } from "@/contexts/layoutContext";
import { signIn } from 'next-auth/react'
import { withoutAuth } from '@/components/withoutAuth'

const SignIn: React.FC = () => {
    const { setUseDefaultLayout } = useLayout();
    useEffect(() => {
        return () => setUseDefaultLayout(false);
    }, [setUseDefaultLayout]);
    const router = useRouter();
    const [loading, setLoading] = useState(false);
    const [showPassword, setShowPassword] = useState(false);

    const onFinish = async (values: { email: string; password: string; }) => {
        setLoading(true)
        const result = await signIn('credentials', {
            redirect: false,
            email: values.email,
            password: values.password,
        })

        if (result?.error) {
            console.log('result', result);
            if (result.error === 'CredentialsSignin') {
                message.error('邮箱或密码不正确')
            } else {
                message.error(result.error || '登录失败')
            }
        } else {
            Notification({
                type: 'success',
                message: '登录成功!',
                placement: 'top'
            })
            router.push('/system/user')
        }
        setLoading(false)
    }


    // 切换密码显示隐藏
    const togglePasswordVisibility = () => {
        setShowPassword((prev) => !prev);
    };

    return (
        <div className="min-h-screen flex items-center justify-center bg-gray-100" id="notification-root">
            <div className="w-full max-w-4xl rounded-md border border-gray-300 bg-white shadow-lg dark:border-gray-700 dark:bg-gray-800">
                <div className="flex flex-col xl:flex-row">
                    <div className="hidden md:hidden w-full xl:w-1/2 xl:block">
                        <div className="p-10 text-center">
                            <span className="mt-10 inline-block">
                                <Image
                                    width={350}
                                    height={350}
                                    src={loginBg}
                                    alt="Login"
                                    loading="lazy"
                                />
                            </span>
                        </div>
                    </div>

                    <div className="w-full xl:w-1/2 border-t border-gray-300 dark:border-gray-700 xl:border-t-0 xl:border-l">
                        <div className="p-8 sm:p-12 xl:p-16">
                            <div className="flex items-center justify-center flex-col">
                                <h2 className="text-2xl font-bold text-gray-900 dark:text-gray-100 mt-2">
                                    登录 COM ADMIN
                                </h2>
                                <Link className="mb-6 inline-block" href="/">
                                    <Image
                                        width={140}
                                        height={18}
                                        src={logoIcon}
                                        alt="Logo"
                                        loading="lazy"
                                    />
                                </Link>
                            </div>

                            <Form
                                name="login"
                                onFinish={onFinish}
                                autoComplete="off"
                                className="mt-8"
                            >
                                <Form.Item
                                    name="email"
                                    rules={[{ required: true, message: '请输入邮箱!' }]}
                                >
                                    <Input
                                        type="email"
                                        prefix={<UserOutlined className="text-gray-400" />}
                                        placeholder="请输入邮箱..."
                                        className="w-full rounded-lg border border-gray-300 bg-transparent py-3 pl-4 pr-10 text-gray-900 focus:border-blue-500 focus:ring focus:ring-blue-200 dark:border-gray-600 dark:bg-gray-700 dark:text-gray-100"
                                    />
                                </Form.Item>

                                <Form.Item
                                    name="password"
                                    rules={[{ required: true, message: '请输入密码!' }, { min: 6, message: '密码长度不能少于6个字符' }]}
                                >
                                    <Input
                                        prefix={<LockOutlined className="text-gray-400" />}
                                        type={showPassword ? "text" : "password"}
                                        placeholder="请输入密码..."
                                        className="w-full rounded-lg border border-gray-300 bg-transparent py-3 pl-4 pr-10 text-gray-900 focus:border-blue-500 focus:ring focus:ring-blue-200 dark:border-gray-600 dark:bg-gray-700 dark:text-gray-100"
                                        suffix={
                                            <span className="cursor-pointer" onClick={togglePasswordVisibility}>
                                                {showPassword ? <EyeOutlined /> : <EyeInvisibleOutlined />}
                                            </span>
                                        }
                                    />
                                </Form.Item>

                                <Form.Item>
                                    <Button
                                        type="primary"
                                        htmlType="submit"
                                        className={`w-full cursor-pointer rounded-lg border border-blue-500 bg-blue-500 py-3 text-white transition hover:bg-blue-600 flex justify-center items-center ${loading ? 'opacity-50 cursor-not-allowed' : ''}`}
                                        disabled={loading}
                                    >
                                        {loading ? (
                                            <LoadingOutlined className="animate-spin h-5 w-5 mr-3" />
                                        ) : (
                                            '登录'
                                        )}
                                    </Button>
                                </Form.Item>
                            </Form>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    );
};

export default withoutAuth(SignIn)
