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

type IGroupMemberRepo interface {
	Query(obj *entity.GroupMember) (*entity.GroupMember, error)
	QueryAll(obj *entity.GroupMember) ([]entity.GroupMember, error)
	Save(obj *entity.GroupMember) error
	Delete(obj *entity.GroupMember) error
	BeginTx() *gorm.DB
}
type GroupMemberService struct {
	repo IGroupMemberRepo
}

func NewGroupMemberService() *GroupMemberService {
	return &GroupMemberService{
		repo: repository.NewGroupMemberRepo(),
	}
}
func QueryGroupMember(obj *entity.GroupMember, repo IGroupMemberRepo) (*entity.GroupMember, error) {
	return repo.Query(obj)
}

// selectMembers 从服务器获取群成员
func (_self *GroupMemberService) selectMembers(gId uint64) ([]entity.GroupMember, *utils.Error) {
	if conf.GetLoginInfo().User == nil || conf.GetLoginInfo().User.Id == 0 {
		return nil, log.WithError(utils.ERR_QUERY_FAIL)
	}
	resultDTO, err := Post("/api/group/selectMembers", map[string]interface{}{"id": gId})
	if err != nil {
		return nil, log.WithError(err)
	}
	var members []entity.GroupMember
	e := util.Str2Obj(resultDTO.Data.(string), &members)
	if e != nil {
		return nil, log.WithError(utils.ERR_QUERY_FAIL)
	}
	//遍历members
	for i := 0; i < len(members); i++ {
		if members[i].User == nil || members[i].User.Id == 0 {
			continue
		}
		//先查询 是否存在 存在就不添加了
		userService := NewUserService()
		sysUser, e := QueryUser(members[i].UserId, userService.repo)
		if e != nil {
			return nil, log.WithError(utils.ERR_QUERY_FAIL)
		}
		if sysUser != nil && sysUser.Id != 0 {
			continue
		}
		e = userService.Save(members[i].User)
		if e != nil {
			return nil, log.WithError(utils.ERR_QUERY_FAIL)
		}
	}
	return members, nil
}
func (_self *GroupMemberService) QueryAll() ([]entity.GroupMember, error) {
	if conf.GetLoginInfo().User == nil || conf.GetLoginInfo().User.Id == 0 {
		return []entity.GroupMember{}, nil
	}
	return _self.repo.QueryAll(&entity.GroupMember{})
}
