package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserBadges struct {
	UserBadgesId                    primitive.ObjectID `bson:"_id,omitempty" json:"user_badges_id"`
	UserBadgesUser                  primitive.ObjectID `bson:"user_badges_user,omitempty" json:"user_badges_user"`
	UserBadgesBadge                 primitive.ObjectID `bson:"user_badges_badge,omitempty" json:"user_badges_badge"`
	UserBadgesRewilding             primitive.ObjectID `bson:"user_badges_rewilding,omitempty" json:"user_badges_rewilding"`
	UserBadgesEvents                primitive.ObjectID `bson:"user_badges_events,omitempty" json:"user_badges_events"`
	UserBadgesEventsParticipantUser primitive.ObjectID `bson:"user_badges_events_participant_user,omitempty" json:"user_badges_events_participant_user"`
	UserBadgesCreatedAt             primitive.DateTime `bson:"user_badges_created_at,omitempty" json:"user_badges_created_at"`
}
