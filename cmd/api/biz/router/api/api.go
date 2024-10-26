// Code generated by hertz generator. DO NOT EDIT.

package api

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	api "github.com/wushiling50/aster/cmd/api/biz/handler/api"
)

/*
 This file will register all the routes of the services in the master idl.
 And it will update automatically when you use the "update" command for the idl.
 So don't modify the contents of the file, or your code will be deleted when it is updated.
*/

// Register register routes based on the IDL 'api.${HTTP Method}' annotation.
func Register(r *server.Hertz) {

	root := r.Group("/", rootMw()...)
	{
		_aster := root.Group("/aster", _asterMw()...)
		_aster.GET("/search", append(_searchMw(), api.Search)...)
		{
			_nation := _aster.Group("/nation", _nationMw()...)
			_nation.GET("/guess", append(_nationguessMw(), api.NationGuess)...)
		}
		{
			_talent := _aster.Group("/talent", _talentMw()...)
			{
				_rank := _talent.Group("/rank", _rankMw()...)
				_rank.GET("/", append(_talentrankMw(), api.TalentRank)...)
			}
		}
	}
}
