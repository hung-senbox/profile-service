package cache

import "profile-service/pkg/constants"

// ==============================
// === Common Cache Key Utils ===
// ==============================

// entity data
func UserCacheKey(userID string) string {
	return constants.ProfileCachePrefix + "user:" + userID
}

func StudentCacheKey(studentID string) string {
	return constants.ProfileCachePrefix + "student:" + studentID
}

func TeacherCacheKey(teacherID string) string {
	return constants.ProfileCachePrefix + "teacher:" + teacherID
}

func StaffCacheKey(staffID string) string {
	return constants.ProfileCachePrefix + "staff:" + staffID
}

func ParentCacheKey(parentID string) string {
	return constants.ProfileCachePrefix + "parent:" + parentID
}

func ChildCacheKey(childID string) string {
	return constants.ProfileCachePrefix + "child:" + childID
}

func TeacherByUserAndOrgCacheKey(userID, orgID string) string {
	return constants.ProfileCachePrefix + "teacher-by:" + userID + ":" + orgID
}

// code mapping
func UserCodeCacheKey(userID string) string {
	return constants.ProfileCachePrefix + "user_code:" + userID
}

func StudentCodeCacheKey(studentID string) string {
	return constants.ProfileCachePrefix + "student_code:" + studentID
}

func TeacherCodeCacheKey(teacherID string) string {
	return constants.ProfileCachePrefix + "teacher_code:" + teacherID
}

func StaffCodeCacheKey(staffID string) string {
	return constants.ProfileCachePrefix + "staff_code:" + staffID
}

func ParentCodeCacheKey(parentID string) string {
	return constants.ProfileCachePrefix + "parent_code:" + parentID
}

func ChildCodeCacheKey(childID string) string {
	return constants.ProfileCachePrefix + "child_code:" + childID
}

// generic
func GenericCacheKey(prefix, id string) string {
	return constants.ProfileCachePrefix + prefix + ":" + id
}
