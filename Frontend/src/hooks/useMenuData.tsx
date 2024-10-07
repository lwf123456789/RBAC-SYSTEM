'use client';

import { useMenuContext } from '@/contexts/menuContext';
import dynamic from 'next/dynamic';
import { MenuItem } from '@/types/menu';

export const useMenuData = () => {
  const { menuData } = useMenuContext();

  const findMenuItem = (path: string): MenuItem | undefined => {
    const findItem = (items: MenuItem[]): MenuItem | undefined => {
      for (const item of items) {
        if (item.path === path) return item;
        if (item.children) {
          const found = findItem(item.children);
          if (found) return found;
        }
      }
      return undefined;
    };

    return findItem(menuData);
  };

  const getComponentForSlug = (slug: string, subSlug: string) => {
    const path = subSlug ? `/${slug}/${subSlug}` : `/${slug}`;
    const menuItem = findMenuItem(path);
    if (!menuItem || !menuItem.component_path) return null;

    return dynamic(() => import(`@/components/${menuItem.component_path}`).then(mod => mod.default), {
      ssr: false // 禁用服务器端渲染
    });
  };

  return { menuData, getComponentForSlug };
};