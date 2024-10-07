import { notification } from 'antd';
import { CheckCircleOutlined, WarningOutlined, InfoCircleOutlined } from '@ant-design/icons';

interface NotificationProps {
  type: 'success' | 'error' | 'info' | 'warning';
  message: string;
  description?: string;
  duration?: number;
  placement?: 'top' | 'topLeft' | 'topRight' | 'bottomLeft' | 'bottomRight';
}

const Notification: React.FC<NotificationProps> = ({ type, message, description, duration = 1.5, placement = 'topRight' }) => {
  const iconMap: Record<NotificationProps['type'], React.ReactNode> = {
    success: <CheckCircleOutlined style={{ color: '#52c41a' }} />,
    error: <WarningOutlined style={{ color: '#ff4d4f' }} />,
    info: <InfoCircleOutlined style={{ color: '#1890ff' }} />,
    warning: <WarningOutlined style={{ color: '#faad14' }} />,
  };

  const backgroundColorMap: Record<NotificationProps['type'], string> = {
    success: '#f6ffed',
    error: '#fff1f0',
    info: '#e6f7ff',
    warning: '#fffbe6',
  };

  const borderColorMap: Record<NotificationProps['type'], string> = {
    success: '#b7eb8f',
    error: '#ffa39e',
    info: '#91d5ff',
    warning: '#ffe58f',
  };

  notification.open({
    message,
    description,
    icon: iconMap[type],
    duration,
    placement,
    style: {
      backgroundColor: backgroundColorMap[type],
      border: `1px solid ${borderColorMap[type]}`
    },
  });
  return null
};

export default Notification;
