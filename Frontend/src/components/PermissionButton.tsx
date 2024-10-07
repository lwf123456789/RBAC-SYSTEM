import React, { memo } from 'react';
import { Button, ButtonProps } from 'antd';
import { usePermissions } from '@/hooks/usePermissions';

interface PermissionButtonProps extends ButtonProps {
  permissionCode: string;
}

const PermissionButton: React.FC<PermissionButtonProps> = memo(({ permissionCode, children, ...props }) => {
  const { hasPermission } = usePermissions();

  if (!hasPermission(permissionCode)) {
    return null;
  }

  return <Button {...props}>{children}</Button>;
});

PermissionButton.displayName = 'PermissionButton';

export default PermissionButton;