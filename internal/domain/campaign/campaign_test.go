package campaign

import (
	"testing"
	"time"

	"github.com/jaswdr/faker/v2"
	"github.com/stretchr/testify/assert"
)

var (
	name     = "Campaign X"
	content  = "content X"
	contacts = []string{"email@mail.com", "email2@mail.com"}
	fake     = faker.New()
)

func Test_New_Campaign(t *testing.T) {
	assert := assert.New(t)

	campaign, _ := NewCampaign(name, content, contacts)

	assert.Equal(campaign.Name, name)
	assert.Equal(campaign.Content, content)
	assert.Equal(len(campaign.Contacts), len(contacts))
}

func Test_New_Campaign_Id_Not_Nil(t *testing.T) {
	assert := assert.New(t)

	campaign, _ := NewCampaign(name, content, contacts)

	assert.NotNil(campaign.Id)
}

func Test_New_Campaign_CreateOn_Must_Be_Now(t *testing.T) {
	assert := assert.New(t)

	now := time.Now().Add(-time.Minute)

	campaign, _ := NewCampaign(name, content, contacts)

	assert.Greater(campaign.CreatedOn, now)
}

func Test_New_Campaign_Must_Validate_Name_Min(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign("", content, contacts)

	assert.Equal("name is required with min 5", err.Error())
}

func Test_New_Campaign_Must_Validate_Name_Max(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign(fake.Lorem().Text(30), content, contacts)

	assert.Equal("name is required with max 24", err.Error())
}

func Test_New_Campaign_Must_Validate_Content_Min(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign(name, "", contacts)

	assert.Equal("content is required with min 5", err.Error())
}

func Test_New_Campaign_Must_Validate_Content_Max(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign(name, fake.Lorem().Text(1040), contacts)

	assert.Equal("content is required with max 1024", err.Error())
}

func Test_New_Campaign_Must_Validate_Contacts_Min(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign(name, content, nil)

	assert.Equal("contacts is required with min 1", err.Error())
}

func Test_New_Campaign_Must_Validate_Contacts_Email(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign(name, content, []string{"email_invalid"})

	assert.Equal("email is invalid", err.Error())
}
