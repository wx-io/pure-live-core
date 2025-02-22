package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/iyear/pure-live/model"
	"github.com/iyear/pure-live/pkg/ecode"
	"github.com/iyear/pure-live/pkg/format"
	"github.com/iyear/pure-live/service/srv_os"
)

func GetOSInfo(c *gin.Context) {
	info, err := srv_os.GetOSInfo()
	if err != nil {
		format.HTTP(c, ecode.ErrorGetOSInfo, err, nil)
		return
	}
	format.HTTP(c, ecode.Success, nil, info)
}

func GetSysMem(c *gin.Context) {
	r, err := srv_os.GetSysMem()
	if err != nil {
		format.HTTP(c, ecode.ErrorGetSysMem, err, nil)
		return
	}
	format.HTTP(c, ecode.Success, nil, r)
}

func GetSelfMem(c *gin.Context) {
	r, err := srv_os.GetSelfMem()
	if err != nil {
		format.HTTP(c, ecode.ErrorGetSelfMem, err, nil)
		return
	}
	format.HTTP(c, ecode.Success, nil, r)
}

func GetSysCPU(c *gin.Context) {
	r, err := srv_os.GetSysCPU()
	if err != nil {
		format.HTTP(c, ecode.ErrorGetSysCPU, err, nil)
		return
	}
	format.HTTP(c, ecode.Success, nil, r)
}

func GetSelfCPU(c *gin.Context) {
	r, err := srv_os.GetSelfCPU()
	if err != nil {
		format.HTTP(c, ecode.ErrorGetSelfCPU, err, nil)
		return
	}
	format.HTTP(c, ecode.Success, nil, r)
}

func GetOSAll(c *gin.Context) {
	info := &model.OSInfo{}
	sysCPU := &model.SysCPU{}
	selfCPU := &model.SelfCPU{}
	sysMem := &model.SysMem{}
	selfMem := &model.SelfMem{}

	err := func() error {
		var err error
		if info, err = srv_os.GetOSInfo(); err != nil {
			return err
		}
		if sysCPU, err = srv_os.GetSysCPU(); err != nil {
			return err
		}
		if selfCPU, err = srv_os.GetSelfCPU(); err != nil {
			return err
		}
		if sysMem, err = srv_os.GetSysMem(); err != nil {
			return err
		}
		if selfMem, err = srv_os.GetSelfMem(); err != nil {
			return err
		}
		return nil
	}()

	if err != nil {
		format.HTTP(c, ecode.ErrorGetOsAll, err, nil)
		return
	}

	format.HTTP(c, ecode.Success, nil, gin.H{
		"info":     info,
		"sys_cpu":  sysCPU,
		"self_cpu": selfCPU,
		"sys_mem":  sysMem,
		"self_mem": selfMem,
	})
}
