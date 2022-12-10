package cache

import (
	"context"
	"encoding/json"

	controller "github.com/cal1co/jpi/controllers"
	"github.com/go-redis/redis/v9"
)

var ctx = context.Background()

func LookupAndCache(trieData controller.Output, word string) []controller.Entry {
	cache := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	cachedTranslations, err := cache.Get(ctx, word).Result()
	if err == nil {
		var data []controller.Entry
		json.Unmarshal([]byte(cachedTranslations), &data)
		return data
	} else {
		// If the word is not in the cache, perform dictionary lookup
		translations := controller.Fetch(&trieData.Nodes, trieData.Slice, word)
		bytes, _ := json.Marshal(translations)
		cache.Set(ctx, word, bytes, 0).Err()
		return translations
	}
}
