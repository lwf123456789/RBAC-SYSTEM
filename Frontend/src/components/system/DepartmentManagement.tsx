'use client';
import React, { useEffect, useState, useCallback } from 'react';
import { Table, Button, Form, Input, Popconfirm, message, Space, Drawer, Select, InputNumber, Tree, Row, Col, TreeSelect, Dropdown } from 'antd';
import { PlusOutlined, EditOutlined, DeleteOutlined, SearchOutlined, SettingOutlined, ReloadOutlined, ApartmentOutlined, TeamOutlined, UserOutlined, BankOutlined, GroupOutlined, SolutionOutlined } from '@ant-design/icons';
import Notification from '@/components/Notification'
import Pagination from '@/components/Pagination'
import { useLayout } from '@/contexts/layoutContext';
import { $clientReq } from '@/utils/clientRequest';
import { DataNode } from 'antd/es/tree';

interface Department {
    id: number;
    name: string;
    parent_id: number | null;
    icon: string;
    created_at: string;
    updated_at: string;
}

const DepartmentManagement: React.FC = () => {
    const { setUseDefaultLayout } = useLayout();
    const [departments, setDepartments] = useState<Department[]>([]);
    const [currentPage, setCurrentPage] = useState<number>(1);
    const [pageSize, setPageSize] = useState<number>(10);
    const [totalItems, setTotalItems] = useState(0);
    const [loading, setLoading] = useState(false);
    const [isDrawerOpen, setIsDrawerOpen] = useState(false);
    const [editingDepartment, setEditingDepartment] = useState<Department | null>(null);
    const [form] = Form.useForm();
    const [searchName, setSearchName] = useState('');
    const [treeData, setTreeData] = useState<DataNode[]>([]);

    useEffect(() => {
        return () => setUseDefaultLayout(true);
    }, [setUseDefaultLayout]);

    const handlePageChange = (page: number) => {
        setCurrentPage(page);
    };

    const handlePageSizeChange = (size: number) => {
        setPageSize(size);
        setCurrentPage(1);
    };

    const [allDepartments, setAllDepartments] = useState<Department[]>([]);

    const fetchAllDepartments = useCallback(async () => {
        try {
            const data = await $clientReq.get('/department/get');
            if (data && data.departments) {
                setAllDepartments(data.departments);
                setTreeData(buildTreeData(data.departments));
            }
        } catch (error) {
        }
    }, []);

    useEffect(() => {
        fetchAllDepartments();
    }, [fetchAllDepartments]);

    const fetchDepartments = useCallback(async (name = searchName) => {
        setLoading(true);
        const url = `/department/get?page=${currentPage}&pageSize=${pageSize}&name=${name}`;
        try {
          const data = await $clientReq.get(url);
          if (data) {
            setDepartments(data.departments);
            setTotalItems(data.total);
          }
        } catch (error) {
        } finally {
          setLoading(false);
        }
      }, [currentPage, pageSize, searchName]);

    useEffect(() => {
        fetchDepartments();
    }, [fetchDepartments]);

    const buildTreeData = (departments: Department[]): DataNode[] => {
        const treeMap = new Map<number, DataNode>();
        const rootNodes: DataNode[] = [];

        departments.forEach(dept => {
            const node: any = {
                title: dept.name,
                key: dept.id,
                value: dept.id,  // 添加 value 属性
                children: [],
                icon: getIconComponent(dept.icon),
            };
            treeMap.set(dept.id, node);
        });

        departments.forEach(dept => {
            const node = treeMap.get(dept.id)!;
            if (dept.parent_id === null) {
                rootNodes.push(node);
            } else {
                const parentNode = treeMap.get(dept.parent_id);
                if (parentNode) {
                    parentNode.children?.push(node);
                }
            }
        });

        return rootNodes;
    };

    const getIconComponent = (iconName: string) => {
        switch (iconName) {
            case 'ApartmentOutlined':
                return <ApartmentOutlined />;
            case 'TeamOutlined':
                return <TeamOutlined />;
            case 'UserOutlined':
                return <UserOutlined />;
            case 'BankOutlined':
                return <BankOutlined />;
            case 'GroupOutlined':
                return <GroupOutlined />;
            case 'SolutionOutlined':
                return <SolutionOutlined />;
            default:
                return <ApartmentOutlined />;
        }
    };

    const handleAdd = () => {
        setEditingDepartment(null);
        setIsDrawerOpen(true);
        form.resetFields();
    };

    const handleEdit = (record: Department) => {
        setEditingDepartment(record);
        setIsDrawerOpen(true);
        form.setFieldsValue(record);
    };

    const [saveLoading, setSaveLoading] = useState<boolean>(false);
    const handleSave = async () => {
        form.validateFields().then(async values => {
            setSaveLoading(true);
            try {
                const url = editingDepartment ? '/department/update' : '/department/create';
                const method = editingDepartment ? $clientReq.put : $clientReq.post;

                await method(url, {
                    ...values,
                    id: editingDepartment?.id
                });

                Notification({
                    type: 'success',
                    message: editingDepartment ? '更新成功!' : '添加成功!',
                    description: editingDepartment ? '部门信息已成功更新。' : '新部门已成功添加。',
                    placement: 'topRight',
                });
            } finally {
                setSaveLoading(false);
                setIsDrawerOpen(false);
                fetchDepartments();
            }
        });
    };

    const handleDelete = async (id: number) => {
        await $clientReq.delete(`/department/del?id=${id}`);
        Notification({
            type: 'success',
            message: '删除成功!',
            description: '部门已删除',
            placement: 'topRight'
        });
        fetchDepartments();
    };

    const handleSearch = () => {
        setCurrentPage(1);
        fetchDepartments(searchName);
    };

    const handleReset = () => {
        setSearchName('');
        setCurrentPage(1);
        fetchDepartments('');
    };

    const columns: any = [
        {
            title: '部门名称',
            dataIndex: 'name',
            key: 'name',
            render: (text: string, record: Department) => (
                <Space>
                    {getIconComponent(record.icon)}
                    <span>{text}</span>
                </Space>
            ),
        },
        {
            title: '创建时间',
            dataIndex: 'created_at',
            key: 'created_at',
            render: (text: string) => new Date(text).toLocaleString(),
        },
        {
            title: '更新时间',
            dataIndex: 'updated_at',
            key: 'updated_at',
            render: (text: string) => new Date(text).toLocaleString(),
        },
        {
            title: '操作',
            key: 'action',
            render: (_: any, record: Department) => (
                <Space size="middle">
                    <Button type="link" icon={<EditOutlined />} onClick={() => handleEdit(record)}>
                        编辑
                    </Button>
                    <Popconfirm
                        title="确定要删除此部门吗?"
                        onConfirm={() => handleDelete(record.id)}
                        okText="确定"
                        cancelText="取消"
                    >
                        <Button type="link" danger icon={<DeleteOutlined />}>
                            删除
                        </Button>
                    </Popconfirm>
                </Space>
            ),
        },
    ];

    return (
        <div className="p-4">
            <Row gutter={[16, 16]}>
                <Col span={6}>
                    <Input
                        placeholder="搜索部门名称"
                        value={searchName}
                        onChange={(e) => setSearchName(e.target.value)}
                        prefix={<SearchOutlined />}
                    />
                </Col>
                <Col span={18}>
                    <Space>
                        <Button type="primary" icon={<SearchOutlined />} onClick={handleSearch}>
                            搜索
                        </Button>
                        <Button icon={<ReloadOutlined />} onClick={handleReset}>
                            重置
                        </Button>
                        <Button type="primary" icon={<PlusOutlined />} onClick={handleAdd}>
                            添加部门
                        </Button>
                        <Dropdown menu={columns}>
                            <Button icon={<SettingOutlined />}>自定义列</Button>
                        </Dropdown>
                    </Space>
                </Col>
            </Row>

            <Table
                columns={columns}
                dataSource={departments}
                rowKey={(record) => record.id.toString()}
                loading={loading}
                pagination={false}
                style={{ marginTop: 16 }}
            />
            <Pagination
                currentPage={currentPage}
                pageSize={pageSize}
                totalItems={totalItems}
                onPageChange={handlePageChange}
                onPageSizeChange={handlePageSizeChange}
            />

            <Drawer
                title={editingDepartment ? '编辑部门' : '添加部门'}
                placement="right"
                onClose={() => setIsDrawerOpen(false)}
                open={isDrawerOpen}
                width={400}
                footer={
                    <div style={{ textAlign: 'right' }}>
                        <Button onClick={() => setIsDrawerOpen(false)} style={{ marginRight: 8 }}>
                            取消
                        </Button>
                        <Button onClick={handleSave} type="primary" loading={saveLoading}>
                            保存
                        </Button>
                    </div>
                }
            >
                <Form form={form} layout="vertical">
                    <Form.Item name="name" label="部门名称" rules={[{ required: true, message: '请输入部门名称' }]}>
                        <Input />
                    </Form.Item>
                    <Form.Item name="parent_id" label="上级部门">
                        <TreeSelect
                            treeData={treeData}
                            placeholder="请选择上级部门"
                            allowClear
                            treeDefaultExpandAll
                        />
                    </Form.Item>
                    <Form.Item name="icon" label="图标" rules={[{ required: true, message: '请选择图标' }]}>
                        <Select>
                            <Select.Option value="ApartmentOutlined">
                                <Space><ApartmentOutlined /> 公司</Space>
                            </Select.Option>
                            <Select.Option value="TeamOutlined">
                                <Space><TeamOutlined /> 团队</Space>
                            </Select.Option>
                            <Select.Option value="UserOutlined">
                                <Space><UserOutlined /> 个人</Space>
                            </Select.Option>
                            <Select.Option value="BankOutlined">
                                <Space><BankOutlined /> 机构</Space>
                            </Select.Option>
                            <Select.Option value="GroupOutlined">
                                <Space><GroupOutlined /> 群组</Space>
                            </Select.Option>
                            <Select.Option value="SolutionOutlined">
                                <Space><SolutionOutlined /> 解决方案</Space>
                            </Select.Option>
                        </Select>
                    </Form.Item>
                </Form>
            </Drawer>
        </div>
    );
};

export default DepartmentManagement;