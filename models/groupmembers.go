package models

type GroupMembers struct {
	GroupMemberId            uint   `gorm:"primarykey;column:group_member_id" json:"group_member_id"`
	GroupId                  string `gorm:"size:100;not null" json:"group_id"`
	UserId                   string `gorm:"size:100;not null" json:"user_id"`
	GroupMemberName          string `gorom:"size:100" json:"group_member_name"`
	GroupMemberEmailPassword string `gorom:"size:100" json:"group_member_email_password"`
}

func (g *GroupMembers) SaveGroupMember() (*GroupMembers, error) {
	var err error
	err = DB.Create(&g).Error
	if err != nil {
		return &GroupMembers{}, err
	}
	return g, nil
}
