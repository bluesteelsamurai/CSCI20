package main

import (
	"fmt"
	"testing"
)

func makeUser(screenName string) User {
	return NewUser(screenName)
}

func makeTweet() *Tweet {
	u := makeUser("CoderX")
	t := NewTweet("I am the greatest Go programmer!", u, "goboss")
	return t
}

func sameHashtags(sliceOne, sliceTwo []string) bool {
	if len(sliceOne) != len(sliceTwo) {
		return false
	}

	for i, ht := range sliceOne {
		if ht != sliceTwo[i] {
			return false
		}
	}

	return true
}

func sameMentions(sliceOne, sliceTwo []User) bool {
	if len(sliceOne) != len(sliceTwo) {
		return false
	}

	for i, u := range sliceOne {
		if u.ScreenName != sliceTwo[i].ScreenName {
			return false
		}
	}

	return true
}

func TestUser(t *testing.T) {
	u := makeUser("TestUser")
	if u.ScreenName != "TestUser" {
		t.Error("Expected screen name 'TestUser', got", u.ScreenName)
	}
}

func TestBasicTweet(t *testing.T) {
	tweet := makeTweet()
	expText := "I am the greatest Go programmer!"
	if tweet.Text != expText {
		t.Errorf("Expected '%s', got '%s'\n", expText, tweet.Text)
	}
	if tweet.Author.ScreenName != "CoderX" {
		t.Error("Expected author 'CoderX', got", tweet.Author.ScreenName)
	}
	expHashtags := []string{"goboss"}
	if !sameHashtags(tweet.Hashtags, expHashtags) {
		t.Errorf("Expected hashtags %v, got %v\n", expHashtags, tweet.Hashtags)
	}
	expMentions := []User{}
	if !sameMentions(tweet.Mentions, expMentions) {
		t.Errorf("Expected mentions %v, got %v\n", expMentions, tweet.Mentions)
	}
}

func TestDuplicateHashtags(t *testing.T) {
	tweet := makeTweet()

	tweet.AddHashtag("goboss")
	expHashtags := []string{"goboss"}
	if !sameHashtags(tweet.Hashtags, expHashtags) {
		t.Errorf("Expected hashtags %v, got %v\n", expHashtags, tweet.Hashtags)
	}
}

func TestDuplicateMentions(t *testing.T) {
	tweet := makeTweet()

	u := makeUser("TestUser")
	tweet.AddMention(u)
	tweet.AddMention(u)
	expMentions := []User{u}
	if !sameMentions(tweet.Mentions, expMentions) {
		t.Errorf("Expected mentions %v, got %v\n", expMentions, tweet.Mentions)
	}
}

func TestTweetWithMention(t *testing.T) {
	tweet := makeTweet()

	tweet.AddHashtag("goisgood")
	expHashtags := []string{"goboss", "goisgood"}
	if !sameHashtags(tweet.Hashtags, expHashtags) {
		t.Errorf("Expected hashtags %v, got %v\n", expHashtags, tweet.Hashtags)
	}

	u := makeUser("CoderPlus")
	tweet.AddMention(u)
	expMentions := []User{u}
	if !sameMentions(tweet.Mentions, expMentions) {
		t.Errorf("Expected mentions %v, got %v\n", expMentions, tweet.Mentions)
	}

	expStr := "CoderX\nI am the greatest Go programmer!\n@goboss @goisgood \n1 mention(s), last mentioned by CoderPlus"
	if tweet.ToString() != expStr {
		t.Errorf("Expected |%s|, got |%s|\n", expStr, tweet.ToString())
	}
}

func TestTweetWithManyMentions(t *testing.T) {
	tweet := makeTweet()

	tweet.AddHashtag("goisfun")
	tweet.AddHashtag("goisgood")
	expHashtags := []string{"goboss", "goisfun", "goisgood"}
	if !sameHashtags(tweet.Hashtags, expHashtags) {
		t.Errorf("Expected hashtags %v, got %v\n", expHashtags, tweet.Hashtags)
	}

	var u User
	var expMentions []User
	for i := 1; i < 42; i++ {
		u = makeUser(fmt.Sprintf("Coder%d", i))
		expMentions = append(expMentions, u)
		tweet.AddMention(u)
	}

	u = makeUser("CoderPlus")
	tweet.AddMention(u)
	expMentions = append(expMentions, u)
	if !sameMentions(tweet.Mentions, expMentions) {
		t.Errorf("Expected mentions %v, got %v\n", expMentions, tweet.Mentions)
	}

	expStr := "CoderX\nI am the greatest Go programmer!\n@goboss @goisfun @goisgood \n42 mention(s), last mentioned by CoderPlus"
	if tweet.ToString() != expStr {
		t.Errorf("Expected |%s|, got |%s|\n", expStr, tweet.ToString())
	}
}
