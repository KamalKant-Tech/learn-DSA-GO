package main

import (
	"encoding/json"
	"fmt"
)

type Organization struct {
	ID                 string        `json:"id"`
	Name               string        `json:"name"`
	ParentOrganization *Organization `json:"parentOrganization,omitempty"`
	//SubOrganizations   []Organization `json:"subOrganizations,omitempty"`
}

type OrgSuborgs struct {
	OrgId   string        `json:"organizationId"`
	OrgName string        `json:"organizationName"`
	SubOrg  []*OrgSuborgs `json:"subOrganizations"`
}

type Data struct {
	Orgs []Organization `json:"orgs"`
}

func main() {
	jsonInput := `{"orgs":[{"id":"56150002","name":"Org1","parentOrganization":{"id":"46409635","name":"ParentOrg"}},{"id":"23456789","name":"Org2","parentOrganization":{"id":"56150002","name":"Org1"}},{"id":"23456788","name":"Org3","parentOrganization":{"id":"56150002","name":"Org1"}},{"id":"1234567","name":"Org4","parentOrganization":{"id":"23456789","name":"Org2"}},{"id":"56150001","name":"Org5","parentOrganization":{"id":"46409635","name":"ParentOrg"}},{"id":"46409635","name":"ParentOrg", "parentOrganization": null}, {"id":"12309876","name":"Another admin Org", "parentOrganization": null}, {"id":"12309877","name":"Org6","parentOrganization":{"id":"12309876","name":"Another admin Org"}},{"id":"23456790","name":"Org7","parentOrganization":{"id":"23456789","name":"Org2"}},{"id":"23456791","name":"Org8","parentOrganization":{"id":"23456788","name":"Org3"}}, {"id":"23456792","name":"Org9","parentOrganization":{"id":"1234567","name":"Org4"}},{"id":"23456793","name":"Org10","parentOrganization":{"id":"23456792","name":"Org9"}}, {"id":"23456794","name":"Org11","parentOrganization":{"id":"23456791","name":"Org8"}}, {"id":"23456795","name":"New added org"}]}`

	var data Data
	err := json.Unmarshal([]byte(jsonInput), &data)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	orgsData := generateOrgSuborgs(data.Orgs)
	jsonData, err := json.MarshalIndent(orgsData, "", "   ")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(string(jsonData))
}

func generateOrgSuborgs(orgs []Organization) []*OrgSuborgs {
	orgMap := make(map[string]*OrgSuborgs)
	subOrgIds := make(map[string]bool)

	for _, org := range orgs {
		if org.ParentOrganization == nil {
			if _, ok := orgMap[org.ID]; !ok {
				orgMap[org.ID] = &OrgSuborgs{
					OrgId:   org.ID,
					OrgName: org.Name,
					SubOrg:  []*OrgSuborgs{},
				}
			}
			continue
		}

		subOrgIds[org.ID] = true

		subOrg, ok := orgMap[org.ID]
		if !ok {
			subOrg = &OrgSuborgs{
				OrgId:   org.ID,
				OrgName: org.Name,
				SubOrg:  []*OrgSuborgs{},
			}
			orgMap[org.ID] = subOrg
		}

		parentOrg := org.ParentOrganization
		if parentOrg == nil {
			continue
		}

		parentSubOrg, ok := orgMap[parentOrg.ID]
		if !ok {
			parentSubOrg = &OrgSuborgs{
				OrgId:   parentOrg.ID,
				OrgName: parentOrg.Name,
				SubOrg:  []*OrgSuborgs{subOrg},
			}
			orgMap[parentOrg.ID] = parentSubOrg
		} else {
			parentSubOrg.SubOrg = append(parentSubOrg.SubOrg, subOrg)
		}
	}

	rootOrgs := []*OrgSuborgs{}
	for _, org := range orgs {
		if org.ParentOrganization == nil {
			rootOrgs = append(rootOrgs, orgMap[org.ID])
		}
	}

	return rootOrgs
}
