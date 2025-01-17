// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

/*
 * Cadence - The resource-oriented smart contract programming language
 *
 * Copyright 2019-2021 Dapper Labs, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * Based on https://github.com/wk8/go-ordered-map, Copyright Jean Rougé
 *
 */

package orderedmap

import "container/list"

// StringFruitOrderedMap
//
type StringFruitOrderedMap struct {
	pairs map[string]*StringFruitPair
	list  *list.List
}

// NewStringFruitOrderedMap creates a new StringFruitOrderedMap.
func NewStringFruitOrderedMap() *StringFruitOrderedMap {
	return &StringFruitOrderedMap{
		pairs: make(map[string]*StringFruitPair),
		list:  list.New(),
	}
}

// Clear removes all entries from this ordered map.
func (om *StringFruitOrderedMap) Clear() {
	om.list.Init()
	// NOTE: Range over map is safe, as it is only used to delete entries
	for key := range om.pairs { //nolint:maprangecheck
		delete(om.pairs, key)
	}
}

// Get returns the value associated with the given key.
// Returns nil if not found.
// The second return value indicates if the key is present in the map.
func (om *StringFruitOrderedMap) Get(key string) (result *Fruit, present bool) {
	var pair *StringFruitPair
	if pair, present = om.pairs[key]; present {
		return pair.Value, present
	}
	return
}

// GetPair returns the key-value pair associated with the given key.
// Returns nil if not found.
func (om *StringFruitOrderedMap) GetPair(key string) *StringFruitPair {
	return om.pairs[key]
}

// Set sets the key-value pair, and returns what `Get` would have returned
// on that key prior to the call to `Set`.
func (om *StringFruitOrderedMap) Set(key string, value *Fruit) (oldValue *Fruit, present bool) {
	var pair *StringFruitPair
	if pair, present = om.pairs[key]; present {
		oldValue = pair.Value
		pair.Value = value
		return
	}

	pair = &StringFruitPair{
		Key:   key,
		Value: value,
	}
	pair.element = om.list.PushBack(pair)
	om.pairs[key] = pair

	return
}

// Delete removes the key-value pair, and returns what `Get` would have returned
// on that key prior to the call to `Delete`.
func (om *StringFruitOrderedMap) Delete(key string) (oldValue *Fruit, present bool) {
	var pair *StringFruitPair
	pair, present = om.pairs[key]
	if !present {
		return
	}

	om.list.Remove(pair.element)
	delete(om.pairs, key)
	oldValue = pair.Value

	return
}

// Len returns the length of the ordered map.
func (om *StringFruitOrderedMap) Len() int {
	return len(om.pairs)
}

// Oldest returns a pointer to the oldest pair.
func (om *StringFruitOrderedMap) Oldest() *StringFruitPair {
	return listElementToStringFruitPair(om.list.Front())
}

// Newest returns a pointer to the newest pair.
func (om *StringFruitOrderedMap) Newest() *StringFruitPair {
	return listElementToStringFruitPair(om.list.Back())
}

// Foreach iterates over the entries of the map in the insertion order, and invokes
// the provided function for each key-value pair.
func (om *StringFruitOrderedMap) Foreach(f func(key string, value *Fruit)) {
	for pair := om.Oldest(); pair != nil; pair = pair.Next() {
		f(pair.Key, pair.Value)
	}
}

// StringFruitPair
//
type StringFruitPair struct {
	Key   string
	Value *Fruit

	element *list.Element
}

// Next returns a pointer to the next pair.
func (p *StringFruitPair) Next() *StringFruitPair {
	return listElementToStringFruitPair(p.element.Next())
}

// Prev returns a pointer to the previous pair.
func (p *StringFruitPair) Prev() *StringFruitPair {
	return listElementToStringFruitPair(p.element.Prev())
}

func listElementToStringFruitPair(element *list.Element) *StringFruitPair {
	if element == nil {
		return nil
	}
	return element.Value.(*StringFruitPair)
}
