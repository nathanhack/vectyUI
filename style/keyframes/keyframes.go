package keyframes

import (
	"github.com/nathanhack/vectyUI/internal"
	"github.com/nathanhack/vectyUI/style/animationName"
	"strconv"
	"strings"
	"syscall/js"
)

var keyframeHistory = map[animationName.Type]map[string]bool{}

type KeyframeStyler interface {
	AddTo(m map[string]string)
}

var count = 0

type Keyframe struct {
	Sel    Selector
	Styles []KeyframeStyler
}

type Selector string

const (
	From Selector = "from"
	To   Selector = "to"
)

func Percent(percent interface{}) Selector {
	return Selector(internal.Stringify(percent, "%"))
}

//Keyframes creates @keyframes used in animation
// NOTE currently keyframes are only inserted
// old keyframes are not removed and are intended to be
// create at the start of the app and never changed.
// However, if keyframes are used in a dynamic
// way (changing after startup) then
// there is a possible out of memory issue.
type Keyframes struct {
	Name         animationName.Type
	Frames       []Keyframe
	genKeyframes []string
}

func (k *Keyframes) Mount() {
	//instead of styles we load the keyframes!
	if k.Name == "" {
		k.generateName()
	}

	k.generateKeyframes()
	k.addKeyframesRules()
}

func (k *Keyframes) generateName() {
	//we do our best to make a unique name
	k.Name = animationName.Type("vecytUIKeyframes" + strconv.Itoa(count))
	count++
}

func (k *Keyframes) generateKeyframes() {
	//we don't recreate the keyframes
	// if the have already been created
	if len(k.genKeyframes) != 0 {
		return
	}

	k.genKeyframes = make([]string, 0)
	for _, prefix := range []string{"-webit-", "-moz-", "-ms-", "-o-", ""} {
		sb := strings.Builder{}
		sb.WriteString("@")
		sb.WriteString(prefix)
		sb.WriteString("keyframes ")
		sb.WriteString(string(k.Name))
		sb.WriteString(" { ")

		for _, f := range k.Frames {
			sb.WriteString(string(f.Sel))
			sb.WriteString(" { ")
			styleMap := make(map[string]string)
			for _, style := range f.Styles {
				style.AddTo(styleMap)
			}
			for k, s := range styleMap {
				sb.WriteString(k + ":" + s + ";")
			}

			sb.WriteString(" } ")
		}

		sb.WriteString(" }")
		k.genKeyframes = append(k.genKeyframes, sb.String())
	}
}

func (k *Keyframes) addKeyframesRules() {
	doc := js.Global().Get("document")
	// styleSheets = doc.Get("styleSheets")

	for _, keyframe := range k.genKeyframes {
		//for each keyframe we check our history
		// if it's new we add it
		// this could occur if a component containing Keyframes{} was remounted
		// we just mitigate dups by keeping a history of our own
		// note this doesn't fix the dup issue if the keyframe was changed between remounted

		if _, has := keyframeHistory[k.Name]; !has {
			keyframeHistory[k.Name] = make(map[string]bool)
		}

		if _, has := keyframeHistory[k.Name][keyframe]; !has {
			s := doc.Call("createElement", "style")
			s.Set("innerHTML", keyframe)
			doc.Call("getElementsByTagName", "head").Index(0).Call("appendChild", s)
		}
		// instead of always inserting like above
		// if we know the 0th index of stylesheets is actually
		// a style and not a link then we just use
		// the insertRule method to add the keyframes
		//styleSheets.Index(0).Call("insertRule", k, 0)
	}
}

// the doesn't work. using .Get() explodes and is not recover()'able yet in wasm
//func (k *Keyframes) deleteOldRules() {
//	//we need to find the keyframes in the styles and remove it
//	styleSheets := js.Global().Get("document").Get("styleSheets")
//
//	if !styleSheets.IsNull() {
//		for i := 0; i < styleSheets.Length(); i++ {
//			styleSheet := styleSheets.Index(i)
//			if styleSheet.IsNull() {
//				continue
//			}
//			rules := styleSheet.Get("cssRules")
//			if  rules.IsNull() {
//				continue
//			}
//			for j := 0; j < rules.Length(); j++ {
//				rule := rules.Index(j)
//				if rule.IsNull() {
//					continue
//				}
//				cssType := rule.Get("type")
//				if  cssType.IsNull() || cssType.Int() != 7 {
//					continue
//				}
//
//				name := rule.Get("name")
//				if name.IsNull() || name.String() != string(k.Name) {
//					continue
//				}
//
//				//well we found it, time to remove it
//				rules.Call("deleteRule", j)
//			}
//		}
//	}
//}

// recover() currently doesn't work for wasm
//func styleSheetHasRules(index int) (has bool) {
//	has = false
//	defer func() {
//		if r := recover(); r != nil {
//			fmt.Println("Recovered in f", r)
//		}
//	}()
//	doc := js.Global().Get("document")
//	styleSheets := doc.Get("styleSheets")
//	sheet := styleSheets.Index(index)
//	sheet.Get("cssRules")
//	return true
//}
