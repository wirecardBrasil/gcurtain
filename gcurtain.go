package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/go-redis/redis"
)

type GCurtain struct {
	redisClient *redis.Client
}

func (g *GCurtain) Init(uri string) {
	g.redisClient = redis.NewClient(&redis.Options{
		Addr: uri,
	})
}

func (g *GCurtain) IsOpen(feature string, users ...string) bool {
	if existingUser(users...) > 0 {
		return g.isOpenForUser(feature, users...)
	}
	return g.isFeatureOpen(feature)
}

func (g *GCurtain) IsOpenFeature(feature string) bool {
	return g.isFeatureOpen(feature)
}

func (g *GCurtain) IsOpenFeatureUser(feature, user string) bool {
	return g.isOpenForUser(feature, user) || g.isFeatureOpen(feature)
}

func (g *GCurtain) CreateFeature(feature, user string) {
	g.redisClient.SAdd(featureUsers(feature), user)
}

func (g *GCurtain) CreatePercentage(feature, percentage string) {
	num, _ := strconv.Atoi(percentage)
	g.redisClient.Do("SET", featurePercentage(feature), num)
}

func (g *GCurtain) isOpenForUser(feature string, users ...string) bool {
	for _, user := range users {
		if !g.redisClient.SIsMember(featureUsers(feature), user).Val() {
			return false
		}
	}
	return true
}

func (g *GCurtain) getFeaturePercentage(feature string) int {
	featurePercentage := g.redisClient.Get(featurePercentage(feature)).Val()
	if featurePercentage == "" {
		return 0
	}
	num, _ := strconv.Atoi(featurePercentage)
	return num
}

func (g *GCurtain) isFeatureOpen(feature string) bool {
	return randomPercentage() <= g.getFeaturePercentage(feature)
}

func randomPercentage() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(100) + 1
}

func featureUsers(feature string) string {
	return fmt.Sprintf("feature:%s:users", feature)
}

func featurePercentage(feature string) string {
	return fmt.Sprintf("feature:%s:percentage", feature)
}

func existingUser(users ...string) int {
	if len(users) > 0 {
		return 1
	}
	return 0
}

func GetClient(url string) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: url,
	})
}
