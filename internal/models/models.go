package models

import "time"

type User struct {
	ID          string    `json:"id"`
	Email       string    `json:"email"`
	Password    string    `json:"-"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	DateOfBirth time.Time `json:"date_of_birth"`
	AvatarURL   string    `json:"avatar_url,omitempty"`
	Nickname    string    `json:"nickname,omitempty"`
    Gender      string    `json:"gender"`
	AboutMe     string    `json:"about_me,omitempty"`
	IsPublic    bool      `json:"is_public"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type LoginData struct {
    Email       string    `json:"email"`
	Password    string    `json:"-"`
}

type Follower struct {
    ID         string    `json:"id"`
    FollowerID string    `json:"follower_id"`
    FollowedID string    `json:"followed_id"`
    CreatedAt  time.Time `json:"created_at"`
    Status     string    `json:"status"` // "pending", "accepted"
}

type Session struct {
    SessionID string    `json:"session_id"`
    UserID    string    `json:"user_id"`
    ExpiresAt time.Time `json:"expires_at"`
}

type Post struct {
    ID          string    `json:"id"`
    UserID      string    `json:"user_id"`
    Content     string    `json:"content"`
    ImageURL    string    `json:"image_url,omitempty"`
    Privacy     string    `json:"privacy"` // "public", "private", "almost private"
    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
}

type Comment struct {
    ID        string    `json:"id"`
    PostID    string    `json:"post_id"`
    UserID    string    `json:"user_id"`
    Content   string    `json:"content"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}

type Group struct {
    ID          string    `json:"id"`
    CreatorID   string    `json:"creator_id"`
    Title       string    `json:"title"`
    Description string    `json:"description"`
    CreatedAt   time.Time `json:"created_at"`
}

type GroupMember struct {
    ID        string    `json:"id"`
    GroupID   string    `json:"group_id"`
    UserID    string    `json:"user_id"`
    Role      string    `json:"role"` // "member", "admin"
    JoinedAt  time.Time `json:"joined_at"`
}


type GroupPost struct {
    ID          string    `json:"id"`
    GroupID     string    `json:"group_id"`
    UserID      string    `json:"user_id"`
    Content     string    `json:"content"`
    ImageURL    string    `json:"image_url,omitempty"`
    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
}


type Event struct {
    ID          string    `json:"id"`
    GroupID     string    `json:"group_id"`
    Title       string    `json:"title"`
    Description string    `json:"description"`
    DateTime    time.Time `json:"date_time"`
    CreatedAt   time.Time `json:"created_at"`
}


type EventRSVP struct {
    ID        string    `json:"id"`
    EventID   string    `json:"event_id"`
    UserID    string    `json:"user_id"`
    Response  string    `json:"response"` // "going", "not going"
    CreatedAt time.Time `json:"created_at"`
}

type Notification struct {
    ID        string    `json:"id"`
    UserID    string    `json:"user_id"`
    Message   string    `json:"message"`
    Link      string    `json:"link"`
    Type      string    `json:"type"` // "follow_request", "group_invite", etc.
    Read      bool      `json:"read"`
    CreatedAt time.Time `json:"created_at"`
}


type PrivateMessage struct {
    ID          string    `json:"id"`
    SenderID    string    `json:"sender_id"`
    RecipientID string    `json:"recipient_id"`
    Message     string    `json:"message"`
    SentAt      time.Time `json:"sent_at"`
}


type GroupMessage struct {
    ID        string    `json:"id"`
    GroupID   string    `json:"group_id"`
    SenderID  string    `json:"sender_id"`
    Message   string    `json:"message"`
    SentAt    time.Time `json:"sent_at"`
}

