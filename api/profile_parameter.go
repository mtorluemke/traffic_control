// Copyright 2015 Comcast Cable Communications Management, LLC

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

// http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This file was initially generated by gen_to_start.go (add link), as a start
// of the Traffic Ops golang data model

package api

import (
	"encoding/json"
	_ "github.com/Comcast/traffic_control/traffic_ops/experimental/server/output_format" // needed for swagger
	"github.com/jmoiron/sqlx"
	"log"
	"time"
)

type ProfileParameter struct {
	LastUpdated time.Time `db:"last_updated" json:"lastUpdated"`
	Links       struct {
		Self    string `db:"self" json:"_self"`
		Profile struct {
			ID  int64  `db:"profile" json:"id"`
			Ref string `db:"profile_id_ref" json:"_ref"`
		} `json:"profile" db:-`
		Parameter struct {
			ID  int64  `db:"parameter" json:"id"`
			Ref string `db:"parameter_id_ref" json:"_ref"`
		} `json:"parameter" db:-`
	} `json:"_links" db:-`
}

// @Title getProfileParameterById
// @Description retrieves the profile_parameter information for a certain id
// @Accept  application/json
// @Param   id              path    int     false        "The row id"
// @Success 200 {array}    ProfileParameter
// @Resource /api/2.0
// @Router /api/2.0/profile_parameter/{id} [get]
func getProfileParameterById(id int, db *sqlx.DB) (interface{}, error) {
	ret := []ProfileParameter{}
	arg := ProfileParameter{}
	arg.Links.Profile.ID = int64(id)
	queryStr := "select *, concat('" + API_PATH + "profile_parameter/', id) as self "
	queryStr += ", concat('" + API_PATH + "profile/', profile) as profile_id_ref"
	queryStr += ", concat('" + API_PATH + "parameter/', parameter) as parameter_id_ref"
	queryStr += " from profile_parameter where Links.Profile.ID=:Links.Profile.ID"
	nstmt, err := db.PrepareNamed(queryStr)
	err = nstmt.Select(&ret, arg)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	nstmt.Close()
	return ret, nil
}

// @Title getProfileParameters
// @Description retrieves the profile_parameter
// @Accept  application/json
// @Success 200 {array}    ProfileParameter
// @Resource /api/2.0
// @Router /api/2.0/profile_parameter [get]
func getProfileParameters(db *sqlx.DB) (interface{}, error) {
	ret := []ProfileParameter{}
	queryStr := "select *, concat('" + API_PATH + "profile_parameter/', id) as self "
	queryStr += ", concat('" + API_PATH + "profile/', profile) as profile_id_ref"
	queryStr += ", concat('" + API_PATH + "parameter/', parameter) as parameter_id_ref"
	queryStr += " from profile_parameter"
	err := db.Select(&ret, queryStr)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return ret, nil
}

// @Title postProfileParameter
// @Description enter a new profile_parameter
// @Accept  application/json
// @Param                 Body body     ProfileParameter   true "ProfileParameter object that should be added to the table"
// @Success 200 {object}    output_format.ApiWrapper
// @Resource /api/2.0
// @Router /api/2.0/profile_parameter [post]
func postProfileParameter(payload []byte, db *sqlx.DB) (interface{}, error) {
	var v ProfileParameter
	err := json.Unmarshal(payload, &v)
	if err != nil {
		log.Println(err)
	}
	sqlString := "INSERT INTO profile_parameter("
	sqlString += "profile"
	sqlString += ",parameter"
	sqlString += ") VALUES ("
	sqlString += ":profile"
	sqlString += ",:parameter"
	sqlString += ")"
	result, err := db.NamedExec(sqlString, v)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return result, err
}

// @Title putProfileParameter
// @Description modify an existing profile_parameterentry
// @Accept  application/json
// @Param   id              path    int     true        "The row id"
// @Param                 Body body     ProfileParameter   true "ProfileParameter object that should be added to the table"
// @Success 200 {object}    output_format.ApiWrapper
// @Resource /api/2.0
// @Router /api/2.0/profile_parameter/{id}  [put]
func putProfileParameter(id int, payload []byte, db *sqlx.DB) (interface{}, error) {
	var v ProfileParameter
	err := json.Unmarshal(payload, &v)
	v.Links.Profile.ID = int64(id) // overwrite the id in the payload
	if err != nil {
		log.Println(err)
		return nil, err
	}
	v.LastUpdated = time.Now()
	sqlString := "UPDATE profile_parameter SET "
	sqlString += "profile = :profile"
	sqlString += ",parameter = :parameter"
	sqlString += ",last_updated = :last_updated"
	sqlString += " WHERE Links.Profile.ID=:Links.Profile.ID"
	result, err := db.NamedExec(sqlString, v)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return result, err
}

// @Title delProfileParameterById
// @Description deletes profile_parameter information for a certain id
// @Accept  application/json
// @Param   id              path    int     false        "The row id"
// @Success 200 {array}    ProfileParameter
// @Resource /api/2.0
// @Router /api/2.0/profile_parameter/{id} [delete]
func delProfileParameter(id int, db *sqlx.DB) (interface{}, error) {
	arg := ProfileParameter{}
	arg.Links.Profile.ID = int64(id)
	result, err := db.NamedExec("DELETE FROM profile_parameter WHERE id=:id", arg)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return result, err
}
