'use client';

import React, { createContext, useContext, useState, ReactNode, useEffect } from 'react';
import DefaultLayout from '@/components/Layouts/DefaultLayout';
import { usePathname } from 'next/navigation';
import { shouldUseDefaultLayout } from '@/utils/layoutConfig';

interface LayoutContextType {
  useDefaultLayout: boolean;
  setUseDefaultLayout: (use: boolean) => void;
}

const LayoutContext = createContext<LayoutContextType | undefined>(undefined);

export const LayoutProvider: React.FC<{ children: ReactNode }> = ({ children }) => {
  const pathname = usePathname();
  const [useDefaultLayout, setUseDefaultLayout] = useState(() => shouldUseDefaultLayout(pathname));

  useEffect(() => {
    setUseDefaultLayout(shouldUseDefaultLayout(pathname));
  }, [pathname]);

  return (
    <LayoutContext.Provider value={{ useDefaultLayout, setUseDefaultLayout }}>
      {useDefaultLayout ? <DefaultLayout>{children}</DefaultLayout> : children}
    </LayoutContext.Provider>
  );
};

export const useLayout = () => {
  const context = useContext(LayoutContext);
  if (context === undefined) {
    throw new Error('useLayout must be used within a LayoutProvider');
  }
  return context;
};