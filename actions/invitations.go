package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
	"github.com/invitation/models"
	"github.com/pkg/errors"
)

// This file is generated by Buffalo. It offers a basic structure for
// adding, editing and deleting a page. If your model is more
// complex or you need more than the basic implementation you need to
// edit this file.

// Following naming logic is implemented in Buffalo:
// Model: Singular (Invitation)
// DB Table: Plural (invitations)
// Resource: Plural (Invitations)
// Path: Plural (/invitations)
// View Template Folder: Plural (/templates/invitations/)

// InvitationsResource is the resource for the Invitation model
type InvitationsResource struct {
	buffalo.Resource
}

// List gets all Invitations. This function is mapped to the path
// GET /invitations
func (v InvitationsResource) List(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}
	invitations := &models.Invitations{}
	u := c.Value("current_user").(*models.User)

	// Paginate results. Params "page" and "per_page" control pagination.
	// Default values are "page=1" and "per_page=20".
	q := tx.PaginateFromParams(c.Params())
	// Retrieve all Invitations from the DB
	if err := q.Where("userid = ?", u.ID).All(invitations); err != nil {
		return errors.WithStack(err)
	}

	// Add the paginator to the context so it can be used in the template.
	c.Set("pagination", q.Paginator)

	return c.Render(200, r.Auto(c, invitations))
}

// Show gets the data for one Invitation. This function is mapped to
// the path GET /invitations/{invitation_id}
func (v InvitationsResource) Show(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	// Allocate an empty Invitation
	invitation := &models.Invitation{}

	// To find the Invitation the parameter invitation_id is used.
	if err := tx.Find(invitation, c.Param("invitation_id")); err != nil {
		return c.Error(404, err)
	}

	return c.Render(200, r.Auto(c, invitation))
}

// New renders the form for creating a new Invitation.
// This function is mapped to the path GET /invitations/new
func (v InvitationsResource) New(c buffalo.Context) error {
	return c.Render(200, r.Auto(c, &models.Invitation{}))
}

// Create adds a Invitation to the DB. This function is mapped to the
// path POST /invitations
func (v InvitationsResource) Create(c buffalo.Context) error {
	// Allocate an empty Invitation
	invitation := &models.Invitation{}
	u := c.Value("current_user").(*models.User)
	invitation.UserID = u.ID
	// Bind invitation to the html form elements
	if err := c.Bind(invitation); err != nil {
		return errors.WithStack(err)
	}
	invitation.Mailtext += c.Params().Get("name0")

	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	// Validate the data from the html form
	verrs, err := tx.ValidateAndCreate(invitation)
	if err != nil {
		return errors.WithStack(err)
	}

	if verrs.HasAny() {
		// Make the errors available inside the html template
		c.Set("errors", verrs)

		// Render again the new.html template that the user can
		// correct the input.
		return c.Render(422, r.Auto(c, invitation))
	}

	// If there are no errors set a success message
	c.Flash().Add("success", "Invitation was created successfully")

	// and redirect to the invitations index page
	return c.Render(201, r.Auto(c, invitation))
}

// Edit renders a edit form for a Invitation. This function is
// mapped to the path GET /invitations/{invitation_id}/edit
func (v InvitationsResource) Edit(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	// Allocate an empty Invitation
	invitation := &models.Invitation{}

	if err := tx.Find(invitation, c.Param("invitation_id")); err != nil {
		return c.Error(404, err)
	}

	return c.Render(200, r.Auto(c, invitation))
}

// Update changes a Invitation in the DB. This function is mapped to
// the path PUT /invitations/{invitation_id}
func (v InvitationsResource) Update(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	// Allocate an empty Invitation
	invitation := &models.Invitation{}

	if err := tx.Find(invitation, c.Param("invitation_id")); err != nil {
		return c.Error(404, err)
	}

	// Bind Invitation to the html form elements
	if err := c.Bind(invitation); err != nil {
		return errors.WithStack(err)
	}

	verrs, err := tx.ValidateAndUpdate(invitation)
	if err != nil {
		return errors.WithStack(err)
	}

	if verrs.HasAny() {
		// Make the errors available inside the html template
		c.Set("errors", verrs)

		// Render again the edit.html template that the user can
		// correct the input.
		return c.Render(422, r.Auto(c, invitation))
	}

	// If there are no errors set a success message
	c.Flash().Add("success", "Invitation was updated successfully")

	// and redirect to the invitations index page
	return c.Render(200, r.Auto(c, invitation))
}

// Destroy deletes a Invitation from the DB. This function is mapped
// to the path DELETE /invitations/{invitation_id}
func (v InvitationsResource) Destroy(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	// Allocate an empty Invitation
	invitation := &models.Invitation{}

	// To find the Invitation the parameter invitation_id is used.
	if err := tx.Find(invitation, c.Param("invitation_id")); err != nil {
		return c.Error(404, err)
	}

	if err := tx.Destroy(invitation); err != nil {
		return errors.WithStack(err)
	}

	// If there are no errors set a flash message
	c.Flash().Add("success", "Invitation was destroyed successfully")

	// Redirect to the invitations index page
	return c.Render(200, r.Auto(c, invitation))
}
