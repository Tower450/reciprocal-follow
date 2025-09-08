import instaloader

# Initialize
L = instaloader.Instaloader()

# Login
USERNAME = "your_username"  # 👈 Change this
L.login(USERNAME, input("🔐 Enter your Instagram password: "))

# Load profile
profile = instaloader.Profile.from_username(L.context, USERNAME)

# Get followers and followees
followers = set(f.username for f in profile.get_followers())
following = set(f.username for f in profile.get_followees())

# Mutuals and not-following-back
mutuals = followers & following
not_following_back = following - followers
not_followed_by_you = followers - following

# Output
print(f"\n✅ Reciprocal Followers ({len(mutuals)}):")
for user in sorted(mutuals):
    print("  -", user)

print(f"\n❌ Not Following You Back ({len(not_following_back)}):")
for user in sorted(not_following_back):
    print("  -", user)

print(f"\n🤷 You Don't Follow Back ({len(not_followed_by_you)}):")
for user in sorted(not_followed_by_you):
    print("  -", user)

