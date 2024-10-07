'use client';

import React, { createContext, useState, useContext, useMemo } from 'react';
import { MenuItem } from '@/types/menu';
import { useMenuDataFetch } from '@/hooks/useMenuDataFetch';

interface MenuContextType {
  menuData: MenuItem[];
  permissionsMap: Map<string, string[]>;
}

const MenuContext = createContext<MenuContextType | undefined>(undefined);

export const MenuProvider: React.FC<{ children: React.ReactNode }> = ({ children }) => {
  const { menuData } = useMenuDataFetch();

  const permissionsMap = useMemo(() => {
    const map = new Map<string, string[]>();
    const processMenuItem = (item: MenuItem) => {
      if (item.permissions) {
        map.set(item.path, JSON.parse(item.permissions));
      }
      if (item.children) {
        item.children.forEach(processMenuItem);
      }
    };
    menuData.forEach(processMenuItem);
    return map;
  }, [menuData]);

  const value = useMemo(() => ({ menuData, permissionsMap }), [menuData, permissionsMap]);

  return (
    <MenuContext.Provider value={value}>
      {children}
    </MenuContext.Provider>
  );
};

export const useMenuContext = () => {
  const context = useContext(MenuContext);
  if (context === undefined) {
    throw new Error('useMenuContext must be used within a MenuProvider');
  }
  return context;
};