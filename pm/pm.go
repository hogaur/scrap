package pm

type Pm struct {
	Name          string
	CurrentRole   string
	PreviousRoles string
}

func BuildPmInfo(name string, currentRole string, previousRoles string) Pm {
	return Pm{Name: name, CurrentRole: currentRole, PreviousRoles: previousRoles}
}

func (p Pm) AddPmToList(pms []Pm) []Pm {
	for _, pm := range pms {
		if p.Name == pm.Name {
			return pms
		}
	}
	pms = append(pms, p)
	return pms
}
