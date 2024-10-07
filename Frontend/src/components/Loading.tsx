import React from 'react';
import { Spin } from 'antd';

const contentStyle: React.CSSProperties = {
  padding: 50,
  background: 'rgba(0, 0, 0, 0.05)',
  borderRadius: 4,
};

interface LoadingProps {
  isShow: boolean;
}

const Loading: React.FC<LoadingProps> = ({ isShow }) => {
  if (!isShow) {
    return null;
  }

  return (
    <div style={{ display: 'flex', justifyContent: 'center', alignItems: 'center' }}>
      <Spin tip="Loading">
        <div style={contentStyle} />
      </Spin>
    </div>
  );
};

export default Loading;
