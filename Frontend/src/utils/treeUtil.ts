import * as Icons from '@ant-design/icons'; // 导入所有图标
import React, { useEffect, useState } from 'react'
import { DataNode } from 'antd/es/tree';
import { Department } from '@/types/department';

// 动态加载图标
export const getIconComponent = (iconName: string): React.ReactNode | null => {
    // 确保 iconName 是 Icons 中的合法键
    const IconComponent = (Icons as any)[iconName];
    
    // 确保 IconComponent 是有效的 React 组件，并返回
    return IconComponent ? React.createElement(IconComponent) : null;
  };
  
  
  // 递归构建树形结构
export const buildTree = (list: Department[], parentId: number | null = null): DataNode[] => {
    return list
      .filter(item => item.parent_id === parentId)
      .map(item => ({
        title: item.name,
        key: item.id,
        icon: getIconComponent(item.icon), // 动态加载图标
        children: buildTree(list, item.id) // 递归构建子部门
      }));
  };
  