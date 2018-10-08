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
	Mailtext   string
	Salutation int
	Name0      string
	Gender0    int
	Mail0      string
	GuestCount string
	Mail1      string
	Name1      string
	Gender1    int
	Mail2      string
	Name2      string
	Gender2    int
}

func (as *ActionSuite) Test_InvitationsResource_Create() {
	as.LoadFixture("Test data")
	u := &models.User{}
	err := as.DB.Where("email = ?", "sonja@example.com").First(u)
	as.Session.Set("current_user_id", u.ID)
	as.NoError(err)

	i := &invitationTest{
		Mailtext:   "Sie sind herzlich eingeladen! Mit freundlichen Gruessen",
		Salutation: 2,
		GuestCount: "3",
		Name0:      "Alfred",
		Name1:      "Harald",
		Name2:      "Alex",
		Gender0:    1,
		Gender1:    2,
		Gender2:    3,
		Mail0:      "alfred@example.com",
		Mail1:      "harald@example.com",
		Mail2:      "alex@example.com",
	}

	res := as.HTML("/invitations").Post(i)
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
