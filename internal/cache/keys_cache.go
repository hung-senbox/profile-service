package cache

import "profile-service/pkg/constants"

// ==============================
// === Common Cache Key Utils ===
// ==============================

// entity data
func UserCacheKey(userID string) string {
	return constants.CachePrefix + "user:" + userID
}

func StudentCacheKey(studentID string) string {
	return constants.CachePrefix + "student:" + studentID
}

func TeacherCacheKey(teacherID string) string {
	return constants.CachePrefix + "teacher:" + teacherID
}

func StaffCacheKey(staffID string) string {
	return constants.CachePrefix + "staff:" + staffID
}

func ParentCacheKey(parentID string) string {
	return constants.CachePrefix + "parent:" + parentID
}

func ChildCacheKey(childID string) string {
	return constants.CachePrefix + "child:" + childID
}

// code mapping
func UserCodeCacheKey(userID string) string {
	return constants.CachePrefix + "user_code:" + userID
}

func StudentCodeCacheKey(studentID string) string {
	return constants.CachePrefix + "student_code:" + studentID
}

func TeacherCodeCacheKey(teacherID string) string {
	return constants.CachePrefix + "teacher_code:" + teacherID
}

func StaffCodeCacheKey(staffID string) string {
	return constants.CachePrefix + "staff_code:" + staffID
}

func ParentCodeCacheKey(parentID string) string {
	return constants.CachePrefix + "parent_code:" + parentID
}

func ChildCodeCacheKey(childID string) string {
	return constants.CachePrefix + "child_code:" + childID
}

// generic
func GenericCacheKey(prefix, id string) string {
	return constants.CachePrefix + prefix + ":" + id
}
