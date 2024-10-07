'use client'
import React, { useRef, useEffect, useState } from "react";
import { usePathname, useRouter } from "next/navigation";
import Link from "next/link";
import Image from "next/image";
import SidebarLinkGroup from "./SidebarLinkGroup";
import { MenuItem } from "@/types/menu";
import { Tooltip } from "antd";
import LOGO from '@/public/icons/logo.png'
import { Icon } from '@iconify/react';

interface SidebarProps {
    sidebarOpen: boolean;
    setSidebarOpen: (open: boolean) => void;
    menuData: MenuItem[];
    sidebarExpanded: boolean;
    setSidebarExpanded: (expanded: boolean) => void;
}

const buildTree = (items: MenuItem[]): MenuItem[] => {
    const itemMap = new Map<number, MenuItem>();
    const rootItems: MenuItem[] = [];

    // 首先，按照 sort 字段对所有项目进行排序
    const sortedItems = [...items].sort((a:any, b:any) => {
        if (a.sort === b.sort) {
            // 如果 sort 值相同，则按照 id 排序
            return a.id - b.id;
        }
        return a.sort - b.sort;
    });

    sortedItems.forEach(item => {
        itemMap.set(item.id, { ...item, children: [] });
    });

    itemMap.forEach(item => {
        if (item.parent_id === null) {
            rootItems.push(item);
        } else {
            const parent = itemMap.get(item.parent_id);
            if (parent) {
                parent.children?.push(item);
            }
        }
    });

    // 对每个父项的子项进行排序
    const sortChildren = (items: MenuItem[]) => {
        items.forEach(item => {
            if (item.children && item.children.length > 0) {
                item.children.sort((a:any, b:any) => {
                    if (a.sort === b.sort) {
                        return a.id - b.id;
                    }
                    return a.sort - b.sort;
                });
                sortChildren(item.children);
            }
        });
    };

    sortChildren(rootItems);

    return rootItems;
};

const Sidebar: React.FC<SidebarProps> = ({
    sidebarOpen,
    setSidebarOpen,
    menuData,
    sidebarExpanded,
    setSidebarExpanded
}) => {
    const pathname = usePathname();
    const router = useRouter();
    const sidebar = useRef<HTMLDivElement | null>(null);
    const trigger = useRef<HTMLButtonElement | null>(null);
    const [activeMenu, setActiveMenu] = useState<string | null>(null);
    const [expandedMenus, setExpandedMenus] = useState<Set<string>>(new Set());

    useEffect(() => {
        const clickHandler = ({ target }: MouseEvent) => {
            if (!sidebar.current || !trigger.current || !sidebarOpen) return;
            if (!(target instanceof Node) || sidebar.current.contains(target) || trigger.current.contains(target)) {
                return;
            }
            setSidebarOpen(false);
        };

        document.addEventListener("click", clickHandler);
        return () => document.removeEventListener("click", clickHandler);
    }, [sidebarOpen, setSidebarOpen]);

    useEffect(() => {
        const keyHandler = ({ key }: KeyboardEvent) => {
            if (key === "Escape") setSidebarOpen(false);
        };
        document.addEventListener("keydown", keyHandler);
        return () => document.removeEventListener("keydown", keyHandler);
    }, [setSidebarOpen]);

    useEffect(() => {
        localStorage.setItem("sidebar-expanded", sidebarExpanded.toString());
        document.body.classList.toggle("sidebar-expanded", sidebarExpanded);
    }, [sidebarExpanded]);

    useEffect(() => {
        const currentPath = pathname;
        const findActiveMenu = (items: MenuItem[]): string | null => {
            for (const item of items) {
                if (item.path === currentPath) {
                    if (item.parent_id) {
                        setExpandedMenus(prev => new Set(prev).add(item.parent_id!.toString()));
                    }
                    return item.path;
                }
                if (item.children) {
                    const result = findActiveMenu(item.children);
                    if (result) {
                        setExpandedMenus(prev => new Set(prev).add(item.path));
                        return result;
                    }
                }
            }
            return null;
        };
        const activeMenuPath = findActiveMenu(buildTree(menuData));
        setActiveMenu(activeMenuPath);
    }, [pathname, menuData]);

    const handleMenuClick = (path: string, hasChildren: boolean) => {
        if (hasChildren) {
            setExpandedMenus(prev => {
                const newSet = new Set(prev);
                if (newSet.has(path)) {
                    newSet.delete(path);
                } else {
                    newSet.add(path);
                }
                return newSet;
            });
        } else {
            setActiveMenu(path);
        }
        if (!sidebarExpanded) setSidebarExpanded(true);
    };

    const renderMenuItem = (item: MenuItem, isSubMenu = false) => {
        const isActive = pathname === item.path;
        const isExpanded = expandedMenus.has(item.path);
        const hasChildren = item.children && item.children.length > 0;

        return (
            <SidebarLinkGroup key={item.id} activeCondition={isActive}>
                {(handleClick, open) => (
                    <>
                        <Link
                            href={item.path}
                            className={`group flex items-center p-2 rounded-md transition-all duration-200 ${isActive
                                ? "bg-indigo-400 text-white"
                                : "text-gray-300 hover:bg-indigo-400 hover:text-white"
                                }`}
                            onClick={(e) => {
                                if (hasChildren) {
                                    e.preventDefault();
                                }
                                handleMenuClick(item.path, hasChildren || false);
                                handleClick();
                            }}
                        >
                            {item.icon && (!isSubMenu || sidebarExpanded) && (
                                <Tooltip placement="rightTop" title={sidebarExpanded ? '' : item.title} color="geekblue" className="z-9999">
                                    <Icon icon={item.icon} width="20" height="20" className="ml-1" />
                                </Tooltip>
                            )}
                            {sidebarExpanded && (
                                <span
                                    className="ml-3 text-sm font-medium transition-all duration-200"
                                    style={{ whiteSpace: "nowrap" }}
                                >
                                    {item.title}
                                </span>
                            )}
                            {hasChildren && sidebarExpanded && (
                                <svg
                                    className={`ml-auto transition-transform duration-300 ${isExpanded ? "rotate-180" : ""}`}
                                    width="16"
                                    height="16"
                                    fill="currentColor"
                                    viewBox="0 0 20 20"
                                >
                                    <path fillRule="evenodd" d="M5.293 7.293a1 1 0 011.414 0L10 10.586l3.293-3.293a1 1 0 111.414 1.414l-4 4a1 1 0 01-1.414 0l-4-4a1 1 0 010-1.414z" clipRule="evenodd" />
                                </svg>
                            )}
                        </Link>
                        {hasChildren && sidebarExpanded && (
                            <div className={`ml-4 mt-2 space-y-1 transition-all duration-200 ${!isExpanded ? "hidden" : ""}`}>
                                {item.children?.map((subItem) => renderMenuItem(subItem, true))}
                            </div>
                        )}
                    </>
                )}
            </SidebarLinkGroup>
        );
    };

    const treeMenuData = buildTree(menuData);

    return (
        <aside
            ref={sidebar}
            className={`absolute left-0 top-0 z-99 flex h-screen ${sidebarExpanded ? "w-64" : "w-16"
                } flex-col bg-gray-900 text-gray-200 overflow-x-hidden overflow-y-auto duration-300 ease-linear shadow-lg lg:static lg:translate-x-0 ${sidebarOpen ? "translate-x-0" : "-translate-x-full"
                }`}
        >
            <div className="flex items-center justify-between px-4 py-5 bg-gray-800">
                <Link href="/">
                    <div className="flex items-center">
                        <Tooltip placement="rightTop" title="RBAC Admin" color="geekblue">
                            <Image
                                className="ml-1"
                                width={sidebarExpanded ? 45 : 30}
                                height={sidebarExpanded ? 45 : 30}
                                src={LOGO}
                                alt="Logo"
                                priority
                            />
                        </Tooltip>
                        <span
                            className={`ml-3 text-xl font-bold text-indigo-400 transition-all duration-300 ${sidebarExpanded ? "opacity-100 translate-x-0" : "opacity-0 -translate-x-10"
                                }`}
                            style={{ whiteSpace: "nowrap" }}
                        >
                            RBAC Admin
                        </span>
                    </div>
                </Link>
                <button
                    ref={trigger}
                    onClick={() => setSidebarOpen(!sidebarOpen)}
                    className="block lg:hidden text-gray-400 hover:text-gray-300"
                >
                    {/* Mobile toggle button content */}
                </button>
            </div>

            <nav className="flex-1 px-3 py-4 space-y-1">
                {treeMenuData.map((item) => renderMenuItem(item))}
            </nav>
        </aside>
    );
};

export default Sidebar;