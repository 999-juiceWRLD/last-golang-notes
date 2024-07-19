package main

import (
	"errors"
)

type userstruct struct {
	name					string
	number 					int
	scheduledForDeletion	bool
}

/*
if the user doesn't exist in the map, return the error not found.
if they exist but aren't scheduled for deletion, return deleted as false with no errors.
if they exist and are scheduled for deletion, return deleted as true with no errors and delete their record from the map.
*/

func deleteIfNecessary(users map[string]userstruct, name string) (deleted bool, err error) {
	someUser, ok := users[name];
	if !ok {
		return false, errors.New("not found");
	}
	isScheduled := someUser.scheduledForDeletion;
	if isScheduled {
		delete(users, name);
		return true, nil;
	} else if !isScheduled {
		return false, nil;
	}
	return;
}

func getCounts(userIDs [] string) map[string]int {
	countsMap := make(map[string]int);
	for _, id := range userIDs {
		count := countsMap[id]
		count++
		countsMap[id] = count
	}
	return countsMap;
}
