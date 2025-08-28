# State

> State is a behavioral design pattern that lets an object alter its behavior when its internal state changes. It appears as if the object changed its class. (Refactoring Guru)

> The state pattern is a behavioral software design pattern that allows an object to alter its behavior when its internal state changes. This pattern is close to the concept of finite-state machines. The state pattern can be interpreted as a strategy pattern, which is able to switch a strategy through invocations of methods defined in the pattern's interface. (Wikipedia)

## Example

### Basic example

```go
package main

import "fmt"

// Other resources
type Song struct {
 title string
}

func NewSong(title string) *Song {
 return &Song{title}
}

// Context
type Player struct {
 state       PlayerState
 Playlist    []*Song
 currentSong int

 ready   PlayerState
 playing PlayerState
 paused  PlayerState
}

func NewPlayer(playlist []*Song) *Player {
 if len(playlist) < 1 {
  panic("Playlist cannot be empty.")
 }
 p := &Player{
  Playlist: playlist,
 }
 p.ready = &ReadyState{p}
 p.paused = &PausedState{p}
 p.playing = &PlayingState{p}
 p.ChangeState(p.ready)
 return p
}

func (p *Player) ChangeState(state PlayerState) {
 p.state = state
 p.state.Report()
}

func (p *Player) AddTrack(song *Song) {
 p.Playlist = append(p.Playlist, song)
}

func (p *Player) CurrentSong() string {
 return p.Playlist[p.currentSong].title
}

// State Interface
type PlayerState interface {
 Play()
 Next()
 Previous()
 Report()
}

// Concrete state - ready
type ReadyState struct {
 player *Player
}

func (s *ReadyState) Play() {
 fmt.Println("[ACTION] Click play")
 s.player.ChangeState(s.player.playing)
}
func (s *ReadyState) Next() {
 fmt.Println("[ACTION] Click next")
 s.Report()
}
func (s *ReadyState) Previous() {
 fmt.Println("[ACTION] Click previous")
 s.Report()
}
func (s *ReadyState) Report() {
 fmt.Println("[STATUS] The player is not reproducing any.")
}

// Concrete state - playing
type PlayingState struct {
 player *Player
}

func (s *PlayingState) Play() {
 fmt.Println("[ACTION] Click play")
 s.player.ChangeState(s.player.paused)
}
func (s *PlayingState) Next() {
 fmt.Println("[ACTION] Click next")
 if len(s.player.Playlist)-1 <= s.player.currentSong {
  fmt.Println("This is the last song on the playlist")
  return
 }

 s.player.currentSong++
 s.Report()
}
func (s *PlayingState) Previous() {
 fmt.Println("[ACTION] Click previous")
 if s.player.currentSong == 0 {
  fmt.Println("This is the first song on the playlist")
  return
 }

 s.player.currentSong--
 s.Report()
}
func (s *PlayingState) Report() {
 fmt.Printf("[STATUS] Playing \"%s\"\n", s.player.CurrentSong())
}

// Concrete state - ready
type PausedState struct {
 player *Player
}

func (s *PausedState) Play() {
 fmt.Println("[ACTION] Click play")
 s.player.ChangeState(s.player.playing)
}
func (s *PausedState) Next() {
 fmt.Println("[ACTION] Click next")
 if len(s.player.Playlist)-1 <= s.player.currentSong {
  fmt.Println("This is the last song on the playlist")
  return
 }

 s.player.currentSong++
 s.player.ChangeState(s.player.playing)
}
func (s *PausedState) Previous() {
 fmt.Println("[ACTION] Click previous")
 if s.player.currentSong == 0 {
  fmt.Println("This is the first song on the playlist")
  return
 }

 s.player.currentSong--
 s.player.ChangeState(s.player.playing)
}
func (s *PausedState) Report() {
 fmt.Printf("[STATUS] Paused \"%s\"\n", s.player.CurrentSong())
}

func main() {
 player := NewPlayer([]*Song{
  NewSong("I Believe in Miracles"),
  NewSong("Zero Zero UFO"),
  NewSong("Don't Bust My Chops"),
 })

 player.state.Play()
 player.state.Next()
 player.state.Play()
 player.state.Previous()
 player.state.Next()

 // Play the remaing songs
 for player.currentSong < len(player.Playlist)-1 {
  player.currentSong++
  player.state.Report()
 }
}
```

### Switch base example

```go
package main

import (
 "bufio"
 "fmt"
 "os"
 "strings"
)

type State int

const (
 Locked State = iota
 Failed
 Unlocked
)

func main() {
 code := "1234"
 state := Locked
 entry := strings.Builder{}

 for {
  switch state {
  case Locked:
   r, _, _ := bufio.NewReader(os.Stdin).ReadRune()
   entry.WriteRune(r)

   if entry.String() == code {
    state = Unlocked
    break
   }

   if strings.Index(code, entry.String()) != 0 {
    state = Failed
   }
  case Failed:
   fmt.Println("FAILED")
   entry.Reset()
   state = Locked
  case Unlocked:
   fmt.Println("UNLOCKED")
   return
  }
 }
}
```
