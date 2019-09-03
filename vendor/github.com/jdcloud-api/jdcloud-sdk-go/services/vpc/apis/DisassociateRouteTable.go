// Copyright 2018 JDCLOUD.COM
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// NOTE: This class is auto generated by the jdcloud code generator program.

package apis

import (
    "github.com/jdcloud-api/jdcloud-sdk-go/core"
)

type DisassociateRouteTableRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* RouteTable ID  */
    RouteTableId string `json:"routeTableId"`

    /* 路由表要解绑的子网ID，解绑后子网绑定默认路由表  */
    SubnetId string `json:"subnetId"`
}

/*
 * param regionId: Region ID (Required)
 * param routeTableId: RouteTable ID (Required)
 * param subnetId: 路由表要解绑的子网ID，解绑后子网绑定默认路由表 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDisassociateRouteTableRequest(
    regionId string,
    routeTableId string,
    subnetId string,
) *DisassociateRouteTableRequest {

	return &DisassociateRouteTableRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/routeTables/{routeTableId}:disassociateRouteTable",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        RouteTableId: routeTableId,
        SubnetId: subnetId,
	}
}

/*
 * param regionId: Region ID (Required)
 * param routeTableId: RouteTable ID (Required)
 * param subnetId: 路由表要解绑的子网ID，解绑后子网绑定默认路由表 (Required)
 */
func NewDisassociateRouteTableRequestWithAllParams(
    regionId string,
    routeTableId string,
    subnetId string,
) *DisassociateRouteTableRequest {

    return &DisassociateRouteTableRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/routeTables/{routeTableId}:disassociateRouteTable",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        RouteTableId: routeTableId,
        SubnetId: subnetId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDisassociateRouteTableRequestWithoutParam() *DisassociateRouteTableRequest {

    return &DisassociateRouteTableRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/routeTables/{routeTableId}:disassociateRouteTable",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *DisassociateRouteTableRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param routeTableId: RouteTable ID(Required) */
func (r *DisassociateRouteTableRequest) SetRouteTableId(routeTableId string) {
    r.RouteTableId = routeTableId
}

/* param subnetId: 路由表要解绑的子网ID，解绑后子网绑定默认路由表(Required) */
func (r *DisassociateRouteTableRequest) SetSubnetId(subnetId string) {
    r.SubnetId = subnetId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DisassociateRouteTableRequest) GetRegionId() string {
    return r.RegionId
}

type DisassociateRouteTableResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DisassociateRouteTableResult `json:"result"`
}

type DisassociateRouteTableResult struct {
}