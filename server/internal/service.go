package internal

import (
	"server/entity"
	"server/tools/redis"
	"strconv"
)

type Service struct {
	repo        *Repository
	redisClient *redis.RedisClient
}

func NewService(repo *Repository, redisClient *redis.RedisClient) *Service {
	return &Service{
		repo:        repo,
		redisClient: redisClient,
	}
}

func (s *Service) CreateValue(value *entity.Value) (*entity.Value, error) {
	val, err := s.repo.Create(value)
	if err != nil {
		return nil, err
	}
	s.setToRedisAndPublish(val)
	return val, nil
}

func (s *Service) GetAll() ([]entity.Value, error) {
	return s.repo.GetAll()
}

func (s *Service) GetCurrent() (map[string]string, error) {
	return s.getCurrentFromRedis()
}

func (s *Service) setToRedisAndPublish(value *entity.Value) {
	val := strconv.Itoa(value.Values)
	s.redisClient.HSet("values", val, "Nothing yet!")
	s.redisClient.Publish("insert", val)
}

func (s *Service) getCurrentFromRedis() (map[string]string, error) {
	return s.redisClient.HGetAll("values")
}
