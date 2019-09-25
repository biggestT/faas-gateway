// omg it rhymes
package dict

import (
  "container/list"
  "math/rand"
  "bufio"
  "strings"
  "regexp"
  "os"
)

type Sound struct {
  alternatives *list.List
  current *list.Element
}

func NewSound(word string) *Sound {
  sound := new(Sound)
  sound.alternatives = list.New()
  sound.alternatives.PushFront(word)
  sound.current = sound.alternatives.Front()
  return sound
}

type rhymeMap = map[string]*Sound

type Dict struct {
  rhymes rhymeMap
  verbs []string
  numVerbs int
  rhymeSoundExp *regexp.Regexp
}

func (d *Dict) getSyllable(word string) string {
  return d.rhymeSoundExp.FindString(word) 
}

func (d *Dict) NextRhyme(word string) string {
  syll := d.getSyllable(word)
  sound := d.rhymes[syll]
  for sound == nil {
    syll = syll[1:]
    sound = d.rhymes[syll]
  }
  rhyme := sound.current.Value.(string)
  next := sound.current.Next()
  if next == nil {
    next = sound.alternatives.Front()
  }
  sound.current = next 
  verb := d.verbs[rand.Intn(d.numVerbs)]
  sentence := []string{word, verb, rhyme}
  return strings.Join(sentence, " ")
}

func parseLine(txt string) (string, string) {
  parts := strings.Split(txt, " ")
  return parts[0], parts[1]
}

func New(file string) *Dict {
  f, _ := os.Open(file)
  dict := new(Dict)
  dict.rhymes = make(rhymeMap)
	dict.rhymeSoundExp, _ = regexp.Compile(`[aeiouyåäö][a-zåäö]{1,3}$`)
  scanner := bufio.NewScanner(f)
  for scanner.Scan() {
    txt := scanner.Text()
    word, class := parseLine(txt)
    if class == "nn" {
      syllable := dict.getSyllable(word) 
      sound, exists := dict.rhymes[syllable]
      if !exists {
        dict.rhymes[syllable] = NewSound(word)
      } else {
        sound.alternatives.PushBack(word)
      }
    }
    if class == "vb" {
      dict.verbs = append(dict.verbs, word)
    }
  }
  // this won't change
  dict.numVerbs = len(dict.verbs)
  return dict
}
