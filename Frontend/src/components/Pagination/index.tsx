import React from 'react';
import { Button, Tooltip, Select, InputNumber } from 'antd';
import { LeftOutlined, RightOutlined, DoubleLeftOutlined, DoubleRightOutlined } from '@ant-design/icons';

interface PaginationProps {
    currentPage: number;
    pageSize: number;
    totalItems: number;
    onPageChange: (page: number) => void;
    onPageSizeChange?: (size: number) => void; // 添加分页大小改变回调
}

const Pagination: React.FC<PaginationProps> = ({ currentPage, pageSize, totalItems, onPageChange, onPageSizeChange }) => {
    const totalPages = Math.ceil(totalItems / pageSize);
    const paginationItems = [];

    const renderPageButton = (page: number) => (
        <Tooltip title={`第 ${page} 页`} key={page}>
            <Button
                type={currentPage === page ? 'primary' : 'default'}
                shape="circle"
                onClick={() => onPageChange(page)}
                className={`!mx-1 ${currentPage === page ? 'bg-blue-600 text-white' : 'text-gray-900 hover:text-blue-600'}`}
            >
                {page}
            </Button>
        </Tooltip>
    );

    // 构建分页按钮逻辑
    if (totalPages <= 7) {
        // 少于7页时全部显示
        for (let i = 1; i <= totalPages; i++) {
            paginationItems.push(renderPageButton(i));
        }
    } else {
        // 多于7页时显示省略号
        paginationItems.push(renderPageButton(1));
        if (currentPage > 4) {
            paginationItems.push(<span key="left-dots" className="mx-1">...</span>);
        }
        const start = Math.max(2, currentPage - 2);
        const end = Math.min(totalPages - 1, currentPage + 2);
        for (let i = start; i <= end; i++) {
            paginationItems.push(renderPageButton(i));
        }
        if (currentPage < totalPages - 3) {
            paginationItems.push(<span key="right-dots" className="mx-1">...</span>);
        }
        paginationItems.push(renderPageButton(totalPages));
    }

    return (
        <div className="bg-white px-4 py-3 sm:px-6 shadow-lg dark:border-gray-700 dark:bg-gray-800">
            <div className="flex flex-col sm:flex-row sm:items-center sm:justify-between">
                {/* 显示信息 */}
                <div className="mb-4 sm:mb-0">
                    <p className="text-sm text-gray-700 dark:text-gray-300">
                        显示第
                        <span className="font-medium"> {((currentPage - 1) * pageSize) + 1} </span>
                        至
                        <span className="font-medium"> {Math.min(currentPage * pageSize, totalItems)} </span>
                        条，共
                        <span className="font-medium"> {totalItems} </span>
                        条记录
                    </p>
                </div>

                <div className="flex items-center justify-center space-x-2">
                    {/* 页面大小选择器 */}
                    {onPageSizeChange && (
                        <Select
                            defaultValue={pageSize}
                            style={{ width: 100 }}
                            onChange={onPageSizeChange}
                        >
                            <Select.Option value={5}>5 条/页</Select.Option>
                            <Select.Option value={10}>10 条/页</Select.Option>
                            <Select.Option value={20}>20 条/页</Select.Option>
                            <Select.Option value={50}>50 条/页</Select.Option>
                            <Select.Option value={100}>100 条/页</Select.Option>
                        </Select>
                    )}

                    {/* 跳转到第一页按钮 */}
                    {currentPage > 1 && (
                        <Tooltip title="第一页">
                            <Button
                                icon={<DoubleLeftOutlined />}
                                shape="circle"
                                size="large"
                                onClick={() => onPageChange(1)}
                                className="!mx-2 bg-gray-100 hover:bg-gray-200 dark:bg-gray-700 dark:hover:bg-gray-600"
                            />
                        </Tooltip>
                    )}

                    {/* 上一页按钮 */}
                    {currentPage > 1 && (
                        <Tooltip title="上一页">
                            <Button
                                icon={<LeftOutlined />}
                                shape="circle"
                                size="large"
                                onClick={() => onPageChange(currentPage - 1)}
                                className="!mx-2 bg-gray-100 hover:bg-gray-200 dark:bg-gray-700 dark:hover:bg-gray-600"
                            />
                        </Tooltip>
                    )}

                    {/* 页码按钮 */}
                    {paginationItems}

                    {/* 下一页按钮 */}
                    {currentPage < totalPages && (
                        <Tooltip title="下一页">
                            <Button
                                icon={<RightOutlined />}
                                shape="circle"
                                size="large"
                                onClick={() => onPageChange(currentPage + 1)}
                                className="!mx-2 bg-gray-100 hover:bg-gray-200 dark:bg-gray-700 dark:hover:bg-gray-600"
                            />
                        </Tooltip>
                    )}

                    {/* 跳转到最后一页按钮 */}
                    {currentPage < totalPages && (
                        <Tooltip title="最后一页">
                            <Button
                                icon={<DoubleRightOutlined />}
                                shape="circle"
                                size="large"
                                onClick={() => onPageChange(totalPages)}
                                className="!mx-2 bg-gray-100 hover:bg-gray-200 dark:bg-gray-700 dark:hover:bg-gray-600"
                            />
                        </Tooltip>
                    )}

                    {/* 页码跳转 */}
                    <InputNumber
                        min={1}
                        max={totalPages}
                        defaultValue={currentPage}
                        onChange={value => onPageChange(value || 1)}
                        className="w-20 ml-2"
                        placeholder="跳转页码"
                    />
                </div>
            </div>
        </div>
    );
};

export default Pagination;
