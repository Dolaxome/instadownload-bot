package storylink

import "fmt"

type InstStory struct {
	Url       string
	Timestamp int64
}

type InstUser struct {
	Username string
	Id       int64
	Stories  []InstStory
	Title    string
}

func itemToStories(items []TrayItem) (stories []InstStory) {
	for _, item := range items {
		story := InstStory{
			Timestamp: item.TakenAt,
		}

		if len(item.VideoVersions) > 0 {
			story.Url = item.VideoVersions[0].Url
		} else {
			story.Url = item.ImageVersions2.Candidates[0].Url
		}

		stories = append(stories, story)
	}
	return
}

func GetUnread() (users []InstUser, err error) {
	trays, err := GetReelsTray()
	if err != nil {
		return
	}

	for _, tray := range trays {
		user := InstUser{
			Id:       tray.Id,
			Username: tray.User.Username,
		}
		user.Stories = itemToStories(tray.Items)

		users = append(users, user)
	}
	return
}

func checkUserStories(user *InstUser, c chan int) {
	tray, err := GetUStoriesTray(user.Id)
	if err != nil {
		fmt.Println("failed to fetch" + user.Username)
		c <- 1
		return
	}
	user.Stories = itemToStories(tray.Items)
	c <- 1
}

func GetStoriesAll() (users []InstUser, err error) {
	users, err = GetUnread()
	if err != nil {
		return
	}

	c := make(chan int)
	EmptyUsers := 0
	for index := range users {
		if len(users[index].Stories) == 0 {
			EmptyUsers++
			go checkUserStories(&users[index], c)
		}
	}

	for i := 0; i < EmptyUsers; i++ {
		<-c
	}

	return
}
