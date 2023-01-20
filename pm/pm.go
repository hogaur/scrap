package pm

import "strings"

type TopPm struct {
	Name          string
	CurrentRole   CurrentRole
	PreviousRoles []string
}

type CurrentRole struct {
	Position string
	Detail   string
	Company  string
}

func NewTopPm(name string, currentRole string, previousRoles string) TopPm {
	position, detail, company := currentRoleDissection(currentRole)
	return TopPm{
		Name:          name,
		CurrentRole:   CurrentRole{Position: position, Detail: detail, Company: company},
		PreviousRoles: strings.Split(previousRoles, ",")}
}

func currentRoleDissection(currentRole string) (position, detail, company string) {
	currentRoleArray := strings.Split(currentRole, ",")
	currentRoleSize := len(currentRoleArray)
	company = currentRoleArray[currentRoleSize-1]
	position = currentRoleArray[0]
	detail = ""
	if currentRoleSize > 2 {
		detail = currentRoleArray[1]
	}
	return position, detail, company
}

func (p TopPm) AddPmToList(pms []TopPm) []TopPm {
	for _, pm := range pms {
		if p.Name == pm.Name {
			return pms
		}
	}
	pms = append(pms, p)
	return pms
}
