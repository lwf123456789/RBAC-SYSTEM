'use client';
import React, { useEffect, useState } from 'react';
import Notification from '@/components/Notification'
import {
    Table,
    Button,
    Modal,
    Form,
    Input,
    Popconfirm,
    message,
    Space,
    TreeSelect,
    Checkbox,
    Dropdown,
    Menu,
    Tag,
    Tooltip,
    Switch,
    InputNumber,
    Drawer,
    Select
} from 'antd';
import {
    PlusOutlined,
    EditOutlined,
    DeleteOutlined,
    SettingOutlined,
    MinusCircleOutlined,
    SlidersOutlined
} from '@ant-design/icons';
import { useLayout } from '@/contexts/layoutContext';
import IconSelector from '@/components/IconSelect/IconSelector';
import { Icon } from '@iconify/react';
import Pagination from '../Pagination';
import { $clientReq } from '@/utils/clientRequest';
import PermissionConfigDrawer from '../PermissionConfigDrawer';

interface MenuItem {
    id: number;
    title: string;
    path: string;
    component_path?: string;
    icon?: string;
    parent_id?: number;
    children?: MenuItem[];
    status?: 0 | 1;
    sort?: number;
    created_at?: string;
    updated_at?: string;
    roles?: any;
}

const MenuManagement: React.FC = () => {
    const { setUseDefaultLayout } = useLayout();
    const [currentPage, setCurrentPage] = useState<number>(1);
    const [pageSize, setPageSize] = useState<number>(10);
    const [totalItems, setTotalItems] = useState(0);
    const [loading, setLoading] = useState(false);
    const [menus, setMenus] = useState<MenuItem[]>([]);
    const [drawerVisible, setDrawerVisible] = useState(false);

    // 关联角色，菜单所有关联角色
    const [visibleRoles, setVisibleRoles] = useState<{ id: number; name: string }[]>([]);
    const [isRoleModalVisible, setIsRoleModalVisible] = useState(false);
    // 打开关联角色弹窗
    const showRoleModal = (roles: { id: number; name: string }[]) => {
        setVisibleRoles(roles);
        setIsRoleModalVisible(true);
    };
    // 关闭关联角色弹窗
    const handleRoleModalClose = () => {
        setIsRoleModalVisible(false);
    };

    const handlePageChange = (page: number) => {
        setCurrentPage(page);
    };

    const handlePageSizeChange = (size: number) => {
        setPageSize(size);
        setCurrentPage(1);
    };

    const fetchMenus = async () => {
        setLoading(true);
        try {
            const url = `/menu/get?page=${currentPage}&pageSize=${pageSize}`;
            const data = await $clientReq.get(url);
            const treeData = buildMenuTree(data.menus);
            setMenus(treeData);
            setTotalItems(data.total);
        } catch (error) {
        } finally {
            setLoading(false);
        }
    }

    const [expandedRowKeys, setExpandedRowKeys] = useState<React.Key[]>([]);

    const onExpand = (expanded: boolean, record: MenuItem) => {
      if (expanded) {
        // 如果是展开操作，只保留当前展开的行
        setExpandedRowKeys([record.id]);
      } else {
        // 如果是折叠操作，清空展开的行
        setExpandedRowKeys([]);
      }
    };
    const buildMenuTree = (menuList: MenuItem[]): MenuItem[] => {
        const menuMap = new Map<number, MenuItem>();
        const rootMenus: MenuItem[] = [];

        menuList.forEach(menu => {
            menuMap.set(menu.id, { ...menu, children: [] });
        });

        menuList.forEach(menu => {
            if (menu.parent_id) {
                const parentMenu = menuMap.get(menu.parent_id);
                if (parentMenu) {
                    parentMenu.children?.push(menuMap.get(menu.id)!);
                }
            } else {
                rootMenus.push(menuMap.get(menu.id)!);
            }
        });

        return rootMenus;
    }

    useEffect(() => {
        fetchMenus();
    }, [currentPage, pageSize]);

    const [editingMenu, setEditingMenu] = useState<MenuItem | null>(null);
    // 用于显示和隐藏自定义列
    const [visibleColumns, setVisibleColumns] = useState({
        path: true,
        component_path: true,
        status: true,
        icon: true
    });
    // 用于显示和隐藏图标选择器
    const [iconSelectorVisible, setIconSelectorVisible] = useState(false);

    const [form] = Form.useForm();

    // 用于在组件卸载时恢复默认布局
    useEffect(() => {
        return () => setUseDefaultLayout(true);
    }, [setUseDefaultLayout]);

    // 添加菜单
    const handleAdd = () => {
        setEditingMenu(null);
        setDrawerVisible(true);
        form.resetFields();
    };

    // 编辑菜单
    const handleEdit = (menu: MenuItem) => {
        setEditingMenu(menu);
        setDrawerVisible(true);
        form.setFieldsValue(menu);
    };

    // 删除菜单
    const handleDelete = async (id: number) => {
        const response = await fetch(`/api/menu/del?id=${id}`, {
            method: 'DELETE',
        });
        const result = await response.json()
        if (!response.ok) {
            message.error(result.error)
        } else {
            Notification({
                type: 'success',
                message: '删除成功!',
                description: '菜单已删除',
                placement: 'topRight'
            });
        }
        fetchMenus();
    };

    /** 权限配置-start */
    const [permissionDrawerVisible, setPermissionDrawerVisible] = useState(false);
    const [currentMenuId, setCurrentMenuId] = useState<number | null>(null);
    const handleConfigPermissions = (menuId: number) => {
        setCurrentMenuId(menuId);
        setPermissionDrawerVisible(true);
    };
    /** 权限配置-end */

    // 保存菜单
    const [saveLoading, setSaveLoading] = useState<boolean>(false)
    const handleSave = async () => {
        setSaveLoading(true);
        try {
            const values = await form.validateFields();
            const newMenu: Partial<MenuItem> = {
                ...(editingMenu?.id ? { id: editingMenu.id } : {}),
                ...values,
            };

            const url = editingMenu ? '/api/menu/update' : '/api/menu/create';
            const method = editingMenu ? 'PUT' : 'POST';

            const response = await fetch(url, {
                method,
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify(newMenu),
            });

            const result = await response.json();
            if (!response.ok) {
                message.error(result.error || '操作失败');
                return;
            }


            await fetchMenus();
            setDrawerVisible(false);
            message.success(editingMenu ? '菜单已更新' : '新菜单已添加');
        } catch (error) {
            if (error instanceof Error) {
                message.error(error.message);
            } else {
                message.error('发生未知错误');
            }
        } finally {
            setSaveLoading(false);
        }
    };

    // 用于显示和隐藏自定义列
    const handleColumnVisibilityChange = (columnKey: keyof typeof visibleColumns) => {
        setVisibleColumns(prev => ({
            ...prev,
            [columnKey]: !prev[columnKey],
        }));
    };

    // 选择图标
    const handleIconSelect = (iconName: string) => {
        form.setFieldsValue({ icon: iconName });
        setIconSelectorVisible(false);
    };

    // 用于配置自定义列
    const columnDefinitions = [
        { title: '页面路径', dataIndex: 'path', key: 'path' },
        { title: '组件路径', dataIndex: 'component_path', key: 'component_path' },
        { title: '状态', dataIndex: 'status', key: 'status' },
        { title: '图标', dataIndex: 'icon', key: 'icon' }
    ];

    // 表格列定义
    const columns: any = [
        {
            title: '菜单标题',
            dataIndex: 'title',
            key: 'title',
            align: 'left',
            render: (text: string, record: MenuItem) => (
                <Space>
                    {record.icon && <Icon icon={record.icon} className="text-lg" />}
                    <span>{text}</span>
                </Space>
            )
        },
        {
            title: '页面路径',
            dataIndex: 'path',
            key: 'path',
            align: 'center',
            visible: visibleColumns.path,
        },
        {
            title: '组件路径',
            dataIndex: 'component_path',
            key: 'component_path',
            align: 'center',
            visible: visibleColumns.component_path,
        },
        {
            title: '排序',
            dataIndex: 'sort',
            key: 'sort',
            align: 'center',
            width: 80,
        },
        {
            title: '状态',
            dataIndex: 'status',
            key: 'status',
            align: 'center',
            width: 100,
            render: (text: number) => (
                <Tag color={text === 1 ? 'green' : 'red'}>
                    {text === 1 ? '启用' : '禁用'}
                </Tag>
            ),
            visible: visibleColumns.status,
        },
        {
            title: '操作',
            key: 'action',
            align: 'center',
            width: 200,
            render: (text: any, record: MenuItem) => (
                <Space size="middle">
                    <Button type="link" icon={<EditOutlined />} onClick={() => handleEdit(record)}>编辑</Button>
                    <Button type="link" icon={<SlidersOutlined />} onClick={() => handleConfigPermissions(record.id)}>配置权限</Button>
                    <Popconfirm title="确定要删除吗?" okText="确定" cancelText="取消" onConfirm={() => handleDelete(record.id)}>
                        <Button type="link" danger icon={<DeleteOutlined />}>删除</Button>
                    </Popconfirm>
                </Space>
            ),
        },
    ].filter(column => column.visible !== false);

    // 自定义列
    const columnMenu = (
        <Menu>
            {columnDefinitions.map((col) => (
                <Menu.Item key={col.dataIndex}>
                    <Checkbox
                        checked={visibleColumns[col.dataIndex as keyof typeof visibleColumns]}
                        onChange={() => handleColumnVisibilityChange(col.dataIndex as keyof typeof visibleColumns)}
                    >
                        {col.title}
                    </Checkbox>
                </Menu.Item>
            ))}
        </Menu>
    );

    return (
        <div className="p-4">
            <Space>
                <Button type="primary" icon={<PlusOutlined />} onClick={handleAdd}>
                    添加菜单
                </Button>
                <Dropdown overlay={columnMenu}>
                    <Button icon={<SettingOutlined />}>自定义列</Button>
                </Dropdown>
            </Space>

            <Table
                style={{ marginTop: 20 }}
                loading={loading}
                columns={columns}
                dataSource={menus}
                rowKey="id"
                pagination={false}
                bordered
                expandable={{
                    expandedRowKeys,
                    onExpand,
                    defaultExpandAllRows: true,
                    indentSize: 20,
                }}
            />
            <Pagination
                currentPage={currentPage}
                pageSize={pageSize}
                totalItems={totalItems}
                onPageChange={handlePageChange}
                onPageSizeChange={handlePageSizeChange}
            />

            <Modal
                title="关联角色"
                open={isRoleModalVisible}
                onCancel={handleRoleModalClose}
                footer={null}
            >
                {visibleRoles.map(role => (
                    <Tag color="blue" key={role.id} style={{ margin: '5px' }}>
                        {role.name}
                    </Tag>
                ))}
            </Modal>

            <Drawer
                title={editingMenu ? '编辑菜单' : '添加菜单'}
                placement="right"
                onClose={() => setDrawerVisible(false)}
                open={drawerVisible}
                width={400}
                footer={
                    <div style={{ textAlign: 'right' }}>
                        <Button onClick={() => setDrawerVisible(false)} style={{ marginRight: 8 }}>
                            取消
                        </Button>
                        <Button onClick={handleSave} type="primary" loading={saveLoading}>
                            保存
                        </Button>
                    </div>
                }
            >
                <Form form={form} layout="vertical">
                    <Form.Item name="title" label="菜单标题" rules={[{ required: true, message: '请输入菜单标题' }]}>
                        <Input />
                    </Form.Item>
                    <Form.Item name="path" label="菜单路径" rules={[{ required: true, message: '请输入路径' }]}>
                        <Input />
                    </Form.Item>
                    <Form.Item name="component_path" label="组件路径"
                        rules={[
                            {
                                validator: (_, value) => {
                                    if (value && value.startsWith('/')) {
                                        return Promise.reject('组件路径不能以 "/" 开头');
                                    }
                                    return Promise.resolve();
                                }
                            }
                        ]}
                    >
                        <Input placeholder="请输入组件路径,不要以 / 开头" />
                    </Form.Item>
                    <Form.Item name="parent_id" label="上级菜单">
                        <TreeSelect
                            showSearch
                            treeDefaultExpandAll
                            treeData={menus.map(menu => ({
                                title: menu.title,
                                value: menu.id,
                                key: menu.id,
                                children: menu.children?.map(child => ({
                                    title: child.title,
                                    value: child.id,
                                    key: child.id,
                                })),
                            }))}
                            placeholder="选择上级菜单"
                            allowClear
                        />
                    </Form.Item>
                    <Form.Item name="icon" label="图标">
                        <Input
                            readOnly
                            value={form.getFieldValue('icon')}
                            addonAfter={
                                <Button onClick={() => setIconSelectorVisible(true)}>
                                    选择图标
                                </Button>
                            }
                            addonBefore={
                                form.getFieldValue('icon') && (
                                    <Icon icon={`${form.getFieldValue('icon')}`} width="24" height="24" />
                                )
                            }
                        />
                    </Form.Item>
                    <Form.Item name="status" label="状态" valuePropName="checked" getValueFromEvent={(checked: boolean) => checked ? 1 : 0}>
                        <Switch checkedChildren="启用" unCheckedChildren="禁用" />
                    </Form.Item>
                    <Form.Item
                        name="sort"
                        label="排序"
                        rules={[
                            { required: true, message: '请输入排序' },
                            { type: 'number', message: '请输入有效的数字' }
                        ]}
                    >
                        <InputNumber
                            min={0}
                            precision={0}
                            style={{ width: '100%' }}
                            placeholder="请输入排序数字"
                        />
                    </Form.Item>
                </Form>
            </Drawer>

            <PermissionConfigDrawer
                visible={permissionDrawerVisible}
                onClose={() => setPermissionDrawerVisible(false)}
                menuId={currentMenuId}
            />

            <IconSelector
                visible={iconSelectorVisible}
                onClose={() => setIconSelectorVisible(false)}
                onSelect={handleIconSelect}
            />
        </div>
    );
};

export default MenuManagement;