import { createContext, useContext, useState, useEffect, ReactNode } from 'react';
import { usePathname } from 'next/navigation';
import { MenuItem } from '@/types/menu';
import { useMenuContext } from './menuContext';

interface BreadcrumbItem {
  title: string;
  href?: string;
  icon?: ReactNode;
}

interface BreadcrumbContextProps {
  breadcrumbs: BreadcrumbItem[];
}

const BreadcrumbContext = createContext<BreadcrumbContextProps | undefined>(undefined);

export const BreadcrumbProvider = ({ children }: { children: React.ReactNode }) => {
  const [breadcrumbs, setBreadcrumbs] = useState<BreadcrumbItem[]>([]);
  const pathname = usePathname();
  const { menuData } = useMenuContext();

  useEffect(() => {
    if (pathname) {
      const breadcrumbList: BreadcrumbItem[] = [];
      breadcrumbList.push({ title: "首页", href: "/home" });

      const findMatchingMenuItem = (items: MenuItem[]): boolean => {
        for (const item of items) {
          if (item.path === pathname) {
            breadcrumbList.push({ 
              title: item.title, 
              href: item.path,
              icon: item.icon // 假设 MenuItem 类型中包含 icon 属性
            });
            return true;
          }
          if (item.children && item.children.length > 0) {
            if (findMatchingMenuItem(item.children)) {
              breadcrumbList.unshift({ 
                title: item.title, 
                href: item.path,
                icon: item.icon // 假设 MenuItem 类型中包含 icon 属性
              });
              return true;
            }
          }
        }
        return false;
      };

      findMatchingMenuItem(menuData);

      setBreadcrumbs(breadcrumbList);
    }
  }, [pathname, menuData]); // 添加 menuData 作为依赖项

  return (
    <BreadcrumbContext.Provider value={{ breadcrumbs }}>
      {children}
    </BreadcrumbContext.Provider>
  );
};

export const useBreadcrumb = () => {
  const context = useContext(BreadcrumbContext);
  if (!context) {
    throw new Error('useBreadcrumb must be used within a BreadcrumbProvider');
  }
  return context;
};