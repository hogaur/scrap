package pm

type Pm struct {
	Name          string
	CurrentRole   string
	PreviousRoles string
}

func BuildPmInfo(name string, currentRole string, previousRoles string) Pm {
	return Pm{Name: name, CurrentRole: currentRole, PreviousRoles: previousRoles}
}
