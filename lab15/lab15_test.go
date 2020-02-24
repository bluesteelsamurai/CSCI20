package main

import "testing"

func makeMovie() *Movie {
	return NewMovie("Godzilla", "action")
}

func TestNewMovie(t *testing.T) {
	m := makeMovie()
	if m.Title != "Godzilla" {
		t.Error("Expected 'Godzilla', got", m.Title)
	}
	if m.Genre != "action" {
		t.Error("Expected 'action', got", m.Genre)
	}
	if m.Likes != 0 {
		t.Error("Expected 0 likes, got", m.Likes)
	}
	if m.Dislikes != 0 {
		t.Error("Expected 0 dislikes, got", m.Likes)
	}
	if m.Views != 0 {
		t.Error("Expected 0 views, got", m.Views)
	}
	if m.GetPercentageLikes() != 0.0 {
		t.Error("Expected 0% likes, got", m.GetPercentageLikes())
	}
	if m.GetPercentageDislikes() != 0.0 {
		t.Error("Expected 0% dislikes, got", m.GetPercentageDislikes())
	}
}

func TestLikeViews(t *testing.T) {
	m := makeMovie()
	m.AddLikeView()
	if m.Views != 1 {
		t.Error("Expected 1 view, got", m.Views)
	}
	if m.Likes != 1 {
		t.Error("Expected 1 like, got", m.Likes)
	}
	if m.GetPercentageLikes() != 100.0 {
		t.Errorf("Expected 100.00%% likes, got %.2f%%", m.GetPercentageLikes())
	}
}

func TestDislikeViews(t *testing.T) {
	m := makeMovie()
	m.AddDislikeView()
	if m.Views != 1 {
		t.Error("Expected 1 view, got", m.Views)
	}
	if m.Dislikes != 1 {
		t.Error("Expected 1 like, got", m.Dislikes)
	}
	if m.GetPercentageDislikes() != 100.0 {
		t.Errorf("Expected 100.00%% dislikes, got %.2f%%", m.GetPercentageDislikes())
	}
}

func TestLikesAndDislikes(t *testing.T) {
	m := makeMovie()
	m.AddLikeView()
	m.AddDislikeView()
	m.AddLikeView()
	m.AddDislikeView()
	m.AddLikeView()

	if m.Views != 5 {
		t.Error("Expected 5 views, got", m.Views)
	}
	if m.Likes != 3 {
		t.Error("Expected 3 likes, got", m.Likes)
	}
	if m.Dislikes != 2 {
		t.Error("Expected 2 dislikes, got", m.Dislikes)
	}
	if m.GetPercentageLikes() != 60.0 {
		t.Errorf("Expected 60.00%% likes, got %.2f%%", m.GetPercentageLikes())
	}
	if m.GetPercentageDislikes() != 40.0 {
		t.Errorf("Expected 40.00%% dislikes, got %.2f%%", m.GetPercentageDislikes())
	}
}

func TestToString(t *testing.T) {
	m := makeMovie()
	m.AddLikeView()
	m.AddLikeView()
	m.AddLikeView()
	m.AddDislikeView()

	mStr := m.ToString()
	expStr := "Godzilla, action, 4 views, 75.00% likes"
	if mStr != expStr {
		t.Errorf("Expected '%s', got '%s'", expStr, mStr)
	}
}
