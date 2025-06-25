package system

type Menu struct {
	Id            int64   `json:"id"`
	ParentId      int64   `json:"parentId"`                    // 父Id
	Name          string  `json:"name"`                        // 名称
	Level         int8    `json:"level"`                       // 层级
	Type          int8    `json:"type"`                        // 类型: 1MENU, 2BUTTON
	Path          string  `json:"path"`                        // 前端路由地址
	PermissionKey string  `json:"permissionKey"`               // 权限标识 如 user:add
	OrderNum      int     `json:"orderNum"`                    // 排序
	Visible       int8    `json:"visible"`                     // 是否可见 1是 0否
	Children      []*Menu `json:"children,omitempty" gorm:"-"` // 子菜单
}
