'use client';
import { Button, Result } from 'antd';
import Link from 'next/link';
import { SmileOutlined } from '@ant-design/icons';
import { useLayout } from '@/contexts/layoutContext';
import { useEffect } from 'react';

const Custom404 = () => {
  const { setUseDefaultLayout } = useLayout();
  useEffect(() => {
    return () => setUseDefaultLayout(false);
  }, [setUseDefaultLayout]);
  return (
    <div
      className="min-h-screen flex flex-col items-center justify-center bg-cover bg-center"
      style={{ backgroundImage: 'url(https://tailwindui.com/img/beams-basic.png)' }}
    >
      <Result
        icon={<SmileOutlined style={{ fontSize: '80px', color: '#1890ff' }} />}
        title={<h1 className="text-6xl font-extrabold text-blue-800 mb-4 animate-bounce">404</h1>}
        subTitle={
          <p className="text-xl text-gray-700 mb-4">
            抱歉，您访问的页面不存在。
          </p>
        }
        extra={
          <div className="space-y-4 space-x-4">
            <Link href="/home">
              <Button
                type="primary"
                size="large"
                className="bg-blue-600 hover:bg-blue-700 transition duration-300 transform hover:scale-105 shadow-lg"
              >
                返回首页
              </Button>
            </Link>
            <Link href="/home">
              <Button
                type="default"
                size="large"
                className="transition duration-300 transform hover:scale-105 shadow-lg"
              >
                联系我们
              </Button>
            </Link>
          </div>
        }
        className="shadow-2xl rounded-xl p-10 bg-white border border-gray-300 transform hover:scale-105 transition duration-500 ease-in-out"
      />
    </div>
  );
};

export default Custom404;
