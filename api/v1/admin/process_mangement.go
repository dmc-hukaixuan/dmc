package admin

import (
	model "dmc/kernel/model/admin"
	"dmc/kernel/model/common/request"
	"dmc/kernel/model/common/response"
	"dmc/kernel/service/admin/process"
	"dmc/kernel/util"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
	yaml "gopkg.in/yaml.v2"
)

type ProcessManagementApi struct {
	// BaseController
}

// get permission process
func (p *ProcessManagementApi) ProcessOverview(c *gin.Context) {

	if typelist, total, err := process.ProcessTypeList(); err != nil {
		response.FailWithMessage("获取失败", c)
	} else {
		fmt.Println("list ----------------:", typelist, "total:", total)
		// response.SuccessWithDetailed(response.PageResult{
		// 	List:  list,
		// 	Total: total,
		// }, "获取成功", c)
	}

	if list, total, err := process.ProcessList(); err != nil {
		response.FailWithMessage("获取失败", c)
	} else {
		fmt.Println("list :", list, "total:", total)
		response.SuccessWithDetailed(gin.H{
			"List":  list,
			"Total": total,
		}, "获取成功", c)
	}
}

// process type edit
func (p *ProcessManagementApi) ProcessTypeEdit(c *gin.Context) {
	//get process type id
	tid := c.Param("id")
	id, _ := strconv.Atoi(tid)
	var processTypeData model.ProcessType
	// get data
	if id > 0 {
		processType, err1 := process.ProcessTypeGet(id)
		if err1 != nil {
			response.FailWithMessage("获取失败", c)
		}
		processTypeData = processType
	}
	// process type edit
	a := map[string]*model.FieldData{
		"name": &model.FieldData{
			Name:      "name",
			Default:   processTypeData.Name,
			FieldType: "Text",
			Label:     "Process type name",
			Display:   1,
		},
		"valid": &model.FieldData{
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
		if _, err := process.ProcessTypeUpdate(&pt); err != nil {
			response.FailWithMessage("Update failded,Please try again later", c)
		} else {
			response.SuccessWithData(pt, c)
		}
	} else {
		pt.CreateBy = user_id.(int)
		pt.ChangeBy = user_id.(int)
		if _, err := process.ProcessTypeAdd(&pt); err != nil {
			response.FailWithMessage("Update failded,", c)
		} else {
			response.SuccessWithData(pt, c)
		}
	}
}

// porcess management main function
func (p *ProcessManagementApi) ProcessManagement(c *gin.Context) {
	var pm request.SubActionData
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
	fmt.Println("processData ", processData)
	initData := make(map[string]interface{})

	//  process detail info
	if processID > 0 {
		process, err1 := process.ProcessGet(processID)

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

	response.SuccessWithDetailed(gin.H{
		// "fieldOrder":         [...]string{"name", "description", "processState", "processType"},
		"baseData": map[string]string{
			"name":         "",
			"description":  "",
			"processState": "",
			"processType":  "",
		},
		"templateList": "",
		"porcessNode":  initData["porcessNode"],
		"ndoes":        initData["nodes"],
	}, "获取成功", c)
}

// process save
// add process data
// jsonStu, err := json.Marshal(pd["nodeDelete"])
// nodeLocation, _ := yaml.Marshal(pd["nodeLocation"])
// processConfig, _ := yaml.Marshal(pd["processConfig"])
// transitionValue, _ := yaml.Marshal(pd["transitionValue"])
// jsonStu, err := json.Marshal(pd["nodeDelete"])
// nodeValue1, _ := json.Marshal(pd["nodeValue"])
// transitionValue, _ := json.Marshal(pd["transitionValue"])
// nodeValue, _ := yaml.Marshal(pd["nodeValue"])
// nodeinfo := make(map[string]nodeDetail)
// yaml.Unmarshal(nodeValue, &nodeinfo)
func processSave(pd map[string]interface{}, c *gin.Context) {
	var pl model.PorcessLayout

	mapstructure.Decode(pd, &pl)
	// process data
	Process := &model.Process{}
	// process add
	if pd["processEntityID"] == "" {
		// process data
		proceelayout, _ := json.Marshal(pl.FlowData)
		proceeConfig, _ := json.Marshal(pl.ProcessData)
		Process = &model.Process{
			Name:          pl.ProcessData.Name,
			Description:   pl.ProcessData.Comments,
			EntityID:      "Process-" + util.GenerateRandomString(32),
			StateEntityID: pl.ProcessData.ProcessState,
			Layout:        string(proceelayout),
			Config:        string(proceeConfig),
			ProcessType:   pl.ProcessData.ProcessTypeID,
			CreateTime:    time.Now().Format("2006-01-02 15:04:05"),
			CreateBy:      1,
			ChangeTime:    time.Now().Format("2006-01-02 15:04:05"),
			ChangeBy:      1,
		}
		if typelist, err := process.ProcessAdd(Process); err != nil {
			//response.FailWithMessage("获取失败", c)
		} else {
			fmt.Println("list ----------------:", typelist, "total:", typelist)
			//response.SuccessWithMessage("添加成功", c)
		}
	} else {
		processID, _ := strconv.Atoi(pd["processID"].(string))
		proceelayout, _ := json.Marshal(pl.FlowData)
		proceeConfig, _ := json.Marshal(pl.ProcessData)
		// process data
		Process = &model.Process{
			ID:            processID,
			Name:          pl.ProcessData.Name,
			Description:   pl.ProcessData.Comments,
			EntityID:      "Process-" + util.GenerateRandomString(32),
			StateEntityID: pd["processStateEntityID"].(string),
			Layout:        string(proceelayout),
			Config:        string(proceeConfig),
			ProcessType:   pl.ProcessData.ProcessTypeID,
			ChangeTime:    time.Now().Format("2006-01-02 15:04:05"),
			ChangeBy:      1,
		}
		fmt.Println("Process :", Process)
		// process update
		if _, err := process.ProcessUpdate(Process); err != nil {
			fmt.Println("ProcessUpdate  err :", err)
			//response.FailWithMessage("获取失败", c)
		} else {
			//fmt.Println("list ----------------:", typelist, "total:", typelist)
			//response.SuccessWithMessage("添加成功", c)
		}
	}

	/*
		--------------------
			process node add
			and check ndoe usage, if nodes link any process delete it
		--------------------
	*/
	nodeUpdate(pl.FlowData.NodeLayout, pl.ProcessData, pl.ProcessData.ProcessID)
	// process transition
	transitionValue, _ := json.Marshal(pd["transitionValue"])
	traninfo := make(map[string]model.Transations)

	// assign json data to struct
	json.Unmarshal(transitionValue, &traninfo)
	transitionUpdate(pl.FlowData.Edges, pl.ProcessData, pl.ProcessData.ProcessID)

	//  process transition action
	response.SuccessWithMessage("添加成功", c)
}

// node add or update or delete
func nodeUpdate(nodeLayout []model.NodeLayout, processData model.ProcessData, ProcessID string) {
	nodelist, _ := process.NodeProcessListGet(ProcessID)
	var processNode []model.ProcessNode
	for _, v := range nodeLayout {
		config, _ := yaml.Marshal(v)
		if v.ID != "" {
			if _, usageNodeID := nodelist[v.ID]; usageNodeID {
				delete(nodelist, v.ID)
			}
			NodeUpdate := model.ProcessNode{
				ID:         0,
				Name:       processData.Nodes[v.ID].Name,
				ProcessID:  processData.ProcessID,
				NodeID:     v.ID,
				Config:     string(config),
				ChangeTime: time.Now().Format("2006-01-02 15:04:05"),
				ChangeBy:   1,
			}

			// link activity role to link
			// ask db
			if _, err := process.NodeUpdate(&NodeUpdate, processData.Nodes[v.ID].TemplateList); err != nil {
				fmt.Println("==================:", err)
			} else {
				fmt.Println("================sss==:", err)
				//fmt.Println("list ----------------:", typelist, "total:", typelist)
			}
		} else {
			// assemble the node data into an slice
			processNode = append(processNode, model.ProcessNode{
				Name:       processData.Nodes[v.ID].Name,
				ProcessID:  processData.ProcessID,
				NodeID:     v.ID,
				Config:     string(config),
				CreateTime: time.Now().Format("2006-01-02 15:04:05"),
				CreateBy:   1,
				ChangeTime: time.Now().Format("2006-01-02 15:04:05"),
				ChangeBy:   1,
			})
		}
	}
	// delete unusage node
	for k, _ := range nodelist {
		process.NodeDelete(k)
	}
	if len(processNode) > 0 {
		if typelist, err := process.NodeAdd(processNode); err != nil {
			//response.FailWithMessage("获取失败", c)
		} else {
			fmt.Println("list ----------------:", typelist, "total:", typelist)
			//response.SuccessWithMessage("添加成功", c)
		}
	}
}

func transitionUpdate(edges []model.Edges, processData model.ProcessData, ProcessID string) {
	var processTransition []model.ProcessTransition

	// get transition list
	transitionList, _ := process.TransitionListbyProceeIDGet(ProcessID)

	// add transation action detail info to db
	for _, v := range edges {
		transitionLayout, _ := json.Marshal(v)
		transitionConfig, _ := json.Marshal(processData.Transition[v.ID])
		transitionActionConfig, _ := json.Marshal(processData.TransitionAction[v.ID])
		if v.ID != "" {
			if _, usageNodeID := transitionList[v.ID]; usageNodeID {
				delete(transitionList, v.ID)
			}
			processTransition := &model.ProcessTransition{
				Name:         processData.Transition[v.ID].Name,
				ProcessID:    processData.ProcessID,
				TransitionID: v.ID,
				Layout:       string(transitionLayout),
				Config:       string(transitionConfig),
				ChangeTime:   time.Now().Format("2006-01-02 15:04:05"),
				ChangeBy:     1,
			}

			if _, err := process.TransitionUpdate(processTransition); err != nil {

				fmt.Println("TransitionUpdate rereeeee-------:", err)
				//response.FailWithMessage("获取失败", c)
			} else {
				//fmt.Println("list ----------------:", typelist, "total:", typelist)
				//response.SuccessWithMessage("添加成功", c)
			}
			// process transition action
			processTransitionAction := &model.ProcessTransitionAction{
				Name:         processData.TransitionAction[v.ID].Name,
				ProcessID:    processData.ProcessID,
				TransitionID: processTransition.TransitionID,
				Config:       string(transitionActionConfig),
				ChangeTime:   time.Now().Format("2006-01-02 15:04:05"),
				ChangeBy:     1,
			}
			if err := process.TransitionActionAdd(processTransitionAction); err != nil {
				fmt.Println("TransitionUpdate rereeeee-------:", err)
				//response.FailWithMessage("获取失败", c)
			} else {
				//fmt.Println("list ----------------:", typelist, "total:", typelist)
				//response.SuccessWithMessage("添加成功", c)
			}
		} else {
			// perfermore impore
			processTransition = append(processTransition, model.ProcessTransition{
				Name:         processData.Transition[v.ID].Name,
				ProcessID:    processData.ProcessID,
				TransitionID: v.ID,
				Config:       string(transitionConfig),
				Layout:       string(transitionLayout),
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
		if typelist, err := process.TransitionAdd(processTransition); err != nil {
			//response.FailWithMessage("获取失败", c)
		} else {
			fmt.Println("list ----------------:", typelist, "total:", typelist)
			//response.SuccessWithMessage("添加成功", c)
		}
	}

	// delete unusage transition and transitionaction
	for k, _ := range transitionList {
		process.TransitionActionDelete(k)
	}
}

// get detail list on transition
// assemble the data structure
func processTransitionList(processID int) *map[string]interface{} {
	transitionList, _ := process.TransitionListbyProceeID(processID)
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

	nodeList, _ := process.NodeListByProcessID(processID)
	tl := map[string]interface{}{}

	// for each build up node list data
	for _, v := range nodeList {
		templatelist := []string{}
		json.Unmarshal([]byte(v.Config), &templatelist)
		tl[v.NodeID] = model.NodeDetail{
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
	a := map[string]*model.Nodeinfo{
		"start": &model.Nodeinfo{
			Name:   "Start",
			NodeID: "Node-" + util.GenerateRandomString(32),
		},
		"middle": &model.Nodeinfo{
			Name:   "Process node",
			NodeID: "Node-" + util.GenerateRandomString(32),
		},
		"end": &model.Nodeinfo{
			Name:   "End",
			NodeID: "Node-" + util.GenerateRandomString(32),
		},
	}

	// after the page is initalized, the position of the inital node
	// on the canvas
	layout := [...]model.ActivityNote{
		model.ActivityNote{
			ID:       "Node-" + util.GenerateRandomString(32),
			Label:    "start",
			NodeType: "bizFlowNode",
			Style:    "",
			Left:     295,
			Right:    100,
		},
		model.ActivityNote{
			ID:       "Node-" + util.GenerateRandomString(32),
			Label:    "middle",
			NodeType: "bizFlowNode",
			Style:    "",
			Left:     350,
			Right:    250,
		},
		model.ActivityNote{
			ID:       "Node-" + util.GenerateRandomString(32),
			Label:    "end",
			NodeType: "bizFlowNode",
			Style:    "",
			Left:     350,
			Right:    400,
		},
	}

	startTransition := "Transition-" + util.GenerateRandomString(32)
	endTransition := "Transition-" + util.GenerateRandomString(32)
	return map[string]interface{}{
		"porcessNode":       a,
		"nodes":             layout,
		"processTransition": map[string]interface{}{
			// startTransition: &model.Transition{
			// 	ConditionName: "Process transtion",
			// 	EntityID:      a["start"].NodeID,
			// },
			// endTransition: &model.Transition{
			// 	ConditionName: "Process transtion",
			// 	EntityID:      a["middle"].NodeID,
			// },
		},
		"processConfig": &model.Path{
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
		"conditionFieldValue": &model.FieldData{
			Name:      "conditionFieldValue",
			Default:   "",
			FieldType: "text",
			Label:     "dropdown",
			Display:   1,
			Options:   map[string]string{},
		},
		"conditionType": &model.FieldData{
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
		"conditionCompare": &model.FieldData{
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
		"conditionName": &model.FieldData{
			Name:      "conditionName",
			Default:   "Process transition",
			FieldType: "text",
			Label:     "Transition name",
			Display:   2,
		},
		"conditionLinking": &model.FieldData{
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
		"conditionLabel": &model.FieldData{
			Name:      "templateList",
			Default:   "",
			FieldType: "dropdown",
			Label:     "Node operation name",
			Display:   2,
			Options:   map[string]string{},
		},
		"conditionFieldName": &model.FieldData{
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
		"name":         "",
		"templateList": []string{},
		// "name": &FeildData{
		// 	Name:      "name",
		// 	Default:   "",
		// 	FieldType: "text",
		// 	Label:     "Node name",
		// 	Display:   1,
		// },
		// "templateList": &FeildData{
		// 	Name:      "templateList",
		// 	Default:   "",
		// 	FieldType: "dropdown",
		// 	Label:     "Node operation name",
		// 	Display:   1,
		// 	Options:   map[string]string{},
		// },
	}
}
