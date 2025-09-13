package main

import (
	"gorm.io/gorm"
)

type TLDRDBCached struct {
	provider TLDRProvider
	db       *gorm.DB
}

func (t *TLDRDBCached) Retrieve(key string) string {
	// First, check if the key exists in cache
	var entity TLDREntity
	result := t.db.Where("key = ?", key).First(&entity)

	if result.Error == nil {
		// Found in cache, return cached value
		return entity.Val
	}

	// Not found in cache, get from provider
	value := t.provider.Retrieve(key)

	// Store in cache for future use
	entity = TLDREntity{
		Key: key,
		Val: value,
	}
	t.db.Create(&entity)

	return value
}

func (t *TLDRDBCached) List() []string {
	// Get all available items from provider
	allItems := t.provider.List()

	// Get all cached keys from database
	var cachedEntities []TLDREntity
	t.db.Find(&cachedEntities)

	// Create a map of cached keys for quick lookup
	cachedKeys := make(map[string]bool)
	for _, entity := range cachedEntities {
		cachedKeys[entity.Key] = true
	}

	// Separate cached and non-cached items
	var cachedItems []string
	var nonCachedItems []string

	for _, item := range allItems {
		if cachedKeys[item] {
			cachedItems = append(cachedItems, item)
		} else {
			nonCachedItems = append(nonCachedItems, item)
		}
	}

	// Return cached items first, then non-cached items
	result := make([]string, 0, len(allItems))
	result = append(result, cachedItems...)
	result = append(result, nonCachedItems...)

	return result
}

func NewTLDRDBCached(nonCachedProvider TLDRProvider) TLDRProvider {
	return &TLDRDBCached{
		provider: nonCachedProvider,
		db:       GetConnection(),
	}
}
