'use client'
import React, { useCallback, useEffect, useState } from 'react';
import { Table, Button, Modal, Form, Input, Popconfirm, message, Drawer, Tree } from 'antd';
import Pagination from '../Pagination';
import { EditOutlined, SettingOutlined, DeleteOutlined } from '@ant-design/icons';
import Notification from '@/components/Notification'
import { useLayout } from '@/contexts/layoutContext';
import { $clientReq } from '@/utils/clientRequest';
import { DataNode } from 'antd/es/tree';
import { Icon } from '@iconify/react';

// 扩展 DataNode 类型
interface ExtendedDataNode extends DataNode {
  checked?: boolean;
}

interface Role {
  id: number;
  name: string;
  description: string;
  created_at: string;
  updated_at: string;
  users: any;
}
interface MenuItem {
  id: number;
  title: string;
  icon?: string;
  parent_id: number | null;
  children?: MenuItem[];
}

const RoleManagement: React.FC = () => {
  const { setUseDefaultLayout } = useLayout();
  useEffect(() => {
    return () => setUseDefaultLayout(true);
  }, [setUseDefaultLayout]);
  const [roles, setRoles] = useState<Role[]>([]);
  const [currentPage, setCurrentPage] = useState<number>(1);
  const [pageSize, setPageSize] = useState<number>(5);
  const [totalItems, setTotalItems] = useState(0);
  const [loading, setLoading] = useState(false);
  const [editingRole, setEditingRole] = useState<Role | null>(null);
  const [isDrawerOpen, setIsDrawerOpen] = useState(false);
  const [form] = Form.useForm();

  /** 菜单树模块-start */
  const [menus, setMenus] = useState<MenuItem[]>([])
  const [menuTree, setMenuTree] = useState<DataNode[]>([]);
  const [checkedKeys, setCheckedKeys] = useState<React.Key[]>([]);
  const [assignedMenuIds, setAssignedMenuIds] = useState<number[]>([]);
  const [isMenuDrawerOpen, setIsMenuDrawerOpen] = useState(false);
  const [currentRole, setCurrentRole] = useState<Role | null>(null);
  const convertToTreeData = useCallback((menus: MenuItem[], selectedMenuIds: number[]): DataNode[] => {
    const menuMap = new Map<number, DataNode>();
    const treeData: DataNode[] = [];

    menus.forEach(menu => {
      const node: DataNode = {
        title: menu.title,
        key: menu.id,
        children: [],
        icon: menu.icon ? <Icon icon={menu.icon} /> : undefined,
      };
      menuMap.set(menu.id, node);
    });

    menus.forEach(menu => {
      const node = menuMap.get(menu.id)!;
      if (menu.parent_id === null) {
        treeData.push(node);
      } else {
        const parentNode = menuMap.get(menu.parent_id);
        if (parentNode) {
          parentNode.children?.push(node);
        }
      }
    });

    return treeData;
  }, []);


  const fetchMenus = useCallback(async () => {
    try {
      const menuRes = await $clientReq.get('/role/getMenus?page=1&pageSize=999');
      setMenus(menuRes.menus);
      setMenuTree(convertToTreeData(menuRes.menus, []));
    } catch (error) {
      message.error('获取菜单列表失败');
    }
  }, [convertToTreeData]);

  const fetchAssignedMenus = useCallback(async (roleId: number) => {
    const menuIds = await $clientReq.get(`/role/getAssignedMenus?id=${roleId}`);
    setAssignedMenuIds(menuIds.menu_ids);
  }, []);

  const handleAssignMenus = useCallback((record: Role) => {
    setCurrentRole(record);
    fetchAssignedMenus(record.id);
    setIsMenuDrawerOpen(true);
  }, [fetchAssignedMenus]);

  const handleMenuChange = useCallback((checkedKeys: React.Key[] | { checked: React.Key[]; halfChecked: React.Key[] }, info: any) => {
    const { node, checked } = info;
    let newCheckedKeys: React.Key[] = Array.isArray(checkedKeys) ? checkedKeys : checkedKeys.checked;
  
    if (node.children) {
      const updateChildrenKeys = (parentKey: React.Key, shouldCheck: boolean) => {
        const updateKeys = (items: DataNode[]) => {
          items.forEach(item => {
            if (shouldCheck) {
              newCheckedKeys.push(item.key);
            } else {
              newCheckedKeys = newCheckedKeys.filter(key => key !== item.key);
            }
            if (item.children) {
              updateKeys(item.children);
            }
          });
        };
  
        const parent = menuTree.find(item => item.key === parentKey);
        if (parent && parent.children) {
          updateKeys(parent.children);
        }
      };
  
      updateChildrenKeys(node.key, checked);
    }
  
    setAssignedMenuIds(Array.from(new Set(newCheckedKeys)).map(key => Number(key)));
  }, [menuTree]);

  const handleSaveAssignedMenus = async () => {
    if (!currentRole) return;
    try {
      await $clientReq.post('/role/assignMenus', {
        role_id: currentRole.id,
        menu_ids: assignedMenuIds,
      });
      Notification({
        type: 'success',
        message: '菜单分配成功!',
        description: '角色的菜单权限已更新。',
        placement: 'topRight',
      });
      setIsMenuDrawerOpen(false);
    } catch (error) {
      message.error('菜单分配失败');
    } finally {
      fetchRoles();
    }
  };

  useEffect(() => {
    fetchMenus();
  }, [fetchMenus]);
  /** 菜单树模块-end */

  const handlePageChange = (page: number) => {
    setCurrentPage(page);
  };
  const handlePageSizeChange = (size: number) => {
    setPageSize(size);
    setCurrentPage(1);
  };

  const fetchRoles = useCallback(async () => {
    setLoading(true);
    const url = `/role/get?page=${currentPage}&pageSize=${pageSize}`;
    const data = await $clientReq.get(url);
    if (data) {
      setRoles(data.roles);
      setTotalItems(data.total);
      setLoading(false);
    }
  }, [currentPage, pageSize])

  useEffect(() => {
    fetchRoles();
  }, [fetchRoles]);

  const handleAdd = () => {
    setEditingRole(null);
    setIsDrawerOpen(true);
    form.resetFields();
  };

  const handleEdit = (record: Role) => {
    setEditingRole(record);
    setIsDrawerOpen(true);
    form.setFieldsValue({
      name: record.name,
      description: record.description,
    });
  };

  const [saveLoading, setSaveLoading] = useState<boolean>(false)
  const handleSave = async () => {
    form.validateFields().then(async values => {
      setSaveLoading(true);
      try {
        const url = editingRole ? '/api/role/update' : '/api/role/create';
        const method = editingRole ? 'PUT' : 'POST';

        const response = await fetch(url, {
          method,
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify({
            ...values,
            id: editingRole?.id
          }),
        });

        const result = await response.json();
        if (result.error) {
          message.error(result.error);
          return;
        }
        Notification({
          type: 'success',
          message: editingRole ? '更新成功!' : '添加成功!',
          description: editingRole ? '角色信息已成功更新。' : '角色信息已成功添加。',
          placement: 'topRight',
        });
      } finally {
        setSaveLoading(false);
        setIsDrawerOpen(false);
        fetchRoles();
      }
    });
  };


  const handleDelete = async (id: number) => {
    const response = await fetch(`/api/role/del?id=${id}`, {
      method: 'DELETE',
    });
    const result = await response.json()
    if (!response.ok) {
      message.error(result.error)
    } else {
      Notification({
        type: 'success',
        message: '删除成功!',
        description: '角色已删除',
        placement: 'topRight'
      });
    }
    fetchRoles();
  };

  const columns: any[] = [
    {
      title: '角色名称',
      dataIndex: 'name',
      key: 'name',
      align: 'center'
    },
    {
      title: '角色描述',
      dataIndex: 'description',
      key: 'description',
      align: 'center'
    },
    {
      title: '操作',
      key: 'action',
      align: 'center',
      render: (text: any, record: Role) => (
        <>
          <Button type="link" icon={<EditOutlined />} onClick={() => handleEdit(record)}>编辑</Button>
          <Button type="link" icon={<SettingOutlined />} onClick={() => handleAssignMenus(record)}>分配菜单</Button>
          <Popconfirm
            title="确定删除此数据吗?"
            onConfirm={() => handleDelete(record.id)}
            okText="确定"
            cancelText="取消"
          >
            <Button type="link" icon={<DeleteOutlined />} danger>删除</Button>
          </Popconfirm>
        </>
      ),
    },
  ];

  return (
    <div className='p-4'>
      <Button type="primary" onClick={handleAdd}>添加角色</Button>
      <Table
        columns={columns}
        dataSource={roles}
        style={{ marginTop: 20 }}
        loading={loading}
        rowKey={(record) => record.id.toString()} // 使用 id 作为唯一键
        pagination={false} // 使用自定义的 Pagination
      />
      <Pagination
        currentPage={currentPage}
        pageSize={pageSize}
        totalItems={totalItems}
        onPageChange={handlePageChange}
        onPageSizeChange={handlePageSizeChange}
      />

      <Drawer
        title={editingRole ? '编辑角色' : '添加角色'}
        placement="right"
        open={isDrawerOpen}
        onClose={() => setIsDrawerOpen(false)}
        width={400}
        footer={
          <div style={{ textAlign: 'right' }}>
            <Button onClick={() => setIsDrawerOpen(false)} style={{ marginRight: 8 }}>
              取消
            </Button>
            <Button type="primary" loading={saveLoading} onClick={handleSave}>保存</Button>
          </div>
        }
      >
        <Form form={form} layout="vertical">
          <Form.Item
            name="name"
            label="角色名称"
            rules={[{ required: true, message: '请输入角色名称' }]}
          >
            <Input />
          </Form.Item>
          <Form.Item
            name="description"
            label="角色描述"
            rules={[{ required: true, message: '请输入角色描述' }]}
          >
            <Input />
          </Form.Item>
        </Form>
      </Drawer>

      <Drawer
        title={`为 ${currentRole?.name || ''} 分配菜单`}
        placement="right"
        width={400}
        onClose={() => setIsMenuDrawerOpen(false)}
        open={isMenuDrawerOpen}
        footer={
          <div style={{ textAlign: 'right' }}>
            <Button onClick={() => setIsMenuDrawerOpen(false)} style={{ marginRight: 8 }}>
              取消
            </Button>
            <Button type="primary" onClick={handleSaveAssignedMenus}>
              保存
            </Button>
          </div>
        }
      >
        <Tree
          showLine={true}
          showIcon={true}
          checkable
          checkStrictly
          checkedKeys={assignedMenuIds}
          onCheck={handleMenuChange}
          treeData={menuTree}
        />
      </Drawer>
    </div>
  );
};

export default RoleManagement;
