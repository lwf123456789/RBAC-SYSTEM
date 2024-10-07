import { useMenuContext } from '@/contexts/menuContext';
import { usePathname } from 'next/navigation';
import { useMemo } from 'react';

export const usePermissions = () => {
  const { permissionsMap } = useMenuContext();
  const pathname = usePathname();

  const permissions = useMemo(() => {
    return permissionsMap.get(pathname) || [];
  }, [permissionsMap, pathname]);

  const hasPermission = useMemo(() => {
    return (permissionCode: string) => permissions.includes(permissionCode);
  }, [permissions]);

  return { hasPermission };
};