import { Modal, Tree } from 'antd';
import { useState, useEffect } from 'react';

interface MenuAssignModalProps {
  visible: boolean;
  onCancel: () => void;
  onOk: (selectedMenus: string[]) => void;
  roleId: number;
  roleName: string;
}

const MenuAssignModal: React.FC<MenuAssignModalProps> = ({
  visible,
  onCancel,
  onOk,
  roleId,
  roleName,
}) => {
  const [menuTree, setMenuTree] = useState([]);
  const [selectedMenus, setSelectedMenus] = useState<string[]>([]);

  useEffect(() => {
    // 获取菜单树结构
    fetchMenuTree();
    // 获取该角色已分配的菜单
    fetchAssignedMenus(roleId);
  }, [roleId]);

  const fetchMenuTree = async () => {
    // 调用API获取菜单树结构
    // setMenuTree(response.data);
  };

  const fetchAssignedMenus = async (roleId: number) => {
    // 调用API获取该角色已分配的菜单
    // setSelectedMenus(response.data);
  };

  const onCheck = (checkedKeys: any) => {
    setSelectedMenus(checkedKeys);
  };

  return (
    <Modal
      title={`为 ${roleName} 分配菜单`}
      visible={visible}
      onCancel={onCancel}
      onOk={() => onOk(selectedMenus)}
    >
      <Tree
        checkable
        checkedKeys={selectedMenus}
        onCheck={onCheck}
        treeData={menuTree}
      />
    </Modal>
  );
};

export default MenuAssignModal;