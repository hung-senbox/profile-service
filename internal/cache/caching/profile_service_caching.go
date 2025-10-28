package caching

import (
	"context"
	"profile-service/internal/cache"
)

type CachingService interface {
	SetUserCode(ctx context.Context, userID, code string) error
	SetStudentCode(ctx context.Context, studentID, code string) error
	SetTeacherCode(ctx context.Context, teacherID, code string) error
	SetStaffCode(ctx context.Context, staffID, code string) error
	SetParentCode(ctx context.Context, parentID, code string) error
	SetChildCode(ctx context.Context, childID, code string) error

	InvalidateUserCode(ctx context.Context, userID string) error
	InvalidateStudentCode(ctx context.Context, studentID string) error
	InvalidateTeacherCode(ctx context.Context, teacherID string) error
	InvalidateStaffCode(ctx context.Context, staffID string) error
	InvalidateParentCode(ctx context.Context, parentID string) error
	InvalidateChildCode(ctx context.Context, childID string) error
}

type cachingService struct {
	cache      *cache.RedisCache
	defaultTTL int
}

func NewCachingService(redisCache *cache.RedisCache, ttlSeconds int) CachingService {
	return &cachingService{
		cache:      redisCache,
		defaultTTL: ttlSeconds,
	}
}

func (s *cachingService) setByKey(ctx context.Context, key, code string) error {
	return s.cache.Set(ctx, key, code, s.defaultTTL)
}

func (s *cachingService) deleteByKey(ctx context.Context, key string) error {
	return s.cache.Delete(ctx, key)
}

// ========================
// === SET CODE CACHE ===
// ========================
func (s *cachingService) SetUserCode(ctx context.Context, userID, code string) error {
	key := cache.UserCodeCacheKey(userID)
	return s.setByKey(ctx, key, code)
}

func (s *cachingService) SetStudentCode(ctx context.Context, studentID, code string) error {
	key := cache.StudentCodeCacheKey(studentID)
	return s.setByKey(ctx, key, code)
}

func (s *cachingService) SetTeacherCode(ctx context.Context, teacherID, code string) error {
	key := cache.TeacherCodeCacheKey(teacherID)
	return s.setByKey(ctx, key, code)
}

func (s *cachingService) SetStaffCode(ctx context.Context, staffID, code string) error {
	key := cache.StaffCodeCacheKey(staffID)
	return s.setByKey(ctx, key, code)
}

func (s *cachingService) SetParentCode(ctx context.Context, parentID, code string) error {
	key := cache.ParentCodeCacheKey(parentID)
	return s.setByKey(ctx, key, code)
}

func (s *cachingService) SetChildCode(ctx context.Context, childID, code string) error {
	key := cache.ChildCodeCacheKey(childID)
	return s.setByKey(ctx, key, code)
}

// ========================
// === INVALIDATE CACHE ===
// ========================
func (s *cachingService) InvalidateUserCode(ctx context.Context, userID string) error {
	return s.deleteByKey(ctx, cache.UserCodeCacheKey(userID))
}

func (s *cachingService) InvalidateStudentCode(ctx context.Context, studentID string) error {
	return s.deleteByKey(ctx, cache.StudentCodeCacheKey(studentID))
}

func (s *cachingService) InvalidateTeacherCode(ctx context.Context, teacherID string) error {
	return s.deleteByKey(ctx, cache.TeacherCodeCacheKey(teacherID))
}

func (s *cachingService) InvalidateStaffCode(ctx context.Context, staffID string) error {
	return s.deleteByKey(ctx, cache.StaffCodeCacheKey(staffID))
}

func (s *cachingService) InvalidateParentCode(ctx context.Context, parentID string) error {
	return s.deleteByKey(ctx, cache.ParentCodeCacheKey(parentID))
}

func (s *cachingService) InvalidateChildCode(ctx context.Context, childID string) error {
	return s.deleteByKey(ctx, cache.ChildCodeCacheKey(childID))
}
