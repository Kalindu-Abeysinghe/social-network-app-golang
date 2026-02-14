package db

import (
	"context"
	"fmt"
	"log"
	"math/rand"

	"github.com/Kalindu-Abeysinghe/social-app.git/internal/store"
)

var usernames = []string{
	"shadowhawk", "cyberphoenix", "nightrunner", "stormbreaker",
	"ironwolf", "blazeforge", "frostbyte", "thunderstrike",
	"voidwalker", "crimsonblade", "silverarrow", "darkspectre", "neonghhost", "steelreaper",
	"cosmicninja", "quantumshift", "ravenclaw", "lunarshadow",
	"emberstorm", "ghostrider", "titanforge", "novaflare",
	"obsidianedge", "voltstriker", "mysticviper", "astralfox",
	"zerohour", "infinityedge", "blazehunter", "echowarrior",
	"plasmabolt", "vortexking", "shadowfang", "stormchaser",
	"ironclad", "nightshade", "crimsonwolf", "thunderbolt",
	"frostfang", "nebulastar", "darkmatter", "silverphoenix",
	"cybersamurai", "voidhunter", "blazewolf", "quantumecho",
	"stealthreaper", "cosmicfury", "titanstrike", "shadowreaver",
}

var titles = []string{
	"Why I Switched from React to Vue and Never Looked Back",
	"The Ultimate Guide to Microservices Architecture",
	"10 VS Code Extensions That Changed My Life",
	"How I Built a SaaS App in 30 Days",
	"Docker vs Kubernetes: Which Should You Learn First?",
	"The Dark Side of Agile Development",
	"My Journey from Junior to Senior Developer",
	"Stop Using Nested If Statements: A Better Approach",
	"How I Debug Production Issues at 3 AM",
	"Is TypeScript Really Worth the Hype?",
	"Building a Real-Time Chat App with WebSockets",
	"The Best Programming Books I've Ever Read",
	"Why I Deleted All My Social Media Apps",
	"Understanding Big O Notation Once and For All",
	"How to Negotiate Your Developer Salary",
	"The Fastest Way to Learn a New Programming Language",
	"My Home Office Setup for Maximum Productivity",
	"GraphQL vs REST: A Practical Comparison",
	"How I Passed the AWS Solutions Architect Exam",
	"The Importance of Code Reviews Nobody Talks About",
	"Building Your First Chrome Extension",
	"Why I Use Linux for Development",
	"The Truth About Coding Bootcamps",
	"How to Handle Imposter Syndrome as a Developer",
	"My Favorite Git Commands and Aliases",
	"Building a Portfolio That Actually Gets You Hired",
	"The Rise of AI-Powered Code Assistants",
	"How I Stay Up to Date with Technology",
	"The Best Side Projects for Your Resume",
	"Understanding OAuth 2.0 in Simple Terms",
	"Why Clean Code Matters More Than You Think",
	"My Experience Working at a FAANG Company",
	"How to Build and Deploy a Serverless API",
	"The Problem with Tutorial Hell",
	"Database Indexing Explained with Real Examples",
	"How I Automated My Entire Workflow",
	"The Future of Web Development in 2025",
	"Why I Started Writing Technical Blog Posts",
	"Understanding Redis and When to Use It",
	"How to Contribute to Open Source Projects",
	"The Best Practices for API Design",
	"My Journey Learning Machine Learning",
	"How to Build a Scalable Backend Architecture",
	"The Truth About Remote Work as a Developer",
	"Understanding Design Patterns with Practical Examples",
	"How I Optimized My Website to Load in Under 1 Second",
	"The Importance of Testing Your Code",
	"Building a CI/CD Pipeline from Scratch",
	"How to Land Your First Developer Job",
	"The Best Resources for Learning Data Structures",
}

var contents = []string{
	"The Midnight Chronicles", "Echoes of Tomorrow",
	"Beyond the Horizon", "Whispers in the Dark",
	"The Last Frontier", "Fragments of Time",
	"Shattered Realities", "The Silent Witness",
	"Dreams of Yesterday", "The Infinite Loop",
	"Shadows and Light", "The Forgotten Path",
	"Voices from Beyond", "The Breaking Point",
	"Eternal Wanderers",
}

var tags = []string{
	"golang", "javascript", "python", "typescript", "react",
	"vue", "angular", "nodejs", "docker", "kubernetes",
	"aws", "azure", "gcp", "devops", "microservices",
	"rest", "graphql", "api", "backend", "frontend",
	"fullstack", "webdev", "programming", "coding", "software",
	"development", "tutorial", "guide", "tips", "best-practices",
	"architecture", "design-patterns", "clean-code", "testing", "cicd",
	"git", "linux", "productivity", "career", "learning",
	"database", "sql", "nosql", "mongodb", "postgresql",
	"redis", "machine-learning", "ai", "opensource", "tutorial-hell",
	"imposter-syndrome", "remote-work", "agile", "scrum", "code-review",
}

var sampleComments = []string{
	"Great article! This helped me a lot.",
	"I disagree with this approach. Have you considered using X instead?",
	"Thanks for sharing! Bookmarked for later.",
	"This is exactly what I was looking for!",
	"Could you elaborate more on the third point?",
	"I tried this and it worked perfectly. Thank you!",
	"This changed my entire workflow. Thank you so much!",
	"Not sure I agree with all of this, but good points overall.",
	"Can you provide a code example for this?",
	"This is a common misconception. Actually...",
	"Saved me hours of debugging. You're a lifesaver!",
	"I've been doing this wrong for years. Mind blown!",
	"Great tutorial, but the link in step 3 is broken.",
	"This doesn't work for my use case unfortunately.",
	"Amazing content as always. Keep it up!",
	"I have a question about the performance implications.",
	"This is gold. Why don't more people talk about this?",
	"I wish I had read this before starting my project.",
	"The example code has a small typo in line 12.",
	"This is too complicated. There's a simpler way.",
	"Exactly what the documentation should have explained!",
	"I'm confused about the difference between A and B here.",
	"This approach saved us 40% on AWS costs!",
	"Tried this in production and it caused issues.",
	"Best explanation of this concept I've seen.",
	"Could you make a video tutorial on this?",
	"I implemented this and our team loves it.",
	"This conflicts with what I learned in course X.",
	"Finally someone explains this properly!",
	"This should be required reading for all developers.",
	"I've seen this done better elsewhere, honestly.",
	"Thank you! This solved my exact problem.",
	"Looking forward to part 2 of this series!",
}

func Seed(store store.Storage) {
	ctx := context.Background()

	users := generateUsers(100)

	for _, user := range users {
		if err := store.Users.Create(ctx, user); err != nil {
			log.Fatal("Error creating user: ", err)
			return
		}
	}

	posts := generatePosts(200, users)
	for _, post := range posts {
		if err := store.Posts.Create(ctx, post); err != nil {
			log.Fatal("Error creating post: ", err)
			return
		}
	}

	comments := generateComments(500, users, posts)
	for _, comment := range comments {
		if err := store.Comments.Create(ctx, comment); err != nil {
			log.Fatal("Error creating Comment: ", err)
			return
		}
	}
	log.Println("Seeding complete")
}

func generateUsers(num int) []*store.User {
	users := make([]*store.User, num)
	for i := 0; i < num; i++ {
		username := usernames[i%len(usernames)] + fmt.Sprintf("%d", i)
		users[i] = &store.User{
			Username: username,
			Email:    username + "@email.com",
			Password: "123456",
		}
	}
	return users
}

func generatePosts(num int, users []*store.User) []*store.Post {
	posts := make([]*store.Post, num)
	for i := 0; i < num; i++ {
		user := users[rand.Intn(len(users))]
		posts[i] = &store.Post{
			Content: contents[rand.Intn(len(contents))],
			Title:   titles[rand.Intn(len(titles))],
			UserID:  user.ID,
			Tags: []string{
				tags[rand.Intn(len(tags))],
				tags[rand.Intn(len(tags))],
			},
		}
	}
	return posts
}

func generateComments(num int, users []*store.User, posts []*store.Post) []*store.Comment {
	comments := make([]*store.Comment, num)
	for i := 0; i < num; i++ {
		comments[i] = &store.Comment{
			Content: sampleComments[rand.Intn(len(sampleComments))],
			UserID:  users[rand.Intn(len(users))].ID,
			PostID:  posts[rand.Intn(len(posts))].ID,
		}
	}
	return comments
}
