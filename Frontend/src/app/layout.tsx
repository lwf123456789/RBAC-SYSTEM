"use client";
import React, { useEffect, useState } from "react";
import { BreadcrumbProvider } from '@/contexts/breadcrumbContext';
import { MenuProvider } from '@/contexts/menuContext';
import { LayoutProvider } from '@/contexts/layoutContext';
import '@/styles/globals.css'
import { SessionProvider } from "next-auth/react"
import { ConfigProvider } from 'antd';
import dayjs from 'dayjs';
import 'dayjs/locale/zh-cn';
import customParseFormat from 'dayjs/plugin/customParseFormat';
import utc from 'dayjs/plugin/utc';
import timezone from 'dayjs/plugin/timezone';

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {

  return (
    <html lang="en">
      <body suppressHydrationWarning={true}>
        <ConfigProvider>
          <SessionProvider refetchInterval={10 * 60}>
            <MenuProvider>
              <BreadcrumbProvider>
                <LayoutProvider>
                  <div>
                    {children}
                  </div>
                </LayoutProvider>
              </BreadcrumbProvider>
            </MenuProvider>
          </SessionProvider>
        </ConfigProvider>
      </body>
    </html>
  );
}