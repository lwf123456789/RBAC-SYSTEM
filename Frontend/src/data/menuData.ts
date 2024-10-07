import { MenuItem } from '@/types/menu'

const menuData: MenuItem[] = [
  {
      "id": 1,
      "title": "系统管理",
      "path": "/system",
      "component_path": "",
      "icon": "emojione:gear",
      "parent_id": null,
      "sort": 1,
      "status": 1,
      "created_at": "2024-09-21T16:12:18+08:00",
      "updated_at": "2024-09-21T16:12:18+08:00"
  },
  {
      "id": 2,
      "title": "用户管理",
      "path": "/system/user",
      "component_path": "system/UserManagement",
      "icon": "emojione:busts-in-silhouette",
      "parent_id": 1,
      "sort": 1,
      "status": 1,
      "created_at": "2024-09-21T16:12:39+08:00",
      "updated_at": "2024-09-27T15:50:19.893+08:00"
  },
  {
      "id": 3,
      "title": "角色管理",
      "path": "/system/role",
      "component_path": "system/RoleManagement",
      "icon": "emojione:shield",
      "parent_id": 1,
      "sort": 2,
      "status": 1,
      "created_at": "2024-09-21T16:12:39+08:00",
      "updated_at": "2024-09-21T16:12:39+08:00"
  },
  {
      "id": 4,
      "title": "菜单管理",
      "path": "/system/menu",
      "component_path": "system/MenuManagement",
      "icon": "emojione:clipboard",
      "parent_id": 1,
      "sort": 3,
      "status": 1,
      "created_at": "2024-09-21T16:12:39+08:00",
      "updated_at": "2024-09-28T14:55:36.3+08:00"
  }
]

export default menuData;