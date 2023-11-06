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

type IFriendRepo interface {
	Query(obj *entity.Friend) (*entity.Friend, error)
	QueryAll(obj *entity.Friend) ([]entity.Friend, error)
	Save(obj *entity.Friend) error
	Delete(obj *entity.Friend) error
	BeginTx() *gorm.DB
}
type FriendService struct {
	repo IFriendRepo
}

func NewFriendService() *FriendService {
	return &FriendService{
		repo: repository.NewFriendRepo(),
	}
}
func QueryFriend(id uint64, repo IFriendRepo) (*entity.Friend, error) {
	return repo.Query(&entity.Friend{Id: id})
}
func QueryFriendAll(repo IFriendRepo) ([]entity.Friend, error) {
	if conf.GetLoginInfo().User == nil || conf.GetLoginInfo().User.Id == 0 {
		return []entity.Friend{}, nil
	}
	return repo.QueryAll(&entity.Friend{Me: conf.GetLoginInfo().User.Id})
}
func (_self *FriendService) updateOne(he, me uint64) *utils.Error {
	var req = make(map[string]uint64)
	req["he"] = he
	req["me"] = me
	resultDTO, err := util.Post("/api/friend/selectOne", req)
	if err != nil {
		return log.WithError(err)
	}
	var fa entity.Friend
	_ = util.Str2Obj(resultDTO.Data.(string), &fa)
	if fa.Id != 0 {
		//保存到数据库
		e := _self.repo.Save(&fa)
		if e != nil {
			return log.WithError(utils.ERR_OPERATION_FAIL)
		}
	}
	return nil
}
func (_self *FriendService) Del(id uint64) *utils.Error {
	//先通过服务器删除
	req := make(map[string]uint64)
	req["id"] = id
	_, err := util.Post("/api/friend/selectOne", req)
	if err != nil {
		return log.WithError(err)
	}
	tx := _self.repo.BeginTx()
	if err := tx.Error; err != nil {
		return log.WithError(utils.ERR_OPERATION_FAIL)
	}
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	err = func() *utils.Error {
		friend, e := QueryFriend(id, _self.repo)
		if e != nil {
			return log.WithError(utils.ERR_OPERATION_FAIL)
		}
		//再删除本地好友记录 不删用户
		e = _self.repo.Delete(friend)
		if e != nil {
			return log.WithError(utils.ERR_OPERATION_FAIL)
		}
		//再修改好友申请为拒绝
		faService := NewFriendApplyService()
		e = faService.updateReject(friend.Me, friend.He)
		if e != nil {
			return log.WithError(utils.ERR_OPERATION_FAIL)
		}
		e = faService.updateReject(friend.He, friend.Me)
		if e != nil {
			return log.WithError(utils.ERR_OPERATION_FAIL)
		}
		e = tx.Commit().Error
		if e != nil {
			return log.WithError(utils.ERR_OPERATION_FAIL)
		}
		return nil
	}()
	if err != nil {
		tx.Rollback()
	}
	return err
}
func (_self *FriendService) SelectOne(id uint64) (*entity.Friend, *utils.Error) {
	friend, e := QueryFriend(id, _self.repo)
	if e != nil {
		return nil, log.WithError(utils.ERR_QUERY_FAIL)
	}
	if friend != nil {
		userService := NewUserService()
		user, e := QueryUser(friend.He, userService.repo)
		if e != nil || user == nil {
			return nil, log.WithError(utils.ERR_QUERY_FAIL)
		}
		friend.HeUser = user
	}
	return friend, nil
}
func (_self *FriendService) SelectAll() ([]entity.Friend, *utils.Error) {
	friends, e := QueryFriendAll(_self.repo)
	if e != nil {
		return nil, log.WithError(utils.ERR_QUERY_FAIL)
	}
	if len(friends) != 0 {
		//封装好友信息
		for i := 0; i < len(friends); i++ {
			userService := NewUserService()
			user, e := QueryUser(friends[i].He, userService.repo)
			if e != nil || user == nil {
				return nil, log.WithError(utils.ERR_QUERY_FAIL)
			}
			friends[i].HeUser = user
			if friends[i].Name == "" {
				friends[i].Name = user.Nickname
			}
		}
	}
	//没用查到 就从后台查一次
	resultDTO, err := util.Post("/api/friend/selectAll", nil)
	if err != nil {
		return nil, log.WithError(err)
	}
	var fs []entity.Friend
	e = util.Str2Obj(resultDTO.Data.(string), &fs)
	if e != nil || fs == nil {
		return []entity.Friend{}, nil
	}
	tx := _self.repo.BeginTx()
	if err := tx.Error; err != nil {
		return nil, log.WithError(utils.ERR_QUERY_FAIL)
	}
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	r, err := func() ([]entity.Friend, *utils.Error) {
		// 保存到数据库
		for _, v := range fs {
			e := _self.repo.Save(&v)
			if e != nil {
				return nil, log.WithError(utils.ERR_QUERY_FAIL)
			}
			//先查询 是否存在 存在就不添加了
			//保存对应的用户信息
			userService := NewUserService()
			sysUser, e := QueryUser(v.He, userService.repo)
			if e != nil {
				return nil, log.WithError(utils.ERR_QUERY_FAIL)
			}
			if sysUser != nil {
				continue
			}
			e = userService.Save(v.HeUser)
			if e != nil {
				return nil, log.WithError(utils.ERR_QUERY_FAIL)
			}
		}
		e = tx.Commit().Error
		if e != nil {
			return nil, log.WithError(utils.ERR_QUERY_FAIL)
		}
		return fs, nil
	}()
	if err != nil {
		tx.Rollback()
	}
	return r, err
}
func (_self *FriendService) UpdateName(id uint64, name string) *utils.Error {
	friend, e := QueryFriend(id, _self.repo)
	if e != nil {
		return log.WithError(utils.ERR_OPERATION_FAIL)
	}
	friend.Name = name
	//服务器修改
	_, err := util.Post("/api/friend/edit", friend)
	if err != nil {
		return log.WithError(utils.ERR_OPERATION_FAIL)
	}
	e = _self.repo.Save(friend)
	if e != nil {
		return log.WithError(utils.ERR_OPERATION_FAIL)
	}
	return nil
}
