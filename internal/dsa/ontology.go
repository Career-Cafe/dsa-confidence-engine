package dsa

import (
	"strings"
	"sync"
)

type Ontology struct {
	mu           sync.RWMutex
	concepts     map[string]*Concept
	aliasIndex   map[string]*Concept
	keywordIndex map[string][]*Concept
	categories   map[string][]*Concept
}

func NewOntology() *Ontology {
	return &Ontology{
		concepts:     make(map[string]*Concept),
		aliasIndex:   make(map[string]*Concept),
		keywordIndex: make(map[string][]*Concept),
		categories:   make(map[string][]*Concept),
	}
}

func (o *Ontology) AddConcept(c *Concept) {
	o.mu.Lock()
	defer o.mu.Unlock()

	o.concepts[c.ID] = c

	normalizedID := NormalizeText(c.ID)
	o.aliasIndex[normalizedID] = c
	normalizedName := NormalizeText(c.Name)
	o.aliasIndex[normalizedName] = c

	for _, alias := range c.Aliases {
		norm := NormalizeText(alias)
		if norm != "" {
			o.aliasIndex[norm] = c
		}
	}

	for _, kw := range c.Keywords {
		norm := NormalizeText(kw)
		if norm != "" {
			o.keywordIndex[norm] = append(o.keywordIndex[norm], c)
		}
	}

	cat := c.Category
	if cat == "" {
		cat = c.Parent
	}
	if cat != "" {
		o.categories[cat] = append(o.categories[cat], c)
	}
}

func (o *Ontology) FindConcept(id string) (*Concept, bool) {
	o.mu.RLock()
	defer o.mu.RUnlock()
	c, ok := o.concepts[id]
	return c, ok
}

func (o *Ontology) FindConceptByAlias(alias string) (*Concept, bool) {
	o.mu.RLock()
	defer o.mu.RUnlock()
	norm := NormalizeText(alias)
	c, ok := o.aliasIndex[norm]
	return c, ok
}

func (o *Ontology) IsAncestor(ancestorID, descendantID string) bool {
	o.mu.RLock()
	defer o.mu.RUnlock()

	curr := descendantID
	visited := make(map[string]bool)
	for curr != "" {
		if visited[curr] {
			break
		}
		visited[curr] = true
		if curr == ancestorID {
			return true
		}
		c, ok := o.concepts[curr]
		if !ok {
			break
		}
		if c.Parent == ancestorID {
			return true
		}
		curr = c.Parent
	}
	return false
}

func (o *Ontology) GetLineage(id string) []string {
	o.mu.RLock()
	defer o.mu.RUnlock()

	var lineage []string
	curr := id
	visited := make(map[string]bool)
	for curr != "" {
		if visited[curr] {
			break
		}
		visited[curr] = true
		lineage = append(lineage, curr)
		c, ok := o.concepts[curr]
		if !ok {
			break
		}
		curr = c.Parent
	}
	return lineage
}

func (o *Ontology) AllConcepts() []*Concept {
	o.mu.RLock()
	defer o.mu.RUnlock()

	list := make([]*Concept, 0, len(o.concepts))
	for _, c := range o.concepts {
		list = append(list, c)
	}
	return list
}

func (o *Ontology) AllAliases() map[string]*Concept {
	o.mu.RLock()
	defer o.mu.RUnlock()

	aliases := make(map[string]*Concept, len(o.aliasIndex))
	for k, v := range o.aliasIndex {
		aliases[k] = v
	}
	return aliases
}

func NormalizeText(s string) string {
	s = strings.ToLower(s)
	s = strings.ReplaceAll(s, "_", " ")
	s = strings.ReplaceAll(s, "-", " ")
	s = strings.ReplaceAll(s, "/", " ")
	var b strings.Builder
	for _, r := range s {
		if (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') || r == ' ' {
			b.WriteRune(r)
		}
	}
	fields := strings.Fields(b.String())
	return strings.Join(fields, " ")
}
