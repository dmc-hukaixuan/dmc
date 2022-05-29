package admin

import (
	model "dmc/kernel/model/admin"
	"dmc/kernel/model/common/response"
	"dmc/kernel/service/admin"
	"dmc/kernel/util"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	yaml "gopkg.in/yaml.v2"
)

type ProcessManagementApi struct {
	// BaseController
}

type nodeinfo struct {
	Name         string   `json:"name"`
	NodeID       string   `json:"nodeID"`
	ID           int      `json:"id"`
	TemplateList []string `json:"templateList"`
}

type transition struct {
	ConditionName string `json:"conditionName"`
	Conditions    string `json:"conditions"`
	ID            int    `json:"id"`
	EntityID      string `json:"entityID"`
}

type path struct {
	StartNode string                            `json:"startNode"`
	EndNode   string                            `json:"endNode"`
	Path      map[string]map[string]interface{} `json:"path"`
}

// process node struct detail
type nodeDetail struct {
	ID           int      `json:"id"`
	TemplateList []string `json:"templateList"`
	Name         string   `json:"name"`
}

// transation struct detail
type transations struct {
	ID               int    `json:"id"`
	Condition        string `json:"condition"`
	ConditionName    string `json:"conditionName"`
	ConditionLinking string `json:"conditionLinking"`
	SourceP          string `json:"sourceP"`
	TargetP          string `json:"targetP"`
	TransitionID     string `json:"transitionID"`
	Count            string `json:"Count"`
}

// get permission process
func (p *ProcessManagementApi) ProcessOverview(c *gin.Context) {

	if typelist, total, err := admin.ProcessModelA.ProcessTypeList(); err != nil {
		response.FailWithMessage("获取失败", c)
	} else {
		fmt.Println("list ----------------:", typelist, "total:", total)
		// response.OkWithDetailed(response.PageResult{
		// 	List:  list,
		// 	Total: total,
		// }, "获取成功", c)
	}

	if list, total, err := admin.ProcessModelA.ProcessList(); err != nil {
		response.FailWithMessage("获取失败", c)
	} else {
		fmt.Println("list :", list, "total:", total)
		// response.OkWithDetailed(response.PageResult{
		// 	List:  list,
		// 	Total: total,
		// }, "获取成功", c)
	}
}

// 在 tag 中添加 omitempty 忽略空值
// 注意这里 hobby,omitempty 合起来是 json tag 值，中间用英文逗号分隔
type FeildData struct {
	Name                 string            `json:"name"`
	Default              string            `json:"default"`
	FieldType            string            `json:"type"`
	Label                string            `json:"label"`
	Placeholder          string            `json:"placeholder,omitempty"`
	Display              int               `json:"display"`
	Impacts              []string          `json:"impacts,omitempty"`
	DependsOn            []string          `json:"dependsOn,omitempty"`
	PromptCode           int               `json:"promptCode,omitempty"`
	PromptMessage        string            `json:"PromptMessage,omitempty"`
	AutoComplete         bool              `json:"autoComplete,omitempty"`
	Options              map[string]string `json:"options,omitempty"`
	OptionsType          string            `json:"optionsType,omitempty"`
	OptionsValueComments map[string]string `json:"optionsValueComments,omitempty"`
	HintMessage          string            `json:"hint,omitempty"`
	HintType             int               `json:"hint,omitempty"`
	RegexError           string            `json:"regexError,omitempty"`
	Regex                string            `json:"regex,omitempty"`
}

// process type edit
func (p *ProcessManagementApi) ProcessTypeEdit(c *gin.Context) {
	//get process type id
	tid := c.Param("id")
	id, _ := strconv.Atoi(tid)
	var processTypeData model.ProcessType
	// get data
	if id > 0 {
		processType, err1 := admin.ProcessModelA.ProcessTypeGet(id)
		if err1 != nil {
			response.FailWithMessage("获取失败", c)
		}
		processTypeData = processType
	}
	// process type edit
	a := map[string]*FeildData{
		"name": &FeildData{
			Name:      "name",
			Default:   processTypeData.Name,
			FieldType: "Text",
			Label:     "Process type name",
			Display:   1,
		},
		"valid": &FeildData{
			Name:      "valid",
			Default:   strconv.Itoa(processTypeData.Valid),
			FieldType: "dropdown",
			Options: map[string]string{
				"1": "Valid",
				"2": "Invalid",
			},
			Label: "Validity",
		},
	}
	fmt.Println("a ", a)
	response.SuccessWithDetailed(gin.H{
		"fieldOrder": [...]string{"name", "valid"},
		"fieldData":  a,
	}, "获取成功", c)
}

// process type save
func (p *ProcessManagementApi) ProcessTypeSave(c *gin.Context) {
	var pt model.ProcessType
	_ = c.ShouldBindJSON(&pt)
	fmt.Println("pttttt: ", pt)
	user_id, ok := c.Get("userID")
	fmt.Println("user_id :", user_id, " is ok:", ok)
	// add or update type
	if pt.ID > 0 {
		pt.ChangeBy = user_id.(int)
		if _, err := admin.ProcessModelA.ProcessTypeUpdate(&pt); err != nil {
			response.FailWithMessage("Update failded,Please try again later", c)
		} else {
			response.SuccessWithData(pt, c)
		}
	} else {
		pt.CreateBy = user_id.(int)
		pt.ChangeBy = user_id.(int)
		if _, err := admin.ProcessModelA.ProcessTypeAdd(&pt); err != nil {
			response.FailWithMessage("Update failded,", c)
		} else {
			response.SuccessWithData(pt, c)
		}
	}
}

// 其实这里应该是一个通用的提交按钮
// 这里先在这里
// @TODO : 需要把这个抽象到 model 层
type ProcessManagement struct {
	SubAction string                 `json:"subaction"`
	Data      map[string]interface{} `json:"data"`
}

// porcess management main function
func (p *ProcessManagementApi) ProcessManagement(c *gin.Context) {
	var pm ProcessManagement
	_ = c.ShouldBindJSON(&pm)

	// get processid
	processID, ok := pm.Data["processID"]

	// user, _ := c.Get("token")
	user_id, ok := c.Get("userID")
	fmt.Println("user_id ----------------------:", user_id, " is ok:", ok)

	if pm.SubAction == "edit" && ok {
		fmt.Println("pm ", pm, processID)
		processBaseData(int(processID.(float64)), c)
	} else if pm.SubAction == "save" {
		processSave(pm.Data, c)
	}
}

// get process base data
func processBaseData(processID int, c *gin.Context) {
	var processData model.Process
	initData := make(map[string]interface{})

	//  process detail info
	if processID > 0 {
		process, err1 := admin.ProcessModelA.ProcessGet(processID)

		if err1 != nil {
			response.FailWithMessage("获取失败", c)
		}
		layout := map[string]map[string]string{}
		config := map[string]interface{}{}
		// transition := map[string]interface{}{}
		json.Unmarshal([]byte(process.Layout), &layout)
		json.Unmarshal([]byte(process.Config), &config)
		// json.Marshal()
		//nodelist := processTransitionList(process.ID)
		//initData["porcessNode"] = config
		initData["layout"] = layout
		initData["processTransition"] = processTransitionList(process.ID)
		initData["processConfig"] = config
		initData["porcessNode"] = processNodeList(process.ID)

		//	initData["processConfig"] = process.ProcessConfig
		processData = process
	} else {
		initData = processNewCanvas()
	}
	// fmt.Println("processData :", processData)
	a := map[string]*FeildData{
		"name": &FeildData{
			Name:      "name",
			Default:   processData.Name,
			FieldType: "text",
			Label:     "Process name",
			Display:   2,
		},
		"description": &FeildData{
			Name:      "description",
			Default:   processData.Description,
			FieldType: "textarea",
			Label:     "Description",
			Display:   1,
		},
		"processState": &FeildData{
			Name:      "processState",
			Default:   processData.StateEntityID,
			FieldType: "dropdown",
			Options: map[string]string{
				"1": "Valid",
				"2": "Invalid",
			},
			Label: "Validity",
		},
		"processType": &FeildData{
			Name:      "processType",
			Default:   strconv.Itoa(processData.ProcessType),
			FieldType: "dropdown",
			Options: map[string]string{
				"1": "Valid",
				"2": "Invalid",
			},
			Label: "Select business type for process",
		},
	}

	response.SuccessWithDetailed(gin.H{
		"fieldOrder":         [...]string{"name", "description", "processState", "processType"},
		"fieldData":          a,
		"porcessNode":        initData["porcessNode"],
		"layout":             initData["layout"],
		"processConfig":      initData["processConfig"],
		"processTransition":  initData["processTransition"],
		"nodeBaseData":       processNodeBaseData(),
		"transitionBaseData": processTransitionBaseData(),
	}, "获取成功", c)
}

// process save
func processSave(pd map[string]interface{}, c *gin.Context) {

	// add process data
	// jsonStu, err := json.Marshal(pd["nodeDelete"])
	// nodeLocation, _ := yaml.Marshal(pd["nodeLocation"])
	// processConfig, _ := yaml.Marshal(pd["processConfig"])
	// transitionValue, _ := yaml.Marshal(pd["transitionValue"])
	// jsonStu, err := json.Marshal(pd["nodeDelete"])
	// nodeValue1, _ := json.Marshal(pd["nodeValue"])
	// transitionValue, _ := json.Marshal(pd["transitionValue"])
	// nodeValue, _ := yaml.Marshal(pd["nodeValue"])
	nodeLocation, _ := json.Marshal(pd["nodeLocation"])
	processConfig, _ := json.Marshal(pd["processConfig"])
	ptid, _ := strconv.Atoi(pd["processType"].(string))
	// process data
	Process := &model.Process{}

	// nodeinfo := make(map[string]nodeDetail)
	// yaml.Unmarshal(nodeValue, &nodeinfo)
	// json fromat node value
	nodeValue1, _ := json.Marshal(pd["nodeValue"])
	nodeinfo1 := make(map[string]nodeDetail)
	json.Unmarshal(nodeValue1, &nodeinfo1)
	// process add
	if pd["processEntityID"] == "" {
		// process data
		Process = &model.Process{
			Name:          pd["processName"].(string),
			Description:   pd["processDescription"].(string),
			EntityID:      "Process-" + util.GenerateRandomString(32),
			StateEntityID: pd["processStateEntityID"].(string),
			Layout:        string(nodeLocation),
			Config:        string(processConfig),
			ProcessType:   ptid,
			CreateTime:    time.Now().Format("2006-01-02 15:04:05"),
			CreateBy:      1,
			ChangeTime:    time.Now().Format("2006-01-02 15:04:05"),
			ChangeBy:      1,
		}
		if typelist, err := admin.ProcessModelA.ProcessAdd(Process); err != nil {
			//response.FailWithMessage("获取失败", c)
		} else {
			fmt.Println("list ----------------:", typelist, "total:", typelist)
			//response.SuccessWithMessage("添加成功", c)
		}
	} else {
		processID, _ := strconv.Atoi(pd["processID"].(string))
		// process data
		Process = &model.Process{
			ID:            processID,
			Name:          pd["processName"].(string),
			Description:   pd["processDescription"].(string),
			EntityID:      "Process-" + util.GenerateRandomString(32),
			StateEntityID: pd["processStateEntityID"].(string),
			Layout:        string(nodeLocation),
			Config:        string(processConfig),
			ProcessType:   ptid,
			ChangeTime:    time.Now().Format("2006-01-02 15:04:05"),
			ChangeBy:      1,
		}
		fmt.Println("Process :", Process)
		// process update
		if _, err := admin.ProcessModelA.ProcessUpdate(Process); err != nil {
			fmt.Println("ProcessUpdate  err :", err)
			//response.FailWithMessage("获取失败", c)
		} else {
			//fmt.Println("list ----------------:", typelist, "total:", typelist)
			//response.SuccessWithMessage("添加成功", c)
		}
	}

	// process node info
	var processNode []model.ProcessNode
	for k, v := range nodeinfo1 {
		//tt := v.TemplateList
		config, _ := yaml.Marshal(v.TemplateList)
		if v.ID > 0 {
			NodeUpdate := model.ProcessNode{
				ID:         v.ID,
				Name:       v.Name,
				ProcessID:  Process.ID,
				NodeID:     k,
				Config:     string(config),
				ChangeTime: time.Now().Format("2006-01-02 15:04:05"),
				ChangeBy:   1,
			}
			// ask db
			if _, err := admin.ProcessModelA.NodeUpdate(&NodeUpdate); err != nil {
				fmt.Println("==================:", err)
			} else {
				fmt.Println("================sss==:", err)
				//fmt.Println("list ----------------:", typelist, "total:", typelist)
			}
		} else {
			// assemble the node data into an slice
			processNode = append(processNode, model.ProcessNode{
				Name:       v.Name,
				ProcessID:  Process.ID,
				NodeID:     k,
				Config:     string(config),
				CreateTime: time.Now().Format("2006-01-02 15:04:05"),
				CreateBy:   1,
				ChangeTime: time.Now().Format("2006-01-02 15:04:05"),
				ChangeBy:   1,
			})
		}
	}

	if len(processNode) > 0 {
		if typelist, err := admin.ProcessModelA.NodeAdd(processNode); err != nil {
			//response.FailWithMessage("获取失败", c)
		} else {
			fmt.Println("list ----------------:", typelist, "total:", typelist)
			//response.SuccessWithMessage("添加成功", c)
		}
	}

	// process transition
	transitionValue, _ := json.Marshal(pd["transitionValue"])
	traninfo := make(map[string]transations)

	// assign json data to struct
	json.Unmarshal(transitionValue, &traninfo)
	var processTransition []model.ProcessTransition

	// add transation action detail info to db
	for k, v := range traninfo {
		// format the condition data as json
		transitionValue, _ := json.Marshal(map[string]interface{}{
			"condition":        v.Condition,
			"conditionName":    v.ConditionName,
			"conditionLinking": v.ConditionLinking,
			"sourceP":          v.SourceP,
			"targetP":          v.TargetP,
			"count":            v.Count,
		})
		if v.ID > 0 {
			processTransition := &model.ProcessTransition{
				ID:           v.ID,
				Name:         v.ConditionName,
				ProcessID:    Process.ID,
				TransitionID: k,
				Config:       string(transitionValue),
				ChangeTime:   time.Now().Format("2006-01-02 15:04:05"),
				ChangeBy:     1,
			}
			if _, err := admin.ProcessModelA.TransitionUpdate(processTransition); err != nil {
				fmt.Println("TransitionUpdate rereeeee-------:", err)
				//response.FailWithMessage("获取失败", c)
			} else {
				//fmt.Println("list ----------------:", typelist, "total:", typelist)
				//response.SuccessWithMessage("添加成功", c)
			}
		} else {
			// perfermore impore
			fmt.Println(" configjson ", string(transitionValue))
			processTransition = append(processTransition, model.ProcessTransition{
				Name:         v.ConditionName,
				ProcessID:    Process.ID,
				TransitionID: k,
				Config:       string(transitionValue),
				CreateTime:   time.Now().Format("2006-01-02 15:04:05"),
				CreateBy:     1,
				ChangeTime:   time.Now().Format("2006-01-02 15:04:05"),
				ChangeBy:     1,
			})
		}
	}
	// when edit a process, if there is newly added process node, use the method
	// of betch of batch insert to the db
	if len(processTransition) > 0 {
		if typelist, err := admin.ProcessModelA.TransitionAdd(processTransition); err != nil {
			//response.FailWithMessage("获取失败", c)
		} else {
			fmt.Println("list ----------------:", typelist, "total:", typelist)
			//response.SuccessWithMessage("添加成功", c)
		}
	}
	response.SuccessWithMessage("添加成功", c)
}

// get detail list on transition
// assemble the data structure
func processTransitionList(processID int) *map[string]interface{} {
	transitionList, _ := admin.ProcessModelA.TransitionListbyProceeID(processID)
	tl := map[string]interface{}{}

	for _, v := range transitionList {
		config := map[string]interface{}{}
		json.Unmarshal([]byte(v.Config), &config)
		tl[v.TransitionID] = map[string]interface{}{
			"ID":            v.ID,
			"ConditionName": v.Name,
			"Condition":     config,
			"TransitionID":  v.TransitionID,
		}
	}

	return &tl
}

// get process each node detail data
// assemble the data structure
func processNodeList(processID int) *map[string]interface{} {

	nodeList, _ := admin.ProcessModelA.NodeListByProcessID(processID)
	tl := map[string]interface{}{}

	// for each build up node list data
	for _, v := range nodeList {
		templatelist := []string{}
		json.Unmarshal([]byte(v.Config), &templatelist)
		tl[v.NodeID] = nodeDetail{
			ID:           v.ID,
			Name:         v.Name,
			TemplateList: templatelist,
		}
	}

	return &tl
}

// create a new process, initialize to generate three nodes
func processNewCanvas() map[string]interface{} {
	// process node
	a := map[string]*nodeinfo{
		"start": &nodeinfo{
			Name:   "Start",
			NodeID: "Node-" + util.GenerateRandomString(32),
		},
		"middle": &nodeinfo{
			Name:   "Process node",
			NodeID: "Node-" + util.GenerateRandomString(32),
		},
		"end": &nodeinfo{
			Name:   "End",
			NodeID: "Node-" + util.GenerateRandomString(32),
		},
	}

	// after the page is initalized, the position of the inital node
	// on the canvas
	layout := map[string]map[string]string{
		a["start"].NodeID: {
			"left": "295px",
			"top":  "100px",
		},
		a["middle"].NodeID: {
			"left": "350px",
			"top":  "250px",
		},
		a["end"].NodeID: {
			"left": "350px",
			"top":  "4000px",
		},
	}
	startTransition := "Transition-" + util.GenerateRandomString(32)
	endTransition := "Transition-" + util.GenerateRandomString(32)
	return map[string]interface{}{
		"porcessNode": a,
		"layout":      layout,
		"processTransition": map[string]interface{}{
			startTransition: &transition{
				ConditionName: "Process transtion",
				EntityID:      a["start"].NodeID,
			},
			endTransition: &transition{
				ConditionName: "Process transtion",
				EntityID:      a["middle"].NodeID,
			},
		},
		"processConfig": &path{
			StartNode: a["start"].NodeID,
			EndNode:   a["end"].NodeID,
			Path: map[string]map[string]interface{}{
				a["start"].NodeID: map[string]interface{}{
					startTransition: map[string]interface{}{
						"NextNode":         a["middle"].NodeID,
						"TransitionAction": []string{},
					},
				},
				a["middle"].NodeID: {
					endTransition: map[string]interface{}{
						"NextNode":         a["end"].NodeID,
						"TransitionAction": []string{},
					},
				},
				a["end"].NodeID: {},
			},
		},
	}
}

// transition base data for add new transitaion condition
func processTransitionBaseData() map[string]interface{} {
	return map[string]interface{}{
		"conditionFieldValue": &FeildData{
			Name:      "conditionFieldValue",
			Default:   "",
			FieldType: "text",
			Label:     "dropdown",
			Display:   1,
			Options:   map[string]string{},
		},
		"conditionType": &FeildData{
			Name:      "conditionType",
			Default:   "",
			FieldType: "dropdown",
			Label:     "TypeIdentifier Search",
			Display:   2,
			Options: map[string]string{
				"or":  "or",
				"and": "and",
			},
		},
		"conditionCompare": &FeildData{
			Name:      "conditionCompare",
			Default:   "",
			FieldType: "dropdown",
			Label:     "type",
			Display:   2,
			Options: map[string]string{
				"ne":      "Not equals",
				"ib":      "Is between",
				"inb":     "Is not between",
				"le":      "Smaller than equals",
				"ioo":     "Is one of",
				"inoo":    "Is not one of",
				"null":    "Null",
				"lt":      "Smaller than",
				"notNull": "Not null",
				"gt":      "Greater than",
				"eq":      "Equals",
				"ge":      "Greater than equals",
			},
		},
		"conditionName": &FeildData{
			Name:      "conditionName",
			Default:   "Process transition",
			FieldType: "text",
			Label:     "Transition name",
			Display:   2,
		},
		"conditionLinking": &FeildData{
			Name:      "conditionLinking",
			Default:   "",
			FieldType: "dropdown",
			Label:     "Type of Linking between Conditions",
			Display:   2,
			Options: map[string]string{
				"1": "Valid",
				"2": "Invalid",
			},
		},
		"conditionLabel": &FeildData{
			Name:      "templateList",
			Default:   "",
			FieldType: "dropdown",
			Label:     "Node operation name",
			Display:   2,
			Options:   map[string]string{},
		},
		"conditionFieldName": &FeildData{
			Name:      "conditionFieldName",
			Default:   "",
			FieldType: "dropdown",
			Label:     "Name",
			Display:   2,
			Options:   map[string]string{},
		},
	}
}

// base node model for
func processNodeBaseData() map[string]interface{} {
	// get ticket template list
	return map[string]interface{}{
		"name": &FeildData{
			Name:      "name",
			Default:   "",
			FieldType: "text",
			Label:     "Node name",
			Display:   1,
		},
		"templateList": &FeildData{
			Name:      "templateList",
			Default:   "",
			FieldType: "dropdown",
			Label:     "Node operation name",
			Display:   1,
			Options:   map[string]string{},
		},
	}
}
