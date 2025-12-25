package models

type Groupslist struct {
	GroupId         uint   `gorm:"primarykey;column:group_id" json:"group_id"`
	GroupName       string `gorm:"size:100;not null" json:"group_name"`
	CreatedByUserId string `gorm:"size:100;not null" json:"created_by_user_id"`
	Description     string `gorom:"size:150;"json:"description"`
}

func (g *Groupslist) SaveGroup() (*Groupslist, error) {
	var err error
	err = DB.Create(&g).Error
	if err != nil {
		return &Groupslist{}, err
	}
	return g, nil
}

func GetGroupsCreatedByUserID(userID uint) ([]Groupslist, error) {
	var groups []Groupslist

	// Use Where() to filter by the CreatedByUserId column
	result := DB.Where("created_by_user_id = ?", userID).Find(&groups)

	if result.Error != nil {
		return nil, result.Error
	}

	// Handle case where user is found but created no groups (returns empty slice, not an error)
	if len(groups) == 0 {
		return groups, nil // You might return an error here if required, but nil is standard for "found nothing"
	}

	return groups, nil
}
