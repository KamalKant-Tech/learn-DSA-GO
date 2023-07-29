package main

import (
	"encoding/json"
	"fmt"
)

type Organization struct {
	ID                 string         `json:"id"`
	Name               string         `json:"name"`
	ParentOrganization *Organization  `json:"parentOrganization,omitempty"`
	SubOrganizations   []Organization `json:"subOrganizations,omitempty"`
}

type OrgSuborgs struct {
	OrgId   string       `json:"organizationId"`
	OrgName string       `json:"organizationName"`
	SubOrg  []OrgSuborgs `json:"subOrganizations"`
}

type Data struct {
	Orgs []Organization `json:"orgs"`
}

func main() {
	jsonInput := `{"orgs":[{"id":"56150002","name":"Org1","parentOrganization":{"id":"46409635","name":"ParentOrg"}},{"id":"23456789","name":"Org2","parentOrganization":{"id":"56150002","name":"Org1"}},{"id":"23456788","name":"Org3","parentOrganization":{"id":"56150002","name":"Org1"}},{"id":"1234567","name":"Org4","parentOrganization":{"id":"23456789","name":"Org2"}},{"id":"56150001","name":"Org5","parentOrganization":{"id":"46409635","name":"ParentOrg"}},{"id":"46409635","name":"ParentOrg", "parentOrganization": null}, {"id":"12309876","name":"Another admin Org", "parentOrganization": null}, {"id":"12309877","name":"Org6","parentOrganization":{"id":"12309876","name":"Another admin Org"}},{"id":"23456790","name":"Org7","parentOrganization":{"id":"23456789","name":"Org2"}},{"id":"23456791","name":"Org8","parentOrganization":{"id":"23456788","name":"Org3"}}, {"id":"23456792","name":"Org9","parentOrganization":{"id":"1234567","name":"Org4"}},{"id":"23456793","name":"Org10","parentOrganization":{"id":"23456792","name":"Org9"}}]}`

	var data Data
	var orgsData []OrgSuborgs
	err := json.Unmarshal([]byte(jsonInput), &data)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	orgs := data.Orgs
	var subOrgIds = make(map[string]bool)
	//fmt.Println(reflect.TypeOf(orgs))
	var storeParentOrgIndex = make(map[string]OrgSuborgs)
	for _, v := range orgs {
		var subOrgsList OrgSuborgs
		//var mainOrg OrgSuborgs
		if v.ParentOrganization != nil {
			subOrgIds[v.ID] = true
			subOrgsList.OrgId = v.ID
			subOrgsList.OrgName = v.Name
			recApproch(orgs, &v, &subOrgsList)

			subOrgsList = OrgSuborgs{OrgId: v.ParentOrganization.ID, OrgName: v.ParentOrganization.Name, SubOrg: []OrgSuborgs{subOrgsList}}

			recApproch(orgs, v.ParentOrganization, &subOrgsList)
			//fmt.Println("Final Org Lost", subOrgsList, v.ID)
			if storeParent, ok := storeParentOrgIndex[v.ParentOrganization.ID]; !ok {
				storeParentOrgIndex[v.ParentOrganization.ID] = subOrgsList
			} else {
				storeParent.SubOrg = append(storeParentOrgIndex[v.ParentOrganization.ID].SubOrg, subOrgsList.SubOrg...)
				storeParentOrgIndex[v.ParentOrganization.ID] = storeParent
			}
		} else {
			if _, ok := storeParentOrgIndex[v.ID]; !ok {
				storeParentOrgIndex[v.ID] = OrgSuborgs{OrgId: v.ID, OrgName: v.Name, SubOrg: []OrgSuborgs{}}
			}
		}
	}
	for orgIndex, v := range storeParentOrgIndex {
		if _, ok := subOrgIds[orgIndex]; !ok {
			orgsData = append(orgsData, v)
		}
	}

	jsonData, err := json.Marshal(orgsData)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	//fmt.Println(subOrgIds)
	fmt.Println(string(jsonData))
}

func recApproch(orgs []Organization, currentOrg *Organization, subOrg *OrgSuborgs) {
	if currentOrg.ParentOrganization == nil {
		return
	}
	for _, org := range orgs {
		if org.ParentOrganization != nil && org.ParentOrganization.ID == currentOrg.ID {
			subOrg.SubOrg = append(subOrg.SubOrg, OrgSuborgs{OrgId: org.ID, OrgName: org.Name, SubOrg: []OrgSuborgs{}})
			recApproch(orgs, &org, &subOrg.SubOrg[len(subOrg.SubOrg)-1])
		}
	}

}
