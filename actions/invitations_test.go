package actions

import "github.com/invitation/models"

func (as *ActionSuite) Test_InvitationsResource_List() {
	u := &models.User{}
	err := as.DB.Where("email = ?", "sonja@example.com").First(u)
	as.Session.Set("current_user_id", u.ID)
	as.NoError(err)

	res := as.HTML("/invitations").Get()
	as.Equal(res.Code, 200)
	as.Contains(res.Body.String(), "Show")
}

func (as *ActionSuite) Test_InvitationsResource_Show() {
	u := &models.User{}
	err := as.DB.Eager().Where("email = ?", "sonja@example.com").First(u)
	as.Session.Set("current_user_id", u.ID)
	as.NoError(err)
	i := u.Invitations[0].ID

	res := as.HTML("/invitations/" + i.String()).Get()
	as.Equal(res.Code, 200)
	as.Contains(res.Body.String(), "Sie sind herzlich eingeladen!")
}

func (as *ActionSuite) Test_InvitationsResource_New() {
	u := &models.User{}
	err := as.DB.Where("email = ?", "sonja@example.com").First(u)
	as.Session.Set("current_user_id", u.ID)
	as.NoError(err)
}

func (as *ActionSuite) Test_InvitationsResource_Create() {
	u := &models.User{}
	err := as.DB.Where("email = ?", "sonja@example.com").First(u)
	as.Session.Set("current_user_id", u.ID)
	as.NoError(err)
}

func (as *ActionSuite) Test_InvitationsResource_Edit() {
	u := &models.User{}
	err := as.DB.Where("email = ?", "sonja@example.com").First(u)
	as.Session.Set("current_user_id", u.ID)
	as.NoError(err)
}

func (as *ActionSuite) Test_InvitationsResource_Update() {
	u := &models.User{}
	err := as.DB.Where("email = ?", "sonja@example.com").First(u)
	as.Session.Set("current_user_id", u.ID)
	as.NoError(err)
}

func (as *ActionSuite) Test_InvitationsResource_Destroy() {
	u := &models.User{}
	err := as.DB.Where("email = ?", "sonja@example.com").First(u)
	as.Session.Set("current_user_id", u.ID)
	as.NoError(err)
}
