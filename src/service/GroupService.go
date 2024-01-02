package service

import (
	"IM-Service/src/configs/conf"
	utils "IM-Service/src/configs/err"
	"IM-Service/src/configs/log"
	"IM-Service/src/entity"
	"IM-Service/src/repository"
	"IM-Service/src/util"
	"gorm.io/gorm"
)

type IGroupRepo interface {
	Query(obj *entity.Group) (*entity.Group, error)
	QueryAll(obj *entity.Group) ([]entity.Group, error)
	Save(obj *entity.Group) error
	Delete(obj *entity.Group) error
	BeginTx() *gorm.DB
}
type GroupService struct {
	repo IGroupRepo
}

func NewGroupService() *GroupService {
	return &GroupService{
		repo: repository.NewGroupRepo(),
	}
}
func QueryGroup(obj *entity.Group, repo IGroupRepo) (*entity.Group, error) {
	return repo.Query(obj)
}

// Invite 邀请好友进群 ids是用户ID
func (_self *GroupService) Invite(id uint64, ids []uint64) *utils.Error {
	if conf.GetLoginInfo().User == nil || conf.GetLoginInfo().User.Id == 0 {
		return log.WithError(utils.ERR_NOT_LOGIN)
	}
	//先从服务器创建
	_, err := Post("/api/group/invite", map[string]interface{}{"ids": ids, "id": id})
	if err != nil {
		return log.WithError(err)
	}
	return nil
}

// Create 创建群聊
func (_self *GroupService) Create(ids []uint64, tp int, password string) (*entity.Group, *utils.Error) {
	if conf.GetLoginInfo().User == nil || conf.GetLoginInfo().User.Id == 0 {
		return nil, log.WithError(utils.ERR_NOT_LOGIN)
	}
	//先从服务器创建
	resultDTO, err := Post("/api/group/create", map[string]interface{}{"ids": ids, "type": tp, "password": password})
	if err != nil {
		return nil, log.WithError(err)
	}
	var group entity.Group
	e := util.Str2Obj(resultDTO.Data.(string), &group)
	if e != nil {
		return nil, log.WithError(utils.ERR_OPERATION_FAIL)
	}
	if group.Id != 0 {
		//保存到数据库
		e := _self.repo.Save(&group)
		if e != nil {
			return nil, log.WithError(utils.ERR_OPERATION_FAIL)
		}
	}
	return &group, nil
}
func (_self *GroupService) SelectAll() ([]entity.Group, error) {
	if conf.GetLoginInfo().User == nil || conf.GetLoginInfo().User.Id == 0 {
		return []entity.Group{}, nil
	}
	return _self.repo.QueryAll(&entity.Group{UserId: conf.GetLoginInfo().User.Id})
}

func (_self *GroupService) SelectOne(target uint64, refresh bool) (*entity.Group, *utils.Error) {
	group, e := _self.repo.Query(&entity.Group{Id: target})
	if e != nil {
		return nil, log.WithError(utils.ERR_GROUP_GET_FAIL)
	}
	if group == nil || refresh {
		resultDTO, err := Post("/api/group/selectOne", map[string]uint64{"id": target})
		if err != nil {
			return nil, log.WithError(err)
		}
		//如果服务器获取失败
		if resultDTO.Data == nil {
			return nil, log.WithError(utils.ERR_GROUP_GET_FAIL)
		}
		var g entity.Group
		e := util.Str2Obj(resultDTO.Data.(string), &g)
		if e != nil {
			return nil, log.WithError(utils.ERR_GROUP_GET_FAIL)
		}
		if g.Id != 0 {
			//保存到数据库
			g.UserId = conf.GetLoginInfo().User.Id
			e := _self.repo.Save(&g)
			if e != nil {
				return nil, log.WithError(utils.ERR_OPERATION_FAIL)
			}
		}
		return &g, nil
	}
	return group, nil
}

func (_self *GroupService) NeedPassword(tp string, target uint64) string {
	if "group" != tp {
		return "-1"
	}
	group, e := _self.SelectOne(target, false)
	if e != nil || group == nil {
		return "-1"
	}
	if group.Type == 2 {
		//需要密码 但是有密码
		if conf.Conf.Pwds[tp+"_"+util.Uint642Str(target)] != "" {
			return "2"
		}
		return "1" //前端 仅在返回1时 打开密码输入框
	}
	return "-1"
}
