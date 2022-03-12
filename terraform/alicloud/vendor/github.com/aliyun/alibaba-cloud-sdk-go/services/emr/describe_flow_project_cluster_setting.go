package emr

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// DescribeFlowProjectClusterSetting invokes the emr.DescribeFlowProjectClusterSetting API synchronously
func (client *Client) DescribeFlowProjectClusterSetting(request *DescribeFlowProjectClusterSettingRequest) (response *DescribeFlowProjectClusterSettingResponse, err error) {
	response = CreateDescribeFlowProjectClusterSettingResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeFlowProjectClusterSettingWithChan invokes the emr.DescribeFlowProjectClusterSetting API asynchronously
func (client *Client) DescribeFlowProjectClusterSettingWithChan(request *DescribeFlowProjectClusterSettingRequest) (<-chan *DescribeFlowProjectClusterSettingResponse, <-chan error) {
	responseChan := make(chan *DescribeFlowProjectClusterSettingResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeFlowProjectClusterSetting(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// DescribeFlowProjectClusterSettingWithCallback invokes the emr.DescribeFlowProjectClusterSetting API asynchronously
func (client *Client) DescribeFlowProjectClusterSettingWithCallback(request *DescribeFlowProjectClusterSettingRequest, callback func(response *DescribeFlowProjectClusterSettingResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeFlowProjectClusterSettingResponse
		var err error
		defer close(result)
		response, err = client.DescribeFlowProjectClusterSetting(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// DescribeFlowProjectClusterSettingRequest is the request struct for api DescribeFlowProjectClusterSetting
type DescribeFlowProjectClusterSettingRequest struct {
	*requests.RpcRequest
	ClusterId string `position:"Query" name:"ClusterId"`
	ProjectId string `position:"Query" name:"ProjectId"`
}

// DescribeFlowProjectClusterSettingResponse is the response struct for api DescribeFlowProjectClusterSetting
type DescribeFlowProjectClusterSettingResponse struct {
	*responses.BaseResponse
	RequestId    string                                       `json:"RequestId" xml:"RequestId"`
	GmtCreate    int64                                        `json:"GmtCreate" xml:"GmtCreate"`
	GmtModified  int64                                        `json:"GmtModified" xml:"GmtModified"`
	ProjectId    string                                       `json:"ProjectId" xml:"ProjectId"`
	ClusterId    string                                       `json:"ClusterId" xml:"ClusterId"`
	K8sClusterId string                                       `json:"K8sClusterId" xml:"K8sClusterId"`
	DefaultUser  string                                       `json:"DefaultUser" xml:"DefaultUser"`
	DefaultQueue string                                       `json:"DefaultQueue" xml:"DefaultQueue"`
	UserList     UserListInDescribeFlowProjectClusterSetting  `json:"UserList" xml:"UserList"`
	QueueList    QueueListInDescribeFlowProjectClusterSetting `json:"QueueList" xml:"QueueList"`
	HostList     HostListInDescribeFlowProjectClusterSetting  `json:"HostList" xml:"HostList"`
}

// CreateDescribeFlowProjectClusterSettingRequest creates a request to invoke DescribeFlowProjectClusterSetting API
func CreateDescribeFlowProjectClusterSettingRequest() (request *DescribeFlowProjectClusterSettingRequest) {
	request = &DescribeFlowProjectClusterSettingRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Emr", "2016-04-08", "DescribeFlowProjectClusterSetting", "emr", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeFlowProjectClusterSettingResponse creates a response to parse from DescribeFlowProjectClusterSetting response
func CreateDescribeFlowProjectClusterSettingResponse() (response *DescribeFlowProjectClusterSettingResponse) {
	response = &DescribeFlowProjectClusterSettingResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}