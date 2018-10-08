package actions

import "github.com/invitation/models"

func (as *ActionSuite) Test_InvitationsResource_List() {
	as.LoadFixture("Test data")
	u := &models.User{}
	err := as.DB.Where("email = ?", "sonja@example.com").First(u)
	as.Session.Set("current_user_id", u.ID)
	as.NoError(err)

	res := as.HTML("/invitations").Get()
	as.Equal(res.Code, 200)
	as.Contains(res.Body.String(), "Show")
}

func (as *ActionSuite) Test_InvitationsResource_Show() {
	as.LoadFixture("Test data")
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
	as.LoadFixture("Test data")
	u := &models.User{}
	err := as.DB.Where("email = ?", "sonja@example.com").First(u)
	as.Session.Set("current_user_id", u.ID)
	as.NoError(err)

	res := as.HTML("/invitations/new").Get()
	as.Equal(res.Code, 200)
	as.Contains(res.Body.String(), "Salutation")
}

type invitationTest struct {
	mailtext   string
	salutation int
	name0      string
	gender0    int
	mail0      string
	guestcount int
	mail1      string
	name1      string
	gender1    int
	mail2      string
	name2      string
	gender2    int
}

func (as *ActionSuite) Test_InvitationsResource_Create() {
	as.LoadFixture("Test data")
	u := &models.User{}
	err := as.DB.Where("email = ?", "sonja@example.com").First(u)
	as.Session.Set("current_user_id", u.ID)
	as.NoError(err)

	i := invitationTest{
		mailtext:   "Sie sind herzlich eingeladen! Mit freundlichen Gruessen",
		salutation: 2,
		guestcount: 3,
		name0:      "Alfred",
		name1:      "Harald",
		name2:      "Alex",
		gender0:    1,
		gender1:    2,
		gender2:    3,
		mail0:      "alfred@example.com",
		mail1:      "harald@example.com",
		mail2:      "alex@example.com",
	}

	res := as.HTML("/invitations/new").Post(i)
	as.Contains(res.Body.String(), "Invitation was created successfully")
	count, err := as.DB.Count("invitations")
	as.NoError(err)
	as.Equal(count, 3)

	count, err = as.DB.Count("guests")
	as.NoError(err)
	as.Equal(count, 5)
}

func (as *ActionSuite) Test_InvitationsResource_Edit() {
	as.LoadFixture("Test data")
	u := &models.User{}
	err := as.DB.Where("email = ?", "sonja@example.com").First(u)
	as.Session.Set("current_user_id", u.ID)
	as.NoError(err)
}

func (as *ActionSuite) Test_InvitationsResource_Update() {
	as.LoadFixture("Test data")
	u := &models.User{}
	err := as.DB.Where("email = ?", "sonja@example.com").First(u)
	as.Session.Set("current_user_id", u.ID)
	as.NoError(err)
}

func (as *ActionSuite) Test_InvitationsResource_Destroy() {
	as.LoadFixture("Test data")
	u := &models.User{}
	err := as.DB.Where("email = ?", "sonja@example.com").First(u)
	as.Session.Set("current_user_id", u.ID)
	as.NoError(err)
}
