'use client';

import { useState, useEffect } from 'react';
import { $clientReq } from '@/utils/clientRequest';
import { usePathname } from 'next/navigation';

export const useMenuDataFetch = () => {
  const [menuData, setMenuData] = useState([]);
  const pathname = usePathname();

  const fetchMenuData = async () => {
    if (pathname === '/') return;
    try {
      const data = await $clientReq.get('/menu/getMenusByUser');
      setMenuData(data);
    } catch (error: any) {
    }
  };

  useEffect(() => {
    fetchMenuData();
  }, [pathname]);

  return { menuData };
};