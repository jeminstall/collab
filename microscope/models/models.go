package models

import (
	"collab/common/jsutil"
	"fmt"

	"github.com/gopherjs/gopherjs/js"
)

type Deck struct {
	*js.Object           // this is needed for bidirectional data bindings
	Name       string    `js:"name"`
	Periods    []*Period `js:"periods"`
}

func NewDeck() *Deck {
	d := &Deck{Object: jsutil.NewJSObject()}
	d.Periods = []*Period{}
	return d
}

type Period struct {
	*js.Object          // this is needed for bidirectional data bindings
	Title      string   `js:"title"`
	Events     []*Event `js:"events"`
}

func NewPeriod() *Period {
	p := &Period{Object: jsutil.NewJSObject()}
	p.Events = []*Event{}
	return p
}

type Event struct {
	*js.Object          // this is needed for bidirectional data bindings
	Title      string   `js:"title"`
	Scenes     []*Scene `js:"scenes"`
}

func NewEvent() *Event {
	e := &Event{Object: jsutil.NewJSObject()}
	e.Scenes = []*Scene{}
	return e
}

type Scene struct {
	*js.Object        // this is needed for bidirectional data bindings
	Question   string `js:"question"`
	Setting    string `js:"setting"`
	Answer     string `js:"answer"`
}

func NewScene() *Scene {
	return &Scene{Object: jsutil.NewJSObject()}
}

func MockDeck() *Deck {
	d := NewDeck()
	d.Name = "Mock Deck"

	for pi := 0; pi < 20; pi++ {
		p := NewPeriod()
		p.Title = fmt.Sprintf("Period Title %d", pi)
		d.Periods = append(d.Periods, p)

		for ei := 0; ei < 4; ei++ {
			e := NewEvent()
			e.Title = fmt.Sprintf("Event Title %d", ei)
			p.Events = append(p.Events, e)

			for si := 0; si < 2; si++ {
				s := NewScene()
				s.Question = fmt.Sprintf("Why is %d equal to %d?", si, si)
				s.Setting = fmt.Sprintf("Setting for %d people", si)
				s.Answer = fmt.Sprintf("Because it just is equal to %d", si)
				e.Scenes = append(e.Scenes, s)
			}
		}
	}

	return d
}
