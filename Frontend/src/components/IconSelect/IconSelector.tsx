import React, { useState, useEffect } from 'react';
import { Modal, Input, Pagination } from 'antd';
import { Icon } from '@iconify/react';
import { searchIcons } from '@/utils/iconLoader';

interface IconSelectorProps {
  visible: boolean;
  onClose: () => void;
  onSelect: (iconName: string) => void;
}

const IconSelector: React.FC<IconSelectorProps> = ({ visible, onClose, onSelect }) => {
  const [searchTerm, setSearchTerm] = useState('');
  const [currentPage, setCurrentPage] = useState(1);
  const [iconList, setIconList] = useState<string[]>([]);
  const pageSize = 48; // 每页显示的图标数量

  useEffect(() => {
    const allIcons = searchIcons(searchTerm);
    setIconList(allIcons);
    setCurrentPage(1);
  }, [searchTerm]);

  const paginatedIcons = iconList.slice((currentPage - 1) * pageSize, currentPage * pageSize);

  return (
    <Modal
      title="选择图标"
      open={visible}
      onCancel={onClose}
      footer={null}
      width={800}
    >
      <Input
        placeholder="搜索图标"
        value={searchTerm}
        onChange={(e) => setSearchTerm(e.target.value)}
        className="mb-4"
      />
      <div className="h-96 overflow-y-auto mb-4">
        <div className="grid grid-cols-6 gap-4">
          {paginatedIcons.map((name) => (
            <div
              key={name}
              onClick={() => onSelect(name)}
              className="flex flex-col items-center justify-center p-3 cursor-pointer hover:bg-gray-100 rounded-lg transition-colors"
            >
              <Icon icon={name} className="text-4xl" />
              <div className="mt-2 text-xs text-center truncate w-full">{name.split(':')[1]}</div>
            </div>
          ))}
        </div>
      </div>
      <Pagination
        current={currentPage}
        total={iconList.length}
        pageSize={pageSize}
        onChange={(page) => setCurrentPage(page)}
        showSizeChanger={false}
        className="text-center"
      />
    </Modal>
  );
};

export default IconSelector;