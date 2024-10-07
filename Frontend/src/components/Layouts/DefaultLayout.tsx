"use client";
import React, { useState, ReactNode } from "react";
import Sidebar from "@/components/Sidebar";
import Header from "@/components/Header";
// import menuData from "@/data/menuData";
import { useMenuContext } from "@/contexts/menuContext";
import { Spin } from "antd";
export default function DefaultLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  const { menuData } = useMenuContext();
  const [sidebarOpen, setSidebarOpen] = useState(false);
  const [sidebarExpanded, setSidebarExpanded] = useState(true); // 初始为展开状态
  if (!menuData || menuData.length === 0) {
    return (
      <div className="flex items-center justify-center h-screen bg-gray-100">
        <Spin size="large" />
      </div>
    );
  }
  return (
    <>
      <div className="flex h-screen overflow-hidden">
        {/* 侧边栏 */}
        <Sidebar sidebarOpen={sidebarOpen} setSidebarOpen={setSidebarOpen} menuData={menuData} sidebarExpanded={sidebarExpanded} setSidebarExpanded={setSidebarExpanded} />
        {/* 侧边栏 end */}

        {/* 主体区域 */}
        <div className="relative flex flex-1 flex-col overflow-y-auto overflow-x-hidden">
          {/* 头部 */}
          <Header sidebarOpen={sidebarOpen} setSidebarOpen={setSidebarOpen} sidebarExpanded={sidebarExpanded} setSidebarExpanded={setSidebarExpanded} />
          {/* 头部end */}

          {/* 内容 */}
          <main>
            <div className="mx-auto max-w-screen-2xl p-2 md:p-4 2xl:p-6">
              {children}
            </div>
          </main>
          {/* 内容end */}
        </div>
        {/* <!-- 主体区域end */}
      </div>
    </>
  );
}
