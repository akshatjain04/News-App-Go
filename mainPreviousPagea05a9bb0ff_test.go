package main

import (
    "testing"
)

type Search struct {
    page int
}

func (s *Search) CurrentPage() int {
    return s.page
}

func (s *Search) PreviousPage() int {
    return s.CurrentPage() - 1
}

func TestPreviousPagea05a9bb0ff(t *testing.T) {
    t.Run("test when current page is greater than 1", func(t *testing.T) {
        search := &Search{page: 2}
        expected := 1
        got := search.PreviousPage()

        if got != expected {
            t.Errorf("Expected previous page to be %d, but got %d", expected, got)
        }
    })

    t.Run("test when current page is 1", func(t *testing.T) {
        search := &Search{page: 1}
        expected := 0
        got := search.PreviousPage()

        if got != expected {
            t.Errorf("Expected previous page to be %d, but got %d", expected, got)
        }
    })

    t.Run("test when current page is 0", func(t *testing.T) {
        search := &Search{page: 0}
        expected := -1
        got := search.PreviousPage()

        if got != expected {
            t.Errorf("Expected previous page to be %d, but got %d", expected, got)
        }
    })
}
