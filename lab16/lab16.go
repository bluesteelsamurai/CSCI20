// An adaptation of Tweet JSON ->
// https://developer.twitter.com/en/docs/tweets/data-dictionary/overview/intro-to-tweet-json
package main

// User represents a Twitter user.
type User struct {
	ScreenName string
}

// NewUser returns a User with the given ScreenName.
func NewUser(screenName string) User {
	return User{screenName}
}

// Tweet represents a simplified Twitter tweet.
type Tweet struct {
	Text     string
	Author   User
	Hashtags []string
	Mentions []User
}

// NewTweet returns a tweet with the given text, author, and hashtag.
// The initial hashtag is the first appened to the Tweet's Hashtags.
func NewTweet(text string, author User, hashtag string) *Tweet {
	return &Tweet{text, author, nil, nil}
}

// AddHashtag appends a hashtag to the Tweet's Hashtags.
// If the hashtag is already in Hashtags, it is not appended.
func (t *Tweet) AddHashtag(hashtag string) {
	t.Hashtags = append(t.Hashtags, hashtag)
}

// AddMention appends a User to the Tweet's Mentions.
// If the User is already in Mentions, it is not appended.
func (t *Tweet) AddMention(mention User) {

	for _, usr := range t.Mentions {
		if usr != mention {

		}

	}

}

// ToString returns a string representation of the Movie.
//
// TEMPLATE: author.screen_name
//           text
//           @hashtag1 @hashtag2, ...
//           __ mentions, last mentioned by __
//
// EXAMPLE:  CoderX
//           I am the greatest Go programmer!
//           @goboss @goisgood
//           42 mentions, last mentioned by CoderPlus
func (t *Tweet) ToString() string {
	return ""
}

func main() {

}
